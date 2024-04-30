package doc

import "strings"

type TypeDoc struct {
	Ref    string      `json:"ref"`
	Name   string      `json:"name"`
	Doc    []string    `json:"doc"`
	Fields []*FieldDoc `json:"fields"`
}

var (
	ORPCThis = &FieldDoc{
		Name: "This",
		Doc:  []string{"This: ORPCTHIS structure that is used to send ORPC extension data to the server."},
	}

	ORPCThat = &FieldDoc{
		Name: "That",
		Doc:  []string{"That: ORPCTHAT structure that is used to return ORPC extension data to the client."},
	}
)

func (td *TypeDoc) ObjectField(n string) (*FieldDoc, bool) {

	if n == "This" {
		return ORPCThis, true
	}

	if n == "That" {
		return ORPCThat, true
	}

	if n == "ReturnValue" {
		n = "Return Values"
	}

	return td.Field(n)
}

func (td *TypeDoc) Field(n string) (*FieldDoc, bool) {
	if td == nil {
		return nil, false
	}

	for _, field := range td.Fields {
		if field.Name == n || field.Name == strings.TrimPrefix(n, "__") {
			return field, true
		}
	}
	return nil, false
}

func (td *TypeDoc) AddField(field, doc string) {
	fd := &FieldDoc{Name: field}
	if doc != "" {
		fd.AddDoc(doc)
	}
	td.Fields = append(td.Fields, fd)
}

func (td *TypeDoc) AddDoc(doc string) {
	if doc == "" || doc == "msdn link" || doc == "Server Processing Rules" {
		return
	}
	if len(td.Fields) > 0 {
		td.Fields[len(td.Fields)-1].AddDoc(doc)
		return
	}
	td.Doc = append(td.Doc, doc)
}

type FieldDoc struct {
	Name string   `json:"name"`
	Doc  []string `json:"doc"`
}

func (td *FieldDoc) AddDoc(doc string) {
	if doc == "" {
		return
	}
	td.Doc = append(td.Doc, doc)
}
