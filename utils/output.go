package utils

import (
	"fmt"

	"github.com/alfaifiisa/DomainParser/models"
)

// FormatedDomainSTDOutput outputs the domain strucure in a formatted way.
func FormatedDomainSTDOutput(domain models.Domain) {
	var categories string
	if len(domain.Categories) > 0 {
		for _, category := range domain.Categories {
			categories += "\t" + category
		}
	} else {
		categories = "\tNON"
	}
	fmt.Println(domain.Cname, categories)
}
