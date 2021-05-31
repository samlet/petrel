package alfin

// StringPtr Trivial helper to get is a string pointer from a string (because you're not
// allowed to take the address of a literal directly).
func StringPtr(s string) *string {
	return &s
}
