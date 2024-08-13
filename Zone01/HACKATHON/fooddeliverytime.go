package pescine

type food struct {
	burger  int
	chips   int
	nuggets int
}

func FoodDeliveryTime(order string) int {
	plate := food{
		burger:  15,
		chips:   10,
		nuggets: 12,
	}
	if order == "burger" {
		return plate.burger
	} else if order == "chips" {
		return plate.chips
	} else if order == "nuggets" {
		return plate.nuggets
	}
	return 404
}
