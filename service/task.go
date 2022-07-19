package service

import (
	"awesomeProject7/module"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var loc = time.Local

func CreateNewTask(RepeatTime uint) *module.Task {
	return &module.Task{
		RepeatTime:          RepeatTime,
		LocalTime:           loc,
		LastRun:             time.Unix(0, 0),
		NextRun:             time.Unix(0, 0),
		ActiveTimeByWeekday: time.Sunday,
		TaskFunction:        make(map[string]interface{}),
		TaskTag:             []string{},
	}
}

func GetFunctionName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func SetTime(t string) (hour, min, sec int, err error) {
	ts := strings.Split(t, ":")
	if len(ts) < 2 || len(ts) > 3 {
		return 0, 0, 0, module.ErrTimeSet
	}

	if hour, err = strconv.Atoi(ts[0]); err != nil {
		return 0, 0, 0, err
	}
	if min, err = strconv.Atoi(ts[1]); err != nil {
		return 0, 0, 0, err
	}
	if len(ts) == 3 {
		if sec, err = strconv.Atoi(ts[2]); err != nil {
			return 0, 0, 0, err
		}
	}

	if hour < 0 || hour > 23 || min < 0 || min > 59 || sec < 0 || sec > 59 {
		return 0, 0, 0, module.ErrTimeSet
	}

	return hour, min, sec, nil
}
func Create(RepateTime uint) *module.Task {
	t := module.Task{}
	task := CreateNewTask(RepateTime).Loc(t.LocalTime)
	return task
}
