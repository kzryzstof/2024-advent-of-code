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

	deltaSign := 0

	for levelIndex := 1; levelIndex < len(r.levels); levelIndex++ {

		/* Reads the current level as well as the previous one for comparison */
		previousLevel := r.levels[levelIndex-1]
		level := r.levels[levelIndex]

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

	deltaSign := 0

	delta := int(level) - int(previousLevel)
	unsignedDelta := math.Abs(float64(delta))

	if unsignedDelta < MinDelta || unsignedDelta > MaxDelta {
		return 0, fmt.Errorf("delta out of bounds: %g", unsignedDelta)
	}

	if delta <= 0 {
		deltaSign = -1
	} else {
		deltaSign = 1
	}

	return deltaSign, nil
}
