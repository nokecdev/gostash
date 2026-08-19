package cache_test

import (
	"cache"
	"testing"
	"time"
)

func TestExpired(t *testing.T) {
	//Initialize test cases
	tests := []struct {
		name       string
		expiration int64
		expected   bool
	}{
		{
			name:       "not expired",
			expiration: time.Now().Add(time.Hour).UnixNano(),
			expected:   false,
		},
		{
			name:       "expired",
			expiration: time.Now().Add(-time.Hour).UnixNano(),
			expected:   true,
		},
		{
			name:       "no expiration",
			expiration: 0,
			expected:   false,
		},
	}

	// Create and run subtests
	for _, tt := range tests {
		t.Run(tt.name, func(st *testing.T) {
			item := cache.Item{
				Expiration: tt.expiration,
			}

			if got := item.Expired(); got != tt.expected {
				st.Errorf(
					"Expired() = %v, expected %v",
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestCacheExpiry(t *testing.T) {
	c := cache.NewCache(
		10*time.Millisecond,
		5*time.Millisecond,
	)

	c.Set("foo", "bar", time.Millisecond)

	// value, exists = value not needed, test only the existence
	if _, ok := c.Get("foo"); !ok {
		t.Fatal("expected item to exist")
	}

	time.Sleep(20 * time.Millisecond)

	if _, ok := c.Get("foo"); ok {
		t.Fatal("expected item to be expired")
	}
}
