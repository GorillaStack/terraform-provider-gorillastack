// This is from the aws-sdk-go https://raw.githubusercontent.com/aws/aws-sdk-go/master/aws/awsutil/prettify.go
// To help with serialization

package util

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
)

var callStringMethodOn = []string{
	"gorillastack.StringArrayOrNull",
	"gorillastack.IntOrString",
}

// StringValue returns the string representation of a value.
func StringValue(i interface{}) string {
	var buf bytes.Buffer
	stringValue(reflect.ValueOf(i), &buf)
	return buf.String()
}

func decapitalize(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

// stringValue will recursively walk value v to build a textual
// representation of the value.
func stringValue(v reflect.Value, buf *bytes.Buffer) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		strtype := v.Type().String()
		if strtype == "time.Time" {
			fmt.Fprintf(buf, "%s", v.Interface())
			break
		} else if strings.HasPrefix(strtype, "io.") {
			buf.WriteString("<buffer>")
			break
		} else if Contains(callStringMethodOn, strtype) {
			tmp := v.MethodByName("String").Call([]reflect.Value{})
			buf.WriteString(tmp[0].String())
			break
		}

		buf.WriteString("{")

		names := []string{}
		for i := 0; i < v.Type().NumField(); i++ {
			name := v.Type().Field(i).Name
			f := v.Field(i)
			if name[0:1] == strings.ToLower(name[0:1]) {
				continue // ignore unexported fields
			}
			if (f.Kind() == reflect.Ptr || f.Kind() == reflect.Slice || f.Kind() == reflect.Map) && f.IsNil() {
				continue // ignore unset fields
			}
			names = append(names, name)
		}

		for i, n := range names {
			val := v.FieldByName(n)
			buf.WriteString("\"" + decapitalize(n) + "\":")
			stringValue(val, buf)

			if i < len(names)-1 {
				buf.WriteString(",")
			}
		}

		buf.WriteString("}")
	case reflect.Slice:
		strtype := v.Type().String()
		if strtype == "[]uint8" {
			fmt.Fprintf(buf, "<binary> len %d", v.Len())
			break
		}

		buf.WriteString("[")
		for i := 0; i < v.Len(); i++ {
			stringValue(v.Index(i), buf)

			if i < v.Len()-1 {
				buf.WriteString(",")
			}
		}

		buf.WriteString("]")
	case reflect.Map:
		buf.WriteString("{")

		for i, k := range v.MapKeys() {
			buf.WriteString("\"" + decapitalize(k.String()) + "\":")
			stringValue(v.MapIndex(k), buf)

			if i < v.Len()-1 {
				buf.WriteString(",")
			}
		}

		buf.WriteString("}")
	default:
		if !v.IsValid() {
			fmt.Fprint(buf, "<invalid value>")
			return
		}
		format := "%v"
		switch v.Interface().(type) {
		case string:
			format = "%q"
		case io.ReadSeeker, io.Reader:
			format = "buffer(%p)"
		}
		fmt.Fprintf(buf, format, v.Interface())
	}
}
