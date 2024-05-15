package com.metavirus.adminserver.contract;

import io.reactivex.Flowable;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import org.web3j.abi.EventEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Address;
import org.web3j.abi.datatypes.Event;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.Utf8String;
import org.web3j.abi.datatypes.generated.Bytes32;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.abi.datatypes.generated.Uint64;
import org.web3j.abi.datatypes.generated.Uint8;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.DefaultBlockParameter;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.RemoteFunctionCall;
import org.web3j.protocol.core.methods.request.EthFilter;
import org.web3j.protocol.core.methods.response.BaseEventResponse;
import org.web3j.protocol.core.methods.response.Log;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.ContractGasProvider;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 1.5.3.
 */
@SuppressWarnings("rawtypes")
public class NexGami extends Contract {
    public static final String BINARY = "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152503480156200004457600080fd5b50620000556200005b60201b60201c565b620001cf565b60006200006d6200016560201b60201c565b90508060000160089054906101000a900460ff1615620000b9576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614620001625767ffffffffffffffff8160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff604051620001599190620001b2565b60405180910390a15b50565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600067ffffffffffffffff82169050919050565b620001ac816200018d565b82525050565b6000602082019050620001c96000830184620001a1565b92915050565b608051612642620001f960003960008181610d2501528181610d7a0152610f3501526126426000f3fe6080604052600436106101095760003560e01c8063715018a611610095578063a9059cbb11610064578063a9059cbb1461034c578063ad3cb1cc14610389578063c4d66de8146103b4578063dd62ed3e146103dd578063f2fde38b1461041a57610109565b8063715018a6146102b657806379cc6790146102cd5780638da5cb5b146102f657806395d89b411461032157610109565b8063313ce567116100dc578063313ce567146101de57806342966c68146102095780634f1ef2861461023257806352d1902d1461024e57806370a082311461027957610109565b806306fdde031461010e578063095ea7b31461013957806318160ddd1461017657806323b872dd146101a1575b600080fd5b34801561011a57600080fd5b50610123610443565b6040516101309190611b36565b60405180910390f35b34801561014557600080fd5b50610160600480360381019061015b9190611c00565b6104e4565b60405161016d9190611c5b565b60405180910390f35b34801561018257600080fd5b5061018b610507565b6040516101989190611c85565b60405180910390f35b3480156101ad57600080fd5b506101c860048036038101906101c39190611ca0565b61051f565b6040516101d59190611c5b565b60405180910390f35b3480156101ea57600080fd5b506101f361054e565b6040516102009190611d0f565b60405180910390f35b34801561021557600080fd5b50610230600480360381019061022b9190611d2a565b610557565b005b61024c60048036038101906102479190611e8c565b61056b565b005b34801561025a57600080fd5b5061026361058a565b6040516102709190611f01565b60405180910390f35b34801561028557600080fd5b506102a0600480360381019061029b9190611f1c565b6105bd565b6040516102ad9190611c85565b60405180910390f35b3480156102c257600080fd5b506102cb610614565b005b3480156102d957600080fd5b506102f460048036038101906102ef9190611c00565b610628565b005b34801561030257600080fd5b5061030b610648565b6040516103189190611f58565b60405180910390f35b34801561032d57600080fd5b50610336610680565b6040516103439190611b36565b60405180910390f35b34801561035857600080fd5b50610373600480360381019061036e9190611c00565b610721565b6040516103809190611c5b565b60405180910390f35b34801561039557600080fd5b5061039e610744565b6040516103ab9190611b36565b60405180910390f35b3480156103c057600080fd5b506103db60048036038101906103d69190611f1c565b61077d565b005b3480156103e957600080fd5b5061040460048036038101906103ff9190611f73565b6109bc565b6040516104119190611c85565b60405180910390f35b34801561042657600080fd5b50610441600480360381019061043c9190611f1c565b610a51565b005b6060600061044f610ad7565b905080600301805461046090611fe2565b80601f016020809104026020016040519081016040528092919081815260200182805461048c90611fe2565b80156104d95780601f106104ae576101008083540402835291602001916104d9565b820191906000526020600020905b8154815290600101906020018083116104bc57829003601f168201915b505050505091505090565b6000806104ef610aff565b90506104fc818585610b07565b600191505092915050565b600080610512610ad7565b9050806002015491505090565b60008061052a610aff565b9050610537858285610b19565b610542858585610bad565b60019150509392505050565b60006012905090565b610568610562610aff565b82610ca1565b50565b610573610d23565b61057c82610e09565b6105868282610e14565b5050565b6000610594610f33565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b6000806105c8610ad7565b90508060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054915050919050565b61061c610fba565b6106266000611041565b565b61063a82610634610aff565b83610b19565b6106448282610ca1565b5050565b600080610653611118565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b6060600061068c610ad7565b905080600401805461069d90611fe2565b80601f01602080910402602001604051908101604052809291908181526020018280546106c990611fe2565b80156107165780601f106106eb57610100808354040283529160200191610716565b820191906000526020600020905b8154815290600101906020018083116106f957829003601f168201915b505050505091505090565b60008061072c610aff565b9050610739818585610bad565b600191505092915050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6000610787611140565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff161480156107d55750825b9050600060018367ffffffffffffffff1614801561080a575060003073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610818575080155b1561084f576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550831561089f5760018560000160086101000a81548160ff0219169083151502179055505b6109136040518060400160405280600781526020017f4e657847616d69000000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f4e45584700000000000000000000000000000000000000000000000000000000815250611168565b61091b61117e565b61092486611188565b61092c61119c565b6109583361093861054e565b600a6109449190612175565b633b9aca0061095391906121c0565b6111a6565b83156109b45760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516109ab919061225b565b60405180910390a15b505050505050565b6000806109c7610ad7565b90508060010160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205491505092915050565b610a59610fba565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610acb5760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610ac29190611f58565b60405180910390fd5b610ad481611041565b50565b60007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00905090565b600033905090565b610b148383836001611228565b505050565b6000610b2584846109bc565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610ba75781811015610b97578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401610b8e93929190612276565b60405180910390fd5b610ba684848484036000611228565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c1f5760006040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610c169190611f58565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c915760006040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610c889190611f58565b60405180910390fd5b610c9c83838361140e565b505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d135760006040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610d0a9190611f58565b60405180910390fd5b610d1f8260008361140e565b5050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480610dd057507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610db761164d565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610e07576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610e11610fba565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610e7c57506040513d601f19601f82011682018060405250810190610e7991906122d9565b60015b610ebd57816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401610eb49190611f58565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114610f2457806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401610f1b9190611f01565b60405180910390fd5b610f2e83836116a4565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610fb8576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610fc2610aff565b73ffffffffffffffffffffffffffffffffffffffff16610fe0610648565b73ffffffffffffffffffffffffffffffffffffffff161461103f57611003610aff565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016110369190611f58565b60405180910390fd5b565b600061104b611118565b905060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b60007f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b611170611717565b61117a8282611757565b5050565b611186611717565b565b611190611717565b61119981611794565b50565b6111a4611717565b565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036112185760006040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161120f9190611f58565b60405180910390fd5b6112246000838361140e565b5050565b6000611232610ad7565b9050600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16036112a65760006040517fe602df0500000000000000000000000000000000000000000000000000000000815260040161129d9190611f58565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036113185760006040517f94280d6200000000000000000000000000000000000000000000000000000000815260040161130f9190611f58565b60405180910390fd5b828160010160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508115611407578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040516113fe9190611c85565b60405180910390a35b5050505050565b6000611418610ad7565b9050600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff160361146e57818160020160008282546114629190612306565b92505081905550611547565b60008160000160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156114fd578481846040517fe450d38c0000000000000000000000000000000000000000000000000000000081526004016114f493929190612276565b60405180910390fd5b8281038260000160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611592578181600201600082825403925050819055506115e2565b818160000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161163f9190611c85565b60405180910390a350505050565b600061167b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61181a565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6116ad82611824565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a260008151111561170a5761170482826118f1565b50611713565b611712611975565b5b5050565b61171f6119b2565b611755576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b61175f611717565b6000611769610ad7565b90508281600301908161177c91906124dc565b508181600401908161178e91906124dc565b50505050565b61179c611717565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361180e5760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016118059190611f58565b60405180910390fd5b61181781611041565b50565b6000819050919050565b60008173ffffffffffffffffffffffffffffffffffffffff163b0361188057806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016118779190611f58565b60405180910390fd5b806118ad7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61181a565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000808473ffffffffffffffffffffffffffffffffffffffff168460405161191b91906125f5565b600060405180830381855af49150503d8060008114611956576040519150601f19603f3d011682016040523d82523d6000602084013e61195b565b606091505b509150915061196b8583836119d2565b9250505092915050565b60003411156119b0576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60006119bc611140565b60000160089054906101000a900460ff16905090565b6060826119e7576119e282611a61565b611a59565b60008251148015611a0f575060008473ffffffffffffffffffffffffffffffffffffffff163b145b15611a5157836040517f9996b315000000000000000000000000000000000000000000000000000000008152600401611a489190611f58565b60405180910390fd5b819050611a5a565b5b9392505050565b600081511115611a745780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081519050919050565b600082825260208201905092915050565b60005b83811015611ae0578082015181840152602081019050611ac5565b60008484015250505050565b6000601f19601f8301169050919050565b6000611b0882611aa6565b611b128185611ab1565b9350611b22818560208601611ac2565b611b2b81611aec565b840191505092915050565b60006020820190508181036000830152611b508184611afd565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611b9782611b6c565b9050919050565b611ba781611b8c565b8114611bb257600080fd5b50565b600081359050611bc481611b9e565b92915050565b6000819050919050565b611bdd81611bca565b8114611be857600080fd5b50565b600081359050611bfa81611bd4565b92915050565b60008060408385031215611c1757611c16611b62565b5b6000611c2585828601611bb5565b9250506020611c3685828601611beb565b9150509250929050565b60008115159050919050565b611c5581611c40565b82525050565b6000602082019050611c706000830184611c4c565b92915050565b611c7f81611bca565b82525050565b6000602082019050611c9a6000830184611c76565b92915050565b600080600060608486031215611cb957611cb8611b62565b5b6000611cc786828701611bb5565b9350506020611cd886828701611bb5565b9250506040611ce986828701611beb565b9150509250925092565b600060ff82169050919050565b611d0981611cf3565b82525050565b6000602082019050611d246000830184611d00565b92915050565b600060208284031215611d4057611d3f611b62565b5b6000611d4e84828501611beb565b91505092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611d9982611aec565b810181811067ffffffffffffffff82111715611db857611db7611d61565b5b80604052505050565b6000611dcb611b58565b9050611dd78282611d90565b919050565b600067ffffffffffffffff821115611df757611df6611d61565b5b611e0082611aec565b9050602081019050919050565b82818337600083830152505050565b6000611e2f611e2a84611ddc565b611dc1565b905082815260208101848484011115611e4b57611e4a611d5c565b5b611e56848285611e0d565b509392505050565b600082601f830112611e7357611e72611d57565b5b8135611e83848260208601611e1c565b91505092915050565b60008060408385031215611ea357611ea2611b62565b5b6000611eb185828601611bb5565b925050602083013567ffffffffffffffff811115611ed257611ed1611b67565b5b611ede85828601611e5e565b9150509250929050565b6000819050919050565b611efb81611ee8565b82525050565b6000602082019050611f166000830184611ef2565b92915050565b600060208284031215611f3257611f31611b62565b5b6000611f4084828501611bb5565b91505092915050565b611f5281611b8c565b82525050565b6000602082019050611f6d6000830184611f49565b92915050565b60008060408385031215611f8a57611f89611b62565b5b6000611f9885828601611bb5565b9250506020611fa985828601611bb5565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611ffa57607f821691505b60208210810361200d5761200c611fb3565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b60018511156120995780860481111561207557612074612013565b5b60018516156120845780820291505b808102905061209285612042565b9450612059565b94509492505050565b6000826120b2576001905061216e565b816120c0576000905061216e565b81600181146120d657600281146120e05761210f565b600191505061216e565b60ff8411156120f2576120f1612013565b5b8360020a91508482111561210957612108612013565b5b5061216e565b5060208310610133831016604e8410600b84101617156121445782820a90508381111561213f5761213e612013565b5b61216e565b612151848484600161204f565b9250905081840481111561216857612167612013565b5b81810290505b9392505050565b600061218082611bca565b915061218b83611cf3565b92506121b87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846120a2565b905092915050565b60006121cb82611bca565b91506121d683611bca565b92508282026121e481611bca565b915082820484148315176121fb576121fa612013565b5b5092915050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000819050919050565b600061224561224061223b84612202565b612220565b61220c565b9050919050565b6122558161222a565b82525050565b6000602082019050612270600083018461224c565b92915050565b600060608201905061228b6000830186611f49565b6122986020830185611c76565b6122a56040830184611c76565b949350505050565b6122b681611ee8565b81146122c157600080fd5b50565b6000815190506122d3816122ad565b92915050565b6000602082840312156122ef576122ee611b62565b5b60006122fd848285016122c4565b91505092915050565b600061231182611bca565b915061231c83611bca565b925082820190508082111561233457612333612013565b5b92915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261239c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261235f565b6123a6868361235f565b95508019841693508086168417925050509392505050565b60006123d96123d46123cf84611bca565b612220565b611bca565b9050919050565b6000819050919050565b6123f3836123be565b6124076123ff826123e0565b84845461236c565b825550505050565b600090565b61241c61240f565b6124278184846123ea565b505050565b5b8181101561244b57612440600082612414565b60018101905061242d565b5050565b601f821115612490576124618161233a565b61246a8461234f565b81016020851015612479578190505b61248d6124858561234f565b83018261242c565b50505b505050565b600082821c905092915050565b60006124b360001984600802612495565b1980831691505092915050565b60006124cc83836124a2565b9150826002028217905092915050565b6124e582611aa6565b67ffffffffffffffff8111156124fe576124fd611d61565b5b6125088254611fe2565b61251382828561244f565b600060209050601f8311600181146125465760008415612534578287015190505b61253e85826124c0565b8655506125a6565b601f1984166125548661233a565b60005b8281101561257c57848901518255600182019150602085019450602081019050612557565b868310156125995784890151612595601f8916826124a2565b8355505b6001600288020188555050505b505050505050565b600081519050919050565b600081905092915050565b60006125cf826125ae565b6125d981856125b9565b93506125e9818560208601611ac2565b80840191505092915050565b600061260182846125c4565b91508190509291505056fea2646970667358221220843a63a8c407247403ba5c84bc038b34b5446dccccd1adcb251ea300b8f8cf1c64736f6c63430008170033";

