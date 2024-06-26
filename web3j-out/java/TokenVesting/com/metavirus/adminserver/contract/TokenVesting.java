package com.metavirus.adminserver.contract;

import io.reactivex.Flowable;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import org.web3j.abi.EventEncoder;
import org.web3j.abi.FunctionEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Address;
import org.web3j.abi.datatypes.Event;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.generated.Uint256;
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
 * <p>Generated with web3j version 1.5.0.
 */
@SuppressWarnings("rawtypes")
public class TokenVesting extends Contract {
    public static final String BINARY = "0x60806040523480156200001157600080fd5b5060405162001655380380620016558339818101604052810190620000379190620004b6565b33600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000ad5760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000a491906200054f565b60405180910390fd5b620000be816200030860201b60201c565b5060018081905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff160362000138576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200012f90620005f3565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603620001aa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001a1906200068b565b60405180910390fd5b60008111620001f0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001e790620006fd565b60405180910390fd5b8082111562000236576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200022d9062000795565b60405180910390fd5b84600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826004819055508183620002cd9190620007e6565b6005819055508060068190555060006007819055506000600860006101000a81548160ff021916908315150217905550505050505062000821565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003fe82620003d1565b9050919050565b60006200041282620003f1565b9050919050565b620004248162000405565b81146200043057600080fd5b50565b600081519050620004448162000419565b92915050565b6200045581620003f1565b81146200046157600080fd5b50565b60008151905062000475816200044a565b92915050565b6000819050919050565b62000490816200047b565b81146200049c57600080fd5b50565b600081519050620004b08162000485565b92915050565b600080600080600060a08688031215620004d557620004d4620003cc565b5b6000620004e58882890162000433565b9550506020620004f88882890162000464565b94505060406200050b888289016200049f565b93505060606200051e888289016200049f565b925050608062000531888289016200049f565b9150509295509295909350565b6200054981620003f1565b82525050565b60006020820190506200056660008301846200053e565b92915050565b600082825260208201905092915050565b7f546f6b656e56657374696e673a20746f6b656e20697320746865207a65726f2060008201527f6164647265737300000000000000000000000000000000000000000000000000602082015250565b6000620005db6027836200056c565b9150620005e8826200057d565b604082019050919050565b600060208201905081810360008301526200060e81620005cc565b9050919050565b7f546f6b656e56657374696e673a2062656e65666963696172792069732074686560008201527f207a65726f206164647265737300000000000000000000000000000000000000602082015250565b600062000673602d836200056c565b9150620006808262000615565b604082019050919050565b60006020820190508181036000830152620006a68162000664565b9050919050565b7f546f6b656e56657374696e673a206475726174696f6e20697320300000000000600082015250565b6000620006e5601b836200056c565b9150620006f282620006ad565b602082019050919050565b600060208201905081810360008301526200071881620006d6565b9050919050565b7f546f6b656e56657374696e673a20636c696666206973206c6f6e67657220746860008201527f616e206475726174696f6e000000000000000000000000000000000000000000602082015250565b60006200077d602b836200056c565b91506200078a826200071f565b604082019050919050565b60006020820190508181036000830152620007b0816200076e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620007f3826200047b565b915062000800836200047b565b92508282019050808211156200081b576200081a620007b7565b5b92915050565b610e2480620008316000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80635b94008114610067578063715018a61461008557806386d1a69f1461008f5780638da5cb5b14610099578063b6549f75146100b7578063f2fde38b146100c1575b600080fd5b61006f6100dd565b60405161007c9190610958565b60405180910390f35b61008d6100ec565b005b610097610100565b005b6100a16102e9565b6040516100ae91906109b4565b60405180910390f35b6100bf610312565b005b6100db60048036038101906100d69190610a00565b610547565b005b60006100e76105cd565b905090565b6100f4610608565b6100fe600061068f565b565b610108610753565b600860009054906101000a900460ff1615610158576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014f90610ab0565b60405180910390fd5b60006101626105cd565b9050600081116101a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019e90610b1c565b60405180910390fd5b80600760008282546101b99190610b6b565b92505081905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161023f929190610b9f565b6020604051808303816000875af115801561025e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102829190610c00565b507fc7798891864187665ac6dd119286e44ec13f014527aeeb2b8eb3fd413df93179600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16826040516102d6929190610b9f565b60405180910390a1506102e7610799565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61031a610608565b600860009054906101000a900460ff161561036a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036190610c9f565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016103c791906109b4565b602060405180830381865afa1580156103e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104089190610ceb565b905060006104146105cd565b9050600081836104249190610d18565b90506001600860006101000a81548160ff021916908315150217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6104876102e9565b836040518363ffffffff1660e01b81526004016104a5929190610b9f565b6020604051808303816000875af11580156104c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e89190610c00565b507f39983c6d4d174a7aee564f449d4a5c3c7ac9649d72b7793c56901183996f8af6600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405161053a91906109b4565b60405180910390a1505050565b61054f610608565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105c15760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016105b891906109b4565b60405180910390fd5b6105ca8161068f565b50565b6000600860009054906101000a900460ff16156105ed5760009050610605565b6007546105f86107a2565b6106029190610d18565b90505b90565b610610610937565b73ffffffffffffffffffffffffffffffffffffffff1661062e6102e9565b73ffffffffffffffffffffffffffffffffffffffff161461068d57610651610937565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161068491906109b4565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60026001540361078f576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600181905550565b60018081905550565b60006005544210156107b75760009050610934565b6006546004546107c79190610b6b565b421061087057600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161082891906109b4565b602060405180830381865afa158015610845573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108699190610ceb565b9050610934565b600654600454426108819190610d18565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016108dc91906109b4565b602060405180830381865afa1580156108f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091d9190610ceb565b6109279190610d4c565b6109319190610dbd565b90505b90565b600033905090565b6000819050919050565b6109528161093f565b82525050565b600060208201905061096d6000830184610949565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061099e82610973565b9050919050565b6109ae81610993565b82525050565b60006020820190506109c960008301846109a5565b92915050565b600080fd5b6109dd81610993565b81146109e857600080fd5b50565b6000813590506109fa816109d4565b92915050565b600060208284031215610a1657610a156109cf565b5b6000610a24848285016109eb565b91505092915050565b600082825260208201905092915050565b7f546f6b656e56657374696e673a20636f6e7472616374206973207265766f6b6560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b6000610a9a602183610a2d565b9150610aa582610a3e565b604082019050919050565b60006020820190508181036000830152610ac981610a8d565b9050919050565b7f546f6b656e56657374696e673a206e6f20746f6b656e73206172652064756500600082015250565b6000610b06601f83610a2d565b9150610b1182610ad0565b602082019050919050565b60006020820190508181036000830152610b3581610af9565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b768261093f565b9150610b818361093f565b9250828201905080821115610b9957610b98610b3c565b5b92915050565b6000604082019050610bb460008301856109a5565b610bc16020830184610949565b9392505050565b60008115159050919050565b610bdd81610bc8565b8114610be857600080fd5b50565b600081519050610bfa81610bd4565b92915050565b600060208284031215610c1657610c156109cf565b5b6000610c2484828501610beb565b91505092915050565b7f546f6b656e56657374696e673a20636f6e747261637420616c7265616479207260008201527f65766f6b65640000000000000000000000000000000000000000000000000000602082015250565b6000610c89602683610a2d565b9150610c9482610c2d565b604082019050919050565b60006020820190508181036000830152610cb881610c7c565b9050919050565b610cc88161093f565b8114610cd357600080fd5b50565b600081519050610ce581610cbf565b92915050565b600060208284031215610d0157610d006109cf565b5b6000610d0f84828501610cd6565b91505092915050565b6000610d238261093f565b9150610d2e8361093f565b9250828203905081811115610d4657610d45610b3c565b5b92915050565b6000610d578261093f565b9150610d628361093f565b9250828202610d708161093f565b91508282048414831517610d8757610d86610b3c565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610dc88261093f565b9150610dd38361093f565b925082610de357610de2610d8e565b5b82820490509291505056fea2646970667358221220f4e02084e8555f84f44e9082352ea05189630b922c4412020606fa9312529b6964736f6c63430008170033";

