// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Oracle is an auto generated low-level Go binding around an user-defined struct.
type Oracle struct {
	Name       common.Address
	Ip         string
	Port       string
	Reputation *big.Int
}

// RegisterMetaData contains all meta data concerning the Register contract.
var RegisterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_registerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reportFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"daoOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blacklistedOracle\",\"type\":\"address\"}],\"name\":\"OracleBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"OracleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reportedOracle\",\"type\":\"address\"}],\"name\":\"OracleReported\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"name\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"}],\"internalType\":\"structOracle[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistrationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOracleRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"reportOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateRegistrationFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateReportFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161216c38038061216c83398181016040528101906100319190610118565b82600181905550816002819055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050610168565b5f5ffd5b5f819050919050565b61009d8161008b565b81146100a7575f5ffd5b50565b5f815190506100b881610094565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100e7826100be565b9050919050565b6100f7816100dd565b8114610101575f5ffd5b50565b5f81519050610112816100ee565b92915050565b5f5f5f6060848603121561012f5761012e610087565b5b5f61013c868287016100aa565b935050602061014d868287016100aa565b925050604061015e86828701610104565b9150509250925092565b611ff7806101755f395ff3fe608060405260043610610085575f3560e01c806369366c481161005857806369366c48146101375780639a6bea1814610167578063af582c6b14610191578063e1704d83146101c1578063fe575a87146101f157610085565b80630946e807146100895780630d1db1f0146100b35780633ffbd47f146100dd57806340884c521461010d575b5f5ffd5b348015610094575f5ffd5b5061009d61022d565b6040516100aa9190611338565b60405180910390f35b3480156100be575f5ffd5b506100c7610236565b6040516100d4919061136b565b60405180910390f35b6100f760048036038101906100f291906114d1565b61044a565b604051610104919061136b565b60405180910390f35b348015610118575f5ffd5b5061012161064f565b60405161012e9190611717565b60405180910390f35b610151600480360381019061014c9190611761565b610828565b60405161015e919061136b565b60405180910390f35b348015610172575f5ffd5b5061017b6108c8565b6040516101889190611338565b60405180910390f35b6101ab60048036038101906101a69190611761565b6108d1565b6040516101b8919061136b565b60405180910390f35b6101db60048036038101906101d691906117b6565b610971565b6040516101e8919061136b565b60405180910390f35b3480156101fc575f5ffd5b50610217600480360381019061021291906117b6565b610f44565b604051610224919061136b565b60405180910390f35b5f600154905090565b5f5f5f90505b5f80549050811015610442575f5f828154811061025c5761025b6117e1565b5b905f5260205f2090600402016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820180546102d79061183b565b80601f01602080910402602001604051908101604052809291908181526020018280546103039061183b565b801561034e5780601f106103255761010080835404028352916020019161034e565b820191905f5260205f20905b81548152906001019060200180831161033157829003601f168201915b505050505081526020016002820180546103679061183b565b80601f01602080910402602001604051908101604052809291908181526020018280546103939061183b565b80156103de5780601f106103b5576101008083540402835291602001916103de565b820191905f5260205f20905b8154815290600101906020018083116103c157829003601f168201915b5050505050815260200160038201548152505090503373ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff160361043457600192505050610447565b50808060010191505061023c565b505f90505b90565b5f600154341015610490576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610487906118c5565b60405180910390fd5b60065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff161561051a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051190611953565b60405180910390fd5b5f60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020016001815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816105d29190611b11565b5060408201518160020190816105e89190611b11565b506060820151816003015550503373ffffffffffffffffffffffffffffffffffffffff167f8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c9848460405161063d929190611c18565b60405180910390a26001905092915050565b60605f805480602002602001604051908101604052809291908181526020015f905b8282101561081f578382905f5260205f2090600402016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820180546106f69061183b565b80601f01602080910402602001604051908101604052809291908181526020018280546107229061183b565b801561076d5780601f106107445761010080835404028352916020019161076d565b820191905f5260205f20905b81548152906001019060200180831161075057829003601f168201915b505050505081526020016002820180546107869061183b565b80601f01602080910402602001604051908101604052809291908181526020018280546107b29061183b565b80156107fd5780601f106107d4576101008083540402835291602001916107fd565b820191905f5260205f20905b8154815290600101906020018083116107e057829003601f168201915b5050505050815260200160038201548152505081526020019060010190610671565b50505050905090565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108af90611c97565b60405180910390fd5b8160028190555060019050919050565b5f600254905090565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610961576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095890611c97565b60405180910390fd5b8160018190555060019050919050565b5f5f5f90505f5f90505b5f80549050811015610a64573373ffffffffffffffffffffffffffffffffffffffff165f82815481106109b1576109b06117e1565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148015610a49575060065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16155b15610a575760019150610a64565b808060010191505061097b565b5080610aa5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9c90611cff565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0a90611d67565b60405180910390fd5b60065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610b9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9490611dcf565b60405180910390fd5b60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610c62576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5990611e37565b60405180910390fd5b600254341015610ca7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9e90611e9f565b60405180910390fd5b600160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190610d8490611eea565b91905055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d60405160405180910390a3600160025f80549050610df69190611f5e565b610e009190611f8e565b60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410610f3a57600160065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508273ffffffffffffffffffffffffffffffffffffffff167f8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf60405160405180910390a260025460075f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f299190611f8e565b92505081905550610f3983610f96565b5b6001915050919050565b5f60065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f5f90505f5f90505b5f805490508110156110915760045f5f8381548110610fc157610fc06117e1565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff161561108457818061108090611eea565b9250505b8080600101915050610f9f565b505f816001546110a19190611f5e565b90505f5f90505b5f8054905081101561131a5760045f5f83815481106110ca576110c96117e1565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff161561130d575f8181548110611191576111906117e1565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc8360075f5f86815481106111ef576111ee6117e1565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546112619190611f8e565b90811502906040515f60405180830381858888f19350505050158015611289573d5f5f3e3d5ffd5b505f60075f5f84815481106112a1576112a06117e1565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505b80806001019150506110a8565b50505050565b5f819050919050565b61133281611320565b82525050565b5f60208201905061134b5f830184611329565b92915050565b5f8115159050919050565b61136581611351565b82525050565b5f60208201905061137e5f83018461135c565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6113e38261139d565b810181811067ffffffffffffffff82111715611402576114016113ad565b5b80604052505050565b5f611414611384565b905061142082826113da565b919050565b5f67ffffffffffffffff82111561143f5761143e6113ad565b5b6114488261139d565b9050602081019050919050565b828183375f83830152505050565b5f61147561147084611425565b61140b565b90508281526020810184848401111561149157611490611399565b5b61149c848285611455565b509392505050565b5f82601f8301126114b8576114b7611395565b5b81356114c8848260208601611463565b91505092915050565b5f5f604083850312156114e7576114e661138d565b5b5f83013567ffffffffffffffff81111561150457611503611391565b5b611510858286016114a4565b925050602083013567ffffffffffffffff81111561153157611530611391565b5b61153d858286016114a4565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61159982611570565b9050919050565b6115a98161158f565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6115e1826115af565b6115eb81856115b9565b93506115fb8185602086016115c9565b6116048161139d565b840191505092915050565b61161881611320565b82525050565b5f608083015f8301516116335f8601826115a0565b506020830151848203602086015261164b82826115d7565b9150506040830151848203604086015261166582826115d7565b915050606083015161167a606086018261160f565b508091505092915050565b5f611690838361161e565b905092915050565b5f602082019050919050565b5f6116ae82611547565b6116b88185611551565b9350836020820285016116ca85611561565b805f5b8581101561170557848403895281516116e68582611685565b94506116f183611698565b925060208a019950506001810190506116cd565b50829750879550505050505092915050565b5f6020820190508181035f83015261172f81846116a4565b905092915050565b61174081611320565b811461174a575f5ffd5b50565b5f8135905061175b81611737565b92915050565b5f602082840312156117765761177561138d565b5b5f6117838482850161174d565b91505092915050565b6117958161158f565b811461179f575f5ffd5b50565b5f813590506117b08161178c565b92915050565b5f602082840312156117cb576117ca61138d565b5b5f6117d8848285016117a2565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061185257607f821691505b6020821081036118655761186461180e565b5b50919050565b5f82825260208201905092915050565b7f496e73756666696369656e7420726567697374726174696f6e206665650000005f82015250565b5f6118af601d8361186b565b91506118ba8261187b565b602082019050919050565b5f6020820190508181035f8301526118dc816118a3565b9050919050565b7f43616e6e6f74207265676973746572206120626c61636b6c6973746564206f725f8201527f61636c6500000000000000000000000000000000000000000000000000000000602082015250565b5f61193d60248361186b565b9150611948826118e3565b604082019050919050565b5f6020820190508181035f83015261196a81611931565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026119cd7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611992565b6119d78683611992565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611a12611a0d611a0884611320565b6119ef565b611320565b9050919050565b5f819050919050565b611a2b836119f8565b611a3f611a3782611a19565b84845461199e565b825550505050565b5f5f905090565b611a56611a47565b611a61818484611a22565b505050565b5b81811015611a8457611a795f82611a4e565b600181019050611a67565b5050565b601f821115611ac957611a9a81611971565b611aa384611983565b81016020851015611ab2578190505b611ac6611abe85611983565b830182611a66565b50505b505050565b5f82821c905092915050565b5f611ae95f1984600802611ace565b1980831691505092915050565b5f611b018383611ada565b9150826002028217905092915050565b611b1a826115af565b67ffffffffffffffff811115611b3357611b326113ad565b5b611b3d825461183b565b611b48828285611a88565b5f60209050601f831160018114611b79575f8415611b67578287015190505b611b718582611af6565b865550611bd8565b601f198416611b8786611971565b5f5b82811015611bae57848901518255600182019150602085019450602081019050611b89565b86831015611bcb5784890151611bc7601f891682611ada565b8355505b6001600288020188555050505b505050505050565b5f611bea826115af565b611bf4818561186b565b9350611c048185602086016115c9565b611c0d8161139d565b840191505092915050565b5f6040820190508181035f830152611c308185611be0565b90508181036020830152611c448184611be0565b90509392505050565b7f4e6f7420617574686f72697a65640000000000000000000000000000000000005f82015250565b5f611c81600e8361186b565b9150611c8c82611c4d565b602082019050919050565b5f6020820190508181035f830152611cae81611c75565b9050919050565b7f4e6f7420616e20617574686f72697a6564206f7261636c6500000000000000005f82015250565b5f611ce960188361186b565b9150611cf482611cb5565b602082019050919050565b5f6020820190508181035f830152611d1681611cdd565b9050919050565b7f43616e6e6f74207265706f727420796f757273656c66000000000000000000005f82015250565b5f611d5160168361186b565b9150611d5c82611d1d565b602082019050919050565b5f6020820190508181035f830152611d7e81611d45565b9050919050565b7f4f7261636c6520697320616c726561647920626c61636b6c69737465640000005f82015250565b5f611db9601d8361186b565b9150611dc482611d85565b602082019050919050565b5f6020820190508181035f830152611de681611dad565b9050919050565b7f4f7261636c6520616c7265616479207265706f7274656420627920796f7500005f82015250565b5f611e21601e8361186b565b9150611e2c82611ded565b602082019050919050565b5f6020820190508181035f830152611e4e81611e15565b9050919050565b7f496e73756666696369656e74207265706f7274206665650000000000000000005f82015250565b5f611e8960178361186b565b9150611e9482611e55565b602082019050919050565b5f6020820190508181035f830152611eb681611e7d565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611ef482611320565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611f2657611f25611ebd565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611f6882611320565b9150611f7383611320565b925082611f8357611f82611f31565b5b828204905092915050565b5f611f9882611320565b9150611fa383611320565b9250828201905080821115611fbb57611fba611ebd565b5b9291505056fea2646970667358221220b45e5ea9bd7f2eaa80ad3abc906ab492fc2db96a5de9021c87cc8169982adb7f64736f6c634300081c0033",
}

