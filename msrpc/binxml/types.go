package binxml

import (
	"fmt"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

// An event query or subscription returns multiple events in the result set.
// The result set is a buffer containing one or more variable length
// EVENT_DESCRIPTOR structures, as specified in [MS-DTYP] section 2.3.1.
// Methods that return multiple events always return an array of offsets into
// the buffer for the individual events.
type ResultSet struct {
	// A 32-bit unsigned integer that contains the total size in bytes of this
	// structure, including the header.
	TotalSize uint32 `json:"total_size"`
	// This MUST always be set to 0x00000010.
	HeaderSize uint32 `json:"header_size"`
	// This MUST always be set to 0x00000010.
	EventOffset uint32 `json:"event_offset"`
	// A 32-bit unsigned integer that contains the byte offset from the start
	// of this structure to bookMarkData.
	BookmarkOffset uint32 `json:"bookmark_offset"`
	// Size in bytes of the BinXml data in the eventData field.
	BinXMLSize uint32 `json:"binxml_size"`
	// A byte-array that contains variable length BinXml data.
	Document *Document `json:"document"`
	// Number of subqueryIDs fields that follow. This is 0 if the event
	// is selected by an XPath expression (rather than a structured XML query).
	SubqueriesLength uint32 `json:"subqueries_length"`
	// An array of subquery IDs. Events that are selected using a structured
	// XML query can be selected by one or more subqueries. Each subquery has
	// either an ID specified in the XML element that defines the subquery,
	// or defaults to 0xFFFFFFFF. This list has an entry for each subquery
	// that matches the event. If two subqueries select the event, and both
	// use the same ID, the ID only is listed once.
	Subqueries []uint32 `json:"subqueries"`
	// A byte-array that contains variable length bookmark data.
	Bookmark *Bookmark `json:"bookmark"`
}

func (o *ResultSet) Decode(r *Reader) error {

	r.Begin("result_set")

	r.Read(&o.TotalSize)

	if r.Read(&o.HeaderSize); o.HeaderSize != 0x10 {
		return r.WithErrf("invalid header_size: %d", o.HeaderSize)
	}

	if r.Read(&o.EventOffset); o.EventOffset != 0x10 {
		return r.WithErrf("invalid event_offset: %d", o.EventOffset)
	}

	r.Read(&o.BookmarkOffset)

	r.ReadWithLen(func(r *Reader) error {
		r.SetSize(&o.BinXMLSize) // capture bin_xml_size from read_with_len.
		r.Read(&o.Document)
		return nil
	})

	r.Read(&o.SubqueriesLength)
	o.Subqueries = make([]uint32, o.SubqueriesLength)
	for i := 0; i < int(o.SubqueriesLength); i++ {
		r.Read(&o.Subqueries[i])
	}

	r.ReadWithLen(&o.Bookmark, ExcludeSize)

	return r.Done()
}

// A query can refer to several channels or backup event logs. A subscription
// can refer to several channels. To accurately record the state of a query,
// it is necessary to know where the file cursor (bookmark) is with respect
// to those channels or backup event logs. The bookmark data is encoded as
// follows. Note that all integer fields in this structure MUST be in
// little-endian byte order (that is, least significant byte first).
type Bookmark struct {
	// A 32-bit unsigned integer that contains the total size in bytes of the
	// bookmark, including the header and logRecordNumbers.
	BookmarkSize uint32 `json:"bookmark_size"`
	// A 32-bit unsigned integer, and MUST be set to 0x00000018.
	HeaderSize uint32 `json:"header_size"`
	// A 32-bit unsigned integer that contains the number of channels in the
	// query. This is the number of elements in logRecordNumbers.
	LogRecordNumbersSize uint32 `json:"log_record_numbers_size"`
	// A 32-bit unsigned integer that indicates what channel the current
	// event is from.
	CurrentChannel uint32 `json:"current_channel"`
	// A 32-bit unsigned integer that contains the read direction. 0x00000000
	// indicates chronological order based on time written, and 0x00000001
	// indicates reverse order.
	ReadDirection uint32 `json:"read_direction"`
	// A 32-bit unsigned integer that contains the byte offset from the start
	// of the header to logRecordNumbers.
	RecordIDsOffset uint32 `json:"record_ids_offset"`
	// An array of 64-bit unsigned integers that contain the record numbers
	// for each of the channels or backup event logs. The order of the record
	// numbers MUST match the order of the channels or backup event logs in
	// the query (for example, the first channel in the query corresponds to
	// the first member of the array).
	LogRecordNumbers []uint64 `json:"log_record_numbers"`
}

func (o *Bookmark) Decode(r *Reader) error {

	r.Begin("bookmark")

	r.SetSize(&o.BookmarkSize)

	if r.Read(&o.HeaderSize); o.HeaderSize != 0x18 {
		return r.WithErrf("invalid header_size: %d", o.HeaderSize)
	}

	r.Read(&o.LogRecordNumbersSize)
	r.Read(&o.CurrentChannel)
	r.Read(&o.ReadDirection)
	r.Read(&o.RecordIDsOffset)

	o.LogRecordNumbers = make([]uint64, o.LogRecordNumbersSize)
	for i := 0; i < int(o.LogRecordNumbersSize); i++ {
		r.Read(&o.LogRecordNumbers[i])
	}

	return r.Done()
}

type Document struct {
	Prolog   *ProcInst `json:"prolog,omitempty"`
	Fragment *Fragment `json:"fragment"`
	Misc     *ProcInst `json:"misc,omitempty"`
}

func (o *Document) Decode(r *Reader) error {

	r.Begin("document")

	if r.ReadTag(0x0A, 0x0F) == 0x0A {
		r.Read(&o.Prolog)
		r.ReadTag(0x0F)
	}

	r.Read(&o.Fragment)

	if r.ReadTag(0x00, 0x0A) == 0x0A {
		r.Read(&o.Misc)
		r.ReadTag(0x00)
	}

	return r.Done()
}

type ProcInst struct {
	Target *Name  `json:"target"`
	Data   string `json:"data"`
}

func (o *ProcInst) Decode(r *Reader) error {

	r.Begin("proc_inst")

	r.Read(&o.Target)
	r.ReadTag(0x0B)
	r.ReadUTF16StringWithSize(&o.Data)

	return r.Done()
}

type Name struct {
	Hash uint16 `json:"name_hash"`
	Name string `json:"name"`
}

func (n *Name) String() string { return n.Name }

func (n *Name) Decode(r *Reader) error {

	r.Begin("name")

	r.Read(&n.Hash)
	r.ReadUTF16NStringWithSize(&n.Name)

	if NameHash(n.Name) != n.Hash {
		return r.WithErrf("invalid has for name %q", n.Name)
	}

	return r.Done()
}

func NameHash(s string) uint16 {
	h := uint32(0)
	for _, chr := range []byte(s) {
		h = h*65599 + uint32(chr)
	}
	return uint16(h)
}

type Fragment struct {
	MajorVersion uint8     `json:"major_ver"`
	MinorVersion uint8     `json:"minor_ver"`
	Flags        uint8     `json:"flags"`
	Template     *Template `json:"template,omitempty"`
	Element      *Element  `json:"element,omitempty"`
}

func (o *Fragment) Decode(r *Reader) error {

	r.Begin("fragment")

	r.Read(&o.MajorVersion)
	r.Read(&o.MinorVersion)
	r.Read(&o.Flags)

	switch tag, more := r.ReadTagAndMore(0x0C, 0x01); tag {
	case 0x0C:
		r.Read(&o.Template)
	case 0x01:
		o.Element = &Element{HasAttr: more}
		r.Read(&o.Element)
	}

	return r.Done()
}

type Element struct {
	HasAttr      bool         `json:"has_attr"`
	IsOpenClose  bool         `json:"is_open_close,omitempty"`
	DependencyID int16        `json:"dependency_id"`
	Name         *Name        `json:"name"`
	Attributes   []*Attribute `json:"attributes,omitempty"`
	Content      []*Content   `json:"content,omitempty"`
}

func (o *Element) Decode(r *Reader) error {

	r.Begin("element")

	r.Read(&o.DependencyID)

	r.ReadWithLen(func(r *Reader) error {

		r.Read(&o.Name)

		if o.HasAttr {
			r.ReadWithLen(func(r *Reader) error {
				for {
					attr := &Attribute{}
					_, attr.More = r.ReadTagAndMore(0x06)
					r.Begin(fmt.Sprintf("attr_%d", len(o.Attributes)))
					r.Read(&attr)
					r.Done()
					if o.Attributes = append(o.Attributes, attr); !attr.More {
						break
					}
				}
				return nil
			})
		}

		switch r.ReadTag(0x02, 0x03) {
		case 0x02:
			for {
				c := &Content{}
				if c.ReadTag(r); c.Type == 0x04 {
					break
				}
				r.Begin(fmt.Sprintf("content_%d", len(o.Content)))
				r.Read(&c)
				o.Content = append(o.Content, c)
				r.Done()
			}
		case 0x03:
			o.IsOpenClose = true
		}
		return nil
	})

	return r.Done()
}

type Attribute struct {
	More bool       `json:"has_more"`
	Name *Name      `json:"name"`
	Data []*Content `json:"data"`
}

func (a *Attribute) Decode(r *Reader) error {

	r.Begin("attr")

	r.Read(&a.Name)

	for {
		c := &Content{}
		c.ReadTag(r)
		r.Read(&c)
		if a.Data = append(a.Data, c); !c.More {
			break
		}
	}

	return r.Done()
}

type Substitution struct {
	IsOptional bool   `json:"optional,omitempty"`
	ID         uint16 `json:"id"`
	Type       uint8  `json:"type"`
}

func (s *Substitution) Decode(r *Reader) error {

	r.Begin("substitution")

	r.Read(&s.ID)
	r.Read(&s.Type)

	return r.Done()
}

type Template struct {
	ID       *uuid.UUID      `json:"template_id"`
	Length   uint32          `json:"template_size"`
	Fragment *Fragment       `json:"fragment"`
	Values   *TemplateValues `json:"values"`
}

func (o *Template) Decode(r *Reader) error {

	r.Begin("template")

	// this is pad, actually.
	r.ReadTag(0x00)

	o.ID = new(uuid.UUID)
	r.Read(o.ID)

	r.Read(&o.Length)

	// first read template bytes.
	b := make([]byte, o.Length)
	r.Read(b)

	// decode values.
	r.Read(&o.Values)

	// decode template.
	r.ReadWithBytes(b, func(r *Reader) error {
		r.ReadTag(0x0F)
		r.Read(&o.Fragment)
		r.ReadTag(0x00)
		return nil
	})

	return r.Done()
}

type TemplateValues struct {
	Values       []*Value `json:"values"`
	ValuesLength uint32   `json:"values_length"`
}

func (o *TemplateValues) Decode(r *Reader) error {

	r.Begin("template_values")

	r.Read(&o.ValuesLength)
	o.Values = make([]*Value, o.ValuesLength)

	for i := range o.Values {
		o.Values[i] = &Value{ID: i}
		r.Read(&o.Values[i])
	}

	for _, v := range o.Values {
		if v.Length != 0 {
			b := make([]byte, v.Length)
			r.Read(b)
			if err := v.DecodeValue(b); err != nil {
				return r.WithErr(err)
			}
		}
	}

	return r.Done()
}

var (
	ElementType              uint8 = 0x01
	ValueTextType            uint8 = 0x05
	AttributeType            uint8 = 0x06
	CDATAType                uint8 = 0x07
	CharRefType              uint8 = 0x08
	EntityRefType            uint8 = 0x09
	ProcInstType             uint8 = 0x0B
	NormalSubstitutionType   uint8 = 0x0D
	OptionalSubstitutionType uint8 = 0x0E
)

type Content struct {
	Type         uint8         `json:"type"`
	More         bool          `json:"-"`
	Element      *Element      `json:"element,omitempty"`
	Text         string        `json:"text_value,omitempty"`
	Substitution *Substitution `json:"substitution,omitempty"`
	CharRef      uint16        `json:"char_ref,omitempty"`
	EntityRef    *Name         `json:"entitiy_ref,omitempty"`
	CDATA        string        `json:"cdata,omitempty"`
	ProcInst     *ProcInst     `json:"pidata,omitempty"`
}

func (c *Content) ReadTag(r *Reader) error {
	r.Begin("content_tag")
	c.Type, c.More = r.ReadTagAndMore(0x01, 0x04, 0x05, 0x07, 0x08, 0x09, 0x0B, 0x0D, 0x0E)
	return r.Done()
}

func (c *Content) Decode(r *Reader) error {

	r.Begin("content")

	switch c.Type {
	case 0x01:
		c.Element = &Element{HasAttr: c.More}
		r.Read(&c.Element)
	case 0x05:
		r.ReadTag(0x01)
		r.ReadUTF16StringWithSize(&c.Text)
	case 0x07:
		r.ReadUTF16StringWithSize(&c.CDATA)
	case 0x08:
		r.Read(&c.CharRef)
	case 0x09:
		r.Read(&c.EntityRef)
	case 0x0D:
		r.Read(&c.Substitution)
	case 0x0E:
		c.Substitution = &Substitution{IsOptional: true}
		r.Read(&c.Substitution)
	case 0x0B:
		r.Read(&c.ProcInst)
	}

	return r.Done()
}
