package constant

import (
	"crypto/rsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

//JWTPublicKey key
var JWTPublicKey *rsa.PublicKey

//JWTPrivateKey key
var JWTPrivateKey *rsa.PrivateKey

func init() {
	JWTPrivateKey, _ = LoadRSAPrivateKeyFromDisk("keys/jwt")
	JWTPublicKey, _ = LoadRSAPublicKeyFromDisk("keys/jwt.pub")
}

//LoadRSAPrivateKeyFromDisk load private key
func LoadRSAPrivateKeyFromDisk(location string) (*rsa.PrivateKey, error) {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	return jwt.ParseRSAPrivateKeyFromPEM(keyData)
}

//LoadRSAPublicKeyFromDisk load public key
func LoadRSAPublicKeyFromDisk(location string) (*rsa.PublicKey, error) {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	return jwt.ParseRSAPublicKeyFromPEM(keyData)
}
