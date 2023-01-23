// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// FairswapABI is the input ABI used to generate the binding from.
const FairswapABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EncInputRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"circuitGatesOperation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"receiverAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiverGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Now\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"senderAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recEntryKey\",\"type\":\"bytes32\"}],\"name\":\"initializeRecieverAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"complaint\",\"type\":\"bytes32[]\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Mverify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitmentOfKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"revealKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"complaint\",\"type\":\"bytes32[][]\"},{\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"complain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"receiverEntryKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nextStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"senderGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_commitmentOfKey\",\"type\":\"bytes32\"},{\"name\":\"_EncInputRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiverEntryKey\",\"type\":\"bytes32\"},{\"name\":\"circuitGatesOperationArray\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FairswapBin is the compiled bytecode used for deploying new contracts.
const FairswapBin = `0x6080604052600a805460ff191690553480156200001b57600080fd5b50604051620014af380380620014af8339810180604052620000419190810190620001ab565b60078054600160a060020a031916331790556001859055600284905560008381555b8151811015620000a05781818151811015156200007c57fe5b60209081029091018101516000838152600690925260409091205560010162000063565b506004829055620000b9640100000000620000c4810204565b50505050506200028e565b600a5460ff166006811115620000d657fe5b6001016006811115620000e557fe5b600a805460ff19166001836006811115620000fc57fe5b0217905550600a4204600a01600955565b6000601f820183136200011f57600080fd5b81516200013662000130826200026a565b62000243565b915081818352602084019350602081019050838560208402820111156200015c57600080fd5b60005b838110156200018c578162000175888262000196565b84525060209283019291909101906001016200015f565b5050505092915050565b6000620001a482516200028b565b9392505050565b600080600080600060a08688031215620001c457600080fd5b6000620001d2888862000196565b9550506020620001e58882890162000196565b9450506040620001f88882890162000196565b93505060606200020b8882890162000196565b92505060808601516001604060020a038111156200022857600080fd5b62000236888289016200010d565b9150509295509295909350565b6040518181016001604060020a03811182821017156200026257600080fd5b604052919050565b60006001604060020a038211156200028157600080fd5b5060209081020190565b90565b611211806200029e6000396000f3fe608060405260043610610131576000357c0100000000000000000000000000000000000000000000000000000000900480639ab164e5116100bd578063d654723611610081578063d6547236146102af578063de30b907146102cf578063eaa94830146102ef578063ee3743ab14610304578063f6fa345f1461031957610131565b80639ab164e514610216578063a035b1fe14610236578063a26303721461024b578063b1c9fe6e14610278578063ca641f861461029a57610131565b80633943380c116101045780633943380c146101ad5780633e37aeb6146101c257806344d4fd19146101d75780634fc84791146101ec57806370dea79a1461020157610131565b8063062a1398146101365780630b90457d1461016157806316fed3e2146101815780632852b71c146101a3575b600080fd5b34801561014257600080fd5b5061014b61032e565b6040516101589190611139565b60405180910390f35b34801561016d57600080fd5b5061014b61017c3660046110cd565b610334565b34801561018d57600080fd5b50610196610346565b604051610158919061111d565b6101ab610355565b005b3480156101b957600080fd5b5061014b6103fa565b3480156101ce57600080fd5b506101ab610400565b3480156101e357600080fd5b5061014b6104af565b3480156101f857600080fd5b506101966104b5565b34801561020d57600080fd5b5061014b6104c4565b34801561022257600080fd5b506101ab6102313660046110cd565b6104ca565b34801561024257600080fd5b5061014b6104fa565b34801561025757600080fd5b5061026b610266366004611086565b610500565b604051610158919061112b565b34801561028457600080fd5b5061028d610718565b6040516101589190611147565b3480156102a657600080fd5b5061014b610721565b3480156102bb57600080fd5b506101ab6102ca3660046110cd565b610727565b3480156102db57600080fd5b5061026b6102ea36600461101d565b610860565b3480156102fb57600080fd5b5061014b610b9b565b34801561031057600080fd5b506101ab610ba1565b34801561032557600080fd5b506101ab610be7565b60005481565b60066020526000908152604090205481565b600854600160a060020a031681565b600854600160a060020a0316600180600a5460ff16600681111561037557fe5b1461037f57600080fd5b600954600a42041061039057600080fd5b33600160a060020a038316146103a557600080fd5b6001543410156103ee5760405133903480156108fc02916000818181858888f193505050501580156103db573d6000803e3d6000fd5b50600a805460ff191660061790556103f6565b6103f6610ba1565b5050565b60035481565b600854600160a060020a0316600280600a5460ff16600681111561042057fe5b1461042a57600080fd5b600954600a42041061043b57600080fd5b33600160a060020a0383161461045057600080fd5b600954600a42041161046157600080fd5b600854600154604051600160a060020a039092169181156108fc0291906000818181858888f1935050505015801561049d573d6000803e3d6000fd5b5050600a805460ff1916600517905550565b60055481565b600754600160a060020a031681565b60095481565b60045481146104d857600080fd5b506008805473ffffffffffffffffffffffffffffffffffffffff191633179055565b60015481565b60008083600081518110151561051257fe5b602090810290910101516040805181815260608181018352929350600092839260059290602082018180388339019050509050600092505b6001820383101561070557600283900a87811604600114156106325760005b60208110156105cf57888460010181518110151561058357fe5b906020019060200201518160208110151561059a57fe5b1a60f860020a0282828151811015156105af57fe5b906020010190600160f860020a031916908160001a905350600101610569565b5060005b6020811015610620578581602081106105e857fe5b1a60f860020a02828260200181518110151561060057fe5b906020010190600160f860020a031916908160001a9053506001016105d3565b508051602082012094508493506106fa565b60005b602081101561067f5785816020811061064a57fe5b1a60f860020a02828281518110151561065f57fe5b906020010190600160f860020a031916908160001a905350600101610635565b5060005b60208110156106ec57888460010181518110151561069d57fe5b90602001906020020151816020811015156106b457fe5b1a60f860020a0282826020018151811015156106cc57fe5b906020010190600160f860020a031916908160001a905350600101610683565b508051602082012094508493505b60019092019161054a565b6000548514955050505050505b92915050565b600a5460ff1681565b60025481565b600754600160a060020a0316600280600a5460ff16600681111561074757fe5b1461075157600080fd5b600954600a42041061076257600080fd5b33600160a060020a0383161461077757600080fd5b6040805160208082528183019092526000916060919060208201818038833901905050905060005b60208110156107e9578581602081106107b457fe5b1a60f860020a0282828151811015156107c957fe5b906020010190600160f860020a031916908160001a90535060010161079f565b5080516020820120600254909250821461084c57600854600154604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015610839573d6000803e3d6000fd5b50600a805460ff19166005179055610859565b6003859055610859610ba1565b5050505050565b600854600090600160a060020a0316600380600a5460ff16600681111561088357fe5b1461088d57600080fd5b600954600a42041061089e57600080fd5b33600160a060020a038316146108b357600080fd5b60008060606000808860008151811015156108ca57fe5b90602001906020020151905060018a5103604051908082528060200260200182016040528015610904578160200160208202803883390190505b5092506109418a600081518110151561091957fe5b906020019060200201518a600081518110151561093257fe5b90602001906020020151610500565b151561099f57600a805460ff19166006179055600754600154604051600160a060020a03909216916108fc82150291906000818181858888f19350505050158015610990573d6000803e3d6000fd5b50600097505050505050610b93565b600660008a60008151811015156109b257fe5b906020019060200201518152602001908152602001600020549350610a068a60008151811015156109df57fe5b9060200190602002015160008151811015156109f757fe5b90602001906020020151610c8d565b915060005b60018b5103811015610ad357610a408b82815181101515610a2857fe5b906020019060200201518b8381518110151561093257fe5b1515610a9f57600a805460ff19166006179055600754600154604051600160a060020a03909216916108fc82150291906000818181858888f19350505050158015610a8f573d6000803e3d6000fd5b5060009850505050505050610b93565b610ab38b826001018151811015156109df57fe5b8482815181101515610ac157fe5b60209081029091010152600101610a0b565b50610ade8484610d46565b9450818514610b3f57600a805460ff19166005179055600854600154604051600160a060020a03909216916108fc82150291906000818181858888f19350505050158015610b30573d6000803e3d6000fd5b50600197505050505050610b93565b600a805460ff19166006179055600754600154604051600160a060020a03909216916108fc82150291906000818181858888f19350505050158015610b88573d6000803e3d6000fd5b506000975050505050505b505092915050565b60045481565b600a5460ff166006811115610bb257fe5b6001016006811115610bc057fe5b600a805460ff19166001836006811115610bd657fe5b0217905550600a4204600a01600955565b600754600160a060020a0316600380600a5460ff166006811115610c0757fe5b14610c1157600080fd5b600954600a420410610c2257600080fd5b33600160a060020a03831614610c3757600080fd5b600954600a420411610c4857600080fd5b600754600154604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015610c84573d6000803e3d6000fd5b506103f6610ba1565b604080516021808252606082810190935260009291839182919060208201818038833901905050925060005b6020811015610d05576003548160208110610cd057fe5b1a60f860020a028482815181101515610ce557fe5b906020010190600160f860020a031916908160001a905350600101610cb9565b50825160009084906020908110610d1857fe5b906020010190600160f860020a031916908160001a9053505081516020909201919091208318915050919050565b6000808360021415610e4e5760408051818152606081810183529160208201818038833901905050905060005b6020811015610dd757846000815181101515610d8b57fe5b9060200190602002015181602081101515610da257fe5b1a60f860020a028282815181101515610db757fe5b906020010190600160f860020a031916908160001a905350600101610d73565b5060005b6020811015610e4257846001815181101515610df357fe5b9060200190602002015181602081101515610e0a57fe5b1a60f860020a028282602001815181101515610e2257fe5b906020010190600160f860020a031916908160001a905350600101610ddb565b50805160209091012090505b9392505050565b6000601f82018313610e6657600080fd5b8135610e79610e748261117c565b611155565b81815260209384019390925082018360005b83811015610eb75781358601610ea18882610ec1565b8452506020928301929190910190600101610e8b565b5050505092915050565b6000601f82018313610ed257600080fd5b8135610ee0610e748261117c565b91508181835260208401935060208101905083856020840282011115610f0557600080fd5b60005b83811015610eb75781610f1b8882611011565b8452506020928301929190910190600101610f08565b6000601f82018313610f4257600080fd5b8135610f50610e748261117c565b91508181835260208401935060208101905083856020840282011115610f7557600080fd5b60005b83811015610eb75781610f8b8882611011565b8452506020928301929190910190600101610f78565b6000601f82018313610fb257600080fd5b8135610fc0610e748261117c565b91508181835260208401935060208101905083856020840282011115610fe557600080fd5b60005b83811015610eb75781610ffb8882611011565b8452506020928301929190910190600101610fe8565b6000610e4e82356111ad565b6000806040838503121561103057600080fd5b823567ffffffffffffffff81111561104757600080fd5b61105385828601610e55565b925050602083013567ffffffffffffffff81111561107057600080fd5b61107c85828601610fa1565b9150509250929050565b6000806040838503121561109957600080fd5b823567ffffffffffffffff8111156110b057600080fd5b6110bc85828601610f31565b925050602061107c85828601611011565b6000602082840312156110df57600080fd5b60006110eb8484611011565b949350505050565b6110fc8161119d565b82525050565b6110fc816111a8565b6110fc816111ad565b6110fc816111cc565b6020810161071282846110f3565b602081016107128284611102565b60208101610712828461110b565b602081016107128284611114565b60405181810167ffffffffffffffff8111828210171561117457600080fd5b604052919050565b600067ffffffffffffffff82111561119357600080fd5b5060209081020190565b6000610712826111c0565b151590565b90565b6000600782106111bc57fe5b5090565b600160a060020a031690565b6000610712826111b056fea265627a7a7230582000b92f9a75717d80a8e5968cf9c8a3f214fc54da64ac192fdb09c43710c66bbf6c6578706572696d656e74616cf50037`

