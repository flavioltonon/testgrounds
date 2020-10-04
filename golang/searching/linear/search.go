package linear

func Search(value int, source []int) bool {
	for _, v := range source {
		if v == value {
			return true
		}
	}

	return false
}
