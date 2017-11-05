package models

// Domain reporesents a domain record with some categories.
type Domain struct {
	Cname      string
	Categories []string
}
