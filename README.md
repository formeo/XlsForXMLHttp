# XlsForXMLHttp service convert XLS files to XML for HTTP Server

## Description

* Service starts HTTP server
* Server scans the source directory, select the target files
* Analyzes files and generates XML
* Client refers to the address on the server HTTP and server issues formed XML

## Usage

  * http://127.0.0.1:8081/payorder/files -- make XML
  * http://127.0.0.1:8081/payorder/backup -- copy files from source dir to backup dir
## Install
go get github.com/formeo/XlsForXMLHttp

## License
MIT:

## Author
Gordienko Roman
