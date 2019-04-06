package main

import (
	"fmt"
	"os"

	"github.com/sethvargo/go-password/password"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	pNumber      int
	pLength      int
	pDigits      int
	pSymbols     int
	pUpperLower  bool
	pRepeatChars bool
)

func main() {
	app := setApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func genPassword(length int, digits int, symbols int, upperlower bool, repeatchars bool) string {
	res, err := password.Generate(length, digits, symbols, upperlower, repeatchars)
	if err != nil {
		log.Fatal(err)
	}
	password := fmt.Sprintf(res)

	return password
}

func setApp() *cli.App {
	app := cli.NewApp()
	app.Name = "genpwd"
	app.Usage = "generate random passwords of specified length"
	app.Flags = getFlags()
	app.Action = func(c *cli.Context) error {
		i := 0
		for i < pNumber {
			fmt.Println(genPassword(pLength, pDigits, pSymbols, pUpperLower, pRepeatChars))
			i++
		}
		return nil
	}

	return app
}

func getFlags() []cli.Flag {
	var flags []cli.Flag

	flags = append(flags, cli.IntFlag{
		Name:        "length, l",
		Value:       16,
		Usage:       "Password length",
		Destination: &pLength,
	},
		cli.IntFlag{
			Name:        "number, n",
			Value:       1,
			Usage:       "Number of passwords to generate",
			Destination: &pNumber,
		},
		cli.IntFlag{
			Name:        "digits, d",
			Value:       4,
			Usage:       "Number of digits in password",
			Destination: &pDigits,
		},
		cli.IntFlag{
			Name:        "symbols, s",
			Value:       2,
			Usage:       "Number of symbols in password",
			Destination: &pSymbols,
		},
		cli.BoolTFlag{
			Name:        "upper, u",
			Usage:       "Use upper and lower case",
			Destination: &pUpperLower,
		},
		cli.BoolTFlag{
			Name:        "repeat, r",
			Usage:       "Repeat characters",
			Destination: &pRepeatChars,
		},
	)

	return flags
}
