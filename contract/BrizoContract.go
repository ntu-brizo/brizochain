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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BrizoABI is the input ABI used to generate the binding from.
const BrizoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"content\",\"type\":\"string\"}],\"name\":\"writeData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"readData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashKey\",\"type\":\"string\"}],\"name\":\"readDataFromHashDict\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hashKey\",\"type\":\"string\"},{\"name\":\"content\",\"type\":\"string\"}],\"name\":\"writeDataToHashDict\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BrizoFuncSigs maps the 4-byte function signature to its string representation.
var BrizoFuncSigs = map[string]string{
	"8ada066e": "getCounter()",
	"62630e83": "readData(address,uint256)",
	"832cbfb5": "readDataFromHashDict(string)",
	"134bc876": "writeData(string)",
	"bb96db73": "writeDataToHashDict(string,string)",
}

// BrizoBin is the compiled bytecode used for deploying new contracts.
var BrizoBin = "0x6080604052600060025534801561001557600080fd5b50610a81806100256000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663134bc876811461007157806362630e8314610093578063832cbfb5146100c95780638ada066e146100f6578063bb96db7314610118575b600080fd5b34801561007d57600080fd5b5061009161008c3660046106cc565b610138565b005b34801561009f57600080fd5b506100b36100ae366004610692565b6101a6565b6040516100c09190610930565b60405180910390f35b3480156100d557600080fd5b506100e96100e43660046106cc565b610312565b6040516100c0919061091f565b34801561010257600080fd5b5061010b6104c8565b6040516100c09190610981565b34801561012457600080fd5b50610091610133366004610709565b6104cf565b80516000106101655760405160e560020a62461bcd02815260040161015c90610941565b60405180910390fd5b33600090815260208181526040822080546001810180835591845292829020845191936101989391019190850190610587565b505060028054600101905550565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040812054606091106101ef5760405160e560020a62461bcd02815260040161015c90610961565b60008210158015610224575073ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205482105b15156102455760405160e560020a62461bcd02815260040161015c90610971565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902080548390811061027657fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156103045780601f106102d957610100808354040283529160200191610304565b820191906000526020600020905b8154815290600101906020018083116102e757829003601f168201915b505050505090505b92915050565b606060006001836040518082805190602001908083835b602083106103485780518252601f199092019160209182019101610329565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205492909211915061039b90505760405160e560020a62461bcd02815260040161015c90610951565b6001826040518082805190602001908083835b602083106103cd5780518252601f1990920191602091820191016103ae565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185208054808402870184019092528186529350915060009084015b828210156104bd5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156104a95780601f1061047e576101008083540402835291602001916104a9565b820191906000526020600020905b81548152906001019060200180831161048c57829003601f168201915b505050505081526020019060010190610412565b505050509050919050565b6002545b90565b80516000106104f35760405160e560020a62461bcd02815260040161015c90610941565b6001826040518082805190602001908083835b602083106105255780518252601f199092019160209182019101610506565b51815160209384036101000a600019018019909216911617905292019485525060405193849003810190932080546001810180835560009283529185902086519295610578955091019250850190610587565b50506002805460010190555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105c857805160ff19168380011785556105f5565b828001600101855582156105f5579182015b828111156105f55782518255916020019190600101906105da565b50610601929150610605565b5090565b6104cc91905b80821115610601576000815560010161060b565b600061062b82356109e8565b9392505050565b6000601f8201831361064357600080fd5b8135610656610651826109b6565b61098f565b9150808252602083016020830185838301111561067257600080fd5b61067d838284610a01565b50505092915050565b600061062b82356104cc565b600080604083850312156106a557600080fd5b60006106b1858561061f565b92505060206106c285828601610686565b9150509250929050565b6000602082840312156106de57600080fd5b813567ffffffffffffffff8111156106f557600080fd5b61070184828501610632565b949350505050565b6000806040838503121561071c57600080fd5b823567ffffffffffffffff81111561073357600080fd5b61073f85828601610632565b925050602083013567ffffffffffffffff81111561075c57600080fd5b6106c285828601610632565b6000610773826109e4565b8084526020840193508360208202850161078c856109de565b60005b848110156107c35783830388526107a78383516107cf565b92506107b2826109de565b60209890980197915060010161078f565b50909695505050505050565b60006107da826109e4565b8084526107ee816020860160208601610a0d565b6107f781610a3d565b9093016020019392505050565b601181527f436f6e74656e7420697320656d70747921000000000000000000000000000000602082015260400190565b602881527f4e6f206461746120776173207772697474656e207573696e672074686973206860208201527f617368206b657921000000000000000000000000000000000000000000000000604082015260600190565b602481527f4e6f206461746120776173207772697474656e2062792074686973206164647260208201527f6573732100000000000000000000000000000000000000000000000000000000604082015260600190565b601381527f496e646578206f7574206f662072616e67652100000000000000000000000000602082015260400190565b610919816104cc565b82525050565b6020808252810161062b8184610768565b6020808252810161062b81846107cf565b6020808252810161030c81610804565b6020808252810161030c81610834565b6020808252810161030c8161088a565b6020808252810161030c816108e0565b6020810161030c8284610910565b60405181810167ffffffffffffffff811182821017156109ae57600080fd5b604052919050565b600067ffffffffffffffff8211156109cd57600080fd5b506020601f91909101601f19160190565b60200190565b5190565b73ffffffffffffffffffffffffffffffffffffffff1690565b82818337506000910152565b60005b83811015610a28578181015183820152602001610a10565b83811115610a37576000848401525b50505050565b601f01601f1916905600a265627a7a723058201d3ca0728b9cc772979e5055e6330a03e23ee5260a94f350d046b299f1a2cb636c6578706572696d656e74616cf50037"

