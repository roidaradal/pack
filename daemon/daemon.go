// Package daemon contains daemon types and functions
package daemon

import (
	"time"

	"github.com/zeroibot/pack/dict"
)

type Instance struct {
	start    *dict.SyncMap[string, time.Time]
	last     *dict.SyncMap[string, time.Time]
	duration *dict.SyncMap[string, time.Duration]
}

// NewInstance creates a new Daemon Instance
func NewInstance() *Instance {
	return new(Instance{
		start:    dict.NewSyncMap[string, time.Time](),
		last:     dict.NewSyncMap[string, time.Time](),
		duration: dict.NewSyncMap[string, time.Duration](),
	})
}

/*
Note: DaemonConfig is expected to have this structure:
Use int for Interval and Margin values so we can disable a daemon by setting the value < 0.

type Config struct {
	<Domain> struct {
		<FeatureInterval> int
		...
	}
	...
}
*/

// LoadConfig loads a Daemon Config which is expected to follow the required structure.
// This checks if any Interval values are 0 (invalid).
//func LoadConfig[T any](path string) (*T, error) {
//
//}
