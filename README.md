# HTTPServerTest

Testing tool for contents of an HTTP response and log file which an HTTP server outputs

## Abstract

HTTPServerTest is a testing tool for an HTTP server. It can check not only an HTTP response but also a log file which an HTTP server outputs.

## Installation

Download a binary file on the [GitHub Releases page](https://github.com/keijiyoshida/httpservertest/releases) and deploy it to anywhere you want on your server.

## Usage

```shell
$ httpservertest -u http://yourdomain.com -t /your/test/cases/file/path -p /your/parameters/file/path
```

## Test Cases File

A test cases file is a YAML file which describes test cases. [sample_files/tests.sample.yml](https://github.com/keijiyoshida/httpservertest/blob/master/sample_files/tests.sample.yml) is a sample of a test cases file. See [the GoDoc page of the httpservertest package](https://godoc.org/github.com/keijiyoshida/httpservertest/httpservertest) for available settings of a test cases file.

## Parameters File

A parameters file is a YAML file which describes parameters which are injected into a test cases file. Parameters on a parameters file can be accessed by writing `{{.Params.Xxxx}}` on a test cases file. [sample_files/params.sample.yml](https://github.com/keijiyoshida/httpservertest/blob/master/sample_files/params.sample.yml) is a sample of a parameters file.

## Cross Compile

```shell
$ ./cross_compile.sh "Version of HTTPServerTest like 0.0.1"
```
