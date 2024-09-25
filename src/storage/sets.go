package storage

import "sync"

var sets = map[string]string{}
var setsMu = sync.RWMutex{}

func Set(k, v string) {
    setsMu.Lock()
    sets[k] = v
    setsMu.Unlock()
}

func Get(k string) (string, bool) {
    v, ok := sets[k]

    return v, ok
}

func Del(k string) {
    delete(sets, k)
}
