package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"context"
	"time"
	"github.com/Nik-U/pbc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/scrypt"

	store "../contracts" // for demo
	functions "../func"
)

var pairing *pbc.Pairing
var pk_F *pbc.Element
var sk_F *pbc.Element
var c1 *pbc.Element
var c2 *pbc.Element
var rk_FC *pbc.Element
var c11 *pbc.Element
var c12 *pbc.Element
var nonce []byte
var ciphertext []byte
var c []byte
var rct []byte

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

func timeTrack(start time.Time, name string) {
        elapsed := time.Since(start)
        fmt.Printf("\nFunction %s took: %s \n", name,elapsed)
}


func main() {

	DId := "D101"

	//generate key for AES
	salt := []byte{0xd8, 0x26, 0xc2, 0x57, 0xa7, 0x6b, 0xad, 0x7d}
	key, err := scrypt.Key([]byte("Example Key"), salt, 1<<15, 8, 1, 32) //returns a []byte as key
	functions.CheckError(err)

	instContract, auth, connection := ContractVars()
	fmt.Println("connection Returned: ", connection)
	fmt.Println("Auth Returned: ", auth)
	ctx := context.Background()

	inputFilePath := "../InputFiles/File8.txt"
	q := func() {
        defer timeTrack(time.Now(), "q")
	nonce, ciphertext = functions.AESencrypt(key, inputFilePath)
	}
	q()
	
	r := func() {
        defer timeTrack(time.Now(), "r")
	
	tx, err := instContract.SetNonce(auth, nonce)
	functions.CheckError(err)
	fmt.Println("tx=",tx.Hash().Hex())
	
	_, err = bind.WaitMined(ctx, connection, tx)
	functions.CheckError(err)
	}
	r()

	GetParamsByte_val, GetG_val, err := instContract.GetParams(nil)
	functions.CheckError(err)
	
	pairing, err = pbc.NewPairingFromString(string(GetParamsByte_val))
	functions.CheckError(err)

	g := pairing.NewG1().SetBytes(GetG_val)
	m := pairing.NewGT().SetBytes(key)
	
	s := func() {
        defer timeTrack(time.Now(), "s")

	sk_F = pairing.NewZr().Rand()
	pk_F = pairing.NewG1().PowZn(g, sk_F)
	}
	s()

	functions.FileCrWr("../tmp/private/fog/skF", sk_F.Bytes())

	t := func() {
        defer timeTrack(time.Now(), "t")
	
	tx1, err := instContract.SetPKF(auth, pk_F.Bytes())
	functions.CheckError(err)
	fmt.Println("tx1=",tx1.Hash().Hex())
	
	_, err = bind.WaitMined(ctx, connection, tx1)
	functions.CheckError(err)
	}
	t()

	GetPKC_val, err := instContract.GetPKC(nil)
	functions.CheckError(err)
	pk_C := pairing.NewG1().SetBytes(GetPKC_val)
	
	u := func() {
        defer timeTrack(time.Now(), "u")
	r := pairing.NewZr().Rand()
	c1 = pairing.NewG1().PowZn(pk_F, r)
	c2 = pairing.NewGT().Mul((pairing.NewGT().PowZn((pairing.NewGT().Pair(g, g)), r)),m)
	}
	u()

	c = append(c, c1.Bytes()...)
	c = append(c, c2.Bytes()...)
	
	accessP := "Client"

	v := func() {
        defer timeTrack(time.Now(), "v")
	tx2, err := instContract.SetAccessP(auth, accessP)
	functions.CheckError(err)
	fmt.Println("tx2=",tx2.Hash().Hex())
	
	_, err = bind.WaitMined(ctx, connection, tx2)
	functions.CheckError(err)
	}
	v()

	w := func() {
        defer timeTrack(time.Now(), "w")
	hash1 := crypto.Keccak256Hash(ciphertext)
	hash2 := crypto.Keccak256Hash(c)

	tx3, err := instContract.SetMeta1(auth, DId, hash1.Hex(), hash2.Bytes())
	functions.CheckError(err)
	fmt.Println("tx3=",tx3.Hash().Hex())
	
	_, err = bind.WaitMined(ctx, connection, tx3)
	functions.CheckError(err)
	}
	w()

	message, err := ioutil.ReadFile("../tmp/private/fog/Result")
	functions.CheckError(err)

	if (string(message) == "Acknowledged") {
	
		x := func() {
        	defer timeTrack(time.Now(), "x")
		rk_FC = pairing.NewG1().PowZn(pk_C, pairing.NewZr().Invert(sk_F))
		}
		x()
			
		y := func() {
        	defer timeTrack(time.Now(), "y")
		t := pairing.NewZr().Rand()
		c11 = pairing.NewG1().PowZn(rk_FC, pairing.NewZr().Invert(t))
		c12 = pairing.NewG1().PowZn(c1, t)
		}
		y()

		functions.FileCrWr("../tmp/private/client/SharedData/c11", c11.Bytes())
		functions.FileCrWr("../tmp/private/client/SharedData/c12", c12.Bytes())
		functions.FileCrWr("../tmp/private/client/SharedData/c2", c2.Bytes())

		rct = append(rct, c11.Bytes()...)
		rct = append(rct, c12.Bytes()...)
		rct = append(rct, c2.Bytes()...)
				
		z := func() {
		defer timeTrack(time.Now(), "z")
		hash3 := crypto.Keccak256Hash(rct)	
		tx4, err := instContract.SetMeta2(auth, DId, hash3.Hex())
		functions.CheckError(err)
		fmt.Println("tx4=",tx4.Hash().Hex())
	
		_, err = bind.WaitMined(ctx, connection, tx4)
		functions.CheckError(err)
		}
		z()

	} else {
		fmt.Println("Abort Message sent from Cloud")
	}
}
