package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func hashAndSalt(plainPwd string) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	bytePassword := []byte(plainPwd)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if err != nil {
		golog.Err("Error doing bcrypt.GenerateFromPassword : %v", err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	// we need to convert strings to byte slice
	byteHash := []byte(hashedPwd)
	bytePassword := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		golog.Err("Error doing bcrypt.CompareHashAndPassword %v ", err)
		return false
	}

	return true
}

func sha256Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {

	if len(os.Args) > 1 {
		password := os.Args[1]
		fmt.Printf("\nYour original password is   : %s", password)
		passwordHash := sha256Hash(password)
		fmt.Printf("\nThe html client side will calculate the SHA256 hash of your password and send it to server")
		fmt.Printf("\nThe password sha256 hash is : %s", passwordHash)
		fmt.Printf("\nThe server receives the SHA256 hash of your password and calculate a salted hash")
		fmt.Printf("\nThe salted hash is always different for the same password, to disallow password guessing")
		goHash := hashAndSalt(passwordHash)
		fmt.Printf("\nYour BCRYPT password hash (salted) to store in DB is :\n%s", goHash)
		ok := comparePasswords(goHash, passwordHash)
		fmt.Printf("\nThe verify_password with this hash any our password IS : %v\n", ok)
	} else {
		fmt.Println("This utility program will give you a salted bcrypt hash of your password ")
		fmt.Println("Expecting 2 parameters : Password and Your_Salted_Hash_FromDB_2Test")
	}
}
