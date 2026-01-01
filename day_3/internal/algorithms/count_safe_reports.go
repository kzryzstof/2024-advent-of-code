package algorithms

import (
	"day_3/internal/abstractions"
)

func CountSafeReports(
	reports []abstractions.Report,
) uint {

	safeReportsCount := uint(0)

	for _, report := range reports {
		if report.GetStatus() == abstractions.StatusSafe {
			safeReportsCount++
		}
	}

	return safeReportsCount
}
