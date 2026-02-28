package period

import (
	"time"
)

// PeriodValid 期间有效
type PeriodValid struct {
	ValidFrom *time.Time `json:"valid_from,omitempty"`
	ValidTo   *time.Time `json:"valid_to,omitempty"`
}

// IsExpired 检查标签是否过期
func (t *PeriodValid) IsExpired() bool {
	now := time.Now()
	if t.ValidFrom != nil && now.Before(*t.ValidFrom) {
		return true
	}
	if t.ValidTo != nil && now.After(*t.ValidTo) {
		return true
	}
	return false
}
