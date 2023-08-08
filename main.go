package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"home-ai-backend/pkg/handler"
	"home-ai-backend/server/config"
	"home-ai-backend/server/flags"
	"home-ai-backend/server/logging"
	"home-ai-backend/server/mysql"
	"home-ai-backend/server/worker"
)

func main() {
	// setup log styling
	logging.Initialize()

	log.InfoLevelStyle = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("12")).
		Foreground(lipgloss.Color("256"))

	// init flags
	flags.Get()

	// import config xml
	config.Importer()

	mysql.InitializeDatabase()

	worker.InitializeCron()

	handler.HandleRequests()
}
