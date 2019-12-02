// Package qmk provides a Go wrapper to QMK's asynchronous API that Web and GUI tools can use to compile arbitrary keymaps for any keyboard supported by QMK.
package qmk

import (
	"encoding/json"
	"io/ioutil"
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

func TestKeyboardData(t *testing.T) {
	var kb Keyboard
	var empty Keyboard
	preonic := `{"git_hash":"f0991420040b8c43529bd1d863c1987fcb8850e4","last_updated":"2019-11-30 13:05:20 UTC","keyboards":{"preonic/rev3":{"processor":"STM32F303","keyboard_name":"Preonic rev. 3","width":12,"device_ver":"0x0003","identifier":"unknown:unknown:0x0003","readme":true,"platform":"STM32","maintainer":"jackhumbert","keymaps":["cranium","blake-newman","seph","fsck","kuatsure","mikethetiger","kjwon15","that_canadian","choromanski","laurentlaurent","senseored","CMD-Preonic","kinesis","dlaroe","xulkal","muzfuz","trigotometry","mguterl","juno","ekis_isa","smt","pitty","jacwib","dudeofawesome","bucktooth","0xdec","fig-r","zach","default","egstad","boy314","nikchi"],"keyboard_folder":"preonic/rev3","processor_type":"arm","url":"https://olkb.com/preonic","layouts":{"LAYOUT_preonic_1x2uC":{"layout":[{"y":0,"x":0,"label":"k00"},{"y":0,"x":1,"label":"k01"},{"y":0,"x":2,"label":"k02"},{"y":0,"x":3,"label":"k03"},{"y":0,"x":4,"label":"k04"},{"y":0,"x":5,"label":"k05"},{"y":0,"x":6,"label":"k06"},{"y":0,"x":7,"label":"k07"},{"y":0,"x":8,"label":"k08"},{"y":0,"x":9,"label":"k09"},{"y":0,"x":10,"label":"k0a"},{"y":0,"x":11,"label":"k0b"},{"y":1,"x":0,"label":"k10"},{"y":1,"x":1,"label":"k11"},{"y":1,"x":2,"label":"k12"},{"y":1,"x":3,"label":"k13"},{"y":1,"x":4,"label":"k14"},{"y":1,"x":5,"label":"k15"},{"y":1,"x":6,"label":"k16"},{"y":1,"x":7,"label":"k17"},{"y":1,"x":8,"label":"k18"},{"y":1,"x":9,"label":"k19"},{"y":1,"x":10,"label":"k1a"},{"y":1,"x":11,"label":"k1b"},{"y":2,"x":0,"label":"k20"},{"y":2,"x":1,"label":"k21"},{"y":2,"x":2,"label":"k22"},{"y":2,"x":3,"label":"k23"},{"y":2,"x":4,"label":"k24"},{"y":2,"x":5,"label":"k25"},{"y":2,"x":6,"label":"k26"},{"y":2,"x":7,"label":"k27"},{"y":2,"x":8,"label":"k28"},{"y":2,"x":9,"label":"k29"},{"y":2,"x":10,"label":"k2a"},{"y":2,"x":11,"label":"k2b"},{"y":3,"x":0,"label":"k30"},{"y":3,"x":1,"label":"k31"},{"y":3,"x":2,"label":"k32"},{"y":3,"x":3,"label":"k33"},{"y":3,"x":4,"label":"k34"},{"y":3,"x":5,"label":"k35"},{"y":3,"x":6,"label":"k36"},{"y":3,"x":7,"label":"k37"},{"y":3,"x":8,"label":"k38"},{"y":3,"x":9,"label":"k39"},{"y":3,"x":10,"label":"k3a"},{"y":3,"x":11,"label":"k3b"},{"y":4,"x":0,"label":"k40"},{"y":4,"x":1,"label":"k41"},{"y":4,"x":2,"label":"k42"},{"y":4,"x":3,"label":"k43"},{"y":4,"x":4,"label":"k44"},{"y":4,"x":5,"label":"k45","w":2},{"y":4,"x":7,"label":"k47"},{"y":4,"x":8,"label":"k48"},{"y":4,"x":9,"label":"k49"},{"y":4,"x":10,"label":"k4a"},{"y":4,"x":11,"label":"k4b"}],"key_count":59},"LAYOUT_preonic_2x2u":{"layout":[{"y":0,"x":0,"label":"k00"},{"y":0,"x":1,"label":"k01"},{"y":0,"x":2,"label":"k02"},{"y":0,"x":3,"label":"k03"},{"y":0,"x":4,"label":"k04"},{"y":0,"x":5,"label":"k05"},{"y":0,"x":6,"label":"k06"},{"y":0,"x":7,"label":"k07"},{"y":0,"x":8,"label":"k08"},{"y":0,"x":9,"label":"k09"},{"y":0,"x":10,"label":"k0a"},{"y":0,"x":11,"label":"k0b"},{"y":1,"x":0,"label":"k10"},{"y":1,"x":1,"label":"k11"},{"y":1,"x":2,"label":"k12"},{"y":1,"x":3,"label":"k13"},{"y":1,"x":4,"label":"k14"},{"y":1,"x":5,"label":"k15"},{"y":1,"x":6,"label":"k16"},{"y":1,"x":7,"label":"k17"},{"y":1,"x":8,"label":"k18"},{"y":1,"x":9,"label":"k19"},{"y":1,"x":10,"label":"k1a"},{"y":1,"x":11,"label":"k1b"},{"y":2,"x":0,"label":"k20"},{"y":2,"x":1,"label":"k21"},{"y":2,"x":2,"label":"k22"},{"y":2,"x":3,"label":"k23"},{"y":2,"x":4,"label":"k24"},{"y":2,"x":5,"label":"k25"},{"y":2,"x":6,"label":"k26"},{"y":2,"x":7,"label":"k27"},{"y":2,"x":8,"label":"k28"},{"y":2,"x":9,"label":"k29"},{"y":2,"x":10,"label":"k2a"},{"y":2,"x":11,"label":"k2b"},{"y":3,"x":0,"label":"k30"},{"y":3,"x":1,"label":"k31"},{"y":3,"x":2,"label":"k32"},{"y":3,"x":3,"label":"k33"},{"y":3,"x":4,"label":"k34"},{"y":3,"x":5,"label":"k35"},{"y":3,"x":6,"label":"k36"},{"y":3,"x":7,"label":"k37"},{"y":3,"x":8,"label":"k38"},{"y":3,"x":9,"label":"k39"},{"y":3,"x":10,"label":"k3a"},{"y":3,"x":11,"label":"k3b"},{"y":4,"x":0,"label":"k40"},{"y":4,"x":1,"label":"k41"},{"y":4,"x":2,"label":"k42"},{"y":4,"x":3,"label":"k43"},{"y":4,"x":4,"label":"k44","w":2},{"y":4,"x":6,"label":"k46","w":2},{"y":4,"x":8,"label":"k48"},{"y":4,"x":9,"label":"k49"},{"y":4,"x":10,"label":"k4a"},{"y":4,"x":11,"label":"k4b"}],"key_count":58},"LAYOUT_preonic_1x2uR":{"layout":[{"y":0,"x":0,"label":"k00"},{"y":0,"x":1,"label":"k01"},{"y":0,"x":2,"label":"k02"},{"y":0,"x":3,"label":"k03"},{"y":0,"x":4,"label":"k04"},{"y":0,"x":5,"label":"k05"},{"y":0,"x":6,"label":"k06"},{"y":0,"x":7,"label":"k07"},{"y":0,"x":8,"label":"k08"},{"y":0,"x":9,"label":"k09"},{"y":0,"x":10,"label":"k0a"},{"y":0,"x":11,"label":"k0b"},{"y":1,"x":0,"label":"k10"},{"y":1,"x":1,"label":"k11"},{"y":1,"x":2,"label":"k12"},{"y":1,"x":3,"label":"k13"},{"y":1,"x":4,"label":"k14"},{"y":1,"x":5,"label":"k15"},{"y":1,"x":6,"label":"k16"},{"y":1,"x":7,"label":"k17"},{"y":1,"x":8,"label":"k18"},{"y":1,"x":9,"label":"k19"},{"y":1,"x":10,"label":"k1a"},{"y":1,"x":11,"label":"k1b"},{"y":2,"x":0,"label":"k20"},{"y":2,"x":1,"label":"k21"},{"y":2,"x":2,"label":"k22"},{"y":2,"x":3,"label":"k23"},{"y":2,"x":4,"label":"k24"},{"y":2,"x":5,"label":"k25"},{"y":2,"x":6,"label":"k26"},{"y":2,"x":7,"label":"k27"},{"y":2,"x":8,"label":"k28"},{"y":2,"x":9,"label":"k29"},{"y":2,"x":10,"label":"k2a"},{"y":2,"x":11,"label":"k2b"},{"y":3,"x":0,"label":"k30"},{"y":3,"x":1,"label":"k31"},{"y":3,"x":2,"label":"k32"},{"y":3,"x":3,"label":"k33"},{"y":3,"x":4,"label":"k34"},{"y":3,"x":5,"label":"k35"},{"y":3,"x":6,"label":"k36"},{"y":3,"x":7,"label":"k37"},{"y":3,"x":8,"label":"k38"},{"y":3,"x":9,"label":"k39"},{"y":3,"x":10,"label":"k3a"},{"y":3,"x":11,"label":"k3b"},{"y":4,"x":0,"label":"k40"},{"y":4,"x":1,"label":"k41"},{"y":4,"x":2,"label":"k42"},{"y":4,"x":3,"label":"k43"},{"y":4,"x":4,"label":"k44"},{"y":4,"x":5,"label":"k45"},{"y":4,"x":6,"label":"k47","w":2},{"y":4,"x":8,"label":"k48"},{"y":4,"x":9,"label":"k49"},{"y":4,"x":10,"label":"k4a"},{"y":4,"x":11,"label":"k4b"}],"key_count":59},"LAYOUT_preonic_1x2uL":{"layout":[{"y":0,"x":0,"label":"k00"},{"y":0,"x":1,"label":"k01"},{"y":0,"x":2,"label":"k02"},{"y":0,"x":3,"label":"k03"},{"y":0,"x":4,"label":"k04"},{"y":0,"x":5,"label":"k05"},{"y":0,"x":6,"label":"k06"},{"y":0,"x":7,"label":"k07"},{"y":0,"x":8,"label":"k08"},{"y":0,"x":9,"label":"k09"},{"y":0,"x":10,"label":"k0a"},{"y":0,"x":11,"label":"k0b"},{"y":1,"x":0,"label":"k10"},{"y":1,"x":1,"label":"k11"},{"y":1,"x":2,"label":"k12"},{"y":1,"x":3,"label":"k13"},{"y":1,"x":4,"label":"k14"},{"y":1,"x":5,"label":"k15"},{"y":1,"x":6,"label":"k16"},{"y":1,"x":7,"label":"k17"},{"y":1,"x":8,"label":"k18"},{"y":1,"x":9,"label":"k19"},{"y":1,"x":10,"label":"k1a"},{"y":1,"x":11,"label":"k1b"},{"y":2,"x":0,"label":"k20"},{"y":2,"x":1,"label":"k21"},{"y":2,"x":2,"label":"k22"},{"y":2,"x":3,"label":"k23"},{"y":2,"x":4,"label":"k24"},{"y":2,"x":5,"label":"k25"},{"y":2,"x":6,"label":"k26"},{"y":2,"x":7,"label":"k27"},{"y":2,"x":8,"label":"k28"},{"y":2,"x":9,"label":"k29"},{"y":2,"x":10,"label":"k2a"},{"y":2,"x":11,"label":"k2b"},{"y":3,"x":0,"label":"k30"},{"y":3,"x":1,"label":"k31"},{"y":3,"x":2,"label":"k32"},{"y":3,"x":3,"label":"k33"},{"y":3,"x":4,"label":"k34"},{"y":3,"x":5,"label":"k35"},{"y":3,"x":6,"label":"k36"},{"y":3,"x":7,"label":"k37"},{"y":3,"x":8,"label":"k38"},{"y":3,"x":9,"label":"k39"},{"y":3,"x":10,"label":"k3a"},{"y":3,"x":11,"label":"k3b"},{"y":4,"x":0,"label":"k40"},{"y":4,"x":1,"label":"k41"},{"y":4,"x":2,"label":"k42"},{"y":4,"x":3,"label":"k43"},{"y":4,"x":4,"label":"k44","w":2},{"y":4,"x":6,"label":"k46"},{"y":4,"x":7,"label":"k47"},{"y":4,"x":8,"label":"k48"},{"y":4,"x":9,"label":"k49"},{"y":4,"x":10,"label":"k4a"},{"y":4,"x":11,"label":"k4b"}],"key_count":59},"LAYOUT_preonic_grid":{"layout":[{"y":0,"label":"k00","x":0,"w":1},{"y":0,"label":"k01","x":1,"w":1},{"y":0,"label":"k02","x":2,"w":1},{"y":0,"label":"k03","x":3,"w":1},{"y":0,"label":"k04","x":4,"w":1},{"y":0,"label":"k05","x":5,"w":1},{"y":0,"label":"k06","x":6,"w":1},{"y":0,"label":"k07","x":7,"w":1},{"y":0,"label":"k08","x":8,"w":1},{"y":0,"label":"k09","x":9,"w":1},{"y":0,"label":"k0a","x":10,"w":1},{"y":0,"label":"k0b","x":11,"w":1},{"y":1,"label":"k10","x":0,"w":1},{"y":1,"label":"k11","x":1,"w":1},{"y":1,"label":"k12","x":2,"w":1},{"y":1,"label":"k13","x":3,"w":1},{"y":1,"label":"k14","x":4,"w":1},{"y":1,"label":"k15","x":5,"w":1},{"y":1,"label":"k16","x":6,"w":1},{"y":1,"label":"k17","x":7,"w":1},{"y":1,"label":"k18","x":8,"w":1},{"y":1,"label":"k19","x":9,"w":1},{"y":1,"label":"k1a","x":10,"w":1},{"y":1,"label":"k1b","x":11,"w":1},{"y":2,"label":"k20","x":0,"w":1},{"y":2,"label":"k21","x":1,"w":1},{"y":2,"label":"k22","x":2,"w":1},{"y":2,"label":"k23","x":3,"w":1},{"y":2,"label":"k24","x":4,"w":1},{"y":2,"label":"k25","x":5,"w":1},{"y":2,"label":"k26","x":6,"w":1},{"y":2,"label":"k27","x":7,"w":1},{"y":2,"label":"k28","x":8,"w":1},{"y":2,"label":"k29","x":9,"w":1},{"y":2,"label":"k2a","x":10,"w":1},{"y":2,"label":"k2b","x":11,"w":1},{"y":3,"label":"k30","x":0,"w":1},{"y":3,"label":"k31","x":1,"w":1},{"y":3,"label":"k32","x":2,"w":1},{"y":3,"label":"k33","x":3,"w":1},{"y":3,"label":"k34","x":4,"w":1},{"y":3,"label":"k35","x":5,"w":1},{"y":3,"label":"k36","x":6,"w":1},{"y":3,"label":"k37","x":7,"w":1},{"y":3,"label":"k38","x":8,"w":1},{"y":3,"label":"k39","x":9,"w":1},{"y":3,"label":"k3a","x":10,"w":1},{"y":3,"label":"k3b","x":11,"w":1},{"y":4,"label":"k40","x":0,"w":1},{"y":4,"label":"k41","x":1,"w":1},{"y":4,"label":"k42","x":2,"w":1},{"y":4,"label":"k43","x":3,"w":1},{"y":4,"label":"k44","x":4,"w":1},{"y":4,"label":"k45","x":5,"w":1},{"y":4,"label":"k46","x":6,"w":1},{"y":4,"label":"k47","x":7,"w":1},{"y":4,"label":"k48","x":8,"w":1},{"y":4,"label":"k49","x":9,"w":1},{"y":4,"label":"k4a","x":10,"w":1},{"y":4,"label":"k4b","x":11,"w":1}],"key_count":60}},"bootloader":"stm32-dfu","height":5}}}`
	_ = json.Unmarshal([]byte(preonic), &kb)
	type args struct {
		keyboard string
	}
	tests := []struct {
		name    string
		args    args
		want    Keyboard
		wantErr bool
	}{
		{
			name:    "Preonic",
			args:    args{keyboard: "preonic/rev3"},
			want:    kb,
			wantErr: false,
		},
		{
			name:    "NonExistentKeyboard",
			args:    args{keyboard: "NonExistentKeyboard"},
			want:    empty,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KeyboardData(tt.args.keyboard)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyboardData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyboardData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyboardReadme(t *testing.T) {
	var empty []byte
	readme, _ := ioutil.ReadFile("plaid_readme_test")

	type args struct {
		keyboard string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Plaid",
			args: args{
				keyboard: "plaid",
			},
			want:    readme,
			wantErr: false,
		},
		{
			name: "Nonens",
			args: args{
				keyboard: "Nonens",
			},
			want:    empty,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KeyboardReadme(tt.args.keyboard)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyboarReadme() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyboardReadme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeymapData(t *testing.T) {
	var plaidKb Keymap
	wantKb := []byte(`{"git_hash":"f0991420040b8c43529bd1d863c1987fcb8850e4","last_updated":"2019-11-30 13:05:20 UTC","keyboards":{"plaid":{"processor":"atmega328p","keyboard_name":"Plaid // Through Hole","vendor_id":"0x16C0","width":12,"product_id":"0x27DB","device_ver":"0x0002","manufacturer":"dm9records","readme":true,"platform":"unknown","description":"12x4 ortholinear keyboard with through hole components","maintainer":"hsgw","keymaps":{"thehalfdeafchef":{"layers":[],"keymap_name":"thehalfdeafchef","keyboard_name":"plaid","layout_macro":"LAYOUT_planck_mit","keymap_folder":"qmk_firmware/keyboards/plaid/keymaps"}},"keyboard_folder":"plaid","processor_type":"avr","url":"https://github.com/hsgw/plaid","layouts":{"LAYOUT_ortho_4x12":{"layout":[{"y":0,"x":0,"w":1},{"y":0,"x":1,"w":1},{"y":0,"x":2,"w":1},{"y":0,"x":3,"w":1},{"y":0,"x":4,"w":1},{"y":0,"x":5,"w":1},{"y":0,"x":6,"w":1},{"y":0,"x":7,"w":1},{"y":0,"x":8,"w":1},{"y":0,"x":9,"w":1},{"y":0,"x":10,"w":1},{"y":0,"x":11,"w":1},{"y":1,"x":0,"w":1},{"y":1,"x":1,"w":1},{"y":1,"x":2,"w":1},{"y":1,"x":3,"w":1},{"y":1,"x":4,"w":1},{"y":1,"x":5,"w":1},{"y":1,"x":6,"w":1},{"y":1,"x":7,"w":1},{"y":1,"x":8,"w":1},{"y":1,"x":9,"w":1},{"y":1,"x":10,"w":1},{"y":1,"x":11,"w":1},{"y":2,"x":0,"w":1},{"y":2,"x":1,"w":1},{"y":2,"x":2,"w":1},{"y":2,"x":3,"w":1},{"y":2,"x":4,"w":1},{"y":2,"x":5,"w":1},{"y":2,"x":6,"w":1},{"y":2,"x":7,"w":1},{"y":2,"x":8,"w":1},{"y":2,"x":9,"w":1},{"y":2,"x":10,"w":1},{"y":2,"x":11,"w":1},{"y":3,"x":0,"w":1},{"y":3,"x":1,"w":1},{"y":3,"x":2,"w":1},{"y":3,"x":3,"w":1},{"y":3,"x":4,"w":1},{"y":3,"x":5,"w":1},{"y":3,"x":6,"w":1},{"y":3,"x":7,"w":1},{"y":3,"x":8,"w":1},{"y":3,"x":9,"w":1},{"y":3,"x":10,"w":1},{"y":3,"x":11,"w":1}],"key_count":48},"LAYOUT_plaid_mit":{"layout":[{"y":0,"x":0,"w":1},{"y":0,"x":1,"w":1},{"y":0,"x":2,"w":1},{"y":0,"x":3,"w":1},{"y":0,"x":4,"w":1},{"y":0,"x":5,"w":1},{"y":0,"x":6,"w":1},{"y":0,"x":7,"w":1},{"y":0,"x":8,"w":1},{"y":0,"x":9,"w":1},{"y":0,"x":10,"w":1},{"y":0,"x":11,"w":1},{"y":1,"x":0,"w":1},{"y":1,"x":1,"w":1},{"y":1,"x":2,"w":1},{"y":1,"x":3,"w":1},{"y":1,"x":4,"w":1},{"y":1,"x":5,"w":1},{"y":1,"x":6,"w":1},{"y":1,"x":7,"w":1},{"y":1,"x":8,"w":1},{"y":1,"x":9,"w":1},{"y":1,"x":10,"w":1},{"y":1,"x":11,"w":1},{"y":2,"x":0,"w":1},{"y":2,"x":1,"w":1},{"y":2,"x":2,"w":1},{"y":2,"x":3,"w":1},{"y":2,"x":4,"w":1},{"y":2,"x":5,"w":1},{"y":2,"x":6,"w":1},{"y":2,"x":7,"w":1},{"y":2,"x":8,"w":1},{"y":2,"x":9,"w":1},{"y":2,"x":10,"w":1},{"y":2,"x":11,"w":1},{"y":3,"x":0,"w":1},{"y":3,"x":1,"w":1},{"y":3,"x":2,"w":1},{"y":3,"x":3,"w":1},{"y":3,"x":4,"w":1},{"y":3,"x":5,"w":2},{"y":3,"x":7,"w":1},{"y":3,"x":8,"w":1},{"y":3,"x":9,"w":1},{"y":3,"x":10,"w":1},{"y":3,"x":11,"w":1}],"key_count":47},"KEYMAP":{"layout":[{"y":0,"x":0,"w":1},{"y":0,"x":1,"w":1},{"y":0,"x":2,"w":1},{"y":0,"x":3,"w":1},{"y":0,"x":4,"w":1},{"y":0,"x":5,"w":1},{"y":0,"x":6,"w":1},{"y":0,"x":7,"w":1},{"y":0,"x":8,"w":1},{"y":0,"x":9,"w":1},{"y":0,"x":10,"w":1},{"y":0,"x":11,"w":1},{"y":1,"x":0,"w":1},{"y":1,"x":1,"w":1},{"y":1,"x":2,"w":1},{"y":1,"x":3,"w":1},{"y":1,"x":4,"w":1},{"y":1,"x":5,"w":1},{"y":1,"x":6,"w":1},{"y":1,"x":7,"w":1},{"y":1,"x":8,"w":1},{"y":1,"x":9,"w":1},{"y":1,"x":10,"w":1},{"y":1,"x":11,"w":1},{"y":2,"x":0,"w":1},{"y":2,"x":1,"w":1},{"y":2,"x":2,"w":1},{"y":2,"x":3,"w":1},{"y":2,"x":4,"w":1},{"y":2,"x":5,"w":1},{"y":2,"x":6,"w":1},{"y":2,"x":7,"w":1},{"y":2,"x":8,"w":1},{"y":2,"x":9,"w":1},{"y":2,"x":10,"w":1},{"y":2,"x":11,"w":1},{"y":3,"x":0,"w":1},{"y":3,"x":1,"w":1},{"y":3,"x":2,"w":1},{"y":3,"x":3,"w":1},{"y":3,"x":4,"w":1},{"y":3,"x":5,"w":1},{"y":3,"x":6,"w":1},{"y":3,"x":7,"w":1},{"y":3,"x":8,"w":1},{"y":3,"x":9,"w":1},{"y":3,"x":10,"w":1},{"y":3,"x":11,"w":1}],"key_count":48},"LAYOUT_plaid_grid":{"layout":[{"y":0,"x":0,"w":1},{"y":0,"x":1,"w":1},{"y":0,"x":2,"w":1},{"y":0,"x":3,"w":1},{"y":0,"x":4,"w":1},{"y":0,"x":5,"w":1},{"y":0,"x":6,"w":1},{"y":0,"x":7,"w":1},{"y":0,"x":8,"w":1},{"y":0,"x":9,"w":1},{"y":0,"x":10,"w":1},{"y":0,"x":11,"w":1},{"y":1,"x":0,"w":1},{"y":1,"x":1,"w":1},{"y":1,"x":2,"w":1},{"y":1,"x":3,"w":1},{"y":1,"x":4,"w":1},{"y":1,"x":5,"w":1},{"y":1,"x":6,"w":1},{"y":1,"x":7,"w":1},{"y":1,"x":8,"w":1},{"y":1,"x":9,"w":1},{"y":1,"x":10,"w":1},{"y":1,"x":11,"w":1},{"y":2,"x":0,"w":1},{"y":2,"x":1,"w":1},{"y":2,"x":2,"w":1},{"y":2,"x":3,"w":1},{"y":2,"x":4,"w":1},{"y":2,"x":5,"w":1},{"y":2,"x":6,"w":1},{"y":2,"x":7,"w":1},{"y":2,"x":8,"w":1},{"y":2,"x":9,"w":1},{"y":2,"x":10,"w":1},{"y":2,"x":11,"w":1},{"y":3,"x":0,"w":1},{"y":3,"x":1,"w":1},{"y":3,"x":2,"w":1},{"y":3,"x":3,"w":1},{"y":3,"x":4,"w":1},{"y":3,"x":5,"w":1},{"y":3,"x":6,"w":1},{"y":3,"x":7,"w":1},{"y":3,"x":8,"w":1},{"y":3,"x":9,"w":1},{"y":3,"x":10,"w":1},{"y":3,"x":11,"w":1}],"key_count":48},"LAYOUT_planck_mit":{"layout":[{"y":0,"x":0,"w":1},{"y":0,"x":1,"w":1},{"y":0,"x":2,"w":1},{"y":0,"x":3,"w":1},{"y":0,"x":4,"w":1},{"y":0,"x":5,"w":1},{"y":0,"x":6,"w":1},{"y":0,"x":7,"w":1},{"y":0,"x":8,"w":1},{"y":0,"x":9,"w":1},{"y":0,"x":10,"w":1},{"y":0,"x":11,"w":1},{"y":1,"x":0,"w":1},{"y":1,"x":1,"w":1},{"y":1,"x":2,"w":1},{"y":1,"x":3,"w":1},{"y":1,"x":4,"w":1},{"y":1,"x":5,"w":1},{"y":1,"x":6,"w":1},{"y":1,"x":7,"w":1},{"y":1,"x":8,"w":1},{"y":1,"x":9,"w":1},{"y":1,"x":10,"w":1},{"y":1,"x":11,"w":1},{"y":2,"x":0,"w":1},{"y":2,"x":1,"w":1},{"y":2,"x":2,"w":1},{"y":2,"x":3,"w":1},{"y":2,"x":4,"w":1},{"y":2,"x":5,"w":1},{"y":2,"x":6,"w":1},{"y":2,"x":7,"w":1},{"y":2,"x":8,"w":1},{"y":2,"x":9,"w":1},{"y":2,"x":10,"w":1},{"y":2,"x":11,"w":1},{"y":3,"x":0,"w":1},{"y":3,"x":1,"w":1},{"y":3,"x":2,"w":1},{"y":3,"x":3,"w":1},{"y":3,"x":4,"w":1},{"y":3,"x":5,"w":2},{"y":3,"x":7,"w":1},{"y":3,"x":8,"w":1},{"y":3,"x":9,"w":1},{"y":3,"x":10,"w":1},{"y":3,"x":11,"w":1}],"key_count":47}},"bootloader":"USBasp","height":4,"identifier":"0x16C0:0x27DB:0x0002"}}}`)

	json.Unmarshal(wantKb, &plaidKb)

	type args struct {
		keyboard string
		keymap   string
	}
	tests := []struct {
		name    string
		args    args
		want    Keymap
		wantErr bool
	}{
		{
			name: "Plaid",
			args: args{
				keyboard: "plaid",
				keymap:   "thehalfdeafchef",
			},
			want:    plaidKb,
			wantErr: false,
		},
		{
			name: "Nonense",
			args: args{
				keyboard: "Nonense",
				keymap:   "ohashoi",
			},
			want:    Keymap{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KeymapData(tt.args.keyboard, tt.args.keymap)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeymapData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeymapData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeymapReadme(t *testing.T) {
	var empty []byte
	readme, _ := ioutil.ReadFile("plaid_thehalfdeafchef_readme_test")

	type args struct {
		keyboard string
		keymap   string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Plaid",
			args: args{
				keyboard: "plaid",
				keymap:   "thehalfdeafchef",
			},
			want:    readme,
			wantErr: false,
		},
		{
			name: "Nonens",
			args: args{
				keyboard: "Nonens",
				keymap:   "none",
			},
			want:    empty,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KeymapReadme(tt.args.keyboard, tt.args.keymap)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeymapReadme() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeymapReadme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyboardLayoutBuildStatus(t *testing.T) {
	i, err := KeyboardLayoutBuildStatus()

	if reflect.TypeOf(i) != reflect.TypeOf(BuildStatus{}) {
		t.Errorf("Got %s, wants %s", reflect.TypeOf(i), reflect.TypeOf(BuildStatus{}))
	}

	if err != nil {
		t.Error(err)
	}
}

func TestLayoutBuildLog(t *testing.T) {
	i, err := LayoutBuildLog()

	if reflect.TypeOf(i) != reflect.TypeOf(BuildLog{}) {
		t.Errorf("Got %s, wants %s", reflect.TypeOf(i), reflect.TypeOf(BuildLog{}))
	}

	if err != nil {
		t.Error(err)
	}
}
