package roles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecurity_Check(t *testing.T) {
	sec,err:=NewSecurity()
	if err != nil {
		panic(err)
	}
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read" // the operation that the user performs on the resource.
	ok, err :=sec.Check(sub,obj,act)
	if err != nil {
		panic(err)
	}
	println(ok)
	assert.False(t, ok)

	_, err = sec.e.AddPolicy("alice", "data1", "read")
	ok, err =sec.Check(sub,obj,act)
	if err != nil {
		panic(err)
	}
	println(ok)
	assert.True(t, ok)

	// Remove the added rule.
	_, err = sec.e.RemovePolicy("alice", "data1", "read")
	if err != nil {
		panic(err)
	}
	ok, err =sec.Check(sub,obj,act)
	if err != nil {
		panic(err)
	}
	println(ok)
	assert.False(t, ok)
}

/**
refs:
	- https://github.com/casbin/xorm-adapter/blob/master/adapter_test.go
 */

