package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

// ColumnSizes struct stores data about the length of columns.
type ColumnSizes struct {
	Name int
	Debt int
}

func main() {
	directory := flag.String("dir", "./", "File directory")
	flag.Parse()

	files, err := ioutil.ReadDir(*directory)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		name := strings.Split(file.Name(), ".")

		// Сheck the file extension. If file is json extension, then parse file and display the data.
		if (len(name) > 1) && ((name[1] == "json") || (name[1] == "js")) {
			jsonFile, err := os.Open(filepath.Join(*directory, file.Name()))
			if err != nil {
				fmt.Println(err)
				return
			}

			defer jsonFile.Close()

			byteValue, _ := ioutil.ReadAll(jsonFile)

			if err != nil {
				fmt.Println(err)
				return
			}

			data := make(map[string][]map[string]interface{})

			err = json.Unmarshal(byteValue, &data)

			if err != nil {
				fmt.Println(err)
				return
			}

			length := getLengthOfColumns(data)

			fmt.Printf("|%"+fmt.Sprint(length.Name+10)+"v|%15v|%15v|%"+fmt.Sprint(length.Debt+10)+"v|\n", "Name", "Group", "Avg", "Debt")
			fmt.Printf("|%"+fmt.Sprint(length.Name+10)+"v|%15v|%15v|%"+fmt.Sprint(length.Debt+10)+"v|\n", "", "", "", "")

			for _, item := range data["items"] {
				var debts string

				debtsInterface, err := getSliceFromInterface(item["debt"])

				if err != nil {
					debts = item["debt"].(string)
				} else {
					for _, debt := range debtsInterface {
						debts = debts + " " + debt.(string)
					}
				}

				fmt.Printf("|%"+fmt.Sprint(length.Name+10)+"v|%15v|%15v|%"+fmt.Sprint(length.Debt+10)+"v|\n", item["name"], item["group"], item["avg"], debts)
			}
		}
	}
}

// getSliceFromInterface func get []interface{} from interface if it posible.
func getSliceFromInterface(slice interface{}) ([]interface{}, error) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil, fmt.Errorf("InterfaceSlice() given a non-slice type")
	}

	if s.IsNil() {
		return nil, fmt.Errorf("InterfaceSlice() given a nil Interface")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret, nil
}

// getLengthOfColumns get length of columns depending on the length of the string.
func getLengthOfColumns(data map[string][]map[string]interface{}) (cs ColumnSizes) {
	var lengthsForName []int

	var lengthsForDebt []int

	for _, item := range data["items"] {
		lengthsForName = append(lengthsForName, len(item["name"].(string)))

		var debts string

		debtsInterface, err := getSliceFromInterface(item["debt"])

		if err != nil {
			debts = item["debt"].(string)
		} else {
			for _, deb := range debtsInterface {
				debts += deb.(string)
			}
		}

		lengthsForDebt = append(lengthsForDebt, len(debts))
	}

	sort.Ints(lengthsForName)

	sort.Ints(lengthsForDebt)

	cs.Name = lengthsForName[len(lengthsForName)-1]

	cs.Debt = lengthsForDebt[len(lengthsForDebt)-1]

	return cs
}
