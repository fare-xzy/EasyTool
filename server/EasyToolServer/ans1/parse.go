package ans1

import (
	"EasyToolServer/log"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Asn1Type struct {
	Tag, Type string
	Byte      byte
}

type ParseError struct {
	Msg string
}

func (e ParseError) Error() string { return "asn1: syntax error: " + e.Msg }

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
			resultValue, _ := parseInt(value)
			fmt.Println(resultValue)
		case IA5String:
			valLen := asn1Byte[cursor : cursor+2]
			valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
			cursor += 2
			value := asn1Byte[cursor : cursor+valLenInt*2]
			cursor = cursor + valLenInt*2
			ret, _ := parseIA5String(value)
			fmt.Println(ret)
		case UTF8String:
			valLen := asn1Byte[cursor : cursor+2]
			valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
			cursor += 2
			value := asn1Byte[cursor : cursor+valLenInt*2]
			cursor = cursor + valLenInt*2
			ret, _ := parseUTF8String(value)
			fmt.Println(ret)
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
			generalizedTime, _ := parseGeneralizedTime(value)
			fmt.Println(generalizedTime)
			log.Infof("HexValue = %s", parseString(value))
		case ObjectIdentifier:
			valLen := asn1Byte[cursor : cursor+2]
			valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
			cursor += 2
			value := asn1Byte[cursor : cursor+valLenInt*2]
			cursor = cursor + valLenInt*2
			log.Infof("HexValue = %s", parseString(value))
		case Boolean:
			valLen := asn1Byte[cursor : cursor+2]
			valLenInt, _ := strconv.ParseInt(string(valLen), 16, 0)
			cursor += 2
			value := asn1Byte[cursor : cursor+valLenInt*2]
			cursor = cursor + valLenInt*2
			log.Infof("HexValue = %t", parseBoolean(value))
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

// 转换为整型
func parseInt(value []byte) (resultValue int64, err error) {
	resultValue, err = strconv.ParseInt(string(value), 16, 0)
	return
}

// 转换为字符串
func parseString(value []byte) string {
	decodeByte, err := hex.DecodeString(string(value))
	if err != nil {
		return ""
	}
	// 将字节数组转换为字符串
	str := string(decodeByte)
	return str
}

func parseBoolean(value []byte) bool {
	if strings.EqualFold(string(value), "00") {
		return false
	} else {
		return true
	}
}

func parseGeneralizedTime(bytes []byte) (ret time.Time, err error) {
	decodeByte, err := hex.DecodeString(string(bytes))
	const formatStr = "20060102150405Z0700"
	s := string(decodeByte)

	if ret, err = time.Parse(formatStr, s); err != nil {
		return
	}

	if serialized := ret.Format(formatStr); serialized != s {
		err = fmt.Errorf("asn1: time did not serialize back to the original value and may be invalid: given %q, but serialized as %q", s, serialized)
	}

	return
}

func parseIA5String(bytes []byte) (ret string, err error) {
	for _, b := range bytes {
		if b >= utf8.RuneSelf {
			err = ParseError{"IA5String contains invalid character"}
			return
		}
	}
	ret = string(bytes)
	return
}

func parseUTF8String(bytes []byte) (ret string, err error) {
	if !utf8.Valid(bytes) {
		err = ParseError{"asn1: invalid UTF-8 string"}
		return
	}
	return string(bytes), nil
}
