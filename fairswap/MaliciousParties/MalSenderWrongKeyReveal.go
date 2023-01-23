package main

import(
	"fmt"
	kec "github.com/ethereum/go-ethereum/crypto"
	"os"
	"io"
	"log"
	"encoding/json"
	//"reflect"

	"time"
  "context"
  "crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core"
	"crypto/rand"
  "github.com/ethereum/go-ethereum/common"
	"math/big"
	mathlib "math"
   import_contracts "./contracts"
)

const noOfInputGates=8 //total number of input gates
const buffer_size=32   //byte array size (size of each byte array sent from sender to the receiver)
const totalNumberOfGates=16  //total number of gates in the entire circuit
const maxLinesToGate=2			//maximum number of inputs a gate can take
const totNumOfEncOutVecToMer=16	//total number of encrypted gate outputs(equals totalNumberOfGates)
const waitingTime=5   //max time to complete a stage
const keySize=32			//key size
const hashFunctionOutputBitSize=32 //hash function's output size(32 bytes in this case)

var contractAddress common.Address

//these struct and SetCircuitTuples function is common to both sender and receiver
type circuit struct{
	Index int
	Operation int
	InputToTheGate []int
}

type SenderToReceiverStruct struct {
		Id int                                        //id
		Keycommit [hashFunctionOutputBitSize]byte														//key commitment
		EncodedOutputOfGates [totNumOfEncOutVecToMer][hashFunctionOutputBitSize]byte		//Encoded output of gates
		ReceiverEntryKey [keySize]byte
}

type SenderToContractStruct struct{
	Id int
	Price big.Int
	Keycommit [hashFunctionOutputBitSize]byte
	//MerkleRootOfCircuit [hashFunctionOutputBitSize]byte
	MerkleRootOfEncInp [hashFunctionOutputBitSize]byte
	ConReceiverEntryKey [keySize]byte
}

func main(){

		//channels
		channel_SenToRec:=make(chan []byte)				//sender sends the byterarray(of Id,Keycommit,EncodedOutputOfGates)
		channel_RecToMainIni:=make(chan string)		//Receiver sends a notificaion once he has finished his entire tasks
		channel_RecToMainRev:=make(chan string)
		channel_SenToMain:=make(chan string)

		//sender address generation
		authSender,senderAddress := AuthAndAddressGeneration()
		fmt.Println("\n->Sender\nSender Address :",senderAddress)

		//receiver address generation
		authReceiver,receiverAddress := AuthAndAddressGeneration()
		fmt.Println("\n->Receiver\nReceiver Address :",receiverAddress)

		//blockchain creation
		alloc := make(core.GenesisAlloc)
		alloc[authSender.From] = core.GenesisAccount{Balance: big.NewInt(11000000000)}  //1.1 ether
		alloc[authReceiver.From] = core.GenesisAccount{Balance: big.NewInt(25000000000)} //2.5 ether
		//client
		client := backends.NewSimulatedBackend(alloc,60000000000)

		// /account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

		//Calculating sender and receiver initial account balances
		senderInitialBalance, err := client.BalanceAt(context.Background(),senderAddress, nil)
		if err != nil{
				log.Fatal(err)
		}
		fmt.Println("-->Main\nSenderInitialAccaountBalance : ",senderInitialBalance,"\n")

		receiverInitialBalance, err := client.BalanceAt(context.Background(),receiverAddress, nil)
		if err != nil{
				log.Fatal(err)
		}
		fmt.Println("-->Main\nReceiverInitialAccountBalance : ",receiverInitialBalance,"\n")


		//go routines
		go maliciousSender_WrongKeyReveal(authSender,client,channel_SenToRec,channel_SenToMain)
		go Receiver(authReceiver,client,channel_SenToRec,channel_RecToMainIni,channel_RecToMainRev)


		//channels... Parties communicate to the main function
		ReceicverInittilaizationMessage:=<-channel_RecToMainIni
		fmt.Println("\n->Main fn\n",ReceicverInittilaizationMessage)

		SenderMessageToMain:=<-channel_SenToMain
		fmt.Println("\n->Main fn\n",SenderMessageToMain)

		ReceiverRevealMessage:=<-channel_RecToMainRev
		fmt.Println("\n->Main fn\n",ReceiverRevealMessage)

		//getting the parties account balance at the end of protocol
		senderFinalBalance, err := client.BalanceAt(context.Background(),senderAddress, nil)
		if err != nil{
				log.Fatal(err)
		}

		senderDifferenceInFinalAndInitialBalance:=new(big.Int).Sub(senderInitialBalance,senderFinalBalance)
		fmt.Println("\n-->Main\nSenderFinalAccountBalance : ",senderFinalBalance,"\nSenderDifferenceInFinalAndInitialBalance : ",senderDifferenceInFinalAndInitialBalance)

		receiverFinalBalance, err := client.BalanceAt(context.Background(),receiverAddress, nil)
		if err != nil{
				log.Fatal(err)
		}
		receiverDifferenceInFinalAndInitialBalance:=new(big.Int).Sub(receiverInitialBalance,receiverFinalBalance)
		fmt.Println("\n-->Main\nReceiverFinalAccountBalance : ",receiverFinalBalance,"\nReceiverDifferenceInFinalAndInitialBalance : ",receiverDifferenceInFinalAndInitialBalance ,"\n")

}

