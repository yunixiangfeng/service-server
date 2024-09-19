package treeutil

import (
	"fmt"
	"github.com/attains/go-kit/helper/diff"
	"github.com/attains/go-kit/helper/uniq"
	"gorm.io/gorm"
	"reflect"
	"sort"
)

type Entity struct {
	Orm           *gorm.DB
	IdKey         string
	ParentIdKey   string
	DescendantKey string
	HasChildKey   string
}

func (e *Entity) Add(v interface{}, id, parentId string) error {
	return e.addDescendantRecursively(reflect.ValueOf(v), id, parentId, true)
}

func (e *Entity) Remove(v interface{}, id, parentId string) error {
	return e.removeDescendantRecursively(reflect.ValueOf(v), id, parentId, true)
}

func (e *Entity) addDescendantRecursively(value reflect.Value, id, parentId string, first bool) error {
	prepared, err := e.prepare(value.Elem().Type(), id, parentId, first)
	if err != nil || prepared == nil {
		return err
	}
	value = prepared.(reflect.Value)
	elem := value.Elem()
	vId := elem.FieldByName(e.IdKey).String()
	if vId == "" || vId == RootId {
		return nil
	}

	parentV := reflect.New(elem.Type())
	parent := parentV.Elem()
	result := e.Orm.Where(e.IdKey, parentId).Limit(1).
		Find(parent.Addr().Interface())
	if result.Error != nil {
		return result.Error
	}
	descendant := elem.FieldByName(e.DescendantKey)
	descendant = reflect.AppendSlice(parent.FieldByName(e.DescendantKey), reflect.Append(
		descendant,
		reflect.ValueOf(id).Convert(descendant.Type().Elem()),
	))
	descendant = uniq.Any(descendant)
	sort.Slice(descendant.Interface(), func(i, j int) bool {
		return descendant.Index(i).String() < descendant.Index(j).String()
	})
	updateDataV := reflect.New(elem.Type())
	updateData := updateDataV.Elem()
	updateData.FieldByName(e.HasChildKey).SetBool(true)
	updateData.FieldByName(e.DescendantKey).Set(descendant)
	result = e.Orm.Model(parent.Addr().Interface()).
		Select([]string{e.HasChildKey, e.DescendantKey}).
		Updates(updateData.Interface())
	if result.Error != nil {
		return result.Error
	}
	return e.addDescendantRecursively(value, id, parent.FieldByName(e.ParentIdKey).String(), false)
}

func (e *Entity) removeDescendantRecursively(value reflect.Value, id, parentId string, first bool) error {
	prepared, err := e.prepare(value.Elem().Type(), id, parentId, first)
	if err != nil || prepared == nil {
		return err
	}
	value = prepared.(reflect.Value)
	elem := value.Elem()
	vId := elem.FieldByName(e.IdKey).String()
	if vId == "" || vId == RootId {
		return nil
	}

	parentV := reflect.New(elem.Type())
	parent := parentV.Elem()
	result := e.Orm.Where(e.IdKey, parentId).Limit(1).
		Find(parent.Addr().Interface())
	if result.Error != nil {
		return result.Error
	}
	descendant := elem.FieldByName(e.DescendantKey)
	descendantElemType := descendant.Type().Elem()
	descendant, _ = diff.Any(
		parent.FieldByName(e.DescendantKey),
		reflect.Append(
			descendant,
			reflect.ValueOf(id).Convert(descendantElemType),
		),
	)
	sort.Slice(descendant.Interface(), func(i, j int) bool {
		return descendant.Index(i).String() < descendant.Index(j).String()
	})
	updateDataV := reflect.New(elem.Type())
	updateData := updateDataV.Elem()
	updateData.FieldByName(e.HasChildKey).SetBool(descendant.Len() > 0)
	updateData.FieldByName(e.DescendantKey).Set(descendant)
	result = e.Orm.Model(parent.Addr().Interface()).
		Select([]string{e.HasChildKey, e.DescendantKey}).
		Updates(updateData.Interface())
	if result.Error != nil {
		return result.Error
	}
	return e.removeDescendantRecursively(value, id, parent.FieldByName(e.ParentIdKey).String(), false)
}

