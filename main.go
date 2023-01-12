package stringutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type jsonData struct {
	Name string `json: "name"`
}

var testInfo []jsonData

func convertOrdinalNumber(i int) string {
	var ordinalValue string
	if i == 11 || i == 12 || i == 13 {
		ordinalValue = "th"
	} else if i%10 == 1 {
		ordinalValue = "st"
	} else if i%10 == 2 {
		ordinalValue = "nd"
	} else if i%10 == 3 {
		ordinalValue = "rd"
	} else {
		ordinalValue = "th"
	}

	return ordinalValue
}

func separateStringAndNumber(name string) (letters, numbers string) {
	var l, n []rune
	for _, c := range name {
		switch {
		case c >= 'A' && c <= 'Z':
			l = append(l, c)
		case c >= 'a' && c <= 'z':
			l = append(l, c)
		case c >= '0' && c <= '9':
			n = append(n, c)
		}
	}
	return string(l), string(n)
}

func stringutils(url string) {
	JsonData, _ := os.Open(url)
	byteData, _ := ioutil.ReadAll(JsonData)
	defer JsonData.Close()
	json.Unmarshal(byteData, &testInfo)
	for i := 0; i < len(testInfo); i++ {
		ordinalValue := convertOrdinalNumber(i + 1)
		strVale, numValue := separateStringAndNumber(testInfo[i].Name)
		fmt.Println(strconv.Itoa(i+1)+ordinalValue, "name is ", strVale, numValue)
	}

}
