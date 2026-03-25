package arrays

import "slices"

func CompareProductLists(oldList, newList []string) (added, removed []string) {
	added = make([]string, 0)
	removed = make([]string, 0)

	//added, _ = lo.Difference(newList, oldList)
	//removed, _ = lo.Difference(oldList, newList)

	for _, newV := range newList {

		isConstainsNewInOld := slices.Contains(oldList, newV)

		if isConstainsNewInOld == false {
			added = append(added, newV)
		}
	}

	for _, oldV := range oldList {

		isConstainsOldInNew := slices.Contains(newList, oldV)

		if isConstainsOldInNew == false {
			removed = append(removed, oldV)
		}
	}

	return added, removed
}
