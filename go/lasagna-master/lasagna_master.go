package lasagna

// define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgPrepTime
}

// define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities []float64
	for i := 0; i < len(quantities); i++ {
		scaledQuantities = append(scaledQuantities, quantities[i]*float64(portions)/2)
	}
	return scaledQuantities
}

// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
