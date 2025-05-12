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

type Bangs map[string]string

func config_path() string {
	config_path := os.Getenv("BANG_CONFIG_PATH")
	if config_path == "" {
		config_path = "./bangs.json"
	}

	return config_path
}

func init_bangs() Bangs {
	config_path := config_path()
	bangs := Bangs{}

	cont, err := os.ReadFile(config_path)
	if err != nil {
		log.Printf("Can't find the config (%s) file: %s", config_path, err)
		return Bangs{}
	}

	err = json.Unmarshal(cont, &bangs)
	if err != nil {
		log.Println("Can't parse json. May be corrpted. ", err)
		return Bangs{}
	}

	log.Println("Imported Config from", config_path)

	return bangs
}

func (b Bangs) add(key, url string) error {
	// Verify that the kry start with `!`
	if !strings.HasPrefix(key, "!") {
		return errors.New("Invalid key")
	}

	if b[key] != "" {
		return errors.New("Key already exist!. Dupilcated not allowed.")
	}

	b[key] = url
	json, err := json.Marshal(b)
	if err != nil {
		return err
	}
	config_path := config_path()

	os.WriteFile(config_path, json, 0644)
	return nil
}
