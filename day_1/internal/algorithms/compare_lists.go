package algorithms

import (
	"day_1/internal/abstractions"
	"fmt"
	"math"
)

func CompareLists(
	office *abstractions.Office,
) (abstractions.Distance, error) {

	if office.List.Len() != office.OtherList.Len() {
		return 0, fmt.Errorf("lists have different lengths: %d vs %d", office.List.Len(), office.OtherList.Len())
	}

	totalDistance := abstractions.Distance(0)

	for locationIndex := uint(0); locationIndex < office.List.Len(); locationIndex++ {
		totalDistance += abstractions.Distance(math.Abs(float64(office.List.Get(locationIndex)) - float64(office.OtherList.Get(locationIndex))))
	}

	return totalDistance, nil
}
