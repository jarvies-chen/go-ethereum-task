// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
	_ = abi.ConvertType
)

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CountIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CountReset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"setCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f5f819055506104f5806100625f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c80638da5cb5b14610059578063a87d942c14610077578063d09de08a14610095578063d14e62b81461009f578063d826f88f146100bb575b5f5ffd5b6100616100c5565b60405161006e919061030a565b60405180910390f35b61007f6100ea565b60405161008c919061033b565b60405180910390f35b61009d6100f2565b005b6100b960048036038101906100b49190610382565b610159565b005b6100c36101f1565b005b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f54905090565b5f5f815480929190610103906103da565b91905055503373ffffffffffffffffffffffffffffffffffffffff167f66e68b0f81b27905187c1253498474c586eef78d11987ee7ea04427459d2daa95f5460405161014f919061033b565b60405180910390a2565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101df906104a1565b60405180910390fd5b805f8190555050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610280576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610277906104a1565b60405180910390fd5b5f5f819055503373ffffffffffffffffffffffffffffffffffffffff167fa5ee6258204973c56c5a39c4ac31e61723f410d84f9e8117ba52b76b7cea990c60405160405180910390a2565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102f4826102cb565b9050919050565b610304816102ea565b82525050565b5f60208201905061031d5f8301846102fb565b92915050565b5f819050919050565b61033581610323565b82525050565b5f60208201905061034e5f83018461032c565b92915050565b5f5ffd5b61036181610323565b811461036b575f5ffd5b50565b5f8135905061037c81610358565b92915050565b5f6020828403121561039757610396610354565b5b5f6103a48482850161036e565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6103e482610323565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610416576104156103ad565b5b600182019050919050565b5f82825260208201905092915050565b7f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f5f8201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f61048b602183610421565b915061049682610431565b604082019050919050565b5f6020820190508181035f8301526104b88161047f565b905091905056fea264697066735822122045a9d164bc2261e025af335573a3242d6fde4d9c1184a0a1cd21fd71030f57a664736f6c634300081e0033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterCallerSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Counter *CounterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Counter *CounterSession) Owner() (common.Address, error) {
	return _Counter.Contract.Owner(&_Counter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Counter *CounterCallerSession) Owner() (common.Address, error) {
	return _Counter.Contract.Owner(&_Counter.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterTransactorSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_Counter *CounterTransactorSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _count) returns()
func (_Counter *CounterTransactor) SetCount(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "setCount", _count)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _count) returns()
func (_Counter *CounterSession) SetCount(_count *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.SetCount(&_Counter.TransactOpts, _count)
}

// SetCount is a paid mutator transaction binding the contract method 0xd14e62b8.
//
// Solidity: function setCount(uint256 _count) returns()
func (_Counter *CounterTransactorSession) SetCount(_count *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.SetCount(&_Counter.TransactOpts, _count)
}

// CounterCountIncrementedIterator is returned from FilterCountIncremented and is used to iterate over the raw logs and unpacked data for CountIncremented events raised by the Counter contract.
type CounterCountIncrementedIterator struct {
	Event *CounterCountIncremented // Event containing the contract specifics and raw log

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
func (it *CounterCountIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterCountIncremented)
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
		it.Event = new(CounterCountIncremented)
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
func (it *CounterCountIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterCountIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterCountIncremented represents a CountIncremented event raised by the Counter contract.
type CounterCountIncremented struct {
	Caller   common.Address
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountIncremented is a free log retrieval operation binding the contract event 0x66e68b0f81b27905187c1253498474c586eef78d11987ee7ea04427459d2daa9.
//
// Solidity: event CountIncremented(address indexed caller, uint256 newValue)
func (_Counter *CounterFilterer) FilterCountIncremented(opts *bind.FilterOpts, caller []common.Address) (*CounterCountIncrementedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Counter.contract.FilterLogs(opts, "CountIncremented", callerRule)
	if err != nil {
		return nil, err
	}
	return &CounterCountIncrementedIterator{contract: _Counter.contract, event: "CountIncremented", logs: logs, sub: sub}, nil
}

// WatchCountIncremented is a free log subscription operation binding the contract event 0x66e68b0f81b27905187c1253498474c586eef78d11987ee7ea04427459d2daa9.
//
// Solidity: event CountIncremented(address indexed caller, uint256 newValue)
func (_Counter *CounterFilterer) WatchCountIncremented(opts *bind.WatchOpts, sink chan<- *CounterCountIncremented, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Counter.contract.WatchLogs(opts, "CountIncremented", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterCountIncremented)
				if err := _Counter.contract.UnpackLog(event, "CountIncremented", log); err != nil {
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

// ParseCountIncremented is a log parse operation binding the contract event 0x66e68b0f81b27905187c1253498474c586eef78d11987ee7ea04427459d2daa9.
//
// Solidity: event CountIncremented(address indexed caller, uint256 newValue)
func (_Counter *CounterFilterer) ParseCountIncremented(log types.Log) (*CounterCountIncremented, error) {
	event := new(CounterCountIncremented)
	if err := _Counter.contract.UnpackLog(event, "CountIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterCountResetIterator is returned from FilterCountReset and is used to iterate over the raw logs and unpacked data for CountReset events raised by the Counter contract.
type CounterCountResetIterator struct {
	Event *CounterCountReset // Event containing the contract specifics and raw log

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
func (it *CounterCountResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterCountReset)
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
		it.Event = new(CounterCountReset)
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
func (it *CounterCountResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterCountResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterCountReset represents a CountReset event raised by the Counter contract.
type CounterCountReset struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCountReset is a free log retrieval operation binding the contract event 0xa5ee6258204973c56c5a39c4ac31e61723f410d84f9e8117ba52b76b7cea990c.
//
// Solidity: event CountReset(address indexed caller)
func (_Counter *CounterFilterer) FilterCountReset(opts *bind.FilterOpts, caller []common.Address) (*CounterCountResetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Counter.contract.FilterLogs(opts, "CountReset", callerRule)
	if err != nil {
		return nil, err
	}
	return &CounterCountResetIterator{contract: _Counter.contract, event: "CountReset", logs: logs, sub: sub}, nil
}

// WatchCountReset is a free log subscription operation binding the contract event 0xa5ee6258204973c56c5a39c4ac31e61723f410d84f9e8117ba52b76b7cea990c.
//
// Solidity: event CountReset(address indexed caller)
func (_Counter *CounterFilterer) WatchCountReset(opts *bind.WatchOpts, sink chan<- *CounterCountReset, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Counter.contract.WatchLogs(opts, "CountReset", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterCountReset)
				if err := _Counter.contract.UnpackLog(event, "CountReset", log); err != nil {
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

// ParseCountReset is a log parse operation binding the contract event 0xa5ee6258204973c56c5a39c4ac31e61723f410d84f9e8117ba52b76b7cea990c.
//
// Solidity: event CountReset(address indexed caller)
func (_Counter *CounterFilterer) ParseCountReset(log types.Log) (*CounterCountReset, error) {
	event := new(CounterCountReset)
	if err := _Counter.contract.UnpackLog(event, "CountReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