    private static String librariesLinkedBinary;

    public static final String FUNC_UPGRADE_INTERFACE_VERSION = "UPGRADE_INTERFACE_VERSION";

    public static final String FUNC_ALLOWANCE = "allowance";

    public static final String FUNC_APPROVE = "approve";

    public static final String FUNC_BALANCEOF = "balanceOf";

    public static final String FUNC_BURN = "burn";

    public static final String FUNC_BURNFROM = "burnFrom";

    public static final String FUNC_DECIMALS = "decimals";

    public static final String FUNC_INITIALIZE = "initialize";

    public static final String FUNC_NAME = "name";

    public static final String FUNC_OWNER = "owner";

    public static final String FUNC_PROXIABLEUUID = "proxiableUUID";

    public static final String FUNC_RENOUNCEOWNERSHIP = "renounceOwnership";

    public static final String FUNC_SYMBOL = "symbol";

    public static final String FUNC_TOTALSUPPLY = "totalSupply";

    public static final String FUNC_TRANSFER = "transfer";

    public static final String FUNC_TRANSFERFROM = "transferFrom";

    public static final String FUNC_TRANSFEROWNERSHIP = "transferOwnership";

    public static final String FUNC_UPGRADETOANDCALL = "upgradeToAndCall";

    public static final Event APPROVAL_EVENT = new Event("Approval", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}, new TypeReference<Uint256>() {}));
    ;

    public static final Event INITIALIZED_EVENT = new Event("Initialized", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Uint64>() {}));
    ;

    public static final Event OWNERSHIPTRANSFERRED_EVENT = new Event("OwnershipTransferred", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}));
    ;

    public static final Event TRANSFER_EVENT = new Event("Transfer", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}, new TypeReference<Uint256>() {}));
    ;

    public static final Event UPGRADED_EVENT = new Event("Upgraded", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}));
    ;

    @Deprecated
    protected NexGami(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected NexGami(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected NexGami(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected NexGami(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static List<ApprovalEventResponse> getApprovalEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = staticExtractEventParametersWithLog(APPROVAL_EVENT, transactionReceipt);
        ArrayList<ApprovalEventResponse> responses = new ArrayList<ApprovalEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            ApprovalEventResponse typedResponse = new ApprovalEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.owner = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.spender = (String) eventValues.getIndexedValues().get(1).getValue();
            typedResponse.value = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public static ApprovalEventResponse getApprovalEventFromLog(Log log) {
        Contract.EventValuesWithLog eventValues = staticExtractEventParametersWithLog(APPROVAL_EVENT, log);
        ApprovalEventResponse typedResponse = new ApprovalEventResponse();
        typedResponse.log = log;
        typedResponse.owner = (String) eventValues.getIndexedValues().get(0).getValue();
        typedResponse.spender = (String) eventValues.getIndexedValues().get(1).getValue();
        typedResponse.value = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
        return typedResponse;
    }

    public Flowable<ApprovalEventResponse> approvalEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(log -> getApprovalEventFromLog(log));
    }

    public Flowable<ApprovalEventResponse> approvalEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(APPROVAL_EVENT));
        return approvalEventFlowable(filter);
    }

    public static List<InitializedEventResponse> getInitializedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = staticExtractEventParametersWithLog(INITIALIZED_EVENT, transactionReceipt);
        ArrayList<InitializedEventResponse> responses = new ArrayList<InitializedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            InitializedEventResponse typedResponse = new InitializedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.version = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public static InitializedEventResponse getInitializedEventFromLog(Log log) {
        Contract.EventValuesWithLog eventValues = staticExtractEventParametersWithLog(INITIALIZED_EVENT, log);
        InitializedEventResponse typedResponse = new InitializedEventResponse();
        typedResponse.log = log;
        typedResponse.version = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
        return typedResponse;
    }

    public Flowable<InitializedEventResponse> initializedEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(log -> getInitializedEventFromLog(log));
    }

    public Flowable<InitializedEventResponse> initializedEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(INITIALIZED_EVENT));
        return initializedEventFlowable(filter);
    }

    public static List<OwnershipTransferredEventResponse> getOwnershipTransferredEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = staticExtractEventParametersWithLog(OWNERSHIPTRANSFERRED_EVENT, transactionReceipt);
        ArrayList<OwnershipTransferredEventResponse> responses = new ArrayList<OwnershipTransferredEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            OwnershipTransferredEventResponse typedResponse = new OwnershipTransferredEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.previousOwner = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.newOwner = (String) eventValues.getIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public static OwnershipTransferredEventResponse getOwnershipTransferredEventFromLog(Log log) {
        Contract.EventValuesWithLog eventValues = staticExtractEventParametersWithLog(OWNERSHIPTRANSFERRED_EVENT, log);
        OwnershipTransferredEventResponse typedResponse = new OwnershipTransferredEventResponse();
        typedResponse.log = log;
        typedResponse.previousOwner = (String) eventValues.getIndexedValues().get(0).getValue();
        typedResponse.newOwner = (String) eventValues.getIndexedValues().get(1).getValue();
        return typedResponse;
    }

    public Flowable<OwnershipTransferredEventResponse> ownershipTransferredEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(log -> getOwnershipTransferredEventFromLog(log));
    }

    public Flowable<OwnershipTransferredEventResponse> ownershipTransferredEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(OWNERSHIPTRANSFERRED_EVENT));
        return ownershipTransferredEventFlowable(filter);
    }

    public static List<TransferEventResponse> getTransferEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = staticExtractEventParametersWithLog(TRANSFER_EVENT, transactionReceipt);
        ArrayList<TransferEventResponse> responses = new ArrayList<TransferEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            TransferEventResponse typedResponse = new TransferEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.from = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.to = (String) eventValues.getIndexedValues().get(1).getValue();
            typedResponse.value = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public static TransferEventResponse getTransferEventFromLog(Log log) {
        Contract.EventValuesWithLog eventValues = staticExtractEventParametersWithLog(TRANSFER_EVENT, log);
        TransferEventResponse typedResponse = new TransferEventResponse();
        typedResponse.log = log;
        typedResponse.from = (String) eventValues.getIndexedValues().get(0).getValue();
        typedResponse.to = (String) eventValues.getIndexedValues().get(1).getValue();
        typedResponse.value = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
        return typedResponse;
    }

    public Flowable<TransferEventResponse> transferEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(log -> getTransferEventFromLog(log));
    }

    public Flowable<TransferEventResponse> transferEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(TRANSFER_EVENT));
        return transferEventFlowable(filter);
    }

    public static List<UpgradedEventResponse> getUpgradedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = staticExtractEventParametersWithLog(UPGRADED_EVENT, transactionReceipt);
        ArrayList<UpgradedEventResponse> responses = new ArrayList<UpgradedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            UpgradedEventResponse typedResponse = new UpgradedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.implementation = (String) eventValues.getIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public static UpgradedEventResponse getUpgradedEventFromLog(Log log) {
        Contract.EventValuesWithLog eventValues = staticExtractEventParametersWithLog(UPGRADED_EVENT, log);
        UpgradedEventResponse typedResponse = new UpgradedEventResponse();
        typedResponse.log = log;
        typedResponse.implementation = (String) eventValues.getIndexedValues().get(0).getValue();
        return typedResponse;
    }

    public Flowable<UpgradedEventResponse> upgradedEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(log -> getUpgradedEventFromLog(log));
    }

    public Flowable<UpgradedEventResponse> upgradedEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(UPGRADED_EVENT));
        return upgradedEventFlowable(filter);
    }

    public RemoteFunctionCall<String> UPGRADE_INTERFACE_VERSION() {
        final Function function = new Function(FUNC_UPGRADE_INTERFACE_VERSION, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<BigInteger> allowance(String owner, String spender) {
        final Function function = new Function(FUNC_ALLOWANCE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, owner), 
                new org.web3j.abi.datatypes.Address(160, spender)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<TransactionReceipt> approve(String spender, BigInteger value) {
        final Function function = new Function(
                FUNC_APPROVE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, spender), 
                new org.web3j.abi.datatypes.generated.Uint256(value)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<BigInteger> balanceOf(String account) {
        final Function function = new Function(FUNC_BALANCEOF, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, account)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<TransactionReceipt> burn(BigInteger value) {
        final Function function = new Function(
                FUNC_BURN, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(value)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> burnFrom(String account, BigInteger value) {
        final Function function = new Function(
                FUNC_BURNFROM, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, account), 
                new org.web3j.abi.datatypes.generated.Uint256(value)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<BigInteger> decimals() {
        final Function function = new Function(FUNC_DECIMALS, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint8>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<TransactionReceipt> initialize(String initialOwner) {
        final Function function = new Function(
                FUNC_INITIALIZE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, initialOwner)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<String> name() {
        final Function function = new Function(FUNC_NAME, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<String> owner() {
        final Function function = new Function(FUNC_OWNER, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<byte[]> proxiableUUID() {
        final Function function = new Function(FUNC_PROXIABLEUUID, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bytes32>() {}));
        return executeRemoteCallSingleValueReturn(function, byte[].class);
    }

    public RemoteFunctionCall<TransactionReceipt> renounceOwnership() {
        final Function function = new Function(
                FUNC_RENOUNCEOWNERSHIP, 
                Arrays.<Type>asList(), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<String> symbol() {
        final Function function = new Function(FUNC_SYMBOL, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<BigInteger> totalSupply() {
        final Function function = new Function(FUNC_TOTALSUPPLY, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<TransactionReceipt> transfer(String to, BigInteger value) {
        final Function function = new Function(
                FUNC_TRANSFER, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, to), 
                new org.web3j.abi.datatypes.generated.Uint256(value)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> transferFrom(String from, String to, BigInteger value) {
        final Function function = new Function(
                FUNC_TRANSFERFROM, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, from), 
                new org.web3j.abi.datatypes.Address(160, to), 
                new org.web3j.abi.datatypes.generated.Uint256(value)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> transferOwnership(String newOwner) {
        final Function function = new Function(
                FUNC_TRANSFEROWNERSHIP, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, newOwner)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> upgradeToAndCall(String newImplementation, byte[] data, BigInteger weiValue) {
        final Function function = new Function(
                FUNC_UPGRADETOANDCALL, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, newImplementation), 
                new org.web3j.abi.datatypes.DynamicBytes(data)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function, weiValue);
    }

    @Deprecated
    public static NexGami load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new NexGami(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static NexGami load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new NexGami(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static NexGami load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new NexGami(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static NexGami load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new NexGami(contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<NexGami> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return deployRemoteCall(NexGami.class, web3j, credentials, contractGasProvider, getDeploymentBinary(), "");
    }

    public static RemoteCall<NexGami> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return deployRemoteCall(NexGami.class, web3j, transactionManager, contractGasProvider, getDeploymentBinary(), "");
    }

    @Deprecated
    public static RemoteCall<NexGami> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(NexGami.class, web3j, credentials, gasPrice, gasLimit, getDeploymentBinary(), "");
    }

    @Deprecated
    public static RemoteCall<NexGami> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(NexGami.class, web3j, transactionManager, gasPrice, gasLimit, getDeploymentBinary(), "");
    }

    public static void linkLibraries(List<Contract.LinkReference> references) {
        librariesLinkedBinary = linkBinaryWithReferences(BINARY, references);
    }

    public static String getDeploymentBinary() {
        if (librariesLinkedBinary != null) {
            return librariesLinkedBinary;
        } else {
            return BINARY;
        }
    }

    public static class ApprovalEventResponse extends BaseEventResponse {
        public String owner;

        public String spender;

        public BigInteger value;
    }

    public static class InitializedEventResponse extends BaseEventResponse {
        public BigInteger version;
    }

    public static class OwnershipTransferredEventResponse extends BaseEventResponse {
        public String previousOwner;

        public String newOwner;
    }

    public static class TransferEventResponse extends BaseEventResponse {
        public String from;

        public String to;

        public BigInteger value;
    }

    public static class UpgradedEventResponse extends BaseEventResponse {
        public String implementation;
    }
}
