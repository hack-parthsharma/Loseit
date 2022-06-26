# loseit
loseit is a simple command line tool to calculate an estimate of daily calorie consumption to meet a weekly weight loss goal.

## Requirements
* [Go](https://github.com/golang/go)

## Install
The latest version of bfm can be installed using `go get`.

```
go get -u github.com/hack-parthsharma/loseit
```

Make sure `$GOPATH` is set correctly that and that `$GOPATH/bin` is in your `$PATH`.

The `loseit` executable will be installed under the `$GOPATH/bin` directory.

## Overview
`loseit` is a simple command line tool that takes a tdee calculation
given as an argument and provides a rough estimate of how many
calories to consume daily in order to achieve the desired weight
loss per week.

`loseit` is designed to be used with input piped from a tool
such as `tdee`. Refer to the [tdee repo](https://github.com/lgug2z/tdee)
for usage instructions.

