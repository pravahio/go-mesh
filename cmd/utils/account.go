package utils

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"

	cli "gopkg.in/urfave/cli.v1"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	libcrypto "github.com/libp2p/go-libp2p-core/crypto"
	peer "github.com/libp2p/go-libp2p-core/peer"
	eth "github.com/upperwal/go-mesh/driver/eth"
	utils "github.com/upperwal/go-mesh/util/account"
)

var (
	CREATE_ACCOUNT          = "c"
	ACCOUNT_OUTPUT_FILENAME = "o"
	PRINT_ACCOUNT_ID        = "p"
	PARSE                   = "parse"
)

var (
	CreateAccountFlag = cli.BoolFlag{
		Name:  CREATE_ACCOUNT,
		Usage: "creates a new account",
	}

	AccountOutputFileNameFlag = cli.StringFlag{
		Name:  ACCOUNT_OUTPUT_FILENAME,
		Usage: "account output file",
	}

	PrintPublicKeyFlag = cli.BoolFlag{
		Name:  PRINT_ACCOUNT_ID,
		Usage: "print account id of the newly created account",
	}

	ParseAccountFile = cli.StringFlag{
		Name:  PARSE,
		Usage: "parse an account file",
	}
)

func AccountCommandHandler(ctx *cli.Context) {
	if ctx.Bool(CREATE_ACCOUNT) {
		createAccount(ctx)
		return
	}
	if f := ctx.String(PARSE); f != "" {
		parse(f)
		return
	}

	cli.ShowCommandHelpAndExit(ctx, "account", 0)
}

func createAccount(ctx *cli.Context) error {
	keyMap := make(map[string][]byte)

	ethPrivKey, err := eth.CreateNewAccount()
	if err != nil {
		return err
	}

	_, privKey, _, err := utils.GenerateAccount()
	if err != nil {
		return err
	}

	peerId, err := peer.IDFromPrivateKey(privKey)
	if err != nil {
		return err
	}

	ethAdd := ethcrypto.PubkeyToAddress(ethPrivKey.PublicKey)

	encodeEthPrivKey := ethcrypto.FromECDSA(ethPrivKey)
	if encodeEthPrivKey == nil {
		fmt.Println("ETH key is nil")
		return errors.New("ETH key is nil")
	}
	encodedLibp2pPrivKey, err := libcrypto.MarshalPrivateKey(privKey)
	if err != nil {
		fmt.Println(err)
		return err
	}

	keyMap["RAPrivKey"] = encodeEthPrivKey
	keyMap["Libp2pPrivKey"] = encodedLibp2pPrivKey

	fn := "mesh.msa"
	if n := ctx.String(ACCOUNT_OUTPUT_FILENAME); n != "" {
		fn = n
	}

	err = writeToFile(fn, keyMap)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if ctx.Bool(PRINT_ACCOUNT_ID) {
		fmt.Println("Peer ID: ", peerId)
		fmt.Println("Auth Add: ", ethAdd.String())
	}

	return nil
}

func writeToFile(fileName string, v map[string][]byte) error {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)

	err := e.Encode(v)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, b.Bytes(), 0644)
}

func readFromFile(fileName string) (map[string][]byte, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(b)
	d := gob.NewDecoder(buf)

	var m map[string][]byte

	err = d.Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func parse(fileName string) {

	libPriv, ethPriv, err := GetLibp2pAndRAPrivKey(fileName)

	peerID, err := peer.IDFromPrivateKey(libPriv)
	if err != nil {
		fmt.Println(err)
	}

	ethAdd := ethcrypto.PubkeyToAddress(ethPriv.PublicKey)

	fmt.Println("Peer ID: ", peerID)
	fmt.Println("Auth Add: ", ethAdd.String())
}

func GetLibp2pAndRAPrivKey(filename string) (libcrypto.PrivKey, *ecdsa.PrivateKey, error) {
	m, err := readFromFile(filename)
	if err != nil {
		return nil, nil, err
	}

	if len(m) != 2 {
		return nil, nil, errors.New("need 2 keys in .msa file")
	}

	ethPriv, err := ethcrypto.ToECDSA(m["RAPrivKey"])
	if err != nil {
		return nil, nil, err
	}

	libPriv, err := libcrypto.UnmarshalPrivateKey(m["Libp2pPrivKey"])
	if err != nil {
		return nil, nil, err
	}

	return libPriv, ethPriv, nil
}
