//
// ws.go - A simple web server for static files and limit server side JavaScript
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	// Local package
	"github.com/caltechlibrary/ws"
)

// Flag options
var (
	showHelp    bool
	showVersion bool
	showLicense bool
	initialize  bool
	uri         string
	htdocs      string
	sslkey      string
	sslcert     string
	cfg         *ws.Configuration
)

const license = `
%s

Copyright (c) 2016, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`

func usage() {
	fmt.Println(`
 USAGE: ws [OPTIONS] [HTDOCS]

 OVERVIEW

 ws is a utility for prototyping web services and sites. Start
 it with the "init" options will generate a default directory structure
 in your current path along with selfsigned certs if the url to listen
 for uses the https protocol (e.g. ws -url https://localhost:8443 init).

 OPTIONS

`)
	flag.VisitAll(func(f *flag.Flag) {
		if len(f.Name) > 1 {
			switch f.Name {
			case "htdocs":
				fmt.Printf("    -%s, -%s\t%s\n", strings.ToUpper(f.Name[0:1]), f.Name, f.Usage)
			default:
				fmt.Printf("    -%s, -%s\t%s\n", f.Name[0:1], f.Name, f.Usage)
			}
		}
	})

	fmt.Printf(`

 EXAMPLES

 Run a static web server using the content in the current directory
 (assumes the environment variables WS_HTDOCS is not defined).

   ws

 Setup a SSL base site saving the configuration in setup.bash.

   ws -url https://localhost:8443 -init
   . setup.bash
   ws

 Setup a standard project without SSL.

   ws -url http://localhost:8000 -init
   . setup.bash
   ws

 Saving the setup for later...

   ws -url https://localhost:8443 \
      -key /etc/ssl/sites/mysite.key \
      -cert /etc/ssl/sites/mysite.crt \
      -htdocs $HOME/Sites \
	  -init

   . setup.bash
   ws

 Version: %s
`, ws.Version)
	os.Exit(0)
}

func init() {
	cfg = new(ws.Configuration)
	cfg.Getenv()
	uri = cfg.URL.String()
	if uri == "" {
		uri = "http://localhost:8000"
	}
	htdocs = cfg.HTDocs
	sslkey = cfg.SSLKey
	sslcert = cfg.SSLCert

	flag.BoolVar(&showHelp, "h", false, "Display this help message")
	flag.BoolVar(&showHelp, "help", false, "Display this help message")
	flag.BoolVar(&showVersion, "v", false, "Should version info")
	flag.BoolVar(&showVersion, "version", false, "Should version info")
	flag.BoolVar(&showLicense, "l", false, "Show license info")
	flag.BoolVar(&showLicense, "license", false, "Show license info")
	flag.BoolVar(&initialize, "i", false, "Initialize a project")
	flag.BoolVar(&initialize, "init", false, "Initialize a project")
	flag.StringVar(&htdocs, "H", htdocs, "Set the htdocs path")
	flag.StringVar(&htdocs, "htdocs", htdocs, "Set the htdocs path")
	flag.StringVar(&uri, "u", uri, "The protocal and hostname listen for as a URL")
	flag.StringVar(&uri, "url", uri, "The protocal and hostname listen for as a URL")
	flag.StringVar(&sslkey, "k", sslkey, "Set the path for the SSL Key")
	flag.StringVar(&sslkey, "key", sslkey, "Set the path for the SSL Key")
	flag.StringVar(&sslcert, "c", sslcert, "Set the path for the SSL Cert")
	flag.StringVar(&sslcert, "cert", sslcert, "Set the path for the SSL Cert")
}

func logRequest(r *http.Request) {
	log.Printf("Request: %s Path: %s RemoteAddr: %s UserAgent: %s\n", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		next.ServeHTTP(w, r)
	})
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()

	// Process flags and update the environment as needed.
	if showHelp == true {
		usage()
	}
	if showLicense == true {
		fmt.Printf(license, appName)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Printf("ws version %s\n", ws.Version)
		os.Exit(0)
	}
	args := flag.Args()
	if len(args) > 0 {
		htdocs = args[0]
	}
	if uri != "" {
		os.Setenv("WS_URL", uri)
	}
	if htdocs != "" {
		os.Setenv("WS_HTDOCS", htdocs)
	}
	if sslkey != "" {
		os.Setenv("WS_SSL_KEY", sslkey)
	}
	if sslcert != "" {
		os.Setenv("WS_SSL_CERT", sslcert)
	}
	// Merge the environment changes
	cfg.Getenv()

	// Run through initialization process if requested.
	if initialize == true {
		if cfg.HTDocs == "" || cfg.HTDocs == "." {
			cfg.HTDocs = "htdocs"
		}
		if cfg.URL.Scheme == "https" {
			if cfg.SSLKey == "" {
				cfg.SSLKey = "etc/ssl/site.key"
			}
			if cfg.SSLCert == "" {
				cfg.SSLCert = "etc/ssl/site.crt"
			}
		}
		setup, err := cfg.InitializeProject()
		if err != nil {
			log.Fatalf("%s", err)
		}
		// Do a sanity check before generating project
		err = cfg.Validate()
		if err != nil {
			log.Fatalf("Proposed configuration not valid, %s", err)
		}
		ioutil.WriteFile("setup.bash", []byte(setup), 0660)
		log.Println("Wrote setup to setup.bash")
		os.Exit(0)
	}

	// Do a final sanity check before starting up web server
	err := cfg.Validate()
	if err != nil {
		log.Fatalf("Invalid configuration, %s", err)
	}

	log.Printf("HTDocs %s", cfg.HTDocs)
	log.Printf("Listening for %s", cfg.URL.Host)
	http.Handle("/", http.FileServer(http.Dir(cfg.HTDocs)))
	if cfg.URL.Scheme == "https" {
		http.ListenAndServeTLS(cfg.URL.Host, cfg.SSLCert, cfg.SSLKey, logger(http.DefaultServeMux))
	} else {
		http.ListenAndServe(cfg.URL.Host, logger(http.DefaultServeMux))
	}
}
