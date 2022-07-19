package module

import (
	"awesomeProject7/service"
	"reflect"
	"time"
)

func (t *Task) Err() error {
	return t.err
}

func (t *Task) RunNowIfShould() bool {
	return time.Now().Unix() >= t.NextRun.Unix()
}

func (t *Task) Run(task interface{}, params ...interface{}) error {
	if t.err != nil {
		return t.err
	}
	typ := reflect.TypeOf(task)
	if typ.Kind() != reflect.Func {
		return ErrNotAFunction
	}
	fname := service.GetFunctionName(task)
	t.TaskFunction[fname] = task
	t.TaskFunctionParam[fname] = params
	t.TaskFunctionMain = fname
	return nil
}

func (t *Task) At(timing string) *Task {
	hour, min, sec, err := service.SetTime(timing)
	if err != nil {
		t.err = ErrTimeSet
		return t
	}
	t.ActiveTime = time.Duration(hour)*time.Hour + time.Duration(min)*time.Minute + time.Duration(sec)*time.Second
	return t
}

func (t *Task) Loc(loc *time.Location) *Task {
	t.LocalTime = loc
	return t
}

func (t *Task) TagSet(tag string) {
	t.TaskTag = append(t.TaskTag, tag)
}

func (t *Task) SetTime(time uint) *Task {
	t.TimeUnit = time
	return t
}

func (t *Task) Seconds(repTime uint) *Task {
	t.RepeatTime = repTime
	return t.SetTime(seconds)
}

func (t *Task) Minutes(repTime uint) *Task {
	t.RepeatTime = repTime
	return t.SetTime(minutes)
}

func (t *Task) Hours(repTime uint) *Task {
	t.RepeatTime = repTime
	return t.SetTime(hours)
}

func (t *Task) Days(repTime uint) *Task {
	t.RepeatTime = repTime
	return t.SetTime(days)
}

func (t *Task) Weeks(repTime uint) *Task {
	t.RepeatTime = repTime
	return t.SetTime(weeks)
}

func (t *Task) Weekday(repTime uint, StartDay time.Weekday) *Task {
	t.RepeatTime = repTime
	t.ActiveTimeByWeekday = StartDay
	return t.Weeks(1)
}

func (t *Task) Monday(repTime uint) *Task {
	return t.Weekday(repTime, time.Monday)
}

func (t *Task) Tuesday(repTime uint) *Task {
	return t.Weekday(repTime, time.Tuesday)
}

func (t *Task) Wednesday(repTime uint) *Task {
	return t.Weekday(repTime, time.Wednesday)
}

func (t *Task) Thursday(repTime uint) *Task {
	return t.Weekday(repTime, time.Thursday)
}

func (t *Task) Friday(repTime uint) *Task {
	return t.Weekday(repTime, time.Friday)
}

func (t *Task) Satday(repTime uint) *Task {
	return t.Weekday(repTime, time.Saturday)
}

func (t *Task) Sunday(repTime uint) *Task {
	return t.Weekday(repTime, time.Sunday)
}
