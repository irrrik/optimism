// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const BridgedETHStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/celo/BridgedETH.sol:BridgedETH\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/celo/BridgedETH.sol:BridgedETH\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1002,\"contract\":\"src/celo/BridgedETH.sol:BridgedETH\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/celo/BridgedETH.sol:BridgedETH\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1004,\"contract\":\"src/celo/BridgedETH.sol:BridgedETH\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var BridgedETHStorageLayout = new(solc.StorageLayout)

var BridgedETHDeployedBin = "0x608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e14610395578063e78cea9214610349578063ee9a31a2146103db57600080fd5b8063ae1f6aaf14610349578063c01e1bd61461036f578063d6c0b2c41461036f57600080fd5b80639dc29fac116100bd5780639dc29fac14610310578063a457c2d714610323578063a9059cbb1461033657600080fd5b806370a08231146102d257806395d89b411461030857600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461026e57806340c10f191461028157806354fd4d501461029657600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a36600461117d565b610402565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f86104f3565b60405161019b91906111c6565b61018f610213366004611262565b610585565b6002545b60405190815260200161019b565b61018f61023836600461128c565b61059d565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161019b565b61018f61027c366004611262565b6105c1565b61029461028f366004611262565b61060d565b005b6101f86040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b61021c6102e03660046112c8565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101f8610735565b61029461031e366004611262565b610744565b61018f610331366004611262565b61085b565b61018f610344366004611262565b61092c565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c6103a33660046112e3565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000085168314806104bb57507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b806104ea57507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b60606003805461050290611316565b80601f016020809104026020016040519081016040528092919081815260200182805461052e90611316565b801561057b5780601f106105505761010080835404028352916020019161057b565b820191906000526020600020905b81548152906001019060200180831161055e57829003601f168201915b5050505050905090565b60003361059381858561093a565b5060019392505050565b6000336105ab858285610aee565b6105b6858585610bc5565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906105939082908690610608908790611398565b61093a565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084015b60405180910390fd5b6106e18282610e78565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161072991815260200190565b60405180910390a25050565b60606004805461050290611316565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610809576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084016106ce565b6108138282610f98565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405161072991815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561091f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016106ce565b6105b6828686840361093a565b600033610593818585610bc5565b73ffffffffffffffffffffffffffffffffffffffff83166109dc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff8216610a7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610bbf5781811015610bb2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016106ce565b610bbf848484840361093a565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c68576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff8216610d0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610dc1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220858503905591851681529081208054849290610e05908490611398565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610e6b91815260200190565b60405180910390a3610bbf565b73ffffffffffffffffffffffffffffffffffffffff8216610ef5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016106ce565b8060026000828254610f079190611398565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610f41908490611398565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff821661103b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040902054818110156110f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016106ce565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040812083830390556002805484929061112d9084906113b0565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610ae1565b60006020828403121561118f57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146111bf57600080fd5b9392505050565b600060208083528351808285015260005b818110156111f3578581018301518582016040015282016111d7565b81811115611205576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461125d57600080fd5b919050565b6000806040838503121561127557600080fd5b61127e83611239565b946020939093013593505050565b6000806000606084860312156112a157600080fd5b6112aa84611239565b92506112b860208501611239565b9150604084013590509250925092565b6000602082840312156112da57600080fd5b6111bf82611239565b600080604083850312156112f657600080fd5b6112ff83611239565b915061130d60208401611239565b90509250929050565b600181811c9082168061132a57607f821691505b602082108103611363577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156113ab576113ab611369565b500190565b6000828210156113c2576113c2611369565b50039056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(BridgedETHStorageLayoutJSON), BridgedETHStorageLayout); err != nil {
		panic(err)
	}

	layouts["BridgedETH"] = BridgedETHStorageLayout
	deployedBytecodes["BridgedETH"] = BridgedETHDeployedBin
	immutableReferences["BridgedETH"] = true
}
