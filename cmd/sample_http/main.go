package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/peterbourgon/ff"
	"github.com/peterbourgon/ff/fftoml"

	"github.com/bagadatp/sample_server/pkg/data"
)

var version = "v0.0.1-dev"

const (
	cfgUsage = `
Configuration is parsed from: command line flags (highest priority),
configuration file, and environment variables (lowest priority).

Flag names are parsed from config file top-level fields. Values are passed to
flags unprocessed, as if they were strings. List values will be processed as
repeated flag with corresponding list elements as values.

Flag names are prefixed with SAMPLE_HTTP_, capitalized, and separator characters
are converted to underscores to match with env vars.`
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	start_time := time.Now().Unix()
	fmt.Printf("%d path %s\n", start_time, r.URL.Path)
	time.Sleep(10 * time.Second)
	end_time := time.Now().Unix()
	fmt.Fprintf(w, "hello... <%v, %v>", start_time, end_time)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	resp:= data.GetBasicData(r.URL.Query())
	fmt.Fprintf(w, resp)
}

func main() {

	l := log.New(os.Stdout, fmt.Sprintf("sample_http/%d ", os.Getpid()), 0)
	//// Init
	var (
		cfg     = flag.NewFlagSet("", flag.ExitOnError)
		cfgPath = cfg.String("config", "", "path to config `file`")
		httpAddr     = cfg.String("http", ":http", "listen on this `addr` for HTTP traffic")
		//httpsAddr    = cfg.String("https", ":https", "listen on this `addr` for HTTPS traffic")
	)
	cfg.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Sample HTTP Server %s\n", version)
		cfg.PrintDefaults()
		_, _ = fmt.Fprintln(os.Stderr, cfgUsage)
	}
	cfgOpts := []ff.Option{
		ff.WithConfigFileFlag("config"),
		ff.WithEnvVarPrefix("SAMPLE_HTTP"),
	}

	// Pre-parse config file location to disable config file parsing if none provided.
	err := cfg.Parse(os.Args[1:])
	if err != nil {
		l.Fatalf("error parsing command arguments: %s", err)
	}
	if *cfgPath != "" {
		cfgOpts = append(cfgOpts, ff.WithConfigFileParser(fftoml.Parser))
	}
	// Process all config inputs.
	if err := ff.Parse(cfg, os.Args[1:], cfgOpts...); err != nil {
		l.Fatal("error parsing config: ", err)
	}


	//// Main
	http.HandleFunc("/healthcheck", Healthcheck)
	http.HandleFunc("/get", GetData)
	s := &http.Server{
		Addr:         *httpAddr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	l.Print("serving HTTP on ", s.Addr)
	log.Fatal(s.ListenAndServe())
}
