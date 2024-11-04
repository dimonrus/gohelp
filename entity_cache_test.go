package gohelp

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type TestEntity struct {
	Id      int32    `json:"id"`
	Name    *string  `json:"name"`
	Decimal *float64 `json:"decimal"`
}

func RefreshTestEntity(id ...int32) (map[int32]*TestEntity, error) {
	if len(id) == 0 {
		for i := 1; i < GetRndNumber(5, 20); i++ {
			id = append(id, int32(i))
		}
	}
	var result = make(map[int32]*TestEntity, len(id))
	for _, key := range id {
		result[key] = &TestEntity{
			Id:      key,
			Name:    Ptr(RandString(10)),
			Decimal: Ptr(rand.Float64()),
		}
	}
	for _, entity := range result {
		fmt.Println(*entity.Name)
	}
	fmt.Println("Cache refreshed")

	return result, nil
}

var testEntityMap = map[int32]*TestEntity{
	1: {Id: 1, Name: Ptr("Первый"), Decimal: Ptr(10.01)},
	2: {Id: 2, Name: Ptr("Сегодня"), Decimal: Ptr(12.69)},
	3: {Id: 3, Name: Ptr("Возможно"), Decimal: Ptr(122.21)},
	4: {Id: 4, Name: Ptr("Классно"), Decimal: Ptr(1.07)},
	5: {Id: 5, Name: Ptr("Яма"), Decimal: Ptr(5.55)},
	6: {Id: 6, Name: Ptr("Foo"), Decimal: Ptr(17.31)},
	7: {Id: 7, Name: Ptr("Bar"), Decimal: Ptr(56.11)},
	8: {Id: 8, Name: Ptr("qux"), Decimal: Ptr(100.01)},
}

func FixedRefreshEntities(id ...int32) (map[int32]*TestEntity, error) {
	return testEntityMap, nil
}

func TestEntityCache_Idle(t *testing.T) {
	itemIds := []int32{10, 12, 13, 15}
	cache := NewEntityCache[int32, TestEntity](5, RefreshTestEntity)
	cache.SetItemIds(itemIds...)
	go func() {
		cache.Idle(context.Background())
	}()
	time.Sleep(time.Second * 10)
	ids := cache.GetItemIds()
	if len(ids) != 4 {
		t.Fatal("wrong refresh item number")
	}
	for _, id := range ids {
		if !ExistsInArray(id, itemIds) {
			t.Fatal("wrong id in item ids")
		}
	}
}

func TestEntityCache_Refresh(t *testing.T) {
	itemIds := []int32{10, 12, 13, 15}
	cache := NewEntityCache[int32, TestEntity](5, RefreshTestEntity)
	cache.SetItemIds(itemIds...)
	cache.Refresh()
	ids := cache.GetItemIds()
	if len(ids) != 4 {
		t.Fatal("wrong refresh item number")
	}
	for _, id := range ids {
		if !ExistsInArray(id, itemIds) {
			t.Fatal("wrong id in item ids")
		}
	}
	items := cache.GetAll()
	for id, item := range items {
		if id != item.Id {
			t.Fatal("wrong refresh")
		}
		it := cache.GetEntity(id)
		if *item.Name != *it.Name {
			t.Fatal("wrong item")
		}
	}
	cache.Set(nil)
	cache.SetCallback(nil)
	cache.Refresh()
	if len(cache.GetItemIds()) != 4 {
		t.Fatal("refresh must not be executed")
	}
	cache.SetCallback(RefreshTestEntity)
	cache.AddItemIds(33, 44)
	cache.Refresh()
	if len(cache.GetItemIds()) != 6 {
		t.Fatal("refresh items must be 6")
	}
}

func TestEntityCache_SortOrder(t *testing.T) {
	t.Run("sort_order", func(t *testing.T) {
		cache := NewEntityCache[int32, TestEntity](5, FixedRefreshEntities)
		cache.Refresh()
		order := cache.SortOrder(func(a, b *TestEntity) bool {
			return *a.Name < *b.Name
		})
		expected := []int32{7, 6, 8, 3, 4, 1, 2, 5}
		for i := range order {
			if expected[i] != order[i] {
				t.Fatal("wrong order")
			}
		}
	})
}
