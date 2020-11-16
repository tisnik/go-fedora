package main

import (
	"log"
	"os"

	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
)

const defaultOutputFile = "flat.parquet"

// Record represents one record stored in Parquet file
type Record struct {
	Active bool `parquet:"name=active, type=BOOLEAN"`
}

func writeRecords(pw *writer.ParquetWriter, n int) {
	// create report structure to be stored in Parquet file
	record := Record{
		Active: false,
	}

	for i := 0; i < n; i++ {
		record.Active = i%2 == 0
		// write the record structure into Parquet file
		err := pw.Write(record)
		if err != nil {
			log.Println("Write into Parquet error", err)
		}
	}
}

// stopWrite function stop writing into Parquet file
func stopWrite(pw *writer.ParquetWriter) {
	err := pw.WriteStop()

	// most write errors are caught at this time
	if err != nil {
		log.Println("WriteStop error", err)
	}
}

func createAndWriteIntoParquetFile(filename string, records int) {
	w, err := os.Create(filename)
	if err != nil {
		log.Println("Can't create local file", err)
		return
	}

	defer w.Close()

	// initialize Parquet file writer
	pw, err := writer.NewParquetWriterFromWriter(w, new(Record), 1)
	if err != nil {
		log.Println("Can't create parquet writer", err)
		return
	}

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = parquet.CompressionCodec_UNCOMPRESSED

	defer stopWrite(pw)

	writeRecords(pw, records)

	log.Println("Write Finished")
}

func main() {
	createAndWriteIntoParquetFile("0records.parquet", 0)
	createAndWriteIntoParquetFile("1record.parquet", 1)
	createAndWriteIntoParquetFile("10records.parquet", 10)
	createAndWriteIntoParquetFile("100records.parquet", 100)
	createAndWriteIntoParquetFile("1000records.parquet", 1000)
	createAndWriteIntoParquetFile("10000records.parquet", 10000)
	createAndWriteIntoParquetFile("100000records.parquet", 100000)
}