package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func pickWinner(lines []string) {

	//We split the strings
	onlyString := strings.Join(lines, " ")
	stringSlice := strings.Split(onlyString, " ")

	var intarray = []int{}

	i := 1

	for i <= len(stringSlice) {
		//we convert the string arrays to int arrays
		j, err := strconv.Atoi(stringSlice[i])
		if err != nil {
			panic(err)
		}
		intarray = append(intarray, j)

		//we increment by 2 because we want to ignore the names
		i = i + 2
	}

	sort.Sort(sort.Reverse(sort.IntSlice(intarray)))

	k := 0

	fmt.Print("The highest scores in order from first to last are: \n")

	for k < 3 {

		fmt.Println(intarray[k])
		k++
	}

}

func openFile(text string) {

	file, err := os.Open(text)

	if err != nil {
		log.Fatalf("Failed opening the file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	pickWinner(txtlines)
}

func main() {

	openFile("result.txt")

}
