package scheduler

import "time"

type JobFunc func() error
type Scheduler struct {
	Duration time.Duration
	Job      JobFunc
}

func NewScheduler(duration time.Duration, job JobFunc) *Scheduler {
	return &Scheduler{
		Duration: duration,
		Job:      job,
	}
}

func (s *Scheduler) DoScheduleJob() error {
	ticker := time.NewTicker(s.Duration)
	for {
		select {
		case <-ticker.C:
			if err := s.Job(); err != nil {
				return err
			}
		}
	}
}
