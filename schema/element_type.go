package schema

type ElementType struct {
	Bool   *BoolType             `json:"bool,omitempty"`
	List   *ListType             `json:"list,omitempty"`
	Map    *MapType              `json:"map,omitempty"`
	Object []ObjectAttributeType `json:"object,omitempty"`
	String *StringType           `json:"string,omitempty"`
}
