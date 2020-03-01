package brizochain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	BrizoContract "github.com/ntu-brizo/brizochain/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// TODO: make config more flexible
const rpcUrl = "http://140.112.18.197:8051"
const contractAddr = "0xFce98B78cfF8884f3dd61bE4653A3f409a44931f"
const privKey = "63c9dbad87b6b367c0a16eb1d928b68d5a1ae9d993ab81b2b7d665d6fad127b9"

type brizoChain struct {
	client           *ethclient.Client
	txSigner         *bind.TransactOpts
	contractInstance *BrizoContract.Brizo
}

// NewBrizoChain returns a instance of brizoChain (factory)
func NewBrizoChain() (*brizoChain, error) {
	rpcDial, err := rpc.Dial(rpcUrl)
	if err != nil {
		return nil, errors.New("Failed to create a RPC client.")
	}

	client := ethclient.NewClient(rpcDial)

	contract, err := BrizoContract.NewBrizo(common.HexToAddress(contractAddr), client)
	if err != nil {
		return nil, errors.New("Failed to instantiate contract.")
	}

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, errors.New("Invalid private key format.")
	}
	signer := bind.NewKeyedTransactor(privateKey)
	signer.GasLimit = 200000

	return &brizoChain{client: client, txSigner: signer, contractInstance: contract}, nil
}

// Write func writes content to Brizo Chain and returns error if error occurs.
func (brizo *brizoChain) Write(msg string) error {
	tx, err := brizo.contractInstance.WriteData(brizo.txSigner, msg)
	if err != nil {
		return errors.New("Failed to broadcast transaction to Brizo Chain.")
	}

	_, err = bind.WaitMined(context.Background(), brizo.client, tx)
	if err != nil {
		return errors.New("Transaction mining error.")
	}

	return nil
}

// WriteByHashKey func write content to Brizo Chain with a key and returns error if error occurs.
func (brizo *brizoChain) WriteByHashKey(key string, content string) error {
	//contentSliceArray := []string{}
	const SLICE_EVERY_N = 50
	for i := 0; i <= len(content)/SLICE_EVERY_N; i++ {
		if i == len(content)/SLICE_EVERY_N {
			tx, err := brizo.contractInstance.WriteDataToHashDict(brizo.txSigner, key, content[SLICE_EVERY_N*i:len(content)])
			fmt.Println(content[SLICE_EVERY_N*i : len(content)])
			if err != nil {
				return errors.New("Failed to broadcast transaction to Brizo Chain.")
			}

			_, err = bind.WaitMined(context.Background(), brizo.client, tx)
			if err != nil {
				return errors.New("Transaction mining error.")
			}
		} else {
			tx, err := brizo.contractInstance.WriteDataToHashDict(brizo.txSigner, key, content[SLICE_EVERY_N*i:SLICE_EVERY_N*(i+1)])
			fmt.Println(content[SLICE_EVERY_N*i : SLICE_EVERY_N*(i+1)])
			if err != nil {
				return errors.New("Failed to broadcast transaction to Brizo Chain.")
			}

			_, err = bind.WaitMined(context.Background(), brizo.client, tx)
			if err != nil {
				return errors.New("Transaction mining error.")
			}
		}
	}

	return nil
}

// Read returns data on Brizo Chain according to the given index (start at 0)
func (brizo *brizoChain) Read(index int64) (string, error) {
	return brizo.ReadByAddress(brizo.txSigner.From, index)
}

// ReadByAddress returns data on Brizo Chain by passing address (common.Address) and index
func (brizo *brizoChain) ReadByAddress(address common.Address, index int64) (string, error) {
	data, err := brizo.contractInstance.ReadData(&bind.CallOpts{}, address, big.NewInt(index))
	if err != nil {
		return "", errors.New("Failed to read data from contract.")
	}
	return data, nil
}

// ReadByHexAddress returns data on Brizo Chain by passing hex address (string) and index
func (brizo *brizoChain) ReadByHexAddress(hexAddress string, index int64) (string, error) {
	return brizo.ReadByAddress(common.HexToAddress(hexAddress), index)
}

func (brizo *brizoChain) ReadDataFromHashDict(key string) (string, error) {
	content, err := brizo.contractInstance.ReadDataFromHashDict(&bind.CallOpts{}, key)
	return strings.Join(content, ""), err
}
