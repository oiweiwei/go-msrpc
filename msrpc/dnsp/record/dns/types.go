package dns

import (
	"encoding/hex"
	"fmt"
	"net"
	"strconv"

	"github.com/miekg/dns"
)

var (
	TypeATMA  uint16 = 0x22
	TypeWINS  uint16 = 0xFF01
	TypeWINSR uint16 = 0xFF02
	TypeWKS   uint16 = 0x0B
)

func init() {
	dns.TypeToString[TypeATMA] = "ATMA"
	dns.TypeToString[TypeWINSR] = "WINSR"
	dns.TypeToString[TypeWINS] = "WKS"
	dns.TypeToString[TypeWKS] = "WINS"
}

type ATMA struct {
	dns.ANY
	Format  uint8
	Address []byte
}

func (rr *ATMA) String() string {
	s := rr.Hdr.String()
	switch rr.Format {
	case 1:
		return s + "E164" + " " + string(rr.Address)
	case 2:
		s += "NSAP" + " "
		for i := range rr.Address {
			s += strconv.FormatUint(uint64(rr.Address[i]), 16)
		}
		return s
	}

	return s + "?" + " " + hex.EncodeToString(rr.Address)
}

type WINS struct {
	dns.ANY
	MappingFlag   uint32
	LookupTimeout uint32
	CacheTimeout  uint32
	WINSServers   []net.IP
}

func (rr *WINS) String() string {
	s := rr.Hdr.String() +
		strconv.Itoa(int(rr.MappingFlag)) +
		" " + strconv.Itoa(int(rr.LookupTimeout)) +
		" " + strconv.Itoa(int(rr.CacheTimeout))
	for _, ip := range rr.WINSServers {
		s += " " + ip.String()
	}

	return s
}

type WINSR struct {
	dns.ANY
	MappingFlag   uint32
	LookupTimeout uint32
	CacheTimeout  uint32
	ResultDomain  string
}

func (rr *WINSR) String() string {
	return rr.Hdr.String() +
		strconv.Itoa(int(rr.MappingFlag)) +
		" " + strconv.Itoa(int(rr.LookupTimeout)) +
		" " + strconv.Itoa(int(rr.CacheTimeout)) +
		" " + rr.ResultDomain
}

type WKS struct {
	dns.ANY
	IPAddress net.IP
	Protocol  uint8
	Services  string
}

func (rr *WKS) String() string {
	return rr.Hdr.String() +
		rr.IPAddress.String() +
		" " + strconv.Itoa(int(rr.Protocol)) +
		" " + fmt.Sprintf("%q", rr.Services)
}
