package main

type KeyValueDatabase interface {
	// Get get a key's value
	//@timeout 12s
	Get(key string) string
	Set(key string, value string)
	Inspect(ctx map[string]interface{})
}

type KeyValueDatabaseMock struct {
	values map[string]string
}

func (db *KeyValueDatabaseMock) Get(key string) string {
	return db.values[key]
}
func (db *KeyValueDatabaseMock) Set(key string, value string) {
	db.values[key] = value
}
