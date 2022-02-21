package main

import (
	"github.com/alecthomas/kong"
	"strconv"
	"time"
)

func currentMonth() string {
	t := time.Now()
	switch t.Month() {
	case 1:
		return "Jan"
	case 2:
		return "Feb"
	case 3:
		return "Mar"
	case 4:
		return "Apr"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "Aug"
	case 9:
		return "Sept"
	case 10:
		return "Oct"
	case 11:
		return "Nov"
	case 12:
		return "Dez"
	default:
		// TODO: consider returning error (error pattern)
		return "Error"
	}
}

var CLI struct {
	Moin struct {
	} `cmd help:"Start tracking time for the day."`

	// TODO: Consider adding extra types for adding break, work
	Add struct {
		Note  string `arg help:"Note to add to the entry."`
		Break bool   `flag help:"Is the added entry of the type break. [default: ${default}]" default:"false"`
		Tags  string `flag help:"A list of tags for the entry, e.g.: tag1,tag2,...,tagN."`
	} `cmd help:"Add a new time tracking entry."`

	// TODO: Consider using json(l) based format instead plain text
	Edit struct {
	} `cmd help:"Edit time tracking file manually."`

	Report struct {
		Type       string `flag enum:"all,work,break" default:"all" help:"Type of report which shall be shown. [choices: ${enum}]"`
		TagsFilter string `flag help:"Flags to filter for, e.g. tag1,tag2,..,tagN."`
		Day        struct {
			DayNumber int `arg help:"Number of the day to be shown. [default: ${day}]" default:"${day}"`
		} `cmd help: "Report the specified day."`
		Week struct {
			WeekNumber int `arg help:"Number of the week to be shown. [default: ${week}]" default:"${week}"`
		} `cmd help: "Report the specified week"`
		Month struct {
			Name string `arg enum:"Jan,Feb,Mar,Apr,May,June,July,Aug,Sept,Oct,Nov,Dec" default:"${month}" help:"Short name of the month to report. [default: ${month}, choices: ${enum}]"`
		} `cmd help: "list the specified month"`
		Period struct {
			Start time.Time `arg format:"2.1.2006" help:"<day>.<month>.<year>"`
			End   time.Time `arg format:"2.1.2006" help:"<day>.<month>.<year> default: today" default:"${today}"`
		} `cmd help: "report for the given time period"`
	} `cmd help:"Show a time report."`
}

func main() {
	var _, weekNumber = time.Now().ISOWeek()
	kong.Parse(&CLI,
		kong.Vars{
			"today": time.Now().Format("2.1.2006"),
			"day":   strconv.FormatInt(int64(time.Now().Day()), 10),
			"week":  strconv.FormatInt(int64(weekNumber), 10),
			"month": currentMonth(),
		})
}
