package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"unicode"
)

func getRandomKey() uint8 {
	randVal := rand.Int() % 25 + 1
	return uint8(randVal)
}

func encode(inputFileName, outputFileName, keyFileName string) {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	keyFile, err := os.Create(keyFileName)
	if err != nil {
		fmt.Println(err)
	}
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		for i := 0; i < len(str); i++ {
			if unicode.IsLetter(rune(uint8(str[i]))) {
				step := getRandomKey()
				fmt.Fprintf(keyFile, "%d ", step)

				if str[i] + step > 'Z' {
					fmt.Fprintf(outputFile, "%s", string(str[i] + step + 'A' - 'Z' - 1))
				} else {
					fmt.Fprintf(outputFile, "%s", string(str[i] + step))
				}
			} else {
				fmt.Fprintf(outputFile, "%c", str[i])
			}
		}
		fmt.Fprintf(outputFile, "\n")
	}
}



func decode(inputFileName, outputFileName, keyFileName string) {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	keyFile, err := os.Open(keyFileName)
	if err != nil {
		fmt.Println(err)
	}
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(str)
		for i := 0; i < len(str); i++ {
			if unicode.IsLetter(rune(uint8(str[i]))) {
				var step uint8
				fmt.Fscan(keyFile, &step)
				step = 26 - step

				if str[i] + step > 'Z' {
					fmt.Fprintf(outputFile, "%s", string(str[i] + step + 'A' - 'Z' - 1))
				} else {
					fmt.Fprintf(outputFile, "%s", string(str[i] + step))
				}
			} else {
				fmt.Fprintf(outputFile, "%c", str[i])
			}
		}
		fmt.Fprintf(outputFile, "\n")
	}
}

func main() {
	encode("Text","Cripted.txt","key.txt")
	decode("Cripted.txt","Decoded.txt","key.txt")


}
