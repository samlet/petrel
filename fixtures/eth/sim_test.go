package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strings"
	"testing"
)

// ref: https://github.com/ethereum/go-ethereum/blob/master/accounts/abi/bind/backends/simulated_test.go
func TestSimulatedBackend_EstimateGas(t *testing.T) {
	/*
		pragma solidity ^0.6.4;
		contract GasEstimation {
		    function PureRevert() public { revert(); }
		    function Revert() public { revert("revert reason");}
		    function OOG() public { for (uint i = 0; ; i++) {}}
		    function Assert() public { assert(false);}
		    function Valid() public {}
		}*/
	const contractAbi = "[{\"inputs\":[],\"name\":\"Assert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OOG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PureRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Revert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Valid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	const contractBin = "0x60806040523480156100115760006000fd5b50610017565b61016e806100266000396000f3fe60806040523480156100115760006000fd5b506004361061005c5760003560e01c806350f6fe3414610062578063aa8b1d301461006c578063b9b046f914610076578063d8b9839114610080578063e09fface1461008a5761005c565b60006000fd5b61006a610094565b005b6100746100ad565b005b61007e6100b5565b005b6100886100c2565b005b610092610135565b005b6000600090505b5b808060010191505061009b565b505b565b60006000fd5b565b600015156100bf57fe5b5b565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f72657665727420726561736f6e0000000000000000000000000000000000000081526020015060200191505060405180910390fd5b565b5b56fea2646970667358221220345bbcbb1a5ecf22b53a78eaebf95f8ee0eceff6d10d4b9643495084d2ec934a64736f6c63430006040033"

	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)
	opts, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))

	sim := backends.NewSimulatedBackend(core.GenesisAlloc{addr: {Balance: big.NewInt(params.Ether)}}, 10000000)
	defer sim.Close()

	parsed, _ := abi.JSON(strings.NewReader(contractAbi))
	contractAddr, _, _, _ := bind.DeployContract(opts, parsed, common.FromHex(contractBin), sim)
	sim.Commit()

	var cases = []struct {
		name        string
		message     ethereum.CallMsg
		expect      uint64
		expectError error
		expectData  interface{}
	}{
		{"plain transfer(valid)", ethereum.CallMsg{
			From:     addr,
			To:       &addr,
			Gas:      0,
			GasPrice: big.NewInt(0),
			Value:    big.NewInt(1),
			Data:     nil,
		}, params.TxGas, nil, nil},

		{"plain transfer(invalid)", ethereum.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      0,
			GasPrice: big.NewInt(0),
			Value:    big.NewInt(1),
			Data:     nil,
		}, 0, errors.New("execution reverted"), nil},

		{"Revert", ethereum.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      0,
			GasPrice: big.NewInt(0),
			Value:    nil,
			Data:     common.Hex2Bytes("d8b98391"),
		}, 0, errors.New("execution reverted: revert reason"), "0x08c379a00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000d72657665727420726561736f6e00000000000000000000000000000000000000"},

		{"PureRevert", ethereum.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      0,
			GasPrice: big.NewInt(0),
			Value:    nil,
			Data:     common.Hex2Bytes("aa8b1d30"),
		}, 0, errors.New("execution reverted"), nil},

		{"OOG", ethereum.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      100000,
			GasPrice: big.NewInt(0),
			Value:    nil,
			Data:     common.Hex2Bytes("50f6fe34"),
		}, 0, errors.New("gas required exceeds allowance (100000)"), nil},

		{"Assert", ethereum.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      100000,
			GasPrice: big.NewInt(0),
			Value:    nil,
			Data:     common.Hex2Bytes("b9b046f9"),
		}, 0, errors.New("invalid opcode: opcode 0xfe not defined"), nil},

		{"Valid", ethereum.CallMsg{
			From:     addr,
			To:       &contractAddr,
			Gas:      100000,
			GasPrice: big.NewInt(0),
			Value:    nil,
			Data:     common.Hex2Bytes("e09fface"),
		}, 21275, nil, nil},
	}
	for _, c := range cases {
		got, err := sim.EstimateGas(context.Background(), c.message)
		if c.expectError != nil {
			if err == nil {
				t.Fatalf("Expect error, got nil")
			}
			if c.expectError.Error() != err.Error() {
				t.Fatalf("Expect error, want %v, got %v", c.expectError, err)
			}
			if c.expectData != nil {
				/*
					if err, ok := err.(*revertError); !ok {
						t.Fatalf("Expect revert error, got %T", err)
					} else if !reflect.DeepEqual(err.ErrorData(), c.expectData) {
						t.Fatalf("Error data mismatch, want %v, got %v", c.expectData, err.ErrorData())
					}
				*/
				fmt.Printf("expectData: %s\n", c.expectData)
			}
			continue
		}
		if got != c.expect {
			t.Fatalf("Gas estimation mismatch, want %d, got %d", c.expect, got)
		}
	}
}

//// revertError is an API error that encompasses an EVM revert with JSON error
//// code and a binary data blob.
//type revertError struct {
//	error
//	reason string // revert reason hex encoded
//}
//// ErrorCode returns the JSON error code for a revert.
//// See: https://github.com/ethereum/wiki/wiki/JSON-RPC-Error-Codes-Improvement-Proposal
//func (e *revertError) ErrorCode() int {
//	return 3
//}
//
//// ErrorData returns the hex encoded revert reason.
//func (e *revertError) ErrorData() interface{} {
//	return e.reason
//}
