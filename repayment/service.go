package repayment

import (
	"github.com/satriarrrrr/tezzz/loan"
)

// Service contains business logic
type Service interface {
	GetProjection(l loan.Loan) (Projection, error)
}

// ServiceWithRepository implements Service that use Repository to store data to the data storage
type ServiceWithRepository struct {
	Repository Repository
}

// GetProjection return the projection of given loan
// - get projection
// - calculate daydiff
// - calculate interest
// - calculate tax
// - store projection
func (s *ServiceWithRepository) GetProjection(l loan.Loan) (Projection, error) {
	var p Projection

	return p, nil
}
