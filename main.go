package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudfoundry/go-cfclient/v3/client"
	"github.com/cloudfoundry/go-cfclient/v3/config"
)

func main() {
	// Load CF credentials from environment variables (or replace with hardcoded values for testing)
	// cfConfig := &client.Config{
	// 	ApiAddress:        os.Getenv("CF_API"),      // e.g. "https://api.bosh-lite.com"
	// 	Username:          os.Getenv("CF_USERNAME"), // e.g. "admin"
	// 	Password:          os.Getenv("CF_PASSWORD"), // e.g. "your-password"
	// 	SkipSslValidation: true,                     // adjust as needed
	// }

	// Create CF client
	cfg, _ := config.NewFromCFHome()
	cf, _ := client.New(cfg)

	apps, err := cf.Applications.ListAll(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error fetching apps: %v", err)
	}

	fmt.Println("Apps:")
	for _, app := range apps {
		fmt.Printf("- %s (%s)\n", app.Name, app.GUID)
	}
}
