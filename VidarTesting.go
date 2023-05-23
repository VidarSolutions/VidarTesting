// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VidarTesting

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

// VidarTestingslice is an auto generated low-level Go binding around an user-defined struct.
type VidarTestingslice struct {
	Len *big.Int
	Ptr *big.Int
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"a\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"address_to_string\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToHexString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_char\",\"type\":\"bytes1\"}],\"name\":\"charToHexChar\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hex_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_avaxAdd\",\"type\":\"string\"}],\"name\":\"checkHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compare\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_string1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_string2\",\"type\":\"string\"}],\"name\":\"getHex_Hash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"getSlice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_len\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ptr\",\"type\":\"uint256\"}],\"internalType\":\"structVidarTesting.slice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isValidAVAXAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"format\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"args\",\"type\":\"string[]\"}],\"name\":\"printf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_a\",\"type\":\"uint256\"}],\"name\":\"setA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"toHexString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052600160005534801561001557600080fd5b50611db3806100256000396000f3fe608060405234801561001057600080fd5b50600436106100ba5760003560e01c80630dbe671f146100bf57806317184d12146100dd578063274695d31461010d5780633a96fdd71461013d578063531915031461016d57806363e1cbea1461019d578063861458d8146101cd578063908c0eea146101eb5780639a329a261461021b578063a3fa67ed1461024b578063b6eec2441461027b578063c68b3787146102ab578063dd005a83146102db578063ee919d501461030b575b600080fd5b6100c761033b565b6040516100d491906110a3565b60405180910390f35b6100f760048036038101906100f2919061112a565b610341565b6040516101049190611166565b60405180910390f35b610127600480360381019061012291906111df565b610387565b604051610134919061129c565b60405180910390f35b610157600480360381019061015291906113f3565b610400565b6040516101649190611486565b60405180910390f35b610187600480360381019061018291906111df565b61041b565b604051610194919061129c565b60405180910390f35b6101b760048036038101906101b291906114cd565b6106a0565b6040516101c4919061129c565b60405180910390f35b6101d56108dc565b6040516101e291906110a3565b60405180910390f35b6102056004803603810190610200919061150d565b61091e565b6040516102129190611594565b60405180910390f35b61023560048036038101906102309190611695565b610952565b604051610242919061129c565b60405180910390f35b610265600480360381019061026091906111df565b610d77565b6040516102729190611486565b60405180910390f35b6102956004803603810190610290919061170d565b610d8a565b6040516102a29190611486565b60405180910390f35b6102c560048036038101906102c091906117ea565b610da8565b6040516102d2919061129c565b60405180910390f35b6102f560048036038101906102f091906113f3565b610f28565b604051610302919061129c565b60405180910390f35b61032560048036038101906103209190611817565b610f6f565b60405161033291906110a3565b60405180910390f35b60005481565b6000600a8260f81c60ff16101561036c5760308260f81c6103629190611880565b60f81b9050610382565b60578260f81c61037c9190611880565b60f81b90505b919050565b606061039282610d77565b156103c2576103bb8273ffffffffffffffffffffffffffffffffffffffff16601460ff166106a0565b90506103fb565b6040518060400160405280601881526020017f4e6f7420612076616c696420417661782061646472657373000000000000000081525090505b919050565b60008180519060200120838051906020012014905092915050565b60606000602a67ffffffffffffffff81111561043a576104396112c8565b5b6040519080825280601f01601f19166020018201604052801561046c5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106104a4576104a36118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610508576105076118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060005b601481101561069657600081601361055291906118e4565b600861055e9190611918565b600261056a9190611a8d565b8573ffffffffffffffffffffffffffffffffffffffff1661058b9190611b07565b9050600060108261059c9190611b38565b905060008160106105ad9190611b69565b836105b89190611ba6565b90506105c382610f7d565b858560026105d19190611918565b60026105dd9190611bdb565b815181106105ee576105ed6118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061062681610f7d565b858560026106349190611918565b60036106409190611bdb565b81518110610651576106506118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350505050808061068e90611c0f565b91505061053a565b5080915050919050565b6060600060028360026106b39190611918565b6106bd9190611bdb565b67ffffffffffffffff8111156106d6576106d56112c8565b5b6040519080825280601f01601f1916602001820160405280156107085781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106107405761073f6118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106107a4576107a36118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026107e49190611918565b6107ee9190611bdb565b90505b600181111561088e577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106108305761082f6118b5565b5b1a60f81b828281518110610847576108466118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061088790611c57565b90506107f1565b50600084146108d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c990611ccc565b60405180910390fd5b8091505092915050565b6000805a905060003a9050600061a4105a846108f891906118e4565b6109029190611bdb565b9050600082826109129190611918565b90508094505050505090565b610926611070565b600082905060006020820190506040518060400160405280835181526020018281525092505050919050565b60606000835167ffffffffffffffff811115610971576109706112c8565b5b60405190808252806020026020018201604052801561099f5781602001602082028036833780820191505090505b50905060005b8451811015610a44578481815181106109c1576109c06118b5565b5b602001015160f81c60f81b8282815181106109df576109de6118b5565b5b60200260200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815250508080610a3c90611c0f565b9150506109a5565b506000805b8451811015610a8e57848181518110610a6557610a646118b5565b5b60200260200101515182610a799190611bdb565b91508080610a8690611c0f565b915050610a49565b506000818351610a9e9190611bdb565b67ffffffffffffffff811115610ab757610ab66112c8565b5b6040519080825280601f01601f191660200182016040528015610ae95781602001600182028036833780820191505090505b509050606060008060005b8651811015610d67577f25000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916878281518110610b5357610b526118b5565b5b60200260200101517effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610ce8577f73000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191687600183610bd39190611bdb565b81518110610be457610be36118b5565b5b60200260200101517effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610ce3578851831015610ce257888381518110610c2f57610c2e6118b5565b5b6020026020010151935060005b8451811015610cc457848181518110610c5857610c576118b5565b5b602001015160f81c60f81b868480610c6f90611c0f565b955081518110610c8257610c816118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610cbc90611c0f565b915050610c3c565b508080610cd090611c0f565b9150508280610cde90611c0f565b9350505b5b610d54565b868181518110610cfb57610cfa6118b5565b5b6020026020010151858381518110610d1657610d156118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508180610d5090611c0f565b9250505b8080610d5f90611c0f565b915050610af4565b5083965050505050505092915050565b600080823b905060008114915050919050565b6000610d9f84610d9a8585610f28565b610400565b90509392505050565b60606000604067ffffffffffffffff811115610dc757610dc66112c8565b5b6040519080825280601f01601f191660200182016040528015610df95781602001600182028036833780820191505090505b50905060005b6020811015610f1e576000848260208110610e1d57610e1c6118b5565b5b1a60f81b60f81c90506000610e40601083610e389190611b38565b60f81b610341565b90506000610e5c601084610e549190611cec565b60f81b610341565b90508185600286610e6d9190611918565b81518110610e7e57610e7d6118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080856001600287610ebe9190611918565b610ec89190611bdb565b81518110610ed957610ed86118b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610f1690611c0f565b915050610dff565b5080915050919050565b606060008383604051602001610f3f929190611d59565b6040516020818303038152906040529050600081805190602001209050610f6581610da8565b9250505092915050565b600081600081905550919050565b6000600a8260ff161015610ff757816040518060400160405280600181526020017f3000000000000000000000000000000000000000000000000000000000000000815250600081518110610fd557610fd46118b5565b5b602001015160f81c60f81b60f81c610fed9190611880565b60f81b905061106b565b600a826040518060400160405280600181526020017f6100000000000000000000000000000000000000000000000000000000000000815250600081518110611043576110426118b5565b5b602001015160f81c60f81b60f81c61105b9190611880565b6110659190611ba6565b60f81b90505b919050565b604051806040016040528060008152602001600081525090565b6000819050919050565b61109d8161108a565b82525050565b60006020820190506110b86000830184611094565b92915050565b6000604051905090565b600080fd5b600080fd5b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b611107816110d2565b811461111257600080fd5b50565b600081359050611124816110fe565b92915050565b6000602082840312156111405761113f6110c8565b5b600061114e84828501611115565b91505092915050565b611160816110d2565b82525050565b600060208201905061117b6000830184611157565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006111ac82611181565b9050919050565b6111bc816111a1565b81146111c757600080fd5b50565b6000813590506111d9816111b3565b92915050565b6000602082840312156111f5576111f46110c8565b5b6000611203848285016111ca565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561124657808201518184015260208101905061122b565b60008484015250505050565b6000601f19601f8301169050919050565b600061126e8261120c565b6112788185611217565b9350611288818560208601611228565b61129181611252565b840191505092915050565b600060208201905081810360008301526112b68184611263565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61130082611252565b810181811067ffffffffffffffff8211171561131f5761131e6112c8565b5b80604052505050565b60006113326110be565b905061133e82826112f7565b919050565b600067ffffffffffffffff82111561135e5761135d6112c8565b5b61136782611252565b9050602081019050919050565b82818337600083830152505050565b600061139661139184611343565b611328565b9050828152602081018484840111156113b2576113b16112c3565b5b6113bd848285611374565b509392505050565b600082601f8301126113da576113d96112be565b5b81356113ea848260208601611383565b91505092915050565b6000806040838503121561140a576114096110c8565b5b600083013567ffffffffffffffff811115611428576114276110cd565b5b611434858286016113c5565b925050602083013567ffffffffffffffff811115611455576114546110cd565b5b611461858286016113c5565b9150509250929050565b60008115159050919050565b6114808161146b565b82525050565b600060208201905061149b6000830184611477565b92915050565b6114aa8161108a565b81146114b557600080fd5b50565b6000813590506114c7816114a1565b92915050565b600080604083850312156114e4576114e36110c8565b5b60006114f2858286016114b8565b9250506020611503858286016114b8565b9150509250929050565b600060208284031215611523576115226110c8565b5b600082013567ffffffffffffffff811115611541576115406110cd565b5b61154d848285016113c5565b91505092915050565b61155f8161108a565b82525050565b60408201600082015161157b6000850182611556565b50602082015161158e6020850182611556565b50505050565b60006040820190506115a96000830184611565565b92915050565b600067ffffffffffffffff8211156115ca576115c96112c8565b5b602082029050602081019050919050565b600080fd5b60006115f36115ee846115af565b611328565b90508083825260208201905060208402830185811115611616576116156115db565b5b835b8181101561165d57803567ffffffffffffffff81111561163b5761163a6112be565b5b80860161164889826113c5565b85526020850194505050602081019050611618565b5050509392505050565b600082601f83011261167c5761167b6112be565b5b813561168c8482602086016115e0565b91505092915050565b600080604083850312156116ac576116ab6110c8565b5b600083013567ffffffffffffffff8111156116ca576116c96110cd565b5b6116d6858286016113c5565b925050602083013567ffffffffffffffff8111156116f7576116f66110cd565b5b61170385828601611667565b9150509250929050565b600080600060608486031215611726576117256110c8565b5b600084013567ffffffffffffffff811115611744576117436110cd565b5b611750868287016113c5565b935050602084013567ffffffffffffffff811115611771576117706110cd565b5b61177d868287016113c5565b925050604084013567ffffffffffffffff81111561179e5761179d6110cd565b5b6117aa868287016113c5565b9150509250925092565b6000819050919050565b6117c7816117b4565b81146117d257600080fd5b50565b6000813590506117e4816117be565b92915050565b600060208284031215611800576117ff6110c8565b5b600061180e848285016117d5565b91505092915050565b60006020828403121561182d5761182c6110c8565b5b600061183b848285016114b8565b91505092915050565b600060ff82169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061188b82611844565b915061189683611844565b9250828201905060ff8111156118af576118ae611851565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006118ef8261108a565b91506118fa8361108a565b925082820390508181111561191257611911611851565b5b92915050565b60006119238261108a565b915061192e8361108a565b925082820261193c8161108a565b9150828204841483151761195357611952611851565b5b5092915050565b60008160011c9050919050565b6000808291508390505b60018511156119b15780860481111561198d5761198c611851565b5b600185161561199c5780820291505b80810290506119aa8561195a565b9450611971565b94509492505050565b6000826119ca5760019050611a86565b816119d85760009050611a86565b81600181146119ee57600281146119f857611a27565b6001915050611a86565b60ff841115611a0a57611a09611851565b5b8360020a915084821115611a2157611a20611851565b5b50611a86565b5060208310610133831016604e8410600b8410161715611a5c5782820a905083811115611a5757611a56611851565b5b611a86565b611a698484846001611967565b92509050818404811115611a8057611a7f611851565b5b81810290505b9392505050565b6000611a988261108a565b9150611aa38361108a565b9250611ad07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846119ba565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611b128261108a565b9150611b1d8361108a565b925082611b2d57611b2c611ad8565b5b828204905092915050565b6000611b4382611844565b9150611b4e83611844565b925082611b5e57611b5d611ad8565b5b828204905092915050565b6000611b7482611844565b9150611b7f83611844565b9250828202611b8d81611844565b9150808214611b9f57611b9e611851565b5b5092915050565b6000611bb182611844565b9150611bbc83611844565b9250828203905060ff811115611bd557611bd4611851565b5b92915050565b6000611be68261108a565b9150611bf18361108a565b9250828201905080821115611c0957611c08611851565b5b92915050565b6000611c1a8261108a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c4c57611c4b611851565b5b600182019050919050565b6000611c628261108a565b915060008203611c7557611c74611851565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000611cb6602083611217565b9150611cc182611c80565b602082019050919050565b60006020820190508181036000830152611ce581611ca9565b9050919050565b6000611cf782611844565b9150611d0283611844565b925082611d1257611d11611ad8565b5b828206905092915050565b600081905092915050565b6000611d338261120c565b611d3d8185611d1d565b9350611d4d818560208601611228565b80840191505092915050565b6000611d658285611d28565b9150611d718284611d28565b9150819050939250505056fea2646970667358221220be4fa72ec829690461d1cc36791f2cbf849c58b37b97eead503f86de9b4a780964736f6c63430008120033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_Api *ApiCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_Api *ApiSession) A() (*big.Int, error) {
	return _Api.Contract.A(&_Api.CallOpts)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_Api *ApiCallerSession) A() (*big.Int, error) {
	return _Api.Contract.A(&_Api.CallOpts)
}

