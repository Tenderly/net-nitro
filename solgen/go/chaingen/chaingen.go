// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chaingen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/tenderly/net-nitro/go-ethereum"
	"github.com/tenderly/net-nitro/go-ethereum/accounts/abi"
	"github.com/tenderly/net-nitro/go-ethereum/accounts/abi/bind"
	"github.com/tenderly/net-nitro/go-ethereum/common"
	"github.com/tenderly/net-nitro/go-ethereum/core/types"
	"github.com/tenderly/net-nitro/go-ethereum/event"
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

// CacheManagerEntry is an auto generated low-level Go binding around an user-defined struct.
type CacheManagerEntry struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}

// CacheManagerMetaData contains all meta data concerning the CacheManager contract.
var CacheManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"}],\"name\":\"AlreadyCached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"asm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queueSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cacheSize\",\"type\":\"uint256\"}],\"name\":\"AsmTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"name\":\"BidTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidsArePaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"MakeSpaceTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotChainOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"DeleteBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"InsertBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"SetCacheSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"decay\",\"type\":\"uint64\"}],\"name\":\"SetDecayRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cacheSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"entries\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"code\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"evictAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"evictPrograms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"code\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"}],\"internalType\":\"structCacheManager.Entry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"}],\"name\":\"getMinBid\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"}],\"name\":\"getMinBid\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"getMinBid\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"getSmallestEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"code\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"}],\"internalType\":\"structCacheManager.Entry[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initCacheSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initDecay\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"makeSpace\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"space\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSize\",\"type\":\"uint64\"}],\"name\":\"setCacheSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDecay\",\"type\":\"uint64\"}],\"name\":\"setDecayRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweepFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161206f61003060003960006105f3015261206f6000f3fe60806040526004361061015f5760003560e01c8063b187bd26116100c0578063c77ed13e11610074578063d29b303e11610059578063d29b303e146103ea578063e49401571461040a578063e9c1bc0f1461041d57600080fd5b8063c77ed13e146103aa578063cadb43e2146103ca57600080fd5b8063bae6c2ad116100a5578063bae6c2ad1461034a578063c1c013c414610377578063c565a2081461038a57600080fd5b8063b187bd26146102cc578063b30906d4146102fd57600080fd5b806354fac919116101175780635c975abb116100fc5780635c975abb14610281578063674a64e014610296578063a8d6fe04146102b757600080fd5b806354fac9191461021e5780635c32e9431461026c57600080fd5b80632dd4f566116101485780632dd4f566146101b157806332052a9b146101d15780633f4ba83a1461020957600080fd5b806317be85c31461016457806320f2f3451461018f575b600080fd5b34801561017057600080fd5b5061017961043d565b6040516101869190611daa565b60405180910390f35b34801561019b57600080fd5b506101af6101aa366004611e33565b6104d0565b005b3480156101bd57600080fd5b506101af6101cc366004611e66565b61073a565b3480156101dd57600080fd5b506101f16101ec366004611ea3565b610810565b6040516001600160c01b039091168152602001610186565b34801561021557600080fd5b506101af610838565b34801561022a57600080fd5b5060035461025390700100000000000000000000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610186565b34801561027857600080fd5b506101af61090a565b34801561028d57600080fd5b506101af6109a2565b3480156102a257600080fd5b506003546102539067ffffffffffffffff1681565b3480156102c357600080fd5b506101af610a7a565b3480156102d857600080fd5b506003546102ed90600160c01b900460ff1681565b6040519015158152602001610186565b34801561030957600080fd5b5061031d610318366004611ec0565b610b5f565b6040805193845267ffffffffffffffff90921660208401526001600160c01b031690820152606001610186565b34801561035657600080fd5b506003546102539068010000000000000000900467ffffffffffffffff1681565b610253610385366004611e66565b610bae565b34801561039657600080fd5b506101f16103a5366004611ec0565b610c81565b3480156103b657600080fd5b506101af6103c5366004611e66565b610c8f565b3480156103d657600080fd5b506101af6103e5366004611ec0565b610d8b565b3480156103f657600080fd5b506101f1610405366004611e66565b610e5d565b6101af610418366004611ea3565b611068565b34801561042957600080fd5b50610179610438366004611ec0565b611138565b60606002805480602002602001604051908101604052809291908181526020016000905b828210156104c757600084815260209081902060408051606081018252600286029092018054835260019081015467ffffffffffffffff8116848601526801000000000000000090046001600160c01b0316918301919091529083529092019101610461565b50505050905090565b600054610100900460ff16158080156104f05750600054600160ff909116105b8061050a5750303b15801561050a575060005460ff166001145b61059b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156105dc576000805461ff0019166101001790555b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001630036106a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c00000000000000000000000000000000000000006064820152608401610592565b6003805467ffffffffffffffff848116700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffff0000000000000000909216908616171790558015610735576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6040516304ddefed60e31b8152336004820152606b906326ef7f6890602401602060405180830381865afa158015610776573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079a9190611ed9565b6107b957604051639531eff160e01b8152336004820152602401610592565b6003805467ffffffffffffffff191667ffffffffffffffff83169081179091556040519081527fca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839906020015b60405180910390a150565b60006108328273ffffffffffffffffffffffffffffffffffffffff163f610c81565b92915050565b6040516304ddefed60e31b8152336004820152606b906326ef7f6890602401602060405180830381865afa158015610874573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108989190611ed9565b6108b757604051639531eff160e01b8152336004820152602401610592565b600380547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff1690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b6040516304ddefed60e31b8152336004820152606b906326ef7f6890602401602060405180830381865afa158015610946573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096a9190611ed9565b61098957604051639531eff160e01b8152336004820152602401610592565b610994600019610d8b565b6109a060026000611d6f565b565b6040516304ddefed60e31b8152336004820152606b906326ef7f6890602401602060405180830381865afa1580156109de573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a029190611ed9565b610a2157604051639531eff160e01b8152336004820152602401610592565b600380547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff16600160c01b1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600080606b73ffffffffffffffffffffffffffffffffffffffff16632d9125e96040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ac9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aed9190611efb565b73ffffffffffffffffffffffffffffffffffffffff164760405160006040518083038185875af1925050503d8060008114610b44576040519150601f19603f3d011682016040523d82523d6000602084013e610b49565b606091505b509150915081610b5b57805160208201fd5b5050565b60028181548110610b6f57600080fd5b60009182526020909120600290910201805460019091015490915067ffffffffffffffff8116906801000000000000000090046001600160c01b031683565b600354600090600160c01b900460ff1615610bf5576040517f8f55c96c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6250000067ffffffffffffffff83161115610c51576040517fe6b801f300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152625000006024820152604401610592565b610c5a826112a9565b50506003546108329067ffffffffffffffff68010000000000000000820481169116611f2e565b60006108326104058361138c565b6040516304ddefed60e31b8152336004820152606b906326ef7f6890602401602060405180830381865afa158015610ccb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cef9190611ed9565b610d0e57604051639531eff160e01b8152336004820152602401610592565b600380547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c90602001610805565b6040516304ddefed60e31b8152336004820152606b906326ef7f6890602401602060405180830381865afa158015610dc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610deb9190611ed9565b610e0a57604051639531eff160e01b8152336004820152602401610592565b60015415801590610e1b5750600081115b15610e5a57600080610e38610e306001611434565b604081901c91565b91509150610e468282611444565b610e51600184611f56565b92505050610e0a565b50565b60035460009067ffffffffffffffff9081169083161115610ec9576003546040517fbcc27c3700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8085166004830152600060248301529091166044820152606401610592565b61100067ffffffffffffffff83161015610ee557611000610ee7565b815b600354909250600090610f1190849068010000000000000000900467ffffffffffffffff16611f69565b60035467ffffffffffffffff9182169250168111610f325750600092915050565b600354600090610f4c9067ffffffffffffffff1683611f56565b905060006110006001610f5f8285611f8a565b610f699190611f56565b610f739190611f9d565b90506000610f8082611138565b905060005b815181101561102857818181518110610fa057610fa0611fbf565b60200260200101516020015167ffffffffffffffff168411610fe157818181518110610fce57610fce611fbf565b6020026020010151604001519550611028565b818181518110610ff357610ff3611fbf565b60200260200101516020015167ffffffffffffffff16846110149190611f56565b93508061102081611fd5565b915050610f85565b50600061103361160b565b905080866001600160c01b03161015611053575060009695505050505050565b61105d8187611fef565b979650505050505050565b600354600160c01b900460ff16156110ac576040517f8f55c96c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81163f6110cd8161163e565b15611107576040517fc7e2d8e500000000000000000000000000000000000000000000000000000000815260048101829052602401610592565b60006111128261138c565b9050600080611120836112a9565b9150915061113182868686856116bb565b5050505050565b60608161114460015490565b10156111505760015491505b600061115d6001846119ed565b9050805167ffffffffffffffff8111156111795761117961200f565b6040519080825280602002602001820160405280156111c457816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816111975790505b50915060005b81518110156112a25760006111fa8383815181106111ea576111ea611fbf565b6020026020010151604081901c91565b91505060028167ffffffffffffffff168154811061121a5761121a611fbf565b600091825260209182902060408051606081018252600293909302909101805483526001015467ffffffffffffffff811693830193909352680100000000000000009092046001600160c01b031691810191909152845185908490811061128357611283611fbf565b602002602001018190525050808061129a90611fd5565b9150506111ca565b5050919050565b6000806112b534611bab565b600254600354919350915060009067ffffffffffffffff165b60035467ffffffffffffffff808316916112f691889168010000000000000000900416611f69565b67ffffffffffffffff16111561132657611313610e306001611434565b935091506113218284611444565b6112ce565b816001600160c01b0316846001600160c01b03161015611385576040517fdf370e480000000000000000000000000000000000000000000000000000000081526001600160c01b03808616600483015283166024820152604401610592565b5050915091565b6040517f4089267f000000000000000000000000000000000000000000000000000000008152600481018290526000908190607190634089267f90602401602060405180830381865afa1580156113e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061140b9190612025565b905061100063ffffffff821610156114255761100061142d565b8063ffffffff165b9392505050565b600061142d826000806002611c06565b600060028267ffffffffffffffff168154811061146357611463611fbf565b600091825260209182902060408051606081018252600293909302909101805480845260019091015467ffffffffffffffff811694840194909452680100000000000000009093046001600160c01b031682820152517fce9720130000000000000000000000000000000000000000000000000000000081526004810192909252915060729063ce97201390602401600060405180830381600087803b15801561150c57600080fd5b505af1158015611520573d6000803e3d6000fd5b505050508060200151600360088282829054906101000a900467ffffffffffffffff1661154d9190611f2e565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080600001517f65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e78483602001516040516115cb9291906001600160c01b0392909216825267ffffffffffffffff16602082015260400190565b60405180910390a260028267ffffffffffffffff16815481106115f0576115f0611fbf565b60009182526020822060029091020181815560010155505050565b60035460009061163990700100000000000000000000000000000000900467ffffffffffffffff164261204b565b905090565b6040517fa72f179b0000000000000000000000000000000000000000000000000000000081526004810182905260009060729063a72f179b90602401602060405180830381865afa158015611697573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108329190611ed9565b60035467ffffffffffffffff808216916116e391859168010000000000000000900416611f69565b67ffffffffffffffff161115611750576003546040517fbcc27c3700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8085166004830152680100000000000000008304811660248301529091166044820152606401610592565b6040805160608101825284815267ffffffffffffffff841660208201526001600160c01b0387168183015290517fe73ac9f200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260729063e73ac9f290602401600060405180830381600087803b1580156117e257600080fd5b505af11580156117f6573d6000803e3d6000fd5b506118229250505067ffffffffffffffff19604088901b1667ffffffffffffffff841617600190611d61565b82600360088282829054906101000a900467ffffffffffffffff166118479190611f69565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506002805490508267ffffffffffffffff160361191257600280546001810182556000829052825191027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace810191909155602082015160408301516001600160c01b0316680100000000000000000267ffffffffffffffff909116177f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf9091015561197b565b8060028367ffffffffffffffff168154811061193057611930611fbf565b60009182526020918290208351600290920201908155908201516040909201516001600160c01b0316680100000000000000000267ffffffffffffffff909216919091176001909101555b6040805173ffffffffffffffffffffffffffffffffffffffff871681526001600160c01b038816602082015267ffffffffffffffff851681830152905185917fb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec919081900360600190a2505050505050565b6060611aad565b6001820360011c611a0b818360069190911b015190565b85108302611a195750611a3a565b600681811b830180516020918201519286901b8501908152015291506119f4565b600682901b81018481526020018390525b50505050565b600060015b83811015611aa0576001810184118101600690811b84015182821b8501511190910180821b840180516020918201519385901b860190815201919091529050600181811b01611a56565b50611131858583856119f4565b5060408051600084815260208082208654815487821882891002821860051b86018085019182529096019390935292939084019190808314155b8015611b74578151845283602001935081840315611b7457602082015160019060011b01838110611b4357600182039150611b3d611b2b838560069190911b015190565b600684901b8501602001518486611a51565b50611ae7565b611b5281870154828486611a51565b600101838114611b3d57611b6b818701548284866119f4565b50600101611ae7565b5050507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08382030160051c83526040525092915050565b600080611bb661160b565b611bc09084611f8a565b90506001600160c01b03811115610832576040517ff6e86d2800000000000000000000000000000000000000000000000000000000815260048101829052602401610592565b60008085548660005260206000206000801986611c5a578715611ccf57878414611c3c5750506001828101895593508181611cdd565b8254898110611c4b5750611cdd565b60039650945060019050611cdd565b60028711611c95578315611ccf5760028703611c8957600184039350838a558383015498508894508315611cdd575b50815493506001611cdd565b60038703611cac5750506001820188558181611cdd565b8894508315611cdd578254898110611cc45750611cdd565b945060019050611cdd565b63a6ca772e6000526004601cfd5b5b83811015611d175780830154600182018085015480831087831011611d035750829050815b938501939093555050600181811b01611cde565b8115611d45576001820360011c935083830154808a10611d375750611d45565b808385015550839150611d17565b6001810115611d545788828401555b5050505094509492505050565b611a4b828260006003611c06565b5080546000825560020290600052602060002090810190610e5a91905b80821115611da65760008082556001820155600201611d8c565b5090565b602080825282518282018190526000919060409081850190868401855b82811015611e09578151805185528681015167ffffffffffffffff16878601528501516001600160c01b03168585015260609093019290850190600101611dc7565b5091979650505050505050565b803567ffffffffffffffff81168114611e2e57600080fd5b919050565b60008060408385031215611e4657600080fd5b611e4f83611e16565b9150611e5d60208401611e16565b90509250929050565b600060208284031215611e7857600080fd5b61142d82611e16565b73ffffffffffffffffffffffffffffffffffffffff81168114610e5a57600080fd5b600060208284031215611eb557600080fd5b813561142d81611e81565b600060208284031215611ed257600080fd5b5035919050565b600060208284031215611eeb57600080fd5b8151801515811461142d57600080fd5b600060208284031215611f0d57600080fd5b815161142d81611e81565b634e487b7160e01b600052601160045260246000fd5b67ffffffffffffffff828116828216039080821115611f4f57611f4f611f18565b5092915050565b8181038181111561083257610832611f18565b67ffffffffffffffff818116838216019080821115611f4f57611f4f611f18565b8082018082111561083257610832611f18565b600082611fba57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052603260045260246000fd5b60006000198203611fe857611fe8611f18565b5060010190565b6001600160c01b03828116828216039080821115611f4f57611f4f611f18565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561203757600080fd5b815163ffffffff8116811461142d57600080fd5b808202811582820484141761083257610832611f1856fea164736f6c6343000811000a",
}

// CacheManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CacheManagerMetaData.ABI instead.
var CacheManagerABI = CacheManagerMetaData.ABI

// CacheManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CacheManagerMetaData.Bin instead.
var CacheManagerBin = CacheManagerMetaData.Bin

// DeployCacheManager deploys a new Ethereum contract, binding an instance of CacheManager to it.
func DeployCacheManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CacheManager, error) {
	parsed, err := CacheManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CacheManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CacheManager{CacheManagerCaller: CacheManagerCaller{contract: contract}, CacheManagerTransactor: CacheManagerTransactor{contract: contract}, CacheManagerFilterer: CacheManagerFilterer{contract: contract}}, nil
}

// CacheManager is an auto generated Go binding around an Ethereum contract.
type CacheManager struct {
	CacheManagerCaller     // Read-only binding to the contract
	CacheManagerTransactor // Write-only binding to the contract
	CacheManagerFilterer   // Log filterer for contract events
}

// CacheManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CacheManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CacheManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CacheManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CacheManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CacheManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CacheManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CacheManagerSession struct {
	Contract     *CacheManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CacheManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CacheManagerCallerSession struct {
	Contract *CacheManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CacheManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CacheManagerTransactorSession struct {
	Contract     *CacheManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CacheManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CacheManagerRaw struct {
	Contract *CacheManager // Generic contract binding to access the raw methods on
}

// CacheManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CacheManagerCallerRaw struct {
	Contract *CacheManagerCaller // Generic read-only contract binding to access the raw methods on
}

// CacheManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CacheManagerTransactorRaw struct {
	Contract *CacheManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCacheManager creates a new instance of CacheManager, bound to a specific deployed contract.
func NewCacheManager(address common.Address, backend bind.ContractBackend) (*CacheManager, error) {
	contract, err := bindCacheManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CacheManager{CacheManagerCaller: CacheManagerCaller{contract: contract}, CacheManagerTransactor: CacheManagerTransactor{contract: contract}, CacheManagerFilterer: CacheManagerFilterer{contract: contract}}, nil
}

// NewCacheManagerCaller creates a new read-only instance of CacheManager, bound to a specific deployed contract.
func NewCacheManagerCaller(address common.Address, caller bind.ContractCaller) (*CacheManagerCaller, error) {
	contract, err := bindCacheManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CacheManagerCaller{contract: contract}, nil
}

// NewCacheManagerTransactor creates a new write-only instance of CacheManager, bound to a specific deployed contract.
func NewCacheManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CacheManagerTransactor, error) {
	contract, err := bindCacheManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CacheManagerTransactor{contract: contract}, nil
}

