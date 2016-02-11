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

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

const pathPrivateKey = "/etc/configurations/keys/private.pem"
const pathPublicKey = "/etc/configurations/keys/public.pem"

var duration, _ = time.ParseDuration("2h")

func keyFunc(token *jwt.Token) (interface{}, error) {
	return publicKey, nil
}

// GenerateToken generate a new jwt token signed with the private key
// and holding all the information related to the user, it's right and
// it's expirency
func GenerateToken(userID uint, scope byte) (string, error) {
	now := time.Now().UTC()
	expiration := now.Add(duration)
	token := jwt.New(jwt.SigningMethodRS512)

	token.Claims["iss"] = "coban"
	token.Claims["exp"] = float64(expiration.Unix())
	token.Claims["iat"] = float64(now.Unix())
	token.Claims["nbf"] = float64(now.Unix())

	token.Claims["user"] = userID
	token.Claims["scope"] = scope

	return token.SignedString(privateKey)
}

// ParseToken parse the received token and check if it is a valid one with
// the server's public key
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, keyFunc)
}

// ParseTokenFromRequest is identical to ParseToken excepti it check the token
// in the HTTP header "Authorization: Bearer <token>"
func ParseTokenFromRequest(r *http.Request) (*jwt.Token, error) {
	return jwt.ParseFromRequest(r, keyFunc)
}

// HashPassword hash a string into the used hash for the database.
func HashPassword(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func init() {
	var data []byte
	var err error

	if data, err = ioutil.ReadFile(pathPrivateKey); err != nil {
		log.Fatal(err)
	}
	if privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(data); err != nil {
		log.Fatal(err)
	}
	if data, err = ioutil.ReadFile(pathPublicKey); err != nil {
		log.Fatal(err)
	}
	if publicKey, err = jwt.ParseRSAPublicKeyFromPEM(data); err != nil {
		log.Fatal(err)
	}
}
