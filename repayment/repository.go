package repayment

type Repository interface {
	GetProjection(loanID string) (Projection, error)
	StoreProjection(p Projection) error
}
