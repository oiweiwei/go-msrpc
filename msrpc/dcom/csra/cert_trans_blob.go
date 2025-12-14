package csra

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	"github.com/oiweiwei/go-msrpc/ndr"
)

type ColumnType uint32

var (
	ColumnTypeSignedInteger ColumnType = 0x00000001
	ColumnTypeTimestamp     ColumnType = 0x00000002
	ColumnTypeBinary        ColumnType = 0x00000003
	ColumnTypeString        ColumnType = 0x00000004
)

type ColumnFlag uint32

var (
	ColumnFlagIndexed ColumnFlag = 0x0001
)

type Column struct {
	Type        ColumnType `json:"type"`
	Flags       ColumnFlag `json:"flags"`
	Index       uint32     `json:"index"`
	MaxLength   uint32     `json:"max_length,omitempty"`
	Name        string     `json:"name,omitempty"`
	DisplayName string     `json:"display_name,omitempty"`
	RawData     []byte     `json:"raw_data,omitempty"`
}

func (c *Column) Value() (any, error) {
	return parseColumnValue(c.Type, c.RawData)
}

func parseColumnValue(colType ColumnType, raw []byte) (any, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	switch colType {
	case ColumnTypeSignedInteger:
		return int32(binary.LittleEndian.Uint32(raw)), nil
	case ColumnTypeTimestamp:
		var filetime dtyp.Filetime
		if err := ndr.Unmarshal(raw, &filetime); err != nil {
			return nil, err
		}
		return filetime.AsTime(), nil
	case ColumnTypeBinary:
		return raw, nil
	case ColumnTypeString:
		var s string
		if err := ndr.ReadUTF16NStringNoSize(context.Background(), ndr.NDR20(raw, ndr.Opaque), &s); err != nil {
			return nil, err
		}
		return s, nil
	default:
		return nil, nil
	}
}

func (o *CertTransportBlob) Columns(count int) ([]*Column, error) {

	var ret []*Column

	for i, r := 0, ndr.NDR20(o.Buffer, ndr.Opaque); i < count; i++ {

		var col CertTransDBColumn

		if err := col.UnmarshalNDR(context.Background(), r); err != nil {
			return nil, err
		}

		var name, displayName string

		if err := ndr.ReadUTF16NStringNoSize(
			context.Background(),
			ndr.NDR20(o.Buffer[col.NameOffset:]), &name); err != nil {
			return nil, err
		}

		if err := ndr.ReadUTF16NStringNoSize(
			context.Background(),
			ndr.NDR20(o.Buffer[col.DisplayNameOffset:]), &displayName); err != nil {
			return nil, err
		}

		ret = append(ret, &Column{
			Flags:       ColumnFlag((col.Type >> 16) & 0xFFFF),
			Type:        ColumnType(col.Type & 0xFFFF),
			Index:       col.Index,
			MaxLength:   col.MaxLength,
			Name:        name,
			DisplayName: displayName,
		})
	}

	return ret, nil
}

type Row struct {
	ID      uint32    `json:"id"`
	Columns []*Column `json:"columns"`
}

func (o *CertTransportBlob) Rows() ([]*Row, bool, error) {

	b := o.Buffer

	var rows []*Row
	var more = true

	for len(b) >= 8 {

		var resultRow CertTransDBResultRow

		r := ndr.NDR20(b, ndr.Opaque)

		if err := r.Unmarshal(context.Background(), &resultRow); err != nil {
			return nil, false, fmt.Errorf("unmarshal result row: %w", err)
		}

		if resultRow.RowID == ^resultRow.ColumnsCount {
			more = false
			break
		}

		row := &Row{
			ID: resultRow.RowID,
		}

		for i := uint32(0); i < resultRow.ColumnsCount; i++ {
			var resultCol CertTransDBResultColumn
			if err := r.Unmarshal(context.Background(), &resultCol); err != nil {
				return nil, false, fmt.Errorf("unmarshal result column: %d: %w", i, err)
			}

			row.Columns = append(row.Columns, &Column{
				Type:    ColumnType(resultCol.Type & 0xFFFF),
				Flags:   ColumnFlag((resultCol.Type >> 16) & 0xFFFF),
				Index:   resultCol.Index,
				RawData: b[resultCol.ValueOffset : resultCol.ValueOffset+resultCol.ValueLength],
			})
		}

		rows, b = append(rows, row), b[resultRow.RowLength:]
	}

	return rows, more, nil
}
