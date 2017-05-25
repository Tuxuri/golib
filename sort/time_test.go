package sort

import (
	"sort"
	"testing"
	"time"
)

var (
	now   = time.Now()
	times = []time.Time{
		now, now.Add(time.Second), now.Add(time.Minute), now.Truncate(time.Hour),
		now.Truncate(time.Millisecond), now.Add(time.Hour), now.Truncate(time.Nanosecond),
	}
)

func TestTimeSlice(t *testing.T) {
	data := times
	a := TimeSlice(data[0:])
	sort.Sort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", times)
		t.Errorf("   got %v", data)
	}
}
