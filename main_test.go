package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func Test(t *testing.T) {
	testURL = "http://" + getEnvValue("HOST") + ":" + getEnvValue("PORT")
}

func TestServerPing(t *testing.T) {
	res, err := http.Get(testURL + "/ping")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	b := string(body)
	if !strings.Contains(b, "Voting Service") {
		t.Fatal()
	}
}

func TestLoadPing(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    testURL + "/ping",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestServerCreate(t *testing.T) {
	data := url.Values{
		"title":   {"RustVSGolang"},
		"options": {"Golang,Rust"},
	}
	res, err := http.PostForm(testURL+"/poll", data)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadCreate(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    testURL + "/poll?title=RustVSGolang&options=Golang,Rust",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestServerGet(t *testing.T) {
	res, err := http.Get(testURL + "/poll?title=RustVSGolang&options=RustVSGolang")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadGet(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    testURL + "/poll?title=RustVSGolang&options=Golang,Rust",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
