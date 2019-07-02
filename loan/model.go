package loan

import "fmt"

// Loan funded loan
type Loan struct {
	ID         string
	Principal  uint64
	APR        float64 // need to use int64 to make sure the calculation is correct
	MatureDate string  // in format "YYYY-MM-DD"
	SoldDate   string  // in format "YYYY-MM-DD"
	TaxRate    float64
}

func (l Loan) String() string {
	return fmt.Sprintf(
		"{ID:%s Principal:%d APR:%.2f%% MatureDate:%s SoldDate:%s TaxRate:%.2f%%}",
		l.ID,
		l.Principal,
		l.APR,
		l.MatureDate,
		l.SoldDate,
		l.TaxRate,
	)
}
