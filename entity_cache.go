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
	// Maximum failed refresh attempts before panic
	maxFailedCount uint16
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

// RefreshUntilSuccess refresh until c.maxFailedCount
func (c *EntityCache[E, T]) RefreshUntilSuccess() error {
	var failedCount uint16
	for {
		err := c.Refresh()
		if err == nil {
			return nil
		}
		if failedCount >= c.maxFailedCount {
			return err
		}
		failedCount++
		time.Sleep(time.Second * time.Duration(c.refreshPeriod))
	}
}

// Idle wait for next time refresh event
func (c *EntityCache[E, T]) Idle(ctx context.Context) error {
	duration := time.Second * time.Duration(c.refreshPeriod)
	ticker := time.NewTicker(duration)
	var failedCount uint16
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			e := c.Refresh()
			if e != nil {
				failedCount++
				if failedCount >= c.maxFailedCount {
					panic(e)
				}
			} else {
				failedCount = 0
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

// UnSetEntity unset single entity
func (c *EntityCache[E, T]) UnSetEntity(id E) {
	c.m.Lock()
	delete(c.entityMap, id)
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

// UnSetItemIds unset item ids
func (c *EntityCache[E, T]) UnSetItemIds(ids ...E) {
	c.m.Lock()
	for _, id := range ids {
		i := Index[E](id, c.ids)
		if i != -1 {
			if i == 0 {
				c.ids = c.ids[1:]
			} else if i == len(c.ids)-1 {
				c.ids = c.ids[:i]
			} else {
				c.ids = append(c.ids[:i], c.ids[i+1:]...)
			}
		}
	}
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

// Map iterator through cache items
// next - single item callback iterator
// ordered - means that ids define order for item iterations
func (c *EntityCache[E, T]) Map(next func(id E, item *T), ordered bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	if ordered {
		for _, id := range c.ids {
			next(id, c.entityMap[id])
		}
	} else {
		for e, t := range c.entityMap {
			next(e, t)
		}
	}
	return
}

// NewEntityCache create new entity cache
// E - identifier type
// T - type of entity cache
// refreshPeriod - how often will be refresh executed
// callback - refresh function
func NewEntityCache[E comparable, T any](refreshPeriod, maxFailedCount uint16, callback RefreshEntityCallback[E, T], id ...E) *EntityCache[E, T] {
	return &EntityCache[E, T]{
		entityMap:      make(map[E]*T),
		refresh:        callback,
		ids:            id,
		refreshPeriod:  refreshPeriod,
		maxFailedCount: maxFailedCount,
	}
}
