package abstractions

import (
	"fmt"
	"math"
)

type Status uint8

const (
	StatusUnknown Status = iota
	StatusSafe
	StatusUnsafe

	MinDelta = 1
	MaxDelta = 3

	BadLevelsTolerance = 1
)

type Report struct {
	Id uint

	levels []Level
}

func NewReport(
	id uint,
	levels []Level,
) *Report {
	return &Report{id, levels}
}

func (r *Report) GetStatus() Status {

	deltaSign := 0

	badLevelsCount := 0
	skipPreviousLevel := false

	for levelIndex := 1; levelIndex < len(r.levels); levelIndex++ {

		/* Reads the current level as well as the previous one for comparison */
		previousLevelIndex := levelIndex - 1

		if skipPreviousLevel {
			/* Skips the previous level was considered bad so we skip it on this step */
			previousLevelIndex -= 1
			skipPreviousLevel = false
		}

		previousLevel := r.levels[previousLevelIndex]
		level := r.levels[levelIndex]

		currentDeltaSign, err := r.getDeltaSign(level, previousLevel)

		if err != nil {

			if r.isBadLevelAcceptable(&badLevelsCount) {
				skipPreviousLevel = true
				continue
			}

			return StatusUnsafe
		}

		if deltaSign == 0 {
			deltaSign = currentDeltaSign
			continue
		}

		if deltaSign != currentDeltaSign {
			
			if r.isBadLevelAcceptable(&badLevelsCount) {
				skipPreviousLevel = true
				continue
			}

			return StatusUnsafe
		}
	}

	return StatusSafe
}

func (r *Report) isBadLevelAcceptable(
	badLevelsCount *int,
) bool {

	*badLevelsCount++

	return *badLevelsCount <= BadLevelsTolerance
}

func (r *Report) getDeltaSign(
	level Level,
	previousLevel Level,
) (int, error) {

	delta := int(level) - int(previousLevel)

	absDelta := math.Abs(float64(delta))

	if absDelta < MinDelta || absDelta > MaxDelta {
		return 0, fmt.Errorf("delta out of bounds: %g", absDelta)
	}

	if delta <= 0 {
		return -1, nil
	}

	return 1, nil
}