// DeployFairswap deploys a new Ethereum contract, binding an instance of Fairswap to it.
func DeployFairswap(auth *bind.TransactOpts, backend bind.ContractBackend, _price *big.Int, _commitmentOfKey [32]byte, _EncInputRoot [32]byte, _receiverEntryKey [32]byte, circuitGatesOperationArray []*big.Int) (common.Address, *types.Transaction, *Fairswap, error) {
	parsed, err := abi.JSON(strings.NewReader(FairswapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FairswapBin), backend, _price, _commitmentOfKey, _EncInputRoot, _receiverEntryKey, circuitGatesOperationArray)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fairswap{FairswapCaller: FairswapCaller{contract: contract}, FairswapTransactor: FairswapTransactor{contract: contract}, FairswapFilterer: FairswapFilterer{contract: contract}}, nil
}

// Fairswap is an auto generated Go binding around an Ethereum contract.
type Fairswap struct {
	FairswapCaller     // Read-only binding to the contract
	FairswapTransactor // Write-only binding to the contract
	FairswapFilterer   // Log filterer for contract events
}

// FairswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type FairswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FairswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FairswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FairswapSession struct {
	Contract     *Fairswap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FairswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FairswapCallerSession struct {
	Contract *FairswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FairswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FairswapTransactorSession struct {
	Contract     *FairswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FairswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type FairswapRaw struct {
	Contract *Fairswap // Generic contract binding to access the raw methods on
}

// FairswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FairswapCallerRaw struct {
	Contract *FairswapCaller // Generic read-only contract binding to access the raw methods on
}

// FairswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FairswapTransactorRaw struct {
	Contract *FairswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFairswap creates a new instance of Fairswap, bound to a specific deployed contract.
func NewFairswap(address common.Address, backend bind.ContractBackend) (*Fairswap, error) {
	contract, err := bindFairswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fairswap{FairswapCaller: FairswapCaller{contract: contract}, FairswapTransactor: FairswapTransactor{contract: contract}, FairswapFilterer: FairswapFilterer{contract: contract}}, nil
}

// NewFairswapCaller creates a new read-only instance of Fairswap, bound to a specific deployed contract.
func NewFairswapCaller(address common.Address, caller bind.ContractCaller) (*FairswapCaller, error) {
	contract, err := bindFairswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FairswapCaller{contract: contract}, nil
}

// NewFairswapTransactor creates a new write-only instance of Fairswap, bound to a specific deployed contract.
func NewFairswapTransactor(address common.Address, transactor bind.ContractTransactor) (*FairswapTransactor, error) {
	contract, err := bindFairswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FairswapTransactor{contract: contract}, nil
}

// NewFairswapFilterer creates a new log filterer instance of Fairswap, bound to a specific deployed contract.
func NewFairswapFilterer(address common.Address, filterer bind.ContractFilterer) (*FairswapFilterer, error) {
	contract, err := bindFairswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FairswapFilterer{contract: contract}, nil
}

// bindFairswap binds a generic wrapper to an already deployed contract.
func bindFairswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FairswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fairswap *FairswapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fairswap.Contract.FairswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fairswap *FairswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.Contract.FairswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fairswap *FairswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fairswap.Contract.FairswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fairswap *FairswapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fairswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fairswap *FairswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fairswap *FairswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fairswap.Contract.contract.Transact(opts, method, params...)
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() constant returns(bytes32)
func (_Fairswap *FairswapCaller) EncInputRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "EncInputRoot")
	return *ret0, err
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() constant returns(bytes32)
func (_Fairswap *FairswapSession) EncInputRoot() ([32]byte, error) {
	return _Fairswap.Contract.EncInputRoot(&_Fairswap.CallOpts)
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) EncInputRoot() ([32]byte, error) {
	return _Fairswap.Contract.EncInputRoot(&_Fairswap.CallOpts)
}

