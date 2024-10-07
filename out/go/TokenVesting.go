// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenvesting

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

// TokenvestingMetaData contains all meta data concerning the Tokenvesting contract.
var TokenvestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTokenVesting.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"}],\"name\":\"VestingScheduleCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_revocable\",\"type\":\"bool\"},{\"internalType\":\"enumTokenVesting.Role\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"enumTokenVesting.ReleaseInterval\",\"name\":\"_releaseInterval\",\"type\":\"uint8\"}],\"name\":\"createVestingSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"getVestingSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slicePeriodSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"enumTokenVesting.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"internalType\":\"structTokenVesting.VestingSchedule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVestingSchedulesTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161147338038061147383398101604081905261002e91610128565b338061005457604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61005d816100d9565b506001600160a01b0381166100b45760405162461bcd60e51b815260206004820152601c60248201527f546f6b656e20616464726573732063616e6e6f74206265207a65726f00000000604482015260640161004b565b600180546001600160a01b0319166001600160a01b0392909216919091179055610155565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f60208284031215610138575f80fd5b81516001600160a01b038116811461014e575f80fd5b9392505050565b611311806101625f395ff3fe608060405234801561000f575f80fd5b50600436106100a6575f3560e01c8063715018a61161006e578063715018a61461011557806374a8f1031461011d5780638da5cb5b146101305780639f82906314610140578063f2fde38b14610160578063fc0c546a14610173575f80fd5b806310fe9ae8146100aa57806319165587146100d45780633ccfd60b146100e957806348deb471146100f15780634979cd1914610102575b5f80fd5b6001546001600160a01b03165b6040516001600160a01b0390911681526020015b60405180910390f35b6100e76100e2366004611024565b610186565b005b6100e761040f565b6003546040519081526020016100cb565b6100e7610110366004611051565b610661565b6100e761098c565b6100e761012b366004611024565b61099f565b5f546001600160a01b03166100b7565b61015361014e366004611024565b610b9c565b6040516100cb9190611109565b6100e761016e366004611024565b610cb6565b6001546100b7906001600160a01b031681565b6001600160a01b0381165f90815260026020526040902054819060ff166101c85760405162461bcd60e51b81526004016101bf906111b2565b60405180910390fd5b6001600160a01b0382165f908152600260205260409020600701548290610100900460ff161561020a5760405162461bcd60e51b81526004016101bf906111e9565b6001600160a01b038084165f908152600260208181526040808420815161016081018352815460ff8082161515835261010091829004909816948201949094526001808301549382019390935293810154606085015260038101546080850152600481015460a0850152600581015460c0850152600681015460e0850152600781015480871615158486015292830486161515610120850152946102e49392869261014085019262010000909204909116908111156102cb576102cb6110d5565b60018111156102dc576102dc6110d5565b905250610cf3565b90505f81116103355760405162461bcd60e51b815260206004820152601e60248201527f4e6f2072656c65617361626c6520616d6f756e7420617661696c61626c65000060448201526064016101bf565b8082600601546103459190611234565b6006830155600154825460405163a9059cbb60e01b81526001600160a01b03610100909204821660048201526024810184905291169063a9059cbb906044016020604051808303815f875af11580156103a0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103c49190611247565b50846001600160a01b03167fb21fb52d5749b80f3182f8c6992236b5e5576681880914484d7f4c9b062e619e8260405161040091815260200190565b60405180910390a25050505050565b335f8181526002602052604090205460ff1661043d5760405162461bcd60e51b81526004016101bf906111b2565b6001600160a01b0381165f90815260026020526040902060070154610100900460ff161561047d5760405162461bcd60e51b81526004016101bf906111e9565b6001600160a01b038082165f908152600260208181526040808420815161016081018352815460ff8082161515835261010091829004909816948201949094526001808301549382019390935293810154606085015260038101546080850152600481015460a0850152600581015460c0850152600681015460e08501526007810154808716151584860152928304861615156101208501529461053e9392869261014085019262010000909204909116908111156102cb576102cb6110d5565b90505f811161058f5760405162461bcd60e51b815260206004820152601e60248201527f4e6f2072656c65617361626c6520616d6f756e7420617661696c61626c65000060448201526064016101bf565b80826006015461059f9190611234565b600683015560015460405163a9059cbb60e01b81526001600160a01b038581166004830152602482018490529091169063a9059cbb906044016020604051808303815f875af11580156105f4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106189190611247565b50826001600160a01b03167fb21fb52d5749b80f3182f8c6992236b5e5576681880914484d7f4c9b062e619e8260405161065491815260200190565b60405180910390a2505050565b610669610dc7565b6001600160a01b0388166106ca5760405162461bcd60e51b815260206004820152602260248201527f42656e65666963696172792063616e6e6f74206265207a65726f206164647265604482015261737360f01b60648201526084016101bf565b5f85116107105760405162461bcd60e51b815260206004820152601460248201527304475726174696f6e206d757374206265203e20360641b60448201526064016101bf565b5f841161075f5760405162461bcd60e51b815260206004820152601860248201527f546f74616c20616d6f756e74206d757374206265203e2030000000000000000060448201526064016101bf565b858510156107af5760405162461bcd60e51b815260206004820152601960248201527f4475726174696f6e206d757374206265203e3d20636c6966660000000000000060448201526064016101bf565b5f6107b982610df3565b90506107c3610fa0565b600181526001600160a01b038a1660208201526107e0888a611234565b6040820152606081018990526080810187905260a0810182905260c081018690528415156101008201526101408101846001811115610821576108216110d5565b90816001811115610834576108346110d5565b9052506001600160a01b038a81165f9081526002602081815260409283902085518154928701516001600160a81b0319909316901515610100600160a81b03191617610100929095168202949094178455918401516001808501919091556060850151918401919091556080840151600384015560a0840151600484015560c0840151600584015560e084015160068401558184015160078401805461012087015161ffff1990911692151561ff00191692909217911515909302178083556101408501518594939092909162ff00001916906201000090849081111561091d5761091d6110d5565b02179055505060035461093291508790611234565b6003556040516001600160a01b038b16907f1ea8125ddaa7ea3322d4ec14f10ea84ba3a89c4306f9bf1610f4d0036ac9f90c9061097890899088908e908d908f90611262565b60405180910390a250505050505050505050565b610994610dc7565b61099d5f610f51565b565b6109a7610dc7565b6001600160a01b0381165f90815260026020526040902054819060ff166109e05760405162461bcd60e51b81526004016101bf906111b2565b6001600160a01b0382165f908152600260205260409020600701548290610100900460ff1615610a225760405162461bcd60e51b81526004016101bf906111e9565b6001600160a01b0383165f908152600260205260409020600781015460ff16610a975760405162461bcd60e51b815260206004820152602160248201527f56657374696e67207363686564756c65206973206e6f74207265766f6361626c6044820152606560f81b60648201526084016101bf565b60078101805461ff001916610100179055600681015460058201545f91610abd91611292565b905080600354610acd9190611292565b6003556001546001600160a01b031663a9059cbb610af25f546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018490526044016020604051808303815f875af1158015610b3c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b609190611247565b50846001600160a01b03167f713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac48260405161040091815260200190565b610ba4610fa0565b6001600160a01b0382165f90815260026020526040902054829060ff16610bdd5760405162461bcd60e51b81526004016101bf906111b2565b6001600160a01b038084165f90815260026020818152604092839020835161016081018552815460ff8082161515835261010091829004909716938201939093526001808301549582019590955292810154606084015260038101546080840152600481015460a0840152600581015460c0840152600681015460e0840152600781015480861615158385015291820485161515610120840152919391926101408501926201000090920490911690811115610c9b57610c9b6110d5565b6001811115610cac57610cac6110d5565b9052509392505050565b610cbe610dc7565b6001600160a01b038116610ce757604051631e4fbdf760e01b81525f60048201526024016101bf565b610cf081610f51565b50565b5f8160400151421015610d0757505f919050565b81610120015115610d1957505f919050565b81608001518260600151610d2d9190611234565b4210610d4d578160e001518260c00151610d479190611292565b92915050565b5f826060015142610d5e9190611292565b90505f8360a0015182610d7191906112a5565b90505f8460a0015182610d8491906112c4565b90505f8560800151828760c00151610d9c91906112c4565b610da691906112a5565b90508560e0015181610db89190611292565b9695505050505050565b919050565b5f546001600160a01b0316331461099d5760405163118cdaa760e01b81523360048201526024016101bf565b5f80826007811115610e0757610e076110d5565b03610e145750603c919050565b6001826007811115610e2857610e286110d5565b03610e365750610e10919050565b6002826007811115610e4a57610e4a6110d5565b03610e59575062015180919050565b6003826007811115610e6d57610e6d6110d5565b03610e7c575062093a80919050565b6004826007811115610e9057610e906110d5565b03610e9f575062278d00919050565b6005826007811115610eb357610eb36110d5565b03610ec257506276a700919050565b6006826007811115610ed657610ed66110d5565b03610ee5575062ed4e00919050565b6007826007811115610ef957610ef96110d5565b03610f0957506301e13380919050565b60405162461bcd60e51b815260206004820152601860248201527f496e76616c69642072656c6561736520696e74657276616c000000000000000060448201526064016101bf565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040518061016001604052805f151581526020015f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f151581526020015f151581526020015f6001811115611009576110096110d5565b905290565b80356001600160a01b0381168114610dc2575f80fd5b5f60208284031215611034575f80fd5b61103d8261100e565b9392505050565b8015158114610cf0575f80fd5b5f805f805f805f80610100898b031215611069575f80fd5b6110728961100e565b97506020890135965060408901359550606089013594506080890135935060a089013561109e81611044565b925060c0890135600281106110b1575f80fd5b915060e0890135600881106110c4575f80fd5b809150509295985092959890939650565b634e487b7160e01b5f52602160045260245ffd5b6002811061110557634e487b7160e01b5f52602160045260245ffd5b9052565b8151151581526101608101602083015161112e60208401826001600160a01b03169052565b5060408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e083015160e083015261010083015161118061010084018215159052565b5061012083015161119661012084018215159052565b506101408301516111ab6101408401826110e9565b5092915050565b6020808252601f908201527f56657374696e67207363686564756c6520646f6573206e6f7420657869737400604082015260600190565b6020808252601b908201527f56657374696e67207363686564756c65206973207265766f6b65640000000000604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d4757610d47611220565b5f60208284031215611257575f80fd5b815161103d81611044565b85815260a0810161127660208301876110e9565b8460408301528360608301528260808301529695505050505050565b81810381811115610d4757610d47611220565b5f826112bf57634e487b7160e01b5f52601260045260245ffd5b500490565b8082028115828204841417610d4757610d4761122056fea26469706673582212202efe011ff3f4f7bbb880ee96b548d8c8d1e5445a6d233d48c5ad04c66ce520e964736f6c634300081a0033",
}

// TokenvestingABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenvestingMetaData.ABI instead.
var TokenvestingABI = TokenvestingMetaData.ABI

// TokenvestingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenvestingMetaData.Bin instead.
var TokenvestingBin = TokenvestingMetaData.Bin

// DeployTokenvesting deploys a new Ethereum contract, binding an instance of Tokenvesting to it.
func DeployTokenvesting(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address) (common.Address, *types.Transaction, *Tokenvesting, error) {
	parsed, err := TokenvestingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenvestingBin), backend, token_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokenvesting{TokenvestingCaller: TokenvestingCaller{contract: contract}, TokenvestingTransactor: TokenvestingTransactor{contract: contract}, TokenvestingFilterer: TokenvestingFilterer{contract: contract}}, nil
}

// Tokenvesting is an auto generated Go binding around an Ethereum contract.
type Tokenvesting struct {
	TokenvestingCaller     // Read-only binding to the contract
	TokenvestingTransactor // Write-only binding to the contract
	TokenvestingFilterer   // Log filterer for contract events
}

// TokenvestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenvestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenvestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenvestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenvestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenvestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenvestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenvestingSession struct {
	Contract     *Tokenvesting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenvestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenvestingCallerSession struct {
	Contract *TokenvestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenvestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenvestingTransactorSession struct {
	Contract     *TokenvestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenvestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenvestingRaw struct {
	Contract *Tokenvesting // Generic contract binding to access the raw methods on
}

// TokenvestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenvestingCallerRaw struct {
	Contract *TokenvestingCaller // Generic read-only contract binding to access the raw methods on
}

// TokenvestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenvestingTransactorRaw struct {
	Contract *TokenvestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenvesting creates a new instance of Tokenvesting, bound to a specific deployed contract.
func NewTokenvesting(address common.Address, backend bind.ContractBackend) (*Tokenvesting, error) {
	contract, err := bindTokenvesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenvesting{TokenvestingCaller: TokenvestingCaller{contract: contract}, TokenvestingTransactor: TokenvestingTransactor{contract: contract}, TokenvestingFilterer: TokenvestingFilterer{contract: contract}}, nil
}

// NewTokenvestingCaller creates a new read-only instance of Tokenvesting, bound to a specific deployed contract.
func NewTokenvestingCaller(address common.Address, caller bind.ContractCaller) (*TokenvestingCaller, error) {
	contract, err := bindTokenvesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenvestingCaller{contract: contract}, nil
}

// NewTokenvestingTransactor creates a new write-only instance of Tokenvesting, bound to a specific deployed contract.
func NewTokenvestingTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenvestingTransactor, error) {
	contract, err := bindTokenvesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenvestingTransactor{contract: contract}, nil
}

