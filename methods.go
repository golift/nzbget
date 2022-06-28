package nzbget

import "time"

// Version returns the NZBGet Version.
// https://nzbget.net/api/version
func (n *NZBGet) Version() (string, int, error) {
	var output string
	size, err := n.GetInto("version", &output)

	return output, size, err
}

// ListFiles returns the NZBGet Files for a download.
// https://nzbget.net/api/listfiles
// nzbID is the NZBID of the group to be returned. Use 0 for all file groups.
func (n *NZBGet) ListFiles(nzbID int64) (*File, int, error) {
	var output File
	size, err := n.GetInto("listfiles", &output, 0, 0, nzbID)

	return &output, size, err
}

// Status returns the NZBGet Status.
// https://nzbget.net/api/status
func (n *NZBGet) Status() (*Status, int, error) {
	var output Status
	size, err := n.GetInto("status", &output)

	return &output, size, err
}

// History returns the NZBGet Download History.
// https://nzbget.net/api/history
func (n *NZBGet) History(hidden bool) ([]*History, int, error) {
	var output []*History
	size, err := n.GetInto("history", &output, hidden)

	return output, size, err
}

// ListGroups returns the NZBGet Download list.
// https://nzbget.net/api/listgroups
func (n *NZBGet) ListGroups() ([]*Group, int, error) {
	var output []*Group
	size, err := n.GetInto("listgroups", &output, 0)

	return output, size, err
}

// Log returns the NZBGet Logs.
// NOTE: only one parameter - either startID or limit - can be specified. The other parameter must be 0.
// https://nzbget.net/api/log
func (n *NZBGet) Log(startID, limit int64) ([]*LogEntry, int, error) {
	var output []*LogEntry
	size, err := n.GetInto("log", &output, startID, limit)

	return output, size, err
}

// LLoadLogog returns the NZBGet log for a specific download.
// NOTE: only one of either startID or limit - can be specified. The other parameter must be 0.
// https://nzbget.net/api/loadlog
func (n *NZBGet) LoadLog(nzbID, startID, limit int64) ([]*LogEntry, int, error) {
	var output []*LogEntry
	size, err := n.GetInto("loadlog", &output, nzbID, startID, limit)

	return output, size, err
}

// Config returns the loaded and active NZBGet configuration parameters.
// https://nzbget.net/api/config
func (n *NZBGet) Config() ([]*Parameter, int, error) {
	var output []*Parameter
	size, err := n.GetInto("config", &output)

	return output, size, err
}

// LoadConfig returns the configuration from disk.
// https://nzbget.net/api/loadconfig
func (n *NZBGet) LoadConfig() ([]*Parameter, int, error) {
	var output []*Parameter
	size, err := n.GetInto("loadconfig", &output)

	return output, size, err
}

// SaveConfig writes new configuration parameters to disk.
// https://nzbget.net/api/saveconfig
func (n *NZBGet) SaveConfig(configs []*Parameter) (bool, int, error) {
	input := []interface{}{}

	for _, config := range configs {
		input = append(input, config)
	}

	var output bool
	size, err := n.GetInto("saveconfig", &output, input...)

	return output, size, err
}

// Shutdown makes NZBGet exit.
// https://nzbget.net/api/shutdown
func (n *NZBGet) Shutdown() (bool, int, error) {
	var output bool
	size, err := n.GetInto("shutdown", &output)

	return output, size, err
}

// Reload makes NZBGet stop all activities and reinitialize.
// https://nzbget.net/api/reload
func (n *NZBGet) Reload() (bool, int, error) {
	var output bool
	size, err := n.GetInto("reload", &output)

	return output, size, err
}

// Rate sets download speed limit.
// https://nzbget.net/api/rate
func (n *NZBGet) Rate(limit int64) (bool, int, error) {
	var output bool
	size, err := n.GetInto("rate", &output, limit)

	return output, size, err
}

// PausePost pauses post processing.
// https://nzbget.net/api/pausepost
func (n *NZBGet) PausePost() (bool, int, error) {
	var output bool
	size, err := n.GetInto("pausepost", &output)

	return output, size, err
}

// ResumePost resumes post processing.
// https://nzbget.net/api/resumepost
func (n *NZBGet) ResumePost() (bool, int, error) {
	var output bool
	size, err := n.GetInto("resumepost", &output)

	return output, size, err
}

// PauseDownload pauses downloads.
// https://nzbget.net/api/pausedownload
func (n *NZBGet) PauseDownload() (bool, int, error) {
	var output bool
	size, err := n.GetInto("pausedownload", &output)

	return output, size, err
}

