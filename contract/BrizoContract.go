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
const BrizoABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"writeData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"readData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"dataCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BrizoFuncSigs maps the 4-byte function signature to its string representation.
var BrizoFuncSigs = map[string]string{
	"d044a6be": "dataCounter(address)",
	"62630e83": "readData(address,uint256)",
	"134bc876": "writeData(string)",
}

// BrizoBin is the compiled bytecode used for deploying new contracts.
var BrizoBin = "0x608060405234801561001057600080fd5b506104db806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063134bc8761461004657806362630e83146100ee578063d044a6be1461018f575b600080fd5b6100ec6004803603602081101561005c57600080fd5b81019060208101813564010000000081111561007757600080fd5b82018360208201111561008957600080fd5b803590602001918460018302840111640100000000831117156100ab57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506101c7945050505050565b005b61011a6004803603604081101561010457600080fd5b506001600160a01b038135169060200135610255565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015457818101518382015260200161013c565b50505050905090810190601f1680156101815780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101b5600480360360208110156101a557600080fd5b50356001600160a01b03166103cd565b60408051918252519081900360200190f35b600081511161021d576040805162461bcd60e51b815260206004820152601d60248201527f54686520636f6e74656e74206f66206461746120697320656d70747921000000604482015290519081900360640190fd5b336000908152602081815260408220805460018101808355918452928290208451919361025093910191908501906103e8565b505050565b6001600160a01b0382166000908152602081905260409020546060906102ac5760405162461bcd60e51b81526004018080602001828103825260238152602001806104846023913960400191505060405180910390fd5b6001600160a01b038316600090815260208190526040902054821061030e576040805162461bcd60e51b8152602060048201526013602482015272496e646578206f7574206f662072616e67652160681b604482015290519081900360640190fd5b6001600160a01b038316600090815260208190526040902080548390811061033257fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156103c05780601f10610395576101008083540402835291602001916103c0565b820191906000526020600020905b8154815290600101906020018083116103a357829003601f168201915b5050505050905092915050565b6001600160a01b031660009081526020819052604090205490565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061042957805160ff1916838001178555610456565b82800160010185558215610456579182015b8281111561045657825182559160200191906001019061043b565b50610462929150610466565b5090565b61048091905b80821115610462576000815560010161046c565b9056fe4e6f206461746120776173207772697474656e20627920746869732061646472657373a265627a7a72315820de5bf82c3eb3c40bfca84d7f44f100cb8cdf1ded636010973fb5881feb4796dd64736f6c634300050b0032"

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

// DataCounter is a free data retrieval call binding the contract method 0xd044a6be.
//
// Solidity: function dataCounter(address addr) constant returns(uint256)
func (_Brizo *BrizoCaller) DataCounter(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Brizo.contract.Call(opts, out, "dataCounter", addr)
	return *ret0, err
}

// DataCounter is a free data retrieval call binding the contract method 0xd044a6be.
//
// Solidity: function dataCounter(address addr) constant returns(uint256)
func (_Brizo *BrizoSession) DataCounter(addr common.Address) (*big.Int, error) {
	return _Brizo.Contract.DataCounter(&_Brizo.CallOpts, addr)
}

// DataCounter is a free data retrieval call binding the contract method 0xd044a6be.
//
// Solidity: function dataCounter(address addr) constant returns(uint256)
func (_Brizo *BrizoCallerSession) DataCounter(addr common.Address) (*big.Int, error) {
	return _Brizo.Contract.DataCounter(&_Brizo.CallOpts, addr)
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

// WriteData is a paid mutator transaction binding the contract method 0x134bc876.
//
// Solidity: function writeData(string data) returns()
func (_Brizo *BrizoTransactor) WriteData(opts *bind.TransactOpts, data string) (*types.Transaction, error) {
	return _Brizo.contract.Transact(opts, "writeData", data)
}

// WriteData is a paid mutator transaction binding the contract method 0x134bc876.
//
// Solidity: function writeData(string data) returns()
func (_Brizo *BrizoSession) WriteData(data string) (*types.Transaction, error) {
	return _Brizo.Contract.WriteData(&_Brizo.TransactOpts, data)
}

// WriteData is a paid mutator transaction binding the contract method 0x134bc876.
//
// Solidity: function writeData(string data) returns()
func (_Brizo *BrizoTransactorSession) WriteData(data string) (*types.Transaction, error) {
	return _Brizo.Contract.WriteData(&_Brizo.TransactOpts, data)
}
