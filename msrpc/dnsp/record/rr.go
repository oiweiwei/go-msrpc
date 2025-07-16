package record

import (
	"context"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/miekg/dns"

	"github.com/oiweiwei/go-msrpc/ndr"

	ms_dns "github.com/oiweiwei/go-msrpc/msrpc/dnsp/record/dns"
)

func (o *NodeName) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(o.DNSName))
}

func (o *NodeName) String() string {
	return string(o.DNSName)
}

func (o *NodeText) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(o.Text))
}

func (o *NodeText) String() string {
	return strings.TrimRight(string(o.Text), "\u0000")
}

func (o *RecordString) StringAt(i int) string {
	if len(o.Data) <= i {
		return ""
	}
	return strings.Trim(string(o.Data[i].Text), "\u0000")
}

func (o *RecordString) Strings() []string {
	s := make([]string, len(o.Data))
	for i := range o.Data {
		s[i] = strings.Trim(string(o.Data[i].Text), "\u0000")
	}
	return s
}

func (o *RecordIPAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(net.IP(o.IPAddress).String())
}

func (o *RecordIPAddress) String() string {
	return net.IP(o.IPAddress).String()
}

// RR returns the dns.RR representation of the node.
func (r *Node) RRs() []dns.RR {
	rrset := make([]dns.RR, len(r.DNSRecords))
	for i, rr := range r.DNSRecords {
		rrset[i] = rr.RR()
		rrset[i].Header().Name = string(r.DNSNodeName.DNSName)
	}
	return rrset
}

// NewNodeName creates a new NodeName from a string.
func NewNodeName(name string) *NodeName {
	return &NodeName{DNSName: []byte(name)}
}

// NewRecordNodeName creates a new RecordNodeName from a string.
func NewRecordNodeName(name string) *RecordNodeName {
	return &RecordNodeName{Name: NewNodeName(name)}
}

// NewRecordString creates a new RecordString from a slice of strings.
func NewRecordString(s ...string) *RecordString {
	data := make([]*NodeText, len(s))
	for i, v := range s {
		data[i] = NewNodeText(v)
	}
	return &RecordString{Data: data}
}

// NewNodeText creates a new NodeText from a string.
func NewNodeText(s string) *NodeText {
	return &NodeText{Text: []byte(s)}
}

// NewIPAddressRecord creates a new RecordIPAddress from a net.IP.
func NewIPAddressRecord(ip net.IP) *RecordIPAddress {
	if ip4 := ip.To4(); ip4 != nil {
		ip = ip4
	}
	return &RecordIPAddress{IPAddress: ip}
}

// NewRecordFromString creates a new Record from a string data.
func NewRecordFromString(typ uint16, ttl uint32, data string) (*Record, error) {

	rr, err := dns.NewRR(fmt.Sprintf(". %d %s %s", ttl, dns.TypeToString[typ], data))
	if err != nil {
		return nil, fmt.Errorf("new record from string: new rr: %w", err)
	}

	record, err := NewRecordFromRR(rr)
	if err != nil {
		return nil, fmt.Errorf("new record from string: new record from rr: %w", err)
	}

	return record, nil
}

