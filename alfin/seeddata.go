package alfin

import "fmt"

// ClientKey is the key for lookup
type ClientKey int

const (
	// CadenceClientKey for retrieving cadence client from context
	CadenceClientKey ClientKey = iota
	WorkloadStoreKey
	SeedsKey
)

type SeedElements struct{
	elements map[string]interface{}
}

func NewSeedElements() *SeedElements {
	return &SeedElements{elements: make(map[string]interface{})}
}

func (t *SeedElements) Put(key string, value interface{}) interface{} {
	t.elements[key]=value
	return value
}

func (t *SeedElements) MustGet(key string) interface{}{
	v,ok:=t.elements[key]
	if !ok{
		panic(fmt.Errorf("Cannot get value from key %s", key))
	}
	return v
}

func (t *SeedElements) Size() int{
	return len(t.elements)
}

