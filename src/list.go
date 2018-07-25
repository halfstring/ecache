package src

import (
	"sync"
	"time"
)

type List struct {
	sync.RWMutex
	key   interface{}
	nodes map[interface{}]*Node
}

func (list *List) Count() int {
	list.Lock()
	defer list.Unlock()

	return len(list.nodes)
}

func (list *List) Foreach(f func(key interface{}, node *Node)) {
	list.Lock()
	defer list.Unlock()
	for k, v := range list.nodes {
		f(k, v)
	}
}

func (list *List) Add(key interface{}, value interface{}, lifeTime time.Duration) *Node {
	node := NewNode(key, value, lifeTime)

	list.Lock()
	defer list.Unlock()
	list.nodes[node.key] = node
	return node
}

func (list *List) Value(key interface{}, args ... interface{} )( *Node, error) {
	list.RLock()
	node, ok := list.nodes[key]
	if !ok {
		return nil, ErrKeyNotFound
		panic("Key no exist....")
	}

	return node, nil
}
