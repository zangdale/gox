package xsync

import "sync"

/*
	通用 带锁的 map
	读写综合优于 sync.Map
*/
type MutexMap[K comparable, V any] struct {
	sy sync.RWMutex
	m  map[K]V
}

func NewMutexMap[K comparable, V any]() *MutexMap[K, V] {
	return &MutexMap[K, V]{
		sy: sync.RWMutex{},
		m:  make(map[K]V),
	}
}

// Store 存储一个 key 和 值（也可以用于更新）
func (m *MutexMap[K, V]) Store(key K, value V) {
	m.sy.Lock()
	defer m.sy.Unlock()
	m.m[key] = value
}

// Load 获取一个 key 的值
func (m *MutexMap[K, V]) Load(key K) (value V, ok bool) {
	m.sy.RLock()
	defer m.sy.RUnlock()
	value, ok = m.m[key]
	return value, ok
}

// LoadOrStore
// 存在返回存在的数据
// 不存在 写入新的值，并返回
func (m *MutexMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	actual, loaded = m.Load(key)
	if !loaded {
		m.Store(key, value)
		actual = value
	}
	return actual, loaded
}

// LoadAndDelete
// 存在就返回值的信息，并删除该条数据
// 不存在，返回空数据
func (m *MutexMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	value, loaded = m.Load(key)
	if loaded {
		m.Delete(key)
	}
	return value, loaded
}

// Delete 删除一条数据（不关心存在不存在）
func (m *MutexMap[K, V]) Delete(key K) {
	m.sy.Lock()
	defer m.sy.Unlock()
	delete(m.m, key)
}

func (m *MutexMap[K, V]) Range(f func(key K, value V) bool) {
	m.sy.RLock()
	defer m.sy.RUnlock()
	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
}

func (m *MutexMap[K, V]) RangeC(f func(key K, value V) bool) {
	m.sy.RLock()
	reader := make(map[K]V, len(m.m))
	for k, v := range m.m {
		reader[k] = v
	}
	m.sy.RUnlock()
	for k, v := range reader {
		if !f(k, v) {
			break
		}
	}
}
