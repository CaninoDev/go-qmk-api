// Package qmk provides a Go wrapper to QMK's asynchronous API that Web and GUI tools can use to compile arbitrary keymaps for any keyboard supported by QMK.
package qmk

import (
	"reflect"
	"testing"
)

func TestCurrentStatus(t *testing.T) {
	i, err := CurrentStatus()

	if reflect.TypeOf(i) != reflect.TypeOf(Status{}) {
		t.Errorf("Got %T, wants %T", reflect.TypeOf(i), reflect.TypeOf(Status{}))
	}

	if err != nil {
		t.Errorf("Got error %s", err)
	}

}

func TestUpdate(t *testing.T) {
	i, err := Update()

	if reflect.TypeOf(i) != reflect.TypeOf(Status{}) {
		t.Errorf("Got %s, wants %s", reflect.TypeOf(i), reflect.TypeOf(Status{}))
	}

	if err != nil {
		t.Errorf("Got error %s", err)
	}

	return

}

func TestConverters(t *testing.T) {
	i, err := Converters()

	if reflect.TypeOf(i) != reflect.TypeOf([]string{}) {
		t.Errorf("Got %s, wants %s", reflect.TypeOf(i), reflect.TypeOf([]string{}))
	}

	if err != nil {
		t.Error(err)
	}

}

func TestKeyboardsList(t *testing.T) {
	i, err := KeyboardsList()

	if reflect.TypeOf(i) != reflect.TypeOf([]string{}) {
		t.Errorf("Got %s, wants %s", reflect.TypeOf(i), reflect.TypeOf([]string{}))
	}

	if err != nil {
		t.Error(err)
	}
}

func TestKeyboardsData(t *testing.T) {
	i, err := KeyboardsData()

	if reflect.TypeOf(i) != reflect.TypeOf(KeyboardsCollection{}) {
		t.Errorf("Got %s, wants %s", reflect.TypeOf(i), reflect.TypeOf(KeyboardsCollection{}))
	}

	if err != nil {
		t.Error(err)
	}
}