//maliciousSender_WrongKeyReveal function
//The sender being malicious reveals a wrong key
//Input parameters:
//authsender(*bind.TransactOpts) - the sender in the blockchain
//client()*backends.SimulatedBackend) - blockchain(here it is a backend blockchain)
//channel_SenToRec(chan []byte) - to send the byte array to the receiver
//channel_SenToMain(chan string) - to communicate with the main function
func maliciousSender_WrongKeyReveal(authSender *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_SenToMain chan string){

	//reading from a file "a.txt" and storing bytearray of 32 bytes in inputVectors of size equal to noOfInputGates
	fmt.Println("\n->Sender\nBegins")
	var currentByte int64 = 0
	fileBuffer := make([]byte, buffer_size)
	file, err := os.Open("a.txt") // For read access.
	if err != nil {
			log.Fatal(err)
	}
	var inputVectors [noOfInputGates][buffer_size]byte
	count:=0
	for {
			n,err := file.ReadAt(fileBuffer, currentByte)
			currentByte += buffer_size
			fileBuffer=fileBuffer[:n]
			if count==noOfInputGates{
					break
			}
			for j:=0;j<n;j++{
				inputVectors[count][j]=fileBuffer[j]
			}
			if err == io.EOF {
					log.Fatal(err)
			}
			count++
	}
	file.Close()

	fmt.Println("->Sender\nInputVectors:",inputVectors)

	//objects
	var SenToConObject SenderToContractStruct
	var SenToRecObject SenderToReceiverStruct
	var circuitObjects [totalNumberOfGates]circuit //array of circuit object

	//sets the circuit
	setCircuitTuples(&circuitObjects)

	/*sender fns call*/
	key:=keyGenerate()
	keyCommit:=fnKeycommit(key)
	MRinp:=createMerkleTree(inputVectors)                          //MRinp (merkle root for input vectors)
	//MRcir:=createMerkleTreeForCir(circuitObjects)									//MRcir (merkle root for ciruit)
	encodedGateOutputs,_:=Encode(inputVectors,key,circuitObjects,MRinp)
	MRencout:=createMerkleTreeForEncInp(encodedGateOutputs)					//MRencout(merkle root for encoded outputs)
	recEntryKey:=keyGenerate()

	fmt.Println("\n->Sender \n Key:",key,"\n\nkeyCommit : ",keyCommit,"\n\nMRinp (merkle root for input vectors): ",MRinp,"\n\n MRencout(merkle root for encoded outputs) : ",MRencout)

	/*sentoconobject initialize*/
	idInitializingForContract(&SenToConObject)
	setPrice(&SenToConObject)
	SenToConObject.Keycommit=keyCommit
	//SenToConObject.MerkleRootOfCircuit=MRcir
	SenToConObject.MerkleRootOfEncInp=MRencout
	SenToConObject.ConReceiverEntryKey=recEntryKey

	//Putting gate operations in an array
	CircuitGateOperationArray:=[]*big.Int{}
	for i:=0;i<totalNumberOfGates;i++{
		op:=big.NewInt(int64(circuitObjects[i].Operation))
		CircuitGateOperationArray=append(CircuitGateOperationArray,op)
	}

	/*Deploy contract*/
	contractAddress, _,_, err=import_contracts.DeployFairswap(
		authSender,
		client,
		&SenToConObject.Price,
		SenToConObject.Keycommit,
		SenToConObject.MerkleRootOfEncInp,
		SenToConObject.ConReceiverEntryKey,
		CircuitGateOperationArray,

	)
	if err!=nil{
		log.Fatal(err)
	}
	client.Commit()

	fmt.Println("\n->Sender \n Contract Deployed")

	/*SenToRecObject initialize*/
	idInitializingForRec(&SenToRecObject)
	SenToRecObject.Keycommit=keyCommit
	SenToRecObject.EncodedOutputOfGates=encodedGateOutputs
	SenToRecObject.ReceiverEntryKey=recEntryKey


	//sending the encodedGateOutputs,keycommitment and receiver entry key to the receiver through a byte array(got after applying json marshal to the SenToRecObject)
	byteArrayFromSenToRec,err := json.Marshal(SenToRecObject)
	if err != nil {
		fmt.Println("error:", err)
	}
	channel_SenToRec<-byteArrayFromSenToRec
	fmt.Println("\n->Sender \nByteArray sent to Receiver")

	fmt.Println("\n->Sender\n Initialization Finished")

	//creating a contract instance
	instance,err:= import_contracts.NewFairswap(contractAddress,client)
	if err != nil {
			log.Fatal(err)
	}

	//The contract is in initialized stage
	//Once the contract is initialized, the receiver has time equal to waitingTime to accept and pay ether to the contract.
	//the for loop checks for the change in stage.
	//case 1 : If the receiver doesn't accept within the time limit, RevealedBool remains false after the for loop. As a result,
	//				"Receiver have not accepted within the timeout, Key not Revealed" is sent to the main function and the receiver doesn't
	//        proceed with the further protocol.
	//case 2 : If the receiver pays less than the price,then the contract stage goes to finishedAndMaliciousReciever.The sender sends
	//				"Malicious Receiver-Dind't pay enough ether" to the main function  and finishes the protocol execution here. The ether is sent back
	//				to the receiver by the contract.
	//case 3: The receiver accepts within the time. The contract stage goes to accepted and the sender has to reveal the key. RevealedBool becomes
	//				true.
	RevealedBool:=false
	for i:=0;i<waitingTime;i++ {
		Phase,err:=instance.Phase(nil)
		if err!=nil{
			log.Fatal(err)
		}
		if(Phase==2){
			RevealedBool=true
			break;
		}
		if(Phase==6){
			fmt.Println("\n->Sender\nReceiver has not paid enough ether")
			channel_SenToMain<-"Malicious Receiver-Dind't pay enough ether "
			break;
		}
		time.Sleep(time.Second *1)

	}

	if(!RevealedBool){
		channel_SenToMain<-"Receiver have not accepted within the timeout, Key not Revealed"

	}else{

					/*malicious act*/
					//Sender Reveals the  wrong key
					notContractKey:=keyGenerate()
					_,err=instance.RevealKey(&bind.TransactOpts{
							From:authSender.From,
							Signer:authSender.Signer,
							Value: nil,
					},notContractKey)
					if err!=nil{
						log.Fatal(err)
					}
					client.Commit()
					fmt.Println("\n->Sender\nReceiver have accepted within the timeout, Key Revealed")
			}

	//The contract is in keyRevealed stage
	//Once the sender has revealed the key, the receiver has time until waitingTime to either register a complaint or
	//walk away(not do any execution or transaction--got the correct file)
	//case 1 : If receiver gives a wrong complaint then the contract stage goes to finishedAndMaliciousReciever. The sender sends
	//				"Receiver Wrong complaint" to the main function and finishes the protocol.
	//case 2 : If the receiver doesn't respond, then the stage contract stage is still in keyRevealed and the sender gets the ether
	//				 after the waitingTime
	for i:=0;i<waitingTime;i++ {

			Phase,err:=instance.Phase(nil)
			if err!=nil{
				log.Fatal(err)
			}
			if(Phase==6){
				fmt.Println("\n->Sender\nReceiver has made a wrong complaint")
				channel_SenToMain<-"Receiver Wrong complaint "
				break;
			}
			time.Sleep(time.Second *1)
	}

	gasLimit := uint64(2100000000)

	//the sender getting the ether
	_,err=instance.SenderGetEther(&bind.TransactOpts{
			From:authSender.From,
			Signer:authSender.Signer,
			Value: nil,
			GasLimit:gasLimit,
		})
	if err!=nil{
				log.Fatal(err)
	}
	client.Commit()

	//the protocol gets finished and the sender sends the message to the main function
	channel_SenToMain<-"Fairswap completed.Sender Got the money"
}

