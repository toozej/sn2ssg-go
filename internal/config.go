package sn2ssg

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getEnvVars() {
	if _, err := os.Stat(".env"); err == nil {
		// Initialize Viper from .env file
		viper.SetConfigFile(".env") // Specify the name of your .env file

		// Read the .env file
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading .env file: %s\n", err)
		}
	}

	// Enable reading environment variables
	viper.AutomaticEnv()

	// get HTTP Basic Auth username and password from Viper
	username = viper.GetString("SN_USERNAME")
	password = viper.GetString("SN_PASSWORD")
	if username == "" {
		log.Fatal("Simplenote username must be provided")
	}

	if password == "" {
		log.Fatal("Simplenote password must be provided")
	}

	// Simplenote config items
	tagToDownload = viper.GetString("TAG_TO_DOWNLOAD")
	continuousNoteTag = viper.GetString("CONTINUOUS_NOTE_TAG")

	// SSG config items
	ssgType = viper.GetString("SSG_TYPE")
	author = viper.GetString("AUTHOR")
	inputDir = viper.GetString("INPUT_DIR")
	outputDir = viper.GetString("OUTPUT_DIR")
	pollingCycle = viper.GetInt("POLLING_CYCLE")

	// notification config items
	gotifyURL = viper.GetString("GOTIFY_URL")
	gotifyToken = viper.GetString("GOTIFY_TOKEN")

	// general config items
	debug = viper.GetBool("DEBUG")
}
