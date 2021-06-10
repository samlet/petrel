.PHONY: build
build: lint
	solc --overwrite --abi --bin --optimize -o build contract.sol
	mkdir -p contract
	abigen -abi build/RedPackage.abi --bin build/RedPackage.bin --pkg contract --out contract/redpkg.go
lint:
	prettier --write *.sol
clean:
	rm -rf build
	rm contract/redpkg.go
deps:
	brew tap ethereum/ethereum
	brew install solidity ethereum node
	npm install -g prettier prettier-plugin-solidity --registry=https://registry.npm.taobao.org