// Mverify is a free data retrieval call binding the contract method 0xa2630372.
//
// Solidity: function Mverify(bytes32[] complaint, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapCaller) Mverify(opts *bind.CallOpts, complaint [][32]byte, _index *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "Mverify", complaint, _index)
	return *ret0, err
}

// Mverify is a free data retrieval call binding the contract method 0xa2630372.
//
// Solidity: function Mverify(bytes32[] complaint, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapSession) Mverify(complaint [][32]byte, _index *big.Int) (bool, error) {
	return _Fairswap.Contract.Mverify(&_Fairswap.CallOpts, complaint, _index)
}

// Mverify is a free data retrieval call binding the contract method 0xa2630372.
//
// Solidity: function Mverify(bytes32[] complaint, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapCallerSession) Mverify(complaint [][32]byte, _index *big.Int) (bool, error) {
	return _Fairswap.Contract.Mverify(&_Fairswap.CallOpts, complaint, _index)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_Fairswap *FairswapCaller) Now(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "Now")
	return *ret0, err
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_Fairswap *FairswapSession) Now() (*big.Int, error) {
	return _Fairswap.Contract.Now(&_Fairswap.CallOpts)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) Now() (*big.Int, error) {
	return _Fairswap.Contract.Now(&_Fairswap.CallOpts)
}

