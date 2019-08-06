package account

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"time"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
)

var templateCertificate = x509.Certificate{
	SerialNumber: big.NewInt(1),
	Subject: pkix.Name{
		Organization: []string{"Soket Labs"},
	},
	NotBefore: time.Now(),
	NotAfter:  time.Now().Add(time.Hour * 168),

	KeyUsage: x509.KeyUsageDigitalSignature,
}

// GenerateAccount generates an account and returns .PEM
// 1. Generate a key pair
// 2. Create a self signed certificate (for now)
// 3. Return the prv key and signed certificate as pem.
func GenerateAccountPEM() ([]byte, error) {

	_, prvKey, certDER, err := GenerateAccount()
	if err != nil {
		return nil, err
	}

	outPEM := &bytes.Buffer{}

	pem.Encode(outPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})

	marshalPrvKey, err := crypto.MarshalPrivateKey(prvKey)
	if err != nil {
		return nil, err
	}
	//block, err := x509.EncryptPEMBlock(rand.Reader, "PRIVATE_KEY")
	pem.Encode(outPEM, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: marshalPrvKey,
	})

	return outPEM.Bytes(), nil
}

func GenerateAccount() (*x509.Certificate, crypto.PrivKey, []byte, error) {
	prvKey, pubKey, err := crypto.GenerateRSAKeyPair(2048, rand.Reader)
	if err != nil {
		return nil, nil, nil, err
	}

	marshalPvtKey, err := prvKey.Raw()
	if err != nil {
		return nil, nil, nil, err
	}
	nativePvtKey, err := x509.ParsePKCS1PrivateKey(marshalPvtKey)
	if err != nil {
		return nil, nil, nil, err
	}

	marshalPubKey, err := pubKey.Raw()
	if err != nil {
		return nil, nil, nil, err
	}
	nativePubKeyInt, err := x509.ParsePKIXPublicKey(marshalPubKey)
	if err != nil {
		fmt.Println("gggg", err)
		return nil, nil, nil, err
	}
	nativePubKey := nativePubKeyInt.(*rsa.PublicKey)

	certDER, err := x509.CreateCertificate(rand.Reader, &templateCertificate, &templateCertificate, nativePubKey, nativePvtKey)
	if err != nil {
		return nil, nil, nil, err
	}

	cer, err := x509.ParseCertificate(certDER)
	if err != nil {
		return nil, nil, nil, err
	}

	return cer, prvKey, certDER, nil
}

func ImportAccount(data []byte) (*x509.Certificate, crypto.PrivKey, error) {

	block, rem := pem.Decode(data)
	if block == nil {
		return nil, nil, errors.New("Error decoding pem")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, err
	}

	block, rem = pem.Decode(rem)
	prvKey, err := crypto.UnmarshalRsaPrivateKey(block.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return cert, prvKey, nil
}