//Receiver function
//The functoin executes honest receiver's role in fairswap
//Input parameters:
//authReceiver(*bind.TransactOpts) - the receiver in the blockchain
//client()*backends.SimulatedBackend) - blockchain(here it is a backend blockchain)
//channel_SenToRec(chan []byte) - to receiver the byte array from the sender
//channel_RecToMainIni(chan string) - to notify the main function for the completion of initialization phase
//channel_RecToMainRev(chan string) - to notify the main function for the completion of reveal phase
func Receiver(authReceiver *bind.TransactOpts,client *backends.SimulatedBackend,channel_SenToRec chan []byte,channel_RecToMainIni chan string,channel_RecToMainRev chan string){


		byteArrayFromSenToRec:=<-channel_SenToRec

		fmt.Println("\n->Receiver\nByteArray got from Sender ")

		var Receiver_SenToRecObj SenderToReceiverStruct
		err:=json.Unmarshal(byteArrayFromSenToRec, &Receiver_SenToRecObj)
		if err != nil {
				log.Fatal(err)
		}

		instance,err := import_contracts.NewFairswap(contractAddress,client)
		if err != nil {
				log.Fatal(err)
		}

		Receiver_MerkleRootOfEncInp,err:=instance.EncInputRoot(nil)
		if err!=nil{
			log.Fatal(err)
		}

		Receiver_EncodedOutputOfGates:=Receiver_SenToRecObj.EncodedOutputOfGates

		//comparing the merkle root of encodedGateOutputs got from contract and the calculated merkleroot from the encodedGateOutputs got from sender
		boolValueEncode:=CheckcreateMerkleTreeForEncInp(Receiver_EncodedOutputOfGates,Receiver_MerkleRootOfEncInp)

		fmt.Println("\n->Receiver\boolValueEncode(comparision between merkleroot of encodedGateOutputs from contract and calulated encodedGateOutputs merkleroot): ",boolValueEncode)

		//if boolValueEncode is false, then the receiver doesn't proceed with the protocol
		if(boolValueEncode){

				//the boolValueEncode was true, so the receiver calls the accept contract function and pays the required price
			_,err=instance.InitializeRecieverAddress(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value: nil,
			},Receiver_SenToRecObj.ReceiverEntryKey)
			if err!=nil{
					log.Fatal(err)
			}

			client.Commit()

			value:=big.NewInt(100)
			gasLimit := uint64(2100000000)

			_	,err=instance.Accept(&bind.TransactOpts{
				From:authReceiver.From,
				Signer:authReceiver.Signer,
				Value:value,
				GasLimit:gasLimit,
			})
			if err != nil {
					log.Fatal(err)
			}

			client.Commit()

			fmt.Println("\n->Receiver\n Finished Accepting and Initialization phase ends")
			//Sending the finished accepted stage message to the main function
			channel_RecToMainIni<-"Accepted stage got over"


			//The contract is in accepted stage
			//The sender
			//case 1 : The sender reveals the wrong key.Contract sents the ether to the receiver and the contract stage to finishedAndMaliciousSender
			//case 2 : The sender doesn't reveal the key within time. AcceptedBool remains false after the for loop. Receiver later calls the contract function
			//			   ReceiverGetEther to get the ether.
			//case 3 : The sender has revealed the correct key within time.The contract stage goes to keyRevealed and AcceptedBool is true after the for loop.
			AcceptedBool:=false
			var contract_Key [keySize]byte
			for i:=0;i<waitingTime;i++ {
				Phase,err:=instance.Phase(nil)
				if err!=nil{
					log.Fatal(err)
				}
				if(Phase==3){
					contract_Key,err=instance.Key(nil)
					if err!=nil{
						log.Fatal(err)
					}
					fmt.Println("\n->Receiver\ncontract_Key",contract_Key)
					AcceptedBool=true
					break;
				}
				if(Phase==5){
					fmt.Println("\n->Receiver\nSender Has revealed the wrong key")
					channel_RecToMainRev<-"Wrong key revealed"

				}
				time.Sleep(time.Second *1)

			}


			if(!AcceptedBool){

					fmt.Println("\n->Receiver\nSender has not revealed the key...ReceiverGetEther function called")
					gasLimit := uint64(2100000000)

					//the receiver getting the ether
					_,err=instance.ReceiverGetEther(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value: nil,
							GasLimit:gasLimit,
						})
					if err!=nil{
								log.Fatal(err)
					}
					client.Commit()

					channel_RecToMainRev<-"Sender hasn't revealed the key within the Timeout"

			} else{


				var circuitObjects [totalNumberOfGates]circuit  //array of circuit object
				//sets the circuit
				setCircuitTuples(&circuitObjects)


				fmt.Println("\n->Receiver\nSender has revealed the key....Proceeding further")

			 //Receiver_EncodedOutputOfGates[9][0]=3
				merkletree:=ReceiverMerkleTreeCreate(Receiver_EncodedOutputOfGates)

				//Extracting function being called
				complain,decodedOutputs,index:=Extract(Receiver_EncodedOutputOfGates,contract_Key,circuitObjects,merkletree,Receiver_MerkleRootOfEncInp)

				fmt.Println("\n->Receiver\nDecoded Ouptuts : ",decodedOutputs)

				//index nil denotes no complain
				if(index==nil){
					fmt.Println("\n->Receiver\nNo complain")
				} else{

					fmt.Println("complain :",complain)
					indexes:=[]*big.Int{}
					for i:=0;i<len(index);i++{
						b:=big.NewInt(int64(index[i]))
						indexes=append(indexes,b)
						}

						//complain function being called
						_,err=instance.Complain(&bind.TransactOpts{
							From:authReceiver.From,
							Signer:authReceiver.Signer,
							Value:nil,
							GasLimit:gasLimit,
						},complain,indexes)
						if err!=nil{
							log.Fatal(err)
						}
						client.Commit()

						//checking whether the given complaint is valid by verifying the stage changes.
						Phase,err:=instance.Phase(nil)
						if err!=nil{
							log.Fatal(err)
						}
						if(Phase==5){
							fmt.Println("\n->Receiver\nSender Has wrongly calulated the encrypted inputs")
							channel_RecToMainRev<-"Wrong encrypted inputs"

						}


				}
			}
			channel_RecToMainRev<-"Reveal Phase Finished"
		}	else{

				channel_RecToMainIni<-"Sender sent wrong encodedGateOutputs,not equal to contract encodedGateOutputs merkleroot,Not gone into accepted stage"
				channel_RecToMainRev<-"Sender sent wrong encodedGateOutputs,not equal to contract encodedGateOutputs merkleroot,Not gone into revealed phase"
		}
}

