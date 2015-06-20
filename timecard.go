/*
A tool for tracking your time.
*/
package main

import (
	"flag"
	"github.com/jinzhu/now"
	"time"
)

type TimeFlag struct {
	time.Time
}

func (t *TimeFlag) Set(s string) (err error) {
	t.Time, err = now.Parse(s)
	return
}

func enter_time(when time.Time, project string) {
}

func list_time(when time.Time) {
}

func main() {
	mode_enter := false
	mode_list := false
	when := TimeFlag{time.Now()}

	flag.BoolVar(&mode_enter, "enter", false, "Mode is entering time (default if arguments)")
	flag.BoolVar(&mode_list, "list", false, "Mode is displaying time (default if no args)")
	flag.Var(&when, "when", "The time to operate on")
	flag.Parse()

	switch {
	case mode_enter:
		enter_time(when.Time, flag.Arg(0))
	case mode_list:
		list_time(when.Time)
	default:
		if flag.NArg() != 0 {
			enter_time(when.Time, flag.Arg(0))
		} else {
			list_time(when.Time)
		}
	}
}
