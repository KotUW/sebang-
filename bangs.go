// First we will do bangs as a json file.
//
// Then use redis robably for user cokkie thing as, we will need to store user info too. (will also test fly.io persistance).
package main

import (
	"encoding/json"
	"errors"
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

func Newbangs() *Bangs {
	cont := getConfig()
	bangs := Bangs{}

	err := json.Unmarshal(cont, &bangs)
	//Caution: Will unmarshal a valid json file that doesn't have any fields.
	if err != nil {
		log.Println("Can't parse json. May be corrpted. ", err)
		return &Bangs{}
	}

	return &bangs
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
		return errors.New("Invalid key")
	}

	_, ok := b.Bang[key]
	if ok {
		return errors.New("Key already exist!. Dupilcation not allowed.")
	}

	b.Bang[key] = url
	json, err := json.Marshal(b)
	if err != nil {
		return err
	}
	config_path := configPath()

	os.WriteFile(config_path, json, 0644)
	return nil
}
