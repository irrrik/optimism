// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressSortedLinkedListWithMedianMetaData contains all meta data concerning the AddressSortedLinkedListWithMedian contract.
var AddressSortedLinkedListWithMedianMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"toAddress\",\"inputs\":[{\"name\":\"b\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toBytes\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"}]",
	Bin: "0x611a3961003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100e95760003560e01c80636eafa6c31161009657806395073a791161007057806395073a7914610213578063c1e728e914610236578063d4a0927214610256578063d938ec7b1461027657600080fd5b80636eafa6c3146101cb5780637c6bb862146101de578063832a2147146101f157600080fd5b8063593b79fe116100c7578063593b79fe1461015257806359d556a8146101965780636cfa3873146101a957600080fd5b80630944c594146100ee5780633118159e1461012b578063341f66231461013e575b600080fd5b6101016100fc3660046116e0565b610289565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101016101393660046116e0565b61029c565b61010161014c3660046116e0565b60601c90565b610188610160366004611722565b60601b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001690565b604051908152602001610122565b6101886101a43660046116e0565b6102ac565b6101bc6101b73660046116e0565b6102c5565b604051610122939291906117df565b6101886101d93660046116e0565b610512565b6101886101ec366004611883565b61051f565b8180156101fd57600080fd5b5061021161020c3660046118af565b61055f565b005b610226610221366004611883565b6105a5565b6040519015158152602001610122565b81801561024257600080fd5b50610211610251366004611883565b6105d7565b81801561026257600080fd5b506102116102713660046118af565b61060b565b6101016102843660046116e0565b61064a565b600061029661014c835490565b92915050565b600061029661014c836005015490565b600061029682600501548361065a90919063ffffffff16565b606080606060006102d585610670565b90506000815167ffffffffffffffff8111156102f3576102f3611906565b60405190808252806020026020018201604052801561031c578160200160208202803683370190505b5090506000825167ffffffffffffffff81111561033b5761033b611906565b604051908082528060200260200182016040528015610364578160200160208202803683370190505b5090506000825167ffffffffffffffff81111561038357610383611906565b6040519080825280602002602001820160405280156103ac578160200160208202803683370190505b50905060005b8451811015610503576103de8582815181106103d0576103d0611935565b602002602001015160601c90565b8482815181106103f0576103f0611935565b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061045685828151811061043f5761043f611935565b60200260200101518a61065a90919063ffffffff16565b83828151811061046857610468611935565b60200260200101818152505088600601600086838151811061048c5761048c611935565b6020026020010151815260200190815260200160002060009054906101000a900460ff168282815181106104c2576104c2611935565b602002602001019060038111156104db576104db61173d565b908160038111156104ee576104ee61173d565b905250806104fb81611993565b9150506103b2565b50919790965090945092505050565b6000610296826002015490565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606082901b1660009081526004830160205260408120545b9392505050565b61059e857fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606087811b821691879187811b8216919087901b1661067b565b5050505050565b6000610558837fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b16610692565b610607827fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b1661069e565b5050565b61059e857fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606087811b821691879187811b8216919087901b166107ad565b600061029661014c836001015490565b6000908152600491909101602052604090205490565b6060610296826109c5565b610685858561069e565b61059e85858585856107ad565b600061055883836109d0565b600282015460009081036106b85760006005840155610794565b6002838101546106c891906119cb565b600003610734576002600083815260068501602052604090205460ff1660038111156106f6576106f661173d565b148061072657506003600083815260068501602052604090205460ff1660038111156107245761072461173d565b145b1561072f575060015b610794565b6001600083815260068501602052604090205460ff16600381111561075b5761075b61173d565b148061078b57506003600083815260068501602052604090205460ff1660038111156107895761078961173d565b145b15610794575060025b61079e83826109ec565b6107a88383610b06565b505050565b6107ba8585858585610b24565b6000848152600386016020526040812060028701549091906001036108265760058701869055600086815260068801602052604090208054600391907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001835b02179055506109b2565b60028781015461083691906119cb565b6001036108f75781541580610871575060018254600090815260068901602052604090205460ff16600381111561086f5761086f61173d565b145b156108b7575060008581526006870160205260409020805460019182917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016828061081c565b600086815260068801602052604090208054600291907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660018361081c565b60018201541580610931575060026001830154600090815260068901602052604090205460ff16600381111561092f5761092f61173d565b145b15610978575060008581526006870160205260409020805460029182917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660018361081c565b6000868152600688016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555b6109bc87826109ec565b50505050505050565b606061029682610d62565b600081815260038301602052604081206002015460ff16610558565b6005820154600090815260038301602052604090206001826002811115610a1557610a1561173d565b03610a62576005830180546000908152600685016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600217905581549055610ac3565b6002826002811115610a7657610a7661173d565b03610ac3576005830180546000908152600685016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915582015490555b50506005810154600090815260069091016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166003179055565b610b108282610d72565b600090815260049091016020526040812055565b8315801590610b335750818414155b8015610b3f5750808414155b8015610b525750610b5085856109d0565b155b610bbd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f696e76616c6964206b657900000000000000000000000000000000000000000060448201526064015b60405180910390fd5b81151580610bca57508015155b80610bd757506002850154155b610c3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f6772656174657220616e64206c6573736572206b6579207a65726f00000000006044820152606401610bb4565b610c4785836109d0565b80610c50575081155b610cb6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964206c6573736572206b657900000000000000000000000000006044820152606401610bb4565b610cc085826109d0565b80610cc9575080155b610d2f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f696e76616c69642067726561746572206b6579000000000000000000000000006044820152606401610bb4565b610d3b85848484610eca565b9092509050610d4c85858484611025565b5050600091825260049092016020526040902055565b6060610296828360020154611573565b600081815260038301602052604090208115801590610da45750600082815260038401602052604090206002015460ff165b610e0a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6b6579206e6f7420696e206c69737400000000000000000000000000000000006044820152606401610bb4565b805415610e3157805460009081526003840160205260409020600180830154910155610e3c565b600180820154908401555b600181015415610e645760018101546000908152600384016020526040902081549055610e69565b805483555b60008281526003840160205260408120818155600180820192909255600290810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055840154610ebd9190611a06565b8360020181905550505050565b60008083158015610ee95750610ee98686868960000160010154611688565b15610efc5750506001840154829061101c565b82158015610f1557508554610f15908790879086611688565b15610f2457505083548161101c565b8315801590610f5157506000848152600387016020526040902060010154610f5190879087908790611688565b15610f735750506000828152600385016020526040902060010154829061101c565b8215801590610f9c57506000838152600387016020526040902054610f9c908790879086611688565b15610fba57505060008181526003850160205260409020548161101c565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f676574206c657373657220616e642067726561746572206661696c75726500006044820152606401610bb4565b94509492505050565b8261108c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4b6579206d75737420626520646566696e6564000000000000000000000000006044820152606401610bb4565b600083815260038501602052604090206002015460ff161561110a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f43616e277420696e7365727420616e206578697374696e6720656c656d656e746044820152606401610bb4565b82821415801561111a5750828114155b6111a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4b65792063616e6e6f74206265207468652073616d652061732070726576696f60448201527f75734b6579206f72206e6578744b6579000000000000000000000000000000006064820152608401610bb4565b60008381526003850160205260408120600280820180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558601549091036111fd5760018501849055838555611554565b8215158061120a57508115155b611296576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4569746865722070726576696f75734b6579206f72206e6578744b6579206d7560448201527f737420626520646566696e6564000000000000000000000000000000000000006064820152608401610bb4565b8281556001810182905582156113f757600083815260038601602052604090206002015460ff16611349576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f49662070726576696f75734b657920697320646566696e65642c206974206d7560448201527f737420657869737420696e20746865206c6973740000000000000000000000006064820152608401610bb4565b60008381526003860160205260409020600181015483146113ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f70726576696f75734b6579206d7573742062652061646a6163656e7420746f2060448201527f6e6578744b6579000000000000000000000000000000000000000000000000006064820152608401610bb4565b6001018490556113ff565b600185018490555b811561155057600082815260038601602052604090206002015460ff166114a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4966206e6578744b657920697320646566696e65642c206974206d757374206560448201527f7869737420696e20746865206c697374000000000000000000000000000000006064820152608401610bb4565b6000828152600386016020526040902080548414611548576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f70726576696f75734b6579206d7573742062652061646a6163656e7420746f2060448201527f6e6578744b6579000000000000000000000000000000000000000000000000006064820152608401610bb4565b849055611554565b8385555b6002850154611564906001611a19565b85600201819055505050505050565b606082600201548211156115e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f6e6f7420656e6f75676820656c656d656e7473000000000000000000000000006044820152606401610bb4565b60008267ffffffffffffffff8111156115fe576115fe611906565b604051908082528060200260200182016040528015611627578160200160208202803683370190505b50845490915060005b8481101561167e578183828151811061164b5761164b611935565b602090810291909101810191909152600092835260038701905260409091205490611677816001611a19565b9050611630565b5090949350505050565b6000808315806116a8575060008481526004870160205260409020548510155b905060008315806116c9575060008481526004880160205260409020548611155b90508180156116d55750805b979650505050505050565b6000602082840312156116f257600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461171d57600080fd5b919050565b60006020828403121561173457600080fd5b610558826116f9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60008151808452602080850194508084016000805b848110156117d357825160048082106117c0577f4e487b7100000000000000000000000000000000000000000000000000000000845260218152602484fd5b5088529683019691830191600101611781565b50959695505050505050565b606080825284519082018190526000906020906080840190828801845b8281101561182e57815173ffffffffffffffffffffffffffffffffffffffff16845292840192908401906001016117fc565b5050508381038285015285518082528683019183019060005b8181101561186357835183529284019291840191600101611847565b50508481036040860152611877818761176c565b98975050505050505050565b6000806040838503121561189657600080fd5b823591506118a6602084016116f9565b90509250929050565b600080600080600060a086880312156118c757600080fd5b853594506118d7602087016116f9565b9350604086013592506118ec606087016116f9565b91506118fa608087016116f9565b90509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036119c4576119c4611964565b5060010190565b600082611a01577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b8181038181111561029657610296611964565b808201808211156102965761029661196456fea164736f6c6343000813000a",
}

// AddressSortedLinkedListWithMedianABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressSortedLinkedListWithMedianMetaData.ABI instead.
var AddressSortedLinkedListWithMedianABI = AddressSortedLinkedListWithMedianMetaData.ABI

// AddressSortedLinkedListWithMedianBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressSortedLinkedListWithMedianMetaData.Bin instead.
var AddressSortedLinkedListWithMedianBin = AddressSortedLinkedListWithMedianMetaData.Bin

// DeployAddressSortedLinkedListWithMedian deploys a new Ethereum contract, binding an instance of AddressSortedLinkedListWithMedian to it.
func DeployAddressSortedLinkedListWithMedian(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressSortedLinkedListWithMedian, error) {
	parsed, err := AddressSortedLinkedListWithMedianMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressSortedLinkedListWithMedianBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressSortedLinkedListWithMedian{AddressSortedLinkedListWithMedianCaller: AddressSortedLinkedListWithMedianCaller{contract: contract}, AddressSortedLinkedListWithMedianTransactor: AddressSortedLinkedListWithMedianTransactor{contract: contract}, AddressSortedLinkedListWithMedianFilterer: AddressSortedLinkedListWithMedianFilterer{contract: contract}}, nil
}

// AddressSortedLinkedListWithMedian is an auto generated Go binding around an Ethereum contract.
type AddressSortedLinkedListWithMedian struct {
	AddressSortedLinkedListWithMedianCaller     // Read-only binding to the contract
	AddressSortedLinkedListWithMedianTransactor // Write-only binding to the contract
	AddressSortedLinkedListWithMedianFilterer   // Log filterer for contract events
}

// AddressSortedLinkedListWithMedianCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressSortedLinkedListWithMedianCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSortedLinkedListWithMedianTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressSortedLinkedListWithMedianTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSortedLinkedListWithMedianFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressSortedLinkedListWithMedianFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSortedLinkedListWithMedianSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSortedLinkedListWithMedianSession struct {
	Contract     *AddressSortedLinkedListWithMedian // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// AddressSortedLinkedListWithMedianCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressSortedLinkedListWithMedianCallerSession struct {
	Contract *AddressSortedLinkedListWithMedianCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// AddressSortedLinkedListWithMedianTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressSortedLinkedListWithMedianTransactorSession struct {
	Contract     *AddressSortedLinkedListWithMedianTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// AddressSortedLinkedListWithMedianRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressSortedLinkedListWithMedianRaw struct {
	Contract *AddressSortedLinkedListWithMedian // Generic contract binding to access the raw methods on
}

// AddressSortedLinkedListWithMedianCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressSortedLinkedListWithMedianCallerRaw struct {
	Contract *AddressSortedLinkedListWithMedianCaller // Generic read-only contract binding to access the raw methods on
}

// AddressSortedLinkedListWithMedianTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressSortedLinkedListWithMedianTransactorRaw struct {
	Contract *AddressSortedLinkedListWithMedianTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressSortedLinkedListWithMedian creates a new instance of AddressSortedLinkedListWithMedian, bound to a specific deployed contract.
