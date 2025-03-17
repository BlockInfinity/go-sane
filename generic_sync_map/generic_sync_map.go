package generic_sync_map

import (
	"sync"
)

// GenericSyncMap is a type-safe, generic wrapper around sync.Map.
type GenericSyncMap[K comparable, V any] struct {
	m sync.Map
}

// Set() stores a key-value pair in the map.
func (s *GenericSyncMap[K, V]) Set(key K, value V) {
	s.m.Store(key, value)
}

// Get() retrieves a value from the map, ensuring type safety.
func (s *GenericSyncMap[K, V]) Get(key K) (V, bool) {
	value, ok := s.m.Load(key)
	if !ok {
		var zero V // Return zero value of V
		return zero, false
	}
	v, ok := value.(V)
	return v, ok
}

// Delete() removes a key from the map.
func (s *GenericSyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)
}

// Range() iterates over all key-value pairs with a type-safe callback.
func (s *GenericSyncMap[K, V]) Range(f func(key K, value V) bool) {
	s.m.Range(func(k, v interface{}) bool {
		key, ok1 := k.(K)
		value, ok2 := v.(V)
		if !ok1 || !ok2 {
			panic("Unexcpeted type found in generic sync map's Range() function.")
		}
		return f(key, value)
	})
}

// ToMap() converts the GenericSyncMap to a standard map[K]V.
func (s *GenericSyncMap[K, V]) ToMap() map[K]V {
	regularMap := make(map[K]V)
	s.Range(func(key K, value V) bool {
		regularMap[key] = value
		return true
	})
	return regularMap
}
