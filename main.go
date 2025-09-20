package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var help bool
	var inputPath, outputPath string

	flag.BoolVar(&help, "h", false, "Tampilkan cara menggunakan")
	flag.StringVar(&inputPath, "i", "", "Lokasi file JSON sebagai input")
	flag.StringVar(&outputPath, "o", "", "Lokasi file JSON sebagai output")

	flag.Parse()

	if help || inputPath == "" || outputPath == "" {
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {

	fmt.Println("Usage: mockdata -i <input file> -o <output file>")
	fmt.Println("-i: File input berupa JSON sebagai template")
	fmt.Println("-o: File output berupa JSON sebagai template")
}
