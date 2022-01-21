package logical

func exclusiveOr(a,b bool) bool{
	return ( a || b) && !(a && b)
}

func OnlyOneOf(list... bool) bool{
	var count = 0
	for _,b := range list {
		if b {
			count++
			if count > 1 {
				return false
			}
		}
	}
	if count == 0 {
		return false
	} else {
		return true
	}
}

func allOf(list... bool) bool{
	for _,b := range list {
		if !b {
			return false
		}
	}
	return true
}

