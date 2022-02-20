package main

import "github.com/alecthomas/kong"

var CLI struct {

	// TODO: Consider adding customization, or at least alternatives  (hello, buongiorno, buenos dias, ...) or suggest alias
	Moin struct {
	} `cmd help:"Start your day by saying moin ;D."`

	// TODO: Consider adding extra types for adding break, work
	Add struct {
	} `cmd help:"Add a period of time to the sheet (work, break, ...)."`

	// TODO: Consider using json(l) based format instead plain text
	Edit struct {
	} `cmd help:"Edit the timesheet manually."`

	// TODO: Consider adding various "filters"
	// Type:
	// - break
	// - work
	// Period(s):
	// - day
	// - week
	// - month
	// - period <start> <end>
	Report struct {
	} `cmd help:"Show a time report."`
}

func main() {
	kong.Parse(&CLI)
}
