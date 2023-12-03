package internal

import (
	"bufio"
	"os"
	"strings"

	"github.com/dwisiswant0/go-stare/pkg/stare"
	"github.com/projectdiscovery/gologger"
)

// Validator to validate options
func Validator(cfg *stare.Config) {
	if isStdin() {
		cfg.URL = bufio.NewScanner(os.Stdin)
	} else if cfg.Target != "" {
		targetURL := cfg.Target
		if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
			// If the target URL doesn't have a protocol, prepend "https://" and "http://"
			targetURL = "https://" + targetURL
			targetURL = "http://" + targetURL
		}

		cfg.URL = bufio.NewScanner(strings.NewReader(targetURL))
	} else {
		gologger.Errorf("No target inputs provided!")
		gologger.Infof("Use -h flag for more info about command.")
		os.Exit(1)
	}
}

func isStdin() bool {
	f, e := os.Stdin.Stat()
	if e != nil {
		return false
	}

	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}