// DeployBrizo deploys a new Ethereum contract, binding an instance of Brizo to it.
func DeployBrizo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Brizo, error) {
	parsed, err := abi.JSON(strings.NewReader(BrizoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BrizoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Brizo{BrizoCaller: BrizoCaller{contract: contract}, BrizoTransactor: BrizoTransactor{contract: contract}, BrizoFilterer: BrizoFilterer{contract: contract}}, nil
}

// Brizo is an auto generated Go binding around an Ethereum contract.
type Brizo struct {
	BrizoCaller     // Read-only binding to the contract
	BrizoTransactor // Write-only binding to the contract
	BrizoFilterer   // Log filterer for contract events
}

// BrizoCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrizoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrizoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrizoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrizoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrizoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrizoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrizoSession struct {
	Contract     *Brizo            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrizoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrizoCallerSession struct {
	Contract *BrizoCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BrizoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrizoTransactorSession struct {
	Contract     *BrizoTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrizoRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrizoRaw struct {
	Contract *Brizo // Generic contract binding to access the raw methods on
}

// BrizoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrizoCallerRaw struct {
	Contract *BrizoCaller // Generic read-only contract binding to access the raw methods on
}

// BrizoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrizoTransactorRaw struct {
	Contract *BrizoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrizo creates a new instance of Brizo, bound to a specific deployed contract.
func NewBrizo(address common.Address, backend bind.ContractBackend) (*Brizo, error) {
	contract, err := bindBrizo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Brizo{BrizoCaller: BrizoCaller{contract: contract}, BrizoTransactor: BrizoTransactor{contract: contract}, BrizoFilterer: BrizoFilterer{contract: contract}}, nil
}

// NewBrizoCaller creates a new read-only instance of Brizo, bound to a specific deployed contract.
func NewBrizoCaller(address common.Address, caller bind.ContractCaller) (*BrizoCaller, error) {
	contract, err := bindBrizo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrizoCaller{contract: contract}, nil
}

// NewBrizoTransactor creates a new write-only instance of Brizo, bound to a specific deployed contract.
func NewBrizoTransactor(address common.Address, transactor bind.ContractTransactor) (*BrizoTransactor, error) {
	contract, err := bindBrizo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrizoTransactor{contract: contract}, nil
}

// NewBrizoFilterer creates a new log filterer instance of Brizo, bound to a specific deployed contract.
func NewBrizoFilterer(address common.Address, filterer bind.ContractFilterer) (*BrizoFilterer, error) {
	contract, err := bindBrizo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrizoFilterer{contract: contract}, nil
}

// bindBrizo binds a generic wrapper to an already deployed contract.
func bindBrizo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrizoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brizo *BrizoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Brizo.Contract.BrizoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brizo *BrizoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brizo.Contract.BrizoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brizo *BrizoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brizo.Contract.BrizoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brizo *BrizoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Brizo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brizo *BrizoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brizo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brizo *BrizoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brizo.Contract.contract.Transact(opts, method, params...)
}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() constant returns(uint256)
func (_Brizo *BrizoCaller) GetCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Brizo.contract.Call(opts, out, "getCounter")
	return *ret0, err
}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() constant returns(uint256)
func (_Brizo *BrizoSession) GetCounter() (*big.Int, error) {
	return _Brizo.Contract.GetCounter(&_Brizo.CallOpts)
}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() constant returns(uint256)
func (_Brizo *BrizoCallerSession) GetCounter() (*big.Int, error) {
	return _Brizo.Contract.GetCounter(&_Brizo.CallOpts)
}

// ReadData is a free data retrieval call binding the contract method 0x62630e83.
//
// Solidity: function readData(address addr, uint256 index) constant returns(string)
func (_Brizo *BrizoCaller) ReadData(opts *bind.CallOpts, addr common.Address, index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Brizo.contract.Call(opts, out, "readData", addr, index)
	return *ret0, err
}

// ReadData is a free data retrieval call binding the contract method 0x62630e83.
//
// Solidity: function readData(address addr, uint256 index) constant returns(string)
func (_Brizo *BrizoSession) ReadData(addr common.Address, index *big.Int) (string, error) {
	return _Brizo.Contract.ReadData(&_Brizo.CallOpts, addr, index)
}

// ReadData is a free data retrieval call binding the contract method 0x62630e83.
//
// Solidity: function readData(address addr, uint256 index) constant returns(string)
func (_Brizo *BrizoCallerSession) ReadData(addr common.Address, index *big.Int) (string, error) {
	return _Brizo.Contract.ReadData(&_Brizo.CallOpts, addr, index)
}

// ReadDataFromHashDict is a free data retrieval call binding the contract method 0x832cbfb5.
//
// Solidity: function readDataFromHashDict(string hashKey) constant returns(string[])
func (_Brizo *BrizoCaller) ReadDataFromHashDict(opts *bind.CallOpts, hashKey string) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Brizo.contract.Call(opts, out, "readDataFromHashDict", hashKey)
	return *ret0, err
}

