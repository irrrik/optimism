// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const GoldTokenStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/celo/GoldToken.sol:GoldToken\",\"label\":\"initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1001,\"contract\":\"src/celo/GoldToken.sol:GoldToken\",\"label\":\"_owner\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"src/celo/GoldToken.sol:GoldToken\",\"label\":\"registry\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_contract(ICeloRegistry)1005\"},{\"astId\":1003,\"contract\":\"src/celo/GoldToken.sol:GoldToken\",\"label\":\"totalSupply_\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1004,\"contract\":\"src/celo/GoldToken.sol:GoldToken\",\"label\":\"allowed\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(ICeloRegistry)1005\":{\"encoding\":\"inplace\",\"label\":\"contract ICeloRegistry\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var GoldTokenStorageLayout = new(solc.StorageLayout)

var GoldTokenDeployedBin = "0x608060405234801561001057600080fd5b50600436106101a35760003560e01c8063715018a6116100ee578063a9059cbb11610097578063c4d66de811610071578063c4d66de8146103e6578063dd62ed3e146103f9578063e1d6aceb1461043f578063f2fde38b1461045257600080fd5b8063a9059cbb146103ad578063a91ee0dc146103c0578063b921e163146103d357600080fd5b80639358928b116100c85780639358928b1461035957806395d89b4114610361578063a457c2d71461039a57600080fd5b8063715018a6146102e75780637b103999146102f15780638da5cb5b1461033657600080fd5b8063313ce5671161015057806342966c681161012a57806342966c681461028657806354255be01461029957806370a08231146102bf57600080fd5b8063313ce56714610251578063395093511461026057806340c10f191461027357600080fd5b806318160ddd1161018157806318160ddd1461022357806323b872dd14610235578063265126bd1461024857600080fd5b806306fdde03146101a8578063095ea7b3146101f3578063158ef93e14610216575b600080fd5b60408051808201909152601181527f43656c6f206e617469766520617373657400000000000000000000000000000060208201525b6040516101ea91906114aa565b60405180910390f35b610206610201366004611524565b610465565b60405190151581526020016101ea565b6000546102069060ff1681565b6002545b6040519081526020016101ea565b61020661024336600461154e565b61055b565b61dead31610227565b604051601281526020016101ea565b61020661026e366004611524565b610943565b610206610281366004611524565b610a73565b61020661029436600461158a565b610d1e565b6040805160018082526020820152600291810191909152600060608201526080016101ea565b6102276102cd3660046115a3565b73ffffffffffffffffffffffffffffffffffffffff163190565b6102ef610d2c565b005b6001546103119073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101ea565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16610311565b610227610d40565b60408051808201909152600481527f43454c4f0000000000000000000000000000000000000000000000000000000060208201526101dd565b6102066103a8366004611524565b610d5f565b6102066103bb366004611524565b610d9b565b6102ef6103ce3660046115a3565b610dae565b6102ef6103e136600461158a565b610ea2565b6102ef6103f43660046115a3565b610f1e565b6102276104073660046115be565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260036020908152604080832093909416825291909152205490565b61020661044d3660046115f1565b610fce565b6102ef6104603660046115a3565b61101f565b600073ffffffffffffffffffffffffffffffffffffffff83166104e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f63616e6e6f742073657420616c6c6f77616e636520666f72203000000000000060448201526064015b60405180910390fd5b33600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff881680855290835292819020869055518581529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a35060015b92915050565b600073ffffffffffffffffffffffffffffffffffffffff8316610600576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f7472616e7366657220617474656d7074656420746f207265736572766564206160448201527f646472657373203078300000000000000000000000000000000000000000000060648201526084016104e0565b73ffffffffffffffffffffffffffffffffffffffff8416318211156106a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f7472616e736665722076616c75652065786365656465642062616c616e63652060448201527f6f662073656e646572000000000000000000000000000000000000000000000060648201526084016104e0565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600360209081526040808320338452909152902054821115610767576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f7472616e736665722076616c75652065786365656465642073656e646572277360448201527f20616c6c6f77616e636520666f72207370656e6465720000000000000000000060648201526084016104e0565b600060fd815a6040805173ffffffffffffffffffffffffffffffffffffffff808b16602083015289169181019190915260608101879052909190608001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107dc91611678565b600060405180830381858888f193505050503d806000811461081a576040519150601f19603f3d011682016040523d82523d6000602084013e61081f565b606091505b5050809150508061088c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43454c4f207472616e73666572206661696c656400000000000000000000000060448201526064016104e0565b73ffffffffffffffffffffffffffffffffffffffff851660009081526003602090815260408083203384529091529020546108c89084906116c3565b73ffffffffffffffffffffffffffffffffffffffff868116600081815260036020908152604080832033845282529182902094909455518681529187169290917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3506001949350505050565b905090565b600073ffffffffffffffffffffffffffffffffffffffff83166109c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f63616e6e6f742073657420616c6c6f77616e636520666f72203000000000000060448201526064016104e0565b33600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054906109fe84836116d6565b33600081815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8b16808552908352928190208590555184815293945090927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3506001949350505050565b60003315610add576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f4f6e6c7920564d2063616e2063616c6c0000000000000000000000000000000060448201526064016104e0565b81600003610aed57506001610555565b73ffffffffffffffffffffffffffffffffffffffff8316610b90576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f6d696e7420617474656d7074656420746f20726573657276656420616464726560448201527f737320307830000000000000000000000000000000000000000000000000000060648201526084016104e0565b81600254610b9e91906116d6565b600255600060fd815a604080516000602082015273ffffffffffffffffffffffffffffffffffffffff89169181019190915260608101879052909190608001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610c1591611678565b600060405180830381858888f193505050503d8060008114610c53576040519150601f19603f3d011682016040523d82523d6000602084013e610c58565b606091505b50508091505080610cc5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43454c4f207472616e73666572206661696c656400000000000000000000000060448201526064016104e0565b60405183815273ffffffffffffffffffffffffffffffffffffffff8516906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020015b60405180910390a35060019392505050565b600061055561dead836110d3565b610d346112d3565b610d3e600061135a565b565b6000803161dead31600254610d5591906116c3565b61093e91906116c3565b33600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152812054816109fe84836116c3565b6000610da783836113d7565b9392505050565b610db66112d3565b73ffffffffffffffffffffffffffffffffffffffff8116610e33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f43616e6e6f7420726567697374657220746865206e756c6c206164647265737360448201526064016104e0565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b90600090a250565b3315610f0a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f4f6e6c7920564d2063616e2063616c6c0000000000000000000000000000000060448201526064016104e0565b80600254610f1891906116d6565b60025550565b60005460ff1615610f8b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f636f6e747261637420616c726561647920696e697469616c697a65640000000060448201526064016104e0565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001178155600255610fc23361135a565b610fcb81610dae565b50565b600080610fdb86866113d7565b90507fe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc848460405161100e9291906116e9565b60405180910390a195945050505050565b6110276112d3565b73ffffffffffffffffffffffffffffffffffffffff81166110ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104e0565b610fcb8161135a565b60003331821115611166576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f7472616e736665722076616c75652065786365656465642062616c616e63652060448201527f6f662073656e646572000000000000000000000000000000000000000000000060648201526084016104e0565b600060fd815a6040805133602082015273ffffffffffffffffffffffffffffffffffffffff89169181019190915260608101879052909190608001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526111d991611678565b600060405180830381858888f193505050503d8060008114611217576040519150601f19603f3d011682016040523d82523d6000602084013e61121c565b606091505b50508091505080611289576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43454c4f207472616e73666572206661696c656400000000000000000000000060448201526064016104e0565b60405183815273ffffffffffffffffffffffffffffffffffffffff85169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610d0c565b60005473ffffffffffffffffffffffffffffffffffffffff610100909104163314610d3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104e0565b6000805473ffffffffffffffffffffffffffffffffffffffff8381166101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff851617855560405193049190911692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a35050565b600073ffffffffffffffffffffffffffffffffffffffff831661147c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f7472616e7366657220617474656d7074656420746f207265736572766564206160448201527f646472657373203078300000000000000000000000000000000000000000000060648201526084016104e0565b610da783836110d3565b60005b838110156114a1578181015183820152602001611489565b50506000910152565b60208152600082518060208401526114c9816040850160208701611486565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461151f57600080fd5b919050565b6000806040838503121561153757600080fd5b611540836114fb565b946020939093013593505050565b60008060006060848603121561156357600080fd5b61156c846114fb565b925061157a602085016114fb565b9150604084013590509250925092565b60006020828403121561159c57600080fd5b5035919050565b6000602082840312156115b557600080fd5b610da7826114fb565b600080604083850312156115d157600080fd5b6115da836114fb565b91506115e8602084016114fb565b90509250929050565b6000806000806060858703121561160757600080fd5b611610856114fb565b935060208501359250604085013567ffffffffffffffff8082111561163457600080fd5b818701915087601f83011261164857600080fd5b81358181111561165757600080fd5b88602082850101111561166957600080fd5b95989497505060200194505050565b6000825161168a818460208701611486565b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055557610555611694565b8082018082111561055557610555611694565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010191905056fea164736f6c6343000813000a"


func init() {
	if err := json.Unmarshal([]byte(GoldTokenStorageLayoutJSON), GoldTokenStorageLayout); err != nil {
		panic(err)
	}

	layouts["GoldToken"] = GoldTokenStorageLayout
	deployedBytecodes["GoldToken"] = GoldTokenDeployedBin
	immutableReferences["GoldToken"] = false
}