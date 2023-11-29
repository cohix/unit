package main

import (
	"fmt"
	"log"

	"github.com/nginx/unit/tools/unitctl/pkg/config"
)

func main() {
	c := config.NewClient("/opt/homebrew/var/run/unit/control.sock")

	config := config.Config{
		Listeners: map[string]config.Listener{
			"*:8080": {
				Pass: "applications/subounit",
			},
		},
		Applications: map[string]config.Application{
			"subounit": {
				Type:       "external",
				WorkDir:    "/Users/c.hicks/Workspaces/suborbital/subounit/www",
				Executable: "app",
			},
		},
	}

	applyResp, err := c.ApplyConfig(config)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to c.ApplyConfig: %w", err))
	}

	fmt.Println(applyResp.Success + "\n")

	newC, err := c.GetConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to c.GetConfig: %w", err))
	}

	fmt.Println(newC.Summary())
}
