package main

import (
	"fmt"
	"os"
	"encoding/pem"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
)

func main(){
	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Error generating key: %s\n", err)
		os.Exit(1)
	}

	publickey := &privatekey.PublicKey

	// dump private key to file
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	privatePem, err := os.Create("private.pem")
	if err != nil {
		fmt.Printf("Error creating private.pem: %s\n", err)
		os.Exit(1)
	}

	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		fmt.Printf("Error encoding private.pem: %s\n", err)
		os.Exit(1)
	}

	// dump public key to file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		fmt.Printf("Error marshalling public key: %s\n", err)
		os.Exit(1)
	}

	publicKeyBlock := &pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicPem, err := os.Create("public.pem")
	if err != nil {
		fmt.Printf("Error creating public.pem: %s\n", err)
		os.Exit(1)
	}

	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		fmt.Printf("Error encoding public.pem: %s\n", err)
		os.Exit(1)
	}

}