// NewRecordFromRR creates a new Record from a dns.RR.
func NewRecordFromRR(rr dns.RR) (*Record, error) {

	val := &Record_Record{}

	switch rr := rr.(type) {
	case *dns.A:
		val.Value = &Record_A{RecordA: (*RecordA)(NewIPAddressRecord(rr.A))}
	case *dns.NS:
		val.Value = &Record_NS{RecordNS: (*RecordNS)(NewRecordNodeName(rr.Ns))}
	case *dns.MD:
		val.Value = &Record_MD{RecordMD: (*RecordMD)(NewRecordNodeName(rr.Md))}
	case *dns.MF:
		val.Value = &Record_MF{RecordMF: (*RecordMF)(NewRecordNodeName(rr.Mf))}
	case *dns.CNAME:
		val.Value = &Record_CNAME{RecordCNAME: (*RecordCNAME)(NewRecordNodeName(rr.Target))}
	case *dns.SOA:
		val.Value = &Record_SOA{RecordSOA: &RecordSOA{
			PrimaryServer:          NewNodeName(rr.Ns),
			ZoneAdministratorEmail: NewNodeName(rr.Mbox),
			SerialNo:               rr.Serial,
			Refresh:                rr.Refresh,
			Retry:                  rr.Retry,
			Expire:                 rr.Expire,
			MinimumTTL:             rr.Minttl,
		}}
	case *dns.MB:
		val.Value = &Record_MB{RecordMB: (*RecordMB)(NewRecordNodeName(rr.Mb))}
	case *dns.MG:
		val.Value = &Record_MG{RecordMG: (*RecordMG)(NewRecordNodeName(rr.Mg))}
	case *dns.MR:
		val.Value = &Record_MR{RecordMR: (*RecordMR)(NewRecordNodeName(rr.Mr))}
	case *dns.NULL:
		val.Value = &Record_NULL{RecordNULL: &RecordNULL{Data: []byte(rr.Data)}}
	case *dns.PTR:
		val.Value = &Record_PTR{RecordPTR: (*RecordPTR)(NewRecordNodeName(rr.Ptr))}
	case *dns.HINFO:
		val.Value = &Record_HINFO{RecordHINFO: (*RecordHINFO)(NewRecordString(rr.Cpu, rr.Os))}
	case *dns.MINFO:
		val.Value = &Record_MINFO{RecordMINFO: &RecordMINFO{
			MailBox:      NewNodeName(rr.Rmail),
			ErrorMailBox: NewNodeName(rr.Email),
		}}
	case *dns.MX:
		val.Value = &Record_MX{RecordMX: &RecordMX{
			Preference: rr.Preference,
			Exchange:   NewNodeName(rr.Mx),
		}}
	case *dns.TXT:
		val.Value = &Record_TXT{RecordTXT: (*RecordTXT)(NewRecordString(rr.Txt...))}
	case *dns.RP:
		val.Value = &Record_RP{RecordRP: &RecordRP{
			MailBox:      NewNodeName(rr.Mbox),
			ErrorMailBox: NewNodeName(rr.Txt),
		}}
	case *dns.AFSDB:
		val.Value = &Record_AFSDB{RecordAFSDB: &RecordAFSDB{
			Preference: rr.Subtype,
			Exchange:   NewNodeName(rr.Hostname),
		}}
	case *dns.X25:
		val.Value = &Record_X25{RecordX25: (*RecordX25)(NewRecordString(strings.Split(rr.PSDNAddress, " ")...))}
	case *dns.ISDN:
		val.Value = &Record_ISDN{RecordISDN: (*RecordISDN)(NewRecordString(rr.Address, rr.SubAddress))}
	case *dns.RT:
		val.Value = &Record_RT{RecordRT: &RecordRT{
			Preference: rr.Preference,
			Exchange:   NewNodeName(rr.Host),
		}}
	case *dns.SIG:
		signatureInfo := make([]byte, base64.StdEncoding.DecodedLen(len(rr.Signature)))
		_, err := base64.StdEncoding.Decode(signatureInfo, []byte(rr.Signature))
		if err != nil {
			return nil, fmt.Errorf("rr: sig: %w", err)
		}
		val.Value = &Record_SIG{RecordSIG: &RecordSIG{
			TypeCovered:   rr.TypeCovered,
			Algorithm:     rr.Algorithm,
			LabelCount:    rr.Labels,
			OriginalTTL:   rr.OrigTtl,
			SigExpiration: rr.Expiration,
			SigInception:  rr.Inception,
			KeyTag:        rr.KeyTag,
			Signer:        NewNodeName(rr.SignerName),
			SignatureInfo: signatureInfo,
		}}
	case *dns.KEY:
		key := make([]byte, base64.StdEncoding.DecodedLen(len(rr.PublicKey)))
		_, err := base64.StdEncoding.Decode(key, []byte(rr.PublicKey))
		if err != nil {
			return nil, fmt.Errorf("rr: key: %w", err)
		}
		val.Value = &Record_KEY{RecordKEY: &RecordKEY{
			Flags:     rr.Flags,
			Protocol:  rr.Protocol,
			Algorithm: rr.Algorithm,
			Key:       key,
		}}
	case *dns.AAAA:
		val.Value = &Record_AAAA{RecordAAAA: (*RecordAAAA)(NewIPAddressRecord(rr.AAAA))}
	case *dns.NXT:
		val.Value = &Record_NXT{RecordNXT: &RecordNXT{
			NextName:  NewNodeName(rr.NextDomain),
			TypeWords: rr.TypeBitMap,
		}}
	case *dns.SRV:
		val.Value = &Record_SRV{RecordSRV: &RecordSRV{
			Priority: rr.Priority,
			Weight:   rr.Weight,
			Port:     rr.Port,
			Target:   NewNodeName(rr.Target),
		}}
	case *dns.NAPTR:
		val.Value = &Record_NAPTR{RecordNAPTR: &RecordNAPTR{
			Order:        rr.Order,
			Preference:   rr.Preference,
			Flags:        NewNodeName(rr.Flags),
			Service:      NewNodeName(rr.Service),
			Substitution: NewNodeName(rr.Regexp),
			Replacement:  NewNodeName(rr.Replacement),
		}}
	case *dns.DNAME:
		val.Value = &Record_DNAME{RecordDNAME: (*RecordDNAME)(NewRecordNodeName(rr.Target))}
	case *dns.DS:
		digest := make([]byte, hex.DecodedLen(len(rr.Digest)))
		_, err := hex.Decode(digest, []byte(rr.Digest))
		if err != nil {
			return nil, fmt.Errorf("rr: ds: %w", err)
		}
		val.Value = &Record_DS{RecordDS: &RecordDS{
			KeyTag:     rr.KeyTag,
			Algorithm:  rr.Algorithm,
			DigestType: rr.DigestType,
			Digest:     digest,
		}}
	case *dns.RRSIG:
		signatureInfo := make([]byte, base64.StdEncoding.DecodedLen(len(rr.Signature)))
		_, err := base64.StdEncoding.Decode(signatureInfo, []byte(rr.Signature))
		if err != nil {
			return nil, fmt.Errorf("rr: rrsig: %w", err)
		}
		val.Value = &Record_RRSIG{RecordRRSIG: &RecordRRSIG{
			TypeCovered:   rr.TypeCovered,
			Algorithm:     rr.Algorithm,
			LabelCount:    rr.Labels,
			OriginalTTL:   rr.OrigTtl,
			SigExpiration: rr.Expiration,
			SigInception:  rr.Inception,
			KeyTag:        rr.KeyTag,
			Signer:        NewNodeName(rr.SignerName),
			SignatureInfo: signatureInfo,
		}}
	case *dns.NSEC:
		val.Value = &Record_NSEC{RecordNSEC: &RecordNSEC{
			Signer:     NewNodeName(rr.NextDomain),
			NSECBitmap: rr.TypeBitMap,
		}}
	case *dns.DNSKEY:
		key := make([]byte, base64.StdEncoding.DecodedLen(len(rr.PublicKey)))
		_, err := base64.StdEncoding.Decode(key, []byte(rr.PublicKey))
		if err != nil {
			return nil, fmt.Errorf("rr: dnskey: %w", err)
		}
		val.Value = &Record_DNSKEY{RecordDNSKEY: &RecordDNSKEY{
			Flags:     rr.Flags,
			Protocol:  rr.Protocol,
			Algorithm: rr.Algorithm,
			Key:       key,
		}}
	case *dns.DHCID:
		dhcid := make([]byte, base64.StdEncoding.DecodedLen(len(rr.Digest)))
		_, err := base64.StdEncoding.Decode(dhcid, []byte(rr.Digest))
		if err != nil {
			return nil, fmt.Errorf("rr: dhcid: %w", err)
		}
		val.Value = &Record_DHCID{RecordDHCID: &RecordDHCID{
			DHCID: dhcid,
		}}
	case *dns.NSEC3:
		salt := make([]byte, hex.DecodedLen(len(rr.Salt)))
		_, err := hex.Decode(salt, []byte(rr.Salt))
		if err != nil {
			return nil, fmt.Errorf("rr: nsec3: %w", err)
		}
		nextHashedOwnerName := make([]byte, base32.StdEncoding.DecodedLen(len(rr.NextDomain)))
		_, err = base32.StdEncoding.Decode(nextHashedOwnerName, []byte(rr.NextDomain))
		if err != nil {
			return nil, fmt.Errorf("rr: nsec3: %w", err)
		}
		val.Value = &Record_NSEC3{RecordNSEC3: &RecordNSEC3{
			Algorithm:           rr.Hash,
			Flags:               rr.Flags,
			Iterations:          rr.Iterations,
			SaltLength:          rr.SaltLength,
			Salt:                salt,
			HashLength:          rr.HashLength,
			NextHashedOwnerName: nextHashedOwnerName,
			Bitmaps:             rr.TypeBitMap,
		}}
	case *dns.NSEC3PARAM:
		salt := make([]byte, hex.DecodedLen(len(rr.Salt)))
		_, err := hex.Decode(salt, []byte(rr.Salt))
		if err != nil {
			return nil, fmt.Errorf("rr: nsec3param: %w", err)
		}
		val.Value = &Record_NSEC3PARAM{RecordNSEC3PARAM: &RecordNSEC3PARAM{
			Algorithm:  rr.Hash,
			Flags:      rr.Flags,
			Iterations: rr.Iterations,
			SaltLength: rr.SaltLength,
			Salt:       salt,
		}}
	case *dns.TLSA:
		cert := make([]byte, hex.DecodedLen(len(rr.Certificate)))
		_, err := hex.Decode(cert, []byte(rr.Certificate))
		if err != nil {
			return nil, fmt.Errorf("rr: tlsa: %w", err)
		}
		val.Value = &Record_TLSA{RecordTLSA: &RecordTLSA{
			CertUsage:                  rr.Usage,
			Selector:                   rr.Selector,
			MatchingType:               rr.MatchingType,
			CertificateAssociationData: cert,
		}}
	case *ms_dns.WINS:
		ips := make([][]byte, len(rr.WINSServers))
		for i := range ips {
			ip := rr.WINSServers[i]
			if ip4 := ip.To4(); ip4 != nil {
				ip = ip4
			}
			ips[i] = ip
		}
		val.Value = &Record_WINS{RecordWINS: &RecordWINS{
			MappingFlag:   rr.MappingFlag,
			LookupTimeout: rr.LookupTimeout,
			CacheTimeout:  rr.CacheTimeout,
			WINSServers:   ips,
		}}
	case *ms_dns.WINSR:
		val.Value = &Record_WINSR{RecordWINSR: &RecordWINSR{
			MappingFlag:   rr.MappingFlag,
			LookupTimeout: rr.LookupTimeout,
			CacheTimeout:  rr.CacheTimeout,
			ResultDomain:  NewNodeName(rr.ResultDomain),
		}}
	case *ms_dns.WKS:
		val.Value = &Record_WKS{RecordWKS: &RecordWKS{
			IPAddress: rr.IPAddress,
			Protocol:  rr.Protocol,
			Services:  NewNodeText(rr.Services),
		}}
	case *ms_dns.ATMA:
		val.Value = &Record_ATMA{RecordATMA: &RecordATMA{
			Format:  rr.Format,
			Address: rr.Address,
		}}
	default:
		return nil, fmt.Errorf("rr: unknown record type: %T", rr)
	}

	typ := TypeRecord(rr.Header().Rrtype)

	w := ndr.NDR20(nil, ndr.Opaque)

	if err := val.MarshalUnionNDR(context.Background(), w, (uint16)(typ)); err != nil {
		return nil, fmt.Errorf("rr: marshal: %w", err)
	}

	buffer := w.Bytes()

	return &Record{
		DataLength: uint16(len(buffer)),
		Type:       typ,
		TTLSeconds: rr.Header().Ttl,
		Record:     val,
		Buffer:     buffer,
	}, nil
}

