list:
	go list -m all
run name:
	go build -o bin/{{name}} fixtures/{{name}}/*.go
	./bin/{{name}}
simple:
	go run fixtures/simpleclient/*.go