//generating a  computing party in the blockchain
func AuthAndAddressGeneration() (*bind.TransactOpts,common.Address){

	privateKey,err:=crypto.GenerateKey()
  if err!=nil{
    log.Fatal(err)
  }
  auth:=bind.NewKeyedTransactor(privateKey)

  publicKey := privateKey.Public()
  publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
  if !ok {
      log.Fatal("error casting public key to ECDSA")
  }
  address := crypto.PubkeyToAddress(*publicKeyECDSA)
  return auth,address
}


																				/*Sender's function*/


//function to model the circuit given....
func setCircuitTuples(circuitObjects *[totalNumberOfGates]circuit){

	//Giving tuple value to input gates
	for i:=0;i<noOfInputGates;i++{
			circuitObjects[i].Index = i
			circuitObjects[i].Operation =1
			circuitObjects[i].InputToTheGate = nil
	}

	//Giving tuple values to !inputgates && !lastgate
	k:=0
	for i:=noOfInputGates;i<totalNumberOfGates-1;i++ {
			circuitObjects[i].Index=i
			circuitObjects[i].Operation = 2
			for j:=0;j<maxLinesToGate;j++{
				circuitObjects[i].InputToTheGate=append(circuitObjects[i].InputToTheGate,k)
				k++
			}
	}

	//Last gate
	circuitObjects[totalNumberOfGates-1].Index = totalNumberOfGates-1
	circuitObjects[totalNumberOfGates-1].Operation = 3
	circuitObjects[totalNumberOfGates-1].InputToTheGate=append(circuitObjects[totalNumberOfGates-1].InputToTheGate,totalNumberOfGates-2)

}


