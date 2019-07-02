package loan

import (
	"testing"
)

func TestNumberOfDays(t *testing.T) {
	type args struct {
		soldDate   string
		matureDate string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		// 1. Invalid format of soldDate
		// 2. Invalid format of matureDate
		// 3. Invalid format of both soldDate & matureDate
		// 4. soldDate is greater than matureDate
		// 5. soldDate == matureDate
		// 6. matureDate is greater than soldDate
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NumberOfDays(tt.args.soldDate, tt.args.matureDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("NumberOfDays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NumberOfDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateInterest(t *testing.T) {
	type args struct {
		principal    uint64
		numberOfDays uint64
		apr          float64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// 1. Number of days = 0
		{
			name: "1. Number of days = 0",
			args: args{
				principal:    550000,
				numberOfDays: 0,
				apr:          16.8,
			},
			want: 0,
		},
		// 2. Number of days > 0
		{
			name: "2. Number of days > 0",
			args: args{
				principal:    550000,
				numberOfDays: 21,
				apr:          7.18,
			},
			want: 2303,
		},
		{
			name: "2. Number of days > 0",
			args: args{
				principal:    550000,
				numberOfDays: 31,
				apr:          7.18,
			},
			want: 3400,
		},
		{
			name: "2. Number of days > 0",
			args: args{
				principal:    550000,
				numberOfDays: 30,
				apr:          7.18,
			},
			want: 3290,
		},
		{
			name: "2. Number of days > 0",
			args: args{
				principal:    100000,
				numberOfDays: 24,
				apr:          7.18,
			},
			want: 478,
		},
		{
			name: "2. Number of days > 0",
			args: args{
				principal:    100000,
				numberOfDays: 31,
				apr:          7.18,
			},
			want: 618,
		},
		{
			name: "2. Number of days > 0",
			args: args{
				principal:    100000,
				numberOfDays: 30,
				apr:          7.18,
			},
			want: 598,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateInterest(tt.args.principal, tt.args.numberOfDays, tt.args.apr); got != tt.want {
				t.Errorf("CalculateInterest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateTax(t *testing.T) {
	type args struct {
		totalInterest uint64
		taxRate       float64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// 1. Total interest: 0, tax rate: 0%				-> 0
		{
			name: "1. Total interest: 0, tax rate: 0%",
			args: args{
				totalInterest: 0,
				taxRate:       0,
			},
			want: 0,
		},
		// 2. Total interest: 1000000, tax rate: 0%			-> 0
		{
			name: "2. Total interest: 1000000, tax rate: 0%",
			args: args{
				totalInterest: 1000000,
				taxRate:       0,
			},
			want: 0,
		},
		// 3. Total interest: 0, tax rate: 15%				-> 0
		{
			name: "3. Total interest: 0, tax rate: 15%",
			args: args{
				totalInterest: 0,
				taxRate:       15,
			},
			want: 0,
		},
		// 4. Total interest: 1000000, tax rate: 15%		-> 150000
		{
			name: "4. Total interest: 1000000, tax rate: 15%",
			args: args{
				totalInterest: 1000000,
				taxRate:       15,
			},
			want: 150000,
		},
		// 5. Total interest: 123123123, tax rate: 12.5%	-> 15390390(.375)
		{
			name: "5. Total interest: 123123123, tax rate: 12.5%",
			args: args{
				totalInterest: 123123123,
				taxRate:       12.5,
			},
			want: 15390390,
		},
		// 6. Total interest: 123123456, tax rate: 13.6%	-> 1674479001(.6)
		{
			name: "6. Total interest: 123123456, tax rate: 13.6%",
			args: args{
				totalInterest: 123123456,
				taxRate:       13.6,
			},
			want: 1674479001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateTax(tt.args.totalInterest, tt.args.taxRate); got != tt.want {
				t.Errorf("CalculateTax() = %v, want %v", got, tt.want)
			}
		})
	}
}
