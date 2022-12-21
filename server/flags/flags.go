package flags

import (
	"flag"
	"shome-backend/server/config"
)

func GetFlags() {
	flag.BoolVar(&config.SETUP, "s", false, "If set to true setup mode will be activated.")
	flag.BoolVar(&config.DEBUG, "d", false, "If set to true debug mode will be activated.")

	flag.StringVar(&config.BROKER_IP, "brokerIP", "localhost", "Sets broker IP")
	flag.StringVar(&config.BROKER_PORT, "brokerPort", "1883", "Sets broker Port")

	flag.Parse()
}

func GetSetupFlag() bool {
	return config.SETUP
}

func GetDebugFlag() bool {
	return config.DEBUG
}
