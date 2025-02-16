package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, existsUnits := units[unit]
	if !existsUnits {
		return false
	}

	_, existsBill := bill[item]
	if existsBill {
		// If item exists, add item count with the unit count.
		bill[item] += units[unit]
	} else {
		// If item does not exist, set item count with the unit count.
		bill[item] = units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, existsUnits := units[unit]
	if !existsUnits {
		return false
	}

	itemQuantity, existsBill := bill[item]
	if !existsBill {
		return false
	}

	itemNewQuantity := itemQuantity - units[unit]
	if itemNewQuantity < 0 {
		// not enough quantity
		return false
	} else if itemNewQuantity == 0 {
		// success to delete
		delete(bill, item)
		return true
	}

	// update bill's item
	bill[item] = itemNewQuantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemNum, existsBill := bill[item]
	if !existsBill {
		return 0, false
	}
	return itemNum, true
}
