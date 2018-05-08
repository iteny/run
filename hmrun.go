package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Unknwon/bra/modules/bindata"
	"github.com/iteny/hmgo/file"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:    "Run",
		Usage:   "make an explosive entrance",
		Version: "1.0.0",
		Action:  runInit,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
func runInit(ctx *cli.Context) error {
	if file.IsExist(".bra.toml") == false {
		var answer string

		// for strings.ToLower(answer) == "yes" || strings.ToLower(answer) == "y" {
		// 	fmt.Println(strings.ToLower(answer) != "y")
		//
		// 	fmt.Scan(&answer)
		// 	if strings.ToLower(answer) == "no" || strings.ToLower(answer) == "n" {
		// 		return nil
		// 	}
		// }
		// fmt.Print("There is a .bra.toml in the work directory, do you want to overwrite?(y/n): ")
		//
		fmt.Println("There is no configuration file in the working directory .bra.toml,Do you need to create a configuration file?(yes/no):")
		fmt.Scan(&answer)
		if strings.ToLower(answer) == "yes" || strings.ToLower(answer) == "y" {
			fmt.Println("create file.")
			return nil
		} else {
			fmt.Println("sorry!你没有配置文件，将无法使用本软件")
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Fail to get work directory: %v", err)
	}

	data, err := bindata.Asset("templates/default.bra.toml")
	if err != nil {
		log.Fatal("Fail to get asset: %v", err)
	}

	appName := filepath.Base(wd)
	if runtime.GOOS == "windows" {
		appName += ".exe"
	}

	data = bytes.Replace(data, []byte("$APP_NAME"), []byte(appName), -1)
	if err := ioutil.WriteFile(".bra.toml", data, os.ModePerm); err != nil {
		log.Fatal("Fail to generate default .bra.toml: %v", err)
	}
	return nil
}
