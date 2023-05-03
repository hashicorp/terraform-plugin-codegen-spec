package schema

type ObjectAttributeType struct {
	Name string `json:"name"`

	Bool   *BoolType             `json:"bool,omitempty"`
	List   *ListType             `json:"list,omitempty"`
	Map    *MapType              `json:"map,omitempty"`
	Object []ObjectAttributeType `json:"object,omitempty"`
	String *StringType           `json:"string,omitempty"`
}