// ResumeDownload resumes downloads.
// https://nzbget.net/api/resumedownload
func (n *NZBGet) ResumeDownload() (bool, int, error) {
	var output bool
	size, err := n.GetInto("resumedownload", &output)

	return output, size, err
}

// PauseScan pauses scanning of directory with incoming nzb-files.
// https://nzbget.net/api/pausescan
func (n *NZBGet) PauseScan() (bool, int, error) {
	var output bool
	size, err := n.GetInto("pausescan", &output)

	return output, size, err
}

// ResumeScan resumes scanning of directory with incoming nzb-files.
// https://nzbget.net/api/resumescan
func (n *NZBGet) ResumeScan() (bool, int, error) {
	var output bool
	size, err := n.GetInto("resumescan", &output)

	return output, size, err
}

// ScheduleResume schedules resuming of all activities after wait duration elapses.
// https://nzbget.net/api/scheduleresume
func (n *NZBGet) ScheduleResume(wait time.Duration) (bool, int, error) {
	var output bool
	size, err := n.GetInto("scheduleresume", &output, wait.Seconds())

	return output, size, err
}

// Scan requests rescanning of incoming directory for nzb-files.
// https://nzbget.net/api/scheduleresume
func (n *NZBGet) Scan() (bool, int, error) {
	var output bool
	size, err := n.GetInto("scan", &output)

	return output, size, err
}

// WriteLog appends a log entry to th server's log and on-screen log buffer.
// https://nzbget.net/api/writelog
func (n *NZBGet) WriteLog(kind LogKind, text string) (bool, int, error) {
	var output bool
	size, err := n.GetInto("writelog", &output, kind, text)

	return output, size, err
}

// ConfigTemplates returns NZBGet configuration file template and
// also extracts configuration sections from all post-processing files.
// This information is for example used by web-interface to build settings
// page or page “Postprocess” in download details dialog.
// https://nzbget.net/api/configtemplates
func (n *NZBGet) ConfigTemplates(loadFromDisk bool) ([]*ConfigTemplate, int, error) {
	var output []*ConfigTemplate
	size, err := n.GetInto("configtemplates", &output, loadFromDisk)

	return output, size, err
}

// ServerVolumes returns download volume statistics per news-server.
// https://nzbget.net/api/servervolumes
// NOTE: The first record (serverid=0) are totals for all servers.
func (n *NZBGet) ServerVolumes() ([]*ServerVolume, int, error) {
	var output []*ServerVolume
	size, err := n.GetInto("servervolumes", &output)

	return output, size, err
}

// ResetServerVolume resets download volume statistics for a specified news-server.
// https://nzbget.net/api/resetservervolume
func (n *NZBGet) ResetServerVolume(serverID int64, sounter string) (bool, int, error) {
	var output bool
	size, err := n.GetInto("resetservervolume", &output, serverID, sounter)

	return output, size, err
}

// AppendInput is the input data for the append method.
// See https://nzbget.net/api/append for more information about this data.
type AppendInput struct {
	Filename   string
	Content    string
	Category   string
	Priority   int64
	AddToTop   bool
	AddPaused  bool
	DupeKey    string // See: https://nzbget.net/rss#duplicate-keys
	DupeScore  int64  // See: https://nzbget.net/rss#duplicate-scores
	DupeMode   string // See: https://nzbget.net/rss#duplicate-modes
	Parameters []*Parameter
}

// Append adds a nzb-file or URL to the download queue.
// https://nzbget.net/api/append
func (n *NZBGet) Append(appendReq *AppendInput) (int64, int, error) {
	input := []interface{}{
		appendReq.Filename,
		appendReq.Content,
		appendReq.Category,
		appendReq.Priority,
		appendReq.AddToTop,
		appendReq.AddPaused,
		appendReq.DupeKey,
		appendReq.DupeScore,
		appendReq.DupeMode,
	}

	for _, param := range appendReq.Parameters {
		input = append(input, param)
	}

	var output int64
	size, err := n.GetInto("append", &output, input...)

	return output, size, err
}

// EditQueue edits items in download queue or in history.
// Read the official docs for how to issue commands, and which commands are available.
// https://nzbget.net/api/editqueue
func (n *NZBGet) EditQueue(command, parameter string, ids []int64) (bool, int, error) {
	var output bool
	size, err := n.GetInto("editqueue", &output, command, parameter, ids)

	return output, size, err
}