// NewCacheManagerFilterer creates a new log filterer instance of CacheManager, bound to a specific deployed contract.
func NewCacheManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CacheManagerFilterer, error) {
	contract, err := bindCacheManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CacheManagerFilterer{contract: contract}, nil
}

// bindCacheManager binds a generic wrapper to an already deployed contract.
func bindCacheManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CacheManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CacheManager *CacheManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CacheManager.Contract.CacheManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CacheManager *CacheManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.Contract.CacheManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CacheManager *CacheManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CacheManager.Contract.CacheManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CacheManager *CacheManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CacheManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CacheManager *CacheManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CacheManager *CacheManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CacheManager.Contract.contract.Transact(opts, method, params...)
}

// CacheSize is a free data retrieval call binding the contract method 0x674a64e0.
//
// Solidity: function cacheSize() view returns(uint64)
func (_CacheManager *CacheManagerCaller) CacheSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "cacheSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CacheSize is a free data retrieval call binding the contract method 0x674a64e0.
//
// Solidity: function cacheSize() view returns(uint64)
func (_CacheManager *CacheManagerSession) CacheSize() (uint64, error) {
	return _CacheManager.Contract.CacheSize(&_CacheManager.CallOpts)
}

// CacheSize is a free data retrieval call binding the contract method 0x674a64e0.
//
// Solidity: function cacheSize() view returns(uint64)
func (_CacheManager *CacheManagerCallerSession) CacheSize() (uint64, error) {
	return _CacheManager.Contract.CacheSize(&_CacheManager.CallOpts)
}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint64)
func (_CacheManager *CacheManagerCaller) Decay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "decay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint64)
func (_CacheManager *CacheManagerSession) Decay() (uint64, error) {
	return _CacheManager.Contract.Decay(&_CacheManager.CallOpts)
}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint64)
func (_CacheManager *CacheManagerCallerSession) Decay() (uint64, error) {
	return _CacheManager.Contract.Decay(&_CacheManager.CallOpts)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size, uint192 bid)
func (_CacheManager *CacheManagerCaller) Entries(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "entries", arg0)

	outstruct := new(struct {
		Code [32]byte
		Size uint64
		Bid  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Code = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Size = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Bid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size, uint192 bid)
func (_CacheManager *CacheManagerSession) Entries(arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}, error) {
	return _CacheManager.Contract.Entries(&_CacheManager.CallOpts, arg0)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size, uint192 bid)
func (_CacheManager *CacheManagerCallerSession) Entries(arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}, error) {
	return _CacheManager.Contract.Entries(&_CacheManager.CallOpts, arg0)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes32,uint64,uint192)[])
func (_CacheManager *CacheManagerCaller) GetEntries(opts *bind.CallOpts) ([]CacheManagerEntry, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getEntries")

	if err != nil {
		return *new([]CacheManagerEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]CacheManagerEntry)).(*[]CacheManagerEntry)

	return out0, err

}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes32,uint64,uint192)[])
func (_CacheManager *CacheManagerSession) GetEntries() ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetEntries(&_CacheManager.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes32,uint64,uint192)[])
func (_CacheManager *CacheManagerCallerSession) GetEntries() ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetEntries(&_CacheManager.CallOpts)
}

// GetMinBid32052a9b is a free data retrieval call binding the contract method 0x32052a9b.
//
// Solidity: function getMinBid(address program) view returns(uint192 min)
func (_CacheManager *CacheManagerCaller) GetMinBid32052a9b(opts *bind.CallOpts, program common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getMinBid", program)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBid32052a9b is a free data retrieval call binding the contract method 0x32052a9b.
//
// Solidity: function getMinBid(address program) view returns(uint192 min)
func (_CacheManager *CacheManagerSession) GetMinBid32052a9b(program common.Address) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBid32052a9b(&_CacheManager.CallOpts, program)
}

// GetMinBid32052a9b is a free data retrieval call binding the contract method 0x32052a9b.
//
// Solidity: function getMinBid(address program) view returns(uint192 min)
func (_CacheManager *CacheManagerCallerSession) GetMinBid32052a9b(program common.Address) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBid32052a9b(&_CacheManager.CallOpts, program)
}

// GetMinBidc565a208 is a free data retrieval call binding the contract method 0xc565a208.
//
// Solidity: function getMinBid(bytes32 codehash) view returns(uint192 min)
func (_CacheManager *CacheManagerCaller) GetMinBidc565a208(opts *bind.CallOpts, codehash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getMinBid0", codehash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBidc565a208 is a free data retrieval call binding the contract method 0xc565a208.
//
// Solidity: function getMinBid(bytes32 codehash) view returns(uint192 min)
func (_CacheManager *CacheManagerSession) GetMinBidc565a208(codehash [32]byte) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidc565a208(&_CacheManager.CallOpts, codehash)
}

// GetMinBidc565a208 is a free data retrieval call binding the contract method 0xc565a208.
//
// Solidity: function getMinBid(bytes32 codehash) view returns(uint192 min)
func (_CacheManager *CacheManagerCallerSession) GetMinBidc565a208(codehash [32]byte) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidc565a208(&_CacheManager.CallOpts, codehash)
}

// GetMinBidd29b303e is a free data retrieval call binding the contract method 0xd29b303e.
//
// Solidity: function getMinBid(uint64 size) view returns(uint192 min)
func (_CacheManager *CacheManagerCaller) GetMinBidd29b303e(opts *bind.CallOpts, size uint64) (*big.Int, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getMinBid1", size)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBidd29b303e is a free data retrieval call binding the contract method 0xd29b303e.
//
// Solidity: function getMinBid(uint64 size) view returns(uint192 min)
func (_CacheManager *CacheManagerSession) GetMinBidd29b303e(size uint64) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidd29b303e(&_CacheManager.CallOpts, size)
}

// GetMinBidd29b303e is a free data retrieval call binding the contract method 0xd29b303e.
//
// Solidity: function getMinBid(uint64 size) view returns(uint192 min)
func (_CacheManager *CacheManagerCallerSession) GetMinBidd29b303e(size uint64) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidd29b303e(&_CacheManager.CallOpts, size)
}

// GetSmallestEntries is a free data retrieval call binding the contract method 0xe9c1bc0f.
//
// Solidity: function getSmallestEntries(uint256 k) view returns((bytes32,uint64,uint192)[] result)
func (_CacheManager *CacheManagerCaller) GetSmallestEntries(opts *bind.CallOpts, k *big.Int) ([]CacheManagerEntry, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getSmallestEntries", k)

	if err != nil {
		return *new([]CacheManagerEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]CacheManagerEntry)).(*[]CacheManagerEntry)

	return out0, err

}

