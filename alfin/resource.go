package alfin

type APIResponse struct {
	// Status is a status code and message. e.g. "200 OK"
	Status string
	// StatusCode is a status code as integer. e.g. 200
	StatusCode int
}
