// 对 go-playground/validator/v10 的拙劣测试
// 使用方法看源码中的 doc.go
// 更多测试看源码中的 _examples & validator_test.go
// 有时间撸下源码也很舒服

package test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var (
	validate *validator.Validate
	err      error
)

func init() {
	validate = validator.New()
}

// eqcsfield			Field Equals Another Field (relative)
// eqfield				Field Equals Another Field
// fieldcontains		NOT DOCUMENTED IN doc.go
// fieldexcludes		NOT DOCUMENTED IN doc.go
// gtcsfield			Field Greater Than Another Relative Field
// gtecsfield			Field Greater Than or Equal To Another Relative Field
// gtefield				Field Greater Than or Equal To Another Field
// gtfield				Field Greater Than Another Field
// ltcsfield			Less Than Another Relative Field
// ltecsfield			Less Than or Equal To Another Relative Field
// ltefield				Less Than or Equal To Another Field
// ltfield				Less Than Another Field
// necsfield			Field Does Not Equal Another Field (relative)
// nefield				Field Does Not Equal Another Field
func TestFieldsTag(t *testing.T) {
	type B struct {
		Z string
	}
	type A struct {
		X   string `validate:"eqfield=Y"`
		Y   string `validate:"eqcsfield=BBB.Z"`
		BBB *B
	}
	param := &A{
		X:   "asd",
		Y:   "asd",
		BBB: &B{Z: "asd"},
	}
	err = validate.Struct(param)
	assert.NoError(t, err)

	param2 := struct {
		X   string `validate:"gtfield=Y"`
		Y   string `validate:"gtcsfield=ZZZ.Z"`
		ZZZ struct {
			Z string
		}
	}{
		X:   "hahaha",
		Y:   "haha",
		ZZZ: struct{ Z string }{Z: "ha"},
	}
	err = validate.Struct(param2)
	assert.NoError(t, err)
}

// cidr					Classless Inter-Domain Routing CIDR
// cidrv4				Classless Inter-Domain Routing CIDRv4
// cidrv6				Classless Inter-Domain Routing CIDRv6
// datauri				Data URL
// fqdn					Full Qualified Domain Name (FQDN)
// hostname				Hostname RFC 952
// hostname_port		HostPort
// hostname_rfc1123		Hostname RFC 1123
// ip					Internet Protocol Address IP
// ip4_addr				Internet Protocol Address IPv4
// ip6_addr				Internet Protocol Address IPv6
// ip_addr				Internet Protocol Address IP
// ipv4					Internet Protocol Address IPv4
// ipv6					Internet Protocol Address IPv6
// mac					Media Access Control Address MAC
// tcp4_addr			Transmission Control Protocol Address TCPv4
// tcp6_addr			Transmission Control Protocol Address TCPv6
// tcp_addr				Transmission Control Protocol Address TCP
// udp4_addr			User Datagram Protocol Address UDPv4
// udp6_addr			User Datagram Protocol Address UDPv6
// udp_addr				User Datagram Protocol Address UDP
// unix_addr			Unix domain socket end point Address
// uri					URI String
// url					URL String
// url_encoded			URL Encoded
// urn_rfc2141			Urn RFC 2141 String
func TestNetworkTag(t *testing.T) {
	hostname := "google.com"
	err = validate.Var(hostname, "hostname")
	assert.NoError(t, err)

	ip := "127.0.0.1"
	err = validate.Var(ip, "ip")
	assert.NoError(t, err)

	uri := "http://google.com"
	err = validate.Var(uri, "uri")
	assert.NoError(t, err)
}

// alpha				Alpha Only
// alphanum				Alphanumeric
// alphanumunicode		Alphanumeric Unicode
// alphaunicode			Alpha Unicode
// ascii				ASCII
// contains				Contains
// containsany			Contains Any
// containsrune			Contains Rune
// endswith				Ends With
// lowercase			Lowercase
// multibyte			Multi-Byte Characters
// number				NOT DOCUMENTED IN doc.go
// numeric				Numeric
// printascii			Printable ASCII
// startswith			Starts With
// uppercase			Uppercase
func TestStringsTag(t *testing.T) {
	contains := "hehe@hehe.com"
	err = validate.Var(contains, "contains=@")
	assert.NoError(t, err)

	endswith := "bye bye,my lover"
	err = validate.Var(endswith, "endswith=lover")
	assert.NoError(t, err)

	number := -123.12
	err = validate.Var(number, "number")
	assert.NoError(t, err)

	numeric := -123.12
	err = validate.Var(numeric, "numeric")
	assert.NoError(t, err)
}