// GetSmallestEntries is a free data retrieval call binding the contract method 0xe9c1bc0f.
//
// Solidity: function getSmallestEntries(uint256 k) view returns((bytes32,uint64,uint192)[] result)
func (_CacheManager *CacheManagerSession) GetSmallestEntries(k *big.Int) ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetSmallestEntries(&_CacheManager.CallOpts, k)
}

// GetSmallestEntries is a free data retrieval call binding the contract method 0xe9c1bc0f.
//
// Solidity: function getSmallestEntries(uint256 k) view returns((bytes32,uint64,uint192)[] result)
func (_CacheManager *CacheManagerCallerSession) GetSmallestEntries(k *big.Int) ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetSmallestEntries(&_CacheManager.CallOpts, k)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CacheManager *CacheManagerCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CacheManager *CacheManagerSession) IsPaused() (bool, error) {
	return _CacheManager.Contract.IsPaused(&_CacheManager.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CacheManager *CacheManagerCallerSession) IsPaused() (bool, error) {
	return _CacheManager.Contract.IsPaused(&_CacheManager.CallOpts)
}

// QueueSize is a free data retrieval call binding the contract method 0xbae6c2ad.
//
// Solidity: function queueSize() view returns(uint64)
func (_CacheManager *CacheManagerCaller) QueueSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "queueSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// QueueSize is a free data retrieval call binding the contract method 0xbae6c2ad.
//
// Solidity: function queueSize() view returns(uint64)
func (_CacheManager *CacheManagerSession) QueueSize() (uint64, error) {
	return _CacheManager.Contract.QueueSize(&_CacheManager.CallOpts)
}

// QueueSize is a free data retrieval call binding the contract method 0xbae6c2ad.
//
// Solidity: function queueSize() view returns(uint64)
func (_CacheManager *CacheManagerCallerSession) QueueSize() (uint64, error) {
	return _CacheManager.Contract.QueueSize(&_CacheManager.CallOpts)
}

// EvictAll is a paid mutator transaction binding the contract method 0x5c32e943.
//
// Solidity: function evictAll() returns()
func (_CacheManager *CacheManagerTransactor) EvictAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "evictAll")
}

// EvictAll is a paid mutator transaction binding the contract method 0x5c32e943.
//
// Solidity: function evictAll() returns()
func (_CacheManager *CacheManagerSession) EvictAll() (*types.Transaction, error) {
	return _CacheManager.Contract.EvictAll(&_CacheManager.TransactOpts)
}

// EvictAll is a paid mutator transaction binding the contract method 0x5c32e943.
//
// Solidity: function evictAll() returns()
func (_CacheManager *CacheManagerTransactorSession) EvictAll() (*types.Transaction, error) {
	return _CacheManager.Contract.EvictAll(&_CacheManager.TransactOpts)
}

// EvictPrograms is a paid mutator transaction binding the contract method 0xcadb43e2.
//
// Solidity: function evictPrograms(uint256 count) returns()
func (_CacheManager *CacheManagerTransactor) EvictPrograms(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "evictPrograms", count)
}

// EvictPrograms is a paid mutator transaction binding the contract method 0xcadb43e2.
//
// Solidity: function evictPrograms(uint256 count) returns()
func (_CacheManager *CacheManagerSession) EvictPrograms(count *big.Int) (*types.Transaction, error) {
	return _CacheManager.Contract.EvictPrograms(&_CacheManager.TransactOpts, count)
}

// EvictPrograms is a paid mutator transaction binding the contract method 0xcadb43e2.
//
// Solidity: function evictPrograms(uint256 count) returns()
func (_CacheManager *CacheManagerTransactorSession) EvictPrograms(count *big.Int) (*types.Transaction, error) {
	return _CacheManager.Contract.EvictPrograms(&_CacheManager.TransactOpts, count)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 initCacheSize, uint64 initDecay) returns()
func (_CacheManager *CacheManagerTransactor) Initialize(opts *bind.TransactOpts, initCacheSize uint64, initDecay uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "initialize", initCacheSize, initDecay)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 initCacheSize, uint64 initDecay) returns()
func (_CacheManager *CacheManagerSession) Initialize(initCacheSize uint64, initDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.Initialize(&_CacheManager.TransactOpts, initCacheSize, initDecay)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 initCacheSize, uint64 initDecay) returns()
func (_CacheManager *CacheManagerTransactorSession) Initialize(initCacheSize uint64, initDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.Initialize(&_CacheManager.TransactOpts, initCacheSize, initDecay)
}

// MakeSpace is a paid mutator transaction binding the contract method 0xc1c013c4.
//
// Solidity: function makeSpace(uint64 size) payable returns(uint64 space)
func (_CacheManager *CacheManagerTransactor) MakeSpace(opts *bind.TransactOpts, size uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "makeSpace", size)
}

// MakeSpace is a paid mutator transaction binding the contract method 0xc1c013c4.
//
// Solidity: function makeSpace(uint64 size) payable returns(uint64 space)
func (_CacheManager *CacheManagerSession) MakeSpace(size uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.MakeSpace(&_CacheManager.TransactOpts, size)
}

