package binxml

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
)

func Unmarshal(b []byte) (*ResultSet, error) {
	rs := &ResultSet{}
	if err := NewReader(nil).ReadWithBytes(b, &rs); err != nil {
		return nil, err
	}

	return rs, nil
}

func RenderXML(b []byte) (string, error) {
	rs, err := Unmarshal(b)
	if err != nil {
		return "", err
	}
	return NewRenderer().Render(rs.Document), nil
}

type Renderer struct {
	w          *strings.Builder
	vs         []*Value
	EscapeText bool
}

func NewRenderer() *Renderer { return &Renderer{w: &strings.Builder{}} }

func (r *Renderer) Render(o interface{}) string {
	r.render(o)
	return r.w.String()
}

func (r *Renderer) withValues(values []*Value) *Renderer {
	return &Renderer{r.w, values, r.EscapeText}
}

func (r *Renderer) render(o interface{}) {

	if reflect.ValueOf(o).IsNil() {
		return
	}

	switch o := o.(type) {
	case *Document:
		r.render(o.Prolog)
		r.render(o.Fragment)
		r.render(o.Misc)
	case *Fragment:
		r.render(o.Template)
		r.render(o.Element)
	case *Template:
		r.withValues(o.Values.Values).render(o.Fragment)
	case *Element:

		var dep []interface{}

		if o.DependencyID != -1 {
			if dep = r.vs[o.DependencyID].Data; r.vs[o.DependencyID].Type == Null {
				return
			}
		}

		for i := 0; dep == nil || i < len(dep); i++ {

			if dep != nil {
				r.vs[o.DependencyID].Data = []interface{}{dep[i]}
			}

			r.w.WriteString("<")
			r.w.WriteString(o.Name.Name)
			for _, attr := range o.Attributes {
				r.render(attr)
			}
			if o.IsOpenClose {
				r.w.WriteString(" />")
				break
			}
			r.w.WriteString(">")
			for _, content := range o.Content {
				r.render(content)
			}

			r.w.WriteString("</")
			r.w.WriteString(o.Name.Name)
			r.w.WriteString(">")

			if dep == nil {
				break
			}
		}

		if dep != nil {
			// set back values.
			r.vs[o.DependencyID].Data = dep
		}

	case *Attribute:

		if len(o.Data) == 0 {
			break
		}

		skip := false
		for _, d := range o.Data {
			if (d.Type == ValueTextType && d.Text == "") ||
				(d.Type == OptionalSubstitutionType && r.vs[d.Substitution.ID].Type == Null) {
				skip = true
				break
			}
		}

		if skip {
			break
		}

		r.w.WriteString(" ")
		r.w.WriteString(o.Name.Name)
		r.w.WriteString("='")
		for _, content := range o.Data {
			r.render(content)
		}
		r.w.WriteString("'")

	case *Content:

		switch o.Type {
		case ElementType:
			r.render(o.Element)
		case ValueTextType:
			r.writeTextValue(o.Text)
		case CharRefType:
			r.w.WriteString("&#")
			r.w.WriteString(strconv.Itoa(int(o.CharRef)))
		case EntityRefType:
			r.w.WriteString("&")
			r.w.WriteString(o.EntityRef.Name)
		case CDATAType:
			r.w.WriteString("<[CDATA[")
			r.w.WriteString(o.CDATA)
			r.w.WriteString("]]")
		case ProcInstType:
			r.render(o.ProcInst)
		case NormalSubstitutionType, OptionalSubstitutionType:
			r.render(r.vs[o.Substitution.ID])
		}

	case *ProcInst:

		r.w.WriteString("<?")
		r.w.WriteString(o.Target.Name)
		r.w.WriteString(" ")
		r.w.WriteString(o.Data)
		r.w.WriteString("?>")

	case *Value:

		for _, value := range o.Data {
			switch o.Type {
			case String, ANSIString:
				r.writeTextValue(value.(string))
			case Int8, Uint8, Int16, Uint16,
				Int32, Uint32, Int64, Uint64, Float32, Float64, Bool, Binary, HexInt32, HexInt64:
				r.w.WriteString(fmt.Sprintf("%v", value))
			case GUID:
				r.w.WriteString("{")
				r.w.WriteString(strings.ToUpper(value.(*uuid.UUID).String()))
				r.w.WriteString("}")
			case FileTime, SysTime:
				r.w.WriteString(value.(time.Time).UTC().Format(time.RFC3339))
			case SID:
				r.w.WriteString(value.(*dtyp.SID).String())
			case BinXML:
				r.render(value)
			default:
				panic("unknown type: " + fmt.Sprintf("%v", o.Type))
			}
		}
	}
}

func (r *Renderer) writeTextValue(s string) {

	if r.EscapeText {
		xml.EscapeText(r.w, []byte(s))
		return
	}

	r.w.WriteString(s)
}
