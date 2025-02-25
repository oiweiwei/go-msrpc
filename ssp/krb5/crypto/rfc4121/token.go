package rfc4121

const (
	CFXFlagSendByAcceptor uint8 = 0x01
	CFXFlagSealed         uint8 = 0x02
	CFXFlagAcceptorSubKey uint8 = 0x04
)

type WrapToken struct {
	TokenID        uint16
	Flags          uint8
	Filler         [1]byte
	EC             uint16
	RRC            uint16
	SequenceNumber []byte
}
