package main

func getMonthlyPrice(tier string) int {
	// ?
	penny := 1
	dollar := 100 * penny
	if tier == "basic" {
		return 100 * dollar
	} else if tier == "premium" {
		return 150 * dollar
	} else if tier == "enterprise" {
		return 500 * dollar
	} else {
		return 0
	}

}



