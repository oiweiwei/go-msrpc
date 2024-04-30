package midl

type Library struct {
	Name  string       `json:"name"`
	Attrs *LibraryAttr `json:"attr,omitempty"`
	Body  LibraryBody  `json:"body,omitempty"`
}

// LibraryBody ...
type LibraryBody struct {
	ComClasses []*ComClass `json:"com_classes,omitempty"`
}

// ComClass structure ...
type ComClass struct {
	Name       string          `json:"name,omitempty"`
	Attrs      *ComClassAttr   `json:"attrs,omitempty"`
	Interfaces []*ComInterface `json:"interfaces,omitempty"`
}

// ComInterface structure ...
type ComInterface struct {
	Name  string            `json:"name,omitempty"`
	Type  *Type             `json:"type,omitempty"`
	Attrs *ComInterfaceAttr `json:"attrs,omitempty"`
}

// DispatchInterface ...
type DispatchInterface struct {
	// The dispatch interface name.
	Name string `json:"name"`
	// The dispatch interface attributes.
	Attrs *DispatchInterfaceAttr `json:"attr,omitempty"`
	// The dispatch interface body.
	Body DispatchInterfaceBody `json:"body,omitempty"`
}

// DispatchInterfaceBody ...
type DispatchInterfaceBody struct {
	// The dispatch interface properties.
	Properties []*Field `json:"properties,omitempty"`
	// The dispatch interfaces methods.
	Methods []*Operation `json:"methods,omitempty"`
}

// ComClassAttr ...
type ComClassAttr struct {
	*InterfaceAttr
	AppObject bool
}

// DispatchInterfaceAttr ...
type DispatchInterfaceAttr struct {
	*InterfaceAttr
}

// LibraryAttr ...
type LibraryAttr struct {
	*InterfaceAttr
}

// ComInterfaceAttr ...
type ComInterfaceAttr struct {
	Default bool
	Source  bool
}