// AddToString is a free data retrieval call binding the contract method 0x53191503.
//
// Solidity: function addToString(address _address) pure returns(string)
func (_Api *ApiCaller) AddToString(opts *bind.CallOpts, _address common.Address) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "addToString", _address)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddToString is a free data retrieval call binding the contract method 0x53191503.
//
// Solidity: function addToString(address _address) pure returns(string)
func (_Api *ApiSession) AddToString(_address common.Address) (string, error) {
	return _Api.Contract.AddToString(&_Api.CallOpts, _address)
}

// AddToString is a free data retrieval call binding the contract method 0x53191503.
//
// Solidity: function addToString(address _address) pure returns(string)
func (_Api *ApiCallerSession) AddToString(_address common.Address) (string, error) {
	return _Api.Contract.AddToString(&_Api.CallOpts, _address)
}

// AddressToString is a free data retrieval call binding the contract method 0x274695d3.
//
// Solidity: function address_to_string(address addr) view returns(string)
func (_Api *ApiCaller) AddressToString(opts *bind.CallOpts, addr common.Address) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "address_to_string", addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddressToString is a free data retrieval call binding the contract method 0x274695d3.
//
// Solidity: function address_to_string(address addr) view returns(string)
func (_Api *ApiSession) AddressToString(addr common.Address) (string, error) {
	return _Api.Contract.AddressToString(&_Api.CallOpts, addr)
}

