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
fileprocessing-test:
    go test routines/fileprocessing/*.go -v

worker name:
    go build -o bin/{{name}} routines/{{name}}/*.go
    ./bin/{{name}} -m worker

dsl name filename:
    go build -o bin/{{name}} routines/{{name}}/*.go
    ./bin/{{name}} -dslConfig routines/{{name}}/{{filename}}.yaml

clear:
	rm -rf output/*

substrate:
    substrate --dev --rpc-external --ws-external --rpc-methods Unsafe --offchain-worker Always