pragma solidity >0.4.0 <=0.6.0;
pragma experimental ABIEncoderV2;

contract fairswap{

	bytes32 public EncInputRoot;  //root of encryptedGateOutputs
	uint public price;						//price of the digital good
	bytes32 public commitmentOfKey; // commitment of the key
	bytes32 public key;			//key
	bytes32 public receiverEntryKey;  //the entry key for the receiver
	uint public Now; //to get the current time


	//the mapping is between the gate index and the operation of the gate
	mapping(uint => uint) public circuitGatesOperation;

	//address of the two parties
	address payable public  senderAddress;
	address payable public  receiverAddress;

	//deadline for completing a stage
  uint public timeout;

	//stages of fairswap
	enum stage{created,initialized,accepted,keyRevealed,finishedAndBothHonestParties,finishedAndMaliciousSender,finishedAndMaliciousReciever}
	//created - sender creates the Contract
	//initialized - sender initializes the state variables
	//stages created and initialized happens together when contract is deployed
	//accepted - receiver has paid the price
	//keyRevealed - sender has revealed the key
	//finishedAndBothHonestParties - fairswap completed and both parties were honest
	//finishedAndMaliciousSender - fairswap not completed as sender was malicious
	//finishedAndMaliciousReciever - fairswap not completed as receiver was malicious

	stage public phase = stage.created;

	//fn used to move to the next stage fron the current stage
	function nextStage() public{
		phase=stage(uint(phase)+1);
		timeout=now/10+ 10 seconds;
	}

	//state variable initilization including the circuit gate's operation
	//stage goes from created to initialized
	constructor(uint _price,bytes32 _commitmentOfKey,bytes32  _EncInputRoot,bytes32 _receiverEntryKey,uint[] memory	 circuitGatesOperationArray) public{

		senderAddress=msg.sender;
		price = _price;
		commitmentOfKey = _commitmentOfKey;
		EncInputRoot = _EncInputRoot;
		for(uint i=0;i<circuitGatesOperationArray.length;i++)
		{
			circuitGatesOperation[i]=circuitGatesOperationArray[i];
		}
		receiverEntryKey=_receiverEntryKey;
		nextStage();
	}

	//the receiver with the help of entry key initializes the receiverAddress.
  function initializeRecieverAddress(bytes32  recEntryKey) public{
		require(receiverEntryKey==recEntryKey);
		receiverAddress=msg.sender;

  }

	//modifier to check the eligibility of party, time and phase for a function execution
	modifier allowed(address _party,stage _stage){
		require(phase==_stage);
		require(now/10<timeout);
		require(msg.sender == _party);
		_;
	}

	//receiver pays the price for the digital good
	//If receiver pays the required ether,then the stage goes from initialized to accepted.
	//In the case when receiver pays lesser amount,the ehter is transferred back to the receiver
	//and the stage goes from initialized to finishedAndMaliciousReciever
	function accept() allowed(receiverAddress,stage.initialized) payable public {

		if(msg.value < price)
		{
			msg.sender.transfer(msg.value);
			phase=stage.finishedAndMaliciousReciever;
		}
		else
		{
			nextStage();
		}
	}

	//sender reveals the key
	//If the sender reveals the right key then the contract stage goes from accepted to keyRevealed
	//If the sender reveals the wrong key, then the ether is transferred to the receiver and the stage
	//goes from accepted to finishedAndMaliciousSender
	function revealKey (bytes32  _key) allowed(senderAddress, stage.accepted) public
	{

		bytes32 Hashed;
		bytes memory toBeHashed;
		toBeHashed=new bytes (32);
		for (uint i=0;i<32;i++)
		{
				toBeHashed[i]=_key[i];
		}
		Hashed=keccak256(toBeHashed);

		if(Hashed!=commitmentOfKey){
			receiverAddress.transfer(price);
			phase=stage.finishedAndMaliciousSender;
		}
		else
		{
			key=_key;
			nextStage();
		}
	}

	//the fucntion is for receiver to register a complain
	//if the receiver gives a valid complain them the ether gets transferred to the receiver and the contract stage
	//goes from keyRevealed to finishedAndMaliciousSender
	//In the case when the receiver gives an invalid complain, then the ether is transferred to the sender and the stage ///goes from keyRevealed to finishedAndMaliciousReciever
 	function complain(bytes32[][] memory complaint,uint[] memory indices) allowed(receiverAddress,stage.keyRevealed) public returns(bool)
	{

		bytes32 operationValue;
		uint operation;
		bytes32[] memory inputVectorsToTheGate;
		bytes32 Out;
		uint index;
		index=indices[0];
		inputVectorsToTheGate = new bytes32[] (complaint.length-1);

		if(!Mverify(complaint[0],indices[0]))
		{
			phase=stage.finishedAndMaliciousReciever;
			senderAddress.transfer(price);
			return false;
		}

		operation=circuitGatesOperation[indices[0]];
		Out=Dec(complaint[0][0]);

		for(uint i=0;i<complaint.length-1;i++)
		{
			if(!Mverify(complaint[i],indices[i]))
			{
				phase=stage.finishedAndMaliciousReciever;
				senderAddress.transfer(price);
				return false;
			}
			inputVectorsToTheGate[i]=Dec(complaint[i+1][0]);
		}
		operationValue=Operation(operation,inputVectorsToTheGate);

		if(Out!=operationValue)
		{

			phase=stage.finishedAndMaliciousSender;
			receiverAddress.transfer(price);
			return true;
		}

		phase=stage.finishedAndMaliciousReciever;
		senderAddress.transfer(price);
		return false;


	}

	//Operation function is circuit specific
	//This function is implemented for the circuit given in the repository.
	//pure functon
	function Operation(uint operation,bytes32[] memory inputVectorsToTheGate) internal pure returns(bytes32)
	{
		bytes32 Hashed;
			if(operation==2)
			{
				bytes memory toBeHashed;

				toBeHashed=new bytes (64);
				for (uint i=0;i<32;i++)
				{
					toBeHashed[i]=inputVectorsToTheGate[0][i];
				}
				for (uint i=0;i<32;i++)
				{
					toBeHashed[i+32]=inputVectorsToTheGate[1][i];
				}
				Hashed=keccak256(toBeHashed);
			}
			return Hashed;
	}

	//The function is used to check whether the individual gate vectors associated with the complaint is actually what the sender sent to the receiver..
	//view function...takes two vectors at a time and hashes them...in this way constructs the tree until the root and /////compares it with the EncInputRoot
	function  Mverify(bytes32[] memory complaint,uint _index)  public view returns(bool)
	{
		bytes32 _value;
		_value=complaint[0];
		bytes32 Hashed;
		uint i ;
		uint depth=5;
		bytes memory toBeHashed;
		toBeHashed=new bytes (64);

		for(i=0; i<depth-1; i++)
		{
			//if condition is used to decide whether to append the _value with the complaint vector or append complaint vector
			//with the _value
			if ((_index&(1<<i))>>i == 1)
			{

				for (uint j=0;j<32;j++)
				{
					toBeHashed[j]=complaint[i+1][j];
				}
				for (uint j=0;j<32;j++)
				{
					toBeHashed[j+32]=_value[j];
				}


				Hashed=keccak256(toBeHashed);
				_value=Hashed;
			}
			else
			{
				for (uint j=0;j<32;j++)
				{
					toBeHashed[j]=_value[j];
				}
				for (uint j=0;j<32;j++)
				{
					toBeHashed[j+32]=complaint[i+1][j];
				}

				Hashed=keccak256(toBeHashed);
				_value=Hashed;
			}
		}
		return (_value==EncInputRoot);
	}


	//used to Dccrypt the encryptedtext with the key
	function Dec(bytes32 encryptedtext) internal view returns(bytes32)
	{
		bytes memory toBeHashedKey;
		bytes32 keyHashed;
		bytes32 decodedtext;

		toBeHashedKey=new bytes (33);
		for(uint i=0;i<32;i++){

			toBeHashedKey[i]=key[i];

		}
		toBeHashedKey[32]=0;
		keyHashed=keccak256(toBeHashedKey);
		decodedtext=keyHashed^encryptedtext;

		return decodedtext;
	}

	//fn is called by sender after the faiswap completion to get the ether
	function senderGetEther() allowed(senderAddress,stage.keyRevealed) public{
		require(now/10>timeout);
		senderAddress.transfer(price);
		nextStage();

	}

	//function is called by receiver in the accepted stage if the receiver doesn't reveal the key within the time. The
	//ether is transferred to the receiver
	function receiverGetEther() allowed(receiverAddress,stage.accepted) public{
		require(now/10>timeout);
		receiverAddress.transfer(price);
		phase=stage.finishedAndMaliciousSender;

	}


}
