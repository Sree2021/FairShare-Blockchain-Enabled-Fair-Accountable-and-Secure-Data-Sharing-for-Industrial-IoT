package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"net"
	"bufio"
	"time"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	store "../contracts" // for demo
	functions "../func"
	fp "../fairswap"
)

func ContractVars() (*store.Store, *bind.TransactOpts, *ethclient.Client) {
	gethPath, nodeKey := functions.GethPathAndKey()
    
	addrByte, err := ioutil.ReadFile("../tmp/contractInfo/contract-address")
	functions.CheckError(err)
    
	connection, err := ethclient.Dial(gethPath)
	functions.CheckError(err)
    
	instContract, err := store.NewStore(common.HexToAddress(string(addrByte)), connection)
	functions.CheckError(err)
    
	auth, err := bind.NewTransactor(strings.NewReader(nodeKey), "account1")
	functions.CheckError(err)

	return instContract, auth, connection
}

func handleConnection(conn net.Conn, result chan string) {
	fmt.Println("Inside function")
	// will listen for message to process ending in newline (\n)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	//conn.Close()
	result<-message
	return
}

func tcpServer(types string, port string, instContract *store.Store) (int, time.Duration) {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen(types, port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Accept connection on port")

	result := make(chan string)	

	// run loop forever (or until ctrl-c)
	for {
		// accept connection on port
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Calling handleConnection")
		go handleConnection(conn,result)
		message := <-result
		fmt.Print("Printing from Cloud ---> Message Received:", string(message))
		msg := "Client"

		start1 := time.Now()
		f_val, err := ioutil.ReadFile("../tmp/private/cloud/EncryptedFiles/Encrypted.txt")
		functions.CheckError(err)
		hash1 := crypto.Keccak256Hash(f_val)
		result, err := instContract.CompareMeta1(nil, hash1.Hex())
		fmt.Println("Result=",result)
		functions.CheckError(err)
		elapsed1 := time.Since(start1)
		log.Printf("\nFile Integrity Verification Took : %s", elapsed1)
		if (result == true) {

			start := time.Now()
			result, err := instContract.CompareAccessPolicy(nil, string(msg))
			functions.CheckError(err)
			fmt.Println("CompareAccessPolicy() Returned:", result)

			// sample process for string received
			newmessage1 := "Acknowledged"
			newmessage2 := "Aborted"
		
			// send decision back to client
			if (result == true && strings.TrimRight(message, "\n") == "D101") {
				conn.Write([]byte(newmessage1 + "\n"))
				elapsed := time.Since(start)
    				functions.FileCrWr("../tmp/private/fog/Result", []byte(newmessage1))
				f:= 1
				return f,elapsed
			} else {
				conn.Write([]byte(newmessage2 + "\n"))
				elapsed := time.Since(start)
				functions.FileCrWr("../tmp/private/fog/Result", []byte(newmessage2))
				f:= 0
				return f,elapsed
			}
		} else {
			newmessage2 := "Aborted"
			conn.Write([]byte(newmessage2 + "\n"))
			elapsed:= time.Since(start1)
			functions.FileCrWr("../tmp/private/fog/Result", []byte(newmessage2))
			f:= 2
			return f,elapsed
		}
	}
}

func main() {

	instContract, auth, connection := ContractVars()
	fmt.Println("connection Returned: ", connection)
	fmt.Println("Auth Returned: ", auth)

	senderAddress := common.HexToAddress(string("0xD900C31957e60bE22f0Ec99b82A7F97fD694d448"))
	receiverAddress := common.HexToAddress(string("0x30212Ff2b011cE7FCF74229Da36a41B3d4B0FF3B"))
	fp.SetSenderReceiver(auth, senderAddress, receiverAddress, connection)

	f, elapsed:= tcpServer("tcp", "127.0.0.1:8081", instContract)
	if (f == 1) {
		log.Printf("\nAccess Policy Verified Successfully with Time : %s", elapsed)
		start2 := time.Now()
		fp.ExecuteFairSwap()
		elapsed2 := time.Since(start2)
		log.Printf("\nFair Swap Took : %s", elapsed2)
	} else if (f == 0) {
		log.Printf("\nAccess Policy Verification Failed with Time %s", elapsed)
	} else {
		log.Printf("\nProcess Aborted")	
	}
}
