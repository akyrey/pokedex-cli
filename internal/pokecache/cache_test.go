package pokecache

import (
	"testing"
	"time"
)

func TestGetCache(t *testing.T) {
	t.Run("should find a previously inserted value in the cache", func(t *testing.T) {
		cache := NewCache(10 * time.Minute)
		cache.Add("test", []byte("test"))

		val, ok := cache.Get("test")
		if !ok {
			t.Errorf("Expected cache to have test key")
		}
		if string(val) != "test" {
			t.Errorf("Expected cache to have test value")
		}
	})

	t.Run("shouldn't find a new value from the cache", func(t *testing.T) {
		cache := NewCache(10 * time.Minute)

		val, ok := cache.Get("test")
		if ok {
			t.Errorf("Expected cache to NOT have test key")
		}
		if val != nil {
			t.Errorf("Expected cache value to be nil")
		}
	})
}

func TestReapLoop(t *testing.T) {
	t.Run("should remove expired values from the cache", func(t *testing.T) {
		cache := NewCache(50 * time.Millisecond)
		cache.Add("test", []byte("test"))

		time.Sleep(100 * time.Millisecond)

		val, ok := cache.Get("test")
		if ok {
			t.Errorf("Expected cache to NOT have test key")
		}
		if val != nil {
			t.Errorf("Expected cache value to be nil")
		}
	})

	t.Run("shouldn't remove not expired values from cache", func(t *testing.T) {
		cache := NewCache(10 * time.Second)
		cache.Add("test", []byte("test"))

		val, ok := cache.Get("test")
		if !ok {
			t.Errorf("Expected cache to have test key")
		}
		if string(val) != "test" {
			t.Errorf("Expected cache to have test value")
		}
	})
}
