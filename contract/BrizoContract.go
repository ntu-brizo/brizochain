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
const BrizoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"content\",\"type\":\"string\"}],\"name\":\"writeData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"readData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashKey\",\"type\":\"string\"}],\"name\":\"readDataFromHashDict\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hashKey\",\"type\":\"string\"},{\"name\":\"content\",\"type\":\"string\"}],\"name\":\"writeDataToHashDict\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BrizoFuncSigs maps the 4-byte function signature to its string representation.
var BrizoFuncSigs = map[string]string{
	"8ada066e": "getCounter()",
	"62630e83": "readData(address,uint256)",
	"832cbfb5": "readDataFromHashDict(string)",
	"134bc876": "writeData(string)",
	"bb96db73": "writeDataToHashDict(string,string)",
}

// BrizoBin is the compiled bytecode used for deploying new contracts.
var BrizoBin = "0x6080604052600060025534801561001557600080fd5b506108f7806100256000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663134bc876811461007157806362630e83146100cc578063832cbfb5146101725780638ada066e146101cb578063bb96db73146101f2575b600080fd5b34801561007d57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100ca9436949293602493928401919081908401838280828437509497506102899650505050505050565b005b3480156100d857600080fd5b506100fd73ffffffffffffffffffffffffffffffffffffffff60043516602435610323565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561013757818101518382015260200161011f565b50505050905090810190601f1680156101645780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561017e57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100fd94369492936024939284019190819084018382808284375094975061051d9650505050505050565b3480156101d757600080fd5b506101e061070a565b60408051918252519081900360200190f35b3480156101fe57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100ca94369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506107119650505050505050565b80516000106102e2576040805160e560020a62461bcd02815260206004820152601d60248201527f54686520636f6e74656e74206f66206461746120697320656d70747921000000604482015290519081900360640190fd5b33600090815260208181526040822080546001810180835591845292829020845191936103159391019190850190610833565b505060028054600101905550565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040812054606091106103c6576040805160e560020a62461bcd028152602060048201526024808201527f4e6f206461746120776173207772697474656e2062792074686973206164647260448201527f6573732100000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600082101580156103fb575073ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205482105b1515610451576040805160e560020a62461bcd02815260206004820152601360248201527f496e646578206f7574206f662072616e67652100000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902080548390811061048257fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156105105780601f106104e557610100808354040283529160200191610510565b820191906000526020600020905b8154815290600101906020018083116104f357829003601f168201915b5050505050905092915050565b606060006001836040518082805190602001908083835b602083106105535780518252601f199092019160209182019101610534565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020805460018160011615610100020316600290049050111515610618576040805160e560020a62461bcd02815260206004820152602860248201527f4e6f206461746120776173207772697474656e207573696e672074686973206860448201527f617368206b657921000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6001826040518082805190602001908083835b6020831061064a5780518252601f19909201916020918201910161062b565b518151600019602094850361010090810a820192831692199390931691909117909252949092019687526040805197889003820188208054601f60026001831615909802909501169590950492830182900482028801820190528187529294509250508301828280156106fe5780601f106106d3576101008083540402835291602001916106fe565b820191906000526020600020905b8154815290600101906020018083116106e157829003601f168201915b50505050509050919050565b6002545b90565b815160001061076a576040805160e560020a62461bcd02815260206004820152601660248201527f4e6f2068617368206b6579207370656369666965642100000000000000000000604482015290519081900360640190fd5b80516000106107c3576040805160e560020a62461bcd02815260206004820152601160248201527f436f6e74656e7420697320656d70747921000000000000000000000000000000604482015290519081900360640190fd5b806001836040518082805190602001908083835b602083106107f65780518252601f1990920191602091820191016107d7565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101909320845161031595919491909101925090505b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061087457805160ff19168380011785556108a1565b828001600101855582156108a1579182015b828111156108a1578251825591602001919060010190610886565b506108ad9291506108b1565b5090565b61070e91905b808211156108ad57600081556001016108b75600a165627a7a7230582034190950a78687be8d3ffe28a96369473715721eaa011545157feceda2dfe1ca0029"

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
// Solidity: function readDataFromHashDict(string hashKey) constant returns(string)
func (_Brizo *BrizoCaller) ReadDataFromHashDict(opts *bind.CallOpts, hashKey string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Brizo.contract.Call(opts, out, "readDataFromHashDict", hashKey)
	return *ret0, err
}

// ReadDataFromHashDict is a free data retrieval call binding the contract method 0x832cbfb5.
//
// Solidity: function readDataFromHashDict(string hashKey) constant returns(string)
func (_Brizo *BrizoSession) ReadDataFromHashDict(hashKey string) (string, error) {
	return _Brizo.Contract.ReadDataFromHashDict(&_Brizo.CallOpts, hashKey)
}

// ReadDataFromHashDict is a free data retrieval call binding the contract method 0x832cbfb5.
//
// Solidity: function readDataFromHashDict(string hashKey) constant returns(string)
func (_Brizo *BrizoCallerSession) ReadDataFromHashDict(hashKey string) (string, error) {
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
