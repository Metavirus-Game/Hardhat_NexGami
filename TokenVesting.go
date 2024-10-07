// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// TokenVestingVestingSchedule is an auto generated low-level Go binding around an user-defined struct.
type TokenVestingVestingSchedule struct {
	Initialized        bool
	Beneficiary        common.Address
	Cliff              *big.Int
	Start              *big.Int
	Duration           *big.Int
	SlicePeriodSeconds *big.Int
	TotalAmount        *big.Int
	Released           *big.Int
	Revocable          bool
	Revoked            bool
	Role               uint8
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Revoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slicePeriodSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_revocable\",\"type\":\"bool\"},{\"internalType\":\"enumTokenVesting.Role\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"enumTokenVesting.ReleaseInterval\",\"name\":\"_releaseInterval\",\"type\":\"uint8\"}],\"name\":\"createVestingSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"getVestingSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slicePeriodSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"enumTokenVesting.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"internalType\":\"structTokenVesting.VestingSchedule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVestingSchedulesTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161122238038061122283398101604081905261002e91610128565b338061005457604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61005d816100d9565b506001600160a01b0381166100b45760405162461bcd60e51b815260206004820152601c60248201527f546f6b656e20616464726573732063616e6e6f74206265207a65726f00000000604482015260640161004b565b600180546001600160a01b0319166001600160a01b0392909216919091179055610155565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f60208284031215610138575f80fd5b81516001600160a01b038116811461014e575f80fd5b9392505050565b6110c0806101625f395ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c806374a8f1031161006357806374a8f103146100da5780638da5cb5b146100ed5780639f82906314610111578063f2fde38b14610131578063fc0c546a14610144575f80fd5b80630357371d1461009457806348deb471146100a95780636b7c60dd146100bf578063715018a6146100d2575b5f80fd5b6100a76100a2366004610e09565b610157565b005b6003546040519081526020015b60405180910390f35b6100a76100cd366004610e3e565b610411565b6100a761074d565b6100a76100e8366004610ecb565b610760565b5f546001600160a01b03165b6040516001600160a01b0390911681526020016100b6565b61012461011f366004610ecb565b610981565b6040516100b69190610f1f565b6100a761013f366004610ecb565b610a9b565b6001546100f9906001600160a01b031681565b6001600160a01b0382165f90815260026020526040902054829060ff166101995760405162461bcd60e51b815260040161019090610fc8565b60405180910390fd5b6001600160a01b0383165f908152600260205260409020600701548390610100900460ff161561020b5760405162461bcd60e51b815260206004820152601b60248201527f56657374696e67207363686564756c65206973207265766f6b656400000000006044820152606401610190565b6001600160a01b038085165f908152600260208181526040808420815161016081018352815460ff8082161515835261010091829004909816948201949094526001808301549382019390935293810154606085015260038101546080850152600481015460a0850152600581015460c0850152600681015460e0850152600781015480871615158486015292830486161515610120850152946102e59392869261014085019262010000909204909116908111156102cc576102cc610eeb565b60018111156102dd576102dd610eeb565b905250610ad8565b9050848110156103465760405162461bcd60e51b815260206004820152602660248201527f43616e6e6f742072656c65617365206d6f7265207468616e2076657374656420604482015265185b5bdd5b9d60d21b6064820152608401610190565b8482600601546103569190611013565b6006830155600154825460405163a9059cbb60e01b81526001600160a01b03610100909204821660048201526024810188905291169063a9059cbb906044016020604051808303815f875af11580156103b1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103d59190611026565b506040518581527ffb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c5659060200160405180910390a1505050505050565b610419610bac565b6001600160a01b03891661047a5760405162461bcd60e51b815260206004820152602260248201527f42656e65666963696172792063616e6e6f74206265207a65726f206164647265604482015261737360f01b6064820152608401610190565b5f86116104c05760405162461bcd60e51b815260206004820152601460248201527304475726174696f6e206d757374206265203e20360641b6044820152606401610190565b5f841161050f5760405162461bcd60e51b815260206004820152601860248201527f546f74616c20616d6f756e74206d757374206265203e203000000000000000006044820152606401610190565b600185101561056a5760405162461bcd60e51b815260206004820152602160248201527f536c69636520706572696f64207365636f6e6473206d757374206265203e3d206044820152603160f81b6064820152608401610190565b868610156105ba5760405162461bcd60e51b815260206004820152601960248201527f4475726174696f6e206d757374206265203e3d20636c696666000000000000006044820152606401610190565b5f6105c482610bd8565b90506105ce610d85565b600181526001600160a01b038b1660208201526105eb898b611013565b6040820152606081018a90526080810188905260a0810182905260c08101869052841515610100820152610140810184600181111561062c5761062c610eeb565b9081600181111561063f5761063f610eeb565b9052506001600160a01b038b81165f9081526002602081815260409283902085518154928701516001600160a81b0319909316901515610100600160a81b03191617610100929095168202949094178455918401516001808501919091556060850151918401919091556080840151600384015560a0840151600484015560c0840151600584015560e084015160068401558184015160078401805461012087015161ffff1990911692151561ff00191692909217911515909302178083556101408501518594939092909162ff00001916906201000090849081111561072857610728610eeb565b02179055505060035461073d91508790611013565b6003555050505050505050505050565b610755610bac565b61075e5f610d36565b565b610768610bac565b6001600160a01b0381165f90815260026020526040902054819060ff166107a15760405162461bcd60e51b815260040161019090610fc8565b6001600160a01b0382165f908152600260205260409020600701548290610100900460ff16156108135760405162461bcd60e51b815260206004820152601b60248201527f56657374696e67207363686564756c65206973207265766f6b656400000000006044820152606401610190565b6001600160a01b0383165f908152600260205260409020600781015460ff166108885760405162461bcd60e51b815260206004820152602160248201527f56657374696e67207363686564756c65206973206e6f74207265766f6361626c6044820152606560f81b6064820152608401610190565b60078101805461ff001916610100179055600681015460058201545f916108ae91611041565b9050806003546108be9190611041565b6003556001546001600160a01b031663a9059cbb6108e35f546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018490526044016020604051808303815f875af115801561092d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109519190611026565b506040517f44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6905f90a15050505050565b610989610d85565b6001600160a01b0382165f90815260026020526040902054829060ff166109c25760405162461bcd60e51b815260040161019090610fc8565b6001600160a01b038084165f90815260026020818152604092839020835161016081018552815460ff8082161515835261010091829004909716938201939093526001808301549582019590955292810154606084015260038101546080840152600481015460a0840152600581015460c0840152600681015460e0840152600781015480861615158385015291820485161515610120840152919391926101408501926201000090920490911690811115610a8057610a80610eeb565b6001811115610a9157610a91610eeb565b9052509392505050565b610aa3610bac565b6001600160a01b038116610acc57604051631e4fbdf760e01b81525f6004820152602401610190565b610ad581610d36565b50565b5f8160400151421015610aec57505f919050565b81610120015115610afe57505f919050565b81608001518260600151610b129190611013565b4210610b32578160e001518260c00151610b2c9190611041565b92915050565b5f826060015142610b439190611041565b90505f8360a0015182610b569190611054565b90505f8460a0015182610b699190611073565b90505f8560800151828760c00151610b819190611073565b610b8b9190611054565b90508560e0015181610b9d9190611041565b9695505050505050565b919050565b5f546001600160a01b0316331461075e5760405163118cdaa760e01b8152336004820152602401610190565b5f80826007811115610bec57610bec610eeb565b03610bf95750603c919050565b6001826007811115610c0d57610c0d610eeb565b03610c1b5750610e10919050565b6002826007811115610c2f57610c2f610eeb565b03610c3e575062015180919050565b6003826007811115610c5257610c52610eeb565b03610c61575062093a80919050565b6004826007811115610c7557610c75610eeb565b03610c84575062278d00919050565b6005826007811115610c9857610c98610eeb565b03610ca757506276a700919050565b6006826007811115610cbb57610cbb610eeb565b03610cca575062ed4e00919050565b6007826007811115610cde57610cde610eeb565b03610cee57506301e13380919050565b60405162461bcd60e51b815260206004820152601860248201527f496e76616c69642072656c6561736520696e74657276616c00000000000000006044820152606401610190565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040518061016001604052805f151581526020015f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f151581526020015f151581526020015f6001811115610dee57610dee610eeb565b905290565b80356001600160a01b0381168114610ba7575f80fd5b5f8060408385031215610e1a575f80fd5b610e2383610df3565b946020939093013593505050565b8015158114610ad5575f80fd5b5f805f805f805f805f6101208a8c031215610e57575f80fd5b610e608a610df3565b985060208a0135975060408a0135965060608a0135955060808a0135945060a08a0135935060c08a0135610e9381610e31565b925060e08a013560028110610ea6575f80fd5b91506101008a013560088110610eba575f80fd5b809150509295985092959850929598565b5f60208284031215610edb575f80fd5b610ee482610df3565b9392505050565b634e487b7160e01b5f52602160045260245ffd5b60028110610f1b57634e487b7160e01b5f52602160045260245ffd5b9052565b81511515815261016081016020830151610f4460208401826001600160a01b03169052565b5060408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e083015160e0830152610100830151610f9661010084018215159052565b50610120830151610fac61012084018215159052565b50610140830151610fc1610140840182610eff565b5092915050565b6020808252601f908201527f56657374696e67207363686564756c6520646f6573206e6f7420657869737400604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610b2c57610b2c610fff565b5f60208284031215611036575f80fd5b8151610ee481610e31565b81810381811115610b2c57610b2c610fff565b5f8261106e57634e487b7160e01b5f52601260045260245ffd5b500490565b8082028115828204841417610b2c57610b2c610fff56fea264697066735822122094a07564d77010e267dd28e565d4981a50773d87e06696fe7290d8d7c2c7438464736f6c634300081a0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, token_)
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
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9f829063.
//
// Solidity: function getVestingSchedule(address _beneficiary) view returns((bool,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint8))
func (_Contract *ContractCaller) GetVestingSchedule(opts *bind.CallOpts, _beneficiary common.Address) (TokenVestingVestingSchedule, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getVestingSchedule", _beneficiary)

	if err != nil {
		return *new(TokenVestingVestingSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenVestingVestingSchedule)).(*TokenVestingVestingSchedule)

	return out0, err

}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9f829063.
//
// Solidity: function getVestingSchedule(address _beneficiary) view returns((bool,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint8))
func (_Contract *ContractSession) GetVestingSchedule(_beneficiary common.Address) (TokenVestingVestingSchedule, error) {
	return _Contract.Contract.GetVestingSchedule(&_Contract.CallOpts, _beneficiary)
}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9f829063.
//
// Solidity: function getVestingSchedule(address _beneficiary) view returns((bool,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint8))
func (_Contract *ContractCallerSession) GetVestingSchedule(_beneficiary common.Address) (TokenVestingVestingSchedule, error) {
	return _Contract.Contract.GetVestingSchedule(&_Contract.CallOpts, _beneficiary)
}

// GetVestingSchedulesTotalAmount is a free data retrieval call binding the contract method 0x48deb471.
//
// Solidity: function getVestingSchedulesTotalAmount() view returns(uint256)
func (_Contract *ContractCaller) GetVestingSchedulesTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getVestingSchedulesTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestingSchedulesTotalAmount is a free data retrieval call binding the contract method 0x48deb471.
//
// Solidity: function getVestingSchedulesTotalAmount() view returns(uint256)
func (_Contract *ContractSession) GetVestingSchedulesTotalAmount() (*big.Int, error) {
	return _Contract.Contract.GetVestingSchedulesTotalAmount(&_Contract.CallOpts)
}

// GetVestingSchedulesTotalAmount is a free data retrieval call binding the contract method 0x48deb471.
//
// Solidity: function getVestingSchedulesTotalAmount() view returns(uint256)
func (_Contract *ContractCallerSession) GetVestingSchedulesTotalAmount() (*big.Int, error) {
	return _Contract.Contract.GetVestingSchedulesTotalAmount(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contract *ContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contract *ContractSession) Token() (common.Address, error) {
	return _Contract.Contract.Token(&_Contract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contract *ContractCallerSession) Token() (common.Address, error) {
	return _Contract.Contract.Token(&_Contract.CallOpts)
}

// CreateVestingSchedule is a paid mutator transaction binding the contract method 0x6b7c60dd.
//
// Solidity: function createVestingSchedule(address _beneficiary, uint256 _start, uint256 _cliff, uint256 _duration, uint256 _slicePeriodSeconds, uint256 _totalAmount, bool _revocable, uint8 _role, uint8 _releaseInterval) returns()
func (_Contract *ContractTransactor) CreateVestingSchedule(opts *bind.TransactOpts, _beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _slicePeriodSeconds *big.Int, _totalAmount *big.Int, _revocable bool, _role uint8, _releaseInterval uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createVestingSchedule", _beneficiary, _start, _cliff, _duration, _slicePeriodSeconds, _totalAmount, _revocable, _role, _releaseInterval)
}

// CreateVestingSchedule is a paid mutator transaction binding the contract method 0x6b7c60dd.
//
// Solidity: function createVestingSchedule(address _beneficiary, uint256 _start, uint256 _cliff, uint256 _duration, uint256 _slicePeriodSeconds, uint256 _totalAmount, bool _revocable, uint8 _role, uint8 _releaseInterval) returns()
func (_Contract *ContractSession) CreateVestingSchedule(_beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _slicePeriodSeconds *big.Int, _totalAmount *big.Int, _revocable bool, _role uint8, _releaseInterval uint8) (*types.Transaction, error) {
	return _Contract.Contract.CreateVestingSchedule(&_Contract.TransactOpts, _beneficiary, _start, _cliff, _duration, _slicePeriodSeconds, _totalAmount, _revocable, _role, _releaseInterval)
}

// CreateVestingSchedule is a paid mutator transaction binding the contract method 0x6b7c60dd.
//
// Solidity: function createVestingSchedule(address _beneficiary, uint256 _start, uint256 _cliff, uint256 _duration, uint256 _slicePeriodSeconds, uint256 _totalAmount, bool _revocable, uint8 _role, uint8 _releaseInterval) returns()
func (_Contract *ContractTransactorSession) CreateVestingSchedule(_beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _slicePeriodSeconds *big.Int, _totalAmount *big.Int, _revocable bool, _role uint8, _releaseInterval uint8) (*types.Transaction, error) {
	return _Contract.Contract.CreateVestingSchedule(&_Contract.TransactOpts, _beneficiary, _start, _cliff, _duration, _slicePeriodSeconds, _totalAmount, _revocable, _role, _releaseInterval)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractTransactor) Release(opts *bind.TransactOpts, _beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "release", _beneficiary, _amount)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractSession) Release(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Release(&_Contract.TransactOpts, _beneficiary, _amount)
}

// Release is a paid mutator transaction binding the contract method 0x0357371d.
//
// Solidity: function release(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) Release(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Release(&_Contract.TransactOpts, _beneficiary, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _beneficiary) returns()
func (_Contract *ContractTransactor) Revoke(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revoke", _beneficiary)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _beneficiary) returns()
func (_Contract *ContractSession) Revoke(_beneficiary common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Revoke(&_Contract.TransactOpts, _beneficiary)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _beneficiary) returns()
func (_Contract *ContractTransactorSession) Revoke(_beneficiary common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Revoke(&_Contract.TransactOpts, _beneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the Contract contract.
type ContractReleasedIterator struct {
	Event *ContractReleased // Event containing the contract specifics and raw log

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
func (it *ContractReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractReleased)
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
		it.Event = new(ContractReleased)
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
func (it *ContractReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractReleased represents a Released event raised by the Contract contract.
type ContractReleased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0xfb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c565.
//
// Solidity: event Released(uint256 amount)
func (_Contract *ContractFilterer) FilterReleased(opts *bind.FilterOpts) (*ContractReleasedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Released")
	if err != nil {
		return nil, err
	}
	return &ContractReleasedIterator{contract: _Contract.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0xfb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c565.
//
// Solidity: event Released(uint256 amount)
func (_Contract *ContractFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *ContractReleased) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Released")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractReleased)
				if err := _Contract.contract.UnpackLog(event, "Released", log); err != nil {
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

// ParseReleased is a log parse operation binding the contract event 0xfb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c565.
//
// Solidity: event Released(uint256 amount)
func (_Contract *ContractFilterer) ParseReleased(log types.Log) (*ContractReleased, error) {
	event := new(ContractReleased)
	if err := _Contract.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the Contract contract.
type ContractRevokedIterator struct {
	Event *ContractRevoked // Event containing the contract specifics and raw log

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
func (it *ContractRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRevoked)
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
		it.Event = new(ContractRevoked)
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
func (it *ContractRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRevoked represents a Revoked event raised by the Contract contract.
type ContractRevoked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_Contract *ContractFilterer) FilterRevoked(opts *bind.FilterOpts) (*ContractRevokedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return &ContractRevokedIterator{contract: _Contract.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_Contract *ContractFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *ContractRevoked) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRevoked)
				if err := _Contract.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: event Revoked()
func (_Contract *ContractFilterer) ParseRevoked(log types.Log) (*ContractRevoked, error) {
	event := new(ContractRevoked)
	if err := _Contract.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
