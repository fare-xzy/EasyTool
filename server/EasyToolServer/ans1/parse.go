package ans1

import (
	"EasyToolServer/log"
	"strconv"
	"strings"
)

type Asn1Type struct {
	Tag, Type string
	Byte      byte
}

func ParseAsn1(asn1Byte []byte) {
	var cursor, entireLen int64
	entireLen = int64(len(asn1Byte))
	log.Infof("entireLen = %d", entireLen)
	for cursor < entireLen {
		tag := asn1Byte[cursor : cursor+2]
		log.Infof("%s", string(tag))
		cursor += 2
		switch string(tag) {
		case Sequence:
			if com := asn1Byte[cursor : cursor+2]; strings.EqualFold(string(com), "82") {
				cursor += 2
				log.Infof("%s", string(com))
				valLen := asn1Byte[cursor : cursor+4]
				valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
				cursor += 4
				ParseAsn1(asn1Byte[cursor : cursor+valLenInt*2])
				cursor = cursor + valLenInt*2
			} else {
				valLen := asn1Byte[cursor : cursor+2]
				valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
				cursor += 2
				ParseAsn1(asn1Byte[cursor : cursor+valLenInt*2])
				cursor = cursor + valLenInt*2
			}
		case Integer:
			valLen := asn1Byte[cursor : cursor+2]
			valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
			cursor += 2
			value := asn1Byte[cursor : cursor+valLenInt*2]
			cursor = cursor + valLenInt*2
			log.Infof("HexValue = %s", string(value))
		case IA5String:
			valLen := asn1Byte[cursor : cursor+2]
			valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
			cursor += 2
			value := asn1Byte[cursor : cursor+valLenInt*2]
			cursor = cursor + valLenInt*2
			log.Infof("HexValue = %s", string(value))
		case UTF8String:
			valLen := asn1Byte[cursor : cursor+2]
			valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
			cursor += 2
			value := asn1Byte[cursor : cursor+valLenInt*2]
			cursor = cursor + valLenInt*2
			log.Infof("HexValue = %s", string(value))
		case OctetString:
			if com := asn1Byte[cursor : cursor+2]; strings.EqualFold(string(com), "82") {
				cursor += 2
				log.Infof("%s", string(com))
				valLen := asn1Byte[cursor : cursor+4]
				valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
				cursor += 4
				value := asn1Byte[cursor : cursor+valLenInt*2]
				cursor = cursor + valLenInt*2
				log.Infof("HexValue = %s", string(value))
			} else {
				valLen := asn1Byte[cursor : cursor+2]
				valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
				cursor += 2
				value := asn1Byte[cursor : cursor+valLenInt*2]
				cursor = cursor + valLenInt*2
				log.Infof("HexValue = %s", string(value))
			}
		case GeneralizedTime:
			valLen := asn1Byte[cursor : cursor+2]
			valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
			cursor += 2
			value := asn1Byte[cursor : cursor+valLenInt*2]
			cursor = cursor + valLenInt*2
			log.Infof("HexValue = %s", string(value))
		case ObjectIdentifier:
			valLen := asn1Byte[cursor : cursor+2]
			valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
			cursor += 2
			value := asn1Byte[cursor : cursor+valLenInt*2]
			cursor = cursor + valLenInt*2
			log.Infof("HexValue = %s", string(value))
		case Boolean:
			valLen := asn1Byte[cursor : cursor+2]
			valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
			cursor += 2
			value := asn1Byte[cursor : cursor+valLenInt*2]
			cursor = cursor + valLenInt*2
			log.Infof("HexValue = %s", string(value))
		case BitString:
			if com := asn1Byte[cursor : cursor+2]; strings.EqualFold(string(com), "82") {
				cursor += 2
				log.Infof("%s", string(com))
				valLen := asn1Byte[cursor : cursor+4]
				valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
				cursor += 4
				value := asn1Byte[cursor : cursor+valLenInt*2]
				cursor = cursor + valLenInt*2
				log.Infof("HexValue = %s", string(value))
			} else {
				valLen := asn1Byte[cursor : cursor+2]
				valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
				cursor += 2
				value := asn1Byte[cursor : cursor+valLenInt*2]
				cursor = cursor + valLenInt*2
				log.Infof("HexValue = %s", string(value))
			}
		case Optional:
			if com := asn1Byte[cursor : cursor+2]; strings.EqualFold(string(com), "82") {
				cursor += 2
				log.Infof("%s", string(com))
				valLen := asn1Byte[cursor : cursor+4]
				valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
				cursor += 4
				ParseAsn1(asn1Byte[cursor : cursor+valLenInt*2])
				cursor = cursor + valLenInt*2
			} else {
				valLen := asn1Byte[cursor : cursor+2]
				valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
				cursor += 2
				ParseAsn1(asn1Byte[cursor : cursor+valLenInt*2])
				cursor = cursor + valLenInt*2
			}
		}
	}
}
