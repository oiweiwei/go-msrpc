package epm

import (
	"bytes"
	"context"
	"encoding/binary"
	"testing"

	"github.com/oiweiwei/go-msrpc/ndr"
)

func le(v ...uint32) []byte {
	b := make([]byte, 4*len(v))
	for i := range v {
		binary.LittleEndian.PutUint32(b[4*i:], v[i])
	}
	return b
}

// tower returns the NDR of twr_t: max count, tower_length, octets padded to 4.
func tower(octets []byte) []byte {
	b := le(uint32(len(octets)), uint32(len(octets)))
	b = append(b, octets...)
	for len(b)%4 != 0 {
		b = append(b, 0)
	}
	return b
}

// mapResponse builds an ept_map response: zero handle, num_towers, towers[max_towers] with
// referent ids ids and the deferred data of the pointers that are not aliases, status.
func mapResponse(ids []uint32, data [][]byte, status uint32) []byte {
	var b bytes.Buffer
	b.Write(make([]byte, 20))
	n := uint32(len(ids))
	b.Write(le(n, n, 0, n))
	b.Write(le(ids...))
	for _, d := range data {
		b.Write(tower(d))
	}
	b.Write(le(status))
	return b.Bytes()
}

func TestMapResponseTowersFullPointerAliasing(t *testing.T) {
	t1 := []byte{1, 0, 0, 0xA}
	t2 := []byte{2, 0, 0, 0xB, 5, 6}

	// distinct ids: every tower carries data.
	var r MapResponse
	if err := r.UnmarshalNDR(context.Background(), ndr.NDR20(mapResponse([]uint32{0x20000, 0x20004}, [][]byte{t1, t2}, 0))); err != nil {
		t.Fatal(err)
	}
	if len(r.Towers) != 2 || !bytes.Equal(r.Towers[0].TowerOctetString, t1) || !bytes.Equal(r.Towers[1].TowerOctetString, t2) {
		t.Fatalf("distinct: %+v", r.Towers)
	}

	// repeated id: [ptr] twr_p_t still aliases, the second pointer carries no data and the status after it is read.
	r = MapResponse{}
	if err := r.UnmarshalNDR(context.Background(), ndr.NDR20(mapResponse([]uint32{0x20000, 0x20000, 0x20004}, [][]byte{t1, t2}, 0x16c9a0d6))); err != nil {
		t.Fatal(err)
	}
	if len(r.Towers) != 3 || !bytes.Equal(r.Towers[0].TowerOctetString, t1) || !bytes.Equal(r.Towers[2].TowerOctetString, t2) || r.Status != 0x16c9a0d6 {
		t.Fatalf("alias: %+v status=%#x", r.Towers, r.Status)
	}
	if r.Towers[1] != nil {
		t.Fatalf("alias: the aliased slot is set at the alias time, as before: %+v", r.Towers[1])
	}
}

func lookupEntry(id uint32, annotation string) []byte {
	var b bytes.Buffer
	b.Write(make([]byte, 16))
	b.Write(le(id))
	a := append([]byte(annotation), 0)
	b.Write(le(0, uint32(len(a))))
	b.Write(a)
	for b.Len()%4 != 0 {
		b.WriteByte(0)
	}
	return b.Bytes()
}

func TestLookupResponseEntryTowers(t *testing.T) {
	t1 := []byte{1, 0, 0, 0xA}
	t2 := []byte{2, 0, 0, 0xB, 5, 6}
	var b bytes.Buffer
	b.Write(make([]byte, 20))
	b.Write(le(2, 2, 0, 2))
	b.Write(lookupEntry(0x20000, "ann1"))
	b.Write(lookupEntry(0x20004, "ann2"))
	b.Write(tower(t1))
	b.Write(tower(t2))
	b.Write(le(0))

	var r LookupResponse
	if err := r.UnmarshalNDR(context.Background(), ndr.NDR20(b.Bytes())); err != nil {
		t.Fatal(err)
	}
	if len(r.Entries) != 2 || r.Entries[0].Annotation != "ann1" || r.Entries[1].Annotation != "ann2" ||
		!bytes.Equal(r.Entries[0].Tower.TowerOctetString, t1) || !bytes.Equal(r.Entries[1].Tower.TowerOctetString, t2) {
		t.Fatalf("entries: %+v", r.Entries)
	}
}
