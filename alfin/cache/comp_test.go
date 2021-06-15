package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestNewContext(t *testing.T) {
	ctx:=context.Background()
	ctx=NewContext(ctx)

	cache:=FromContext(ctx)
	cache.Put("just a item", []string{"values"})
	val:=cache.Get("just a item").([]string)
	fmt.Printf("%s\n", val)
}
