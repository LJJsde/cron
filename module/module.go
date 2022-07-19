package module

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var (
	ErrTimeSet              = errors.New("set time error")
	ErrParamsNotAdapted     = errors.New("the number of params is not adapted")
	ErrNotAFunction         = errors.New("this is not a func")
	ErrPeriodNotSpecified   = errors.New("unspecified job period")
	ErrParameterCannotBeNil = errors.New("nil paramaters cannot be used with reflection")
)

type Task struct {
	err                 error
	RepeatTime          uint
	TimeUnit            uint
	ActiveTime          time.Duration
	LastRun             time.Time
	NextRun             time.Time
	ActiveTimeByWeekday time.Weekday
	LocalTime           *time.Location
	TaskTag             []string
	TaskFunction        map[string]interface{}
	TaskFunctionMain    string
	TaskFunctionParam   map[string]interface{}
}

type TaskResult struct {
	gorm.Model
	Result   string `gorm:"varchar(255)"`
	TaskNote string `gorm:"varchar(255);unique"`
}

const (
	seconds uint = 1
	minutes
	hours
	days
	weeks
)
