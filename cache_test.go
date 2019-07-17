package cache

import (
	`fmt`
	`testing`
	"time"
)

func TestCache_Add(t *testing.T) {
	c := New(DefaultExpiration,DefaultExpiration)
	c.Set("foo","bar",time.Second *2)
	c.Set("baz", 42, NoExpiration)
	fmt.Println(c.items)

	time.Sleep(time.Second *6)

	fmt.Println(c.items)
}
