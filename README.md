# XlsForXMLHttp service convert XLS files to XML for HTTP Server
[![Build Status](https://travis-ci.org/formeo/XlsForXMLHttp.svg?branch=master)](https://travis-ci.org/formeo/XlsForXMLHttp)
![Go](https://github.com/formeo/XlsForXMLHttp/workflows/Go/badge.svg?branch=master)

## Description

* Service starts HTTP server
* Server scans the source directory, select the target files
* Analyzes files and generates XML
* Client refers to the address on the server HTTP and server issues formed XML

## Usage

  * http://127.0.0.1:8081/payorder/files/sber -- make XML from sber xls
  * http://127.0.0.1:8081/payorder/files/zapsib -- make XML from sber zapsib
  * http://127.0.0.1:8081/payorder/backup -- copy files from source dir to backup dir


## Install
go get github.com/formeo/XlsForXMLHttp

## License
MIT:

## Author
Gordienko Roman
