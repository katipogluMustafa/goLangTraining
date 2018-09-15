package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	// Open gzip file that we want to uncompress
	// The file is a reader, but we could use any
	// data source. It is common for web servers
	// to return gzipped contents to save bandwidth
	// and in that case the data is not in a file
	// on the file system but is in a memory buffer
	gzipFile, err := os.Open("text.txt.gz")
	check(err)

	// Create a gzip reader on top of the file reader
	// Again, it could be any type reader though
	gzipReader, err := gzip.NewReader(gzipFile)
	check(err)
	defer gzipReader.Close()

	outfileWriter, err := os.Create("unzipped.txt")
	check(err)
	defer outfileWriter.Close()

	// Copy contents of gzipped file to output file
	_, err = io.Copy(outfileWriter, gzipReader)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