// base64				Base64 String
// base64url			Base64URL String
// btc_addr				Bitcoin Address
// btc_addr_bech32		Bitcoin Bech32 Address (segwit)
// datetime				Datetime
// e164					e164 formatted phone number
// email				E-mail String
// eth_addr				Ethereum Address
// hexadecimal			Hexadecimal String
// hexcolor				Hexcolor String
// hsl					HSL String
// hsla					HSLA String
// html					HTML Tags
// html_encoded			HTML Encoded
// isbn					International Standard Book Number
// isbn10				International Standard Book Number 10
// isbn13				International Standard Book Number 13
// json					JSON
// latitude				Latitude
// longitude			Longitude
// rgb					RGB String
// rgba					RGBA String
// ssn					Social Security Number SSN
// uuid					Universally Unique Identifier UUID
// uuid3				Universally Unique Identifier UUID v3
// uuid3_rfc4122		Universally Unique Identifier UUID v3 RFC4122
// uuid4				Universally Unique Identifier UUID v4
// uuid4_rfc4122		Universally Unique Identifier UUID v4 RFC4122
// uuid5				Universally Unique Identifier UUID v5
// uuid5_rfc4122		Universally Unique Identifier UUID v5 RFC4122
// uuid_rfc4122			Universally Unique Identifier UUID RFC4122
func TestFormatTag(t *testing.T) {
	base64 := "aGFoYWhh"
	err = validate.Var(base64, "base64")
	assert.NoError(t, err)

	datetime := "2021-07-07 15:40:45"
	err = validate.Var(datetime, "datetime=2006-01-02 15:04:05")
	assert.NoError(t, err)

	email := "aaa@gmail.com"
	err = validate.Var(email, "email")
	assert.NoError(t, err)

	html := "<!DOCTYPE html>\n<html lang=\"en\"></html>"
	err = validate.Var(html, "html")
	assert.NoError(t, err)

	json := "{\"zzz\": 1}"
	err = validate.Var(json, "json")
	assert.NoError(t, err)
}

// eq					Equals
// gt					Greater than
// gte					Greater than or equal
// lt					Less Than
// lte					Less Than or Equal
// ne					Not Equal
func TestComparisonsTag(t *testing.T) {
	eq := "asd"
	err = validate.Var(eq, "eq=asd")
	assert.NoError(t, err)

	gte := "qwer"
	err = validate.Var(gte, "gte=4")
	assert.NoError(t, err)
}

// -					Skip Field
// |					Or Operator
// structonly			StructOnly
// nostructlevel		NoStructLevel
// omitempty			Omit Empty
// dive					Dive
// dir					Directory
// endswith				Ends With
// excludes				Excludes
// excludesall			Excludes All
// excludesrune			Excludes Rune
// file					File path
// isdefault			Is Default
// len					Length
// max					Maximum
// min					Minimum
// oneof				One Of
// required				Required
// required_if			Required If
// required_unless		Required Unless
// required_with		Required With
// required_with_all	Required With All
// required_without		Required Without
// required_without_all	Required Without All
// excluded_with		Excluded With
// excluded_with_all	Excluded With All
// excluded_without		Excluded Without
// excluded_without_all	Excluded Without All
// unique				Unique
func TestOtherTag(t *testing.T) {
	omitempty := ""
	err = validate.Var(omitempty, "omitempty")
	assert.NoError(t, err)

	dir := "./"
	err = validate.Var(dir, "dir")
	assert.NoError(t, err)

	file := "./validator_test.go"
	err = validate.Var(file, "file")
	assert.NoError(t, err)

	isdefault := ""
	err = validate.Var(isdefault, "isdefault")
	assert.NoError(t, err)

	lenTag := 99
	err = validate.Var(lenTag, "len=99")
	assert.NoError(t, err)

	lenTagS := "123456"
	err = validate.Var(lenTagS, "len=6")
	assert.NoError(t, err)

	oneof := 1
	err = validate.Var(oneof, "oneof=1 2")
	assert.NoError(t, err)

	required := "1"
	err = validate.Var(required, "required")
	assert.NoError(t, err)
}
