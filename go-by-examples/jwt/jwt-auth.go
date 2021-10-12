package main

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWT key that will be used to create the signature
// TODO(Abu): Replace this with actual private key/passowrd.
var jwtKey = []byte("super_secrect_key")

// Claims struct to encoded to a JWT.
// jwt.StandardClaims used to provide
// for example expiry time
type Claims struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	SSN       string `json:"ssn,omitempty"`
	jwt.StandardClaims
}

// JwtResponse response structure
type JwtResponse struct {
	Name   string    `json:"token"`
	Value  string    `json:"value"`
	Expire time.Time `json:"expireTime"`
}

// CreateJWT handler signes JWT
func CreateJWT() *JwtResponse {

	// Sets token expiration time
	expirationTime := time.Now().Add(30 * time.Minute)

	// Create the JWT claims, which includes the following
	// user info and expiry time.
	claims := &Claims{
		FirstName: "Abu",
		LastName:  "Shumon",
		SSN:       "121212-111A",
		StandardClaims: jwt.StandardClaims{
			// Expiry time is expressed as unix milliseconds in JWT
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	var jwtRes JwtResponse
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		jwtRes.Name = ""
		jwtRes.Value = ""
		jwtRes.Expire = time.Now()

		return &jwtRes
	}

	jwtRes.Name = "token"
	jwtRes.Value = tokenString
	jwtRes.Expire = expirationTime

	return &jwtRes

}

// ValidateJWT handler validates a JWT claim
func ValidateJWT(tkn JwtResponse) *JwtResponse {
	var jwtRes JwtResponse

	if tkn.Value == "" {
		jwtRes.Value = "Invalid token"
	}

	tokenPart := tkn.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenPart, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	// Malformed token
	if err != nil {
		jwtRes.Value = "Malformed authentication token"
	}

	// Token is invalid, maybe not signed on this server
	if !token.Valid {
		jwtRes.Value = "Invalid token!"
	}

	jwtRes.Name = tkn.Name
	jwtRes.Value = tkn.Value
	jwtRes.Expire = tkn.Expire

	return &jwtRes
}
