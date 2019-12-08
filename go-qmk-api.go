// Package qmk provides a Go wrapper to QMK's asynchronous API that Web and GUI tools can use to compile arbitrary keymaps for any keyboard supported by QMK.
package qmk

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

// GET Endpoints

// POST Endpoints

// TODO: /v1/compile
// TODO: /v1/compile/<string:job_id>
// TODO: /v1/compile/<string:job_id>/download
// TODO: /v1/compile/<string:job_id>/hex
// TODO: /v1/compile/<string:job_id>/keymap
// TODO: /v1/compile/<string:job_id>/source
