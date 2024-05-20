// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const FeeCurrencyWhitelistStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/celo/FeeCurrencyWhitelist.sol:FeeCurrencyWhitelist\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/celo/FeeCurrencyWhitelist.sol:FeeCurrencyWhitelist\",\"label\":\"initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/celo/FeeCurrencyWhitelist.sol:FeeCurrencyWhitelist\",\"label\":\"whitelist\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_address)dyn_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_address)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"address[]\",\"numberOfBytes\":\"32\",\"base\":\"t_address\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"}}}"

var FeeCurrencyWhitelistStorageLayout = new(solc.StorageLayout)

var FeeCurrencyWhitelistDeployedBin = "0x608060405234801561001057600080fd5b50600436106100be5760003560e01c80638129fc1c11610076578063d01f63f51161005b578063d01f63f51461019e578063d48bfca7146101b3578063f2fde38b146101c657600080fd5b80638129fc1c146101785780638da5cb5b1461018057600080fd5b806354255be0116100a757806354255be014610112578063715018a6146101385780637ebd1b301461014057600080fd5b806313baf1e6146100c3578063158ef93e146100d8575b600080fd5b6100d66100d1366004610800565b6101d9565b005b6000546100fd9074010000000000000000000000000000000000000000900460ff1681565b60405190151581526020015b60405180910390f35b600180806000604080519485526020850193909352918301526060820152608001610109565b6100d66103f3565b61015361014e36600461082a565b610407565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610109565b6100d661043e565b60005473ffffffffffffffffffffffffffffffffffffffff16610153565b6101a661050b565b6040516101099190610843565b6100d66101c136600461089d565b61057a565b6100d66101d436600461089d565b61062a565b6101e16106e1565b8173ffffffffffffffffffffffffffffffffffffffff166001828154811061020b5761020b6108bf565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1614610299576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e64657820646f6573206e6f74206d6174636800000000000000000000000060448201526064015b60405180910390fd5b60018054906102a881836108ee565b815481106102b8576102b86108bf565b6000918252602090912001546001805473ffffffffffffffffffffffffffffffffffffffff90921691849081106102f1576102f16108bf565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600180548061034a5761034a61092e565b60008281526020908190207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908301810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590910190915560405173ffffffffffffffffffffffffffffffffffffffff851681527fc1f06ffbe5c19d22daa982fd4b3886ced05d83e2bfe7de3b67222728f5234e28910160405180910390a1505050565b6103fb6106e1565b6104056000610762565b565b6001818154811061041757600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60005474010000000000000000000000000000000000000000900460ff16156104c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f636f6e747261637420616c726561647920696e697469616c697a6564000000006044820152606401610290565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905561040533610762565b6060600180548060200260200160405190810160405280929190818152602001828054801561057057602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610545575b5050505050905090565b6105826106e1565b6001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fcf4fe1d1989a39011040b0c33ba1165fdeeca971a1ab2b0340b23550f93727e09060200160405180910390a150565b6106326106e1565b73ffffffffffffffffffffffffffffffffffffffff81166106d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610290565b6106de81610762565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610405576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610290565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107fb57600080fd5b919050565b6000806040838503121561081357600080fd5b61081c836107d7565b946020939093013593505050565b60006020828403121561083c57600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b8181101561089157835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161085f565b50909695505050505050565b6000602082840312156108af57600080fd5b6108b8826107d7565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b81810381811115610928577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a"


func init() {
	if err := json.Unmarshal([]byte(FeeCurrencyWhitelistStorageLayoutJSON), FeeCurrencyWhitelistStorageLayout); err != nil {
		panic(err)
	}

	layouts["FeeCurrencyWhitelist"] = FeeCurrencyWhitelistStorageLayout
	deployedBytecodes["FeeCurrencyWhitelist"] = FeeCurrencyWhitelistDeployedBin
	immutableReferences["FeeCurrencyWhitelist"] = false
}