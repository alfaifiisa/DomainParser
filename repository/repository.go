package repository

import (
	"github.com/alfaifiisa/DomainParser/models"
)

var domains []models.Domain

// StoreDomain creates a Domain object and stores it in the domains slice.
func StoreDomain(cname string, categories []string) {
	domains = append(domains,
		models.Domain{
			Cname:      cname,
			Categories: categories,
		})
}

// FilterByCategory returns only the selected categories
func FilterByCategory(filters ...string) []models.Domain {
	// storing filter in to a map to make the search faster
	filtersMap := make(map[string]bool, 0)
	for _, filter := range filters {
		filtersMap[filter] = true
	}

	// creating the result set
	results := make([]models.Domain, 0)

	// iterating over all stored domains
	for _, domain := range domains {
		// adding current domain as a record to the result set
		record := models.Domain{Cname: domain.Cname}
		// iterating over domain categories to find the match.
		for _, category := range domain.Categories {
			if filtersMap[category] == true {
				// adding the found category match to the record
				record.Categories = append(record.Categories, category)
			}
		}
		// adding the processed records to the results
		results = append(results, record)
	}
	return results
}
