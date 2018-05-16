package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"

	"github.com/BurntSushi/toml"
	"github.com/Unknwon/com"
	"github.com/urfave/cli"
)

var (
	WorkDir string
)
var Cfg struct {
	Run struct {
		InitCmds         [][]string       `toml:"init_cmds"`        //初始化执行命令
		WatchAll         bool             `toml:"watch_all"`        //监视所有子目录
		WatchDirs        []string         `toml:"watch_dirs"`       //需要监视的目录
		WatchExts        []string         `toml:"watch_exts"`       //需要监视的后缀
		IgnoreDirs       []string         `toml:"ignore"`           //忽视的目录
		IgnoreFiles      []string         `toml:"ignore_files"`     //忽视的文件
		IgnoreRegexps    []*regexp.Regexp `toml:"-"`                //忽视的正则
		BuildDelay       int              `toml:"build_delay"`      //事件触发最小间隔
		InterruptTimeout int              `toml:"interrupt_timout"` //超时终止程序
		GracefulKill     bool             `toml:"graceful_kill"`    //等待退出然后杀死
		Cmds             [][]string       `toml:"cmds"`             //文件状态改变执行命令
	} `toml:"run"`
	Sync struct {
		ListenAddr string `toml:"listen_addr"`
		RemoteAddr string `toml:"remote_addr"`
	} `toml:"sync"`
}
var CmdRun = cli.Command{
	Name:   "run",
	Usage:  "start monitoring and notifying",
	Action: runRun,
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

func runRun(ctx *cli.Context) error {
	var err error
	WorkDir, err = os.Getwd()
	if err != nil {
		log.Fatal("Fail to get work directory: %v", err)
	}

	confPath := path.Join(WorkDir, ".hm.toml")
	if !com.IsFile(confPath) {
		log.Fatal(".hm.toml not found in work directory")
	} else if _, err = toml.DecodeFile(confPath, &Cfg); err != nil {
		log.Fatal("Fail to decode .hm.toml: %v", err)
	}
	fmt.Println(&Cfg)
	return nil
}
