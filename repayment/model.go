package repayment

// Projection holds information about
type Projection struct {
	LoanID   string
	Interest uint64
	Tax      uint64
	Total    uint64
}
