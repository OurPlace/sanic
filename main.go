package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/deadlysurgeon/speedtest"
	_ "github.com/lib/pq"
)

const Version = "v0.0.2"

type config struct {
	DSN  string `json:"dsn"`
	Loop string `json:"loop"`
}

func LoadConfig() (config, error) {
	var conf config

	file, err := os.Open("config.json")
	if err != nil {
		return conf, err
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(b, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func StoreResults(db *sql.DB, result speedtest.Results) error {
	_, err := db.Exec("insert into network.speed (status, packet_loss,upload_bandwidth,upload_bytes,upload_elapsed, download_bandwidth, download_bytes, download_elapsed, ping, jitter, isp, server_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		1, result.PacketLoss, result.Upload.Bandwidth, result.Upload.Bytes, result.Upload.Elapsed, result.Download.Bandwidth, result.Download.Bytes, result.Download.Elapsed, result.Ping.Latency, result.Ping.Jitter, result.ISP, result.Server.ID,
	)
	return err
}

func Runspeedtest(db *sql.DB) {
	fmt.Println("running speed test")
	results, _ := speedtest.NewTest()
	fmt.Println("done running speedtest")

	err := StoreResults(db, results)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Sanic: " + Version)

	conf, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	delay, err := time.ParseDuration(conf.Loop)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Delay is %s\n", delay)

	db, err := sql.Open("postgres", conf.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	Runspeedtest(db)

	for range time.Tick(delay) {
		Runspeedtest(db)
	}
}
