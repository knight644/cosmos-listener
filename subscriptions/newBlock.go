package subscriptions

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cosmos-listener/configs"
	"github.com/cosmos-listener/record"
	"github.com/cosmos-listener/txdecoder"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	tenderminttypes "github.com/tendermint/tendermint/types"
)

func NewBlockSub() {
	// Create blank context (Top level)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Create query params for subscription
	queryParams := "tm.event = 'NewBlock'"

	// Subscribe to the events
	blocks, err := Client.Subscribe(ctx, configs.ChainName, queryParams)

	if err != nil {
		log.Println("Failed to subscribe to new block events: ", err)
	} else {
		log.Println("Successfully subscribed to new block events")
	}

	handleNewBlockEvent(blocks)
}

func handleNewBlockEvent(block <-chan coretypes.ResultEvent) {
	for val := range block {
		//log.Println("new block")
		// Parse block data
		data, _ := val.Data.(tenderminttypes.EventDataNewBlock)
		blockDetails, err := GetPrevBlockDetails(data.Block.Height)
		if err != nil {
			log.Println(err)
		}

		txs := blockDetails.Block.Data.Txs

		for _, tx := range txs {
			// log.Println("tx ", txhash)
			gasUsed, gasWanted, txSize, err := getTxGasAndBytes(tx.Hash())
			if err != nil {
				log.Println(err)
			}

			// log.Println(gasUsed, gasWanted, txSize)
			// log.Println(tx.String())

			var txDetails txdecoder.CosmosTx
			err = getTxDetails(fmt.Sprintf("%X", []byte(tx)), &txDetails)
			if err != nil {
				log.Println(err)
			}

			// log.Println(txDetails.Body.Messages)
			msgTypes, multipleMsgs := getMessageType(txDetails.Body.Messages)
			//log.Println(msgTypes, multipleMsgs)

			// Retrieve block height
			height := data.Block.Height

			// Retrieve block time
			time := data.Block.Time

			writeToFile(height, time, msgTypes, multipleMsgs, txSize, gasUsed, gasWanted)
		}
	}
}

func getTxGasAndBytes(txHash []byte) (gasUsed int64, gasWanted int64, txSize int64, err error) {
	tx, err := GetTransactionDetails(txHash)
	if err != nil {
		// log.Println(err)
		return 0, 0, 0, err
	}

	// log.Println(tx)
	gasUsed = tx.TxResult.GasUsed
	gasWanted = tx.TxResult.GasWanted
	txSize = int64(len([]byte(tx.Tx.String())))

	// log.Println("gas used: ", gasUsed, "gas wanted: ", gasWanted)
	// log.Println(txSize, len([]byte(tx.Tx.String())))
	return gasUsed, gasWanted, txSize, nil
}

func getTxDetails(txString string, txDetails *txdecoder.CosmosTx) error {
	decoder := txdecoder.DefaultDecoder

	// log.Println(txString)

	hexTxbytes, err := hex.DecodeString(txString)
	if err != nil {
		// log.Println("1: ", err)
		return err
	}

	txDecoded, err := decoder.Decode(hexTxbytes)
	if err != nil {
		// log.Println("2: ", err)
		return err
	}

	txBytes, err := txDecoded.MarshalToJSON()
	if err != nil {
		// log.Println("3: ", err)
		return err
	}

	err = json.Unmarshal(txBytes, &txDetails)
	if err != nil {
		// log.Println("4: ", err)
		return err
	}

	return nil
}

func getMessageType(msgs []map[string]interface{}) (msgTypes []string, multipleMsgs bool) {
	multipleMsgs = false

	if len(msgs) != 1 {
		// log.Println("multiple messages in transaction")
		multipleMsgs = true
	}

	for _, msg := range msgs {
		msgTypes = append(msgTypes, fmt.Sprintf("%v", msg["@type"]))
	}
	return msgTypes, multipleMsgs
}

func writeToFile(height int64, time time.Time, txTypes []string, multimsg bool, txSize int64, gasUsed int64, gasWanted int64) {
	var monitorDataRow []string
	txType := txTypes[0]

	for index, msg := range txTypes {
		if index < 1 {
			continue
		}
		txType = txType + " + " + msg
	}

	monitorDataRow = append(monitorDataRow, strconv.FormatInt(height, 10), time.String(), txType,
		strconv.FormatInt(txSize, 10), strconv.FormatInt(gasUsed, 10), strconv.FormatInt(gasWanted, 10))
	record.WriteMonitorData(monitorDataRow)
}