// ReadDataFromHashDict is a free data retrieval call binding the contract method 0x832cbfb5.
//
// Solidity: function readDataFromHashDict(string hashKey) constant returns(string[])
func (_Brizo *BrizoSession) ReadDataFromHashDict(hashKey string) ([]string, error) {
	return _Brizo.Contract.ReadDataFromHashDict(&_Brizo.CallOpts, hashKey)
}

// ReadDataFromHashDict is a free data retrieval call binding the contract method 0x832cbfb5.
//
// Solidity: function readDataFromHashDict(string hashKey) constant returns(string[])
func (_Brizo *BrizoCallerSession) ReadDataFromHashDict(hashKey string) ([]string, error) {
	return _Brizo.Contract.ReadDataFromHashDict(&_Brizo.CallOpts, hashKey)
}

// WriteData is a paid mutator transaction binding the contract method 0x134bc876.
//
// Solidity: function writeData(string content) returns()
func (_Brizo *BrizoTransactor) WriteData(opts *bind.TransactOpts, content string) (*types.Transaction, error) {
	return _Brizo.contract.Transact(opts, "writeData", content)
}

// WriteData is a paid mutator transaction binding the contract method 0x134bc876.
//
// Solidity: function writeData(string content) returns()
func (_Brizo *BrizoSession) WriteData(content string) (*types.Transaction, error) {
	return _Brizo.Contract.WriteData(&_Brizo.TransactOpts, content)
}

// WriteData is a paid mutator transaction binding the contract method 0x134bc876.
//
// Solidity: function writeData(string content) returns()
func (_Brizo *BrizoTransactorSession) WriteData(content string) (*types.Transaction, error) {
	return _Brizo.Contract.WriteData(&_Brizo.TransactOpts, content)
}

// WriteDataToHashDict is a paid mutator transaction binding the contract method 0xbb96db73.
//
// Solidity: function writeDataToHashDict(string hashKey, string content) returns()
func (_Brizo *BrizoTransactor) WriteDataToHashDict(opts *bind.TransactOpts, hashKey string, content string) (*types.Transaction, error) {
	return _Brizo.contract.Transact(opts, "writeDataToHashDict", hashKey, content)
}

// WriteDataToHashDict is a paid mutator transaction binding the contract method 0xbb96db73.
//
// Solidity: function writeDataToHashDict(string hashKey, string content) returns()
func (_Brizo *BrizoSession) WriteDataToHashDict(hashKey string, content string) (*types.Transaction, error) {
	return _Brizo.Contract.WriteDataToHashDict(&_Brizo.TransactOpts, hashKey, content)
}

// WriteDataToHashDict is a paid mutator transaction binding the contract method 0xbb96db73.
//
// Solidity: function writeDataToHashDict(string hashKey, string content) returns()
func (_Brizo *BrizoTransactorSession) WriteDataToHashDict(hashKey string, content string) (*types.Transaction, error) {
	return _Brizo.Contract.WriteDataToHashDict(&_Brizo.TransactOpts, hashKey, content)
}
