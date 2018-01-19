package main

import (
	"Matts-Macbook-Pro.local.com/code-advent-2017/challenge"

	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "code-advent-2017"

	app.Action = func(c *cli.Context) {
		var day int
		if c.NArg() > 0 {
			num := c.Args().Get(0)
			i, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("That's not a number!")
				return
			} else {
				day = i
			}
		} else {
			fmt.Println("Requires a challenge number")
			return
		}

		if day == 1 {
			challenge.DayOne()
		} else if day == 2 {
			challenge.DayTwo()
		} else if day == 3 {
			challenge.DayThree()
		} else {
			fmt.Println("Haven't gotten this far!")
			return
		}
	}

	app.Run(os.Args)
}