// AddressToString is a free data retrieval call binding the contract method 0x274695d3.
//
// Solidity: function address_to_string(address addr) view returns(string)
func (_Api *ApiCallerSession) AddressToString(addr common.Address) (string, error) {
	return _Api.Contract.AddressToString(&_Api.CallOpts, addr)
}

// Bytes32ToHexString is a free data retrieval call binding the contract method 0xc68b3787.
//
// Solidity: function bytes32ToHexString(bytes32 _bytes32) pure returns(string)
func (_Api *ApiCaller) Bytes32ToHexString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "bytes32ToHexString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToHexString is a free data retrieval call binding the contract method 0xc68b3787.
//
// Solidity: function bytes32ToHexString(bytes32 _bytes32) pure returns(string)
func (_Api *ApiSession) Bytes32ToHexString(_bytes32 [32]byte) (string, error) {
	return _Api.Contract.Bytes32ToHexString(&_Api.CallOpts, _bytes32)
}

// Bytes32ToHexString is a free data retrieval call binding the contract method 0xc68b3787.
//
// Solidity: function bytes32ToHexString(bytes32 _bytes32) pure returns(string)
func (_Api *ApiCallerSession) Bytes32ToHexString(_bytes32 [32]byte) (string, error) {
	return _Api.Contract.Bytes32ToHexString(&_Api.CallOpts, _bytes32)
}

