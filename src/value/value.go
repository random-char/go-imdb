package value

import (
	"bytes"
	"go-imdb/src/errors"
	"go-imdb/src/resp/prefix"
	"strconv"
	"strings"
)

const (
	type_array byte = iota
	type_bulk
	type_string
	type_null
	type_error
)

type Value struct {
	typ   byte
	str   string
	num   int
	bulk  string
	array []Value
}

func NewStr(str string) Value {
	return Value{
		typ: type_string,
		str: str,
	}
}

func NewBulk(bulk string) Value {
	return Value{
		typ:  type_bulk,
		bulk: bulk,
	}
}

func NewArray(array []Value) Value {
	return Value{
		typ:   type_array,
		array: array,
	}
}

func NewOk() Value {
	return Value{
		typ: type_string,
		str: "OK",
	}
}

func NewNull() Value {
	return Value{
		typ: type_null,
	}
}

func NewError(err error) Value {
	return Value{
		typ: type_error,
		str: err.Error(),
	}
}

func (v *Value) GetBulk() string {
	return v.bulk
}

func (v *Value) Marshal() []byte {
	switch v.typ {
	case type_array:
		return v.marshallArray()
	case type_bulk:
		return v.marshallBulk()
	case type_string:
		return v.marshallString()
	case type_null:
		return v.marshallNull()
	case type_error:
		return v.marshallError()
	default:
		return []byte{}
	}
}

func (v *Value) ExtractCommandAndArgs() (string, []Value, error) {
	if v.typ != type_array {
		return "", nil, errors.ArrayExpectedError
	}

	command := strings.ToUpper(v.array[0].bulk)
	args := v.array[1:]

	return command, args, nil
}

func (v *Value) marshallArray() []byte {
	bb := bytes.Buffer{}

	bb.WriteByte(prefix.ARRAY)
	bb.Write([]byte(strconv.Itoa(len(v.array))))
	bb.WriteByte('\r')
	bb.WriteByte('\n')

	for _, value := range v.array {
		bb.Write(value.Marshal())
	}

	return bb.Bytes()
}

func (v *Value) marshallBulk() []byte {
	bb := bytes.Buffer{}

	bb.WriteByte(prefix.BULK)
	bb.Write([]byte(strconv.Itoa(len(v.bulk))))
	bb.WriteByte('\r')
	bb.WriteByte('\n')
	bb.Write([]byte(v.bulk))
	bb.WriteByte('\r')
	bb.WriteByte('\n')

	return bb.Bytes()
}

func (v *Value) marshallString() []byte {
	bb := bytes.Buffer{}

	bb.WriteByte(prefix.STRING)
	bb.Write([]byte(v.str))
	bb.WriteByte('\r')
	bb.WriteByte('\n')

	return bb.Bytes()
}

func (v *Value) marshallNull() []byte {
	return []byte("$-1\r\n")
}

func (v *Value) marshallError() []byte {
	bb := bytes.Buffer{}

	bb.WriteByte(prefix.ERROR)
	bb.Write([]byte(v.str))
	bb.WriteByte('\r')
	bb.WriteByte('\n')

	return bb.Bytes()
}
