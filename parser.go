// This file convert the .torrent file into the json type format
package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
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
func bencodedSymbolStack(symbol rune, symbolStack []rune) {
	if symbol == 'e' {
		if len(symbolStack) > 0 {
			symbolStack = symbolStack[:len(symbolStack)-1] // Deleted the latest element in the stack if 'e' occurred
		}
	} else {
		symbolStack = append(symbolStack, symbol)
	}
}
func constructMap() map[string]interface{} {
	mp := make(map[string]interface{})
	return mp
}
func constructList() []string {
	//
}

// Unless the symbol is found in symbolStack it will return 1 or else return 0
func checkSymbolStack(symbolStack []rune, symbol rune) int {
	for _, value := range symbolStack {
		if value == symbol {
			return 1
		}
	}
	return 0
}

func bencodedData(data interface{}, symbol rune) {
	// Combine all the sandwiched values in the dictionary, list, int & string into a json type format.
	datatype := data.(type)
	checkSymbolStack := checkSymbolStack(symbol)
	if checkSymbolStack == 1 {
		var latest_symbol_inStack rune = symbolstack[len(symbolstack)-1]
	}
}

func distributeTypes(metaData string) {
	fmt.Println(len(metaData))

	// MetaData is always in bencoded-string format
	for i := 0; i < len(metaData); i++ {
		value := metaData[i]
		value_type := reflect.TypeOf(value)

		if value == 'd' {
			mp := constructMap()
			bencodedSymbolStack(value, symbolstack)
			bencodedData(mp, value)
		}
		if value_type.Name() == "int" {
			j := i + value + 1
			mapString := metaData[i+2 : j]
			i = j + 1
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
