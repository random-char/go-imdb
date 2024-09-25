package handler

import (
	"go-imdb/src/errors"
	"go-imdb/src/storage"
	"go-imdb/src/value"
	"strings"
)

const (
	command_ping    = "PING"
	command_set     = "SET"
	command_get     = "GET"
	command_del     = "DEL"
	command_hset    = "HSET"
	command_hget    = "HGET"
	command_hdel    = "HDEL"
	command_hgetall = "HGETALL"
	command_hdelall = "HDELALL"
)

var Handlers = map[string]func([]value.Value) value.Value{
	command_ping:    ping,
	command_set:     set,
	command_get:     get,
	command_del:     del,
	command_hset:    hset,
	command_hget:    hget,
	command_hdel:    hdel,
	command_hgetall: hgetall,
	command_hdelall: hdelall,
}

func ping(args []value.Value) value.Value {
	str := "PONG"
	if len(args) != 0 {
		sb := strings.Builder{}

		for _, arg := range args {
			sb.WriteString(arg.GetBulk())
			sb.WriteByte(' ')
		}

		str = sb.String()
	}

	return value.NewStr(str)
}

func set(args []value.Value) value.Value {
	if len(args) != 2 {
		return value.NewError(errors.InvalidNumOfArgsToSet)
	}

	k := args[0].GetBulk()
	v := args[1].GetBulk()

	storage.Set(k, v)

	return value.NewOk()
}

func get(args []value.Value) value.Value {
	if len(args) != 1 {
		return value.NewError(errors.InvalidNumOfArgsToGet)
	}

	str, ok := storage.Get(args[0].GetBulk())
	if !ok {
		return value.NewNull()
	}

	return value.NewStr(str)
}

func del(args []value.Value) value.Value {
    if len(args) != 1 {
    return value.NewError(errors.InvalidNumOfArgsToDel)
    }

    k := args[0].GetBulk()
    storage.Del(k)

    return value.NewOk()
}

func hset(args []value.Value) value.Value {
	if len(args) != 3 {
		return value.NewError(errors.InvalidNumOfArgsToHset)
	}

	h := args[0].GetBulk()
	k := args[1].GetBulk()
	v := args[2].GetBulk()

	storage.Hset(h, k, v)

	return value.NewOk()
}

func hget(args []value.Value) value.Value {
	if len(args) != 2 {
		return value.NewError(errors.InvalidNumOfArgsToHget)
	}

	h := args[0].GetBulk()
	k := args[1].GetBulk()

	v, ok := storage.Hget(h, k)
	if !ok {
		return value.NewNull()
	}

	return value.NewStr(v)
}

func hdel(args []value.Value) value.Value {
    if len(args) != 2 {
        return value.NewError(errors.InvalidNumOfArgsToHdel)
    }

    h := args[0].GetBulk()
    k := args[1].GetBulk()

    storage.Hdel(h, k)

    return value.NewOk()
}

func hgetall(args []value.Value) value.Value {
	if len(args) != 1 {
		return value.NewError(errors.InvalidNumOfArgsToHgetall)
	}

	h := args[0].GetBulk()

	m, ok := storage.Hgetall(h)
	if !ok {
		return value.NewNull()
	}

	resultArray := make([]value.Value, len(m)*2)
	var kValue, vValue value.Value

	i := 0
	for k, v := range m {
		kValue = value.NewStr(k)
		vValue = value.NewStr(v)

		resultArray[i] = kValue
		resultArray[i+1] = vValue

		i += 2
	}

	return value.NewArray(resultArray)
}

func hdelall(args []value.Value) value.Value {
    if len(args) != 1 {
        return value.NewError(errors.InvalidNumOfArgsToHdelall)
    }

    h := args[0].GetBulk()
    storage.Hdelall(h)

    return value.NewOk()
}

