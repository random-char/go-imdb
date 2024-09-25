package resp

import (
	"bufio"
	"errors"
	"go-imdb/src/resp/prefix"
	"go-imdb/src/value"
	"io"
	"strconv"
)

var unkownTypeError = errors.New("Unkown type")

type RespReader struct {
	reader *bufio.Reader
}

type RespWriter struct {
	writer io.Writer
}

func NewReader(rd io.Reader) *RespReader {
	return &RespReader{
		reader: bufio.NewReader(rd),
	}
}

func NewWriter(w io.Writer) *RespWriter {
	return &RespWriter{
		writer: w,
	}
}

func (r *RespReader) Read() (v value.Value, err error) {
	typ, err := r.reader.ReadByte()
	if err != nil {
		return v, err
	}

	switch typ {
	case prefix.ARRAY:
		return r.readArray()
	case prefix.BULK:
		return r.readBulk()
	default:
		return v, unkownTypeError
	}
}

func (w *RespWriter) Write(v value.Value) error {
	_, err := w.writer.Write(v.Marshal())

	return err
}

func (r *RespReader) readArray() (v value.Value, err error) {
	length, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	array := make([]value.Value, length)
	for i := 0; i < length; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}

		array[i] = val
	}

	return value.NewArray(array), nil
}

func (r *RespReader) readBulk() (v value.Value, err error) {
	length, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	bulk := make([]byte, length)
	r.reader.Read(bulk)

	r.readLine()

	return value.NewBulk(string(bulk)), nil
}

func (r *RespReader) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}

		n += 1
		line = append(line, b)
		if len(line) > 1 &&
			line[len(line)-2] == '\r' &&
			line[len(line)-1] == '\n' {
			break
		}
	}

	return line[:len(line)-2], n, err
}

func (r *RespReader) readInteger() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}

	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}

	return int(i64), n, err
}
