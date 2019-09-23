package txData

import (
	"encoding/json"
	"strings"
	"tboxcore"
)

type TxArg struct {
	PublicKey   []byte        `json:"PublicKey"`
	ContentHash string        `json:"ContentHash"`
	Signature   []byte        `json:"Signature"`
	Metadata    tboxcore.Meta `json:"Metadata"`
}

func TxArgToJson(txArg *TxArg) string {
	content, _ := json.Marshal(txArg)
	return string(content)
}

func TxArgFromJson(content string) (txArg *TxArg) {
	parser := json.NewDecoder(strings.NewReader(content))
	parser.Decode(&txArg)
	return
}
