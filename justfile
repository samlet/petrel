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

gen name +FLAGS='':
	go build -o bin/{{name}} codegen/{{name}}/*.go
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
eth:
    ganache-cli

login:
	curl -X POST -d 'username=jon' -d 'password=shhh!' localhost:1323/login

dummy:
	rm -rf routines/dummy
	just run generator -w -s flow-def Dummy downloadFile processFile

gen-alpha:
	cd fixtures/alphafac && go generate

tokenProcs:
    cd fixtures/eth && go generate
    just run eth -s tokenProcs