// CharToHexChar is a free data retrieval call binding the contract method 0x17184d12.
//
// Solidity: function charToHexChar(bytes1 _char) pure returns(bytes1)
func (_Api *ApiCaller) CharToHexChar(opts *bind.CallOpts, _char [1]byte) ([1]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "charToHexChar", _char)

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// CharToHexChar is a free data retrieval call binding the contract method 0x17184d12.
//
// Solidity: function charToHexChar(bytes1 _char) pure returns(bytes1)
func (_Api *ApiSession) CharToHexChar(_char [1]byte) ([1]byte, error) {
	return _Api.Contract.CharToHexChar(&_Api.CallOpts, _char)
}

// CharToHexChar is a free data retrieval call binding the contract method 0x17184d12.
//
// Solidity: function charToHexChar(bytes1 _char) pure returns(bytes1)
func (_Api *ApiCallerSession) CharToHexChar(_char [1]byte) ([1]byte, error) {
	return _Api.Contract.CharToHexChar(&_Api.CallOpts, _char)
}

// CheckHash is a free data retrieval call binding the contract method 0xb6eec244.
//
// Solidity: function checkHash(string hex_hash, string _orgMsg, string _avaxAdd) pure returns(bool)
func (_Api *ApiCaller) CheckHash(opts *bind.CallOpts, hex_hash string, _orgMsg string, _avaxAdd string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "checkHash", hex_hash, _orgMsg, _avaxAdd)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckHash is a free data retrieval call binding the contract method 0xb6eec244.
//
// Solidity: function checkHash(string hex_hash, string _orgMsg, string _avaxAdd) pure returns(bool)
func (_Api *ApiSession) CheckHash(hex_hash string, _orgMsg string, _avaxAdd string) (bool, error) {
	return _Api.Contract.CheckHash(&_Api.CallOpts, hex_hash, _orgMsg, _avaxAdd)
}

