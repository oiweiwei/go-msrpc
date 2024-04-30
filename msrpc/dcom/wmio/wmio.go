package wmio

import (
	"bytes"
	"fmt"
	"strings"
)

func Marshal(obj *Object) ([]byte, error) {

	w := NewCodec(nil)

	b, err := w.EncodeBytesWithSize32(nil, obj)
	if err != nil {
		return nil, err
	}
	return append(cimSignature[:], b...), nil
}

func Unmarshal(b []byte) (*Object, error) {
	if !bytes.Equal(cimSignature[:], b[:4]) {
		return nil, fmt.Errorf("invalid signature")
	}
	obj := &Object{}
	return obj, obj.Decode(NewCodec(b[4:]))
}

type Decoration struct {
	ServerName string `json:"server_name,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
}

func (d *Decoration) Decode(r *Codec) error {
	r.Begin("decoration")
	r.ReadData(&d.ServerName)
	r.ReadData(&d.Namespace)
	return r.Done()
}

func (d *Decoration) Encode(r *Codec) error {
	r.Begin("decoration")
	r.WriteData(d.ServerName)
	r.WriteData(d.Namespace)
	return r.Done()
}

type Property struct {
	Name           string       `json:"name,omitempty"`
	Order          uint16       `json:"order"`
	Offset         uint32       `json:"offset,omitempty"`
	ClassOfOrigin  uint32       `json:"class_of_origin,omitempty"`
	Qualifiers     []*Qualifier `json:"qualifiers,omitempty"`
	Nullable       bool         `json:"nullable,omitempty"`
	InheritDefault bool         `json:"inherit_default,omitempty"`
	Inherited      bool         `json:"inherited,omitempty"`
	Value          Value        `json:"value,omitempty"`
}

func (o *Property) Decode(r *Codec) error {
	r.Begin("property")
	r.ReadRef(&o.Name, "name")
	r.ReadRef(func(r *Codec) error {
		r.Begin("property.*")
		r.ReadData((*uint32)(&o.Value.Type))
		o.Value.Type, o.Inherited = CIMType(uint32(o.Value.Type) & ^uint32(0x4000)), uint32(o.Value.Type)&uint32(0x4000) != 0
		r.ReadData(&o.Order)
		r.ReadData(&o.Offset)
		r.ReadData(&o.ClassOfOrigin)
		r.ReadData(&o.Qualifiers)
		return r.Done()
	}, "data")
	return r.Done()
}

func (o *Property) Encode(r *Codec) error {
	r.Begin("property")
	r.WriteRef(o.Name, "name")
	r.WriteRef(func(r *Codec) error {
		r.Begin("property.*")
		typ := uint32(o.Value.Type)
		if o.Inherited {
			typ = typ | 0x4000
		}
		r.WriteData(typ)
		r.WriteData(o.Order)
		r.WriteData(o.Offset)
		r.WriteData(o.ClassOfOrigin)
		r.WriteDataOnHeap(QualifierSetSize(o.Qualifiers), o.Qualifiers) // use pointer to encode on heap.
		return r.Done()
	}, "data")
	return r.Done()
}

type Object struct {
	Decoration *Decoration  `json:"decoration,omitempty"`
	Class      *ObjectClass `json:"class,omitempty"`
	Instance   *Instance    `json:"instance,omitempty"`
}

func (o *Object) Method(n string) (*Object, *Object, error) {

	if o.Class == nil {
		return nil, nil, fmt.Errorf("class does not exist")
	}

	for i := range o.Class.CurrentClassMethods.Methods {
		if m := o.Class.CurrentClassMethods.Methods[i]; m.Name == n {
			in, out := m.InputSignature, m.OutputSignature
			return &in, &out, nil
		}
	}

	return nil, nil, fmt.Errorf("method does not exist")
}

type Values map[string]any

func (o *Object) New(values Values) (*Object, error) {

	if values == nil {
		values = make(map[string]any)
	}

	var cls Class

	if o.Class != nil {
		cls = o.Class.CurrentClass
	} else if o.Instance != nil {
		cls = o.Instance.CurrentClass
	}

	if cls.Name == "" {
		return nil, fmt.Errorf("class does not exist")
	}

	inst := &Instance{
		CurrentClass: cls,
		ClassName:    cls.Name,
		Properties:   make([]*Property, len(cls.Properties)),
	}

	for i := range cls.Properties {
		value, ok := values[inst.CurrentClass.Properties[i].Name]
		if !ok {
			inst.Properties[i] = &Property{InheritDefault: true, Value: Value{Type: cls.Properties[i].Value.Type}}
			continue
		}
		if value == nil {
			inst.Properties[i] = &Property{Nullable: true, Value: Value{Type: cls.Properties[i].Value.Type}}
			continue
		}

		if value, ok := value.(Value); ok {
			inst.Properties[i] = &Property{Value: value}
			continue
		}

		typ, err := ValueType(value)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", inst.CurrentClass.Properties[i].Name, err)
		}

		inst.Properties[i] = &Property{Value: Value{Type: typ, Value: value}}
	}

	return &Object{Instance: inst}, nil
}

func (o *Object) Decode(r *Codec) error {
	r.Begin("object")
	sz := uint32(0)
	r.ReadData(&sz)
	r.DecodeWithSize32(sz, func(r *Codec) error {
		r.Begin("cim")
		var flags uint8
		r.ReadData(&flags)
		if flags & ^uint8(0x07) != 0 {
			return fmt.Errorf("invalid object flags 0x%02x", flags)
		}
		if flags&0x04 != 0 {
			dcr := Decoration{}
			r.ReadData(&dcr)
			o.Decoration = &dcr
		}
		if flags&0x01 != 0 {
			cls := ObjectClass{}
			r.ReadData(&cls)
			o.Class = &cls
		}
		if flags&0x02 != 0 {
			ins := Instance{}
			r.ReadData(&ins)
			o.Instance = &ins
		}
		return r.Done()
	})
	return r.Done()
}

func (o *Object) Encode(r *Codec) error {
	r.Begin("object")

	var flags uint8
	if o.Decoration != nil {
		flags |= 0x04
	}
	if o.Instance != nil {
		flags |= 0x02
	}
	if o.Class != nil {
		flags |= 0x01
	}

	sz := uint32(0)
	b, _ := r.EncodeBytesWithSize32(&sz, func(r *Codec) error {
		r.Begin("cim")
		r.WriteData(flags)
		if o.Decoration != nil {
			r.WriteData(o.Decoration)
		}
		if o.Class != nil {
			r.WriteData(o.Class)
		}
		if o.Instance != nil {
			r.WriteData(o.Instance)
		}
		return r.Done()
	})

	r.WriteData(sz)
	r.WriteData(b)

	return r.Done()
}

type Instance struct {
	CurrentClass Class        `json:"current_class,omitempty"`
	Flags        uint8        `json:"flags,omitempty"`
	ClassName    string       `json:"class_name,omitempty"`
	Properties   []*Property  `json:"properties,omitempty"`
	Qualifiers   []*Qualifier `json:"qualifiers,omitempty"`
}

func (o *Instance) Decode(r *Codec) error {

	r.Begin("instance")

	r.ReadData(&o.CurrentClass)
	r.DecodeWithLength32(func(r *Codec) error {

		r.Begin("body")

		r.refs.Push()
		defer r.refs.Pop()

		r.ReadData(&o.Flags)
		r.ReadRef(&o.ClassName, "class_name")

		pSz := len(o.CurrentClass.Properties)
		o.Properties = make([]*Property, pSz)
		for i := range o.Properties {
			cp := o.CurrentClass.Properties[i]
			o.Properties[i] = &Property{Name: cp.Name, Value: cp.Value}
		}

		r.DecodeWithSize32(o.CurrentClass.NDValueSize, func(r *Codec) error {
			r.Begin("nd_value_table")
			cur := uint8(0)
			for i := range o.Properties {
				if i%4 == 0 {
					r.ReadData(&cur)
				}
				bit := (cur >> ((i % 4) * 2)) & 0x03
				if bit&0x01 != 0 {
					o.Properties[i].Nullable = true
				}
				if bit&0x02 != 0 {
					o.Properties[i].InheritDefault = true
				}
			}

			vT := make([]byte, r.Len())
			r.ReadData(vT)

			for i := range o.Properties {
				if !o.Properties[i].Nullable && !o.Properties[i].InheritDefault {
					r.DecodeWithBytes(vT[o.CurrentClass.Properties[i].Offset:], &o.Properties[i].Value)
				}
			}

			return r.Done()
		})

		r.ReadData(&o.Qualifiers)

		flags := uint8(0)
		r.ReadData(&flags)

		switch flags {
		case 0x01:
		case 0x02:
			for i := range o.Properties {
				r.ReadData(&o.Properties[i].Qualifiers)
			}
		default:
			return fmt.Errorf("instace_property_qualifier_set: invalid flags 0x%02x", flags)
		}

		// instance heap.
		r.ReadHeap()

		r.DecodeHeap()

		return r.Done()
	})

	return r.Done()
}

func (o *Instance) Encode(r *Codec) error {

	r.Begin("instance")

	r.WriteData(&o.CurrentClass)
	r.EncodeWithLength32(func(r *Codec) error {

		r.Begin("body")

		r.refs.Push()
		defer r.refs.Pop()

		r.WriteData(o.Flags)
		r.WriteRef(o.ClassName, "class_name")

		cur := uint8(0)
		for i := range o.CurrentClass.Properties { /* iterate over class set of properties */
			if o.Properties[i].Nullable {
				cur |= (0x01 << ((i % 4) * 2))
			}
			if o.Properties[i].InheritDefault {
				cur |= (0x02 << ((i % 4) * 2))
			}
			if (i+1)%4 == 0 {
				_, cur = r.WriteData(cur), 0
			}
		}
		if len(o.Properties)%4 != 0 {
			_, cur = r.WriteData(cur), 0
		}

		vTSz := 0
		for i := range o.CurrentClass.Properties { /* compute value table size */
			vTSz += o.CurrentClass.Properties[i].Value.EncodeValueSize()
		}

		vT := make([]byte, vTSz)

		for i := range o.CurrentClass.Properties {

			var b []byte

			if o.Properties[i].InheritDefault || o.Properties[i].Nullable {
				b = bytes.Repeat([]byte{0xFF}, o.CurrentClass.Properties[i].Value.EncodeValueSize())
			} else {
				b, _ = r.EncodeBytesWithSize32(nil, &o.Properties[i].Value)
			}
			copy(vT[o.CurrentClass.Properties[i].Offset:], b)
		}

		r.WriteData(vT)

		r.WriteData(o.Qualifiers)

		flags := uint8(0x01)
		for i := range o.Properties {
			if len(o.Properties[i].Qualifiers) > 0 {
				flags = 0x02
				break
			}
		}

		r.WriteData(flags)

		switch flags {
		case 0x01:
		case 0x02:
			for i := range o.Properties {
				r.WriteData(o.Properties[i].Qualifiers)
			}
		}

		// instance heap.
		r.WriteHeap()

		return r.Done()
	})

	return r.Done()
}

type ObjectClass struct {
	ParentClass         Class   `json:"parent_class,omitempty"`
	ParentClassMethods  Methods `json:"parent_class_methods,omitempty"`
	CurrentClass        Class   `json:"current_class,omitempty"`
	CurrentClassMethods Methods `json:"current_class_methods,omitempty"`
}

func (o *ObjectClass) Decode(r *Codec) error {
	r.Begin("object_class")
	r.ReadData(&o.ParentClass)
	r.ReadData(&o.ParentClassMethods)
	r.ReadData(&o.CurrentClass)
	r.ReadData(&o.CurrentClassMethods)
	return r.Done()
}

func (o *ObjectClass) Encode(r *Codec) error {
	r.Begin("object_class")
	r.WriteData(&o.ParentClass)
	r.WriteData(&o.ParentClassMethods)
	r.WriteData(&o.CurrentClass)
	r.WriteData(&o.CurrentClassMethods)
	return r.Done()
}

type Method struct {
	Name            string       `json:"name,omitempty"`
	Flags           uint8        `json:"flags,omitempty"`
	Origin          uint32       `json:"origin,omitempty"`
	Qualifiers      []*Qualifier `json:"qualifiers,omitempty"`
	InputSignature  Object       `json:"input_signature,omitempty"`
	OutputSignature Object       `json:"output_signature,omitempty"`
}

func (o *Method) Decode(r *Codec) error {

	r.Begin("method")

	r.ReadRef(&o.Name, "name")
	r.ReadData(&o.Flags)
	_pad := [3]byte{}
	r.ReadData(&_pad)
	r.ReadData(&o.Origin)
	r.ReadRef(&o.Qualifiers, "qualifiers")
	r.ReadRef(&o.InputSignature, "input_signature")
	r.ReadRef(&o.OutputSignature, "output_signature")

	return r.Done()
}

func (o *Method) Encode(r *Codec) error {

	r.Begin("method")

	r.WriteRef(o.Name, "name")
	r.WriteData(o.Flags)
	_pad := [3]byte{}
	r.WriteData(_pad)
	r.WriteData(o.Origin)
	r.WriteRef(func(r *Codec) error {
		return r.WriteDataOnHeap(QualifierSetSize(o.Qualifiers), o.Qualifiers)
	}, "qualifiers")
	r.WriteRef(&o.InputSignature, "input_signature")
	r.WriteRef(&o.OutputSignature, "output_signature")

	return r.Done()
}

type Methods struct {
	Methods []*Method `json:"methods,omitempty"`
}

func (o *Methods) Decode(r *Codec) error {
	return r.DecodeWithLength32(func(r *Codec) error {
		r.Begin("methods")
		r.refs.Push()
		defer r.refs.Pop()
		mSz := uint16(0)
		r.ReadData(&mSz)
		_pad := uint16(0)
		r.ReadData(&_pad)
		for i := uint16(0); i < mSz; i++ {
			m := Method{}
			r.ReadData(&m)
			o.Methods = append(o.Methods, &m)
		}
		// methods heap.
		r.ReadHeap()
		r.DecodeHeap()
		return r.Done()
	})
}

func (o *Methods) Encode(r *Codec) error {
	return r.EncodeWithLength32(func(r *Codec) error {
		r.Begin("methods")
		r.refs.Push()
		defer r.refs.Pop()
		r.WriteData(uint16(len(o.Methods)))
		r.WriteData(uint16(0))
		for i := range o.Methods {
			r.WriteData(o.Methods[i])
		}
		// method heap.
		r.WriteHeap()
		return r.Done()
	})
}

type Class struct {
	Raw         []byte       `json:"-"`
	Name        string       `json:"name,omitempty"`
	NDValueSize uint32       `json:"-"`
	Derivation  []string     `json:"derivation,omitempty"`
	Qualifiers  []*Qualifier `json:"qualifiers,omitempty"`
	Properties  []*Property  `json:"properties,omitempty"`
}

func (o *Class) Decode(r *Codec) error {
	return r.DecodeWithLength32(func(r *Codec) error {
		// save raw encoding of the class.
		o.Raw = r.buf.Bytes()
		r.Begin("class")
		r.refs.Push()
		defer r.refs.Pop()
		_pad := uint8(0)
		r.ReadData(&_pad)
		r.ReadRef(&o.Name, "name")
		r.ReadData(&o.NDValueSize)
		r.DecodeWithLength32(func(r *Codec) error {
			r.Begin("derivation")
			for i := 0; r.Len() > 0; i++ {
				o.Derivation = append(o.Derivation, "")
				r.ReadData(&o.Derivation[i])
				l := uint32(0)
				r.ReadData(&l)
			}
			return r.Done()
		})

		r.ReadData(&o.Qualifiers)

		pSz := uint32(0)
		r.ReadData(&pSz)

		// skip over propery lookup table.
		pB := make([]byte, pSz*8)
		r.ReadData(pB)

		// skip over nd_value_table.
		ndB := make([]byte, o.NDValueSize)
		r.ReadData(ndB)

		// read class heap.
		r.ReadHeap()

		// decode the properties after heap.
		r.DecodeWithBytes(pB, func(r *Codec) error {
			o.Properties = make([]*Property, pSz)
			for i := uint32(0); i < pSz; i++ {
				p := Property{}
				r.ReadData(&p)
				o.Properties[i] = &p
			}
			return nil
		})

		// class heap.
		r.DecodeHeap()

		r.DecodeWithBytes(ndB, func(r *Codec) error {
			r.Begin("nd_value_table")
			cur := uint8(0)
			for i := range o.Properties {
				if i%4 == 0 {
					r.ReadData(&cur)
				}
				bit := (cur >> ((i % 4) * 2)) & 0x03
				if bit&0x01 != 0 {
					o.Properties[i].Nullable = true
				}
				if bit&0x02 != 0 {
					o.Properties[i].InheritDefault = true
				}
			}

			vT := make([]byte, r.Len())
			r.ReadData(vT)

			for i := range o.Properties {
				// for a specific property, the value here is relevant only
				// if the corresponding NDTable bits for that property are both
				// not set, that is, 0.
				if !o.Properties[i].Nullable && !o.Properties[i].InheritDefault {
					r.DecodeWithBytes(vT[o.Properties[i].Offset:], &o.Properties[i].Value)
				}
			}

			return r.Done()
		})

		// class heap.
		r.DecodeHeap()

		return r.Done()
	})
}

func (o *Class) Encode(r *Codec) error {
	return r.EncodeWithLength32(func(r *Codec) error {
		r.Begin("class")
		// if o.Raw != nil {
		//	r.WriteData(o.Raw)
		//	return r.Done()
		// }
		r.refs.Push()
		defer r.refs.Pop()
		r.WriteData(uint8(0))
		r.WriteRef(o.Name, "name")
		ndSz := len(o.Properties) / 4
		if len(o.Properties)%4 != 0 {
			ndSz++
		}
		vTSz := 0
		for i := range o.Properties {
			vTSz += o.Properties[i].Value.EncodeValueSize()
		}
		r.WriteData((uint32)(ndSz + vTSz))
		r.EncodeWithLength32(func(r *Codec) error {
			r.Begin("derivation")
			for i := range o.Derivation {
				r.WriteData(o.Derivation[i])
				if IsASCII(o.Derivation[i]) {
					r.WriteData(uint32(len(o.Derivation[i]) + 1 /* flags */ + 1 /* null-char */))
				} else {
					r.WriteData(uint32(len(o.Derivation[i])*2 + 1 /* flags */ + 2 /* null-char */))
				}
			}
			return r.Done()
		})

		r.WriteData(o.Qualifiers)

		noOft := true
		for i := range o.Properties { /* determine if all offsets are zero */
			if o.Properties[i].Offset != 0 {
				noOft = false
				break
			}
		}

		if noOft /* fix offsets in value table */ {
			oft := 0
			for i := range o.Properties {
				o.Properties[i].Offset, oft = uint32(oft), oft+o.Properties[i].Value.EncodeValueSize()
			}
		}

		r.WriteData(uint32(len(o.Properties)))

		for i := range o.Properties {
			r.WriteData(o.Properties[i])
		}

		r.Begin("nd_value_table")
		cur := uint8(0)
		for i := range o.Properties {
			if o.Properties[i].Nullable {
				cur |= (0x01 << ((i % 4) * 2))
			}
			if o.Properties[i].InheritDefault {
				cur |= (0x02 << ((i % 4) * 2))
			}
			if (i+1)%4 == 0 {
				_, cur = r.WriteData(cur), 0
			}
		}
		if len(o.Properties)%4 != 0 {
			_, cur = r.WriteData(cur), 0
		}

		vT := make([]byte, vTSz)

		for i := range o.Properties {

			var b []byte

			if o.Properties[i].InheritDefault || o.Properties[i].Nullable {
				b = make([]byte, o.Properties[i].Value.EncodeValueSize())
				for i := range b {
					b[i] = 0xFF
				}
			} else {
				b, _ = r.EncodeBytesWithSize32(nil, &o.Properties[i].Value)
			}
			copy(vT[o.Properties[i].Offset:], b)
		}
		r.WriteData(vT)
		r.Done()

		// class heap.
		r.WriteHeap()

		return r.Done()
	})
}

type Flavor uint8

var (
	PropagateToInstance     Flavor = 0x01
	PropagateToDerivedClass Flavor = 0x02 // ToSubclass
	NotOverridable          Flavor = 0x10 // EnableOverride
	OriginPropagated        Flavor = 0x20
	OriginSystem            Flavor = 0x40
	Amended                 Flavor = 0x80 // Translatable
)

func (o Flavor) String() string {

	ret := []string{}
	if o&PropagateToDerivedClass != 0 {
		ret = append(ret, "tosubclass")
	} else {
		ret = append(ret, "restricted")
	}

	if o&NotOverridable != 0 {
		ret = append(ret, "disableoverride")
	} else {
		ret = append(ret, "enableoverride")
	}

	if o&Amended != 0 {
		ret = append(ret, "translatable")
	}

	return strings.Join(ret, " ")
}

type Qualifier struct {
	Name   string `json:"name,omitempty"`
	Flavor Flavor `json:"flavour,omitempty"`
	Value  Value  `json:"value,omitempty"`
}

func QualifierSetSize(qs []*Qualifier) int {
	/* actual size of the qualifiers.
	qSz := 4 // encoded-length
	for i := range qs {
		qSz += 4 + 1 + 4 + qs[i].Value.EncodeValueSize()
	}
	*/
	return 4 + 20*len(qs)
}

func (o *Qualifier) String() string {
	if o.Value.Type == Bool {
		return fmt.Sprintf("%s", o.Name)
	}
	if o.Value.Type.IsArray() {
		return fmt.Sprintf("%s {%s}", o.Name, o.Value)
	}
	return fmt.Sprintf("%s(%s)", o.Name, o.Value)
}

func (o *Qualifier) Decode(r *Codec) error {
	r.Begin("qualifier")
	r.ReadRef(&o.Name, "name")
	r.ReadData((*uint8)(&o.Flavor))
	r.ReadData((*uint32)(&o.Value.Type))
	o.Value.Type = CIMType(uint32(o.Value.Type) & ^Inherited) /* always unset the inherited bit */
	r.ReadData(&o.Value)
	return r.Done()
}

func (o *Qualifier) Encode(r *Codec) error {
	r.Begin("qualifier")
	r.WriteRef(o.Name, "name")
	r.WriteData((uint8)(o.Flavor))
	r.WriteData(uint32(o.Value.Type) & ^Inherited) /* always unset the inherited bit */
	r.WriteData(&o.Value)
	return r.Done()
}

type ObjectFlag uint8

const (
	ObjectFlagClass      ObjectFlag = 0x01
	ObjectFlagInstance   ObjectFlag = 0x02
	ObjectFlagDecoration ObjectFlag = 0x04
)

var cimDictionary = [...]string{
	0:  "\"",
	1:  "key",
	2:  "\"\"",
	3:  "read",
	4:  "write",
	5:  "volatile",
	6:  "provider",
	7:  "dynamic",
	8:  "cimwin32",
	9:  "DWORD",
	10: "CIMTYPE",
}

var rcimDictionary = map[string]uint32{}

func init() {
	for i, v := range cimDictionary {
		rcimDictionary[v] = uint32(i) | uint32(1<<31)
	}
}

func LookupDictionary(ref uint32) string {
	iref := int(ref & ^uint32(1<<31))
	if iref >= len(cimDictionary) {
		return ""
	}
	return cimDictionary[iref]
}

func ReverseLookupDictionary(s string) uint32 {
	ref, _ := rcimDictionary[s]
	return ref
}

var cimSignature = [...]byte{0x78, 0x56, 0x34, 0x12}
