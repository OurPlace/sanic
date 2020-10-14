package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/deadlysurgeon/speedtest"
	_ "github.com/lib/pq"
)

const Version = "v0.0.2"

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
	db, err := sql.Open("postgres", "host=192.168.1.26 port=5432 user=speed password=password dbname=perf sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for range time.Tick(5 * time.Minute) {
		Runspeedtest(db)
	}
}
