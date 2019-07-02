package loan

// NumberOfDays calculate number of days between sold date to mature date.
// We expect matureDate should always be greater than soldDate.
// We expect both soldDate & matureDate in date format "YYYY-MM-DD".
// If we got error, always return 0 as number of days.
func NumberOfDays(soldDate, matureDate string) (uint64, error) {
	return 0, nil
}

// CalculateInterest calculate total interest based on principal, number of days, and flat apr.
func CalculateInterest(principal, numberOfDays uint64, apr float64) uint64 {
	return 0
}

// CalculateTax calculate tax based on given total interest and tax rate
func CalculateTax(totalInterest uint64, taxRate float64) uint64 {
	return 0
}
