package treeutil

import (
	"fmt"
	"reflect"
)

type FieldKey struct {
	Primary  string
	ParentId string
	Children string
}

var DefaultFieldKey = FieldKey{
	Primary:  "Id",
	ParentId: "ParentId",
	Children: "Children",
}

func FindRootParentId(dst interface{}, fieldKey *FieldKey) (string, error) {
	if fieldKey == nil {
		fieldKey = &DefaultFieldKey
	}
	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	if dstType.Kind() == reflect.Ptr {
		dstType, dstValue = dstType.Elem(), dstValue.Elem()
	}
	if dstType.Kind() != reflect.Slice {
		return "", fmt.Errorf("%v must be slice kind", "dst")
	}
	if dstValue.Len() == 0 {
		return "", fmt.Errorf("%v must be greater than 0", "dst")
	}
	keyByIdItem := make(map[string]reflect.Value)
	for i := 0; i < dstValue.Len(); i++ {
		dstValueNode := dstValue.Index(i)
		id := ""
		if dstValueNode.Type().Kind() == reflect.Ptr {
			id = dstValueNode.Elem().FieldByName(fieldKey.Primary).String()
		} else {
			id = dstValueNode.FieldByName(fieldKey.Primary).String()
		}
		keyByIdItem[id] = dstValueNode
	}
	for i := 0; i < dstValue.Len(); i++ {
		dstValueNode := dstValue.Index(i)
		parentId := ""
		if dstValueNode.Type().Kind() == reflect.Ptr {
			parentId = dstValueNode.Elem().FieldByName(fieldKey.ParentId).String()
		} else {
			parentId = dstValueNode.FieldByName(fieldKey.ParentId).String()
		}
		if _, ok := keyByIdItem[parentId]; !ok {
			return parentId, nil
		}
	}
	return "", fmt.Errorf("cannot found root id")
}

func Generate(dst interface{}, root string, fieldKey *FieldKey) (interface{}, error) {
	if fieldKey == nil {
		fieldKey = &DefaultFieldKey
	}
	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	if dstType.Kind() == reflect.Ptr {
		dstType, dstValue = dstType.Elem(), dstValue.Elem()
	}
	if dstType.Kind() != reflect.Slice {
		return 0, fmt.Errorf("%v must be slice kind", "dst")
	}
	if dstValue.Len() == 0 {
		return 0, fmt.Errorf("%v must be greater than 0", "dst")
	}
	parentIdMap, rootNodes := make(map[string]reflect.Value, 0), reflect.MakeSlice(dstType, 0, 0)
	for i := 0; i < dstValue.Len(); i++ {
		dstNodeValue := dstValue.Index(i)
		parentId := ""
		if dstNodeValue.Type().Kind() == reflect.Ptr {
			parentId = dstNodeValue.Elem().FieldByName(fieldKey.ParentId).String()
		} else {
			parentId = dstNodeValue.FieldByName(fieldKey.ParentId).String()
		}
		parentIdMapGroup, exist := parentIdMap[parentId]
		if !exist {
			parentIdMapGroup = reflect.MakeSlice(dstType, 0, 0)
		}
		parentIdMapGroup = reflect.Append(parentIdMapGroup, dstNodeValue)
		parentIdMap[parentId] = parentIdMapGroup
		if parentId == root {
			rootNodes = reflect.Append(rootNodes, dstNodeValue)
		}
	}

	result := rootNodes.Interface()
	makeTreeLoop(reflect.ValueOf(&result).Elem().Elem(), parentIdMap, fieldKey.Primary, fieldKey.Children)
	return result, nil
}

func makeTreeLoop(children reflect.Value, parentIdMap map[string]reflect.Value, primaryKey, childrenKey string) {
	for i := 0; i < children.Len(); i++ {
		child := children.Index(i)
		childPrimaryValueS := ""
		if child.Type().Kind() == reflect.Ptr {
			childPrimaryValueS = child.Elem().FieldByName(primaryKey).String()
		} else {
			childPrimaryValueS = child.FieldByName(primaryKey).String()
		}
		childGroup, ok := parentIdMap[childPrimaryValueS]
		if !ok {
			continue
		}
		if child.Type().Kind() == reflect.Ptr {
			childChildrenValue := child.Elem().FieldByName(childrenKey)
			childChildrenValue.Set(childGroup)
			makeTreeLoop(childChildrenValue, parentIdMap, primaryKey, childrenKey)
		} else {
			childChildrenValue := child.FieldByName(childrenKey)
			childChildrenValue.Set(childGroup)
			makeTreeLoop(childChildrenValue, parentIdMap, primaryKey, childrenKey)
		}
	}
}
