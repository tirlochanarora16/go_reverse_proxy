package lb

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func CheckConfigFile() string {
	var configFile string

	flag.StringVar(&configFile, "config", "", "YAML file for reverse proxy configuration is missing")

	flag.Parse()

	if strings.TrimSpace(configFile) == "" {
		log.Println("No config file provided. Use -config=path/to/your/file")
		os.Exit(1)
	}

	return configFile
}

func ReadConfigFile(configFile string) {
	_, err := os.Stat(configFile)

	if os.IsNotExist(err) {
		fmt.Println("File not found")
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Some error occured while reading the file")
		os.Exit(1)
	}
}