    public static final String FUNC_OWNER = "owner";

    public static final String FUNC_RELEASABLEAMOUNT = "releasableAmount";

    public static final String FUNC_RELEASE = "release";

    public static final String FUNC_RENOUNCEOWNERSHIP = "renounceOwnership";

    public static final String FUNC_REVOKE = "revoke";

    public static final String FUNC_TRANSFEROWNERSHIP = "transferOwnership";

    public static final Event OWNERSHIPTRANSFERRED_EVENT = new Event("OwnershipTransferred", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}));
    ;

    public static final Event TOKENVESTINGREVOKED_EVENT = new Event("TokenVestingRevoked", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
    ;

    public static final Event TOKENSRELEASED_EVENT = new Event("TokensReleased", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}, new TypeReference<Uint256>() {}));
    ;

    @Deprecated
    protected TokenVesting(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected TokenVesting(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected TokenVesting(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected TokenVesting(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
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

    public static List<TokenVestingRevokedEventResponse> getTokenVestingRevokedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = staticExtractEventParametersWithLog(TOKENVESTINGREVOKED_EVENT, transactionReceipt);
        ArrayList<TokenVestingRevokedEventResponse> responses = new ArrayList<TokenVestingRevokedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            TokenVestingRevokedEventResponse typedResponse = new TokenVestingRevokedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.token = (String) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public static TokenVestingRevokedEventResponse getTokenVestingRevokedEventFromLog(Log log) {
        Contract.EventValuesWithLog eventValues = staticExtractEventParametersWithLog(TOKENVESTINGREVOKED_EVENT, log);
        TokenVestingRevokedEventResponse typedResponse = new TokenVestingRevokedEventResponse();
        typedResponse.log = log;
        typedResponse.token = (String) eventValues.getNonIndexedValues().get(0).getValue();
        return typedResponse;
    }

    public Flowable<TokenVestingRevokedEventResponse> tokenVestingRevokedEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(log -> getTokenVestingRevokedEventFromLog(log));
    }

    public Flowable<TokenVestingRevokedEventResponse> tokenVestingRevokedEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(TOKENVESTINGREVOKED_EVENT));
        return tokenVestingRevokedEventFlowable(filter);
    }

    public static List<TokensReleasedEventResponse> getTokensReleasedEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = staticExtractEventParametersWithLog(TOKENSRELEASED_EVENT, transactionReceipt);
        ArrayList<TokensReleasedEventResponse> responses = new ArrayList<TokensReleasedEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            TokensReleasedEventResponse typedResponse = new TokensReleasedEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.token = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse.amount = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public static TokensReleasedEventResponse getTokensReleasedEventFromLog(Log log) {
        Contract.EventValuesWithLog eventValues = staticExtractEventParametersWithLog(TOKENSRELEASED_EVENT, log);
        TokensReleasedEventResponse typedResponse = new TokensReleasedEventResponse();
        typedResponse.log = log;
        typedResponse.token = (String) eventValues.getNonIndexedValues().get(0).getValue();
        typedResponse.amount = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
        return typedResponse;
    }

    public Flowable<TokensReleasedEventResponse> tokensReleasedEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(log -> getTokensReleasedEventFromLog(log));
    }

    public Flowable<TokensReleasedEventResponse> tokensReleasedEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(TOKENSRELEASED_EVENT));
        return tokensReleasedEventFlowable(filter);
    }

    public RemoteFunctionCall<String> owner() {
        final Function function = new Function(FUNC_OWNER, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<BigInteger> releasableAmount() {
        final Function function = new Function(FUNC_RELEASABLEAMOUNT, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<TransactionReceipt> release() {
        final Function function = new Function(
                FUNC_RELEASE, 
                Arrays.<Type>asList(), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> renounceOwnership() {
        final Function function = new Function(
                FUNC_RENOUNCEOWNERSHIP, 
                Arrays.<Type>asList(), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> revoke() {
        final Function function = new Function(
                FUNC_REVOKE, 
                Arrays.<Type>asList(), 
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

    @Deprecated
    public static TokenVesting load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new TokenVesting(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static TokenVesting load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new TokenVesting(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static TokenVesting load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new TokenVesting(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static TokenVesting load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new TokenVesting(contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<TokenVesting> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider, String token_, String beneficiary_, BigInteger start_, BigInteger cliffDuration_, BigInteger duration_) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, token_), 
                new org.web3j.abi.datatypes.Address(160, beneficiary_), 
                new org.web3j.abi.datatypes.generated.Uint256(start_), 
                new org.web3j.abi.datatypes.generated.Uint256(cliffDuration_), 
                new org.web3j.abi.datatypes.generated.Uint256(duration_)));
        return deployRemoteCall(TokenVesting.class, web3j, credentials, contractGasProvider, BINARY, encodedConstructor);
    }

    public static RemoteCall<TokenVesting> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider, String token_, String beneficiary_, BigInteger start_, BigInteger cliffDuration_, BigInteger duration_) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, token_), 
                new org.web3j.abi.datatypes.Address(160, beneficiary_), 
                new org.web3j.abi.datatypes.generated.Uint256(start_), 
                new org.web3j.abi.datatypes.generated.Uint256(cliffDuration_), 
                new org.web3j.abi.datatypes.generated.Uint256(duration_)));
        return deployRemoteCall(TokenVesting.class, web3j, transactionManager, contractGasProvider, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<TokenVesting> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String token_, String beneficiary_, BigInteger start_, BigInteger cliffDuration_, BigInteger duration_) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, token_), 
                new org.web3j.abi.datatypes.Address(160, beneficiary_), 
                new org.web3j.abi.datatypes.generated.Uint256(start_), 
                new org.web3j.abi.datatypes.generated.Uint256(cliffDuration_), 
                new org.web3j.abi.datatypes.generated.Uint256(duration_)));
        return deployRemoteCall(TokenVesting.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<TokenVesting> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String token_, String beneficiary_, BigInteger start_, BigInteger cliffDuration_, BigInteger duration_) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, token_), 
                new org.web3j.abi.datatypes.Address(160, beneficiary_), 
                new org.web3j.abi.datatypes.generated.Uint256(start_), 
                new org.web3j.abi.datatypes.generated.Uint256(cliffDuration_), 
                new org.web3j.abi.datatypes.generated.Uint256(duration_)));
        return deployRemoteCall(TokenVesting.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public static class OwnershipTransferredEventResponse extends BaseEventResponse {
        public String previousOwner;

        public String newOwner;
    }

    public static class TokenVestingRevokedEventResponse extends BaseEventResponse {
        public String token;
    }

    public static class TokensReleasedEventResponse extends BaseEventResponse {
        public String token;

        public BigInteger amount;
    }
}
