package main

import (
	"fmt"
	"go-imdb/src/errors"
	"go-imdb/src/handler"
	"go-imdb/src/resp"
	"go-imdb/src/value"
	"net"
	"os"
	"strconv"
)

var port = 6379

func init() {
	var err error

	portString := os.Getenv("LISTEN_PORT")
	if portString != "" {
		port, err = strconv.Atoi(portString)

		if err != nil || port < 0 || port > 65535 {
			fmt.Fprintln(os.Stderr, "Invalid port set in env")
			os.Exit(1)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer conn.Close()

	respReader := resp.NewReader(conn)
	respWriter := resp.NewWriter(conn)

	for {
		v, err := respReader.Read()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		command, args, err := v.ExtractCommandAndArgs()
		if err != nil {
			respWriter.Write(value.NewError(errors.ArrayExpectedError))
			continue
		}

		handler, ok := handler.Handlers[command]
		if !ok {
			respWriter.Write(value.NewError(errors.UnknownCommandError))
			continue
		}

		result := handler(args)
		respWriter.Write(result)
	}
}
