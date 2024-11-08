package gohelp

import (
	"context"
	"sort"
	"sync"
	"time"
)

// RefreshEntityCallback refresh entity callback
type RefreshEntityCallback[E comparable, T any] func(id ...E) (map[E]*T, error)

// EntityCache entity data cache
type EntityCache[E comparable, T any] struct {
	// entity map
	entityMap map[E]*T
	// sync mutex
	m sync.RWMutex
	// refresh callback
	refresh RefreshEntityCallback[E, T]
	// entity ids
	ids []E
	// Period until all entity will be refreshed
	refreshPeriod uint16
}

// Refresh cache function
func (c *EntityCache[E, T]) Refresh() error {
	if c.refreshPeriod == 0 || c.refresh == nil {
		return nil
	}
	var err error
	var items map[E]*T
	ids := c.GetItemIds()
	if len(ids) > 0 {
		items, err = c.refresh(ids...)
	} else {
		items, err = c.refresh()
	}
	if err != nil {
		return err
	}
	c.Set(items)
	return nil
}

// Idle wait for next time refresh event
func (c *EntityCache[E, T]) Idle(ctx context.Context) error {
	e := c.Refresh()
	if e != nil {
		return e
	}
	duration := time.Second * time.Duration(c.refreshPeriod)
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			e = c.Refresh()
			if e != nil {
				return e
			}
		}
	}
}

// SortOrder get items sort order according to sort func
func (c *EntityCache[E, T]) SortOrder(s func(a, b *T) bool) []E {
	c.m.RLock()
	defer c.m.RUnlock()
	var keys = make([]E, len(c.entityMap))
	var k int
	for key := range c.entityMap {
		keys[k] = key
		k++
	}
	sort.Slice(keys, func(i, j int) bool {
		return s(c.entityMap[keys[i]], c.entityMap[keys[j]])
	})
	return keys
}

// Len return len of entity map
func (c *EntityCache[E, T]) Len() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return len(c.entityMap)
}

// Set entity map
func (c *EntityCache[E, T]) Set(entityMap map[E]*T) {
	c.m.Lock()
	c.entityMap = entityMap
	c.m.Unlock()
	return
}

// SetEntity set single entity
func (c *EntityCache[E, T]) SetEntity(id E, item *T) {
	c.m.Lock()
	c.entityMap[id] = item
	c.m.Unlock()
	return
}

// GetAll get all items
func (c *EntityCache[E, T]) GetAll() map[E]*T {
	return c.entityMap
}

// GetEntity get single entity
func (c *EntityCache[E, T]) GetEntity(id E) *T {
	c.m.RLock()
	v, ok := c.entityMap[id]
	c.m.RUnlock()
	if ok {
		return v
	}
	return nil
}

// SetItemIds set entity ids for refresh
func (c *EntityCache[E, T]) SetItemIds(id ...E) {
	c.m.Lock()
	c.ids = id
	c.m.Unlock()
	return
}

// AddItemIds add entity ids for refresh
func (c *EntityCache[E, T]) AddItemIds(id ...E) {
	c.m.Lock()
	c.ids = AppendUnique(c.ids, id...)
	c.m.Unlock()
	return
}

// GetItemIds get refresh entity ids
func (c *EntityCache[E, T]) GetItemIds() []E {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.ids
}

// SetCallback set refresh callback
// in case when nil cache will be not refreshed
func (c *EntityCache[E, T]) SetCallback(callback RefreshEntityCallback[E, T]) {
	c.m.Lock()
	c.refresh = callback
	c.m.Unlock()
}

// NewEntityCache create new entity cache
func NewEntityCache[E comparable, T any](refreshPeriod uint16, callback RefreshEntityCallback[E, T], id ...E) *EntityCache[E, T] {
	return &EntityCache[E, T]{
		entityMap:     make(map[E]*T),
		refresh:       callback,
		ids:           id,
		refreshPeriod: refreshPeriod,
	}
}
