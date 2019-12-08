package qmk

import (
	"errors"
)

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

// Converters returns a list of supported format converters
func (c *Client) Converters() ([]string, error) {
	var converters []string
	_, err := c.newRequest("GET", "/converters", converters, nil)

	return converters, err
}

// KLE2QMK take in a map with 'id' or 'raw' ('raw' is not working at the moment) key and corresponding value and returrns Keyboard in QMK preferred format
func (c *Client) KLE2QMK(kleMap map[string]string) (Keyboard, error) {
	if _, ok := kleMap["raw"]; ok {
		return Keyboard{}, errors.New("'raw' type is not working at the moment, check back again later")
	}
	var keyboard Keyboard
	resp, err := c.newRequest("GET", "/converters/kle2qmk", &keyboard, kleMap)
	if code := resp.StatusCode; code >= 200 || code <= 300 {
		return keyboard, nil
	}
	return Keyboard{}, err
}

// USBTable returns a VendorID lookup hash
func (c *Client) USBTable() (USBInfo, error) {
	var usbInfo USBInfo
	_, err := c.newRequest("GET", "/usb", &usbInfo, nil)
	return usbInfo, err
}
