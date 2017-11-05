package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alfaifiisa/DomainParser/processors"
	"github.com/alfaifiisa/DomainParser/repository"
	"github.com/alfaifiisa/DomainParser/utils"
)

func main() {

	var filename string
	flag.StringVar(&filename, "filename", "./input.txt", "filename if the input file")
	flag.Parse()

	// Single function approach:
	fmt.Println("Approach 1:")
	parseDomainInput(filename)

	// Strucural approach:
	fmt.Println("\nApproach 2:")
	err := processors.ProcessInputFromFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, domain := range repository.FilterByCategory("social", "arabic") {
		utils.FormatedDomainSTDOutput(domain)
	}
}

func parseDomainInput(filename string) error {
	// opening file for input.
	input, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer input.Close()

	// scanning the file line by line
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// spliting the line by the spaces
		record := scanner.Text()
		records := strings.Fields(record) // works with white spaces if another character is used, there is a different approach
		// slicing the result and passing it the StoreDomain method.
		cname := records[0]
		categories := records[1:]

		var categoriesString string
		// iterating over the splitted categories
		for _, category := range categories {
			if category == "social" || category == "arabic" {
				// concatinating the found category
				categoriesString += "\t" + category
			}
		}
		// adding a NON value for unmatched categories.
		if categoriesString == "" {
			categoriesString = "\tNON"
		}
		fmt.Println(cname, categoriesString)
	}
	return nil
}
