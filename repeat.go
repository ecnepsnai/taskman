package taskman

import "time"

// Repeat describes a pattern for which a trigger should repeatedly run
type Repeat struct {
	RepeatEvery time.Duration
	Until       time.Duration
}
