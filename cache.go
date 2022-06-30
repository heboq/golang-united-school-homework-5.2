package cache

import (
	"time"
)

type Cache struct {
	Data map[string]*entry
}

type entry struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {
	return Cache{Data: make(map[string]*entry)}
}

func (c Cache) Get(key string) (string, bool) {
	match, ok := c.Data[key]
	if !ok {
		return "", false
	}
	if match.deadline.IsZero() {
		return match.value, true
	} else if !isExpired(match) {
		return match.value, true
	}
	return "", false
}

func (c Cache) Put(key, value string) {
	if match, ok := c.Data[key]; ok {
		match.value = value
	} else {
		c.Data[key] = &entry{value: value}
	}
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.Data))
	for key, entry := range c.Data {
		if entry.deadline.IsZero() {
			keys = append(keys, key)
		} else if !isExpired(entry) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	if match, ok := c.Data[key]; ok {
		match.value = value
	} else {
		c.Data[key] = &entry{value: value, deadline: deadline}
	}
}

func isExpired(entry *entry) bool {
	return !time.Now().Before(entry.deadline)
}
