package flags

import (
	"flag"
	"home-ai-backend/server/config"
	"home-ai-backend/server/logging"
)

func Get() {
	config.ENVIRONMENT = flag.String("ENVIRONMENT", "prod", "Sets the environment (e.g. dev, stage, prod)")
	config.IS_DEBUG = flag.Bool("d", false, "Activates debug mode")

	flag.Parse()

	logging.PrintListDone("Initialize Flags")
}
