
# ws

## A nimble web server

_ws_ is a prototyping platform for web based services and websites.

### _ws_ has a minimal feature set

+ A simple static file webserver 
    + quick startup
    + activity logged to the console
    + supports http2 out of the box
+ A simple server side JavaScript runner for proxing remote JSON resources
  + if you need more, check out [NodeJS](http://nodejs.org)
+ A project setup option called *init*


## Configuration

You can configure _ws_ with command line options or environment variables.
Try `ws -help` for a list of command line options.

### Environment variables

+ WS_URL the URL to listen for by _ws_
  + default is http://localhost:8000
+ WS_HTDOCS the directory of your static content you need to serve
  + the default is ./htdocs
+ WS_JSDOCS the directory for any server side JavaScript processing
  + the default is ./jsdocs (if not found then server side JavaScript is turned off)
+ WS_SSL_KEY the path the the SSL key file (e.g. etc/ssl/site.key)
  + default is empty, only checked if your WS_URL is starts with https://
+ WS_SSL_CERT the path the the SSL cert file (e.g. etc/ssl/site.crt)
  + default is empty, only checked if your WS_URL is starts with https://

### Command line options

+ -url overrides WS_URL
+ -htdocs overrides WS_HTDOCS
+ -jsdocs overrides WS_JSDOCS
+ -ssl-key overrides WS_SSL_KEY
+ -ssl-pem overrides WS_SSL_PEM
+ -init triggers the initialization process and creates a setup.bash file
+ -h, -help displays the help documentation
+ -v, -version display the version of the command
+ -l, -license displays license information

Running _ws_ without environment variables or command line options is an easy way
to server your current working directory's content out as http://localhost:8000.


## The Server Side JavaScript implementation

The server side JavaScript support in _ws_ is intended to make it easy to
pull in JSON resources from the web for integration and testing with your
static website content. It uses a JavaScript engine called [otto](https://github.com/robertkrimen/otto).

[otto](https://github.com/robertkrimen/otto) is a JavaScript virtual machine
written by Robert Krimen. Each JavaScript file in the *jsdocs* directory tree
becomes a URL end point or route. E.g. *jsdocs/example-1.js* becomes the
route */example-1*. *example-1*. Each of the server side JavaScript files
should contain a closure accepting a "Request" and "Response" object as
parameters.  E.g.

```JavaScript
    /* example-1.js - a simple example of Request and Response objects */
    (function (req, res) {
        var header = req.Header;

        res.setHeader("content-type", "text/html");
        res.setContent(
          "<p>Here is the Header array received by this request</p>" +
          "<pre>" + JSON.stringify(header) + "</pre>");
    }(Request, Response));
```

Assuming server side JavaScript is enabled then this end point would rendered 
with a content type of "text/html". The body should be holding the paragraph and
pre element.

Some additional functions are provided to facilitate server side
JavaScript development--

+ http related
  + WS.httpGet(url, array_of_headers) which performs a HTTP GET
  + WS.httpPost(url, array_of_headers, payload) which performs an HTTP POST
+ os related
  + WS.getEnv(varname) which will read an environment variable

The server side JavaScripts cannot call other JavaScript files or modules. This
was a deliberate design decission as [NodeJS](https://nodejs.org) is a widely 
available server side JS environment with a large community and amply documented
at this point in time.

Need to quickly build out a website from Markdown files or other JSON resources?
Take a look at [mkpage](https://caltechlibrary.github.io/mkpage).


## Installation

_ws_ is available as precompile binaries for Linux, Mac OS X, and Windows 10 on Intel.
Additional binaries are provided for Raspbian on ARM6 adn ARM7.  Follow the [INSTALL.md](install.html) 
instructions to download and install the pre-compiled binaries.

If you have Golang installed then _ws_ can be installed with the *go get* command.

```
    go get github.com/caltechlibrary/ws/...
```

## Compiling from source

Required

+ [Golang](http://golang.org) version 1.7 or better
+ A 3rd Party Go package
  + [Otto](https://github.com/robertkrimen/otto) by Robert Krimen, MIT license
    + a JavaScript interpreter written in Go
    + "go gettable" with `go get github.com/robertkrimen/otto/...`


Here's my basic approach to get things setup. Go 1.7 needs to be already installed.

```
  git clone https://github.com/caltechlibrary/ws
  cd ws
  go get -u github.com/robertkrimen/otto
  go test
  go build
  go build cmds/ws/ws.go
```

If everything compiles fine then I do something like this--

```
  go install cmds/ws/ws.go
```


## LICENSE

copyright (c) 2016 Caltech
See [LICENSE](license.html) for details

