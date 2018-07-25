package src

import "sync"

var (
	cache = make(map[interface{}]*List)
	m     sync.RWMutex
)

func Cache(list string) *List {
	m.RLock()
	l, ok := cache[list]
	m.RUnlock()

	if ok {
		return l
	}

	m.Lock()
	defer m.Unlock()
	l = &List{
		key:   list,
		nodes: make(map[interface{}]*Node),
	}
	cache[list] = l

	return l
}
