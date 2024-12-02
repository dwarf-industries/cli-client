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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_registerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reportFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daoOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blacklistedOracle\",\"type\":\"address\"}],\"name\":\"OracleBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTax\",\"type\":\"uint256\"}],\"name\":\"OracleRegisterTaxUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"OracleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reportedOracle\",\"type\":\"address\"}],\"name\":\"OracleReported\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tax\",\"type\":\"uint256\"}],\"name\":\"changeRegisterTax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"name\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"}],\"internalType\":\"structOracle[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistrationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOracleRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"reportOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateRegistrationFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateReportFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405260056002555f600560146101000a81548160ff02191690831515021790555034801561002e575f5ffd5b506040516126ca3803806126ca83398181016040528101906100509190610177565b8360038190555082600481905550815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506101db565b5f5ffd5b5f819050919050565b6100fc816100ea565b8114610106575f5ffd5b50565b5f81519050610117816100f3565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101468261011d565b9050919050565b6101568161013c565b8114610160575f5ffd5b50565b5f815190506101718161014d565b92915050565b5f5f5f5f6080858703121561018f5761018e6100e6565b5b5f61019c87828801610109565b94505060206101ad87828801610109565b93505060406101be87828801610163565b92505060606101cf87828801610163565b91505092959194509250565b6124e2806101e85f395ff3fe60806040526004361061009b575f3560e01c806369366c481161006357806369366c48146101775780637571bdd0146101a75780639a6bea18146101e3578063af582c6b1461020d578063e1704d831461023d578063fe575a871461026d5761009b565b80630946e8071461009f5780630d1db1f0146100c95780632b968958146100f35780633ffbd47f1461011d57806340884c521461014d575b5f5ffd5b3480156100aa575f5ffd5b506100b36102a9565b6040516100c091906116e2565b60405180910390f35b3480156100d4575f5ffd5b506100dd6102b2565b6040516100ea9190611715565b60405180910390f35b3480156100fe575f5ffd5b506101076104c8565b6040516101149190611715565b60405180910390f35b6101376004803603810190610132919061187b565b61057a565b6040516101449190611715565b60405180910390f35b348015610158575f5ffd5b506101616108c7565b60405161016e9190611ac1565b60405180910390f35b610191600480360381019061018c9190611b0b565b610aa1565b60405161019e9190611715565b60405180910390f35b3480156101b2575f5ffd5b506101cd60048036038101906101c89190611b0b565b610b41565b6040516101da9190611715565b60405180910390f35b3480156101ee575f5ffd5b506101f7610c18565b60405161020491906116e2565b60405180910390f35b61022760048036038101906102229190611b0b565b610c21565b6040516102349190611715565b60405180910390f35b61025760048036038101906102529190611b60565b610cc1565b6040516102649190611715565b60405180910390f35b348015610278575f5ffd5b50610293600480360381019061028e9190611b60565b6112e7565b6040516102a09190611715565b60405180910390f35b5f600354905090565b5f5f5f90505b6001805490508110156104c0575f600182815481106102da576102d9611b8b565b5b905f5260205f2090600402016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461035590611be5565b80601f016020809104026020016040519081016040528092919081815260200182805461038190611be5565b80156103cc5780601f106103a3576101008083540402835291602001916103cc565b820191905f5260205f20905b8154815290600101906020018083116103af57829003601f168201915b505050505081526020016002820180546103e590611be5565b80601f016020809104026020016040519081016040528092919081815260200182805461041190611be5565b801561045c5780601f106104335761010080835404028352916020019161045c565b820191905f5260205f20905b81548152906001019060200180831161043f57829003601f168201915b5050505050815260200160038201548152505090503373ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff16036104b2576001925050506104c5565b5080806001019150506102b8565b505f90505b90565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610558576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054f90611c6f565b60405180910390fd5b6001600560146101000a81548160ff0219169083151502179055506001905090565b5f6003543410156105c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b790611cd7565b60405180910390fd5b60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff161561064a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064190611d65565b60405180910390fd5b600160405180608001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020016001815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816107039190611f23565b5060408201518160020190816107199190611f23565b506060820151816003015550505f6064600254600354610739919061201f565b610743919061208d565b90505f813461075291906120bd565b90505f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318a9055c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107be573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107e29190612104565b90508073ffffffffffffffffffffffffffffffffffffffff166108fc8490811502906040515f60405180830381858888f19350505050158015610827573d5f5f3e3d5ffd5b5081600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff167f8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c987876040516108b2929190612167565b60405180910390a26001935050505092915050565b60606001805480602002602001604051908101604052809291908181526020015f905b82821015610a98578382905f5260205f2090600402016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461096f90611be5565b80601f016020809104026020016040519081016040528092919081815260200182805461099b90611be5565b80156109e65780601f106109bd576101008083540402835291602001916109e6565b820191905f5260205f20905b8154815290600101906020018083116109c957829003601f168201915b505050505081526020016002820180546109ff90611be5565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2b90611be5565b8015610a765780601f10610a4d57610100808354040283529160200191610a76565b820191905f5260205f20905b815481529060010190602001808311610a5957829003601f168201915b50505050508152602001600382015481525050815260200190600101906108ea565b50505050905090565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b31576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2890611c6f565b60405180910390fd5b8160048190555060019050919050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bd1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc890611c6f565b60405180910390fd5b816002819055507fa288a9ac9f74fe4c85e8263739645d49297e82ffa010aac11899ba671a54d1ac82604051610c0791906116e2565b60405180910390a160019050919050565b5f600454905090565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cb1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ca890611c6f565b60405180910390fd5b8160038190555060019050919050565b5f600560149054906101000a900460ff1615610d12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d099061220c565b60405180910390fd5b5f5f90505f5f90505b600180549050811015610e06573373ffffffffffffffffffffffffffffffffffffffff1660018281548110610d5357610d52611b8b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148015610deb575060085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16155b15610df95760019150610e06565b8080600101915050610d1b565b5080610e47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3e90612274565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610eb5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eac906122dc565b60405180910390fd5b60085f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610f3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3690612344565b60405180910390fd5b60065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615611004576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ffb906123ac565b60405180910390fd5b600454341015611049576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104090612414565b60405180910390fd5b600160065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060075f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f81548092919061112690612432565b91905055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d60405160405180910390a360016002600180549050611199919061208d565b6111a39190612479565b60075f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054106112dd57600160085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508273ffffffffffffffffffffffffffffffffffffffff167f8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf60405160405180910390a260045460095f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546112cc9190612479565b925050819055506112dc83611339565b5b6001915050919050565b5f60085f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f5f90505f5f90505b6001805490508110156114365760065f6001838154811061136657611365611b8b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff161561142957818061142590612432565b9250505b8080600101915050611342565b505f81600354611446919061208d565b90505f5f90505b6001805490508110156116c45760065f6001838154811061147157611470611b8b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156116b7576001818154811061153957611538611b8b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc8360095f6001868154811061159857611597611b8b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205461160a9190612479565b90811502906040515f60405180830381858888f19350505050158015611632573d5f5f3e3d5ffd5b505f60095f6001848154811061164b5761164a611b8b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505b808060010191505061144d565b50505050565b5f819050919050565b6116dc816116ca565b82525050565b5f6020820190506116f55f8301846116d3565b92915050565b5f8115159050919050565b61170f816116fb565b82525050565b5f6020820190506117285f830184611706565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61178d82611747565b810181811067ffffffffffffffff821117156117ac576117ab611757565b5b80604052505050565b5f6117be61172e565b90506117ca8282611784565b919050565b5f67ffffffffffffffff8211156117e9576117e8611757565b5b6117f282611747565b9050602081019050919050565b828183375f83830152505050565b5f61181f61181a846117cf565b6117b5565b90508281526020810184848401111561183b5761183a611743565b5b6118468482856117ff565b509392505050565b5f82601f8301126118625761186161173f565b5b813561187284826020860161180d565b91505092915050565b5f5f6040838503121561189157611890611737565b5b5f83013567ffffffffffffffff8111156118ae576118ad61173b565b5b6118ba8582860161184e565b925050602083013567ffffffffffffffff8111156118db576118da61173b565b5b6118e78582860161184e565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6119438261191a565b9050919050565b61195381611939565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61198b82611959565b6119958185611963565b93506119a5818560208601611973565b6119ae81611747565b840191505092915050565b6119c2816116ca565b82525050565b5f608083015f8301516119dd5f86018261194a565b50602083015184820360208601526119f58282611981565b91505060408301518482036040860152611a0f8282611981565b9150506060830151611a2460608601826119b9565b508091505092915050565b5f611a3a83836119c8565b905092915050565b5f602082019050919050565b5f611a58826118f1565b611a6281856118fb565b935083602082028501611a748561190b565b805f5b85811015611aaf5784840389528151611a908582611a2f565b9450611a9b83611a42565b925060208a01995050600181019050611a77565b50829750879550505050505092915050565b5f6020820190508181035f830152611ad98184611a4e565b905092915050565b611aea816116ca565b8114611af4575f5ffd5b50565b5f81359050611b0581611ae1565b92915050565b5f60208284031215611b2057611b1f611737565b5b5f611b2d84828501611af7565b91505092915050565b611b3f81611939565b8114611b49575f5ffd5b50565b5f81359050611b5a81611b36565b92915050565b5f60208284031215611b7557611b74611737565b5b5f611b8284828501611b4c565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611bfc57607f821691505b602082108103611c0f57611c0e611bb8565b5b50919050565b5f82825260208201905092915050565b7f4e6f7420617574686f72697a65640000000000000000000000000000000000005f82015250565b5f611c59600e83611c15565b9150611c6482611c25565b602082019050919050565b5f6020820190508181035f830152611c8681611c4d565b9050919050565b7f496e73756666696369656e7420726567697374726174696f6e206665650000005f82015250565b5f611cc1601d83611c15565b9150611ccc82611c8d565b602082019050919050565b5f6020820190508181035f830152611cee81611cb5565b9050919050565b7f43616e6e6f74207265676973746572206120626c61636b6c6973746564206f725f8201527f61636c6500000000000000000000000000000000000000000000000000000000602082015250565b5f611d4f602483611c15565b9150611d5a82611cf5565b604082019050919050565b5f6020820190508181035f830152611d7c81611d43565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611ddf7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611da4565b611de98683611da4565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611e24611e1f611e1a846116ca565b611e01565b6116ca565b9050919050565b5f819050919050565b611e3d83611e0a565b611e51611e4982611e2b565b848454611db0565b825550505050565b5f5f905090565b611e68611e59565b611e73818484611e34565b505050565b5b81811015611e9657611e8b5f82611e60565b600181019050611e79565b5050565b601f821115611edb57611eac81611d83565b611eb584611d95565b81016020851015611ec4578190505b611ed8611ed085611d95565b830182611e78565b50505b505050565b5f82821c905092915050565b5f611efb5f1984600802611ee0565b1980831691505092915050565b5f611f138383611eec565b9150826002028217905092915050565b611f2c82611959565b67ffffffffffffffff811115611f4557611f44611757565b5b611f4f8254611be5565b611f5a828285611e9a565b5f60209050601f831160018114611f8b575f8415611f79578287015190505b611f838582611f08565b865550611fea565b601f198416611f9986611d83565b5f5b82811015611fc057848901518255600182019150602085019450602081019050611f9b565b86831015611fdd5784890151611fd9601f891682611eec565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f612029826116ca565b9150612034836116ca565b9250828202612042816116ca565b9150828204841483151761205957612058611ff2565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f612097826116ca565b91506120a2836116ca565b9250826120b2576120b1612060565b5b828204905092915050565b5f6120c7826116ca565b91506120d2836116ca565b92508282039050818111156120ea576120e9611ff2565b5b92915050565b5f815190506120fe81611b36565b92915050565b5f6020828403121561211957612118611737565b5b5f612126848285016120f0565b91505092915050565b5f61213982611959565b6121438185611c15565b9350612153818560208601611973565b61215c81611747565b840191505092915050565b5f6040820190508181035f83015261217f818561212f565b90508181036020830152612193818461212f565b90509392505050565b7f4f776e657273686970206973207265766f6b65642c206e6f2061646d696e69735f8201527f74726174697665206368616e67657320616c6c6f776564210000000000000000602082015250565b5f6121f6603883611c15565b91506122018261219c565b604082019050919050565b5f6020820190508181035f830152612223816121ea565b9050919050565b7f4e6f7420616e20617574686f72697a6564206f7261636c6500000000000000005f82015250565b5f61225e601883611c15565b91506122698261222a565b602082019050919050565b5f6020820190508181035f83015261228b81612252565b9050919050565b7f43616e6e6f74207265706f727420796f757273656c66000000000000000000005f82015250565b5f6122c6601683611c15565b91506122d182612292565b602082019050919050565b5f6020820190508181035f8301526122f3816122ba565b9050919050565b7f4f7261636c6520697320616c726561647920626c61636b6c69737465640000005f82015250565b5f61232e601d83611c15565b9150612339826122fa565b602082019050919050565b5f6020820190508181035f83015261235b81612322565b9050919050565b7f4f7261636c6520616c7265616479207265706f7274656420627920796f7500005f82015250565b5f612396601e83611c15565b91506123a182612362565b602082019050919050565b5f6020820190508181035f8301526123c38161238a565b9050919050565b7f496e73756666696369656e74207265706f7274206665650000000000000000005f82015250565b5f6123fe601783611c15565b9150612409826123ca565b602082019050919050565b5f6020820190508181035f83015261242b816123f2565b9050919050565b5f61243c826116ca565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361246e5761246d611ff2565b5b600182019050919050565b5f612483826116ca565b915061248e836116ca565b92508282019050808211156124a6576124a5611ff2565b5b9291505056fea2646970667358221220a14de81e03656582a2d82d1330bdafe8f7cd12f634c254346cc01b17a0ea6b9664736f6c634300081c0033",
}

