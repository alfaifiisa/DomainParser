package processors

import (
	"bufio"
	"os"
	"strings"

	"github.com/alfaifiisa/DomainParser/repository"
)

// ProcessInputFromFile parse the input and stores it to the repository.
func ProcessInputFromFile(filename string) error {
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
		records := strings.Fields(record) // works with white spaces, if another character is used, there is a different approach
		// slicing the result and passing it the StoreDomain method.
		repository.StoreDomain(records[0], records[1:])
	}
	return nil
}
