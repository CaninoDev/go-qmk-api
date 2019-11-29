// Package go-qmk-api provides a Go wrapper to QMK's asynchronous API that Web and GUI tools can use to compile
// arbitrary keymaps for any keyboard supported by QMK.
package qmk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"fmt"
)

/* v1 Endpoints
/v1																GET		Returns the API's Status
/v1/update														GET 	Trigger an update of the API
/v1/converters													GET 	Return the list of converters we support
/v1/converters/kle2qmk											POST 	Convert a KLE layout to QMK's layout format
/v1/converters/kle
/v1/keyboards													GET		Return a list of keyboards
/v1/keyboards/all												GET     Return JSON showing all data of all keyboards
/v1/keyboards/<path:keyboard>									GET		Return JSON showing data about a board
/v1/keyboards/<path:keyboard>/readme							GET		Returns the readme for a keyboard
/v1/keyboards/<path:keyboard>/keymaps/<string:keymap>			GET		Return JSON showing data about a board's keymap
/v1/keyboards/<path:keyboard>/keymaps/<string:keymap>/readme	GET		Returns the readme for a keymap
/v1/keyboards/build_status										GET		Returns a dictionary of keyboard/layout pairs.
																		Each entry is True if the keyboard works in
																		configurator and false if it doesn't
/v1/keyboards/build_log											GET		Returns a dictionary of keyboard/layout pairs.
																		Each entry is a dictionary with the following
																		keys:
																		works bool
																		message string
/v1/keyboards/error_log											GET		Return the error log from the last run
/v1/usb															GET		Returns the list of USB device identifiers used
																		in QMK
/v1/compile														POST	Enqeue a compile job
/v1/compile/<string:job_id>										GET		Fetch the status of a compile job
/v1/compile/<string:job_id>/download							GET		Download a compiled firmware
/v1/compile/<string:job_id>/hex
/v1/compile/<string:job_id>/keymap								GET		Download the keymap for a completed compile job
/v1/compile/<string:job_id>/source								GET		Download the full source for a completed compile
																		job
*/

const qmkAPI = "https://api.qmk.fm"
const version = "v1"

var httpClient = &http.Client{
	Timeout: time.Second * 2,
}

type Status struct {
	Result		string		`json:"result",omitempty`
	Children	[]string	`json:"children",omitempty`
	LastPing	string		`json:"last_ping"`
	QueueLength	int			`json:"queue_length"`
	Status		string		`json:"status"`
	Version		string		`json:"version"`
}

// CurrentStatus returns QMK API server status
func CurrentStatus() (Status, error) {
	queryQMK := fmt.Sprintf("%s/%s", qmkAPI, version)
	var body Status
	var rawJSON json.RawMessage
	resp, err := http.Get(queryQMK)
	if err != nil {
		return body, err
	}

	rawJSON, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	err = json.Unmarshal(rawJSON, &body)
	if err != nil {
		return body, err
	}

	return body, nil
}

// Update triggers a QMK API update
func Update() (Status, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s", qmkAPI, version, "update")
	var body Status
	var rawJSON json.RawMessage
	resp, err := http.Get(queryQMK)
	if err != nil {
		return body, err
	}

	rawJSON, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body ,err
	}
	err = json.Unmarshal(rawJSON, &body)
	if err != nil {
		return body, err
	}

	return body, nil
}

// TODO: /v1/converters
// TODO: /v1/converters/kle2qmk
// TODO: /v1/converters/kle
// TODO: /v1/keyboards
// TODO: /v1/keyboards/all
// TODO: /v1/keyboards/<path:keyboard>
// TODO: /v1/keyboards/<path:keyboard>/readme
// TODO: /v1/keyboards/<path:keyboard>/keymaps/<string:keymap>
// TODO: /v1/keyboards/<path:keyboard>/keymaps/<string:keymap>/readme
// TODO: /v1/keyboards/build_status
// TODO: /v1/keyboards/build_log
// TODO: /v1/keyboards/error_log
// TODO: /v1/usb
// TODO: /v1/compile
// TODO: /v1/compile/<string:job_id>
// TODO: /v1/compile/<string:job_id>/download
// TODO: /v1/compile/<string:job_id>/hex
// TODO: /v1/compile/<string:job_id>/keymap
// TODO: /v1/compile/<string:job_id>/source

