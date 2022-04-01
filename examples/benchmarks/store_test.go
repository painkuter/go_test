package benchmarks

import (
	"math/rand"
	"sync"
	"testing"
)

type store struct {
	m     *sync.Mutex
	cache map[int64]int64
}

// NewStore creates and returns a new store
func NewStore() *store {
	return &store{
		m:     &sync.Mutex{},
		cache: make(map[int64]int64),
	}
}

// Get returns the value of a key
func (s *store) Get(k int64) int64 {
	s.m.Lock()
	v := s.cache[k]
	s.m.Unlock()
	return v
}

// Set stores the key value pair
func (s *store) Set(k, v int64) {
	s.m.Lock()
	s.cache[k] = v
	s.m.Unlock()
}

type syncStore struct {
	cache *sync.Map
}

// NewStore creates and returns a new store
func NewSyncStore() *syncStore {
	return &syncStore{
		cache: &sync.Map{},
	}
}

// Get returns the value of a key
func (s *syncStore) Get(k int64) int64 {
	v, ok := s.cache.Load(k)
	if !ok {
		return 0
	}
	return v.(int64)
}

// Set stores the key value pair
func (s *syncStore) Set(k, v int64) {
	s.cache.Store(k, v)
}

func BenchmarkGetStore(b *testing.B) {
	s := NewStore()
	for k, v := range genData() {
		s.Set(k, v)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var i int64
		for pb.Next() {
			i++
			_ = s.Get(i % size)
		}
	})
}

func BenchmarkGetSyncStore(b *testing.B) {
	s := NewSyncStore()
	for k, v := range genData() {
		s.Set(k, v)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var i int64
		for pb.Next() {
			i++
			_ = s.Get(i % size)
		}
	})
}

const size = 60000

func genData() map[int64]int64 {
	result := make(map[int64]int64)
	for i := 0; i < size; i++ {
		result[int64(i)] = rand.Int63()
	}
	return result
}
