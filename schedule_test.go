package scheduler

import (
	"fmt"
	"testing"
	"time"
)

func TestDoScheduleJob(t *testing.T) {
	s := NewScheduler(time.Second*2, func() error {
		fmt.Println("hello")
		return nil
	})
	s.DoScheduleJob()
}