// NewTokenvestingFilterer creates a new log filterer instance of Tokenvesting, bound to a specific deployed contract.
func NewTokenvestingFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenvestingFilterer, error) {
	contract, err := bindTokenvesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenvestingFilterer{contract: contract}, nil
}

// bindTokenvesting binds a generic wrapper to an already deployed contract.
func bindTokenvesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenvestingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenvesting *TokenvestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenvesting.Contract.TokenvestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenvesting *TokenvestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenvesting.Contract.TokenvestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenvesting *TokenvestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenvesting.Contract.TokenvestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenvesting *TokenvestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenvesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenvesting *TokenvestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenvesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenvesting *TokenvestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenvesting.Contract.contract.Transact(opts, method, params...)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Tokenvesting *TokenvestingCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenvesting.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Tokenvesting *TokenvestingSession) GetTokenAddress() (common.Address, error) {
	return _Tokenvesting.Contract.GetTokenAddress(&_Tokenvesting.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Tokenvesting *TokenvestingCallerSession) GetTokenAddress() (common.Address, error) {
	return _Tokenvesting.Contract.GetTokenAddress(&_Tokenvesting.CallOpts)
}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9f829063.
//
// Solidity: function getVestingSchedule(address _beneficiary) view returns((bool,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint8))
func (_Tokenvesting *TokenvestingCaller) GetVestingSchedule(opts *bind.CallOpts, _beneficiary common.Address) (TokenVestingVestingSchedule, error) {
	var out []interface{}
	err := _Tokenvesting.contract.Call(opts, &out, "getVestingSchedule", _beneficiary)

	if err != nil {
		return *new(TokenVestingVestingSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenVestingVestingSchedule)).(*TokenVestingVestingSchedule)

	return out0, err

}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9f829063.
//
// Solidity: function getVestingSchedule(address _beneficiary) view returns((bool,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint8))
func (_Tokenvesting *TokenvestingSession) GetVestingSchedule(_beneficiary common.Address) (TokenVestingVestingSchedule, error) {
	return _Tokenvesting.Contract.GetVestingSchedule(&_Tokenvesting.CallOpts, _beneficiary)
}

// GetVestingSchedule is a free data retrieval call binding the contract method 0x9f829063.
//
// Solidity: function getVestingSchedule(address _beneficiary) view returns((bool,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint8))
func (_Tokenvesting *TokenvestingCallerSession) GetVestingSchedule(_beneficiary common.Address) (TokenVestingVestingSchedule, error) {
	return _Tokenvesting.Contract.GetVestingSchedule(&_Tokenvesting.CallOpts, _beneficiary)
}

// GetVestingSchedulesTotalAmount is a free data retrieval call binding the contract method 0x48deb471.
//
// Solidity: function getVestingSchedulesTotalAmount() view returns(uint256)
func (_Tokenvesting *TokenvestingCaller) GetVestingSchedulesTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenvesting.contract.Call(opts, &out, "getVestingSchedulesTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestingSchedulesTotalAmount is a free data retrieval call binding the contract method 0x48deb471.
//
// Solidity: function getVestingSchedulesTotalAmount() view returns(uint256)
func (_Tokenvesting *TokenvestingSession) GetVestingSchedulesTotalAmount() (*big.Int, error) {
	return _Tokenvesting.Contract.GetVestingSchedulesTotalAmount(&_Tokenvesting.CallOpts)
}

// GetVestingSchedulesTotalAmount is a free data retrieval call binding the contract method 0x48deb471.
//
// Solidity: function getVestingSchedulesTotalAmount() view returns(uint256)
func (_Tokenvesting *TokenvestingCallerSession) GetVestingSchedulesTotalAmount() (*big.Int, error) {
	return _Tokenvesting.Contract.GetVestingSchedulesTotalAmount(&_Tokenvesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenvesting *TokenvestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenvesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenvesting *TokenvestingSession) Owner() (common.Address, error) {
	return _Tokenvesting.Contract.Owner(&_Tokenvesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenvesting *TokenvestingCallerSession) Owner() (common.Address, error) {
	return _Tokenvesting.Contract.Owner(&_Tokenvesting.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Tokenvesting *TokenvestingCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenvesting.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Tokenvesting *TokenvestingSession) Token() (common.Address, error) {
	return _Tokenvesting.Contract.Token(&_Tokenvesting.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Tokenvesting *TokenvestingCallerSession) Token() (common.Address, error) {
	return _Tokenvesting.Contract.Token(&_Tokenvesting.CallOpts)
}

// CreateVestingSchedule is a paid mutator transaction binding the contract method 0x4979cd19.
//
// Solidity: function createVestingSchedule(address _beneficiary, uint256 _start, uint256 _cliff, uint256 _duration, uint256 _totalAmount, bool _revocable, uint8 _role, uint8 _releaseInterval) returns()
func (_Tokenvesting *TokenvestingTransactor) CreateVestingSchedule(opts *bind.TransactOpts, _beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _totalAmount *big.Int, _revocable bool, _role uint8, _releaseInterval uint8) (*types.Transaction, error) {
	return _Tokenvesting.contract.Transact(opts, "createVestingSchedule", _beneficiary, _start, _cliff, _duration, _totalAmount, _revocable, _role, _releaseInterval)
}

// CreateVestingSchedule is a paid mutator transaction binding the contract method 0x4979cd19.
//
// Solidity: function createVestingSchedule(address _beneficiary, uint256 _start, uint256 _cliff, uint256 _duration, uint256 _totalAmount, bool _revocable, uint8 _role, uint8 _releaseInterval) returns()
func (_Tokenvesting *TokenvestingSession) CreateVestingSchedule(_beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _totalAmount *big.Int, _revocable bool, _role uint8, _releaseInterval uint8) (*types.Transaction, error) {
	return _Tokenvesting.Contract.CreateVestingSchedule(&_Tokenvesting.TransactOpts, _beneficiary, _start, _cliff, _duration, _totalAmount, _revocable, _role, _releaseInterval)
}

// CreateVestingSchedule is a paid mutator transaction binding the contract method 0x4979cd19.
//
// Solidity: function createVestingSchedule(address _beneficiary, uint256 _start, uint256 _cliff, uint256 _duration, uint256 _totalAmount, bool _revocable, uint8 _role, uint8 _releaseInterval) returns()
func (_Tokenvesting *TokenvestingTransactorSession) CreateVestingSchedule(_beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _totalAmount *big.Int, _revocable bool, _role uint8, _releaseInterval uint8) (*types.Transaction, error) {
	return _Tokenvesting.Contract.CreateVestingSchedule(&_Tokenvesting.TransactOpts, _beneficiary, _start, _cliff, _duration, _totalAmount, _revocable, _role, _releaseInterval)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address _beneficiary) returns()
func (_Tokenvesting *TokenvestingTransactor) Release(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _Tokenvesting.contract.Transact(opts, "release", _beneficiary)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address _beneficiary) returns()
func (_Tokenvesting *TokenvestingSession) Release(_beneficiary common.Address) (*types.Transaction, error) {
	return _Tokenvesting.Contract.Release(&_Tokenvesting.TransactOpts, _beneficiary)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address _beneficiary) returns()
func (_Tokenvesting *TokenvestingTransactorSession) Release(_beneficiary common.Address) (*types.Transaction, error) {
	return _Tokenvesting.Contract.Release(&_Tokenvesting.TransactOpts, _beneficiary)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenvesting *TokenvestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenvesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenvesting *TokenvestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenvesting.Contract.RenounceOwnership(&_Tokenvesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenvesting *TokenvestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenvesting.Contract.RenounceOwnership(&_Tokenvesting.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _beneficiary) returns()
func (_Tokenvesting *TokenvestingTransactor) Revoke(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _Tokenvesting.contract.Transact(opts, "revoke", _beneficiary)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _beneficiary) returns()
func (_Tokenvesting *TokenvestingSession) Revoke(_beneficiary common.Address) (*types.Transaction, error) {
	return _Tokenvesting.Contract.Revoke(&_Tokenvesting.TransactOpts, _beneficiary)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _beneficiary) returns()
func (_Tokenvesting *TokenvestingTransactorSession) Revoke(_beneficiary common.Address) (*types.Transaction, error) {
	return _Tokenvesting.Contract.Revoke(&_Tokenvesting.TransactOpts, _beneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenvesting *TokenvestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokenvesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenvesting *TokenvestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenvesting.Contract.TransferOwnership(&_Tokenvesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenvesting *TokenvestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenvesting.Contract.TransferOwnership(&_Tokenvesting.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Tokenvesting *TokenvestingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenvesting.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Tokenvesting *TokenvestingSession) Withdraw() (*types.Transaction, error) {
	return _Tokenvesting.Contract.Withdraw(&_Tokenvesting.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Tokenvesting *TokenvestingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Tokenvesting.Contract.Withdraw(&_Tokenvesting.TransactOpts)
}

// TokenvestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokenvesting contract.
type TokenvestingOwnershipTransferredIterator struct {
	Event *TokenvestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenvestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenvestingOwnershipTransferred)
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
		it.Event = new(TokenvestingOwnershipTransferred)
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
func (it *TokenvestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenvestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenvestingOwnershipTransferred represents a OwnershipTransferred event raised by the Tokenvesting contract.
type TokenvestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenvesting *TokenvestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenvestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenvesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenvestingOwnershipTransferredIterator{contract: _Tokenvesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenvesting *TokenvestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenvestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenvesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenvestingOwnershipTransferred)
				if err := _Tokenvesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Tokenvesting *TokenvestingFilterer) ParseOwnershipTransferred(log types.Log) (*TokenvestingOwnershipTransferred, error) {
	event := new(TokenvestingOwnershipTransferred)
	if err := _Tokenvesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenvestingReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the Tokenvesting contract.
type TokenvestingReleasedIterator struct {
	Event *TokenvestingReleased // Event containing the contract specifics and raw log

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
func (it *TokenvestingReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenvestingReleased)
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
		it.Event = new(TokenvestingReleased)
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
func (it *TokenvestingReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenvestingReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenvestingReleased represents a Released event raised by the Tokenvesting contract.
type TokenvestingReleased struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0xb21fb52d5749b80f3182f8c6992236b5e5576681880914484d7f4c9b062e619e.
//
// Solidity: event Released(address indexed beneficiary, uint256 amount)
func (_Tokenvesting *TokenvestingFilterer) FilterReleased(opts *bind.FilterOpts, beneficiary []common.Address) (*TokenvestingReleasedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Tokenvesting.contract.FilterLogs(opts, "Released", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &TokenvestingReleasedIterator{contract: _Tokenvesting.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0xb21fb52d5749b80f3182f8c6992236b5e5576681880914484d7f4c9b062e619e.
//
// Solidity: event Released(address indexed beneficiary, uint256 amount)
func (_Tokenvesting *TokenvestingFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *TokenvestingReleased, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Tokenvesting.contract.WatchLogs(opts, "Released", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenvestingReleased)
				if err := _Tokenvesting.contract.UnpackLog(event, "Released", log); err != nil {
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

// ParseReleased is a log parse operation binding the contract event 0xb21fb52d5749b80f3182f8c6992236b5e5576681880914484d7f4c9b062e619e.
//
// Solidity: event Released(address indexed beneficiary, uint256 amount)
func (_Tokenvesting *TokenvestingFilterer) ParseReleased(log types.Log) (*TokenvestingReleased, error) {
	event := new(TokenvestingReleased)
	if err := _Tokenvesting.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenvestingRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the Tokenvesting contract.
type TokenvestingRevokedIterator struct {
	Event *TokenvestingRevoked // Event containing the contract specifics and raw log

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
func (it *TokenvestingRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenvestingRevoked)
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
		it.Event = new(TokenvestingRevoked)
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
func (it *TokenvestingRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenvestingRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenvestingRevoked represents a Revoked event raised by the Tokenvesting contract.
type TokenvestingRevoked struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0x713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac4.
//
// Solidity: event Revoked(address indexed beneficiary, uint256 amount)
func (_Tokenvesting *TokenvestingFilterer) FilterRevoked(opts *bind.FilterOpts, beneficiary []common.Address) (*TokenvestingRevokedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Tokenvesting.contract.FilterLogs(opts, "Revoked", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &TokenvestingRevokedIterator{contract: _Tokenvesting.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0x713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac4.
//
// Solidity: event Revoked(address indexed beneficiary, uint256 amount)
func (_Tokenvesting *TokenvestingFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *TokenvestingRevoked, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Tokenvesting.contract.WatchLogs(opts, "Revoked", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenvestingRevoked)
				if err := _Tokenvesting.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0x713b90881ad62c4fa8ab6bd9197fa86481fc0c11b2edba60026514281b2dbac4.
//
// Solidity: event Revoked(address indexed beneficiary, uint256 amount)
func (_Tokenvesting *TokenvestingFilterer) ParseRevoked(log types.Log) (*TokenvestingRevoked, error) {
	event := new(TokenvestingRevoked)
	if err := _Tokenvesting.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenvestingVestingScheduleCreatedIterator is returned from FilterVestingScheduleCreated and is used to iterate over the raw logs and unpacked data for VestingScheduleCreated events raised by the Tokenvesting contract.
type TokenvestingVestingScheduleCreatedIterator struct {
	Event *TokenvestingVestingScheduleCreated // Event containing the contract specifics and raw log

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
func (it *TokenvestingVestingScheduleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenvestingVestingScheduleCreated)
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
		it.Event = new(TokenvestingVestingScheduleCreated)
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
func (it *TokenvestingVestingScheduleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenvestingVestingScheduleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenvestingVestingScheduleCreated represents a VestingScheduleCreated event raised by the Tokenvesting contract.
type TokenvestingVestingScheduleCreated struct {
	Beneficiary common.Address
	TotalAmount *big.Int
	Role        uint8
	Start       *big.Int
	Duration    *big.Int
	Cliff       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVestingScheduleCreated is a free log retrieval operation binding the contract event 0x1ea8125ddaa7ea3322d4ec14f10ea84ba3a89c4306f9bf1610f4d0036ac9f90c.
//
// Solidity: event VestingScheduleCreated(address indexed beneficiary, uint256 totalAmount, uint8 role, uint256 start, uint256 duration, uint256 cliff)
func (_Tokenvesting *TokenvestingFilterer) FilterVestingScheduleCreated(opts *bind.FilterOpts, beneficiary []common.Address) (*TokenvestingVestingScheduleCreatedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Tokenvesting.contract.FilterLogs(opts, "VestingScheduleCreated", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &TokenvestingVestingScheduleCreatedIterator{contract: _Tokenvesting.contract, event: "VestingScheduleCreated", logs: logs, sub: sub}, nil
}

// WatchVestingScheduleCreated is a free log subscription operation binding the contract event 0x1ea8125ddaa7ea3322d4ec14f10ea84ba3a89c4306f9bf1610f4d0036ac9f90c.
//
// Solidity: event VestingScheduleCreated(address indexed beneficiary, uint256 totalAmount, uint8 role, uint256 start, uint256 duration, uint256 cliff)
func (_Tokenvesting *TokenvestingFilterer) WatchVestingScheduleCreated(opts *bind.WatchOpts, sink chan<- *TokenvestingVestingScheduleCreated, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Tokenvesting.contract.WatchLogs(opts, "VestingScheduleCreated", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenvestingVestingScheduleCreated)
				if err := _Tokenvesting.contract.UnpackLog(event, "VestingScheduleCreated", log); err != nil {
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

// ParseVestingScheduleCreated is a log parse operation binding the contract event 0x1ea8125ddaa7ea3322d4ec14f10ea84ba3a89c4306f9bf1610f4d0036ac9f90c.
//
// Solidity: event VestingScheduleCreated(address indexed beneficiary, uint256 totalAmount, uint8 role, uint256 start, uint256 duration, uint256 cliff)
func (_Tokenvesting *TokenvestingFilterer) ParseVestingScheduleCreated(log types.Log) (*TokenvestingVestingScheduleCreated, error) {
	event := new(TokenvestingVestingScheduleCreated)
	if err := _Tokenvesting.contract.UnpackLog(event, "VestingScheduleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
