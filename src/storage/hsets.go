package storage

import "sync"

var hsets = map[string]map[string]string{}
var hsetsMu = sync.RWMutex{}

func Hset(h, k, v string) {
	hsetsMu.Lock()

	_, ok := hsets[h]
	if !ok {
		hsets[h] = make(map[string]string)
	}

	hsets[h][k] = v

	hsetsMu.Unlock()
}

func Hget(h, k string) (string, bool) {
	m, ok := hsets[h]
    if !ok {
        return "", false
    }

    v, ok := m[k]

	return v, ok
}

func Hdel(h, k string) {
    m, ok := hsets[h]
    if ok {
        delete(m, k)
    }
}

func Hgetall(h string) (map[string]string, bool) {
	m, ok := hsets[h]

	return m, ok
}

func Hdelall(h string) {
    delete(hsets, h)
}
