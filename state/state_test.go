package state

import "testing"

func TestWeek(t *testing.T) {
	ctx := NewDayContext()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}

	for i := range 8 {
		todayAndNext()
		t.Log(i)
	}
}
