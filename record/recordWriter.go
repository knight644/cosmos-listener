package record

import (
	"encoding/csv"
	"log"
	"os"
)

var monitorCSVWriter *csv.Writer

func CreateMonitorCSVWriter() {
	// TODO: Generate new files everytime programme is run
	csvFile, err := os.Create("monitorTx.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	monitorCSVWriter = csv.NewWriter(csvFile)

	var headerRow []string
	headerRow = append(headerRow, "Block Height", "Time", "Tx type", "Tx Size", "Gas used", "Gas wanted")

	WriteMonitorData(headerRow)
}

func WriteMonitorData(dataRow []string) {
	if err := monitorCSVWriter.Write(dataRow); err != nil {
		log.Fatalln("error writing record to file", err)
	}
	monitorCSVWriter.Flush()
}