// CheckHash is a free data retrieval call binding the contract method 0xb6eec244.
//
// Solidity: function checkHash(string hex_hash, string _orgMsg, string _avaxAdd) pure returns(bool)
func (_Api *ApiCallerSession) CheckHash(hex_hash string, _orgMsg string, _avaxAdd string) (bool, error) {
	return _Api.Contract.CheckHash(&_Api.CallOpts, hex_hash, _orgMsg, _avaxAdd)
}

// Compare is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string a, string b) pure returns(bool)
func (_Api *ApiCaller) Compare(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "compare", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Compare is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string a, string b) pure returns(bool)
func (_Api *ApiSession) Compare(a string, b string) (bool, error) {
	return _Api.Contract.Compare(&_Api.CallOpts, a, b)
}

// Compare is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string a, string b) pure returns(bool)
func (_Api *ApiCallerSession) Compare(a string, b string) (bool, error) {
	return _Api.Contract.Compare(&_Api.CallOpts, a, b)
}

// EstimateGasFee is a free data retrieval call binding the contract method 0x861458d8.
//
// Solidity: function estimateGasFee() view returns(uint256)
func (_Api *ApiCaller) EstimateGasFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "estimateGasFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateGasFee is a free data retrieval call binding the contract method 0x861458d8.
//
// Solidity: function estimateGasFee() view returns(uint256)
func (_Api *ApiSession) EstimateGasFee() (*big.Int, error) {
	return _Api.Contract.EstimateGasFee(&_Api.CallOpts)
}

// EstimateGasFee is a free data retrieval call binding the contract method 0x861458d8.
//
// Solidity: function estimateGasFee() view returns(uint256)
func (_Api *ApiCallerSession) EstimateGasFee() (*big.Int, error) {
	return _Api.Contract.EstimateGasFee(&_Api.CallOpts)
}

// GetHexHash is a free data retrieval call binding the contract method 0xdd005a83.
//
// Solidity: function getHex_Hash(string _string1, string _string2) pure returns(string)
func (_Api *ApiCaller) GetHexHash(opts *bind.CallOpts, _string1 string, _string2 string) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getHex_Hash", _string1, _string2)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHexHash is a free data retrieval call binding the contract method 0xdd005a83.
//
// Solidity: function getHex_Hash(string _string1, string _string2) pure returns(string)
func (_Api *ApiSession) GetHexHash(_string1 string, _string2 string) (string, error) {
	return _Api.Contract.GetHexHash(&_Api.CallOpts, _string1, _string2)
}

