package alfin

import "encoding/json"

// StringPtr Trivial helper to get is a string pointer from a string (because you're not
// allowed to take the address of a literal directly).
func StringPtr(s string) *string {
	return &s
}

func Pretty(v interface{}) {
	r, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(r))
}
