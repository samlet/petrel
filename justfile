export CADENCE_CLI_DOMAIN:="samples-domain"

list:
	go list -m all

# $ just run simpleclient --service find
# $ just run simpleclient -s find
run name +FLAGS='':
	go build -o bin/ft{{name}} fixtures/{{name}}/*.go
	./bin/ft{{name}} {{FLAGS}}

wire name +FLAGS='':
    cd fixtures/{{name}} && go build -o ../../bin/ft{{name}}
    ./bin/ft{{name}} {{FLAGS}} || true

rt name +FLAGS='':
	go build -o bin/rt{{name}} routines/{{name}}/*.go
	./bin/rt{{name}} {{FLAGS}}

gen name +FLAGS='':
	go build -o bin/gn{{name}} codegen/{{name}}/*.go
	./bin/gn{{name}} {{FLAGS}}

alfin name +FLAGS='':
	go build -o bin/al{{name}} alfin/{{name}}/*.go
	./bin/al{{name}} {{FLAGS}}

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
gen-ent name:
    just alfin gentempl -s gen -t alfin/entity_type.tmpl -i alfin/{{name}}.json -o alfin/{{name}}.go

# Web3 Provider from the Environment dropdown, and add “http://127.0.0.1:8545”
geth:
	rm -rf test-chain-dir/
	mkdir test-chain-dir
	geth --datadir test-chain-dir --http --dev --http.corsdomain "https://remix.ethereum.org,http://remix.ethereum.org"

geth-cli:
	geth attach ./test-chain-dir/geth.ipc

tokenProcs:
    cd fixtures/eth && go generate
    just run eth -s tokenProcs

# just wf Seed '"cadence"'
# omit parameter: --domain samples-domain
wf name json_pars:
    cadence workflow run --tl {{name}}Group --wt main.{{name}}Workflow --et 60 -i '{{json_pars}}'

# just meta-cli meta Person
meta-cli +FLAGS='':
	cargo +nightly run --manifest-path ~/sagas/projs/bluecc/entity-seed/meta-cli/Cargo.toml -- {{FLAGS}}


