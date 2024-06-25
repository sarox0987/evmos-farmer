package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

func createWalletFromPrivateKey(privateKeyHex string) (*wallet, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error converting public key to ECDSA type")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &wallet{
		privateKey: privateKey,
		address:    address,
	}, nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	// Function to read a string input
	readString := func(prompt string) string {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}

	privateKeyHex := readString("Enter EVM Private Key: ")
	url := readString("Enter RPC URL: ")

	wallet, err := createWalletFromPrivateKey(privateKeyHex)
	if err != nil {
		fmt.Println("Error creating wallet:", err)
		return
	}
	fmt.Println(wallet.address)
	chainId := big.NewInt(1234)
	c, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	balance, err := c.BalanceAt(context.Background(), wallet.address, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("balance: ", balance)

	i := 0
	for {
		i++
		sendTx(0, c, wallet.address, wallet.address, big.NewInt(int64(i%10+1)), chainId, wallet.privateKey)
		time.Sleep(10 * time.Second)
	}
}

func sendTx(i int, c *ethclient.Client, from, to common.Address, amount *big.Int, chainId *big.Int, pk *ecdsa.PrivateKey) {
	nonce, err := c.NonceAt(context.Background(), from, nil)
	if err != nil {
		log.Printf("Failed to retrieve nonce: %v", err)
		return
	}

	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Failed to suggest gas price: %v", err)
		return
	}

	gasLimit := uint64(22000)

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), pk)
	if err != nil {
		log.Printf("Failed to sign transaction: %v", err)
		return
	}

	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("Failed to send transaction: %v", err)
		return
	}

	fmt.Printf("Address: %d Tx %d: %s\n", i, nonce, tx.Hash())
}
