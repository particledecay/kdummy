package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	heartrate chan time.Duration
	heart     = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kdummy_heartbeat_total",
		Help: "The total number of heartbeats",
	})
)

// Health holds the health status
type Health struct {
	status int
	msg    string
}

// main metric counter for application
func heartBeat(ch chan time.Duration, rate time.Duration) {
	ticker := time.NewTicker(rate)
	defer ticker.Stop()
	var newRate time.Duration

	for {
		select {
		case newRate = <-ch:
			ticker.Reset(newRate)
			log.WithFields(log.Fields{
				"rate": int(newRate.Seconds()),
			}).Info("Set new heart rate.")
		case <-ticker.C:
			heart.Inc()
			log.WithFields(log.Fields{
				"rate": int(newRate.Seconds()),
			}).Info("This is my heartbeat song and I'm gonna play it.")
		}
	}
}

// simple health check
func getHealth(w http.ResponseWriter, r *http.Request) {
	log.Debug("Received a health check request.")

	health := Health{
		status: 200,
		msg:    "OK",
	}

	w.Header().Set("Content-Type", "application/json")

	healthJSON, err := json.Marshal(health)
	if err != nil {
		log.Error("Could not serialize health status into JSON.")

		resp := make(map[string]string)
		resp["status"] = "500"
		resp["msg"] = "Internal Server Error"
		respJSON, err := json.Marshal(resp)
		if err != nil {
			log.Fatal("Can't serialize anything apparently.")
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(respJSON)
		return
	}

	// we're good
	w.WriteHeader(http.StatusOK)
	w.Write(healthJSON)
	return
}

func setHeartRateSeconds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	if rate, ok := vars["rate"]; ok {
		newRate, err := strconv.Atoi(rate)
		if err != nil {
			log.WithFields(log.Fields{
				"rate": rate,
			}).Error("Failed to set new heart rate.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		go func() {
			heartrate <- time.Second * time.Duration(newRate)
		}()

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("{}"))
		return
	}

	log.Debug("Received an invalid heart rate.")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("{}"))
	return
}

func main() {
	verbose := flag.Bool("verbose", false, "show debug messages")
	listenPort := flag.Int("port", 8080, "listen port for server")
	flag.Parse()

	log.SetFormatter(&log.JSONFormatter{})
	// debug mode
	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	mainMux := mux.NewRouter()
	mainMux.HandleFunc("/heart/{rate}", setHeartRateSeconds)

	internalMux := mux.NewRouter()
	internalMux.Handle("/metrics", promhttp.Handler())
	internalMux.HandleFunc("/healthz", getHealth)

	// set up heartbeat goroutine
	heartrate = make(chan time.Duration, 1)
	go heartBeat(heartrate, time.Second*3)
	heartrate <- time.Second * 3

	go func() {
		log.Debug("Listening for internal requests on 9090")
		http.ListenAndServe(":9090", internalMux)
	}()
	log.Info(fmt.Sprintf("Listening for heartbeat changes on %d", *listenPort))
	http.ListenAndServe(fmt.Sprintf(":%d", *listenPort), mainMux)
}
