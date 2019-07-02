package repayment

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/satriarrrrr/tezzz/loan"
)

func TestServiceWithRepository_GetProjection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		l loan.Loan
	}
	tests := []struct {
		name    string
		s       *ServiceWithRepository
		args    args
		want    Projection
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetProjection(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceWithRepository.GetProjection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceWithRepository.GetProjection() = %v, want %v", got, tt.want)
			}
		})
	}
}
