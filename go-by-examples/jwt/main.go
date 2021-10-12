package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	jwt "github.com/dgrijalva/jwt-go"
)

// Token struct for JWT tokens
type Token struct {
	jwt.StandardClaims
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("request:", request.Headers)

	var res = "Got it!"

	// Get token from the header
	authorization := request.Headers["Authorization"]
	fmt.Println("auth:", authorization)

	// If authorization is missing
	if authorization == "" {
		res = "Missing auth token"
	}

	// Authorization comes in two parts
	// Bearer <token>
	splitted := strings.Split(authorization, " ")
	if len(splitted) != 2 {
		res = "Invalid/Malformed auth token"
	}

	// Capture actual token part
	tokenPart := splitted[1]
	fmt.Println("tokenPart:", tokenPart)

	validate := ValidateJWT(tokenPart)
	fmt.Println(validate)

	// token :=CreateJWT();
	// fmt.Println("Token:", token)

	response := events.APIGatewayProxyResponse{
		Body: res,
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