// CircuitGatesOperation is a free data retrieval call binding the contract method 0x0b90457d.
//
// Solidity: function circuitGatesOperation(uint256 ) constant returns(uint256)
func (_Fairswap *FairswapCaller) CircuitGatesOperation(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "circuitGatesOperation", arg0)
	return *ret0, err
}

// CircuitGatesOperation is a free data retrieval call binding the contract method 0x0b90457d.
//
// Solidity: function circuitGatesOperation(uint256 ) constant returns(uint256)
func (_Fairswap *FairswapSession) CircuitGatesOperation(arg0 *big.Int) (*big.Int, error) {
	return _Fairswap.Contract.CircuitGatesOperation(&_Fairswap.CallOpts, arg0)
}

// CircuitGatesOperation is a free data retrieval call binding the contract method 0x0b90457d.
//
// Solidity: function circuitGatesOperation(uint256 ) constant returns(uint256)
func (_Fairswap *FairswapCallerSession) CircuitGatesOperation(arg0 *big.Int) (*big.Int, error) {
	return _Fairswap.Contract.CircuitGatesOperation(&_Fairswap.CallOpts, arg0)
}

// CommitmentOfKey is a free data retrieval call binding the contract method 0xca641f86.
//
// Solidity: function commitmentOfKey() constant returns(bytes32)
func (_Fairswap *FairswapCaller) CommitmentOfKey(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "commitmentOfKey")
	return *ret0, err
}

