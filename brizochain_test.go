package brizochain

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// TestBrizoChain_Write tests Write function of brizochain
func TestBrizoChain_Write(t *testing.T) {
	testcases := []string{
		"1st testing message",
		"2nd testing message",
		"3rd testing message",
	}

	brizoChain, err := NewBrizoChain()
	if err != nil {
		t.Error("Fail to initialize Brizo Chain service.")
	}

	counterBefore, err := brizoChain.contractInstance.DataCounter(&bind.CallOpts{}, brizoChain.txSigner.From)
	if err != nil {
		t.Error("Failed to invoke DataCounter.")
	}
	fmt.Println("Counter before Write =", counterBefore)

	for _, msg := range testcases {
		if err = brizoChain.Write(msg); err != nil {
			t.Error("Failed to write data.")
		}
	}

	counterAfter, err := brizoChain.contractInstance.DataCounter(&bind.CallOpts{}, brizoChain.txSigner.From)
	if err != nil {
		t.Error("Failed to invoke DataCounter.")
	}
	fmt.Println("Counter after Write =", counterAfter)
	if big.NewInt(0).Sub(counterAfter, counterBefore).Cmp(big.NewInt(int64(len(testcases)))) != 0 {
		t.Error("Failed to write data, counterAfter != counterBefore + len(testcases).")
	}

}

// TestBrizoChain_Read tests Read function of brizochain
func TestBrizoChain_Read(t *testing.T) {
	testcases := []struct {
		index int64
		data  string
	}{
		{0, "1st testing message"},
		{1, "2nd testing message"},
		{2, "3rd testing message"},
	}

	brizoChain, err := NewBrizoChain()
	if err != nil {
		t.Error("Fail to initialize blockchain service.")
	}

	for _, mapping := range testcases {
		if msg, err := brizoChain.Read(mapping.index); msg != mapping.data {
			if err != nil {
				t.Error("Failed to read data.")
			}
			t.Error("Data mismatch.")
		}
	}
}
