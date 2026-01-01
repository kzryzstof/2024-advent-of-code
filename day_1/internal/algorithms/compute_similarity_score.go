package algorithms

import (
	"day_1/internal/abstractions"
	"fmt"
)

func CompareSimilarityScore(
	office *abstractions.Office,
) (abstractions.SimilarityScore, error) {

	if office.List.Len() != office.OtherList.Len() {
		return 0, fmt.Errorf("lists have different lengths: %d vs %d", office.List.Len(), office.OtherList.Len())
	}

	similarityScore := abstractions.SimilarityScore(0)
	otherLocationIndex := uint(0)

	for locationIndex := uint(0); locationIndex < office.List.Len(); locationIndex++ {

		locationId := office.List.Get(locationIndex)

		similarNumbers := uint64(0)

		for ; otherLocationIndex < office.OtherList.Len(); otherLocationIndex++ {

			otherLocationId := office.OtherList.Get(otherLocationIndex)

			if locationId < otherLocationId {
				/*
					Location IDs are not equal, and the location ID of the left list is lower than the right list,
					so stop here and try the next location ID from the left list
				*/
				break
			}

			if locationId > otherLocationId {
				/*
					Location IDs are not equal, and the location ID of the left list is higher than the right list,
					so let's continue browsing the right list
				*/
				continue
			}

			/* Both location IDs are equal! Let's count them */
			similarNumbers++
		}

		similarityScore += abstractions.SimilarityScore(
			similarNumbers * uint64(locationId),
		)
	}

	return similarityScore, nil
}
