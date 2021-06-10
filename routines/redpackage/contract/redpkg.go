// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"},{\"name\":\"equalDivision\",\"type\":\"bool\"},{\"name\":\"onlyFriend\",\"type\":\"bool\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"expireDays\",\"type\":\"uint256\"}],\"name\":\"Giving\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MyToken\",\"outputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_friend\",\"type\":\"address\"}],\"name\":\"AddFriend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"}],\"name\":\"getRecord\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"equalDivision\",\"type\":\"bool\"},{\"name\":\"onlyFriend\",\"type\":\"bool\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"remainAmount\",\"type\":\"uint256\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"remainSize\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"expired\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"}],\"name\":\"Revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"list\",\"type\":\"address[]\"}],\"name\":\"AddFriendList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MyFriends\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_friend\",\"type\":\"address\"}],\"name\":\"DelFriend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"NewTokenReceiver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"friendship\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"}],\"name\":\"Grabbing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"}],\"name\":\"CanGrab\",\"outputs\":[{\"name\":\"has\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"word\",\"type\":\"bytes32\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"equalDivision\",\"type\":\"bool\"},{\"name\":\"onlyFriend\",\"type\":\"bool\"},{\"name\":\"size\",\"type\":\"uint256\"},{\"name\":\"expireDays\",\"type\":\"uint256\"}],\"name\":\"Giving\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Send\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Receive\",\"type\":\"event\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x6080604052600060055534801561001557600080fd5b506123cd806100256000396000f3fe6080604052600436106100c15763ffffffff60e060020a6000350416630372388581146100c657806306758f021461010b578063089327de146101445780631396225814610175578063213681cd146101a857806322b3f54914610234578063239fcf0f1461025e5780632bba4dde146102a357806337b37160146103535780635663132e146103b85780635daccf51146103eb57806366031fdf1461040057806386060c021461044f5780638749a6fa14610479578063d3a04056146104a3575b600080fd5b3480156100d257600080fd5b50610109600480360360608110156100e957600080fd5b50600160a060020a038135811691602081013590911690604001356104fe565b005b610109600480360360a081101561012157600080fd5b508035906020810135151590604081013515159060608101359060800135610598565b34801561015057600080fd5b506101596108c1565b60408051600160a060020a039092168252519081900360200190f35b34801561018157600080fd5b506101096004803603602081101561019857600080fd5b5035600160a060020a03166108de565b3480156101b457600080fd5b506101d2600480360360208110156101cb57600080fd5b5035610a4b565b60408051600160a060020a039b8c16815299151560208b015297151589890152959098166060880152608087019390935260a086019190915260c085015260e08401526101008301939093526101208201929092529051908190036101400190f35b34801561024057600080fd5b506101096004803603602081101561025757600080fd5b5035610b59565b34801561026a57600080fd5b506102916004803603602081101561028157600080fd5b5035600160a060020a0316610d1a565b60408051918252519081900360200190f35b3480156102af57600080fd5b50610109600480360360208110156102c657600080fd5b8101906020810181356401000000008111156102e157600080fd5b8201836020820111156102f357600080fd5b8035906020019184602083028401116401000000008311171561031557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610dbb945050505050565b34801561035f57600080fd5b50610368610e9e565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103a457818101518382015260200161038c565b505050509050019250505060405180910390f35b3480156103c457600080fd5b50610109600480360360208110156103db57600080fd5b5035600160a060020a0316610f07565b3480156103f757600080fd5b506101096110ac565b34801561040c57600080fd5b5061043b6004803603604081101561042357600080fd5b50600160a060020a03813581169160200135166111a1565b604080519115158252519081900360200190f35b34801561045b57600080fd5b506101096004803603602081101561047257600080fd5b50356111c1565b34801561048557600080fd5b5061043b6004803603602081101561049c57600080fd5b5035611642565b3480156104af57600080fd5b50610109600480360360e08110156104c657600080fd5b50803590600160a060020a03602082013516906040810135906060810135151590608081013515159060a08101359060c001356117d6565b336000908152600260205260408082205481517f93eb3c62000000000000000000000000000000000000000000000000000000008152600160a060020a038781166004830152868116602483015260448201869052925192909116926393eb3c629260648084019382900301818387803b15801561057b57600080fd5b505af115801561058f573d6000803e3d6000fd5b50505050505050565b6000821180156105a85750600034115b80156105b357508134115b80156105bf5750600081115b1515610615576040805160e560020a62461bcd02815260206004820152601560248201527f696e76616c696420646174612070726f76696465640000000000000000000000604482015290519081900360640190fd5b61061e85611bf7565b15610673576040805160e560020a62461bcd02815260206004820152601260248201527f526564207061636b616765206578697374730000000000000000000000000000604482015290519081900360640190fd5b831561070057813481151561068457fe5b0615610700576040805160e560020a62461bcd02815260206004820152602e60248201527f496e76616c69642076616c756520616e642073697a6520666f7220657175616c60448201527f206469766973696f6e206d6f6465000000000000000000000000000000000000606482015290519081900360840190fd5b6101606040519081016040528033600160a060020a03168152602001851515815260200184151581526020016000600160a060020a03168152602001348152602001348152602001838152602001838152602001428152602001826201518002420181526020016005548152506003600087815260200190815260200160002060008201518160000160006101000a815481600160a060020a030219169083600160a060020a0316021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160000160156101000a81548160ff02191690831515021790555060608201518160010160006101000a815481600160a060020a030219169083600160a060020a031602179055506080820151816002015560a0820151816003015560c0820151816004015560e08201518160050155610100820151816006015561012082015181600701556101408201518160080155905050600560008154809291906001019190505550346000600160a060020a031633600160a060020a03167f93eb3c629eb575edaf0252e4f9fc0c5ccada50496f8c1d32f0f93a65a8257eb560405160405180910390a45050505050565b33600090815260026020526040902054600160a060020a03165b90565b600160a060020a038116331415610965576040805160e560020a62461bcd02815260206004820152602160248201527f43616e27742061646420796f757273656c6620746f20667269656e64206c697360448201527f7400000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b336000908152600160209081526040808320600160a060020a038516845290915290205460ff16156109e1576040805160e560020a62461bcd02815260206004820152601460248201527f467269656e6420616c7265616479206164646564000000000000000000000000604482015290519081900360640190fd5b3360008181526020818152604080832080546001808201835591855283852001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0397909716968717905593835283825280832094835293905291909120805460ff19169091179055565b600080600080600080600080600080610a62611e31565b5050506000988952505060036020818152604098899020895161016081018b528154600160a060020a0380821680845260ff74010000000000000000000000000000000000000000840481161515968501879052750100000000000000000000000000000000000000000090930490921615159c83018d9052600184015416606083018190526002840154608084018190529584015460a08401819052600485015460c08501819052600586015460e08601819052600687015461010087018190526007880154610120880181905260089098015461014090970196909652939f969e9d50919b5095995094975093955092935090565b60008181526003602052604090208054600160a060020a03163314610bee576040805160e560020a62461bcd02815260206004820152602e60248201527f526564207061636b616765206e6f7420657869737473206f7220796f7527726560448201527f206e6f7420746865206f776e6572000000000000000000000000000000000000606482015290519081900360840190fd5b60078101544211610c49576040805160e560020a62461bcd02815260206004820152601760248201527f4f6e6c79207265766f6b652065787069726564206f6e65000000000000000000604482015290519081900360640190fd5b6001810154600160a060020a03161515610c93576002810154604051339180156108fc02916000818181858888f19350505050158015610c8d573d6000803e3d6000fd5b50610c9e565b610c9e336000611cd5565b5060009081526003602081905260408220805475ffffffffffffffffffffffffffffffffffffffffffff1916815560018101805473ffffffffffffffffffffffffffffffffffffffff19169055600281018390559081018290556004810182905560058101829055600681018290556007810182905560080155565b3360009081526002602090815260408083205481517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152925192909116926370a0823192602480840193829003018186803b158015610d8957600080fd5b505afa158015610d9d573d6000803e3d6000fd5b505050506040513d6020811015610db357600080fd5b505192915050565b60005b8151811015610e9a5760008282815181101515610dd757fe5b6020908102909101810151336000908152600183526040808220600160a060020a0384168352909352919091205490915060ff1680610e1e5750600160a060020a03811633145b15610e295750610e92565b3360008181526020818152604080832080546001808201835591855283852001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0397909716968717905593835283825280832094835293905291909120805460ff191690911790555b600101610dbe565b5050565b3360009081526020818152604091829020805483518184028101840190945280845260609392830182828015610efd57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610edf575b5050505050905090565b336000908152600160209081526040808320600160a060020a038516845290915290205460ff161515610f39576110a9565b336000908152600160209081526040808320600160a060020a03851684529091528120805460ff191690555b33600090815260208190526040902054811015610e9a57336000908152602081905260409020805482908110610f9757fe5b600091825260209091200154600160a060020a03838116911614610fba576110a1565b33600090815260208190526040902054811415610ff857336000908152602081905260409020805490610ff1906000198301611ea2565b50506110a9565b3360009081526020819052604090208054600019810190811061101757fe5b60009182526020808320909101543383529082905260409091208054600160a060020a03909216918390811061104957fe5b6000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693909317909255338152908190526040902080549061109f906000198301611ea2565b505b600101610f65565b50565b33600090815260026020526040902054600160a060020a031615611140576040805160e560020a62461bcd02815260206004820152602160248201527f486173206265656e20637265617465642072656365697665722061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b611148611ec6565b604051809103906000f080158015611164573d6000803e3d6000fd5b50336000908152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160209081526000928352604080842090915290825290205460ff1681565b60008181526003602052604090208054600160a060020a03161515611230576040805160e560020a62461bcd02815260206004820152601660248201527f526564207061636b616765206e6f742065786973747300000000000000000000604482015290519081900360640190fd5b600781015442111561128c576040805160e560020a62461bcd02815260206004820152601360248201527f526564207061636b616765206578706972656400000000000000000000000000604482015290519081900360640190fd5b80547501000000000000000000000000000000000000000000900460ff161561132d578054600160a060020a0316600090815260016020908152604080832033845290915290205460ff16151561132d576040805160e560020a62461bcd02815260206004820152601460248201527f4f6e6c7920667269656e642063616e2067726162000000000000000000000000604482015290519081900360640190fd5b6008810154600090815260046020908152604080832033845290915290205460ff16156113a4576040805160e560020a62461bcd02815260206004820152601360248201527f63616e2774206772616262656420747769636500000000000000000000000000604482015290519081900360640190fd5b805460009074010000000000000000000000000000000000000000900460ff16156113e457816004015482600201548115156113dc57fe5b049050611498565b8160050154600114156113fc57506003810154611498565b6003820154600583015460408051336020808301919091528183018590526060820193909352426080808301919091528251808303909101815260a0909101909152805191810191909120909160009181151561145557fe5b06905060008460050154856003015481151561146d57fe5b0490508115156114805760019350611494565b8082111561149057809350611494565b8193505b5050505b6001820154600160a060020a031615156114df57604051339082156108fc029083906000818181858888f193505050501580156114d9573d6000803e3d6000fd5b5061150f565b8154600160a060020a0316331461150f576001820154825461150f91600160a060020a0390811691163384611d93565b6003820180548290039055600582018054600019019081905515156115d75760008381526003602081905260408220805475ffffffffffffffffffffffffffffffffffffffffffff191681556001808201805473ffffffffffffffffffffffffffffffffffffffff19169055600282018490559181018390556004810183905560058101839055600681018390556007810183905560080191909155820154600160a060020a0316156115d25781546115d290600160a060020a03166000611cd5565b611601565b600882015460009081526004602090815260408083203384529091529020805460ff191660011790555b60018201546040518291600160a060020a03169033907ffd19781f43410d9594fd4c02dd1d98dbe99099cbd222d5851a6e183c468d33ca90600090a4505050565b60008181526003602052604081208054600160a060020a031615156116b1576040805160e560020a62461bcd02815260206004820152601660248201527f526564207061636b616765206e6f742065786973747300000000000000000000604482015290519081900360640190fd5b600781015442111561170d576040805160e560020a62461bcd02815260206004820152601360248201527f526564207061636b616765206578706972656400000000000000000000000000604482015290519081900360640190fd5b80547501000000000000000000000000000000000000000000900460ff16156117ae578054600160a060020a0316600090815260016020908152604080832033845290915290205460ff1615156117ae576040805160e560020a62461bcd02815260206004820152601460248201527f4f6e6c7920667269656e642063616e2067726162000000000000000000000000604482015290519081900360640190fd5b60080154600090815260046020908152604080832033845290915290205460ff161592915050565b3360009081526002602090815260408083205481517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038b81166004830152925192909116926370a0823192602480840193829003018186803b15801561184557600080fd5b505afa158015611859573d6000803e3d6000fd5b505050506040513d602081101561186f57600080fd5b50519050858110156118cb576040805160e560020a62461bcd02815260206004820152601460248201527f6e6f742073756666696369656e742066756e6473000000000000000000000000604482015290519081900360640190fd5b6000831180156118db5750600086115b80156118e657508286115b80156118f25750600082115b1515611948576040805160e560020a62461bcd02815260206004820152601560248201527f696e76616c696420646174612070726f76696465640000000000000000000000604482015290519081900360640190fd5b61195188611bf7565b156119a6576040805160e560020a62461bcd02815260206004820152601260248201527f526564207061636b616765206578697374730000000000000000000000000000604482015290519081900360640190fd5b8415611a335782868115156119b757fe5b0615611a33576040805160e560020a62461bcd02815260206004820152602e60248201527f496e76616c69642076616c756520616e642073697a6520666f7220657175616c60448201527f206469766973696f6e206d6f6465000000000000000000000000000000000000606482015290519081900360840190fd5b6101606040519081016040528033600160a060020a031681526020018615158152602001851515815260200188600160a060020a0316815260200187815260200187815260200184815260200184815260200142815260200183620151800242018152602001600554815250600360008a815260200190815260200160002060008201518160000160006101000a815481600160a060020a030219169083600160a060020a0316021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160000160156101000a81548160ff02191690831515021790555060608201518160010160006101000a815481600160a060020a030219169083600160a060020a031602179055506080820151816002015560a0820151816003015560c0820151816004015560e08201518160050155610100820151816006015561012082015181600701556101408201518160080155905050600560008154809291906001019190505550611bb5336001611cd5565b6040518690600160a060020a0389169033907f93eb3c629eb575edaf0252e4f9fc0c5ccada50496f8c1d32f0f93a65a8257eb590600090a45050505050505050565b6000611c01611e31565b50506000908152600360208181526040928390208351610160810185528154600160a060020a0380821680845260ff7401000000000000000000000000000000000000000084048116151596850196909652750100000000000000000000000000000000000000000090920490941615159582019590955260018201549092166060830152600281015460808301529182015460a0820152600482015460c0820152600582015460e08201526006820154610100820152600782015461012082015260089091015461014090910152151590565b600160a060020a038083166000908152600260205260409020541681611d505780600160a060020a03166370e3fffe6040518163ffffffff1660e060020a028152600401600060405180830381600087803b158015611d3357600080fd5b505af1158015611d47573d6000803e3d6000fd5b50505050611d8e565b80600160a060020a03166346620e396040518163ffffffff1660e060020a028152600401600060405180830381600087803b15801561057b57600080fd5b505050565b600160a060020a038084166000908152600260205260408082205481517f93eb3c62000000000000000000000000000000000000000000000000000000008152888516600482015286851660248201526044810186905291519316926393eb3c629260648084019391929182900301818387803b158015611e1357600080fd5b505af1158015611e27573d6000803e3d6000fd5b5050505050505050565b610160604051908101604052806000600160a060020a031681526020016000151581526020016000151581526020016000600160a060020a03168152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b815481835581811115611d8e57600083815260209020611d8e918101908301611ed6565b6040516104ad80611ef583390190565b6108db91905b80821115611ef05760008155600101611edc565b509056fe60806040526000805460ff1916905534801561001a57600080fd5b5060008054610100330261010060a860020a031990911617905561046a806100436000396000f3fe6080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166346620e39811461007c57806359a3ccb81461009357806370a08231146100dc57806370e3fffe1461012157806393eb3c6214610136578063cf30901214610179575b600080fd5b34801561008857600080fd5b506100916101a2565b005b34801561009f57600080fd5b50610091600480360360808110156100b657600080fd5b50600160a060020a038135811691602081013582169160408201351690606001356101cd565b3480156100e857600080fd5b5061010f600480360360208110156100ff57600080fd5b5035600160a060020a031661029f565b60408051918252519081900360200190f35b34801561012d57600080fd5b50610091610334565b34801561014257600080fd5b506100916004803603606081101561015957600080fd5b50600160a060020a0381358116916020810135909116906040013561035c565b34801561018557600080fd5b5061018e610435565b604080519115158252519081900360200190f35b6000546101009004600160a060020a031633146101be57600080fd5b6000805460ff19166001179055565b6000546101009004600160a060020a031633146101e957600080fd5b60005460ff16156101f957600080fd5b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301528481166024830152604482018490529151918616916323b872dd916064808201926020929091908290030181600087803b15801561026d57600080fd5b505af1158015610281573d6000803e3d6000fd5b505050506040513d602081101561029757600080fd5b505050505050565b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600091600160a060020a038416916370a0823191602480820192602092909190829003018186803b15801561030257600080fd5b505afa158015610316573d6000803e3d6000fd5b505050506040513d602081101561032c57600080fd5b505192915050565b6000546101009004600160a060020a0316331461035057600080fd5b6000805460ff19169055565b6000546101009004600160a060020a0316331461037857600080fd5b60005460ff161561038857600080fd5b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561040457600080fd5b505af1158015610418573d6000803e3d6000fd5b505050506040513d602081101561042e57600080fd5b5050505050565b60005460ff168156fea165627a7a72305820de6d5a75974d86d78bac349de32675876342d515777df31aeda33360e84807800029a165627a7a7230582026da83e99650e32e880b14f79f6ea89df4cd33e24f33d89c422d3b66060e20680029"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) view returns(uint256)
func (_Contract *ContractCaller) Balance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "Balance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) view returns(uint256)
func (_Contract *ContractSession) Balance(token common.Address) (*big.Int, error) {
	return _Contract.Contract.Balance(&_Contract.CallOpts, token)
}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) view returns(uint256)
func (_Contract *ContractCallerSession) Balance(token common.Address) (*big.Int, error) {
	return _Contract.Contract.Balance(&_Contract.CallOpts, token)
}