/*Sender Functions*/

//key generate
func keyGenerate() ([keySize]byte){

		keyGen := make([]byte, keySize)
		var key [keySize]byte

		//Here
	  _, err := rand.Read(keyGen)
    if err != nil {
        // handle error here
    }

		for i:=0;i<len(keyGen);i++{
			key[i]=keyGen[i]
		}
		return key
}

//fn fnKeycommit
//commits(hash's) the key
//input parameters
//key : key to be commited
//returns the commitment of key
func fnKeycommit(key [keySize]byte) ([hashFunctionOutputBitSize]byte){

			var toBEHashedKey []byte
			var HashedKey [hashFunctionOutputBitSize]byte

			for j:=0;j<keySize;j++{
				toBEHashedKey=append(toBEHashedKey,key[j])
			}

			HashedAfter:=kec.Keccak256(toBEHashedKey)

			for j:=0;j<hashFunctionOutputBitSize;j++{
					HashedKey[j]=HashedAfter[j]
			}

			return HashedKey
}

//fn createMerkleTree
//creates a merkle tree and returns the merkleroot
func createMerkleTree(inputVectors [noOfInputGates][buffer_size]byte) ([hashFunctionOutputBitSize]byte){

	var merkletree [2*noOfInputGates-1][hashFunctionOutputBitSize]byte

	//input gates
	for i:=0;i<noOfInputGates;i++{
		for j:=0;j<buffer_size;j++{
				merkletree[i][j]=inputVectors[i][j]
		}
	}

	//!inputgates
	k:=0
	for i:=0;i<noOfInputGates-1;i++{

		var toBeHashed []byte
		var Hashed []byte

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed)

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i+noOfInputGates][j]=Hashed[j]
		}
		k=k+2
	}

	return merkletree[2*noOfInputGates-2]
}

//Function  createMerkleTreeForEncInp
//constructs the merkle tree for the enoded inputs and returns the merkleroot
//Input Parameters :
//encryptedGateOutputs : Vectors of encrypted output of the circuit gates
//Returns merkleroot
func createMerkleTreeForEncInp(encryptedGateOutputs [totNumOfEncOutVecToMer][hashFunctionOutputBitSize]byte) ([hashFunctionOutputBitSize]byte){

	var merkletree [2*totNumOfEncOutVecToMer-1][hashFunctionOutputBitSize]byte

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totNumOfEncOutVecToMer;i++{
		for j:=0;j<hashFunctionOutputBitSize;j++{
				merkletree[i][j]=encryptedGateOutputs[i][j]
		}
	}

	//!inputgates
	//merkle tree construction
	k:=0
	for i:=0;i<totNumOfEncOutVecToMer-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i+totNumOfEncOutVecToMer][j]=Hashed[j]
		}
		k=k+2
	}

	  //returns the merkleroot
	return merkletree[2*totNumOfEncOutVecToMer-2]
}

