package main

import "flag"

var configFile = flag.String("f", "config/config.yml", "set config file which viper will loading")

func main()  {
	flag.Parsed()

	app, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}