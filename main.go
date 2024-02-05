package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// Function to generate redeem script
func generateRedeemScript(lockHex string) []byte {
	redeemScript, err := txscript.NewScriptBuilder().
		AddOp(txscript.OP_SHA256).
		AddData([]byte(lockHex)).
		AddOp(txscript.OP_EQUAL).
		Script()
	if err != nil {
		log.Fatal(err)
	}
	return redeemScript
}

// Function to derive address from redeem script
func deriveAddress(redeemScript []byte) string {
	scriptHash := btcutil.Hash160(redeemScript)
	
	p2shScript, err := btcutil.NewAddressScriptHash(scriptHash, &chaincfg.TestNet3Params)
	if err != nil {
		log.Fatal(err)
	}
	return p2shScript.EncodeAddress()
}

// Function to construct transaction to send bitcoins to the address
func constructSendTransaction(address string, amount int64) *wire.MsgTx {
	// This function simulates constructing a transaction
	txBuilder := wire.NewMsgTx(wire.TxVersion)
	txOutput := wire.NewTxOut(amount, []byte(address))
	txBuilder.AddTxOut(txOutput)
	return txBuilder
}

// Function to construct transaction with spending conditions
func constructSpendingTransaction(prvTx *wire.MsgTx, redeemScript []byte) *wire.MsgTx {
	txBuilder := wire.NewMsgTx(wire.TxVersion)
	prvTxHash := prvTx.TxHash()
	outPoint := wire.NewOutPoint(&prvTxHash,0)
	txInput := wire.NewTxIn(outPoint, nil, nil)
	txBuilder.AddTxIn(txInput)

	txOut := wire.NewTxOut(90000, []byte("BtrusBuilderLocking"))
	txBuilder.AddTxOut(txOut)

	script, err := txscript.NewScriptBuilder().
		AddData([]byte("MyGenSignature")).
		AddData(redeemScript).
		Script()
	if err != nil {
		log.Fatal(err)
	}

	txInput.SignatureScript = script

	return txBuilder
}

func main() {
	lockHex := "427472757374204275696c64657273" // Bytes encoding of "Btrust Builders"
	redeemScript := generateRedeemScript(lockHex)
	fmt.Println("Redeem Script:", hex.EncodeToString(redeemScript))

	address := deriveAddress(redeemScript)
	fmt.Println("Derived Address:", address)

	// Construct transaction to send bitcoins
	sendTx := constructSendTransaction(address, 1000000) // 0.01 BTC in Satoshis
	fmt.Println("Send Transaction:", sendTx)

	spendTx := constructSpendingTransaction(sendTx, redeemScript)
	fmt.Println("Spend Transaction:", spendTx)
}

