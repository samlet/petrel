package retail

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"testing"
)

func TestDeployFruitStore2(t *testing.T) {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)
	opts, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))

	sim := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(params.Ether)}}, 10000000)
	defer sim.Close()

	// ...
	Do(opts, sim)
}

func Do(opts *bind.TransactOpts, sim *backends.SimulatedBackend) {
	_, _, c, err := DeployFruitStore2(opts, sim)
	if err != nil {
		log.Fatalf("Failed to deploy new contract: %v", err)
	}
	// Print the current (non existent) and pending name of the contract
	r, _ := c.SetFruitStock(opts, []byte("apple"), big.NewInt(15))
	fmt.Printf("Pre-mining: %s\n", r.Hash())
	// Commit all pending transactions in the simulator
	//sim.Commit()

	num, err:=c.GetStock(&bind.CallOpts{Pending: true}, []byte("apple"))
	if err != nil {
		log.Fatalf("get stock fail: %v", err)
	}
	println("apple stock:", num.String())
}