func (e *Entity) prepare(typ reflect.Type, id, parentId string, first bool) (interface{}, error) {
	if parentId == "" || parentId == RootId {
		return nil, nil
	}
	selfV := reflect.New(typ)
	self := selfV.Elem()
	if err := e.Orm.Where(e.IdKey, id).Limit(1).Find(self.Addr().Interface()).Error; err != nil {
		return nil, err
	}

	if self.FieldByName(e.ParentIdKey).String() == parentId && !first {
		return nil, nil
	}
	return selfV, nil
}

func CreateRecursiveEntity(orm *gorm.DB, options map[string]interface{}) *Entity {
	helper := &Entity{
		Orm:           orm,
		IdKey:         DefaultPrimaryKey,
		ParentIdKey:   DefaultParentIdKey,
		DescendantKey: DefaultDescendantKey,
		HasChildKey:   DefaultHasChildKey,
	}
	if options == nil {
		return helper
	}
	dstTypes := reflect.TypeOf(helper).Elem()
	dstValues := reflect.ValueOf(helper).Elem()
	fieldNum := dstTypes.NumField()
	for i := 0; i < fieldNum; i++ {
		fieldType, fieldValue := dstTypes.Field(i), dstValues.Field(i)
		var paramFieldV interface{}
		existParam := false
		for fieldN, fieldV := range options {
			if fieldType.Name == fieldN {
				existParam, paramFieldV = true, fieldV
				break
			}
		}
		if !existParam {
			continue
		}
		if fieldType.Name != "Orm" {
			if p, ok := paramFieldV.(string); ok && p != "" {
				fieldValue.Set(reflect.ValueOf(p))
			}
		}
	}

	return helper
}

func FilterDisallowed(data interface{}, disallowed interface{}, options map[string]string) (interface{}, error) {
	defaultOptions := map[string]string{
		"primaryKey":  DefaultPrimaryKey,
		"parentIdKey": DefaultParentIdKey,
	}
	if options != nil {
		for k := range defaultOptions {
			if v, ok := options[k]; ok && v != "" {
				defaultOptions[k] = v
			}
		}
	}
	dataV, disallowedV := reflect.ValueOf(data), reflect.ValueOf(disallowed)
	if dataV.Type().Kind() == reflect.Ptr {
		dataV = dataV.Elem()
	}
	if disallowedV.Type().Kind() == reflect.Ptr {
		disallowedV = disallowedV.Elem()
	}
	if dataV.Type().Kind() != reflect.Slice {
		return nil, fmt.Errorf("data must be %q", reflect.Slice.String())
	}
	if disallowedV.Type().Kind() != reflect.Slice {
		return nil, fmt.Errorf("disallowed must be %q", reflect.Slice.String())
	}
	result := reflect.MakeSlice(dataV.Type(), 0, 0)
	for i := 0; i < dataV.Len(); i++ {
		dataVNode := dataV.Index(i)
		dataVNodePrimaryValue := dataVNode.FieldByName(defaultOptions["primaryKey"])
		dataVNodeParentIdValue := dataVNode.FieldByName(defaultOptions["parentIdKey"])
		isDisallowed := false
		for d := 0; d < disallowedV.Len(); d++ {
			disallowedVNode := disallowedV.Index(d)
			if reflect.DeepEqual(dataVNodePrimaryValue.Interface(), disallowedVNode.Interface()) {
				isDisallowed = true
				break
			}
			if reflect.DeepEqual(dataVNodeParentIdValue.Interface(), disallowedVNode.Interface()) {
				isDisallowed = true
				break
			}
		}
		if !isDisallowed {
			result = reflect.Append(result, dataVNode)
		}
	}
	return result.Interface(), nil
}
