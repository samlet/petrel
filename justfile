list:
	go list -m all

# $ just run simpleclient --service find
# $ just run simpleclient -s find
run name +FLAGS='':
	go build -o bin/{{name}} fixtures/{{name}}/*.go
	./bin/{{name}} {{FLAGS}}

rt name +FLAGS='':
	go build -o bin/{{name}} routines/{{name}}/*.go
	./bin/{{name}} {{FLAGS}}

simple:
	go run fixtures/simpleclient/*.go

worker name:
    go build -o bin/{{name}} routines/{{name}}/*.go
    ./bin/{{name}} -m worker

dsl name filename:
    go build -o bin/{{name}} routines/{{name}}/*.go
    ./bin/{{name}} -dslConfig routines/{{name}}/{{filename}}.yaml
