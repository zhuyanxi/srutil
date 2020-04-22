package srutil

import "time"

// ConcurrentMap is a map using channel to garantee concurrency safe
type ConcurrentMap struct {
	m     map[string]*item
	chGet chan map[string]*item
	chAdd chan *item
}

type item struct {
	key   string
	value []byte
}

// NewConcurrentMap init a concurrency safe map
func NewConcurrentMap() *ConcurrentMap {
	cm := &ConcurrentMap{
		m:     make(map[string]*item),
		chGet: make(chan map[string]*item),
		chAdd: make(chan *item),
	}
	go cm.do()
	return cm
}

func (c *ConcurrentMap) do() {
	var k string
	var v []byte
	for {
		select {
		case it := <-c.chAdd:
			k = it.key
			v = it.value
			c.m[k] = &item{
				key:   k,
				value: v,
			}
		case c.chGet <- c.m:
		}
	}
}

// Add :
func (c *ConcurrentMap) Add(key string, value []byte) {
	time.Sleep(time.Millisecond)
	it := &item{
		key:   key,
		value: value,
	}
	c.chAdd <- it
}

// Get :
func (c *ConcurrentMap) Get(key string) ([]byte, bool) {
	cm := <-c.chGet
	v, ok := cm[key]
	if ok {
		return v.value, ok
	}
	return nil, false
}

// Len :
func (c *ConcurrentMap) Len() int {
	return len(c.m)
}