//Encode function
//This fn is used to encrypt the processed input vectors
//For input gates,no operation is performed by the gate on the input vector, it is just a "pass through" gate. So the input vectors are directly encrypted and stored.
//For the remaining higher hierarchical gates the output is calulated according to characteristic operation of a particular gate.And the output is then encrypted.

//The Encode function input parameters :
// inputVectors  :  The set of all input vectors
// key 					 :  The key with which the encryption function will be performed
//circuitObjects :  An array of circuit objects that contains info related to eacg gate
//MRx 					 :  Merkle root of input vectors.Specific to this particular circuit.
//Returns : An array of encrypted vectors
func Encode(inputVectors [noOfInputGates][buffer_size]byte,key [keySize]byte,circuitObjects [totalNumberOfGates]circuit,MRx [hashFunctionOutputBitSize]byte) ([totalNumberOfGates][hashFunctionOutputBitSize]byte,[totalNumberOfGates][32]byte)  {

	var out [totalNumberOfGates][hashFunctionOutputBitSize]byte     //stores the output of each gate
	var z [totalNumberOfGates][hashFunctionOutputBitSize]byte			//stores the encryption of each gate's output

	//Input Gates-No operation performed.Direct encryption of input vectors
	for i:=0;i<noOfInputGates;i++ {
			out[i]=inputVectors[i]
			z[i]=Enc(key,out[i])  //calls the encryption fn
	}

	//Not Input Gates
	//For each gate in this range,the input vectors to a particular gate are processed according to the operation and the ouput and encryption of this output is stored
	for i:=noOfInputGates;i<totalNumberOfGates;i++{
			Op:=circuitObjects[i].Operation  //Operation of gate by the index "i"

			//Operation 2: The gate takes in the output of two of its children, concatenates the two inputs and hashes it.Later the output is encrypted.
			if(Op==2) {
					leftsibling:=circuitObjects[i].InputToTheGate[0]
					rightsibling:=circuitObjects[i].InputToTheGate[1]
					var toBeHashed []byte
					for j:=0;j<hashFunctionOutputBitSize;j++ {
							toBeHashed= append(toBeHashed,out[leftsibling][j])
					}
					for j:=0;j<hashFunctionOutputBitSize;j++ {
							toBeHashed= append(toBeHashed,out[rightsibling][j])
					}
					hash := kec.Keccak256(toBeHashed)
					for j:=0;j<hashFunctionOutputBitSize;j++{
						out[i][j]=hash[j]
					}
			}

			//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
			if(Op==3) {
				for j:=0;j<hashFunctionOutputBitSize;j++{
					if (j==0){
							if(out[i-1][j]==MRx[j]){
								out[i][j]=1
							} else {
								out[i][j]=0
										 }
							} else {
								if(out[i-1][j]==MRx[j]){
									out[i][j]=0
								}	else{
									out[i][j]=1
								}
					}
				}
			}

			z[i]=Enc(key,out[i])  //encryption function is called

		}

	return z,out
}

//Encrypt function
//The function is used to encrypt(XOR) the given vector.
//Input Parameters
//key : The key with which the XORing is done.
//plainText : the input vector
//Returns the encrypt output
//Operation performed : plainText (XOR) key
func Enc(key [keySize]byte,plainText [32]byte) ([hashFunctionOutputBitSize]byte) {

			var keyPlusIndex []byte //appends the key(input parameter) and index "i"(range : 0 to Total Number of gates-1 )

			//In this specific case we are solving,the index("0") always remains the same.The funcion can be easily extended to accomadate varying index
			k:=0
			for i:=0;i<len(key);i++{
				keyPlusIndex=append(keyPlusIndex,key[i])
			}
			keyPlusIndex=append(keyPlusIndex,byte(k))
			key0 := kec.Keccak256(keyPlusIndex)


			var key032 [keySize]byte
			for i:=0;i<len(key0);i++{
				key032[i]=key0[i]
			}

			var encryptedtext [hashFunctionOutputBitSize]byte
			for i:=0;i<hashFunctionOutputBitSize;i++{
				encryptedtext[i]=key032[i]^plainText[i]  //xor operation
			}
			return encryptedtext
}

//initialize id
func idInitializingForRec(object *SenderToReceiverStruct){
	object.Id = 1
}

func idInitializingForContract(object *SenderToContractStruct){
	object.Id = 1
}

//set the price of the good
func setPrice(object *SenderToContractStruct){

	object.Price=*big.NewInt(80)     //800 wei
}


																														/*Receiver associated functions*/

