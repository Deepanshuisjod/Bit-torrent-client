// This file convert the .torrent file into the json type format
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"reflect"
)

// Global declaration of symbol stack array since it is used in many functions
var symbolstack []rune

func readTorrentFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	metaData := strings.Split(string(data), "\n")[0]
	return metaData, nil
}
func bencodedSymbolStack(symbol rune, symbolStack []rune) []rune {
	if symbol == 'e' {
		if len(symbolStack) > 0 {
			symbolStack = symbolStack[:len(symbolStack)-1] // Deleted the latest element in the stack if 'e' occurred
		}
	} else {
		symbolStack = append(symbolStack, symbol)
	}
	return symbolStack
}
func constructMap() map[string]interface{} {
	mp := make(map[string]interface{})
	return mp
}
func constructList() []string {
	//
}

func checkSymbolStack(symbol rune) {
	// We are going to implement binary search in bencodedSymbolStack to check the symbol in SymbolStack
	for _,value rune : range symbolStack {
		if (value == symbol){
			return 1
		}else return 0
	}
}
func bencodedData(map[string]interface{}, symbol rune) {
	// Combine all the sandwiched values in the dictionary, list, int & string into a json type format.
	var checkSymbolStack := checkSymbolStack(symbol)
	if (checkSymbolStack == 1){

	}
}

func distributeTypes(metaData string) {
	fmt.Println(len(metaData))

	// MetaData is always in bencoded-string format
	for _, value := range metaData {
		if value == 'd' {
			mp := constructMap()
			bencodedSymbolStack(value, symbolstack)
			bencodedData(mp,value)
		}
	}
}

func main() {
	fileName := "test_file.torrent"
	metaData, err := readTorrentFile(fileName)
	if err != nil {
		panic(err)
	}
	distributeTypes(metaData)
}
