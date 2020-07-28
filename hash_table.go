package ads

const (
	// hashTableInitialSize is the initial table size.
	hashTableInitialSize int = 8
	// hashTableFMultiplier is the prime p used in the Polynomial Rolling Hash algorithm.
	hashTableFMultiplier int = 53
)

// HashTable implementation using an Open Addressing strategy with Linear Probing.
type HashTable struct {
	data []*hashTableBucket
	// capacity is the number of available buckets
	capacity int
	// length is the number of used buckets
	length int
}

// hasTableBucket stores the key/value pair and a deleted flag.
type hashTableBucket struct {
	key     string
	value   interface{}
	deleted bool
}

// NewHashTable returns a newly initialized hash table.
func NewHashTable() *HashTable {
	return new(HashTable).init()
}

// init initializes hash table.
func (h *HashTable) init() *HashTable {
	h.data = make([]*hashTableBucket, hashTableInitialSize)
	h.capacity = hashTableInitialSize
	h.length = 0
	return h
}

func (h *HashTable) initLazy() {
	if h.data == nil || h.capacity == 0 {
		h.init()
	}
}

// hash returns the integer hash of a given string using Polynomial Rolling Hash algorithm.
func (h *HashTable) hash(k string) int {
	hash := 0
	for _, c := range k {
		hash = (hash*hashTableFMultiplier + int(c)) % h.capacity
	}
	return hash
}

// probe computes the linear probing sequence for a hashed number k at the i-th location.
func (h *HashTable) probe(k, i int) int {
	return (k + i) % h.capacity
}

// Get the value stored in the given key. Returns nil, false if it doesn't exit.
func (h *HashTable) Get(k string) (interface{}, bool) {
	h.initLazy()
	hash := h.hash(k)
	for i, j := 0, h.probe(hash, 0); h.data[j] != nil; i, j = i+1, h.probe(hash, i+1) {
		if !h.data[j].deleted && h.data[j].key == k {
			return h.data[j].value, true
		}
	}
	return nil, false
}

// Set or update a value using given key.
func (h *HashTable) Set(k string, v interface{}) {
	h.initLazy()
	// Maximum load is based in CPython's USABLE_FRACTION
	// https://github.com/python/cpython/blob/master/Objects/dictobject.c#L412
	if h.length >= (h.capacity<<1)/2 {
		h.resize()
	}
	hash := h.hash(k)
	deleted := -1
	j := h.probe(hash, 0)
	for i := 1; h.data[j] != nil; i++ {
		if h.data[j].deleted && deleted == -1 {
			deleted = j
		}
		if !h.data[j].deleted && h.data[j].key == k {
			break
		}
		j = h.probe(hash, i)
	}

	switch {
	case h.data[j] != nil: // Update k with v
		h.data[j].value = v
	case deleted != -1: // Found slot before NIL
		h.data[deleted].key = k
		h.data[deleted].value = v
		h.data[deleted].deleted = false
		h.length++
	default: // Bucket is NIL
		h.data[j] = &hashTableBucket{key: k, value: v, deleted: false}
		h.length++
	}
}

// resize re-allocates every (non-deleted) element in a new table.
func (h *HashTable) resize() {
	// New capacity is based in CPython's 3.4.0-3.6.0 GROWTH_RATE
	// https://github.com/python/cpython/blob/master/Objects/dictobject.c#L427
	h.capacity = (h.length * 2) + (h.capacity / 2)
	h.length = 0
	tmp := h.data
	h.data = make([]*hashTableBucket, h.capacity)
	for _, kv := range tmp {
		if kv != nil && !kv.deleted {
			h.Set(kv.key, kv.value)
		}
	}
}

// Remove the value stored at the given key
func (h *HashTable) Remove(k string) {
	h.initLazy()
	hash := h.hash(k)
	for i, j := 0, h.probe(hash, 0); h.data[j] != nil; i, j = i+1, h.probe(hash, i+1) {
		if !h.data[j].deleted && h.data[j].key == k {
			h.data[j].deleted = true
			h.data[j].value = nil // free reference to removed value
			h.length--
			return
		}
	}
}

// Size returns the number of elements stored in the hash table.
func (h *HashTable) Size() int {
	return h.length
}

// Empty removes all elements from the hash table.
func (h *HashTable) Empty() {
	h.data = nil
	h.init()
}
