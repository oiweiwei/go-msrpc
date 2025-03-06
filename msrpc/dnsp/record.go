package dnsp

import (
	"github.com/miekg/dns"

	"github.com/oiweiwei/go-msrpc/msrpc/dnsp/record"
)

// NewRecordFromString creates a new Record from a string data.
func NewRecordFromString(typ uint16, ttl uint32, data string) (*Record, error) {

	record, err := record.NewRecordFromString(typ, ttl, data)
	if err != nil {
		return nil, err
	}

	return &Record{
		DataLength: record.DataLength,
		Type:       uint16(record.Type),
		Flags:      record.Flags,
		Serial:     record.Serial,
		TTLSeconds: record.TTLSeconds,
		Timestamp:  record.Timestamp,
		Buffer:     record.Buffer,
	}, nil
}

// NewRecordFromRR creates a new Record from a dns.RR.
func NewRecordFromRR(rr dns.RR) (*Record, error) {

	record, err := record.NewRecordFromRR(rr)
	if err != nil {
		return nil, err
	}

	return &Record{
		DataLength: record.DataLength,
		Type:       uint16(record.Type),
		Flags:      record.Flags,
		Serial:     record.Serial,
		TTLSeconds: record.TTLSeconds,
		Timestamp:  record.Timestamp,
		Buffer:     record.Buffer,
	}, nil
}
