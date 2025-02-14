package lasagna

// Calc PreparationTime function
// The function should return the estimate for the total preparation time based on the number of layers as an int.
// If the average preparation time is passed as 0 (the default initial value for an int),
// then the default value of 2 should be used.
func PreparationTime(layers []string, avgTime int) int {
	// If avgTime is zero, avgTime is default time.
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

// Calc Quantities function
// For each noodle layer in your lasagna, you will need 50 grams of noodles.
// For each sauce layer in your lasagna, you will need 0.2 liters of sauce.
func Quantities(layers []string) (int, float64) {
	var noodleQuantity int = 0
	var sauceQuantity float64 = 0.0
	const NoodleGrams int = 50
	const SauceGrams float64 = 0.2
	for _, v := range layers {
		if v == "noodles" {
			noodleQuantity += NoodleGrams
		}
		if v == "sauce" {
			sauceQuantity += SauceGrams
		}
	}
	return noodleQuantity, sauceQuantity
}

// Make my recipe by friend recipe
// Last recipe of friend is special recipe.
// Last recipe of me is empty, then the special recipe set my recipe.
func AddSecretIngredient(friendsList []string, myList []string) {
	secretRecipe := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretRecipe
}

// A slice of float64 amounts needed for 2 portions.
// The scale of portions I want to cook.
// The argument "quantities" is not modified, return the other quantities array.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	const NeedPortion float64 = 2.0
	scale := float64(portions) / NeedPortion
	resultQuantities := []float64{}
	for _, v := range quantities {
		resultQuantities = append(resultQuantities, v*scale)
	}
	return resultQuantities
}
