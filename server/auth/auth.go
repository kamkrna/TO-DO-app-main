package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("Qk64CRujxuKvySD596883YI4akmnDt6B0t8MiZYPbf8cHhHPGyb4UlTLSFjS1DQcKXDO4VixB1x714uc7uuY+g==")

type Claims struct{
	Username string `json:"username"`
	jwt.StandardClaims
}

func HashedPassword(password string) (string,error){
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedpassword),err
}

func CheckPassword(hash, password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
	// if err != nil {
    //     log.Printf("Password mismatch: %v", err)
    //     return false
    // }
    // return true
}

func CreateToken(username string) (string, error){
	expirationTime := time.Now().Add(time.Hour * 24 * 14) // 14 days
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil{
		log.Println("Couldn't create token!")
		return "",err
	}
	log.Printf("Token created at: %v, expires at: %v", time.Now(), expirationTime)
	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error){
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
	if err != nil{
		return nil,err
	}
	if !token.Valid{
		return nil, fmt.Errorf("invalid token")
	}
	log.Printf("Validating token: %s", tokenString)
	log.Printf("Token valid until: %v", time.Unix(claims.ExpiresAt, 0))
	return claims, nil
}

