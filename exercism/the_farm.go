package thefarm

import "errors"
import "fmt"

// To retrieve the total amount of fodder for all the cows.
// This is done by calling the FodderAmount method and passing the number of cows.
// This amount needs to be multiplied with the FatteningFactor. 
// These two values and the number of cows, you can now calculate the amount of food per cow.
func DivideFood(fc FodderCalculator, cowCount int) (float64, error) {
    fmt.Printf("cowCount:%d\n", cowCount)
    fa, err := fc.FodderAmount(cowCount)
    fmt.Printf("fa:%f\n", fa)
    if err != nil {
        return 0.0, err
    }
    ff, err2 := fc.FatteningFactor()
    fmt.Printf("ff:%f\n", ff)
    if err2 != nil {
        return 0.0, err2
    }

    // Calculate the average of cow's fodder
    result := fa * ff / float64(cowCount)
    return result, nil
}

// If the number of cows passed in is greater than 0,
// the function should call DivideFood and return the results of that call.
// If the number of cows is 0 or less,
// the function should return an error with message "invalid number of cows".
func ValidateInputAndDivideFood(fc FodderCalculator , cowCount int) (float64, error) {
    switch cowCount > 0 {
        case true:
        	return DivideFood(fc, cowCount)
        default:
			return 0, errors.New("invalid number of cows")
    }
} 

// InvalidCowsError is custom error
type InvalidCowsError struct {
    cowCount int
    message string
}

func (ice *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", ice.cowCount, ice.message)
}

// If the number of cows is less than 0,
// the function returns an InvalidCowsError with the custom message set to "there are no negative cows".
// If the number of cows is 0,
// the function returns an InvalidCowsError with the custom message set to "no cows don't need food".
// Otherwise, the function returns nil to indicate that the validation was successful.
func ValidateNumberOfCows(cowCount int) error {
    if cowCount < 0 {
        return &InvalidCowsError{cowCount, "there are no negative cows"}
    } else if cowCount == 0 {
        return &InvalidCowsError{cowCount, "no cows don't need food"}
    }
    return nil
}





package thefarm

// This file contains types used in the exercise and tests and should not be modified.

// FodderCalculator provides helper methods to determine the optimal
// amount of fodder to feed cows.
type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}