// RegisterABI is the input ABI used to generate the binding from.
// Deprecated: Use RegisterMetaData.ABI instead.
var RegisterABI = RegisterMetaData.ABI

// RegisterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegisterMetaData.Bin instead.
var RegisterBin = RegisterMetaData.Bin

// DeployRegister deploys a new Ethereum contract, binding an instance of Register to it.
func DeployRegister(auth *bind.TransactOpts, backend bind.ContractBackend, _registerFee *big.Int, _reportFee *big.Int, daoOwner common.Address) (common.Address, *types.Transaction, *Register, error) {
	parsed, err := RegisterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegisterBin), backend, _registerFee, _reportFee, daoOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Register{RegisterCaller: RegisterCaller{contract: contract}, RegisterTransactor: RegisterTransactor{contract: contract}, RegisterFilterer: RegisterFilterer{contract: contract}}, nil
}

// Register is an auto generated Go binding around an Ethereum contract.
type Register struct {
	RegisterCaller     // Read-only binding to the contract
	RegisterTransactor // Write-only binding to the contract
	RegisterFilterer   // Log filterer for contract events
}

// RegisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegisterSession struct {
	Contract     *Register         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegisterCallerSession struct {
	Contract *RegisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegisterTransactorSession struct {
	Contract     *RegisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegisterRaw struct {
	Contract *Register // Generic contract binding to access the raw methods on
}

// RegisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegisterCallerRaw struct {
	Contract *RegisterCaller // Generic read-only contract binding to access the raw methods on
}

// RegisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegisterTransactorRaw struct {
	Contract *RegisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegister creates a new instance of Register, bound to a specific deployed contract.
func NewRegister(address common.Address, backend bind.ContractBackend) (*Register, error) {
	contract, err := bindRegister(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Register{RegisterCaller: RegisterCaller{contract: contract}, RegisterTransactor: RegisterTransactor{contract: contract}, RegisterFilterer: RegisterFilterer{contract: contract}}, nil
}

// NewRegisterCaller creates a new read-only instance of Register, bound to a specific deployed contract.
func NewRegisterCaller(address common.Address, caller bind.ContractCaller) (*RegisterCaller, error) {
	contract, err := bindRegister(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegisterCaller{contract: contract}, nil
}

// NewRegisterTransactor creates a new write-only instance of Register, bound to a specific deployed contract.
func NewRegisterTransactor(address common.Address, transactor bind.ContractTransactor) (*RegisterTransactor, error) {
	contract, err := bindRegister(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegisterTransactor{contract: contract}, nil
}

// NewRegisterFilterer creates a new log filterer instance of Register, bound to a specific deployed contract.
func NewRegisterFilterer(address common.Address, filterer bind.ContractFilterer) (*RegisterFilterer, error) {
	contract, err := bindRegister(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegisterFilterer{contract: contract}, nil
}

// bindRegister binds a generic wrapper to an already deployed contract.
func bindRegister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegisterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Register *RegisterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Register.Contract.RegisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Register *RegisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register.Contract.RegisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Register *RegisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Register.Contract.RegisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Register *RegisterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Register.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Register *RegisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Register *RegisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Register.Contract.contract.Transact(opts, method, params...)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns((address,string,string,uint256)[])
func (_Register *RegisterCaller) GetOracles(opts *bind.CallOpts) ([]Oracle, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "getOracles")

	if err != nil {
		return *new([]Oracle), err
	}

	out0 := *abi.ConvertType(out[0], new([]Oracle)).(*[]Oracle)

	return out0, err

}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns((address,string,string,uint256)[])
func (_Register *RegisterSession) GetOracles() ([]Oracle, error) {
	return _Register.Contract.GetOracles(&_Register.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns((address,string,string,uint256)[])
func (_Register *RegisterCallerSession) GetOracles() ([]Oracle, error) {
	return _Register.Contract.GetOracles(&_Register.CallOpts)
}

// GetRegistrationFee is a free data retrieval call binding the contract method 0x0946e807.
//
// Solidity: function getRegistrationFee() view returns(uint256)
func (_Register *RegisterCaller) GetRegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "getRegistrationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRegistrationFee is a free data retrieval call binding the contract method 0x0946e807.
//
// Solidity: function getRegistrationFee() view returns(uint256)
func (_Register *RegisterSession) GetRegistrationFee() (*big.Int, error) {
	return _Register.Contract.GetRegistrationFee(&_Register.CallOpts)
}

// GetRegistrationFee is a free data retrieval call binding the contract method 0x0946e807.
//
// Solidity: function getRegistrationFee() view returns(uint256)
func (_Register *RegisterCallerSession) GetRegistrationFee() (*big.Int, error) {
	return _Register.Contract.GetRegistrationFee(&_Register.CallOpts)
}

// GetReportFee is a free data retrieval call binding the contract method 0x9a6bea18.
//
// Solidity: function getReportFee() view returns(uint256)
func (_Register *RegisterCaller) GetReportFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "getReportFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportFee is a free data retrieval call binding the contract method 0x9a6bea18.
//
// Solidity: function getReportFee() view returns(uint256)
func (_Register *RegisterSession) GetReportFee() (*big.Int, error) {
	return _Register.Contract.GetReportFee(&_Register.CallOpts)
}

// GetReportFee is a free data retrieval call binding the contract method 0x9a6bea18.
//
// Solidity: function getReportFee() view returns(uint256)
func (_Register *RegisterCallerSession) GetReportFee() (*big.Int, error) {
	return _Register.Contract.GetReportFee(&_Register.CallOpts)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address oracleAddress) view returns(bool)
func (_Register *RegisterCaller) IsBlacklisted(opts *bind.CallOpts, oracleAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "isBlacklisted", oracleAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address oracleAddress) view returns(bool)
func (_Register *RegisterSession) IsBlacklisted(oracleAddress common.Address) (bool, error) {
	return _Register.Contract.IsBlacklisted(&_Register.CallOpts, oracleAddress)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address oracleAddress) view returns(bool)
func (_Register *RegisterCallerSession) IsBlacklisted(oracleAddress common.Address) (bool, error) {
	return _Register.Contract.IsBlacklisted(&_Register.CallOpts, oracleAddress)
}

// IsOracleRegistered is a free data retrieval call binding the contract method 0x0d1db1f0.
//
// Solidity: function isOracleRegistered() view returns(bool)
func (_Register *RegisterCaller) IsOracleRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "isOracleRegistered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOracleRegistered is a free data retrieval call binding the contract method 0x0d1db1f0.
//
// Solidity: function isOracleRegistered() view returns(bool)
func (_Register *RegisterSession) IsOracleRegistered() (bool, error) {
	return _Register.Contract.IsOracleRegistered(&_Register.CallOpts)
}

// IsOracleRegistered is a free data retrieval call binding the contract method 0x0d1db1f0.
//
// Solidity: function isOracleRegistered() view returns(bool)
func (_Register *RegisterCallerSession) IsOracleRegistered() (bool, error) {
	return _Register.Contract.IsOracleRegistered(&_Register.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string ip, string port) payable returns(bool)
func (_Register *RegisterTransactor) Register(opts *bind.TransactOpts, ip string, port string) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "register", ip, port)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string ip, string port) payable returns(bool)
func (_Register *RegisterSession) Register(ip string, port string) (*types.Transaction, error) {
	return _Register.Contract.Register(&_Register.TransactOpts, ip, port)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string ip, string port) payable returns(bool)
func (_Register *RegisterTransactorSession) Register(ip string, port string) (*types.Transaction, error) {
	return _Register.Contract.Register(&_Register.TransactOpts, ip, port)
}

// ReportOracle is a paid mutator transaction binding the contract method 0xe1704d83.
//
// Solidity: function reportOracle(address oracleAddress) payable returns(bool)
func (_Register *RegisterTransactor) ReportOracle(opts *bind.TransactOpts, oracleAddress common.Address) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "reportOracle", oracleAddress)
}

// ReportOracle is a paid mutator transaction binding the contract method 0xe1704d83.
//
// Solidity: function reportOracle(address oracleAddress) payable returns(bool)
func (_Register *RegisterSession) ReportOracle(oracleAddress common.Address) (*types.Transaction, error) {
	return _Register.Contract.ReportOracle(&_Register.TransactOpts, oracleAddress)
}

// ReportOracle is a paid mutator transaction binding the contract method 0xe1704d83.
//
// Solidity: function reportOracle(address oracleAddress) payable returns(bool)
func (_Register *RegisterTransactorSession) ReportOracle(oracleAddress common.Address) (*types.Transaction, error) {
	return _Register.Contract.ReportOracle(&_Register.TransactOpts, oracleAddress)
}

// UpdateRegistrationFee is a paid mutator transaction binding the contract method 0xaf582c6b.
//
// Solidity: function updateRegistrationFee(uint256 fee) payable returns(bool)
func (_Register *RegisterTransactor) UpdateRegistrationFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "updateRegistrationFee", fee)
}

// UpdateRegistrationFee is a paid mutator transaction binding the contract method 0xaf582c6b.
//
// Solidity: function updateRegistrationFee(uint256 fee) payable returns(bool)
func (_Register *RegisterSession) UpdateRegistrationFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateRegistrationFee(&_Register.TransactOpts, fee)
}

