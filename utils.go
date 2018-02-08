package invoicexpress

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	null "gopkg.in/guregu/null.v3"

	"github.com/gorilla/schema"
)

type SchemaMarshaler interface {
	MarshalSchema() string
}

type ToURLValues interface {
	ToURLValues() url.Values
}

func AddQueryParamsToRequest(requestParams interface{}, req *http.Request, skipEmpty bool) error {
	params := url.Values{}

	to, ok := requestParams.(ToURLValues)
	if ok == true {
		params = to.ToURLValues()
	} else {
		encoder := newSchemaEncoder()
		err := encoder.Encode(requestParams, params)
		if err != nil {
			return err
		}
	}

	query := req.URL.Query()
	for k, vals := range params {
		for _, v := range vals {
			if skipEmpty && v == "" {
				continue
			}

			if skipEmpty && v == "0" {
				continue
			}

			query.Add(k, v)
		}
	}

	req.URL.RawQuery = query.Encode()
	// force [ & ] in query parameters
	// req.URL.RawQuery = strings.Replace(req.URL.RawQuery, "%5B", "[", -1)
	// req.URL.RawQuery = strings.Replace(req.URL.RawQuery, "%5D", "]", -1)
	return nil
}

func newSchemaEncoder() *schema.Encoder {
	encoder := schema.NewEncoder()

	// register custom encoders
	encodeSchemaMarshaler := func(v reflect.Value) string {
		marshaler, ok := v.Interface().(SchemaMarshaler)
		if ok == true {
			return marshaler.MarshalSchema()
		}

		stringer, ok := v.Interface().(fmt.Stringer)
		if ok == true {
			return stringer.String()
		}

		return ""
	}

	encodeNullFloat := func(v reflect.Value) string {
		nullFloat, _ := v.Interface().(null.Float)
		if nullFloat.IsZero() {
			return ""
		}
		return strconv.FormatFloat(nullFloat.Float64, 'f', 6, 64)
	}

	encodeNullBool := func(v reflect.Value) string {
		nullBool, _ := v.Interface().(null.Bool)
		if nullBool.IsZero() {
			return ""
		}
		return strconv.FormatBool(nullBool.Bool)
	}

	encoder.RegisterEncoder(InvoiceDocumentType{}, encodeSchemaMarshaler)
	encoder.RegisterEncoder(Date{}, encodeSchemaMarshaler)
	encoder.RegisterEncoder(null.Float{}, encodeNullFloat)
	encoder.RegisterEncoder(null.Bool{}, encodeNullBool)
	return encoder
}
