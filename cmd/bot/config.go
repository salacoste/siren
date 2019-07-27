package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type config struct {
	Website           string `json:"website"`             // one of the following: bongacams, stripchat, chaturbate
	ListenPath        string `json:"listen_path"`         // the path excluding domain to listen to, the good choice is /your-telegram-bot-token
	ListenAddress     string `json:"listen_address"`      // the address to listen to
	ListenDomain      string `json:"listen_domain"`
	BotToken          string `json:"bot_token"`           // your telegram bot token
	PeriodSeconds     int    `json:"period_seconds"`      // the period of querying models statuses
	MaxModels         int    `json:"max_models"`          // maximum models per user
	TimeoutSeconds    int    `json:"timeout_seconds"`     // HTTP timeout
	AdminID           int64  `json:"admin_id"`            // your telegram ID
	DBPath            string `json:"db_path"`             // path to database
	Certificate       string `json:"certificate"`         // your certificate, omit if under a proxy
	Key               string `json:"key"`                 // your key, omit if under a proxy
	NotFoundThreshold int    `json:"not_found_threshold"` // remove a model after a failure to find her this number of times
	BlockThreshold    int    `json:"block_threshold"`     // do not send a message to the user if we fail to do it due to blocking this number of times
	Translation       string `json:"translation"`         // translation strings
	Debug             bool   `json:"debug"`               // debug mode
}

func readConfig(path string) *config {
	file, err := os.Open(filepath.Clean(path))
	checkErr(err)
	defer func() { checkErr(file.Close()) }()
	return parseConfig(file)
}

func parseConfig(r io.Reader) *config {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	cfg := &config{}
	err := decoder.Decode(cfg)
	checkErr(err)
	if !checkConfig(cfg) {
		panic("config error")
	}
	return cfg
}

func checkConfig(cfg *config) bool {
	return true &&
		cfg.ListenPath != "" &&
		cfg.BotToken != "" &&
		cfg.ListenDomain != "" &&
		cfg.PeriodSeconds != 0 &&
		cfg.MaxModels != 0 &&
		cfg.TimeoutSeconds != 0 &&
		cfg.AdminID != 0 &&
		cfg.DBPath != "" &&
		cfg.ListenAddress != "" &&
		cfg.NotFoundThreshold != 0 &&
		cfg.BlockThreshold != 0 &&
		cfg.Website != "" &&
		cfg.Translation != ""
}
