package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	jexia "github.com/baileyjm02/jexia-sdk-go"
	linuxproc "github.com/c9s/goprocinfo/linux"
)

func setupJexia() *jexia.Client {
	client := jexia.NewClient(
		os.Getenv("PROJECT_ID"),
		os.Getenv("PROJECT_ZONE"),
	)
	client.UseAPKToken(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	client.AutoRefreshToken()
	return client
}

type system struct {
	ID        string            `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Hostname  string            `json:"hostname"`
	Platform  string            `json:"platform"`
	CPU       linuxproc.CPUStat `json:"cpu"`
	Processes uint64            `json:"processes"`
	Memory    linuxproc.MemInfo `json:"memory"`
	Uptime    linuxproc.Uptime  `json:"uptime"`
	Load      linuxproc.LoadAvg `json:"load"`
}

func main() {
	client := setupJexia()
	healthDataset := client.GetDataset("health")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		var data *[]system
		err := healthDataset.Select(&data)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe(os.Getenv("HOST"), nil)
}
