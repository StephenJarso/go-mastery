package testing_basics

// CalculateTax calculates the income tax based on brackets:
// Up to $10,000: 10%
// $10,001 to $50,000: 15%
// Above $50,000: 25%
func CalculateTax(income float64) float64 {
	if income <= 0 {
		return 0
	}
	if income <= 10000 {
		return income * 0.10
	}
	if income <= 50000 {
		return 1000 + (income-10000)*0.15
	}
	return 1000 + 6000 + (income-50000)*0.25
}