// GetHexHash is a free data retrieval call binding the contract method 0xdd005a83.
//
// Solidity: function getHex_Hash(string _string1, string _string2) pure returns(string)
func (_Api *ApiCallerSession) GetHexHash(_string1 string, _string2 string) (string, error) {
	return _Api.Contract.GetHexHash(&_Api.CallOpts, _string1, _string2)
}

// GetSlice is a free data retrieval call binding the contract method 0x908c0eea.
//
// Solidity: function getSlice(string _str) pure returns((uint256,uint256))
func (_Api *ApiCaller) GetSlice(opts *bind.CallOpts, _str string) (VidarTestingslice, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSlice", _str)

	if err != nil {
		return *new(VidarTestingslice), err
	}

	out0 := *abi.ConvertType(out[0], new(VidarTestingslice)).(*VidarTestingslice)

	return out0, err

}

// GetSlice is a free data retrieval call binding the contract method 0x908c0eea.
//
// Solidity: function getSlice(string _str) pure returns((uint256,uint256))
func (_Api *ApiSession) GetSlice(_str string) (VidarTestingslice, error) {
	return _Api.Contract.GetSlice(&_Api.CallOpts, _str)
}

// GetSlice is a free data retrieval call binding the contract method 0x908c0eea.
//
// Solidity: function getSlice(string _str) pure returns((uint256,uint256))
func (_Api *ApiCallerSession) GetSlice(_str string) (VidarTestingslice, error) {
	return _Api.Contract.GetSlice(&_Api.CallOpts, _str)
}

// IsValidAVAXAddress is a free data retrieval call binding the contract method 0xa3fa67ed.
//
// Solidity: function isValidAVAXAddress(address _address) view returns(bool)
func (_Api *ApiCaller) IsValidAVAXAddress(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isValidAVAXAddress", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAVAXAddress is a free data retrieval call binding the contract method 0xa3fa67ed.
//
// Solidity: function isValidAVAXAddress(address _address) view returns(bool)
func (_Api *ApiSession) IsValidAVAXAddress(_address common.Address) (bool, error) {
	return _Api.Contract.IsValidAVAXAddress(&_Api.CallOpts, _address)
}

// IsValidAVAXAddress is a free data retrieval call binding the contract method 0xa3fa67ed.
//
// Solidity: function isValidAVAXAddress(address _address) view returns(bool)
func (_Api *ApiCallerSession) IsValidAVAXAddress(_address common.Address) (bool, error) {
	return _Api.Contract.IsValidAVAXAddress(&_Api.CallOpts, _address)
}

// Printf is a free data retrieval call binding the contract method 0x9a329a26.
//
// Solidity: function printf(string format, string[] args) pure returns(string)
func (_Api *ApiCaller) Printf(opts *bind.CallOpts, format string, args []string) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "printf", format, args)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Printf is a free data retrieval call binding the contract method 0x9a329a26.
//
// Solidity: function printf(string format, string[] args) pure returns(string)
func (_Api *ApiSession) Printf(format string, args []string) (string, error) {
	return _Api.Contract.Printf(&_Api.CallOpts, format, args)
}

// Printf is a free data retrieval call binding the contract method 0x9a329a26.
//
// Solidity: function printf(string format, string[] args) pure returns(string)
func (_Api *ApiCallerSession) Printf(format string, args []string) (string, error) {
	return _Api.Contract.Printf(&_Api.CallOpts, format, args)
}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_Api *ApiCaller) ToHexString(opts *bind.CallOpts, value *big.Int, length *big.Int) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "toHexString", value, length)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_Api *ApiSession) ToHexString(value *big.Int, length *big.Int) (string, error) {
	return _Api.Contract.ToHexString(&_Api.CallOpts, value, length)
}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_Api *ApiCallerSession) ToHexString(value *big.Int, length *big.Int) (string, error) {
	return _Api.Contract.ToHexString(&_Api.CallOpts, value, length)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _a) returns(uint256)
func (_Api *ApiTransactor) SetA(opts *bind.TransactOpts, _a *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setA", _a)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _a) returns(uint256)
func (_Api *ApiSession) SetA(_a *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetA(&_Api.TransactOpts, _a)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _a) returns(uint256)
func (_Api *ApiTransactorSession) SetA(_a *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetA(&_Api.TransactOpts, _a)
}