// CanGrab is a free data retrieval call binding the contract method 0x8749a6fa.
//
// Solidity: function CanGrab(bytes32 word) view returns(bool has)
func (_Contract *ContractCaller) CanGrab(opts *bind.CallOpts, word [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "CanGrab", word)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanGrab is a free data retrieval call binding the contract method 0x8749a6fa.
//
// Solidity: function CanGrab(bytes32 word) view returns(bool has)
func (_Contract *ContractSession) CanGrab(word [32]byte) (bool, error) {
	return _Contract.Contract.CanGrab(&_Contract.CallOpts, word)
}

// CanGrab is a free data retrieval call binding the contract method 0x8749a6fa.
//
// Solidity: function CanGrab(bytes32 word) view returns(bool has)
func (_Contract *ContractCallerSession) CanGrab(word [32]byte) (bool, error) {
	return _Contract.Contract.CanGrab(&_Contract.CallOpts, word)
}

// MyFriends is a free data retrieval call binding the contract method 0x37b37160.
//
// Solidity: function MyFriends() view returns(address[])
func (_Contract *ContractCaller) MyFriends(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MyFriends")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// MyFriends is a free data retrieval call binding the contract method 0x37b37160.
//
// Solidity: function MyFriends() view returns(address[])
func (_Contract *ContractSession) MyFriends() ([]common.Address, error) {
	return _Contract.Contract.MyFriends(&_Contract.CallOpts)
}

// MyFriends is a free data retrieval call binding the contract method 0x37b37160.
//
// Solidity: function MyFriends() view returns(address[])
func (_Contract *ContractCallerSession) MyFriends() ([]common.Address, error) {
	return _Contract.Contract.MyFriends(&_Contract.CallOpts)
}

// MyToken is a free data retrieval call binding the contract method 0x089327de.
//
// Solidity: function MyToken() view returns(address token)
func (_Contract *ContractCaller) MyToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MyToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MyToken is a free data retrieval call binding the contract method 0x089327de.
//
// Solidity: function MyToken() view returns(address token)
func (_Contract *ContractSession) MyToken() (common.Address, error) {
	return _Contract.Contract.MyToken(&_Contract.CallOpts)
}

// MyToken is a free data retrieval call binding the contract method 0x089327de.
//
// Solidity: function MyToken() view returns(address token)
func (_Contract *ContractCallerSession) MyToken() (common.Address, error) {
	return _Contract.Contract.MyToken(&_Contract.CallOpts)
}

// Friendship is a free data retrieval call binding the contract method 0x66031fdf.
//
// Solidity: function friendship(address , address ) view returns(bool)
func (_Contract *ContractCaller) Friendship(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "friendship", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Friendship is a free data retrieval call binding the contract method 0x66031fdf.
//
// Solidity: function friendship(address , address ) view returns(bool)
func (_Contract *ContractSession) Friendship(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Contract.Contract.Friendship(&_Contract.CallOpts, arg0, arg1)
}

// Friendship is a free data retrieval call binding the contract method 0x66031fdf.
//
// Solidity: function friendship(address , address ) view returns(bool)
func (_Contract *ContractCallerSession) Friendship(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Contract.Contract.Friendship(&_Contract.CallOpts, arg0, arg1)
}

// GetRecord is a free data retrieval call binding the contract method 0x213681cd.
//
// Solidity: function getRecord(bytes32 word) view returns(address owner, bool equalDivision, bool onlyFriend, address token, uint256 amount, uint256 remainAmount, uint256 size, uint256 remainSize, uint256 timestamp, uint256 expired)
func (_Contract *ContractCaller) GetRecord(opts *bind.CallOpts, word [32]byte) (struct {
	Owner         common.Address
	EqualDivision bool
	OnlyFriend    bool
	Token         common.Address
	Amount        *big.Int
	RemainAmount  *big.Int
	Size          *big.Int
	RemainSize    *big.Int
	Timestamp     *big.Int
	Expired       *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRecord", word)

	outstruct := new(struct {
		Owner         common.Address
		EqualDivision bool
		OnlyFriend    bool
		Token         common.Address
		Amount        *big.Int
		RemainAmount  *big.Int
		Size          *big.Int
		RemainSize    *big.Int
		Timestamp     *big.Int
		Expired       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EqualDivision = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.OnlyFriend = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Token = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RemainAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Size = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.RemainSize = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Expired = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRecord is a free data retrieval call binding the contract method 0x213681cd.
//
// Solidity: function getRecord(bytes32 word) view returns(address owner, bool equalDivision, bool onlyFriend, address token, uint256 amount, uint256 remainAmount, uint256 size, uint256 remainSize, uint256 timestamp, uint256 expired)
func (_Contract *ContractSession) GetRecord(word [32]byte) (struct {
	Owner         common.Address
	EqualDivision bool
	OnlyFriend    bool
	Token         common.Address
	Amount        *big.Int
	RemainAmount  *big.Int
	Size          *big.Int
	RemainSize    *big.Int
	Timestamp     *big.Int
	Expired       *big.Int
}, error) {
	return _Contract.Contract.GetRecord(&_Contract.CallOpts, word)
}

// GetRecord is a free data retrieval call binding the contract method 0x213681cd.
//
// Solidity: function getRecord(bytes32 word) view returns(address owner, bool equalDivision, bool onlyFriend, address token, uint256 amount, uint256 remainAmount, uint256 size, uint256 remainSize, uint256 timestamp, uint256 expired)
func (_Contract *ContractCallerSession) GetRecord(word [32]byte) (struct {
	Owner         common.Address
	EqualDivision bool
	OnlyFriend    bool
	Token         common.Address
	Amount        *big.Int
	RemainAmount  *big.Int
	Size          *big.Int
	RemainSize    *big.Int
	Timestamp     *big.Int
	Expired       *big.Int
}, error) {
	return _Contract.Contract.GetRecord(&_Contract.CallOpts, word)
}

// AddFriend is a paid mutator transaction binding the contract method 0x13962258.
//
// Solidity: function AddFriend(address _friend) returns()
func (_Contract *ContractTransactor) AddFriend(opts *bind.TransactOpts, _friend common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "AddFriend", _friend)
}

// AddFriend is a paid mutator transaction binding the contract method 0x13962258.
//
// Solidity: function AddFriend(address _friend) returns()
func (_Contract *ContractSession) AddFriend(_friend common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddFriend(&_Contract.TransactOpts, _friend)
}

// AddFriend is a paid mutator transaction binding the contract method 0x13962258.
//
// Solidity: function AddFriend(address _friend) returns()
func (_Contract *ContractTransactorSession) AddFriend(_friend common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddFriend(&_Contract.TransactOpts, _friend)
}

// AddFriendList is a paid mutator transaction binding the contract method 0x2bba4dde.
//
// Solidity: function AddFriendList(address[] list) returns()
func (_Contract *ContractTransactor) AddFriendList(opts *bind.TransactOpts, list []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "AddFriendList", list)
}

// AddFriendList is a paid mutator transaction binding the contract method 0x2bba4dde.
//
// Solidity: function AddFriendList(address[] list) returns()
func (_Contract *ContractSession) AddFriendList(list []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddFriendList(&_Contract.TransactOpts, list)
}

// AddFriendList is a paid mutator transaction binding the contract method 0x2bba4dde.
//
// Solidity: function AddFriendList(address[] list) returns()
func (_Contract *ContractTransactorSession) AddFriendList(list []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddFriendList(&_Contract.TransactOpts, list)
}

// DelFriend is a paid mutator transaction binding the contract method 0x5663132e.
//
// Solidity: function DelFriend(address _friend) returns()
func (_Contract *ContractTransactor) DelFriend(opts *bind.TransactOpts, _friend common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "DelFriend", _friend)
}

// DelFriend is a paid mutator transaction binding the contract method 0x5663132e.
//
// Solidity: function DelFriend(address _friend) returns()
func (_Contract *ContractSession) DelFriend(_friend common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DelFriend(&_Contract.TransactOpts, _friend)
}

// DelFriend is a paid mutator transaction binding the contract method 0x5663132e.
//
// Solidity: function DelFriend(address _friend) returns()
func (_Contract *ContractTransactorSession) DelFriend(_friend common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DelFriend(&_Contract.TransactOpts, _friend)
}

// Giving is a paid mutator transaction binding the contract method 0x06758f02.
//
// Solidity: function Giving(bytes32 word, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) payable returns()
func (_Contract *ContractTransactor) Giving(opts *bind.TransactOpts, word [32]byte, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Giving", word, equalDivision, onlyFriend, size, expireDays)
}

// Giving is a paid mutator transaction binding the contract method 0x06758f02.
//
// Solidity: function Giving(bytes32 word, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) payable returns()
func (_Contract *ContractSession) Giving(word [32]byte, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Giving(&_Contract.TransactOpts, word, equalDivision, onlyFriend, size, expireDays)
}

// Giving is a paid mutator transaction binding the contract method 0x06758f02.
//
// Solidity: function Giving(bytes32 word, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) payable returns()
func (_Contract *ContractTransactorSession) Giving(word [32]byte, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Giving(&_Contract.TransactOpts, word, equalDivision, onlyFriend, size, expireDays)
}

// Giving0 is a paid mutator transaction binding the contract method 0xd3a04056.
//
// Solidity: function Giving(bytes32 word, address token, uint256 value, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) returns()
func (_Contract *ContractTransactor) Giving0(opts *bind.TransactOpts, word [32]byte, token common.Address, value *big.Int, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Giving0", word, token, value, equalDivision, onlyFriend, size, expireDays)
}

// Giving0 is a paid mutator transaction binding the contract method 0xd3a04056.
//
// Solidity: function Giving(bytes32 word, address token, uint256 value, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) returns()
func (_Contract *ContractSession) Giving0(word [32]byte, token common.Address, value *big.Int, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Giving0(&_Contract.TransactOpts, word, token, value, equalDivision, onlyFriend, size, expireDays)
}

// Giving0 is a paid mutator transaction binding the contract method 0xd3a04056.
//
// Solidity: function Giving(bytes32 word, address token, uint256 value, bool equalDivision, bool onlyFriend, uint256 size, uint256 expireDays) returns()
func (_Contract *ContractTransactorSession) Giving0(word [32]byte, token common.Address, value *big.Int, equalDivision bool, onlyFriend bool, size *big.Int, expireDays *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Giving0(&_Contract.TransactOpts, word, token, value, equalDivision, onlyFriend, size, expireDays)
}

// Grabbing is a paid mutator transaction binding the contract method 0x86060c02.
//
// Solidity: function Grabbing(bytes32 word) returns()
func (_Contract *ContractTransactor) Grabbing(opts *bind.TransactOpts, word [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Grabbing", word)
}

// Grabbing is a paid mutator transaction binding the contract method 0x86060c02.
//
// Solidity: function Grabbing(bytes32 word) returns()
func (_Contract *ContractSession) Grabbing(word [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Grabbing(&_Contract.TransactOpts, word)
}

// Grabbing is a paid mutator transaction binding the contract method 0x86060c02.
//
// Solidity: function Grabbing(bytes32 word) returns()
func (_Contract *ContractTransactorSession) Grabbing(word [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Grabbing(&_Contract.TransactOpts, word)
}

// NewTokenReceiver is a paid mutator transaction binding the contract method 0x5daccf51.
//
// Solidity: function NewTokenReceiver() returns()
func (_Contract *ContractTransactor) NewTokenReceiver(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "NewTokenReceiver")
}

// NewTokenReceiver is a paid mutator transaction binding the contract method 0x5daccf51.
//
// Solidity: function NewTokenReceiver() returns()
func (_Contract *ContractSession) NewTokenReceiver() (*types.Transaction, error) {
	return _Contract.Contract.NewTokenReceiver(&_Contract.TransactOpts)
}

// NewTokenReceiver is a paid mutator transaction binding the contract method 0x5daccf51.
//
// Solidity: function NewTokenReceiver() returns()
func (_Contract *ContractTransactorSession) NewTokenReceiver() (*types.Transaction, error) {
	return _Contract.Contract.NewTokenReceiver(&_Contract.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x22b3f549.
//
// Solidity: function Revoke(bytes32 word) returns()
func (_Contract *ContractTransactor) Revoke(opts *bind.TransactOpts, word [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "Revoke", word)
}

// Revoke is a paid mutator transaction binding the contract method 0x22b3f549.
//
// Solidity: function Revoke(bytes32 word) returns()
func (_Contract *ContractSession) Revoke(word [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Revoke(&_Contract.TransactOpts, word)
}

// Revoke is a paid mutator transaction binding the contract method 0x22b3f549.
//
// Solidity: function Revoke(bytes32 word) returns()
func (_Contract *ContractTransactorSession) Revoke(word [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Revoke(&_Contract.TransactOpts, word)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x03723885.
//
// Solidity: function WithdrawToken(address token, address _to, uint256 value) returns()
func (_Contract *ContractTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, _to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "WithdrawToken", token, _to, value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x03723885.
//
// Solidity: function WithdrawToken(address token, address _to, uint256 value) returns()
func (_Contract *ContractSession) WithdrawToken(token common.Address, _to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawToken(&_Contract.TransactOpts, token, _to, value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x03723885.
//
// Solidity: function WithdrawToken(address token, address _to, uint256 value) returns()
func (_Contract *ContractTransactorSession) WithdrawToken(token common.Address, _to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawToken(&_Contract.TransactOpts, token, _to, value)
}

// ContractReceiveIterator is returned from FilterReceive and is used to iterate over the raw logs and unpacked data for Receive events raised by the Contract contract.
type ContractReceiveIterator struct {
	Event *ContractReceive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractReceive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractReceive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractReceive represents a Receive event raised by the Contract contract.
type ContractReceive struct {
	Receiver common.Address
	Token    common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceive is a free log retrieval operation binding the contract event 0xfd19781f43410d9594fd4c02dd1d98dbe99099cbd222d5851a6e183c468d33ca.
//
// Solidity: event Receive(address indexed receiver, address indexed token, uint256 indexed value)
func (_Contract *ContractFilterer) FilterReceive(opts *bind.FilterOpts, receiver []common.Address, token []common.Address, value []*big.Int) (*ContractReceiveIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Receive", receiverRule, tokenRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ContractReceiveIterator{contract: _Contract.contract, event: "Receive", logs: logs, sub: sub}, nil
}

// WatchReceive is a free log subscription operation binding the contract event 0xfd19781f43410d9594fd4c02dd1d98dbe99099cbd222d5851a6e183c468d33ca.
//
// Solidity: event Receive(address indexed receiver, address indexed token, uint256 indexed value)
func (_Contract *ContractFilterer) WatchReceive(opts *bind.WatchOpts, sink chan<- *ContractReceive, receiver []common.Address, token []common.Address, value []*big.Int) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Receive", receiverRule, tokenRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractReceive)
				if err := _Contract.contract.UnpackLog(event, "Receive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceive is a log parse operation binding the contract event 0xfd19781f43410d9594fd4c02dd1d98dbe99099cbd222d5851a6e183c468d33ca.
//
// Solidity: event Receive(address indexed receiver, address indexed token, uint256 indexed value)
func (_Contract *ContractFilterer) ParseReceive(log types.Log) (*ContractReceive, error) {
	event := new(ContractReceive)
	if err := _Contract.contract.UnpackLog(event, "Receive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the Contract contract.
type ContractSendIterator struct {
	Event *ContractSend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSend represents a Send event raised by the Contract contract.
type ContractSend struct {
	Sender common.Address
	Token  common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0x93eb3c629eb575edaf0252e4f9fc0c5ccada50496f8c1d32f0f93a65a8257eb5.
//
// Solidity: event Send(address indexed sender, address indexed token, uint256 indexed value)
func (_Contract *ContractFilterer) FilterSend(opts *bind.FilterOpts, sender []common.Address, token []common.Address, value []*big.Int) (*ContractSendIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Send", senderRule, tokenRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ContractSendIterator{contract: _Contract.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0x93eb3c629eb575edaf0252e4f9fc0c5ccada50496f8c1d32f0f93a65a8257eb5.
//
// Solidity: event Send(address indexed sender, address indexed token, uint256 indexed value)
func (_Contract *ContractFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *ContractSend, sender []common.Address, token []common.Address, value []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Send", senderRule, tokenRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSend)
				if err := _Contract.contract.UnpackLog(event, "Send", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSend is a log parse operation binding the contract event 0x93eb3c629eb575edaf0252e4f9fc0c5ccada50496f8c1d32f0f93a65a8257eb5.
//
// Solidity: event Send(address indexed sender, address indexed token, uint256 indexed value)
func (_Contract *ContractFilterer) ParseSend(log types.Log) (*ContractSend, error) {
	event := new(ContractSend)
	if err := _Contract.contract.UnpackLog(event, "Send", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
