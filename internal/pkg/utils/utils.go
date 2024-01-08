package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// keycloak's pubkey
const pubKey = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnIwQK/tDhd7HTres6Kp4Pkfs93xMprzHZ7lYrDPa4eAgJnsGGVe4kAb//fEIQpHcqrfpBjg/8K0UXUfULEibhkAYLSyaKGs6EtvuZ7kA+zqIol9WMtk/BUaEG6KFHlIYsovIJCNygWLUqKd3t6Xmfd0q69uYJ+OuZe3P138dE0E7b7Y88M4RechBWb3lFMyhbT4kz4+jPGolEC7f3s6CgHoZbXNWhpi/zw8eDzGiNx7VJl2yoZXFiWQMnWNPFFBViQaWtNa4bQat3r8l7lys1x7TTYG+g7jLSgtQ3DQio+PidEP4z2H22GDRyC2s5jQTzTrbsgReGZeDXBwqjUvhMQIDAQAB\n-----END PUBLIC KEY-----"

func VerifyJWTToken(target string) (*jwt.Token, error) {

	pubKeyBytes := []byte(pubKey)
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	// Parse the token
	token, err := jwt.Parse(target, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse and verify token: %v", err)
	}

	// Verify the token claims
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, nil
	}

	return nil, fmt.Errorf("token validation failed")
}
