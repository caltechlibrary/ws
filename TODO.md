
# Todo ideas

## Someday Maybe

+ Need to start adding the [Web API](https://developer.mozilla.org/en-US/docs/Web/API) that are appropraite to the server side environment
+ add session and auth support
+ create wsedit for remote editing content of specific files over https connections.
    + look at CodeMirror and AceEditor as candidates for web browser editing
    + research best approach to embedding the editor in the go compiled binary
    + review scripted-editor for general apprach to problem
    + decide how to handle TLS key generation seemlessly
        + use existing certs for host machine
        + create one time self signed certs with signatures in browser display along with onetime URL
+ create a nav generator based on file system organization
    + autogenerate sitemap.xml and siteindex.html for improved search engine ingest
+ develop a generator and JS for browser side site specific search
    + create an inverted word list as JSON file
    + create a sitemap JSON file
+ explore cli tools as CMS to produce static websites
    + a markdown processor for generating static site content
+ explore interfacing with Solr or Bleve interface
+ explore adding Lisp support
    + look at LispEx (a Schema?)
    + look at glisp (a lisp/schema dialect?)
    + Find a CL, FranzLisp (e.g. PC-LISP) or XLisp port to Go.
+ explore a wsphp based on PHP parse in https://github.com/stephens2424/php

