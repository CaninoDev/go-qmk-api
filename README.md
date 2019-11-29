# go-qmk-api : Go wrapper for QMK's API.

## Overview [![GoDoc](https://godoc.org/github.com/caninodev/goqmk?status.svg)](https://godoc.org/github.com/caninodev/goqmk)

go-qmk-api is a Go wrapper for QMK's asynchronous [API](https://docs.api.qmk.fm/) (v1) that Web and GUI tools can use to compile arbitrary keymaps for any keyboard supported by QMK.

## Supported endpoints

```
/v1															
/v1/update													
/v1/converters
/v1/converters/kle2q
/v1/converters/kle               
/v1/keyboards
/v1/keyboards/all
/v1/keyboards/\<path:keyboard\>
/v1/keyboards/\<path:keyboard\>/readme
/v1/keyboards/\<path:keyboard\>/keymaps/\<string:keymap\>
/v1/keyboards/\<path:keyboard\>/keymap/\<string:keymap\>
/v1/keyboards/build_status
/v1/keyboards/build_log
/v1/keyboards/error_log										
/v1/usb														
/v1/compile													
/v1/compile/\<string:job_id\>									
/v1/compile/\<string:job_id\>/download						
/v1/compile/\<string:job_id\>/hex                               
/v1/compile/\<string:job_id\>/keymap							
/v1/compile/\<string:job_id\>/source	
```						

## Install

```
go get github.com/caninodev/go-api-qmk
```

## License

Public Domain.


