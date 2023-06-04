package ans1

import (
	"EasyToolServer/log"
)

type Asn1Type struct {
	Tag, Type string
	Byte      byte
}

func ParseAsn1(asn1Byte []byte) {
	tag := asn1Byte[0:2]
	log.Infof("%s", string(tag))
	switch string(tag) {
	case Sequence:

	}
}
