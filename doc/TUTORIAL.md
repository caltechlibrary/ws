
#  Getting started with _ws_ 

At its simplest _ws_ is a static content web server.  It makes it quick to prototype things
that run browser side.  After you have installed _ws_ server static content is as easy as changing
your directory to your document root and then starting _ws_. 


## Example 

### Simple usage

You have create a directory called /Sites where you plan to develop your website.  To test
your site with _ws_ you need to--

1. Change to the /Sites directory
2. Start _ws_


```shell
    cd /Sites
    ws
```

This should yield output like

```shell
    2016/12/12 11:54:22 DocRoot .
    2016/12/12 11:54:22 Listening for http://localhost:8000
```

You can now point your browser a [http://localhost:8000](http://localhost:8000) and see the contents
of the /Sites directory.

You can press ctrl-C (while holding the key marked "Ctr" or "Ctrl" press the "c" key).  The websere should now stop.


