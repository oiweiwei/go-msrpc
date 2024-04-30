package record

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net"
	"strings"

	"github.com/miekg/dns"

	ms_dns "github.com/oiweiwei/go-msrpc/msrpc/dnsp/record/dns"
)

func (o *NodeName) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(o.DNSName))
}

func (o *NodeName) String() string {
	n := string(o.DNSName)
	if n == "" {
		return "."
	}
	return n
}

func (o *NodeText) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(o.Text))
}

func (o *NodeText) String() string {
	return strings.TrimRight(string(o.Text), "\u0000")
}

func (o *RecordString) StringAt(i int) string {
	if len(o.Data) < i {
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

func (r *Node) RRs() []dns.RR {
	rrset := make([]dns.RR, len(r.DNSRecords))
	for i, rr := range r.DNSRecords {
		rrset[i] = rr.RR()
		rrset[i].Header().Name = string(r.DNSNodeName.DNSName)
	}
	return rrset
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