func NewAddressSortedLinkedListWithMedian(address common.Address, backend bind.ContractBackend) (*AddressSortedLinkedListWithMedian, error) {
	contract, err := bindAddressSortedLinkedListWithMedian(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressSortedLinkedListWithMedian{AddressSortedLinkedListWithMedianCaller: AddressSortedLinkedListWithMedianCaller{contract: contract}, AddressSortedLinkedListWithMedianTransactor: AddressSortedLinkedListWithMedianTransactor{contract: contract}, AddressSortedLinkedListWithMedianFilterer: AddressSortedLinkedListWithMedianFilterer{contract: contract}}, nil
}

// NewAddressSortedLinkedListWithMedianCaller creates a new read-only instance of AddressSortedLinkedListWithMedian, bound to a specific deployed contract.
func NewAddressSortedLinkedListWithMedianCaller(address common.Address, caller bind.ContractCaller) (*AddressSortedLinkedListWithMedianCaller, error) {
	contract, err := bindAddressSortedLinkedListWithMedian(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressSortedLinkedListWithMedianCaller{contract: contract}, nil
}

// NewAddressSortedLinkedListWithMedianTransactor creates a new write-only instance of AddressSortedLinkedListWithMedian, bound to a specific deployed contract.
func NewAddressSortedLinkedListWithMedianTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressSortedLinkedListWithMedianTransactor, error) {
	contract, err := bindAddressSortedLinkedListWithMedian(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressSortedLinkedListWithMedianTransactor{contract: contract}, nil
}

// NewAddressSortedLinkedListWithMedianFilterer creates a new log filterer instance of AddressSortedLinkedListWithMedian, bound to a specific deployed contract.
func NewAddressSortedLinkedListWithMedianFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressSortedLinkedListWithMedianFilterer, error) {
	contract, err := bindAddressSortedLinkedListWithMedian(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressSortedLinkedListWithMedianFilterer{contract: contract}, nil
}

// bindAddressSortedLinkedListWithMedian binds a generic wrapper to an already deployed contract.
func bindAddressSortedLinkedListWithMedian(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressSortedLinkedListWithMedianABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressSortedLinkedListWithMedian.Contract.AddressSortedLinkedListWithMedianCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressSortedLinkedListWithMedian.Contract.AddressSortedLinkedListWithMedianTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressSortedLinkedListWithMedian.Contract.AddressSortedLinkedListWithMedianTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressSortedLinkedListWithMedian.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressSortedLinkedListWithMedian.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressSortedLinkedListWithMedian.Contract.contract.Transact(opts, method, params...)
}

// ToAddress is a free data retrieval call binding the contract method 0x341f6623.
//
// Solidity: function toAddress(bytes32 b) pure returns(address)
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianCaller) ToAddress(opts *bind.CallOpts, b [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AddressSortedLinkedListWithMedian.contract.Call(opts, &out, "toAddress", b)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToAddress is a free data retrieval call binding the contract method 0x341f6623.
//
// Solidity: function toAddress(bytes32 b) pure returns(address)
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianSession) ToAddress(b [32]byte) (common.Address, error) {
	return _AddressSortedLinkedListWithMedian.Contract.ToAddress(&_AddressSortedLinkedListWithMedian.CallOpts, b)
}

// ToAddress is a free data retrieval call binding the contract method 0x341f6623.
//
// Solidity: function toAddress(bytes32 b) pure returns(address)
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianCallerSession) ToAddress(b [32]byte) (common.Address, error) {
	return _AddressSortedLinkedListWithMedian.Contract.ToAddress(&_AddressSortedLinkedListWithMedian.CallOpts, b)
}

// ToBytes is a free data retrieval call binding the contract method 0x593b79fe.
//
// Solidity: function toBytes(address a) pure returns(bytes32)
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianCaller) ToBytes(opts *bind.CallOpts, a common.Address) ([32]byte, error) {
	var out []interface{}
	err := _AddressSortedLinkedListWithMedian.contract.Call(opts, &out, "toBytes", a)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ToBytes is a free data retrieval call binding the contract method 0x593b79fe.
//
// Solidity: function toBytes(address a) pure returns(bytes32)
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianSession) ToBytes(a common.Address) ([32]byte, error) {
	return _AddressSortedLinkedListWithMedian.Contract.ToBytes(&_AddressSortedLinkedListWithMedian.CallOpts, a)
}

// ToBytes is a free data retrieval call binding the contract method 0x593b79fe.
//
// Solidity: function toBytes(address a) pure returns(bytes32)
func (_AddressSortedLinkedListWithMedian *AddressSortedLinkedListWithMedianCallerSession) ToBytes(a common.Address) ([32]byte, error) {
	return _AddressSortedLinkedListWithMedian.Contract.ToBytes(&_AddressSortedLinkedListWithMedian.CallOpts, a)
}
