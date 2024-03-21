// This file convert the .torrent file into the json type format
package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// Global declaration of symbol stack array since it is used in many functions
var symbolstack []rune
var numArray []int

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
	list := make([]string, 1)
	return list
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

func bencodedData(data interface{}, symbol interface{}) {
	// Combine all the sandwiched values in the dictionary, list, int & string into a json type format.
	dataType := reflect.TypeOf(data)
	checkSymbolStack := checkSymbolStack(symbol)
	checkFirstSymbol := symbolstack[0]
	if checkFirstSymbol == 'd' {
		if checkSymbolStack == 1 {
			dataMap, ok := data.(map[string]interface{})
			if !ok {
				if dataType.Kind() == reflect.String {
					inStackValue, count := checknumArray(symbol)
					var keyValue string
					if count%2 != 0 {
						keyValue = inStackValue
						dataMap[keyValue] = ""
					}
					dataMap[keyValue] = inStackValue

				}
			}
		}
	}

}

func distributeTypes(metaData string) {
	fmt.Println(len(metaData))

	// MetaData is always in bencoded-string format
	for i := 0; i < len(metaData); i++ {
		value := metaData[i]
		var numStr string

		// For Dictionary
		if value == 'd' {
			mp := constructMap()
			bencodedSymbolStack(value, symbolstack)
			bencodedData(mp, value)
		}

		// For String
		if value >= '0' && value <= '9' {
			char := value
			for char != ':' {
				numStr += string(value)
				i++
				break
			}
			numInt, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			numArray = append(numArray, numInt)
			numStr = ""

		}

		// For List
		if value == 'l' {
			list := constructList()

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
