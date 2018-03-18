package count_slice_inversions

func CountSliceInversions(s []int) ([]int, int) {
	totalSlice, totalInv := SortAndCountInversions(s, 0)
	return totalSlice, totalInv
}

func SortAndCountInversions(s []int, inv int) ([]int, int) {
	if len(s) == 1 {
		return s, inv
	} else {
		s1, s2 := SplitSlice(s)

		leftS, leftInv := SortAndCountInversions(s1, inv)
		rightS, rightInv := SortAndCountInversions(s2, inv)
		mergedS, mergedInv := MergeAndCountInversions(leftS, rightS)

		totalInv := leftInv + rightInv + mergedInv

		return mergedS, totalInv
	}
}

func MergeAndCountInversions(leftS, rightS []int) ([]int, int) {
	var mergedS []int
	mergedInv := 0
	leftI, rightI := 0, 0

	for i := 0; i < len(leftS)+len(rightS); i++ {
		if leftI == len(leftS) {
			mergedS = append(mergedS, rightS[rightI])
			rightI++
		} else if rightI == len(rightS) {
			mergedS = append(mergedS, leftS[leftI])
			leftI++
		} else if (leftI != len(leftS) && rightI != len(rightS)) && leftS[leftI] < rightS[rightI] {
			mergedS = append(mergedS, leftS[leftI])
			leftI++
		} else if (leftI != len(leftS) && rightI != len(rightS)) && leftS[leftI] > rightS[rightI] {
			mergedS = append(mergedS, rightS[rightI])
			for _, v := range leftS {
				rv := rightS[rightI]
				if v > rv {
					mergedInv++
				}
			}
			rightI++
		}
	}
	return mergedS, mergedInv
}

func SplitSlice (s []int) (s1, s2 []int) {
	if len(s)%2 == 0 {
		s1 = s[:len(s)/2]
		s2 = s[len(s)/2:]
	} else {
		s1 = s[:(len(s)+1)/2]
		s2 = s[(len(s)+1)/2:]
	}
	return
}
