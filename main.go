package main

import (
	"github.com/alecthomas/kong"
	"github.com/nicoretti/flextime/cli"
	"strconv"
	"time"
)

func main() {
	var command_line cli.Cli
	var now = time.Now()
	var _, weekNumber = time.Now().ISOWeek()
	var ctx = kong.Parse(&command_line,
		kong.Vars{
			"today": time.Now().Format("2.1.2006"),
			"day":   strconv.FormatInt(int64(time.Now().Day()), 10),
			"week":  strconv.FormatInt(int64(weekNumber), 10),
			"month": cli.CurrentMonth(&now),
		})
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