// RegisterABI is the input ABI used to generate the binding from.
// Deprecated: Use RegisterMetaData.ABI instead.
var RegisterABI = RegisterMetaData.ABI

// RegisterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegisterMetaData.Bin instead.
var RegisterBin = RegisterMetaData.Bin

// DeployRegister deploys a new Ethereum contract, binding an instance of Register to it.
func DeployRegister(auth *bind.TransactOpts, backend bind.ContractBackend, _registerFee *big.Int, _reportFee *big.Int, _feeSetter common.Address, daoOwner common.Address) (common.Address, *types.Transaction, *Register, error) {
	parsed, err := RegisterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegisterBin), backend, _registerFee, _reportFee, _feeSetter, daoOwner)
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

// ChangeRegisterTax is a paid mutator transaction binding the contract method 0x7571bdd0.
//
// Solidity: function changeRegisterTax(uint256 tax) returns(bool)
func (_Register *RegisterTransactor) ChangeRegisterTax(opts *bind.TransactOpts, tax *big.Int) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "changeRegisterTax", tax)
}

// ChangeRegisterTax is a paid mutator transaction binding the contract method 0x7571bdd0.
//
// Solidity: function changeRegisterTax(uint256 tax) returns(bool)
func (_Register *RegisterSession) ChangeRegisterTax(tax *big.Int) (*types.Transaction, error) {
	return _Register.Contract.ChangeRegisterTax(&_Register.TransactOpts, tax)
}