// MakeSpace is a paid mutator transaction binding the contract method 0xc1c013c4.
//
// Solidity: function makeSpace(uint64 size) payable returns(uint64 space)
func (_CacheManager *CacheManagerTransactorSession) MakeSpace(size uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.MakeSpace(&_CacheManager.TransactOpts, size)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns()
func (_CacheManager *CacheManagerTransactor) Paused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "paused")
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns()
func (_CacheManager *CacheManagerSession) Paused() (*types.Transaction, error) {
	return _CacheManager.Contract.Paused(&_CacheManager.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns()
func (_CacheManager *CacheManagerTransactorSession) Paused() (*types.Transaction, error) {
	return _CacheManager.Contract.Paused(&_CacheManager.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xe4940157.
//
// Solidity: function placeBid(address program) payable returns()
func (_CacheManager *CacheManagerTransactor) PlaceBid(opts *bind.TransactOpts, program common.Address) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "placeBid", program)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xe4940157.
//
// Solidity: function placeBid(address program) payable returns()
func (_CacheManager *CacheManagerSession) PlaceBid(program common.Address) (*types.Transaction, error) {
	return _CacheManager.Contract.PlaceBid(&_CacheManager.TransactOpts, program)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xe4940157.
//
// Solidity: function placeBid(address program) payable returns()
func (_CacheManager *CacheManagerTransactorSession) PlaceBid(program common.Address) (*types.Transaction, error) {
	return _CacheManager.Contract.PlaceBid(&_CacheManager.TransactOpts, program)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x2dd4f566.
//
// Solidity: function setCacheSize(uint64 newSize) returns()
func (_CacheManager *CacheManagerTransactor) SetCacheSize(opts *bind.TransactOpts, newSize uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "setCacheSize", newSize)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x2dd4f566.
//
// Solidity: function setCacheSize(uint64 newSize) returns()
func (_CacheManager *CacheManagerSession) SetCacheSize(newSize uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetCacheSize(&_CacheManager.TransactOpts, newSize)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x2dd4f566.
//
// Solidity: function setCacheSize(uint64 newSize) returns()
func (_CacheManager *CacheManagerTransactorSession) SetCacheSize(newSize uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetCacheSize(&_CacheManager.TransactOpts, newSize)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xc77ed13e.
//
// Solidity: function setDecayRate(uint64 newDecay) returns()
func (_CacheManager *CacheManagerTransactor) SetDecayRate(opts *bind.TransactOpts, newDecay uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "setDecayRate", newDecay)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xc77ed13e.
//
// Solidity: function setDecayRate(uint64 newDecay) returns()
func (_CacheManager *CacheManagerSession) SetDecayRate(newDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetDecayRate(&_CacheManager.TransactOpts, newDecay)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xc77ed13e.
//
// Solidity: function setDecayRate(uint64 newDecay) returns()
func (_CacheManager *CacheManagerTransactorSession) SetDecayRate(newDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetDecayRate(&_CacheManager.TransactOpts, newDecay)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns()
func (_CacheManager *CacheManagerTransactor) SweepFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "sweepFunds")
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns()
func (_CacheManager *CacheManagerSession) SweepFunds() (*types.Transaction, error) {
	return _CacheManager.Contract.SweepFunds(&_CacheManager.TransactOpts)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns()
func (_CacheManager *CacheManagerTransactorSession) SweepFunds() (*types.Transaction, error) {
	return _CacheManager.Contract.SweepFunds(&_CacheManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CacheManager *CacheManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CacheManager *CacheManagerSession) Unpause() (*types.Transaction, error) {
	return _CacheManager.Contract.Unpause(&_CacheManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CacheManager *CacheManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _CacheManager.Contract.Unpause(&_CacheManager.TransactOpts)
}

// CacheManagerDeleteBidIterator is returned from FilterDeleteBid and is used to iterate over the raw logs and unpacked data for DeleteBid events raised by the CacheManager contract.
type CacheManagerDeleteBidIterator struct {
	Event *CacheManagerDeleteBid // Event containing the contract specifics and raw log

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
func (it *CacheManagerDeleteBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerDeleteBid)
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
		it.Event = new(CacheManagerDeleteBid)
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
func (it *CacheManagerDeleteBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerDeleteBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerDeleteBid represents a DeleteBid event raised by the CacheManager contract.
type CacheManagerDeleteBid struct {
	Codehash [32]byte
	Bid      *big.Int
	Size     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteBid is a free log retrieval operation binding the contract event 0x65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7.
//
// Solidity: event DeleteBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) FilterDeleteBid(opts *bind.FilterOpts, codehash [][32]byte) (*CacheManagerDeleteBidIterator, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "DeleteBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return &CacheManagerDeleteBidIterator{contract: _CacheManager.contract, event: "DeleteBid", logs: logs, sub: sub}, nil
}

// WatchDeleteBid is a free log subscription operation binding the contract event 0x65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7.
//
// Solidity: event DeleteBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) WatchDeleteBid(opts *bind.WatchOpts, sink chan<- *CacheManagerDeleteBid, codehash [][32]byte) (event.Subscription, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "DeleteBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerDeleteBid)
				if err := _CacheManager.contract.UnpackLog(event, "DeleteBid", log); err != nil {
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

// ParseDeleteBid is a log parse operation binding the contract event 0x65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7.
//
// Solidity: event DeleteBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) ParseDeleteBid(log types.Log) (*CacheManagerDeleteBid, error) {
	event := new(CacheManagerDeleteBid)
	if err := _CacheManager.contract.UnpackLog(event, "DeleteBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CacheManager contract.
type CacheManagerInitializedIterator struct {
	Event *CacheManagerInitialized // Event containing the contract specifics and raw log

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
func (it *CacheManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerInitialized)
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
		it.Event = new(CacheManagerInitialized)
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
func (it *CacheManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerInitialized represents a Initialized event raised by the CacheManager contract.
type CacheManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CacheManager *CacheManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*CacheManagerInitializedIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CacheManagerInitializedIterator{contract: _CacheManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CacheManager *CacheManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CacheManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerInitialized)
				if err := _CacheManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CacheManager *CacheManagerFilterer) ParseInitialized(log types.Log) (*CacheManagerInitialized, error) {
	event := new(CacheManagerInitialized)
	if err := _CacheManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerInsertBidIterator is returned from FilterInsertBid and is used to iterate over the raw logs and unpacked data for InsertBid events raised by the CacheManager contract.
type CacheManagerInsertBidIterator struct {
	Event *CacheManagerInsertBid // Event containing the contract specifics and raw log

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
func (it *CacheManagerInsertBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerInsertBid)
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
		it.Event = new(CacheManagerInsertBid)
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
func (it *CacheManagerInsertBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerInsertBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerInsertBid represents a InsertBid event raised by the CacheManager contract.
type CacheManagerInsertBid struct {
	Codehash [32]byte
	Program  common.Address
	Bid      *big.Int
	Size     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInsertBid is a free log retrieval operation binding the contract event 0xb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec.
//
// Solidity: event InsertBid(bytes32 indexed codehash, address program, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) FilterInsertBid(opts *bind.FilterOpts, codehash [][32]byte) (*CacheManagerInsertBidIterator, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "InsertBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return &CacheManagerInsertBidIterator{contract: _CacheManager.contract, event: "InsertBid", logs: logs, sub: sub}, nil
}

// WatchInsertBid is a free log subscription operation binding the contract event 0xb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec.
//
// Solidity: event InsertBid(bytes32 indexed codehash, address program, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) WatchInsertBid(opts *bind.WatchOpts, sink chan<- *CacheManagerInsertBid, codehash [][32]byte) (event.Subscription, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "InsertBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerInsertBid)
				if err := _CacheManager.contract.UnpackLog(event, "InsertBid", log); err != nil {
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

// ParseInsertBid is a log parse operation binding the contract event 0xb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec.
//
// Solidity: event InsertBid(bytes32 indexed codehash, address program, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) ParseInsertBid(log types.Log) (*CacheManagerInsertBid, error) {
	event := new(CacheManagerInsertBid)
	if err := _CacheManager.contract.UnpackLog(event, "InsertBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the CacheManager contract.
type CacheManagerPauseIterator struct {
	Event *CacheManagerPause // Event containing the contract specifics and raw log

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
func (it *CacheManagerPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerPause)
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
		it.Event = new(CacheManagerPause)
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
func (it *CacheManagerPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerPause represents a Pause event raised by the CacheManager contract.
type CacheManagerPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CacheManager *CacheManagerFilterer) FilterPause(opts *bind.FilterOpts) (*CacheManagerPauseIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &CacheManagerPauseIterator{contract: _CacheManager.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CacheManager *CacheManagerFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *CacheManagerPause) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerPause)
				if err := _CacheManager.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CacheManager *CacheManagerFilterer) ParsePause(log types.Log) (*CacheManagerPause, error) {
	event := new(CacheManagerPause)
	if err := _CacheManager.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerSetCacheSizeIterator is returned from FilterSetCacheSize and is used to iterate over the raw logs and unpacked data for SetCacheSize events raised by the CacheManager contract.
type CacheManagerSetCacheSizeIterator struct {
	Event *CacheManagerSetCacheSize // Event containing the contract specifics and raw log

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
func (it *CacheManagerSetCacheSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerSetCacheSize)
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
		it.Event = new(CacheManagerSetCacheSize)
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
func (it *CacheManagerSetCacheSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerSetCacheSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerSetCacheSize represents a SetCacheSize event raised by the CacheManager contract.
type CacheManagerSetCacheSize struct {
	Size uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetCacheSize is a free log retrieval operation binding the contract event 0xca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839.
//
// Solidity: event SetCacheSize(uint64 size)
func (_CacheManager *CacheManagerFilterer) FilterSetCacheSize(opts *bind.FilterOpts) (*CacheManagerSetCacheSizeIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "SetCacheSize")
	if err != nil {
		return nil, err
	}
	return &CacheManagerSetCacheSizeIterator{contract: _CacheManager.contract, event: "SetCacheSize", logs: logs, sub: sub}, nil
}

// WatchSetCacheSize is a free log subscription operation binding the contract event 0xca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839.
//
// Solidity: event SetCacheSize(uint64 size)
func (_CacheManager *CacheManagerFilterer) WatchSetCacheSize(opts *bind.WatchOpts, sink chan<- *CacheManagerSetCacheSize) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "SetCacheSize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerSetCacheSize)
				if err := _CacheManager.contract.UnpackLog(event, "SetCacheSize", log); err != nil {
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

// ParseSetCacheSize is a log parse operation binding the contract event 0xca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839.
//
// Solidity: event SetCacheSize(uint64 size)
func (_CacheManager *CacheManagerFilterer) ParseSetCacheSize(log types.Log) (*CacheManagerSetCacheSize, error) {
	event := new(CacheManagerSetCacheSize)
	if err := _CacheManager.contract.UnpackLog(event, "SetCacheSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerSetDecayRateIterator is returned from FilterSetDecayRate and is used to iterate over the raw logs and unpacked data for SetDecayRate events raised by the CacheManager contract.
type CacheManagerSetDecayRateIterator struct {
	Event *CacheManagerSetDecayRate // Event containing the contract specifics and raw log

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
func (it *CacheManagerSetDecayRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerSetDecayRate)
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
		it.Event = new(CacheManagerSetDecayRate)
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
func (it *CacheManagerSetDecayRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerSetDecayRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerSetDecayRate represents a SetDecayRate event raised by the CacheManager contract.
type CacheManagerSetDecayRate struct {
	Decay uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetDecayRate is a free log retrieval operation binding the contract event 0xd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c.
//
// Solidity: event SetDecayRate(uint64 decay)
func (_CacheManager *CacheManagerFilterer) FilterSetDecayRate(opts *bind.FilterOpts) (*CacheManagerSetDecayRateIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "SetDecayRate")
	if err != nil {
		return nil, err
	}
	return &CacheManagerSetDecayRateIterator{contract: _CacheManager.contract, event: "SetDecayRate", logs: logs, sub: sub}, nil
}

// WatchSetDecayRate is a free log subscription operation binding the contract event 0xd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c.
//
// Solidity: event SetDecayRate(uint64 decay)
func (_CacheManager *CacheManagerFilterer) WatchSetDecayRate(opts *bind.WatchOpts, sink chan<- *CacheManagerSetDecayRate) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "SetDecayRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerSetDecayRate)
				if err := _CacheManager.contract.UnpackLog(event, "SetDecayRate", log); err != nil {
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

// ParseSetDecayRate is a log parse operation binding the contract event 0xd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c.
//
// Solidity: event SetDecayRate(uint64 decay)
func (_CacheManager *CacheManagerFilterer) ParseSetDecayRate(log types.Log) (*CacheManagerSetDecayRate, error) {
	event := new(CacheManagerSetDecayRate)
	if err := _CacheManager.contract.UnpackLog(event, "SetDecayRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the CacheManager contract.
type CacheManagerUnpauseIterator struct {
	Event *CacheManagerUnpause // Event containing the contract specifics and raw log

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
func (it *CacheManagerUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerUnpause)
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
		it.Event = new(CacheManagerUnpause)
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
func (it *CacheManagerUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerUnpause represents a Unpause event raised by the CacheManager contract.
type CacheManagerUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CacheManager *CacheManagerFilterer) FilterUnpause(opts *bind.FilterOpts) (*CacheManagerUnpauseIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &CacheManagerUnpauseIterator{contract: _CacheManager.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CacheManager *CacheManagerFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *CacheManagerUnpause) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerUnpause)
				if err := _CacheManager.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CacheManager *CacheManagerFilterer) ParseUnpause(log types.Log) (*CacheManagerUnpause, error) {
	event := new(CacheManagerUnpause)
	if err := _CacheManager.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResourceConstraintManagerMetaData contains all meta data concerning the ResourceConstraintManager contract.
var ResourceConstraintManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"gasTargetPerSec\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"adjustmentWindowSecs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startingBacklogValue\",\"type\":\"uint64\"}],\"name\":\"InvalidPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"gasTargetPerSec\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"adjustmentWindowSecs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startingBacklogValue\",\"type\":\"uint64\"}],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pricingExponent\",\"type\":\"uint64\"}],\"name\":\"PricingExponentTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConstraints\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[3][]\",\"name\":\"constraints\",\"type\":\"uint64[3][]\"}],\"name\":\"setGasPricingConstraints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200143e3803806200143e8339810160408190526200003491620001fb565b6200004160008462000079565b6200006d7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b088362000079565b600255506200023c9050565b62000085828262000089565b5050565b620000a08282620000cc60201b6200078d1760201c565b6000828152600160209081526040909120620000c79183906200082b6200016c821b17901c565b505050565b6000828152602081815260408083206001600160a01b038516845290915290205460ff1662000085576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620001283390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600062000183836001600160a01b0384166200018c565b90505b92915050565b6000818152600183016020526040812054620001d55750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000186565b50600062000186565b80516001600160a01b0381168114620001f657600080fd5b919050565b6000806000606084860312156200021157600080fd5b6200021c84620001de565b92506200022c60208501620001de565b9150604084015190509250925092565b6111f2806200024c6000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c8063a217fddf1161008c578063ca15c87311610066578063ca15c873146101e0578063cc0d556a146101f3578063d547741f14610206578063ec87621c1461021957600080fd5b8063a217fddf146101c7578063ade6e2aa146101cf578063b6549f75146101d857600080fd5b806336568abe116100bd57806336568abe146101525780639010d07c1461016557806391d148541461019057600080fd5b806301ffc9a7146100e4578063248a9ca31461010c5780632f2ff15d1461013d575b600080fd5b6100f76100f2366004610dec565b610240565b60405190151581526020015b60405180910390f35b61012f61011a366004610e2e565b60009081526020819052604090206001015490565b604051908152602001610103565b61015061014b366004610e47565b61029c565b005b610150610160366004610e47565b6102c6565b610178610173366004610e83565b610357565b6040516001600160a01b039091168152602001610103565b6100f761019e366004610e47565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b61012f600081565b61012f60025481565b610150610376565b61012f6101ee366004610e2e565b61041e565b610150610201366004610ea5565b610435565b610150610214366004610e47565b610768565b61012f7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f000000000000000000000000000000000000000000000000000000001480610296575061029682610840565b92915050565b6000828152602081905260409020600101546102b7816108d7565b6102c183836108e4565b505050565b6001600160a01b03811633146103495760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c66000000000000000000000000000000000060648201526084015b60405180910390fd5b6103538282610906565b5050565b600082815260016020526040812061036f9083610928565b9392505050565b6002544210156103b2576040517fd0404f8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f8792701a000000000000000000000000000000000000000000000000000000008152306004820152607090638792701a90602401600060405180830381600087803b15801561040457600080fd5b505af1158015610418573d6000803e3d6000fd5b50505050565b600081815260016020526040812061029690610934565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0861045f816108d7565b81600a81111561049b576040517f4423056400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000805b8281101561069d5760008686838181106104bb576104bb610f1a565b9050606002016000600381106104d3576104d3610f1a565b6020020160208101906104e69190610f4d565b905060008787848181106104fc576104fc610f1a565b90506060020160016003811061051457610514610f1a565b6020020160208101906105279190610f4d565b9050600088888581811061053d5761053d610f1a565b90506060020160026003811061055557610555610f1a565b6020020160208101906105689190610f4d565b9050626acfc08367ffffffffffffffff16108061059257506305f5e1008367ffffffffffffffff16115b156105e5576040517ff12a3edd00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8085166004830152808416602483015282166044820152606401610340565b60058267ffffffffffffffff16108061060a5750620151808267ffffffffffffffff16115b1561065d576040517f6125e37900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8085166004830152808416602483015282166044820152606401610340565b6106678284610f7e565b610673826103e8610f7e565b61067d9190610faa565b6106879086610fdf565b94505050508061069690611000565b905061049f565b50611f408167ffffffffffffffff1611156106f0576040517fc381458900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610340565b6040517fcc0d556a00000000000000000000000000000000000000000000000000000000815260709063cc0d556a9061072f908890889060040161101a565b600060405180830381600087803b15801561074957600080fd5b505af115801561075d573d6000803e3d6000fd5b505050505050505050565b600082815260208190526040902060010154610783816108d7565b6102c18383610906565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16610353576000828152602081815260408083206001600160a01b03851684529091529020805460ff191660011790556107e73390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600061036f836001600160a01b03841661093e565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061029657507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610296565b6108e1813361098d565b50565b6108ee828261078d565b60008281526001602052604090206102c1908261082b565b6109108282610a0b565b60008281526001602052604090206102c19082610a8a565b600061036f8383610a9f565b6000610296825490565b600081815260018301602052604081205461098557508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610296565b506000610296565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16610353576109c9816001600160a01b03166014610ac9565b6109d4836020610ac9565b6040516020016109e59291906110b1565b60408051601f198184030181529082905262461bcd60e51b825261034091600401611132565b6000828152602081815260408083206001600160a01b038516845290915290205460ff1615610353576000828152602081815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b600061036f836001600160a01b038416610cf2565b6000826000018281548110610ab657610ab6610f1a565b9060005260206000200154905092915050565b60606000610ad8836002611165565b610ae390600261117c565b67ffffffffffffffff811115610afb57610afb61118f565b6040519080825280601f01601f191660200182016040528015610b25576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610b5c57610b5c610f1a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610bbf57610bbf610f1a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000610bfb846002611165565b610c0690600161117c565b90505b6001811115610ca3577f303132333435363738396162636465660000000000000000000000000000000085600f1660108110610c4757610c47610f1a565b1a60f81b828281518110610c5d57610c5d610f1a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c93610c9c816111a5565b9050610c09565b50831561036f5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610340565b60008181526001830160205260408120548015610ddb576000610d166001836111bc565b8554909150600090610d2a906001906111bc565b9050818114610d8f576000866000018281548110610d4a57610d4a610f1a565b9060005260206000200154905080876000018481548110610d6d57610d6d610f1a565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610da057610da06111cf565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610296565b6000915050610296565b5092915050565b600060208284031215610dfe57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461036f57600080fd5b600060208284031215610e4057600080fd5b5035919050565b60008060408385031215610e5a57600080fd5b8235915060208301356001600160a01b0381168114610e7857600080fd5b809150509250929050565b60008060408385031215610e9657600080fd5b50508035926020909101359150565b60008060208385031215610eb857600080fd5b823567ffffffffffffffff80821115610ed057600080fd5b818501915085601f830112610ee457600080fd5b813581811115610ef357600080fd5b866020606083028501011115610f0857600080fd5b60209290920196919550909350505050565b634e487b7160e01b600052603260045260246000fd5b803567ffffffffffffffff81168114610f4857600080fd5b919050565b600060208284031215610f5f57600080fd5b61036f82610f30565b634e487b7160e01b600052601160045260246000fd5b67ffffffffffffffff818116838216028082169190828114610fa257610fa2610f68565b505092915050565b600067ffffffffffffffff80841680610fd357634e487b7160e01b600052601260045260246000fd5b92169190910492915050565b67ffffffffffffffff818116838216019080821115610de557610de5610f68565b6000600019820361101357611013610f68565b5060010190565b6020808252818101839052600090846040840183805b87811015611080578284835b600381101561106a5767ffffffffffffffff61105783610f30565b168352918701919087019060010161103c565b5050506060938401939290920191600101611030565b5090979650505050505050565b60005b838110156110a8578181015183820152602001611090565b50506000910152565b7f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008152600083516110e981601785016020880161108d565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000601791840191820152835161112681602884016020880161108d565b01602801949350505050565b602081526000825180602084015261115181604085016020870161108d565b601f01601f19169190910160400192915050565b808202811582820484141761029657610296610f68565b8082018082111561029657610296610f68565b634e487b7160e01b600052604160045260246000fd5b6000816111b4576111b4610f68565b506000190190565b8181038181111561029657610296610f68565b634e487b7160e01b600052603160045260246000fdfea164736f6c6343000811000a",
}

// ResourceConstraintManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ResourceConstraintManagerMetaData.ABI instead.
var ResourceConstraintManagerABI = ResourceConstraintManagerMetaData.ABI

// ResourceConstraintManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ResourceConstraintManagerMetaData.Bin instead.
var ResourceConstraintManagerBin = ResourceConstraintManagerMetaData.Bin

// DeployResourceConstraintManager deploys a new Ethereum contract, binding an instance of ResourceConstraintManager to it.
func DeployResourceConstraintManager(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address, manager common.Address, _expiryTimestamp *big.Int) (common.Address, *types.Transaction, *ResourceConstraintManager, error) {
	parsed, err := ResourceConstraintManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ResourceConstraintManagerBin), backend, admin, manager, _expiryTimestamp)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ResourceConstraintManager{ResourceConstraintManagerCaller: ResourceConstraintManagerCaller{contract: contract}, ResourceConstraintManagerTransactor: ResourceConstraintManagerTransactor{contract: contract}, ResourceConstraintManagerFilterer: ResourceConstraintManagerFilterer{contract: contract}}, nil
}

// ResourceConstraintManager is an auto generated Go binding around an Ethereum contract.
type ResourceConstraintManager struct {
	ResourceConstraintManagerCaller     // Read-only binding to the contract
	ResourceConstraintManagerTransactor // Write-only binding to the contract
	ResourceConstraintManagerFilterer   // Log filterer for contract events
}

// ResourceConstraintManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResourceConstraintManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResourceConstraintManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResourceConstraintManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResourceConstraintManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResourceConstraintManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResourceConstraintManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResourceConstraintManagerSession struct {
	Contract     *ResourceConstraintManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ResourceConstraintManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResourceConstraintManagerCallerSession struct {
	Contract *ResourceConstraintManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ResourceConstraintManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResourceConstraintManagerTransactorSession struct {
	Contract     *ResourceConstraintManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ResourceConstraintManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResourceConstraintManagerRaw struct {
	Contract *ResourceConstraintManager // Generic contract binding to access the raw methods on
}

// ResourceConstraintManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResourceConstraintManagerCallerRaw struct {
	Contract *ResourceConstraintManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ResourceConstraintManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResourceConstraintManagerTransactorRaw struct {
	Contract *ResourceConstraintManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResourceConstraintManager creates a new instance of ResourceConstraintManager, bound to a specific deployed contract.
func NewResourceConstraintManager(address common.Address, backend bind.ContractBackend) (*ResourceConstraintManager, error) {
	contract, err := bindResourceConstraintManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResourceConstraintManager{ResourceConstraintManagerCaller: ResourceConstraintManagerCaller{contract: contract}, ResourceConstraintManagerTransactor: ResourceConstraintManagerTransactor{contract: contract}, ResourceConstraintManagerFilterer: ResourceConstraintManagerFilterer{contract: contract}}, nil
}

// NewResourceConstraintManagerCaller creates a new read-only instance of ResourceConstraintManager, bound to a specific deployed contract.
func NewResourceConstraintManagerCaller(address common.Address, caller bind.ContractCaller) (*ResourceConstraintManagerCaller, error) {
	contract, err := bindResourceConstraintManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResourceConstraintManagerCaller{contract: contract}, nil
}

// NewResourceConstraintManagerTransactor creates a new write-only instance of ResourceConstraintManager, bound to a specific deployed contract.
func NewResourceConstraintManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ResourceConstraintManagerTransactor, error) {
	contract, err := bindResourceConstraintManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResourceConstraintManagerTransactor{contract: contract}, nil
}

// NewResourceConstraintManagerFilterer creates a new log filterer instance of ResourceConstraintManager, bound to a specific deployed contract.
func NewResourceConstraintManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ResourceConstraintManagerFilterer, error) {
	contract, err := bindResourceConstraintManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResourceConstraintManagerFilterer{contract: contract}, nil
}

// bindResourceConstraintManager binds a generic wrapper to an already deployed contract.
func bindResourceConstraintManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResourceConstraintManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResourceConstraintManager *ResourceConstraintManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResourceConstraintManager.Contract.ResourceConstraintManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResourceConstraintManager *ResourceConstraintManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.ResourceConstraintManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResourceConstraintManager *ResourceConstraintManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.ResourceConstraintManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResourceConstraintManager *ResourceConstraintManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResourceConstraintManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResourceConstraintManager *ResourceConstraintManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResourceConstraintManager *ResourceConstraintManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ResourceConstraintManager *ResourceConstraintManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ResourceConstraintManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ResourceConstraintManager *ResourceConstraintManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ResourceConstraintManager.Contract.DEFAULTADMINROLE(&_ResourceConstraintManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ResourceConstraintManager *ResourceConstraintManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ResourceConstraintManager.Contract.DEFAULTADMINROLE(&_ResourceConstraintManager.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_ResourceConstraintManager *ResourceConstraintManagerCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ResourceConstraintManager.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_ResourceConstraintManager *ResourceConstraintManagerSession) MANAGERROLE() ([32]byte, error) {
	return _ResourceConstraintManager.Contract.MANAGERROLE(&_ResourceConstraintManager.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_ResourceConstraintManager *ResourceConstraintManagerCallerSession) MANAGERROLE() ([32]byte, error) {
	return _ResourceConstraintManager.Contract.MANAGERROLE(&_ResourceConstraintManager.CallOpts)
}

// ExpiryTimestamp is a free data retrieval call binding the contract method 0xade6e2aa.
//
// Solidity: function expiryTimestamp() view returns(uint256)
func (_ResourceConstraintManager *ResourceConstraintManagerCaller) ExpiryTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ResourceConstraintManager.contract.Call(opts, &out, "expiryTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiryTimestamp is a free data retrieval call binding the contract method 0xade6e2aa.
//
// Solidity: function expiryTimestamp() view returns(uint256)
func (_ResourceConstraintManager *ResourceConstraintManagerSession) ExpiryTimestamp() (*big.Int, error) {
	return _ResourceConstraintManager.Contract.ExpiryTimestamp(&_ResourceConstraintManager.CallOpts)
}

// ExpiryTimestamp is a free data retrieval call binding the contract method 0xade6e2aa.
//
// Solidity: function expiryTimestamp() view returns(uint256)
func (_ResourceConstraintManager *ResourceConstraintManagerCallerSession) ExpiryTimestamp() (*big.Int, error) {
	return _ResourceConstraintManager.Contract.ExpiryTimestamp(&_ResourceConstraintManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ResourceConstraintManager *ResourceConstraintManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ResourceConstraintManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ResourceConstraintManager *ResourceConstraintManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ResourceConstraintManager.Contract.GetRoleAdmin(&_ResourceConstraintManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ResourceConstraintManager *ResourceConstraintManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ResourceConstraintManager.Contract.GetRoleAdmin(&_ResourceConstraintManager.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ResourceConstraintManager *ResourceConstraintManagerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ResourceConstraintManager.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ResourceConstraintManager *ResourceConstraintManagerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ResourceConstraintManager.Contract.GetRoleMember(&_ResourceConstraintManager.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ResourceConstraintManager *ResourceConstraintManagerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ResourceConstraintManager.Contract.GetRoleMember(&_ResourceConstraintManager.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ResourceConstraintManager *ResourceConstraintManagerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ResourceConstraintManager.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ResourceConstraintManager *ResourceConstraintManagerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ResourceConstraintManager.Contract.GetRoleMemberCount(&_ResourceConstraintManager.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ResourceConstraintManager *ResourceConstraintManagerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ResourceConstraintManager.Contract.GetRoleMemberCount(&_ResourceConstraintManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ResourceConstraintManager *ResourceConstraintManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ResourceConstraintManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ResourceConstraintManager *ResourceConstraintManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ResourceConstraintManager.Contract.HasRole(&_ResourceConstraintManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ResourceConstraintManager *ResourceConstraintManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ResourceConstraintManager.Contract.HasRole(&_ResourceConstraintManager.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ResourceConstraintManager *ResourceConstraintManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ResourceConstraintManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ResourceConstraintManager *ResourceConstraintManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ResourceConstraintManager.Contract.SupportsInterface(&_ResourceConstraintManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ResourceConstraintManager *ResourceConstraintManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ResourceConstraintManager.Contract.SupportsInterface(&_ResourceConstraintManager.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResourceConstraintManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.GrantRole(&_ResourceConstraintManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.GrantRole(&_ResourceConstraintManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResourceConstraintManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.RenounceRole(&_ResourceConstraintManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.RenounceRole(&_ResourceConstraintManager.TransactOpts, role, account)
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_ResourceConstraintManager *ResourceConstraintManagerTransactor) Revoke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResourceConstraintManager.contract.Transact(opts, "revoke")
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_ResourceConstraintManager *ResourceConstraintManagerSession) Revoke() (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.Revoke(&_ResourceConstraintManager.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_ResourceConstraintManager *ResourceConstraintManagerTransactorSession) Revoke() (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.Revoke(&_ResourceConstraintManager.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResourceConstraintManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.RevokeRole(&_ResourceConstraintManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.RevokeRole(&_ResourceConstraintManager.TransactOpts, role, account)
}

// SetGasPricingConstraints is a paid mutator transaction binding the contract method 0xcc0d556a.
//
// Solidity: function setGasPricingConstraints(uint64[3][] constraints) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerTransactor) SetGasPricingConstraints(opts *bind.TransactOpts, constraints [][3]uint64) (*types.Transaction, error) {
	return _ResourceConstraintManager.contract.Transact(opts, "setGasPricingConstraints", constraints)
}

// SetGasPricingConstraints is a paid mutator transaction binding the contract method 0xcc0d556a.
//
// Solidity: function setGasPricingConstraints(uint64[3][] constraints) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerSession) SetGasPricingConstraints(constraints [][3]uint64) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.SetGasPricingConstraints(&_ResourceConstraintManager.TransactOpts, constraints)
}

// SetGasPricingConstraints is a paid mutator transaction binding the contract method 0xcc0d556a.
//
// Solidity: function setGasPricingConstraints(uint64[3][] constraints) returns()
func (_ResourceConstraintManager *ResourceConstraintManagerTransactorSession) SetGasPricingConstraints(constraints [][3]uint64) (*types.Transaction, error) {
	return _ResourceConstraintManager.Contract.SetGasPricingConstraints(&_ResourceConstraintManager.TransactOpts, constraints)
}

// ResourceConstraintManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ResourceConstraintManager contract.
type ResourceConstraintManagerRoleAdminChangedIterator struct {
	Event *ResourceConstraintManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ResourceConstraintManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResourceConstraintManagerRoleAdminChanged)
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
		it.Event = new(ResourceConstraintManagerRoleAdminChanged)
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
func (it *ResourceConstraintManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResourceConstraintManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResourceConstraintManagerRoleAdminChanged represents a RoleAdminChanged event raised by the ResourceConstraintManager contract.
type ResourceConstraintManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ResourceConstraintManager *ResourceConstraintManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ResourceConstraintManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ResourceConstraintManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ResourceConstraintManagerRoleAdminChangedIterator{contract: _ResourceConstraintManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ResourceConstraintManager *ResourceConstraintManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ResourceConstraintManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ResourceConstraintManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResourceConstraintManagerRoleAdminChanged)
				if err := _ResourceConstraintManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ResourceConstraintManager *ResourceConstraintManagerFilterer) ParseRoleAdminChanged(log types.Log) (*ResourceConstraintManagerRoleAdminChanged, error) {
	event := new(ResourceConstraintManagerRoleAdminChanged)
	if err := _ResourceConstraintManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResourceConstraintManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ResourceConstraintManager contract.
type ResourceConstraintManagerRoleGrantedIterator struct {
	Event *ResourceConstraintManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *ResourceConstraintManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResourceConstraintManagerRoleGranted)
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
		it.Event = new(ResourceConstraintManagerRoleGranted)
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
func (it *ResourceConstraintManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResourceConstraintManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResourceConstraintManagerRoleGranted represents a RoleGranted event raised by the ResourceConstraintManager contract.
type ResourceConstraintManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ResourceConstraintManager *ResourceConstraintManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ResourceConstraintManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ResourceConstraintManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ResourceConstraintManagerRoleGrantedIterator{contract: _ResourceConstraintManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ResourceConstraintManager *ResourceConstraintManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ResourceConstraintManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ResourceConstraintManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResourceConstraintManagerRoleGranted)
				if err := _ResourceConstraintManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ResourceConstraintManager *ResourceConstraintManagerFilterer) ParseRoleGranted(log types.Log) (*ResourceConstraintManagerRoleGranted, error) {
	event := new(ResourceConstraintManagerRoleGranted)
	if err := _ResourceConstraintManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResourceConstraintManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ResourceConstraintManager contract.
type ResourceConstraintManagerRoleRevokedIterator struct {
	Event *ResourceConstraintManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ResourceConstraintManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResourceConstraintManagerRoleRevoked)
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
		it.Event = new(ResourceConstraintManagerRoleRevoked)
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
func (it *ResourceConstraintManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResourceConstraintManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResourceConstraintManagerRoleRevoked represents a RoleRevoked event raised by the ResourceConstraintManager contract.
type ResourceConstraintManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ResourceConstraintManager *ResourceConstraintManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ResourceConstraintManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ResourceConstraintManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ResourceConstraintManagerRoleRevokedIterator{contract: _ResourceConstraintManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ResourceConstraintManager *ResourceConstraintManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ResourceConstraintManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ResourceConstraintManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResourceConstraintManagerRoleRevoked)
				if err := _ResourceConstraintManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ResourceConstraintManager *ResourceConstraintManagerFilterer) ParseRoleRevoked(log types.Log) (*ResourceConstraintManagerRoleRevoked, error) {
	event := new(ResourceConstraintManagerRoleRevoked)
	if err := _ResourceConstraintManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
