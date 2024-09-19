package treeutil

const (
	DefaultPrimaryKey    = "Id"
	DefaultParentIdKey   = "ParentId"
	DefaultChildrenKey   = "Children"
	DefaultDescendantKey = "Descendant"
	DefaultHasChildKey   = "HasChild"
	RootId               = "root"
)

type T struct {
	Id       string `json:"-"`
	ParentId string `json:"-"`
	Label    string `json:"label" xml:"label" yaml:"label"`
	Value    string `json:"value" xml:"value" yaml:"value"`
	Children []T    `json:"children,omitempty" xml:"children" yaml:"children"`
}
