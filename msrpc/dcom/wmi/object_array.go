package wmi

import (
	"fmt"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmio"
	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
)

func UnmarshalObjectArrayWithClass(b []byte, cls wmio.Class) (*ObjectArray, error) {
	o := &ObjectArray{}
	return o, o.DecodeWithClass(wmio.NewCodec(b), cls)
}

// The ObjectArray structure MUST be used to encode multiple CIM
// objects that are returned in response to the IWbemWCOSmartEnum::Next
// method. This structure is also used to encode parameters of the
// optimized IWbemObjectSink::Indicate method.<6> To minimize network
// bandwidth, a server SHOULD support the ObjectArray structure when
// an array of CIM objects is sent.
type ObjectArray struct {
	// dwByteOrdering (4 bytes): The byte ordering. It MUST be value 0.
	ByteOrdering uint32
	// abSignature (8 bytes): MUST be set to {0x57, 0x42, 0x45, 0x4D,
	// 0x44, 0x41, 0x54, 0x41} (a byte array containing the unquoted,
	// unterminated ASCII string "WBEMDATA").
	Signature []byte
	// dwSizeOfHeader1 (4 bytes): This stores the total size of these
	// fields: dwByteOrdering, abSignature, dwSizeofHeader1, dwDataSize1,
	// dwFlags, bVersion, and bPacketType.
	// The size of the header MUST be 0x0000001A. Data immediately follows
	// the header.
	SizeOfHeader1 uint32
	// dwDataSize1 (4 bytes): MUST indicate the length, in bytes, of the
	// data that follows this header, starting at the dwSizeOfHeader2 field.
	DataSize1 uint32
	// dwFlags (4 bytes): The flag value MUST be 0x00000000.
	Flags uint32
	// bVersion (1 byte): The version number of the header. The version MUST be 1.
	Version uint8
	// bPacketType (1 byte): The value of this field is dependent on the call
	// context.
	PacketType uint8
	// dwSizeOfHeader2 (4 bytes): This stores the size of these fields:
	// dwSizeofHeader2 and dwDataSize2.
	// This value MUST be 8. Data immediately follows after the field dwDataSize2.
	SizeOfHeader2 uint32
	// dwDataSize2 (4 bytes): MUST be the size, in bytes, of the data that follows
	// this field.
	DataSize2 uint32
	// dwSizeOfHeader3 (4 bytes):  This stores the size of these fields:
	// dwSizeofHeader3, dwDataSize3, and dwNumObjects. This value MUST be 12.
	// Data immediately follows after the field dwNumObjects.
	SizeOfHeader3 uint32
	// dwDataSize3 (4 bytes): MUST indicate the length of the remaining data,
	// starting at the wbemObjects field.
	DataSize3 uint32
	// dwNumObjects (4 bytes): MUST be the number of CIM objects in the
	// ObjectArray.
	NumObjects uint32
	// wbemObjects (variable): The objects array that contains the CIM class
	// definition and CIM instances. These CIM objects MUST be encoded
	// in the WBEM_DATAPACKET_OBJECT structure.
	Objects []*DataPacketObject
}

// The WBEM_DATAPACKET_OBJECT MUST contain the CIM class definition and CIM instances.
type DataPacketObject struct {
	// dwSizeOfHeader (4 bytes): The size, in bytes, of the
	// WBEM_DATAPACKET_OBJECT header, which MUST be 0x00000009.
	SizeOfHeader uint32
	// dwSizeOfData (4 bytes): The size, in bytes, of the data
	// following the WBEM_DATAPACKET_OBJECT header.
	SizeOfData uint32
	// bObjectType (1 byte): The type of data in the data packet.
	// The type MUST take one of the following specified values.
	//
	// * 1:
	//  	Object is type WBEMOBJECT_CLASS:
	//		Structure contains the complete CIM Class definition.
	// * 2:
	// 		Object is type WBEMOBJECT_INSTANCE:
	//		Structure contains the complete CIM Instance definition.
	// * 3:
	//		Object is type WBEMOBJECT_INSTANCE_NOCLASS.
	//		Structure contains CIM Instance without the CIM Class definition.
	ObjectType uint8

	// dwSizeOfHeader (4 bytes): The size, in bytes, of the header, which
	// MUST be 0x00000008 for WBEMOBJECT_CLASS and 0x00000018 for WBEMOBJECT_INSTANCE,
	// and WBEMOBJECT_INSTANCE_NOCLASS.
	SizeOfHeader2 uint32
	// dwSizeOfData (4 bytes): The size, in bytes, of the data that follows the
	/// header.
	SizeOfData2 uint32
	// classID (16 bytes): The unique identifier of the CIM class type.
	ClassID *dcom.ClassID
	// Object (variable): The CIM object carried into the
	// WBEM_DATAPACKET_OBJECT, having dwSizeOfData bytes. The embedded CIM
	// object MUST match the selector field bObjectType.
	Object *wmio.Object
}

func (o *ObjectArray) DecodeWithClass(r *wmio.Codec, cls wmio.Class) error {

	r.ReadData(&o.ByteOrdering)
	o.Signature = make([]byte, 8)
	r.Read(o.Signature)
	r.ReadData(&o.SizeOfHeader1)
	r.ReadData(&o.DataSize1)
	r.ReadData(&o.Flags)
	r.ReadData(&o.Version)
	r.ReadData(&o.PacketType)
	r.ReadData(&o.SizeOfHeader2)
	r.ReadData(&o.DataSize2)
	r.ReadData(&o.SizeOfHeader3)
	r.ReadData(&o.DataSize3)
	r.ReadData(&o.NumObjects)

	for i := 0; i < int(o.NumObjects); i++ {

		var po DataPacketObject

		r.ReadData(&po.SizeOfHeader)
		r.ReadData(&po.SizeOfData)
		r.ReadData(&po.ObjectType)

		r.ReadData(&po.SizeOfHeader2)
		r.ReadData(&po.SizeOfData2)

		if r.Err() != nil {
			return r.Err()
		}

		var err error

		switch po.ObjectType {
		case 1:
			if po.Object, err = wmio.UnmarshalWithClass(r.Bytes()[:po.SizeOfData2], wmio.Class{}); err != nil {
				return err
			}
			cls = po.Object.Class.CurrentClass
		case 2:

			if po.ClassID, err = ReadClassID(r); err != nil {
				return err
			}

			if po.Object, err = wmio.UnmarshalWithClass(r.Bytes()[:po.SizeOfData2], wmio.Class{}); err != nil {
				return err
			}

			cls = po.Object.Instance.CurrentClass
		case 3:
			if cls.Name == "" {
				return fmt.Errorf("class name is empty")
			}

			if po.ClassID, err = ReadClassID(r); err != nil {
				return err
			}

			if po.Object, err = wmio.UnmarshalWithClass(r.Bytes()[:po.SizeOfData2], cls); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown class type %d", po.ObjectType)
		}

		o.Objects, r = append(o.Objects, &po), wmio.NewCodec(r.Bytes()[po.SizeOfData2:])
	}

	return r.Err()
}

func ReadClassID(r *wmio.Codec) (*dcom.ClassID, error) {
	clsID := make([]byte, 16)
	if err := r.Read(clsID); err != nil {
		return nil, fmt.Errorf("class_id: read: %w", err)
	}

	guid, err := dtyp.GUIDFromBytes(clsID)
	if err != nil {
		return nil, fmt.Errorf("class_id: decode: %w", err)
	}

	return (*dcom.ClassID)(guid), nil
}
