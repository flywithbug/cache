package cache

import (
	`fmt`
	`testing`
)

func TestCache_Add(t *testing.T) {
	c := New(DefaultExpiration,DefaultExpiration)
	c.Set("foo","bar",DefaultExpiration)
	c.Set("baz", 42, NoExpiration)

	fmt.Println(c.items)
}
