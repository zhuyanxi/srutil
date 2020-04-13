package srutil

import "sync"

type ConcurrentMap struct {
	sync.RWMutex
	m map[string]struct{}
}