func (r *Record) RR() dns.RR {

	h := dns.RR_Header{
		Rrtype: uint16(r.Type),
		Class:  dns.ClassINET,
		Ttl:    r.TTLSeconds,
	}

	rr := r.Record.GetValue()

	switch h.Rrtype {
	case dns.TypeA:
		return &dns.A{
			Hdr: h,
			A:   rr.(*RecordA).IPAddress,
		}
	case dns.TypeNS:
		return &dns.NS{
			Hdr: h,
			Ns:  rr.(*RecordNS).Name.String(),
		}
	case dns.TypeMD:
		return &dns.MD{
			Hdr: h,
			Md:  rr.(*RecordMD).Name.String(),
		}
	case dns.TypeMF:
		return &dns.MF{
			Hdr: h,
			Mf:  rr.(*RecordMF).Name.String(),
		}
	case dns.TypeCNAME:
		return &dns.CNAME{
			Hdr:    h,
			Target: rr.(*RecordCNAME).Name.String(),
		}
	case dns.TypeSOA:
		rr := rr.(*RecordSOA)
		return &dns.SOA{
			Hdr:     h,
			Ns:      rr.PrimaryServer.String(),
			Mbox:    rr.ZoneAdministratorEmail.String(),
			Serial:  rr.SerialNo,
			Refresh: rr.Refresh,
			Retry:   rr.Retry,
			Expire:  rr.Expire,
			Minttl:  rr.MinimumTTL,
		}
	case dns.TypeMB:
		return &dns.MB{
			Hdr: h,
			Mb:  rr.(*RecordMB).Name.String(),
		}
	case dns.TypeMG:
		return &dns.MG{
			Hdr: h,
			Mg:  rr.(*RecordMG).Name.String(),
		}
	case dns.TypeMR:
		return &dns.MR{
			Hdr: h,
			Mr:  rr.(*RecordMR).Name.String(),
		}
	case dns.TypeNULL:
		return &dns.NULL{
			Hdr:  h,
			Data: string(rr.(*RecordNULL).Data),
		}
	case dns.TypePTR:
		return &dns.PTR{
			Hdr: h,
			Ptr: rr.(*RecordPTR).Name.String(),
		}
	case dns.TypeHINFO:
		rr := rr.(*RecordHINFO)
		return &dns.HINFO{
			Hdr: h,
			Cpu: (*RecordString)(rr).StringAt(0),
			Os:  (*RecordString)(rr).StringAt(1),
		}
	case dns.TypeMINFO:
		rr := rr.(*RecordMINFO)
		return &dns.MINFO{
			Hdr:   h,
			Rmail: rr.MailBox.String(),
			Email: rr.ErrorMailBox.String(),
		}
	case dns.TypeMX:
		rr := rr.(*RecordMX)
		return &dns.MX{
			Hdr:        h,
			Preference: rr.Preference,
			Mx:         rr.Exchange.String(),
		}
	case dns.TypeTXT:
		rr := rr.(*RecordTXT)
		return &dns.TXT{
			Hdr: h,
			Txt: (*RecordString)(rr).Strings(),
		}
	case dns.TypeRP:
		rr := rr.(*RecordRP)
		return &dns.RP{
			Hdr:  h,
			Mbox: rr.MailBox.String(),
			Txt:  rr.ErrorMailBox.String(),
		}
	case dns.TypeAFSDB:
		rr := rr.(*RecordAFSDB)
		return &dns.AFSDB{
			Hdr:      h,
			Subtype:  rr.Preference,
			Hostname: rr.Exchange.String(),
		}
	case dns.TypeX25:
		rr := rr.(*RecordX25)
		return &dns.X25{
			Hdr:         h,
			PSDNAddress: strings.Join((*RecordString)(rr).Strings(), " "),
		}
	case dns.TypeISDN:
		rr := rr.(*RecordISDN)
		return &dns.ISDN{
			Hdr:        h,
			Address:    (*RecordString)(rr).StringAt(0),
			SubAddress: (*RecordString)(rr).StringAt(1),
		}
	case dns.TypeRT:
		rr := rr.(*RecordRT)
		return &dns.RT{
			Hdr:        h,
			Preference: rr.Preference,
			Host:       rr.Exchange.String(),
		}
	case dns.TypeSIG:
		rr := rr.(*RecordSIG)
		return &dns.SIG{
			RRSIG: dns.RRSIG{
				Hdr:         h,
				TypeCovered: rr.TypeCovered,
				Algorithm:   rr.Algorithm,
				Labels:      rr.LabelCount,
				OrigTtl:     rr.OriginalTTL,
				Expiration:  rr.SigExpiration,
				Inception:   rr.SigInception,
				KeyTag:      rr.KeyTag,
				SignerName:  rr.Signer.String(),
				Signature:   base64.StdEncoding.EncodeToString(rr.SignatureInfo),
			},
		}
	case dns.TypeKEY:
		rr := rr.(*RecordKEY)
		return &dns.KEY{
			DNSKEY: dns.DNSKEY{
				Hdr:       h,
				Flags:     rr.Flags,
				Protocol:  rr.Protocol,
				Algorithm: rr.Algorithm,
				PublicKey: base64.StdEncoding.EncodeToString(rr.Key),
			},
		}

	case dns.TypeAAAA:
		return &dns.AAAA{
			Hdr:  h,
			AAAA: rr.(*RecordAAAA).IPAddress,
		}
	case dns.TypeNXT:
		rr := rr.(*RecordNXT)
		return &dns.NXT{
			NSEC: dns.NSEC{
				Hdr:        h,
				NextDomain: rr.NextName.String(),
				TypeBitMap: rr.TypeWords,
			},
		}
	case dns.TypeSRV:
		rr := rr.(*RecordSRV)
		return &dns.SRV{
			Hdr:      h,
			Priority: rr.Priority,
			Weight:   rr.Weight,
			Port:     rr.Port,
			Target:   rr.Target.String(),
		}
	case dns.TypeNAPTR:
		rr := rr.(*RecordNAPTR)
		return &dns.NAPTR{
			Hdr:         h,
			Order:       rr.Order,
			Preference:  rr.Preference,
			Flags:       rr.Flags.String(),
			Service:     rr.Service.String(),
			Regexp:      rr.Substitution.String(),
			Replacement: rr.Replacement.String(),
		}
	case dns.TypeDNAME:
		rr := rr.(*RecordDNAME)
		return &dns.DNAME{
			Hdr:    h,
			Target: rr.Name.String(),
		}
	case dns.TypeDS:
		rr := rr.(*RecordDS)
		return &dns.DS{
			Hdr:        h,
			KeyTag:     rr.KeyTag,
			Algorithm:  rr.Algorithm,
			DigestType: rr.DigestType,
			Digest:     hex.EncodeToString(rr.Digest),
		}
	case dns.TypeRRSIG:
		rr := rr.(*RecordRRSIG)
		return &dns.RRSIG{
			Hdr:         h,
			TypeCovered: rr.TypeCovered,
			Algorithm:   rr.Algorithm,
			Labels:      rr.LabelCount,
			OrigTtl:     rr.OriginalTTL,
			Expiration:  rr.SigExpiration,
			Inception:   rr.SigInception,
			KeyTag:      rr.KeyTag,
			SignerName:  rr.Signer.String(),
			Signature:   base64.StdEncoding.EncodeToString(rr.SignatureInfo),
		}
	case dns.TypeNSEC:
		rr := rr.(*RecordNSEC)
		return &dns.NSEC{
			Hdr:        h,
			NextDomain: rr.Signer.String(),
			TypeBitMap: rr.NSECBitmap,
		}
	case dns.TypeDNSKEY:
		rr := rr.(*RecordDNSKEY)
		return &dns.DNSKEY{
			Hdr:       h,
			Flags:     rr.Flags,
			Protocol:  rr.Protocol,
			Algorithm: rr.Algorithm,
			PublicKey: base64.StdEncoding.EncodeToString(rr.Key),
		}
	case dns.TypeDHCID:
		rr := rr.(*RecordDHCID)
		return &dns.DHCID{
			Hdr:    h,
			Digest: base64.StdEncoding.EncodeToString(rr.DHCID),
		}
	case dns.TypeNSEC3:
		rr := rr.(*RecordNSEC3)
		return &dns.NSEC3{
			Hdr:        h,
			Hash:       rr.Algorithm,
			Flags:      rr.Flags,
			Iterations: rr.Iterations,
			SaltLength: rr.SaltLength,
			Salt:       hex.EncodeToString(rr.Salt),
			HashLength: rr.HashLength,
			NextDomain: base32.StdEncoding.EncodeToString(rr.NextHashedOwnerName),
			TypeBitMap: rr.Bitmaps,
		}
	case dns.TypeNSEC3PARAM:
		rr := rr.(*RecordNSEC3PARAM)
		return &dns.NSEC3PARAM{
			Hdr:        h,
			Hash:       rr.Algorithm,
			Flags:      rr.Flags,
			Iterations: rr.Iterations,
			SaltLength: rr.SaltLength,
			Salt:       hex.EncodeToString(rr.Salt),
		}

	case dns.TypeTLSA:
		rr := rr.(*RecordTLSA)
		return &dns.TLSA{
			Hdr:          h,
			Usage:        rr.CertUsage,
			Selector:     rr.Selector,
			MatchingType: rr.MatchingType,
			Certificate:  hex.EncodeToString(rr.CertificateAssociationData),
		}
	case ms_dns.TypeWINS:
		rr := rr.(*RecordWINS)
		ips := make([]net.IP, len(rr.WINSServers))
		for i := range ips {
			ips[i] = rr.WINSServers[i]
			if ip4 := ips[i].To4(); ip4 != nil {
				ips[i] = ip4
			}
		}
		return &ms_dns.WINS{
			ANY: dns.ANY{
				Hdr: h,
			},
			MappingFlag:   rr.MappingFlag,
			LookupTimeout: rr.LookupTimeout,
			CacheTimeout:  rr.CacheTimeout,
			WINSServers:   ips,
		}
	case ms_dns.TypeWINSR:
		rr := rr.(*RecordWINSR)
		return &ms_dns.WINSR{
			ANY: dns.ANY{
				Hdr: h,
			},
			MappingFlag:   rr.MappingFlag,
			LookupTimeout: rr.LookupTimeout,
			CacheTimeout:  rr.CacheTimeout,
			ResultDomain:  rr.ResultDomain.String(),
		}
	case ms_dns.TypeWKS:
		rr := rr.(*RecordWKS)
		return &ms_dns.WKS{
			ANY: dns.ANY{
				Hdr: h,
			},
			IPAddress: rr.IPAddress,
			Protocol:  rr.Protocol,
			Services:  rr.Services.String(),
		}
	case dns.TypeATMA:
		rr := rr.(*RecordATMA)
		return &ms_dns.ATMA{
			ANY: dns.ANY{
				Hdr: h,
			},
			Format:  rr.Format,
			Address: rr.Address,
		}
	default:

		// *Record_ATMA
		// *Record_Unknown
	}

	return &dns.NULL{
		Hdr:  h,
		Data: hex.EncodeToString(r.Buffer),
	}

}
