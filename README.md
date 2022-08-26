# proto

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/proto.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/proto/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/proto/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/proto)](https://goreportcard.com/report/github.com/gowizzard/proto) [![Go Reference](https://pkg.go.dev/badge/github.com/gowizzard/proto.svg)](https://pkg.go.dev/github.com/gowizzard/proto) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/proto)](https://github.com/gowizzard/proto/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/proto)](https://github.com/gowizzard/proto/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/proto)](https://github.com/gowizzard/proto/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/proto)](https://github.com/gowizzard/proto/blob/master/LICENSE)

The name proto comes from the German word protokollieren `/protokolliéren/`. Accordingly, the whole thing has nothing to do with creating prototypes, but only logging. Since of course also with pleasure for prototypes.

## Install

```console
go get github.com/gowizzard/proto
```

## How to use?

The library should help you to make logging easier for the command line as well as file based. For example, you can add the file information to your error message, so that we can enter the exact file and line of the error. You can also add more attributes via a map.

The functionality is currently being expanded, so check back from time to time, or create an [issue](https://github.com/gowizzard/proto/issues) if desired.

Here you can find an example how to use the library:

```go
c := proto.Config{
    FileInformation: true,
    CommandLine:     false,
    File:            true,
    Path:            "/var/log/proto",
} 

err := c.Log("error", "this is the error message", map[string]any{})
if err != nil {
    log.Fatal(err)
}
```

In this example, we generate the configuration we need as it stores general information. Here you can store for example the path to the folder of the log files, or also whether a log in the command line or only file-based. Both and is also possible.

If you want to log file based, then we create in a specified path a folder per month and for each day a separate log file in this folder, so you always get the best overview of your logs.

Here you can find a sample log:

```console
INFO[2022-08-26T11:51:30]  This is a informational message.	FILE=log_test.go:65
```

We color the log kinds so that you can differentiate them immediately, so `error` is shown as red and `warning` as yellow. Everything else is displayed in cyan.

## Special thanks

Thanks to [JetBrains](https://github.com/JetBrains) for supporting me with this and other [open source projects](https://www.jetbrains.com/community/opensource/#support).