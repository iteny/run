package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/iteny/hmgo/logm"
	"github.com/urfave/cli"
)

var CmdInit = cli.Command{
	Name:   "init",
	Usage:  "initialize config template file",
	Action: runInit,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "file,f",
			Value: "string",
			Usage: "Create a file byte slice!",
		},
		cli.StringFlag{
			Name:  "config,c",
			Value: "go",
			Usage: "Generate configuration file,the default is go language!",
		},
	},
}

func runInit(c *cli.Context) error {
	// name := ""
	// /**
	//  * @English 返回命令行参数数量
	//  * @Chinese returns the number of the command line arguments
	//  */
	// if c.NArg() > 0 {
	// 	*
	// 	 * @English 返回第N个参数，从0开始，如果没有参数返回空字符串
	// 	 * @Chinese Get returns the nth argument, or else a blank string

	// 	name = c.Args().Get(0)
	// }
	file := c.String("file")
	config := c.String("config")
	if file != "string" {
		gopath := os.Getenv("GOPATH")

		toml, err := ioutil.ReadFile(gopath + "/src/github.com/iteny/run/template/" + file + ".hm.toml")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%#v", toml)
		// fmt.Println(goDefault)
		// // fmt.Println(gopath)
	}
	if config == "go" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal("Fail to get work directory: %v", err)
		}
		appName := filepath.Base(wd)
		if runtime.GOOS == "windows" {
			appName += ".exe"
		}
		data := bytes.Replace(goDefault, []byte("$APP_NAME"), []byte(appName), -1)
		ioutil.WriteFile(".hm.toml", data, os.ModePerm)

	} else if config == "color" {
		fmt.Println(logm.Black("fmt.Println(logm.Black(\"前景黑色\"))"))
		fmt.Println(logm.Red("fmt.Println(logm.Black(\"前景红色\"))"))
		fmt.Println(logm.Green("fmt.Println(logm.Black(\"前景绿色\"))"))
		fmt.Println(logm.Yellow("fmt.Println(logm.Black(\"前景黄色\"))"))
		fmt.Println(logm.Blue("fmt.Println(logm.Black(\"前景蓝色\"))"))
		fmt.Println(logm.Magenta("fmt.Println(logm.Black(\"前景紫红色\"))"))
		fmt.Println(logm.Cyan("fmt.Println(logm.Black(\"前景青蓝色\"))"))
		fmt.Println(logm.White("fmt.Println(logm.Black(\"前景白色\"))"))
		fmt.Println(logm.BlackBg("fmt.Println(logm.Black(\"背景黑色\"))"))
		fmt.Println(logm.RedBg("fmt.Println(logm.Black(\"背景红色\"))"))
		fmt.Println(logm.GreenBg("fmt.Println(logm.Black(\"背景绿色\"))"))
		fmt.Println(logm.YellowBg("fmt.Println(logm.Black(\"背景黄色\"))"))
		fmt.Println(logm.BlueBg("fmt.Println(logm.Black(\"背景蓝色\"))"))
		fmt.Println(logm.MagentaBg("fmt.Println(logm.Black(\"背景紫红色\"))"))
		fmt.Println(logm.CyanBg("fmt.Println(logm.Black(\"背景青蓝色\"))"))
		fmt.Println(logm.WhiteBg("fmt.Println(logm.Black(\"背景白色\"))"))
	} else {
		return cli.NewExitError("it is not in the soup", 86)
	}
	return nil
}

// func GetFlag() *[]cli.Flag {
// 	meFlag := []cli.Flag{
// 		cli.StringFlag{
// 			Name:  "lang",
// 			Value: "english",
// 			Usage: "language for the greeting",
// 		},
// 	}
// 	return &meFlag
// }
