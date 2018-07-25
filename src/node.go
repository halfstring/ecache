package src

import (
	"sync"
	"time"
)

type Node struct {
	sync.RWMutex
	key        interface{}
	value      interface{}
	lifeTime   time.Duration
	createTime time.Time
}

func NewNode(key interface{}, value interface{}, lifeTime time.Duration) *Node {
	return &Node{
		key:        key,
		value:      value,
		lifeTime:   lifeTime,
		createTime: time.Now(),
	}
}

func (node *Node) Key() interface{} {
	return node.key
}

func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) LifeTime() time.Duration {
	return node.lifeTime
}

func (node *Node) CreateTime() time.Time {
	return node.createTime
}
