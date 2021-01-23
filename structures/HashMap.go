package structures

import "hash/fnv"

type HashMap struct {
	size    int
	count   int
	table []*LinkedList
}

func NewHashMap() *HashMap {
	return &HashMap{
		table: nil,
	}
}

// LNode singly linked list node
type HashMapNode struct {
	key uint32
	value string
}

func NewHashMapNode(value string) *HashMapNode {
	return &HashMapNode{
		key: hash(value),
		value: value,
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}


// Len returns the count of the elements in the hashmap
func (h *HashMap) Len() int {
	return h.count
}

// Size returns the size of the hashamp
func (h *HashMap) Size() int {
	return h.size
}