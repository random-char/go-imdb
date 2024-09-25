package errors

import stderrors "errors"

var (
	ArrayExpectedError      = stderrors.New("Array type is expected")
	InvalidArrayLengthError = stderrors.New("Invalid array length")
	UnknownCommandError     = stderrors.New("Unknown command")

	InvalidNumOfArgsToSet = stderrors.New("SET requires exactly 2 arguments")
	InvalidNumOfArgsToGet = stderrors.New("GET requires exactly 1 argument")
	InvalidNumOfArgsToDel = stderrors.New("DEL requires exactly 1 argument")

	InvalidNumOfArgsToHset    = stderrors.New("HSET requires exactly 3 arguments")
	InvalidNumOfArgsToHget    = stderrors.New("HGET requires exactly 2 arguments")
	InvalidNumOfArgsToHdel    = stderrors.New("DEL requires exactly 2 arguments")
	InvalidNumOfArgsToHgetall = stderrors.New("HGETALL requires exactly 1 argument")
	InvalidNumOfArgsToHdelall = stderrors.New("HDELALL requires exactly 1 argument")

	NotFound = stderrors.New("Not found")
)
