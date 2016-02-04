package utils

import (
	"crypto/rsa"
	"crypto/sha512"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var PrivateKey	*rsa.PrivateKey
var PublicKey	*rsa.PublicKey

const pathPrivateKey = "/etc/configurations/keys/private.pem"
const pathPublicKey = "/etc/configurations/keys/public.pem"

var duration, _ = time.ParseDuration("2h")

func keyFunc(token *jwt.Token) (interface{}, error) {
	return PublicKey, nil
}

func GenerateToken(userID uint, scope byte) (string, error) {
	now 		:= time.Now().UTC()
	expiration 	:= now.Add(duration)
	token 		:= jwt.New(jwt.SigningMethodRS512)

	token.Claims["iss"] = "coban"
	token.Claims["exp"] = float64(expiration.Unix())
	token.Claims["iat"] = float64(now.Unix())
	token.Claims["nbf"] = float64(now.Unix())

	token.Claims["user"] = userID
	token.Claims["scope"] = scope

	return token.SignedString(PrivateKey)
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, keyFunc)
}

func ParseTokenFromRequest(r *http.Request) (*jwt.Token, error) {
	return jwt.ParseFromRequest(r, keyFunc)
}

func HashPassword(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func init() {
	var data 	[]byte
	var err 	error

	if data, err = ioutil.ReadFile(pathPrivateKey); err != nil {
		log.Fatal(err)
	}
	if PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(data); err != nil {
		log.Fatal(err)
	}
	if data, err = ioutil.ReadFile(pathPublicKey); err != nil {
		log.Fatal(err)
	}
	if PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(data); err != nil {
		log.Fatal(err)
	}
}
