package functions

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"path/filepath"
)

// function to check error
func CheckError(e error) {
    if e != nil {
        panic(e)
    }
}

// function to create directory if not exist
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		CheckError(err)
	}
}

// function to create tmp folder and subsequent inner directories
func Createtmp() {
	CreateDirIfNotExist("./tmp")
	CreateDirIfNotExist("./tmp/contractInfo")
	CreateDirIfNotExist("./tmp/private")
	CreateDirIfNotExist("./tmp/private/fog")
	CreateDirIfNotExist("./tmp/private/cloud")
	CreateDirIfNotExist("./tmp/private/client")
	CreateDirIfNotExist("./tmp/private/cloud/EncryptedFiles")
	CreateDirIfNotExist("./tmp/private/client/DecryptedFiles")
	CreateDirIfNotExist("./tmp/private/client/SharedData")
	}

// function to create file 's' and write 'b' bytes in it
func FileCrWr(s string, b []byte) {
	f, err := os.OpenFile(s, os.O_RDWR|os.O_CREATE, 0755)
	CheckError(err)
	_, err = f.Write(b)
	CheckError(err)
	defer f.Close()
	}

// function to get path of Geth node and keystore key
func GethPathAndKey() (string, string) {

	dir, err := os.Getwd()
	CheckError(err)
	fmt.Printf("Current dir: %q\n", dir)

	//pathParent := path.Dir(dir)
	//pathParent += "/Private_Ethereum/Node2/"
	pathParent := "/home/jayasree/Private_Ethereum/Node2/"
	gethPath := pathParent + "geth.ipc"
	
	var fileName string
	var filePath string
	keyStorePath := pathParent + "keystore/"
	
	err = filepath.Walk(keyStorePath, func(path string, info os.FileInfo, err error) error {
		CheckError(err)
		if !info.IsDir() {
			fmt.Printf("visited file or dir: %q\n", path)
			fileName = info.Name()
			filePath = path
		}
		return nil
	})
	CheckError(err)
	
	var key []byte
	keyRead := func() {
		var err error
		key, err = ioutil.ReadFile(filePath)
		CheckError(err)
    }
    keyRead()
    return gethPath, string(key)
}

func AESencrypt (key []byte, inputFilePath string) ([]byte, []byte) {

	//read plaintext data from a file 
	plaintext, err := ioutil.ReadFile(inputFilePath)
	CheckError(err)

	block, err := aes.NewCipher(key)
	CheckError(err)

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	CheckError(err)

	//Compute ciphertext and write to a File
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	FileCrWr("../tmp/private/cloud/EncryptedFiles/Encrypted.txt", ciphertext)

	return nonce, ciphertext
}

func AESdecrypt (test_key []byte, nonce []byte) {

	block_d, err := aes.NewCipher(test_key)
	CheckError(err)

	aesgcm_d, err := cipher.NewGCM(block_d)
	CheckError(err)

	//read the encrypted data/ciphertext from a file
	ct_val, err := ioutil.ReadFile("../tmp/private/cloud/EncryptedFiles/Encrypted.txt")
	CheckError(err)

	//decrypt the ciphertext and write back the plaintext to another file
	plaintext_decrypted, err := aesgcm_d.Open(nil, nonce, ct_val, nil)
	CheckError(err)
	FileCrWr("../tmp/private/client/DecryptedFiles/Decrypted.txt", plaintext_decrypted)

}


