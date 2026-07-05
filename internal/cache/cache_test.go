package cache

import (
	"testing"
	"time"
)

func TestCacheSetAndGet(t *testing.T) {
	c := NewCache()

	c.Set("users:id:1", "test-value", time.Minute)

	value, ok := c.Get("users:id:1")
	if !ok {
		t.Fatal("expected value in cache")
	}

	if value != "test-value" {
		t.Fatalf("expected test-value, got %v", value)
	}
}

func TestCacheGetMissingKey(t *testing.T) {
	c := NewCache()

	value, ok := c.Get("missing-key")
	if ok {
		t.Fatal("expected missing key")
	}

	if value != nil {
		t.Fatalf("expected nil value, got %v", value)
	}
}

func TestCacheExpiredItem(t *testing.T) {
	c := NewCache()

	c.Set("users:id:1", "test-value", 10*time.Millisecond)

	time.Sleep(20 * time.Millisecond)

	value, ok := c.Get("users:id:1")
	if ok {
		t.Fatal("expected expired value")
	}

	if value != nil {
		t.Fatalf("expected nil value, got %v", value)
	}
}
