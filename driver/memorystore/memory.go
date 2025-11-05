package memorystore

import (
	"errors"
	"sync"
	"time"
)

type item struct {
	value     string
	expiredAt time.Time
}

type Memory struct {
	data map[string]item
	mu   sync.RWMutex
}

// memory store
func New() *Memory {
	c := &Memory{
		data: make(map[string]item),
	}
	go c.cleaner()
	return c
}

func (m *Memory) Set(key, value string, ttl time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = item{
		value:     value,
		expiredAt: time.Now().Add(ttl),
	}
	return nil
}

func (m *Memory) Get(key string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	it, ok := m.data[key]
	if !ok || time.Now().After(it.expiredAt) {
		return "", errors.New("not found")
	}
	return it.value, nil
}

func (m *Memory) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
	return nil
}

func (m *Memory) cleaner() {
	for {
		time.Sleep(30 * time.Second)
		m.mu.Lock()
		for k, it := range m.data {
			if time.Now().After(it.expiredAt) {
				delete(m.data, k)
			}
		}
		m.mu.Unlock()
	}
}
