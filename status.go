package qmk

// Status represents the QMK API's operating status
type Status struct {
	Result      bool     `json:"result,omitempty"`
	Children    []string `json:"children,omitempty"`
	LastPing    string   `json:"last_ping"`
	QueueLength int      `json:"queue_length"`
	Status      string   `json:"status"`
	Version     string   `json:"version"`
}

// ErrorLog represents the various warnings and errors resulting from a keyboard's layout's
type ErrorLog struct {
	Message  string
	Severity string
}

// BuildStatus represents a list of keyboards and their build status against QMK compilation
type BuildStatus struct {
	KeyboardLayout map[string]bool
}

// BuildLog represents a mapping of each keyboard's layout's to their respective logs
type BuildLog struct {
	KeyboardLayout map[string]KeyboardLog
}

// KeyboardLog represents the log of a particular keyboard's layout's compilation
type KeyboardLog struct {
	Works      bool
	LastTested string `json:"last_tested"`
	Message    string `json:"message"`
}

// CurrentStatus returns QMK API server status
func (c *Client) CurrentStatus() (Status, error) {
	var status Status
	_, err := c.newRequest("GET", "", &status, nil)

	return status, err
}

// Update triggers a QMK API update
func (c *Client) Update() (Status, error) {
	var status Status
	_, err := c.newRequest("GET", "/update", status, nil)

	return status, err
}

// ErrorLogs return an array of error logs resulting from compilation of a keyboard's layout
func (c *Client) ErrorLogs() ([]ErrorLog, error) {
	var errorLogs []ErrorLog
	_, err := c.newRequest("GET", "/keybaords/error_log", &errorLogs, nil)

	return errorLogs, err
}

// KeyboardLayoutBuildStatus returns a list of keyboard and their respective layouts build status against QMK compilation
func (c *Client) KeyboardLayoutBuildStatus() (BuildStatus, error) {
	var buildStatus BuildStatus
	_, err := c.newRequest("GET", "/keyboards/build_status", &buildStatus, nil)

	return buildStatus, err
}

// LayoutBuildLog returns every keyboard's layout's build log
func (c *Client) LayoutBuildLog() (BuildLog, error) {
	var buildLog BuildLog
	_, err := c.newRequest("GET", "/keyboards/build_log", &buildLog, nil)
	return buildLog, err
}
