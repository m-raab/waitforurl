package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// Port contains tcpPortWait config parameters
type Config struct {
	urlString    string
	searchString string

	timeout int
	period  int

	runtime int
}

func (config *Config) ParseCommandLine() {
	paramUrlString := flag.String("url", "", "URL to wait for")
	paramSearchString := flag.String("search", "", "Search string")
	paramTimeout := flag.Int("timeout", 300, "Timeout for waiting in seconds for port")
	paramPeriod := flag.Int("period", 30, "Timeperiod in seconds for checking port and content")

	flag.Parse()

	if *paramUrlString == "" {
		fmt.Fprintln(os.Stderr, "Parameter 'url' is empty.")
		flag.CommandLine.Usage()
		os.Exit(101)
	}

	if *paramPeriod > *paramTimeout {
		fmt.Fprintln(os.Stderr, "Parameter 'period' must be smaller than 'timeout'")
		flag.CommandLine.Usage()
		os.Exit(102)
	}

	if *paramPeriod <= 0 {
		fmt.Fprintln(os.Stderr, "Parameter 'period' must be bigger than 0")
		flag.CommandLine.Usage()
		os.Exit(103)
	}
	if *paramTimeout <= 0 {
		fmt.Fprintln(os.Stderr, "Parameter 'timeout' must be bigger than 0")
		flag.CommandLine.Usage()
		os.Exit(104)
	}

	config.urlString = *paramUrlString
	config.searchString = *paramSearchString
	config.timeout = *paramTimeout
	config.period = *paramPeriod
}

func main() {
	config := &Config{}
	config.ParseCommandLine()
	config.runtime = 0

	errUrl := config.CheckForContent(config.urlString)

	if errUrl != nil {
		fmt.Fprintf(os.Stderr, "Erro occured: '%s'", errUrl.Error())
		os.Exit(2)
	}

	os.Exit(0)
}

func (conf *Config) CheckForContent(url string) (err error) {
	err = nil

	for conf.runtime < conf.timeout {

		resp, errget := http.Get(conf.urlString)

		if errget == nil {
			defer resp.Body.Close()
			content, _ := ioutil.ReadAll(resp.Body)
			contentString := string(content)

			if len(conf.searchString) == 0 {
				if len(contentString) > 0 {
					return
				} else {
					err = errors.New("Content of response is empty!")
					return
				}
			} else {
				if strings.Contains(contentString, conf.searchString) {
					return
				} else {
					err = errors.Errorf("Content of response '%s' does not contain search string '%s'!", contentString, conf.searchString)
					return
				}
			}
		}

		conf.runtime += conf.period
		time.Sleep(time.Duration(conf.period) * time.Second)
	}
	err = errors.Errorf("Can not get any content from configured URL '%s' in the configured timeout '%d'.", conf.urlString, conf.timeout)
	return
}
