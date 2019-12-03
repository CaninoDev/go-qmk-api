// Package qmk provides a Go wrapper to QMK's asynchronous API that Web and GUI tools can use to compile arbitrary keymaps for any keyboard supported by QMK.
package qmk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/glog"
)

/*
v1 Endpoints
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

// Status represents the QMK API's operating status
type Status struct {
	Result      bool     `json:"result,omitempty"`
	Children    []string `json:"children,omitempty"`
	LastPing    string   `json:"last_ping"`
	QueueLength int      `json:"queue_length"`
	Status      string   `json:"status"`
	Version     string   `json:"version"`
}

// Mapping represent a key's coordinate and orientation on the board's matrix
type Mapping struct {
	Label string  `json:"label,omitempty"`
	X     float32 `json:"x,omitempty"`
	Y     float32 `json:"y,omitempty"`
	H     float32 `json:"h,omitempty"`
	W     float32 `json:"w,omitempty"`
}

// Layout represent a particular layout mapping and key count
type Layout struct {
	KeyCount int       `json:"key_count"`
	Mapping  []Mapping `json:"layout"`
}

// Keyboard represent information about a keyboard
type Keyboard struct {
	BootLoader     string            `json:"bootloader,omitempty"`
	Description    string            `json:"description,omitempty"`
	Keymaps        map[string]Keymap `json:"keymaps"`
	Identifiers    string            `json:"identifiers"`
	Layouts        map[string]Layout `json:"layouts"`
	URL            string            `json:"url"`
	Height         float32           `json:"height"`
	Width          float32           `json:"width"`
	Readme         bool              `json:"readme"`
	VendorID       string            `json:"vendor_id"`
	Processor      string            `json:"processor"`
	ProcessorType  string            `json:"processor_type"`
	DeviceVersion  string            `json:"device_version"`
	Manufacturer   string            `json:"manufacturer"`
	ProductID      string            `json:"product_id"`
	Maintainer     string            `json:"maintainer"`
	KeyboardFolder string            `json:"keyboard_folder"`
	Platform       string            `json:"platform"`
}

// KeyboardsCollection represents information about all keyboards
type KeyboardsCollection struct {
	LastUpdated string `json:"last_updated"`
	Keyboards   map[string]Keyboard
}

// Keymap represents information about a specific keyboard's keymap
type Keymap struct {
	Layers     []string `json:"layer,omitempty"`
	Name       string   `json:"keymap_name,omitempty"`
	Layout     string   `json:"layout_macro,omitempty"`
	FolderName string   `json:"folder_name,omitempty"`
}

// BuildStatus represents a list of keyboards and their build status against QMK compilation
type BuildStatus struct {
	KeyboardLayout map[string]bool
}

// BuildLog represents a mapping of each keyboard's layout's to their respective logs
type BuildLog struct {
	KeyboardLayout map[string]KeyboardLog
}

// KeyboardLog represents the log of a particular keyboard's layout's
type KeyboardLog struct {
	Works      bool
	LastTested string `json:"last_tested"`
	Message    string `json:"message"`
}

// ErrorLog represents the various warnings and errors resulting from a keyboard's layout's
type ErrorLog struct {
	Message  string
	Severity string
}

// USBInfo represents the Product's hex designation
type USBInfo map[string]map[string]map[string]ProductInfo

// ProductInfo represents the product info as returned by the /api/v1/usb endpoint
type ProductInfo struct {
	Description   string
	Keyboard      string
	Vendor        string `json:"vendor_id"`
	Manufacturer  string
	ProductID     string `json:"product_id"`
	DeviceVersion string `json:"device_ver"`
}

type errMessage struct {
	Message string
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
		return body, err
	}
	err = json.Unmarshal(rawJSON, &body)
	if err != nil {
		return body, err
	}

	return body, nil
}

// Converters returns a list of supported format converters
func Converters() ([]string, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s", qmkAPI, version, "converters")
	var bodyRaw map[string][]interface{}
	var body []string
	var rawJSON json.RawMessage

	resp, err := http.Get(queryQMK)
	if err != nil {
		return body, err
	}

	rawJSON, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	err = json.Unmarshal(rawJSON, &bodyRaw)
	if err != nil {
		return body, err
	}

	body = make([]string, len(bodyRaw["children"]))
	for i, interf := range bodyRaw["children"] {
		body[i] = interf.(string)
	}

	return body, nil
}

// TODO: /v1/converters/kle2qmk
// TODO: /v1/converters/kle

// KeyboardsList returns a list of supported keyboards
func KeyboardsList() ([]string, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s", qmkAPI, version, "keyboards")
	var body []string
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

// KeyboardsData returns a list of keyboards and its specifications
func KeyboardsData() (KeyboardsCollection, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s", qmkAPI, version, "keyboards/all")
	var body KeyboardsCollection
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
		var msg string
		switch t := err.(type) {
		case *json.SyntaxError:
			jsn := string(rawJSON[0:t.Offset])
			jsn += "<--(Invalid Character)"
			msg = fmt.Sprintf("Invalid Character at offset %v\n %s", t.Offset, jsn)
		case *json.UnmarshalTypeError:
			jsn := string(rawJSON[0:t.Offset])
			jsn += " <--(Invalid Type) "
			msg = fmt.Sprintf("Invalid Type at offset %v\n %s", t.Offset, jsn)
		default:
			msg = err.Error()
		}
		glog.Warning(msg)
	}

	return body, nil
}

//KeyboardData returns information about the specified keyboard.
func KeyboardData(keyboard string) (Keyboard, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s/%s", qmkAPI, version, "keyboards", keyboard)
	var body Keyboard
	var rawJSON json.RawMessage
	var errMsg errMessage

	resp, err := http.Get(queryQMK)
	if resp.StatusCode == 404 {
		msg, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(msg, &errMsg)
		return body, errors.New(errMsg.Message)
	}
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

// KeyboardReadme returns a specified keyboard's readme
func KeyboardReadme(keyboard string) ([]byte, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s/%s/%s", qmkAPI, version, "keyboards", keyboard, "readme")
	var rawBlob []byte

	resp, err := http.Get(queryQMK)
	if resp.StatusCode == 500 {
		return rawBlob, errors.New("Keyboard Not Found")
	}
	if err != nil {
		return rawBlob, err
	}
	defer resp.Body.Close()

	rawBlob, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return rawBlob, err
	}

	return rawBlob, nil
}

// KeymapData returns data about the specified keyboard's keymap
func KeymapData(keyboard string, keymap string) (Keymap, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s/%s/%s/%s", qmkAPI, version, "keyboards", keyboard, "keymaps", keymap)
	var keyMap Keymap

	resp, err := http.Get(queryQMK)
	if resp.StatusCode == 404 {
		return keyMap, errors.New("Keyboard Not Found")
	}
	if err != nil {
		return keyMap, err
	}
	defer resp.Body.Close()

	rawJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return keyMap, err
	}

	err = json.Unmarshal(rawJSON, &keyMap)

	return keyMap, nil
}

// KeymapReadme returns a specified keyboard's keymap's readme
func KeymapReadme(keyboard string, keymap string) ([]byte, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", qmkAPI, version, "keyboards", keyboard, "keymaps", keymap, "readme")
	var rawBlob []byte

	resp, err := http.Get(queryQMK)
	if resp.StatusCode == 500 {
		return rawBlob, errors.New("Keyboard or Keymap Not Found")
	}
	if err != nil {
		return rawBlob, err
	}
	defer resp.Body.Close()

	rawBlob, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return rawBlob, err
	}

	return rawBlob, nil
}

// KeyboardLayoutBuildStatus returns a list of keyboard and their respective layouts build status against QMK compilation
func KeyboardLayoutBuildStatus() (BuildStatus, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s/%s", qmkAPI, version, "keyboards", "build_status")
	var buildStatus BuildStatus

	resp, err := http.Get(queryQMK)
	if err != nil {
		return buildStatus, err
	}
	defer resp.Body.Close()

	rawJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return buildStatus, err
	}

	err = json.Unmarshal(rawJSON, &buildStatus)

	return buildStatus, err
}

// LayoutBuildLog returns every keyboard's layout's build log
func LayoutBuildLog() (BuildLog, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s/%s", qmkAPI, version, "keyboards", "build_log")
	var buildLog BuildLog

	resp, err := http.Get(queryQMK)
	if err != nil {
		return buildLog, err
	}
	defer resp.Body.Close()

	rawJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return buildLog, err
	}

	err = json.Unmarshal(rawJSON, &buildLog)

	return buildLog, err
}

// ErrorLogs return an array of error logs resulting from compilation of a keyboard's layout
func ErrorLogs() ([]ErrorLog, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s/%s", qmkAPI, version, "keyboards", "error_log")
	var errorLogs []ErrorLog

	resp, err := http.Get(queryQMK)
	if err != nil {
		return errorLogs, err
	}
	defer resp.Body.Close()

	rawJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errorLogs, err
	}

	err = json.Unmarshal(rawJSON, &errorLogs)

	return errorLogs, err
}

// USBTable returns a VendorID lookup hash
func USBTable() (USBInfo, error) {
	queryQMK := fmt.Sprintf("%s/%s/%s", qmkAPI, version, "usb")
	var usbInfo USBInfo

	resp, err := http.Get(queryQMK)
	if err != nil {
		return usbInfo, err
	}
	defer resp.Body.Close()

	rawJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return usbInfo, err
	}

	err = json.Unmarshal(rawJSON, &usbInfo)

	return usbInfo, err
}

// TODO: /v1/compile
// TODO: /v1/compile/<string:job_id>
// TODO: /v1/compile/<string:job_id>/download
// TODO: /v1/compile/<string:job_id>/hex
// TODO: /v1/compile/<string:job_id>/keymap
// TODO: /v1/compile/<string:job_id>/source
