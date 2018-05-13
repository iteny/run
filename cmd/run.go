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
		InitCmds         [][]string       `toml:"init_cmds"`
		WatchAll         bool             `toml:"watch_all"`
		WatchDirs        []string         `toml:"watch_dirs"`
		WatchExts        []string         `toml:"watch_exts"`
		IgnoreDirs       []string         `toml:"ignore"`
		IgnoreFiles      []string         `toml:"ignore_files"`
		IgnoreRegexps    []*regexp.Regexp `toml:"-"`
		BuildDelay       int              `toml:"build_delay"`
		InterruptTimeout int              `toml:"interrupt_timout"`
		GracefulKill     bool             `toml:"graceful_kill"`
		Cmds             [][]string       `toml:"cmds"`
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
