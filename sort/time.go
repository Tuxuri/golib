package sort

import (
	"sort"
	"time"
)

// TimeSlice implements sort.Interface, sorting in increasing order.
type TimeSlice []time.Time

func (p TimeSlice) Len() int           { return len(p) }
func (p TimeSlice) Less(i, j int) bool { return p[i].Unix() < p[j].Unix() }
func (p TimeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p TimeSlice) Sort() { sort.Sort(p) }

// Reverse is a convenience method.
func (p TimeSlice) Reverse() { sort.Sort(sort.Reverse(p)) }
