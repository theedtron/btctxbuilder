package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"log"
)

// Function to generate redeem script
func generateRedeemScript(lockHex string) []byte {
	redeemScriptBytes, _ := txscript.NewScriptBuilder().
		AddOp(txscript.OP_SHA256).
		AddData([]byte(lockHex)).
		AddOp(txscript.OP_EQUAL).
		Script()
	return redeemScriptBytes
}

// Function to derive address from redeem script
func deriveAddress(redeemScript []byte) string {
	p2shScript, err := txscript.NewScriptBuilder().
		AddOp(txscript.OP_HASH160).
		AddData(redeemScript).
		AddOp(txscript.OP_EQUAL).
		Script()
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(p2shScript)
}

// Function to construct transaction to send bitcoins to the address
func constructSendTransaction(address string, amount int64) *wire.MsgTx {
	// This function simulates constructing a transaction
	tx := wire.NewMsgTx(wire.TxVersion)
	output := wire.NewTxOut(amount, []byte(address))
	tx.AddTxOut(output)
	return tx
}

// Test function to validate all functions
func testFunctions() {
	lockHex := "427472757374204275696c64657273" // Bytes encoding of "Btrust Builders"
	redeemScript := generateRedeemScript(lockHex)
	fmt.Println("Redeem Script:", hex.EncodeToString(redeemScript))

	address := deriveAddress(redeemScript)
	fmt.Println("Derived Address:", address)

	// Construct transaction to send bitcoins
	sendTx := constructSendTransaction(address, 10000000) // 0.1 BTC in Satoshis
	fmt.Println("Send Transaction:", sendTx)
}

func main() {
	testFunctions()
}
