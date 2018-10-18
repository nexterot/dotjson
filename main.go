// main
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"./libdotjson"
)

func main() {
	inFilenamePtr := flag.String("i", "input.json", "input filename")
	outFilenamePtr := flag.String("o", "output.json", "output filename")
	flag.Parse()
	jsonStr, fileErr := ioutil.ReadFile(*inFilenamePtr)
	if fileErr != nil {
		fmt.Println(fileErr)
		os.Exit(1)
	}
	result, formatErr := dotjson.DotFormat(string(jsonStr))
	if formatErr != nil {
		fmt.Println(formatErr)
		os.Exit(1)
	}
	fileErr = ioutil.WriteFile(*outFilenamePtr, []byte(result+"\n"), 0666)
	if fileErr != nil {
		fmt.Println(fileErr)
		os.Exit(1)
	}
}
