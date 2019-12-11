package qmk

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

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
	KeyCount int       `json:"key_count,omitempty"`
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
	KeyboardName   string            `json:"keyboard_name,omitempty"`
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

// KeyboardsList returns a list of supported keyboards
func (c *Client) KeyboardsList() ([]string, error) {
	var keyboardsList []string
	_, err := c.newRequest("GET", "/keyboards", &keyboardsList, nil)

	return keyboardsList, err
}

// KeyboardsData returns a list of keyboards and its specifications
func (c *Client) KeyboardsCollection() (KeyboardsCollection, error) {
	var keyboardsCollection KeyboardsCollection
	_, err := c.newRequest("GET", "/keyboards/all", &keyboardsCollection, nil)

	return keyboardsCollection, err
}

//KeyboardData returns information about the specified keyboard.
func (c *Client) KeyboardData(keyboardName string) (Keyboard, error) {
	var keyboard Keyboard
	endPointURL := fmt.Sprintf("/keyboards/%s", keyboardName)
	_, err := c.newRequest("GET", endPointURL, &keyboard, nil)

	return keyboard, err
}

// KeyboardReadme returns a specified keyboard's readme
func (c *Client) KeyboardReadme(keyboardName string) (string, error) {
	var readme string

	endPointURL := fmt.Sprintf("/keyboards/%s/%s", keyboardName, "readme")
	resp, err := c.newRequest("GET", endPointURL, nil, nil)

	if resp.Body != nil {
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return readme, err
		}
		readme = string(responseBody)
		return readme, nil
	}
	return readme, err
}

// KeymapData returns data about the specified keyboard's keymap
func (c *Client) KeymapData(keyboardName string, keymapName string) (Keymap, error) {
	var keymap Keymap

	endPointURL := fmt.Sprintf("/keyboards/%s/keymaps/%s", keyboardName, keymapName)
	_, err := c.newRequest("GET", endPointURL, &keymap, nil)
	return keymap, err
}

// KeymapReadme returns a specified keyboard's keymap's readme
func (c *Client) KeymapReadme(keyboardName string, keymapName string) (string, error) {

	endpointURL := fmt.Sprintf("/keyboards/%s/keymaps/%s/readme", keyboardName, keymapName)
	resp, err := c.newRequest("GET", endpointURL, nil, nil)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	readme := buf.String()
	return readme, nil

}
