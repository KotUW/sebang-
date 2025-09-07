// First we will do bangs as a json file.
//
// Then use redis robably for user cokkie thing as, we will need to store user info too. (will also test fly.io persistance).
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func configPath() string {
	config_path := os.Getenv("BANG_CONFIG_PATH")
	if config_path == "" {
		config_path = "./public/bangs.json"
	}

	return config_path
}

func getConfig() []byte {
	config_path := configPath()

	data, err := os.ReadFile(config_path)
	if err != nil {
		log.Printf("Can't read the config (%s) file: %s", config_path, err)
		os.Exit(1)
	}

	log.Println("Imported Config from", config_path)
	return data
}

type Bangs struct {
	Default string
	Bang    map[string]string
}

func NewBangs() *Bangs {
	cont := getConfig()
	bangs := &Bangs{
		Bang: make(map[string]string),
	}

	//Caution: Will unmarshal a valid json file that doesn't have any fields.
	if err := json.Unmarshal(cont, bangs); err != nil {
		log.Println("Can't parse json. May be corrpted. ", err)
		bangs.Default = "https://www.google.com/search?q=%s"
	}

	return bangs
}

// Returns either the site releated to query or default site.
func (b *Bangs) Query(query string) string {
	site, ok := b.Bang[query]
	if ok {
		return site
	} else {
		return b.Default
	}
}

func (b *Bangs) Add(key, url string) error {
	// Verify that the kry start with `!`
	if !strings.HasPrefix(key, "!") {
		return errors.New("Invalid key: must start with '!'")
	}

	if _, exists := b.Bang[key]; exists {
		return errors.New("Key already exist!. Dupilcation not allowed.")
	}

	return b.saveConfig()
}

func (b *Bangs) saveConfig() error {
	data, err := json.Marshal(b)
	if err != nil {
		return fmt.Errorf("Failed to marshal new json: %w", err)
	}

	config_path := configPath()

	if err := os.WriteFile(config_path, data, 0644); err != nil {
		return fmt.Errorf("Failed to write %s. Coz: %w", config_path, err)
	}
	return nil
}
