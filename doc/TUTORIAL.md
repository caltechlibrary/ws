
#  Getting started with _ws_ 

At its simplest _ws_ is a static content web server.  It makes it quick to prototype things
that run browser side.  After you have installed _ws_ server static content is as easy as changing
your directory to your document root and then starting _ws_. 


## Example 1

### Simple usage

You have create a directory called /Sites where you plan to develop your website.  To test
your site with _ws_ you need to--

1. Change to the /Sites directory
2. Start _ws_


```shell
    cd /Sites
    ws .
```

This should yield output like

```shell

                   TLS: false
                  Cert: 
                   Key: 
               Docroot: /Sites
                  Host: localhost
                  Port: 8000
                Run as: johndoe

         2014/07/15 17:52:27 Starting http://localhost:8000
```

You can now point your browser a [http://localhost:8000](http://localhost:8000) and see the contents
of the /Sites directory.

You can press ctrl-C (while holding the key marked "Ctr" or "Ctrl" press the "c" key).  The websere should now stop.


## Example 2

### Organizing and doing more

More typically if you are prototyping a website you will organize your code into different folders
based on your build process or tool set.  _wsinit_ can help here.  The *wsinit* will configure
a folder with a simple structure for further develop.  It also will setup up things for more complex
usage of _ws -init_ like accessing Otto Engine or running under SSL.

Here are the normal four steps you take to set things up. We will do the first two, stop, look around
then proceed to steps 3 and 4 to test it out.

1. Change to the directory that will hold your project (e.g. /Sites)
2. Run _ws -init_. For not accept the defaults by pressing enter you "y" and enter when prompted.
3. Source your _etc/config.sh_ file
4. Start _ws_ webserver and test with your browser.

Steps one and two.

```shell
    cd /Sites
    ws -init
```

Take a look at the directories and files created. By default your static content is configured to run
from the _static_ directory. You will find a new _index.html_ created there for you to modify. 

You will also find a directory called _etc_.  This is where you will find your configuration file
_config.sh_ as well as another sub-directory, _etc/ssl_, holding your SSL certificate and key files.

_ws_ draws configuration from either command line options or your shells environment variables.
If you source _etc/config.sh_ you will see anumber of environment vararibles set so _ws_ know
where to find your static content, dynamic route handlers as well as SSL certificates as needed.
Nows a good time to 

You should seem startup information simular to example 1. This time though the index.html file
delivered to your browser was be the in side this *static* sub-directory.

Now you are ready for steps three and four.

```shell
    . etc/config.sh
    ws
```

