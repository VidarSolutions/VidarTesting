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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"address_to_string\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToHexString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_char\",\"type\":\"bytes1\"}],\"name\":\"charToHexChar\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hex_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_avaxAdd\",\"type\":\"string\"}],\"name\":\"checkHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compare\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_string1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_string2\",\"type\":\"string\"}],\"name\":\"getHex_Hash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"getSlice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_len\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ptr\",\"type\":\"uint256\"}],\"internalType\":\"structVidarTesting.slice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isValidAVAXAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"format\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"args\",\"type\":\"string[]\"}],\"name\":\"printf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"toHexString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611d0e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c806317184d12146100a9578063274695d3146100d95780633a96fdd714610109578063531915031461013957806363e1cbea14610169578063861458d814610199578063908c0eea146101b75780639a329a26146101e7578063a3fa67ed14610217578063b6eec24414610247578063c68b378714610277578063dd005a83146102a7575b600080fd5b6100c360048036038101906100be919061107e565b6102d7565b6040516100d091906110ba565b60405180910390f35b6100f360048036038101906100ee9190611133565b61031d565b60405161010091906111f0565b60405180910390f35b610123600480360381019061011e9190611347565b610396565b60405161013091906113da565b60405180910390f35b610153600480360381019061014e9190611133565b6103b1565b60405161016091906111f0565b60405180910390f35b610183600480360381019061017e919061142b565b610636565b60405161019091906111f0565b60405180910390f35b6101a1610872565b6040516101ae919061147a565b60405180910390f35b6101d160048036038101906101cc9190611495565b6108b4565b6040516101de919061151c565b60405180910390f35b61020160048036038101906101fc919061161d565b6108e8565b60405161020e91906111f0565b60405180910390f35b610231600480360381019061022c9190611133565b610d0d565b60405161023e91906113da565b60405180910390f35b610261600480360381019061025c9190611695565b610d20565b60405161026e91906113da565b60405180910390f35b610291600480360381019061028c9190611772565b610d3e565b60405161029e91906111f0565b60405180910390f35b6102c160048036038101906102bc9190611347565b610ebe565b6040516102ce91906111f0565b60405180910390f35b6000600a8260f81c60ff1610156103025760308260f81c6102f891906117db565b60f81b9050610318565b60578260f81c61031291906117db565b60f81b90505b919050565b606061032882610d0d565b15610358576103518273ffffffffffffffffffffffffffffffffffffffff16601460ff16610636565b9050610391565b6040518060400160405280601881526020017f4e6f7420612076616c696420417661782061646472657373000000000000000081525090505b919050565b60008180519060200120838051906020012014905092915050565b60606000602a67ffffffffffffffff8111156103d0576103cf61121c565b5b6040519080825280601f01601f1916602001820160405280156104025781602001600182028036833780820191505090505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811061043a57610439611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061049e5761049d611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060005b601481101561062c5760008160136104e8919061183f565b60086104f49190611873565b600261050091906119e8565b8573ffffffffffffffffffffffffffffffffffffffff166105219190611a62565b905060006010826105329190611a93565b905060008160106105439190611ac4565b8361054e9190611b01565b905061055982610f05565b858560026105679190611873565b60026105739190611b36565b8151811061058457610583611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506105bc81610f05565b858560026105ca9190611873565b60036105d69190611b36565b815181106105e7576105e6611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350505050808061062490611b6a565b9150506104d0565b5080915050919050565b6060600060028360026106499190611873565b6106539190611b36565b67ffffffffffffffff81111561066c5761066b61121c565b5b6040519080825280601f01601f19166020018201604052801561069e5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106106d6576106d5611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061073a57610739611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000600184600261077a9190611873565b6107849190611b36565b90505b6001811115610824577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106107c6576107c5611810565b5b1a60f81b8282815181106107dd576107dc611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061081d90611bb2565b9050610787565b5060008414610868576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085f90611c27565b60405180910390fd5b8091505092915050565b6000805a905060003a9050600061a4105a8461088e919061183f565b6108989190611b36565b9050600082826108a89190611873565b90508094505050505090565b6108bc610ff8565b600082905060006020820190506040518060400160405280835181526020018281525092505050919050565b60606000835167ffffffffffffffff8111156109075761090661121c565b5b6040519080825280602002602001820160405280156109355781602001602082028036833780820191505090505b50905060005b84518110156109da5784818151811061095757610956611810565b5b602001015160f81c60f81b82828151811061097557610974611810565b5b60200260200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152505080806109d290611b6a565b91505061093b565b506000805b8451811015610a24578481815181106109fb576109fa611810565b5b60200260200101515182610a0f9190611b36565b91508080610a1c90611b6a565b9150506109df565b506000818351610a349190611b36565b67ffffffffffffffff811115610a4d57610a4c61121c565b5b6040519080825280601f01601f191660200182016040528015610a7f5781602001600182028036833780820191505090505b509050606060008060005b8651811015610cfd577f25000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916878281518110610ae957610ae8611810565b5b60200260200101517effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610c7e577f73000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191687600183610b699190611b36565b81518110610b7a57610b79611810565b5b60200260200101517effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610c79578851831015610c7857888381518110610bc557610bc4611810565b5b6020026020010151935060005b8451811015610c5a57848181518110610bee57610bed611810565b5b602001015160f81c60f81b868480610c0590611b6a565b955081518110610c1857610c17611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610c5290611b6a565b915050610bd2565b508080610c6690611b6a565b9150508280610c7490611b6a565b9350505b5b610cea565b868181518110610c9157610c90611810565b5b6020026020010151858381518110610cac57610cab611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508180610ce690611b6a565b9250505b8080610cf590611b6a565b915050610a8a565b5083965050505050505092915050565b600080823b905060008114915050919050565b6000610d3584610d308585610ebe565b610396565b90509392505050565b60606000604067ffffffffffffffff811115610d5d57610d5c61121c565b5b6040519080825280601f01601f191660200182016040528015610d8f5781602001600182028036833780820191505090505b50905060005b6020811015610eb4576000848260208110610db357610db2611810565b5b1a60f81b60f81c90506000610dd6601083610dce9190611a93565b60f81b6102d7565b90506000610df2601084610dea9190611c47565b60f81b6102d7565b90508185600286610e039190611873565b81518110610e1457610e13611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080856001600287610e549190611873565b610e5e9190611b36565b81518110610e6f57610e6e611810565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610eac90611b6a565b915050610d95565b5080915050919050565b606060008383604051602001610ed5929190611cb4565b6040516020818303038152906040529050600081805190602001209050610efb81610d3e565b9250505092915050565b6000600a8260ff161015610f7f57816040518060400160405280600181526020017f3000000000000000000000000000000000000000000000000000000000000000815250600081518110610f5d57610f5c611810565b5b602001015160f81c60f81b60f81c610f7591906117db565b60f81b9050610ff3565b600a826040518060400160405280600181526020017f6100000000000000000000000000000000000000000000000000000000000000815250600081518110610fcb57610fca611810565b5b602001015160f81c60f81b60f81c610fe391906117db565b610fed9190611b01565b60f81b90505b919050565b604051806040016040528060008152602001600081525090565b6000604051905090565b600080fd5b600080fd5b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b61105b81611026565b811461106657600080fd5b50565b60008135905061107881611052565b92915050565b6000602082840312156110945761109361101c565b5b60006110a284828501611069565b91505092915050565b6110b481611026565b82525050565b60006020820190506110cf60008301846110ab565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611100826110d5565b9050919050565b611110816110f5565b811461111b57600080fd5b50565b60008135905061112d81611107565b92915050565b6000602082840312156111495761114861101c565b5b60006111578482850161111e565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561119a57808201518184015260208101905061117f565b60008484015250505050565b6000601f19601f8301169050919050565b60006111c282611160565b6111cc818561116b565b93506111dc81856020860161117c565b6111e5816111a6565b840191505092915050565b6000602082019050818103600083015261120a81846111b7565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611254826111a6565b810181811067ffffffffffffffff821117156112735761127261121c565b5b80604052505050565b6000611286611012565b9050611292828261124b565b919050565b600067ffffffffffffffff8211156112b2576112b161121c565b5b6112bb826111a6565b9050602081019050919050565b82818337600083830152505050565b60006112ea6112e584611297565b61127c565b90508281526020810184848401111561130657611305611217565b5b6113118482856112c8565b509392505050565b600082601f83011261132e5761132d611212565b5b813561133e8482602086016112d7565b91505092915050565b6000806040838503121561135e5761135d61101c565b5b600083013567ffffffffffffffff81111561137c5761137b611021565b5b61138885828601611319565b925050602083013567ffffffffffffffff8111156113a9576113a8611021565b5b6113b585828601611319565b9150509250929050565b60008115159050919050565b6113d4816113bf565b82525050565b60006020820190506113ef60008301846113cb565b92915050565b6000819050919050565b611408816113f5565b811461141357600080fd5b50565b600081359050611425816113ff565b92915050565b600080604083850312156114425761144161101c565b5b600061145085828601611416565b925050602061146185828601611416565b9150509250929050565b611474816113f5565b82525050565b600060208201905061148f600083018461146b565b92915050565b6000602082840312156114ab576114aa61101c565b5b600082013567ffffffffffffffff8111156114c9576114c8611021565b5b6114d584828501611319565b91505092915050565b6114e7816113f5565b82525050565b60408201600082015161150360008501826114de565b50602082015161151660208501826114de565b50505050565b600060408201905061153160008301846114ed565b92915050565b600067ffffffffffffffff8211156115525761155161121c565b5b602082029050602081019050919050565b600080fd5b600061157b61157684611537565b61127c565b9050808382526020820190506020840283018581111561159e5761159d611563565b5b835b818110156115e557803567ffffffffffffffff8111156115c3576115c2611212565b5b8086016115d08982611319565b855260208501945050506020810190506115a0565b5050509392505050565b600082601f83011261160457611603611212565b5b8135611614848260208601611568565b91505092915050565b600080604083850312156116345761163361101c565b5b600083013567ffffffffffffffff81111561165257611651611021565b5b61165e85828601611319565b925050602083013567ffffffffffffffff81111561167f5761167e611021565b5b61168b858286016115ef565b9150509250929050565b6000806000606084860312156116ae576116ad61101c565b5b600084013567ffffffffffffffff8111156116cc576116cb611021565b5b6116d886828701611319565b935050602084013567ffffffffffffffff8111156116f9576116f8611021565b5b61170586828701611319565b925050604084013567ffffffffffffffff81111561172657611725611021565b5b61173286828701611319565b9150509250925092565b6000819050919050565b61174f8161173c565b811461175a57600080fd5b50565b60008135905061176c81611746565b92915050565b6000602082840312156117885761178761101c565b5b60006117968482850161175d565b91505092915050565b600060ff82169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006117e68261179f565b91506117f18361179f565b9250828201905060ff81111561180a576118096117ac565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061184a826113f5565b9150611855836113f5565b925082820390508181111561186d5761186c6117ac565b5b92915050565b600061187e826113f5565b9150611889836113f5565b9250828202611897816113f5565b915082820484148315176118ae576118ad6117ac565b5b5092915050565b60008160011c9050919050565b6000808291508390505b600185111561190c578086048111156118e8576118e76117ac565b5b60018516156118f75780820291505b8081029050611905856118b5565b94506118cc565b94509492505050565b60008261192557600190506119e1565b8161193357600090506119e1565b8160018114611949576002811461195357611982565b60019150506119e1565b60ff841115611965576119646117ac565b5b8360020a91508482111561197c5761197b6117ac565b5b506119e1565b5060208310610133831016604e8410600b84101617156119b75782820a9050838111156119b2576119b16117ac565b5b6119e1565b6119c484848460016118c2565b925090508184048111156119db576119da6117ac565b5b81810290505b9392505050565b60006119f3826113f5565b91506119fe836113f5565b9250611a2b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611915565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611a6d826113f5565b9150611a78836113f5565b925082611a8857611a87611a33565b5b828204905092915050565b6000611a9e8261179f565b9150611aa98361179f565b925082611ab957611ab8611a33565b5b828204905092915050565b6000611acf8261179f565b9150611ada8361179f565b9250828202611ae88161179f565b9150808214611afa57611af96117ac565b5b5092915050565b6000611b0c8261179f565b9150611b178361179f565b9250828203905060ff811115611b3057611b2f6117ac565b5b92915050565b6000611b41826113f5565b9150611b4c836113f5565b9250828201905080821115611b6457611b636117ac565b5b92915050565b6000611b75826113f5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611ba757611ba66117ac565b5b600182019050919050565b6000611bbd826113f5565b915060008203611bd057611bcf6117ac565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000611c1160208361116b565b9150611c1c82611bdb565b602082019050919050565b60006020820190508181036000830152611c4081611c04565b9050919050565b6000611c528261179f565b9150611c5d8361179f565b925082611c6d57611c6c611a33565b5b828206905092915050565b600081905092915050565b6000611c8e82611160565b611c988185611c78565b9350611ca881856020860161117c565b80840191505092915050565b6000611cc08285611c83565b9150611ccc8284611c83565b9150819050939250505056fea26469706673582212208536a813ce24e5e0ae5cc8570ccd251d53fa5600933541deab5f26d14146c73964736f6c63430008120033",
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

