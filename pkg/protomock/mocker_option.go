package protomock

import (
	"fmt"
	"time"

	"github.com/nlm/protoc-gen-mock/pkg/pb/mockpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func optionBasedScalarValueMocker(field protoreflect.FieldDescriptor, fieldOptions proto.Message) any {
	var (
		c *mockpb.ConstFieldType
		r mockpb.MockFieldType
		t string
	)
	// Handle MapEntry
	mr := proto.GetExtension(fieldOptions, mockpb.E_Rules).(*mockpb.MockRules)
	if field.Kind() == protoreflect.EnumKind {
		// Enum
		var v protoreflect.EnumValueDescriptor
		if id := mr.GetEnum().GetId(); id != 0 {
			v = field.Enum().Values().ByNumber(protoreflect.EnumNumber(id))
		} else if name := mr.GetEnum().GetName(); name != "" {
			v = field.Enum().Values().ByName(protoreflect.Name(name))
		} else if rand := mr.GetEnum().GetRand(); rand {
			idx := Faker().Number(0, field.Enum().Values().Len()-1)
			v = field.Enum().Values().Get(idx)
		}
		if v != nil {
			return v.Number()
		}
		return nil
	}
	if field.ContainingMessage().IsMapEntry() {
		switch field.Name() {
		case "key":
			c = mr.GetMap().GetKey().GetConst()
			r = mr.GetMap().GetKey().GetMock()
			t = mr.GetMap().GetKey().GetTemplate()
		case "value":
			c = mr.GetMap().GetValue().GetConst()
			r = mr.GetMap().GetValue().GetMock()
			t = mr.GetMap().GetValue().GetTemplate()
		}
	} else {
		c = mr.GetConst()
		r = mr.GetMock()
		t = mr.GetTemplate()
	}
	// Const
	if v := getConstValueFromOption(field, c); v != nil {
		return v
	}
	// Random
	if v := getMockValueFromOption(field, r); v != nil {
		return v
	}
	// Template
	if v := getTemplateValueFromOption(field, t); v != nil {
		return v
	}
	return nil
}

func getTemplateValueFromOption(field protoreflect.FieldDescriptor, template string) any {
	if field.Kind() == protoreflect.StringKind && template != "" {
		return Faker().Generate(template)
	}
	return nil
}

func getConstValueFromOption(field protoreflect.FieldDescriptor, constField *mockpb.ConstFieldType) any {
	switch v := constField.GetValue().(type) {
	case *mockpb.ConstFieldType_Number:
		return convertNumeral(v.Number, field.Kind())
	case *mockpb.ConstFieldType_String_:
		return v.String_
	case *mockpb.ConstFieldType_Bool:
		return v.Bool
	}
	return nil
}

