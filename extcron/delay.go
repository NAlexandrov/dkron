package extcron

import (
	"time"

	"github.com/robfig/cron/v3"
)

// DelayedSchedule represents a delayed recurring duration.
type DelayedSchedule struct {
	Date     time.Time
	Schedule cron.Schedule
}

// After just stores the given time and schedule for this schedule.
func After(date time.Time, schedule cron.Schedule) DelayedSchedule {
	return DelayedSchedule{
		Date:     date,
		Schedule: schedule,
	}
}

// Next conforms to the Schedule interface but this kind of jobs
// doesn't need to be run more than once, so it doesn't return a new date but the existing one.
func (schedule DelayedSchedule) Next(t time.Time) time.Time {
	// If the date set is after the reference time return it.
	// If it's before, return a new date so it runs
	if schedule.Date.After(t) {
		return schedule.Date
	}
	return schedule.Schedule.Next(t)
}