// ChangeRegisterTax is a paid mutator transaction binding the contract method 0x7571bdd0.
//
// Solidity: function changeRegisterTax(uint256 tax) returns(bool)
func (_Register *RegisterTransactorSession) ChangeRegisterTax(tax *big.Int) (*types.Transaction, error) {
	return _Register.Contract.ChangeRegisterTax(&_Register.TransactOpts, tax)
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

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns(bool)
func (_Register *RegisterTransactor) RevokeOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "revokeOwnership")
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns(bool)
func (_Register *RegisterSession) RevokeOwnership() (*types.Transaction, error) {
	return _Register.Contract.RevokeOwnership(&_Register.TransactOpts)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns(bool)
func (_Register *RegisterTransactorSession) RevokeOwnership() (*types.Transaction, error) {
	return _Register.Contract.RevokeOwnership(&_Register.TransactOpts)
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

// RegisterOracleRegisterTaxUpdatedIterator is returned from FilterOracleRegisterTaxUpdated and is used to iterate over the raw logs and unpacked data for OracleRegisterTaxUpdated events raised by the Register contract.
type RegisterOracleRegisterTaxUpdatedIterator struct {
	Event *RegisterOracleRegisterTaxUpdated // Event containing the contract specifics and raw log

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
func (it *RegisterOracleRegisterTaxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterOracleRegisterTaxUpdated)
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
		it.Event = new(RegisterOracleRegisterTaxUpdated)
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
func (it *RegisterOracleRegisterTaxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterOracleRegisterTaxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterOracleRegisterTaxUpdated represents a OracleRegisterTaxUpdated event raised by the Register contract.
type RegisterOracleRegisterTaxUpdated struct {
	NewTax *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOracleRegisterTaxUpdated is a free log retrieval operation binding the contract event 0xa288a9ac9f74fe4c85e8263739645d49297e82ffa010aac11899ba671a54d1ac.
//
// Solidity: event OracleRegisterTaxUpdated(uint256 newTax)
func (_Register *RegisterFilterer) FilterOracleRegisterTaxUpdated(opts *bind.FilterOpts) (*RegisterOracleRegisterTaxUpdatedIterator, error) {

	logs, sub, err := _Register.contract.FilterLogs(opts, "OracleRegisterTaxUpdated")
	if err != nil {
		return nil, err
	}
	return &RegisterOracleRegisterTaxUpdatedIterator{contract: _Register.contract, event: "OracleRegisterTaxUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleRegisterTaxUpdated is a free log subscription operation binding the contract event 0xa288a9ac9f74fe4c85e8263739645d49297e82ffa010aac11899ba671a54d1ac.
//
// Solidity: event OracleRegisterTaxUpdated(uint256 newTax)
func (_Register *RegisterFilterer) WatchOracleRegisterTaxUpdated(opts *bind.WatchOpts, sink chan<- *RegisterOracleRegisterTaxUpdated) (event.Subscription, error) {

	logs, sub, err := _Register.contract.WatchLogs(opts, "OracleRegisterTaxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterOracleRegisterTaxUpdated)
				if err := _Register.contract.UnpackLog(event, "OracleRegisterTaxUpdated", log); err != nil {
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

// ParseOracleRegisterTaxUpdated is a log parse operation binding the contract event 0xa288a9ac9f74fe4c85e8263739645d49297e82ffa010aac11899ba671a54d1ac.
//
// Solidity: event OracleRegisterTaxUpdated(uint256 newTax)
func (_Register *RegisterFilterer) ParseOracleRegisterTaxUpdated(log types.Log) (*RegisterOracleRegisterTaxUpdated, error) {
	event := new(RegisterOracleRegisterTaxUpdated)
	if err := _Register.contract.UnpackLog(event, "OracleRegisterTaxUpdated", log); err != nil {
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