// UpdateRegistrationFee is a paid mutator transaction binding the contract method 0xaf582c6b.
//
// Solidity: function updateRegistrationFee(uint256 fee) payable returns(bool)
func (_Register *RegisterTransactorSession) UpdateRegistrationFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateRegistrationFee(&_Register.TransactOpts, fee)
}

// UpdateReportFee is a paid mutator transaction binding the contract method 0x69366c48.
//
// Solidity: function updateReportFee(uint256 fee) payable returns(bool)
func (_Register *RegisterTransactor) UpdateReportFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "updateReportFee", fee)
}

// UpdateReportFee is a paid mutator transaction binding the contract method 0x69366c48.
//
// Solidity: function updateReportFee(uint256 fee) payable returns(bool)
func (_Register *RegisterSession) UpdateReportFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateReportFee(&_Register.TransactOpts, fee)
}

// UpdateReportFee is a paid mutator transaction binding the contract method 0x69366c48.
//
// Solidity: function updateReportFee(uint256 fee) payable returns(bool)
func (_Register *RegisterTransactorSession) UpdateReportFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateReportFee(&_Register.TransactOpts, fee)
}

// RegisterOracleBlacklistedIterator is returned from FilterOracleBlacklisted and is used to iterate over the raw logs and unpacked data for OracleBlacklisted events raised by the Register contract.
type RegisterOracleBlacklistedIterator struct {
	Event *RegisterOracleBlacklisted // Event containing the contract specifics and raw log

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
func (it *RegisterOracleBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterOracleBlacklisted)
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
		it.Event = new(RegisterOracleBlacklisted)
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
func (it *RegisterOracleBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterOracleBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterOracleBlacklisted represents a OracleBlacklisted event raised by the Register contract.
type RegisterOracleBlacklisted struct {
	BlacklistedOracle common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOracleBlacklisted is a free log retrieval operation binding the contract event 0x8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf.
//
// Solidity: event OracleBlacklisted(address indexed blacklistedOracle)
func (_Register *RegisterFilterer) FilterOracleBlacklisted(opts *bind.FilterOpts, blacklistedOracle []common.Address) (*RegisterOracleBlacklistedIterator, error) {

	var blacklistedOracleRule []interface{}
	for _, blacklistedOracleItem := range blacklistedOracle {
		blacklistedOracleRule = append(blacklistedOracleRule, blacklistedOracleItem)
	}

	logs, sub, err := _Register.contract.FilterLogs(opts, "OracleBlacklisted", blacklistedOracleRule)
	if err != nil {
		return nil, err
	}
	return &RegisterOracleBlacklistedIterator{contract: _Register.contract, event: "OracleBlacklisted", logs: logs, sub: sub}, nil
}

// WatchOracleBlacklisted is a free log subscription operation binding the contract event 0x8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf.
//
// Solidity: event OracleBlacklisted(address indexed blacklistedOracle)
func (_Register *RegisterFilterer) WatchOracleBlacklisted(opts *bind.WatchOpts, sink chan<- *RegisterOracleBlacklisted, blacklistedOracle []common.Address) (event.Subscription, error) {

	var blacklistedOracleRule []interface{}
	for _, blacklistedOracleItem := range blacklistedOracle {
		blacklistedOracleRule = append(blacklistedOracleRule, blacklistedOracleItem)
	}

	logs, sub, err := _Register.contract.WatchLogs(opts, "OracleBlacklisted", blacklistedOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterOracleBlacklisted)
				if err := _Register.contract.UnpackLog(event, "OracleBlacklisted", log); err != nil {
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

// ParseOracleBlacklisted is a log parse operation binding the contract event 0x8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf.
//
// Solidity: event OracleBlacklisted(address indexed blacklistedOracle)
func (_Register *RegisterFilterer) ParseOracleBlacklisted(log types.Log) (*RegisterOracleBlacklisted, error) {
	event := new(RegisterOracleBlacklisted)
	if err := _Register.contract.UnpackLog(event, "OracleBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegisterOracleRegisteredIterator is returned from FilterOracleRegistered and is used to iterate over the raw logs and unpacked data for OracleRegistered events raised by the Register contract.
type RegisterOracleRegisteredIterator struct {
	Event *RegisterOracleRegistered // Event containing the contract specifics and raw log

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
func (it *RegisterOracleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterOracleRegistered)
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
		it.Event = new(RegisterOracleRegistered)
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
func (it *RegisterOracleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterOracleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterOracleRegistered represents a OracleRegistered event raised by the Register contract.
type RegisterOracleRegistered struct {
	OracleAddress common.Address
	Ip            string
	Port          string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleRegistered is a free log retrieval operation binding the contract event 0x8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c9.
//
// Solidity: event OracleRegistered(address indexed oracleAddress, string ip, string port)
func (_Register *RegisterFilterer) FilterOracleRegistered(opts *bind.FilterOpts, oracleAddress []common.Address) (*RegisterOracleRegisteredIterator, error) {

	var oracleAddressRule []interface{}
	for _, oracleAddressItem := range oracleAddress {
		oracleAddressRule = append(oracleAddressRule, oracleAddressItem)
	}

	logs, sub, err := _Register.contract.FilterLogs(opts, "OracleRegistered", oracleAddressRule)
	if err != nil {
		return nil, err
	}
	return &RegisterOracleRegisteredIterator{contract: _Register.contract, event: "OracleRegistered", logs: logs, sub: sub}, nil
}

// WatchOracleRegistered is a free log subscription operation binding the contract event 0x8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c9.
//
// Solidity: event OracleRegistered(address indexed oracleAddress, string ip, string port)
func (_Register *RegisterFilterer) WatchOracleRegistered(opts *bind.WatchOpts, sink chan<- *RegisterOracleRegistered, oracleAddress []common.Address) (event.Subscription, error) {

	var oracleAddressRule []interface{}
	for _, oracleAddressItem := range oracleAddress {
		oracleAddressRule = append(oracleAddressRule, oracleAddressItem)
	}

	logs, sub, err := _Register.contract.WatchLogs(opts, "OracleRegistered", oracleAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterOracleRegistered)
				if err := _Register.contract.UnpackLog(event, "OracleRegistered", log); err != nil {
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

// ParseOracleRegistered is a log parse operation binding the contract event 0x8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c9.
//
// Solidity: event OracleRegistered(address indexed oracleAddress, string ip, string port)
func (_Register *RegisterFilterer) ParseOracleRegistered(log types.Log) (*RegisterOracleRegistered, error) {
	event := new(RegisterOracleRegistered)
	if err := _Register.contract.UnpackLog(event, "OracleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegisterOracleReportedIterator is returned from FilterOracleReported and is used to iterate over the raw logs and unpacked data for OracleReported events raised by the Register contract.
type RegisterOracleReportedIterator struct {
	Event *RegisterOracleReported // Event containing the contract specifics and raw log

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
func (it *RegisterOracleReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterOracleReported)
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
		it.Event = new(RegisterOracleReported)
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
func (it *RegisterOracleReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterOracleReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterOracleReported represents a OracleReported event raised by the Register contract.
type RegisterOracleReported struct {
	Reporter       common.Address
	ReportedOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOracleReported is a free log retrieval operation binding the contract event 0xc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d.
//
// Solidity: event OracleReported(address indexed reporter, address indexed reportedOracle)
func (_Register *RegisterFilterer) FilterOracleReported(opts *bind.FilterOpts, reporter []common.Address, reportedOracle []common.Address) (*RegisterOracleReportedIterator, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var reportedOracleRule []interface{}
	for _, reportedOracleItem := range reportedOracle {
		reportedOracleRule = append(reportedOracleRule, reportedOracleItem)
	}

	logs, sub, err := _Register.contract.FilterLogs(opts, "OracleReported", reporterRule, reportedOracleRule)
	if err != nil {
		return nil, err
	}
	return &RegisterOracleReportedIterator{contract: _Register.contract, event: "OracleReported", logs: logs, sub: sub}, nil
}

// WatchOracleReported is a free log subscription operation binding the contract event 0xc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d.
//
// Solidity: event OracleReported(address indexed reporter, address indexed reportedOracle)
func (_Register *RegisterFilterer) WatchOracleReported(opts *bind.WatchOpts, sink chan<- *RegisterOracleReported, reporter []common.Address, reportedOracle []common.Address) (event.Subscription, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var reportedOracleRule []interface{}
	for _, reportedOracleItem := range reportedOracle {
		reportedOracleRule = append(reportedOracleRule, reportedOracleItem)
	}

	logs, sub, err := _Register.contract.WatchLogs(opts, "OracleReported", reporterRule, reportedOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterOracleReported)
				if err := _Register.contract.UnpackLog(event, "OracleReported", log); err != nil {
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

// ParseOracleReported is a log parse operation binding the contract event 0xc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d.
//
// Solidity: event OracleReported(address indexed reporter, address indexed reportedOracle)
func (_Register *RegisterFilterer) ParseOracleReported(log types.Log) (*RegisterOracleReported, error) {
	event := new(RegisterOracleReported)
	if err := _Register.contract.UnpackLog(event, "OracleReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
