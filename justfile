list:
	go list -m all

# $ just run simpleclient --service find
# $ just run simpleclient -s find
run name +FLAGS='':
	go build -o bin/{{name}} fixtures/{{name}}/*.go
	./bin/{{name}} {{FLAGS}}

simple:
	go run fixtures/simpleclient/*.go



