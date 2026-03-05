package entry

import (
	"time"
)

// PeriodEntry 期间有效
type PeriodEntry struct {
	ValidFrom *time.Time `json:"valid_from,omitempty"`
	ValidTo   *time.Time `json:"valid_to,omitempty"`
}

// IsExpired 检查标签是否过期
func (t *PeriodEntry) IsExpired() bool {
	now := time.Now()
	if t.ValidFrom != nil && now.Before(*t.ValidFrom) {
		return true
	}
	if t.ValidTo != nil && now.After(*t.ValidTo) {
		return true
	}
	return false
}
