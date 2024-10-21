package main

import (
	"encoding/json"
	"fmt"
	"os"

	parrot "github.com/go-parrot/parrot/pkg/app"
	"github.com/go-parrot/parrot/pkg/config"
	v "github.com/go-parrot/parrot/pkg/version"
	"github.com/spf13/pflag"

	"github.com/go-parrot/parrot-starter/internal/tasks"
)

var (
	cfgDir  = pflag.StringP("config dir", "c", "config", "config path.")
	env     = pflag.StringP("env name", "e", "", "env var name.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()
	if *version {
		ver := v.Get()
		marshaled, err := json.MarshalIndent(&ver, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshaled))
		return
	}

	// init config
	c := config.New(*cfgDir, config.WithEnv(*env))
	var cfg parrot.Config
	if err := c.Load("app", &cfg); err != nil {
		panic(err)
	}
	// set global
	parrot.Conf = &cfg

	// load config
	c = config.New(*cfgDir, config.WithEnv(*env))
	var taskCfg tasks.Config
	if err := c.Load("consumer", &taskCfg); err != nil {
		panic(err)
	}

	// start app
	app, cleanup, err := InitApp(&cfg, &cfg.HTTP, &taskCfg)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}
