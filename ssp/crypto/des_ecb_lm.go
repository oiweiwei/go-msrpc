package crypto

import "encoding/binary"

// DES_ECB_LM encrypts the data using the algorithm specified in
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-samr/b1b0094f-2546-431f-b06d-582158a9f2bb
func DES_ECB_LM(k uint32, b []byte) ([]byte, error) {

	key := binary.LittleEndian.AppendUint32(nil, k)

	key1 := []byte{key[0], key[1], key[2], key[3], key[0], key[1], key[2]}
	key2 := []byte{key[3], key[0], key[1], key[2], key[3], key[0], key[1]}

	h1 := DES_ECB(key1, b[:8], false)
	h2 := DES_ECB(key2, b[8:], false)

	return append(h1, h2...), nil
}
