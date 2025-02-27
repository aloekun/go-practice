package interest

// import "fmt"

// BalanceRate data struct
type BalanceRateStruct struct {
	min *float64
	max *float64
	ret *float32
}

// BalanceRate data
// If data will be added, only add this array.
var BalanceRate = []BalanceRateStruct{
	{
		min: nil,
		max: float64Ptr(0.0),
		ret: float32Ptr(3.213),
	},
	{
		min: float64Ptr(0.0),
		max: float64Ptr(1000.0),
		ret: float32Ptr(0.5),
	},
	{
		min: float64Ptr(1000.0),
		max: float64Ptr(5000.0),
		ret: float32Ptr(1.621),
	},
	{
		min: float64Ptr(5000.0),
		max: nil,
		ret: float32Ptr(2.475),
	},
}

func float64Ptr(f float64) *float64 {
	return &f
}
func float32Ptr(f float32) *float32 {
	return &f
}

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	// fmt.Printf("balance: %f\n", balance)
	for _, v := range BalanceRate {
		if v.min != nil && balance < *(v.min) {
			// fmt.Printf("*(v.min): %f, continue\n", *(v.min))
			continue
		}
		if v.max != nil && *(v.max) <= balance {
			// fmt.Printf("*(v.max): %f, continue\n", *(v.max))
			continue
		}
		if v.min != nil {
			// fmt.Printf("*(v.min): %f\n", *(v.min))
		}
		if v.max != nil {
			// fmt.Printf("*(v.max): %f\n", *(v.max))
		}
		if v.ret != nil {
			// fmt.Printf("*(v.ret): %f\n", *(v.ret))
		}
		return *(v.ret)
	}
	// fmt.Printf("return method last\n")
	return 0.0
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	// Calc InterestRate Value
	irPercentage := float64(InterestRate(balance))
	ir := irPercentage / 100.0

	return balance * ir
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	sumBalance := balance
	year := 0
	for {
		if targetBalance <= sumBalance {
			return year
		}
		// Safety for infinite loop
		if 1000 < year {
			return -1
		}
		sumBalance = AnnualBalanceUpdate(sumBalance)
		year++
	}
}