//fn CheckcreateMerkleTreeForEncInp
//constructs merkletree and compares between the merkleroot got from construction and merkleroot passed as function parameter
//input parameters
//encryptedGateOutputs : Vectors of encrypted output of the circuit gates
//merklerootforEncInput : merkle root of Encrypted Input
//Returns bool value,true if merklerootforcircuit matches with the merkleroot got from construction,else false
func CheckcreateMerkleTreeForEncInp(encryptedGateOutputs [totNumOfEncOutVecToMer][hashFunctionOutputBitSize]byte, merklerootforEncInput [hashFunctionOutputBitSize]byte) (bool){

	var merkletree [2*totNumOfEncOutVecToMer-1][hashFunctionOutputBitSize]byte


	//inputgates
	for i:=0;i<totNumOfEncOutVecToMer;i++{
		for j:=0;j<hashFunctionOutputBitSize;j++{
				merkletree[i][j]=encryptedGateOutputs[i][j]
		}
	}


	//!inputgates
	k:=0
	for i:=0;i<totNumOfEncOutVecToMer-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}
		Hashed=kec.Keccak256(toBeHashed)
		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i+totNumOfEncOutVecToMer][j]=Hashed[j]
		}
		k=k+2
	}


	for j:=0;j<hashFunctionOutputBitSize;j++{
		if(merkletree[2*totNumOfEncOutVecToMer-2][j]!=merklerootforEncInput[j]){
			return false
		}
	}

	return true
}


//fn Extract
//The function is used to decrypt the encodedGateOutputs got from the sender. A complaint is generated if there was a malicious gate
//operation performed performed by  the sender.
//input parameters
//encodedGateOutputs([totNumOfEncOutVecToMer][32]byte) - the encodedGateOutputs byte array got from the sender
//key([32]byte) - key used to decrypt
//circuitObjects([totalNumberOfGates]circuit) - need tuple information for each gate to calulate the output of a gate
//merkletree()[2*totNumOfEncOutVecToMer-1][32]byte) - the merkle tree for encodedGateOutputs. Used to give the complaint vectors.
//Receiver_MerkleRootOfEncInp([32]byte) - This parameter is specific to the circuit considered here. The last gate in the circuit compares the
//																			Receiver_MerkleRootOfEncInp and the output from the penultimate gate.

func Extract(encodedGateOutputs [totNumOfEncOutVecToMer][hashFunctionOutputBitSize]byte,key [keySize]byte,circuitObjects [totalNumberOfGates]circuit,merkletree [2*totNumOfEncOutVecToMer-1][hashFunctionOutputBitSize]byte,Receiver_MerkleRootOfEncInp [hashFunctionOutputBitSize]byte) ([][][hashFunctionOutputBitSize]byte,[][buffer_size]byte,[]int) {

	var decodedOutputs [][32]byte  //to store all the decodedOutputs. Dynamic in size due to the fact that a wrong gate operation can lead to stopping the decoding process and
																//generating the complaint
	var complain [][][hashFunctionOutputBitSize]byte
	var out [32]byte
	var indices []int

	//Decryption of input vectors. Have taken the consumption that an error cannot be generated in the input gates. Even if the sender has something fishy in the
	//input gates, this effect will be refelected in the next level of gates.So, a relevant complaint can be given to the contract.
	for i:=0;i<noOfInputGates;i++ {
			out =Decrypt(key,encodedGateOutputs[i])
			decodedOutputs = append(decodedOutputs,out)
	}

	//Decryption of not input gates - gates above the base level.
	for i:=noOfInputGates;i<totalNumberOfGates;i++{
		noOFinputToThisParticularGate:=len(circuitObjects[i].InputToTheGate)

				//putting all the input vectors to a gate in a array
				var operationInputs [][hashFunctionOutputBitSize]byte
				for j:=0;j<noOFinputToThisParticularGate;j++{

								operationInputs=append(operationInputs,decodedOutputs[circuitObjects[i].InputToTheGate[j]])
				}

		out = Decrypt(key,encodedGateOutputs[i])
		decodedOutputs=append(decodedOutputs,out)

						//complaint generation
						//if decodedOutput of a gate doesn't match with the calulatec one. A complaint is created.
						//A complaint constists of merkleproof for encodedGateOutput(output of the gate) of  a particular gate and each individual encodedGateOutput
						//input to this gate.

						if decodedOutputs[i]!=Operation(circuitObjects[i].Operation,operationInputs,decodedOutputs,Receiver_MerkleRootOfEncInp){

						index:=i
						indices=append(indices,index)
						//merkle proof for the output of a complaint gate
						MproofTree:=Mproof(i,merkletree)
						complain=append(complain,MproofTree)

						//merkle proof for the inputs of a complaint gate and appending to the complain vector
						for k:=0;k<noOFinputToThisParticularGate;k++{
								MproofTree:=Mproof(circuitObjects[i].InputToTheGate[k],merkletree)
								complain=append(complain,MproofTree)
								indices=append(indices,circuitObjects[i].InputToTheGate[k])
						}
						return complain,decodedOutputs,indices
				}


	}
	//in this case there is no complain so indices in a nil vector. Entire encodedGateOutput is decoded.
	return complain,decodedOutputs,indices
}

