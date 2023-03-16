package main

import (
	"fmt"
	"log"
	"os"

	"github.com/liuchong/lprocinfo/src"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "lpi",
		Usage:     "linux proc info",
		UsageText: "get | set ",
		Commands: []*cli.Command{
			GetFunc(),
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

//GetFunc - List information
func GetFunc() *cli.Command {
	return &cli.Command{
		Name:      "get",
		Usage:     "display linux system infomation",
		UsageText: "cpu | memory | stack",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "cpu",
				Aliases: []string{"c"},
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("cpu") == "load" {
				nl := src.GetLoad("/proc/loadavg")
				fmt.Printf("1分钟负载：%s\n5分钟负载：%s\n15分钟负载：%s\n运行进程数/总进成数：%s\n运行进程ID：%s\n", nl.Fifteen, nl.Fivemin, nl.Fifteen, nl.Perprocess, nl.Runprocess)
			}
			return nil
		},
	}
}
