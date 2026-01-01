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

	actualStatus := r.getStatus(r.levels)

	if actualStatus != StatusUnsafe {
		return actualStatus
	}

	for removedIndex := range r.levels {

		alteredLevels := make([]Level, 0)

		for levelIndex, level := range r.levels {
			if levelIndex == removedIndex {
				continue
			}
			alteredLevels = append(alteredLevels, level)
		}

		alteredStatus := r.getStatus(alteredLevels)

		if alteredStatus != StatusUnsafe {
			return alteredStatus
		}
	}

	return StatusUnsafe
}

func (r *Report) getStatus(
	levels []Level,
) Status {

	deltaSign := 0

	for levelIndex := 1; levelIndex < len(levels); levelIndex++ {

		/* Reads the current level as well as the previous one for comparison */
		previousLevelIndex := levelIndex - 1

		previousLevel := levels[previousLevelIndex]
		level := levels[levelIndex]

		currentDeltaSign, err := r.getDeltaSign(level, previousLevel)

		if err != nil {
			return StatusUnsafe
		}

		if deltaSign == 0 {
			deltaSign = currentDeltaSign
			continue
		}

		if deltaSign != currentDeltaSign {
			return StatusUnsafe
		}
	}

	return StatusSafe
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
