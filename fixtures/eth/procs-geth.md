## refs
* https://github.com/ethereum/go-ethereum/blob/master/accounts/abi/bind/backends/simulated.go

## start
⊕ [Dev mode | Go Ethereum](https://geth.ethereum.org/docs/getting-started/dev-mode)

```bash
mkdir test-chain-dir
geth --datadir test-chain-dir --http --dev --http.corsdomain "https://remix.ethereum.org,http://remix.ethereum.org"
geth attach <IPC_LOCATION>

#Once geth is running in dev mode, you can interact with it in the same way as when geth is running in other ways.
#For example, create a test account:

> personal.newAccount()

#And transfer ether from the coinbase to the new account:

> eth.sendTransaction({from:eth.coinbase, to:eth.accounts[1], value: web3.toWei(0.05, "ether")})

#And check the balance of the account:

> eth.getBalance(eth.accounts[1])
```
If you want to test your dapps with a realistic block time use the --dev.period option when you start dev mode with the --dev.period 14 argument.
