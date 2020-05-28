package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	defer duration(track("main"))

	clientPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	clientPublicKey := &clientPrivateKey.PublicKey

	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	serverPublicKey := &serverPrivateKey.PublicKey

	client_privkey_bytes := x509.MarshalPKCS1PrivateKey(clientPrivateKey)
	client_privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: client_privkey_bytes,
		},
	)

	client_pubkey_bytes := x509.MarshalPKCS1PublicKey(clientPublicKey)
	client_pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: client_pubkey_bytes,
		},
	)

	server_privkey_bytes := x509.MarshalPKCS1PrivateKey(serverPrivateKey)
	server_privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: server_privkey_bytes,
		},
	)

	server_pubkey_bytes := x509.MarshalPKCS1PublicKey(serverPublicKey)
	server_pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: server_pubkey_bytes,
		},
	)

	fmt.Println("Client Keys:")
	fmt.Println(string(client_privkey_pem))
	fmt.Println(string(client_pubkey_pem))
	fmt.Println("Server Keys:")
	fmt.Println(string(server_privkey_pem))
	fmt.Println(string(server_pubkey_pem))

}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
