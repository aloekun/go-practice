package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	// slice's index exist check
	if 0 <= index && index < len(slice) {
		return slice[index]
	}
	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	// slice's index exist check
	if 0 <= index && index < len(slice) {
		slice[index] = value
	} else {
		// If index does not exist, append tail of slice.
		slice = append(slice, value)
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	// If values are empty, return slice without change.
	if len(values) == 0 {
		return slice
	}
	// Combine slice and values.
	sliceNew := append(values, slice...)
	return sliceNew
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	// If index is out of slice, return slice without change.
	if index < 0 || len(slice) <= index {
		return slice
	}

	slicePre := slice[:index]
	sliceSuffix := slice[index+1:]
	return append(slicePre, sliceSuffix...)
}