func getMockValueFromOption(field protoreflect.FieldDescriptor, mockType mockpb.MockFieldType) any {
	switch mockType {
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
		// FIXME
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
			return Faker().CreditCardNumber(nil)
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
	case mockpb.MockFieldType_phone_number:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Phone()
		}
	case mockpb.MockFieldType_language:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Language()
		}
	case mockpb.MockFieldType_language_short:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().LanguageAbbreviation()
		}
	case mockpb.MockFieldType_language_bcp:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().LanguageBCP()
		}
	case mockpb.MockFieldType_latitude:
		switch field.Kind() {
		case protoreflect.StringKind:
			return fmt.Sprint(Faker().Latitude())
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			return convertNumeral(Faker().Latitude(), field.Kind())
		}
	case mockpb.MockFieldType_longitude:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Longitude()
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			return convertNumeral(Faker().Longitude(), field.Kind())
		}
	case mockpb.MockFieldType_timezone:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().TimeZone()
		}
	case mockpb.MockFieldType_timezone_short:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().TimeZoneAbv()
		}
	case mockpb.MockFieldType_timezone_offset:
		switch field.Kind() {
		case protoreflect.StringKind:
			return fmt.Sprint(Faker().TimeZoneOffset())
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
			return convertNumeral(Faker().TimeZoneOffset(), field.Kind())
		}
	case mockpb.MockFieldType_timezone_region:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().TimeZoneRegion()
		}
	case mockpb.MockFieldType_first_name:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().FirstName()
		}
	case mockpb.MockFieldType_last_name:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().LastName()
		}
	case mockpb.MockFieldType_full_name:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Name()
		}
	case mockpb.MockFieldType_name_prefix:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().NamePrefix()
		}
	case mockpb.MockFieldType_us_ssn:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().SSN()
		}
	case mockpb.MockFieldType_us_state:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().State()
		}
	case mockpb.MockFieldType_us_state_short:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().StateAbr()
		}
	case mockpb.MockFieldType_date:
		switch field.Kind() {
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
			return convertNumeral(Faker().Date().Unix(), field.Kind())
		}
	case mockpb.MockFieldType_year:
		switch field.Kind() {
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind,
			protoreflect.StringKind:
			return convertNumeral(Faker().Year(), field.Kind())
		}
	case mockpb.MockFieldType_hour:
		switch field.Kind() {
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind,
			protoreflect.StringKind:
			return convertNumeral(Faker().Minute(), field.Kind())
			return convertNumeral(Faker().Hour(), field.Kind())
		}
	case mockpb.MockFieldType_minute:
		switch field.Kind() {
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind,
			protoreflect.StringKind:
			return convertNumeral(Faker().Minute(), field.Kind())
		}
	case mockpb.MockFieldType_second:
		switch field.Kind() {
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind,
			protoreflect.StringKind:
			return convertNumeral(Faker().Second(), field.Kind())
		}
	case mockpb.MockFieldType_month:
		switch field.Kind() {
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind,
			protoreflect.StringKind:
			return convertNumeral(Faker().Month(), field.Kind())
		}
	case mockpb.MockFieldType_month_string:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().MonthString()
		}
	case mockpb.MockFieldType_nanosecond:
		switch field.Kind() {
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind,
			protoreflect.StringKind:
			return convertNumeral(Faker().NanoSecond(), field.Kind())
		}
	case mockpb.MockFieldType_date_future:
		switch field.Kind() {
		case protoreflect.StringKind:
			return fmt.Sprint(Faker().DateRange(time.Now(), time.Now().Add(365*24*time.Hour)), field.Kind())
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
			return convertNumeral(Faker().DateRange(time.Now(), time.Now().Add(365*24*time.Hour)).Unix(), field.Kind())
		}
	case mockpb.MockFieldType_date_past:
		switch field.Kind() {
		case protoreflect.StringKind:
			return fmt.Sprint(Faker().DateRange(time.Now().Add(-365*24*time.Hour), time.Now()), field.Kind())
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
			return convertNumeral(Faker().DateRange(time.Now().Add(-365*24*time.Hour), time.Now()).Unix(), field.Kind())
		}
	case mockpb.MockFieldType_date_now:
		switch field.Kind() {
		case protoreflect.StringKind:
			return fmt.Sprint(time.Now())
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
			return convertNumeral(time.Now().Unix(), field.Kind())
		}
	case mockpb.MockFieldType_number:
		switch field.Kind() {
		case protoreflect.DoubleKind, protoreflect.FloatKind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
			protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
			protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
			protoreflect.Fixed64Kind, protoreflect.Uint64Kind,
			protoreflect.StringKind:
			return convertNumeral(Faker().Number(numberLowValue, numberHighValue), field.Kind())
		}
	case mockpb.MockFieldType_author:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().AppAuthor()
		}
	case mockpb.MockFieldType_app_name:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().AppName()
		}
	case mockpb.MockFieldType_version:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().AppName()

		}
	case mockpb.MockFieldType_color:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Color()
		}
	case mockpb.MockFieldType_emoji:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Emoji()
		}
	case mockpb.MockFieldType_emoji_alias:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().EmojiAlias()
		}
	case mockpb.MockFieldType_lorem_ipsum:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().LoremIpsumSentence(loremIpsumWordCount)
		}
	case mockpb.MockFieldType_password:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Password(true, true, true, true, false, 16)
		}
	case mockpb.MockFieldType_phrase:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Phrase()
		}
	case mockpb.MockFieldType_price:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Price(100, 200)
		}
	case mockpb.MockFieldType_job_title:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().JobTitle()
		}
	}
	return nil
}
