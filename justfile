list:
	go list -m all

# $ just run simpleclient --service performFindList
# $ just run simpleclient -s performFindList
run name +FLAGS='':
	go build -o bin/{{name}} fixtures/{{name}}/*.go
	./bin/{{name}} {{FLAGS}}
simple:
	go run fixtures/simpleclient/*.go



