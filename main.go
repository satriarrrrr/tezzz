package main

import (
	"flag"
	"fmt"

	uuid "github.com/satori/go.uuid"

	"github.com/satriarrrrr/tezzz/loan"
	"github.com/satriarrrrr/tezzz/repayment"
)

func main() {
	id := uuid.NewV4()
	loan := loan.Loan{
		ID: id.String(),
	}

	flag.Uint64Var(&loan.Principal, "principal", 0, "-principal=xxxx, default value is 0")
	flag.Float64Var(&loan.APR, "apr", 7.18, "-apr=x.xx, default value is 7.18")
	flag.StringVar(&loan.MatureDate, "mature", "", "-mature=YYYY-MM-DD")
	flag.StringVar(&loan.SoldDate, "sold", "", "-sold=YYYY-MM-DD")
	flag.Float64Var(&loan.TaxRate, "taxrate", 0, "-taxrate=x.xx, default value is 0")
	flag.Parse()

	var r repayment.Repository

	s := repayment.ServiceWithRepository{
		Repository: r,
	}

	p, err := s.GetProjection(loan)
	if err != nil {
		fmt.Printf("you got error: %s", err)
		return
	}

	fmt.Printf("Projection for loan: %s\n", loan)
	fmt.Printf("Principal : %d\n", loan.Principal)
	fmt.Printf("Interest  : %d\n", p.Interest)
	fmt.Printf("Tax       : -%d\n", p.Tax)
	fmt.Printf("---------------------- +\n")
	fmt.Printf("Total     : %d\n", p.Total)
}
