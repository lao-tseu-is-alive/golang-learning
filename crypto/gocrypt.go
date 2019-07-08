package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"runtime"
)

func doCrypt(theSecret string) (string, []uint8) {
	defer golog.Un(golog.Trace("doCrypt()"))
	block, str, iv, encrypted := Encrypt("key4567890keykey", "cipher7890cipher", theSecret)
	secretDecrypted := Decrypt(block, str, iv, encrypted)
	return secretDecrypted, encrypted
}

func Encrypt(key string, cipherTextRaw string, textToEncrypt string) (cipher.Block, []byte, []byte, []uint8) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	str := []byte(textToEncrypt)

	cipherText := []byte(cipherTextRaw)
	iv := cipherText[:aes.BlockSize]
	encrypter := cipher.NewCFBEncrypter(block, iv)

	encrypted := make([]byte, len(str))
	encrypter.XORKeyStream(encrypted, str)
	golog.Info("\n'%s'\n## was encrypted to : ##\n%v\n", str, encrypted)
	return block, str, iv, encrypted
}

func Decrypt(block cipher.Block, str []byte, iv []byte, encrypted []uint8) string {
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(str))
	decrypter.XORKeyStream(decrypted, encrypted)
	golog.Info("\n%v\n## was decrypted to : ##\n'%s'\n", encrypted, decrypted)
	return string(decrypted)
}

func main() {

	info := fmt.Sprintf(`This app binary was build by Go version : %s`, runtime.Version())

	original, enc := doCrypt(info)
	fmt.Println(original, enc)

	/*
		golog.Err("something went terribly wrong here !")
		golog.Warn("NumCPU : %d, GOMAXPROCS : %d  ", runtime.NumCPU(), runtime.GOMAXPROCS(-1))
		golog.Info("just a simple information message to send to log")

	*/
}
