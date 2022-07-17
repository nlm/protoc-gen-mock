package protomock

import (
	"github.com/nlm/protoc-gen-mock/pkg/pb/mockpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func optionBasedScalarValueMocker(field protoreflect.FieldDescriptor) any {
	// TODO: handle field.Name == "key" && field.ContainingMessage().IsMapEntry()
	switch proto.GetExtension(field.Options(), mockpb.E_Type).(mockpb.MockFieldType) {
	case mockpb.MockFieldType_ip, mockpb.MockFieldType_ipv4:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().IPv4Address()
		}
	case mockpb.MockFieldType_ipv6:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().IPv6Address()
		}
	case mockpb.MockFieldType_size:
		return nil
	case mockpb.MockFieldType_uuid:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().UUID()
		}
	case mockpb.MockFieldType_ipnet:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().IPv4Address() + "/32"
		}
	case mockpb.MockFieldType_mac_address:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().MacAddress()
		}
	case mockpb.MockFieldType_url:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().URL()
		}
	case mockpb.MockFieldType_useragent:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().UserAgent()
		}
	case mockpb.MockFieldType_domain:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().DomainName()
		}
	case mockpb.MockFieldType_tld:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().DomainSuffix()
		}
	case mockpb.MockFieldType_email:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Email()
		}
	case mockpb.MockFieldType_file_extension:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().FileExtension()
		}
	case mockpb.MockFieldType_mime_type:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().FileMimeType()
		}
	case mockpb.MockFieldType_http_method:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().HTTPMethod()
		}
	case mockpb.MockFieldType_http_statuscode:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().HTTPStatusCode()
		}
	case mockpb.MockFieldType_http_statuscode_simple:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().HTTPStatusCodeSimple()
		}
	case mockpb.MockFieldType_hexcolor:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().HexColor()
		}
	case mockpb.MockFieldType_image:
		switch field.Kind() {
		case protoreflect.BytesKind:
			return Faker().ImageJpeg(32, 32)
		}
	case mockpb.MockFieldType_image_url:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().ImageURL(32, 32)
		}
	case mockpb.MockFieldType_log_level:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().LogLevel("syslog")
		}
	case mockpb.MockFieldType_street_address:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Street()
		}
	case mockpb.MockFieldType_street_name:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().StreetName()
		}
	case mockpb.MockFieldType_street_number:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().StreetNumber()
		}
	case mockpb.MockFieldType_street_prefix:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().StreetPrefix()
		}
	case mockpb.MockFieldType_street_suffix:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().StreetSuffix()
		}
	case mockpb.MockFieldType_city:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().City()
		}
	case mockpb.MockFieldType_country:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Country()
		}
	case mockpb.MockFieldType_country_short:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().CountryAbr()
		}
	case mockpb.MockFieldType_creditcard_cvv:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().CreditCardCvv()
		}
	case mockpb.MockFieldType_creditcard_number:
		switch field.Kind() {
		case protoreflect.StringKind:
			// return Faker().CreditCardNumber()
			return nil
		}
	case mockpb.MockFieldType_creditcard_type:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().CreditCardType()
		}
	case mockpb.MockFieldType_currency:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().CurrencyLong()
		}
	case mockpb.MockFieldType_currency_short:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().CurrencyShort()
		}
	}
	return nil
}