// CommitmentOfKey is a free data retrieval call binding the contract method 0xca641f86.
//
// Solidity: function commitmentOfKey() constant returns(bytes32)
func (_Fairswap *FairswapSession) CommitmentOfKey() ([32]byte, error) {
	return _Fairswap.Contract.CommitmentOfKey(&_Fairswap.CallOpts)
}

// CommitmentOfKey is a free data retrieval call binding the contract method 0xca641f86.
//
// Solidity: function commitmentOfKey() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) CommitmentOfKey() ([32]byte, error) {
	return _Fairswap.Contract.CommitmentOfKey(&_Fairswap.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes32)
func (_Fairswap *FairswapCaller) Key(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "key")
	return *ret0, err
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes32)
func (_Fairswap *FairswapSession) Key() ([32]byte, error) {
	return _Fairswap.Contract.Key(&_Fairswap.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) Key() ([32]byte, error) {
	return _Fairswap.Contract.Key(&_Fairswap.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Fairswap *FairswapCaller) Phase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "phase")
	return *ret0, err
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Fairswap *FairswapSession) Phase() (uint8, error) {
	return _Fairswap.Contract.Phase(&_Fairswap.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Fairswap *FairswapCallerSession) Phase() (uint8, error) {
	return _Fairswap.Contract.Phase(&_Fairswap.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Fairswap *FairswapCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Fairswap *FairswapSession) Price() (*big.Int, error) {
	return _Fairswap.Contract.Price(&_Fairswap.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) Price() (*big.Int, error) {
	return _Fairswap.Contract.Price(&_Fairswap.CallOpts)
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() constant returns(address)
func (_Fairswap *FairswapCaller) ReceiverAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "receiverAddress")
	return *ret0, err
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() constant returns(address)
func (_Fairswap *FairswapSession) ReceiverAddress() (common.Address, error) {
	return _Fairswap.Contract.ReceiverAddress(&_Fairswap.CallOpts)
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() constant returns(address)
func (_Fairswap *FairswapCallerSession) ReceiverAddress() (common.Address, error) {
	return _Fairswap.Contract.ReceiverAddress(&_Fairswap.CallOpts)
}

// ReceiverEntryKey is a free data retrieval call binding the contract method 0xeaa94830.
//
// Solidity: function receiverEntryKey() constant returns(bytes32)
func (_Fairswap *FairswapCaller) ReceiverEntryKey(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "receiverEntryKey")
	return *ret0, err
}

// ReceiverEntryKey is a free data retrieval call binding the contract method 0xeaa94830.
//
// Solidity: function receiverEntryKey() constant returns(bytes32)
func (_Fairswap *FairswapSession) ReceiverEntryKey() ([32]byte, error) {
	return _Fairswap.Contract.ReceiverEntryKey(&_Fairswap.CallOpts)
}

// ReceiverEntryKey is a free data retrieval call binding the contract method 0xeaa94830.
//
// Solidity: function receiverEntryKey() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) ReceiverEntryKey() ([32]byte, error) {
	return _Fairswap.Contract.ReceiverEntryKey(&_Fairswap.CallOpts)
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() constant returns(address)
func (_Fairswap *FairswapCaller) SenderAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "senderAddress")
	return *ret0, err
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() constant returns(address)
func (_Fairswap *FairswapSession) SenderAddress() (common.Address, error) {
	return _Fairswap.Contract.SenderAddress(&_Fairswap.CallOpts)
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() constant returns(address)
func (_Fairswap *FairswapCallerSession) SenderAddress() (common.Address, error) {
	return _Fairswap.Contract.SenderAddress(&_Fairswap.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Fairswap *FairswapCaller) Timeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "timeout")
	return *ret0, err
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Fairswap *FairswapSession) Timeout() (*big.Int, error) {
	return _Fairswap.Contract.Timeout(&_Fairswap.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) Timeout() (*big.Int, error) {
	return _Fairswap.Contract.Timeout(&_Fairswap.CallOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Fairswap *FairswapTransactor) Accept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "accept")
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Fairswap *FairswapSession) Accept() (*types.Transaction, error) {
	return _Fairswap.Contract.Accept(&_Fairswap.TransactOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Fairswap *FairswapTransactorSession) Accept() (*types.Transaction, error) {
	return _Fairswap.Contract.Accept(&_Fairswap.TransactOpts)
}

// Complain is a paid mutator transaction binding the contract method 0xde30b907.
//
// Solidity: function complain(bytes32[][] complaint, uint256[] indices) returns(bool)
func (_Fairswap *FairswapTransactor) Complain(opts *bind.TransactOpts, complaint [][][32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "complain", complaint, indices)
}

// Complain is a paid mutator transaction binding the contract method 0xde30b907.
//
// Solidity: function complain(bytes32[][] complaint, uint256[] indices) returns(bool)
func (_Fairswap *FairswapSession) Complain(complaint [][][32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _Fairswap.Contract.Complain(&_Fairswap.TransactOpts, complaint, indices)
}

// Complain is a paid mutator transaction binding the contract method 0xde30b907.
//
// Solidity: function complain(bytes32[][] complaint, uint256[] indices) returns(bool)
func (_Fairswap *FairswapTransactorSession) Complain(complaint [][][32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _Fairswap.Contract.Complain(&_Fairswap.TransactOpts, complaint, indices)
}

// InitializeRecieverAddress is a paid mutator transaction binding the contract method 0x9ab164e5.
//
// Solidity: function initializeRecieverAddress(bytes32 recEntryKey) returns()
func (_Fairswap *FairswapTransactor) InitializeRecieverAddress(opts *bind.TransactOpts, recEntryKey [32]byte) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "initializeRecieverAddress", recEntryKey)
}

// InitializeRecieverAddress is a paid mutator transaction binding the contract method 0x9ab164e5.
//
// Solidity: function initializeRecieverAddress(bytes32 recEntryKey) returns()
func (_Fairswap *FairswapSession) InitializeRecieverAddress(recEntryKey [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.InitializeRecieverAddress(&_Fairswap.TransactOpts, recEntryKey)
}

// InitializeRecieverAddress is a paid mutator transaction binding the contract method 0x9ab164e5.
//
// Solidity: function initializeRecieverAddress(bytes32 recEntryKey) returns()
func (_Fairswap *FairswapTransactorSession) InitializeRecieverAddress(recEntryKey [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.InitializeRecieverAddress(&_Fairswap.TransactOpts, recEntryKey)
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapTransactor) NextStage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "nextStage")
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapSession) NextStage() (*types.Transaction, error) {
	return _Fairswap.Contract.NextStage(&_Fairswap.TransactOpts)
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapTransactorSession) NextStage() (*types.Transaction, error) {
	return _Fairswap.Contract.NextStage(&_Fairswap.TransactOpts)
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Fairswap *FairswapTransactor) ReceiverGetEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "receiverGetEther")
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Fairswap *FairswapSession) ReceiverGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.ReceiverGetEther(&_Fairswap.TransactOpts)
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Fairswap *FairswapTransactorSession) ReceiverGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.ReceiverGetEther(&_Fairswap.TransactOpts)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapTransactor) RevealKey(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "revealKey", _key)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapSession) RevealKey(_key [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.RevealKey(&_Fairswap.TransactOpts, _key)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapTransactorSession) RevealKey(_key [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.RevealKey(&_Fairswap.TransactOpts, _key)
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Fairswap *FairswapTransactor) SenderGetEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "senderGetEther")
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Fairswap *FairswapSession) SenderGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.SenderGetEther(&_Fairswap.TransactOpts)
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Fairswap *FairswapTransactorSession) SenderGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.SenderGetEther(&_Fairswap.TransactOpts)
}