//Decrypt function
//The function is used to decrypt(XOR) the given vector.
//Input Parameters
//key : The key with which the XORing is done.
//encryptedtext : the input vector
//Returns the plainText output
//Operation performed : encryptedtext (XOR) key
func Decrypt(key [keySize]byte,encryptedtext [hashFunctionOutputBitSize]byte) ([32]byte){
			var keyPlusIndex []byte

			k:=0
			for i:=0;i<len(key);i++{
				keyPlusIndex=append(keyPlusIndex,key[i])
			}
			keyPlusIndex=append(keyPlusIndex,byte(k))
			key0 := kec.Keccak256(keyPlusIndex)

			var key032 [keySize]byte
			for i:=0;i<len(key0);i++{
				key032[i]=key0[i]
			}

			var plainText [hashFunctionOutputBitSize]byte
			for i:=0;i<hashFunctionOutputBitSize;i++{
				plainText[i]=key032[i]^encryptedtext[i]
			}

			return plainText
}

//fn Mproof
//Input parameters
//inpdex(int) - the index of the gate
//merkletree([2*totNumOfEncOutVecToMer-1][32]byte) - merkle tree for encodedGateOutputs
//Return Parameters
//tree([depth][32]byte) - the merkle proof
func Mproof(index int,merkletree [2*totNumOfEncOutVecToMer-1][hashFunctionOutputBitSize]byte) ([][hashFunctionOutputBitSize]byte){

	depth:=int(mathlib.Log2(totalNumberOfGates))+1
	tree:=make([][32]byte,depth)
	tree[0]=merkletree[index]
	for i:=1;i<depth;i++{
		if(index%2==0){
			tree[i]=merkletree[index+1]
			index=index/2+totNumOfEncOutVecToMer
		} else {
			tree[i]=merkletree[index-1]
			index=index/2+totNumOfEncOutVecToMer
		}
	}
	return tree

}


func Operation(Op int,operationInputs [][hashFunctionOutputBitSize]byte,decodedOutputs [][32]byte,Receiver_MerkleRootOfEncInp [hashFunctionOutputBitSize]byte)  ([32]byte){

		var result [32]byte
		if(Op==2) {
				leftsibling:=operationInputs[0]
				rightsibling:=operationInputs[1]
				var toBeHashed []byte
				for j:=0;j<32;j++ {
						toBeHashed= append(toBeHashed,leftsibling[j])
				}
				for j:=0;j<32;j++ {
						toBeHashed= append(toBeHashed,rightsibling[j])
				}

				hashed := kec.Keccak256(toBeHashed)

				for j:=0;j<hashFunctionOutputBitSize;j++{
					result[j]=hashed[j]
				}
		}

		//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
		if(Op==3) {
			var merkletreeofdecodedinput1_8 [noOfInputGates][hashFunctionOutputBitSize]byte
			for k:=0;k<noOfInputGates;k++{
					merkletreeofdecodedinput1_8[k]=decodedOutputs[k]
			}
			merkleRootofDecodedInputVectors:=createMerkleTree(merkletreeofdecodedinput1_8)
			for j:=0;j<hashFunctionOutputBitSize;j++{
				if (j==0){
						if(operationInputs[0][j]==merkleRootofDecodedInputVectors[j]){
							result[j]=1
						} else {
							result[j]=0
									 }
						} else {
							if(operationInputs[0][j]==merkleRootofDecodedInputVectors[j]){
								result[j]=0
							}	else{
								result[j]=1
							}
				}
			}
		}
		return result

}

func ReceiverMerkleTreeCreate(encryptedGateOutputs [totNumOfEncOutVecToMer][hashFunctionOutputBitSize]byte) ([2*totNumOfEncOutVecToMer-1][hashFunctionOutputBitSize]byte){

	var merkletree [2*totNumOfEncOutVecToMer-1][hashFunctionOutputBitSize]byte

	//input gates
	//no hashing of the input vectors to the fn.
	for i:=0;i<totNumOfEncOutVecToMer;i++{
		for j:=0;j<hashFunctionOutputBitSize;j++{
				merkletree[i][j]=encryptedGateOutputs[i][j]
		}
	}

	//!inputgates
	//merkle tree construction
	k:=0
	for i:=0;i<totNumOfEncOutVecToMer-1;i++{

		var toBeHashed []byte
		var Hashed []byte
		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k][j])
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			toBeHashed=append(toBeHashed,merkletree[k+1][j])
		}

		Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkletree[i+totNumOfEncOutVecToMer][j]=Hashed[j]
		}
		k=k+2
	}

	return merkletree  //returns the merkleroot
}
