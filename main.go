package main

import (
	"log"
	"os"
	"runtime"

	"github.com/iteny/run/cmd"
	"github.com/urfave/cli"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	app := cli.NewApp()
	app.Name = "Run"
	app.Usage = "make an explosive entrance"
	app.Version = "1.4.2.0603"
	// app.Flags = cmd.GetFlag

	// app.Action = cmd.RunAction
	app.Commands = []cli.Command{
		cmd.CmdInit,
		cmd.CmdRun,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	/**
	 * @param1 input a flag file name 输入一个标记，它是文件名
	 * @param2 default flag go
	 * @param3 help infomation
	 */
	// username := flag.String("ss", "asdf", "Input your name")
	// flag.Parse()
	// fmt.Println("gg", *username)
	// gopath := os.Getenv("GOPATH")
	// fmt.Println(gopath)
	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal("Fail to get work directory: %v", err)
	// }
	// fmt.Println("wd", wd)
	// appName := filepath.Base(wd)
	// fmt.Println("appName", appName)
	// if runtime.GOOS == "windows" {
	// 	appName += ".exe"
	// }
	// toml, _ := ioutil.ReadFile(gopath + "/src/github.com/iteny/run/template/default.bra.toml")
	// data := bytes.Replace(toml, []byte("$APP_NAME"), []byte(appName), -1)
	// fmt.Println(data)
	// fmt.Println(string(data))
	// fmt.Println(gopath + "/src/github/iteny/run/template/default.bra.toml")
	// buf, _ := ioutil.ReadFile(gopath + "/src/github.com/iteny/run/template/default.bra.toml")
	// fmt.Println("读取文件", buf)
	// ioutil.WriteFile(".bra.toml", data, os.ModePerm)
	// app := cli.NewApp()
	// app.Name = "Run"
	// app.Usage = "make an explosive entrance"
	// app.Version = "1.0.0"
	// app.Action = runInit

	// err := app.Run(os.Args)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// func runInit(ctx *cli.Context) error {
// 	if file.IsExist(".bra.toml") == false {
// 		var answer string

// 		// for strings.ToLower(answer) == "yes" || strings.ToLower(answer) == "y" {
// 		// 	fmt.Println(strings.ToLower(answer) != "y")
// 		//
// 		// 	fmt.Scan(&answer)
// 		// 	if strings.ToLower(answer) == "no" || strings.ToLower(answer) == "n" {
// 		// 		return nil
// 		// 	}
// 		// }
// 		// fmt.Print("There is a .bra.toml in the work directory, do you want to overwrite?(y/n): ")
// 		//
// 		fmt.Println("There is no configuration file in the working directory .bra.toml,Do you need to create a configuration file?(yes/no):")
// 		fmt.Scan(&answer)
// 		if strings.ToLower(answer) == "yes" || strings.ToLower(answer) == "y" {
// 			fmt.Println("create file.")
// 			return nil
// 		} else {
// 			fmt.Println("sorry!你没有配置文件，将无法使用本软件")
// 		}
// 	}

// 	wd, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal("Fail to get work directory: %v", err)
// 	}
// 	fmt.Println(wd)
// 	data, err := bindata.Asset("templates/default.bra.toml")
// 	if err != nil {
// 		log.Fatal("Fail to get asset: %v", err)
// 	}
// 	// b, err := ioutil.ReadFile("./templates/default.bra.toml")
// 	// if err != nil {
// 	// 	fmt.Print(err)
// 	// }
// 	// fmt.Println(b)
// 	// fmt.Println(data)
// 	// Asset("templates/default.bra.toml")

// 	// b := os.Args[1]
// 	// fmt.Println("asdf", b)
// 	appName := filepath.Base(wd)
// 	if runtime.GOOS == "windows" {
// 		appName += ".exe"
// 	}

// 	data = bytes.Replace(data, []byte("$APP_NAME"), []byte(appName), -1)
// 	if err := ioutil.WriteFile(".bra.toml", data, os.ModePerm); err != nil {
// 		log.Fatal("Fail to generate default .bra.toml: %v", err)
// 	}
// 	return nil
// }
