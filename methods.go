package nzbget

import "time"

// Version returns the NZBGet Version.
// https://nzbget.net/api/version
func (n *NZBGet) Version() (string, error) {
	var output string
	return output, n.GetInto("version", &output)
}

// ListFiles returns the NZBGet Files for a download.
// https://nzbget.net/api/listfiles
// nzbID is the NZBID of the group to be returned. Use 0 for all file groups.
func (n *NZBGet) ListFiles(nzbID int64) (*File, error) {
	var output File
	return &output, n.GetInto("listfiles", &output, 0, 0, nzbID)
}

// Status returns the NZBGet Status.
// https://nzbget.net/api/status
func (n *NZBGet) Status() (*Status, error) {
	var output Status
	return &output, n.GetInto("status", &output)
}

// History returns the NZBGet Download History.
// https://nzbget.net/api/history
func (n *NZBGet) History(hidden bool) ([]*History, error) {
	var output []*History
	return output, n.GetInto("history", &output, hidden)
}

// ListGroups returns the NZBGet Download list.
// https://nzbget.net/api/listgroups
func (n *NZBGet) ListGroups(limit int64) ([]*Group, error) {
	var output []*Group
	return output, n.GetInto("listgroups", &output, limit)
}

// Log returns the NZBGet Logs.
// NOTE: only one parameter - either startID or limit - can be specified. The other parameter must be 0.
// https://nzbget.net/api/log
func (n *NZBGet) Log(startID, limit int64) ([]*LogEntry, error) {
	var output []*LogEntry
	return output, n.GetInto("log", &output, startID, limit)
}

// LLoadLogog returns the NZBGet log for a specific download.
// NOTE: only one of either startID or limit - can be specified. The other parameter must be 0.
// https://nzbget.net/api/loadlog
func (n *NZBGet) LoadLog(nzbID, startID, limit int64) ([]*LogEntry, error) {
	var output []*LogEntry
	return output, n.GetInto("loadlog", &output, nzbID, startID, limit)
}

// Config returns the loaded and active NZBGet configuration parameters.
// https://nzbget.net/api/config
func (n *NZBGet) Config() ([]*Parameter, error) {
	var output []*Parameter
	return output, n.GetInto("config", &output)
}

// LoadConfig returns the configuration from disk.
// https://nzbget.net/api/loadconfig
func (n *NZBGet) LoadConfig() ([]*Parameter, error) {
	var output []*Parameter
	return output, n.GetInto("loadconfig", &output)
}

// SaveConfig writes new configuration parameters to disk.
// https://nzbget.net/api/saveconfig
func (n *NZBGet) SaveConfig(configs []*Parameter) (bool, error) {
	var (
		output = false
		input  = []interface{}{}
	)

	for _, config := range configs {
		input = append(input, config)
	}

	return output, n.GetInto("saveconfig", &output, input...)
}

// Shutdown makes NZBGet exit.
// https://nzbget.net/api/shutdown
func (n *NZBGet) Shutdown() (bool, error) {
	var output bool
	return output, n.GetInto("shutdown", &output)
}

// Reload makes NZBGet stop all activities and reinitialize.
// https://nzbget.net/api/reload
func (n *NZBGet) Reload() (bool, error) {
	var output bool
	return output, n.GetInto("reload", &output)
}

// Rate sets download speed limit.
// https://nzbget.net/api/rate
func (n *NZBGet) Rate(limit int64) (bool, error) {
	var output bool
	return output, n.GetInto("rate", &output, limit)
}

// PausePost pauses post processing.
// https://nzbget.net/api/pausepost
func (n *NZBGet) PausePost() (bool, error) {
	var output bool
	return output, n.GetInto("pausepost", &output)
}

// ResumePost resumes post processing.
// https://nzbget.net/api/resumepost
func (n *NZBGet) ResumePost() (bool, error) {
	var output bool
	return output, n.GetInto("resumepost", &output)
}

// PauseDownload pauses downloads.
// https://nzbget.net/api/pausedownload
func (n *NZBGet) PauseDownload() (bool, error) {
	var output bool
	return output, n.GetInto("pausedownload", &output)
}

// ResumeDownload resumes downloads.
// https://nzbget.net/api/resumedownload
func (n *NZBGet) ResumeDownload() (bool, error) {
	var output bool
	return output, n.GetInto("resumedownload", &output)
}

// PauseScan pauses scanning of directory with incoming nzb-files.
// https://nzbget.net/api/pausescan
func (n *NZBGet) PauseScan() (bool, error) {
	var output bool
	return output, n.GetInto("pausescan", &output)
}

// ResumeScan resumes scanning of directory with incoming nzb-files.
// https://nzbget.net/api/resumescan
func (n *NZBGet) ResumeScan() (bool, error) {
	var output bool
	return output, n.GetInto("resumescan", &output)
}

// ScheduleResume schedules resuming of all activities after wait duration elapses.
// https://nzbget.net/api/scheduleresume
func (n *NZBGet) ScheduleResume(wait time.Duration) (bool, error) {
	var output bool
	return output, n.GetInto("scheduleresume", &output, wait.Seconds())
}

// Scan requests rescanning of incoming directory for nzb-files.
// https://nzbget.net/api/scheduleresume
func (n *NZBGet) Scan() (bool, error) {
	var output bool
	return output, n.GetInto("scan", &output)
}

// WriteLog appends a log entry to th server's log and on-screen log buffer.
// https://nzbget.net/api/writelog
func (n *NZBGet) WriteLog(kind LogKind, text string) (bool, error) {
	var output bool
	return output, n.GetInto("writelog", &output, kind, text)
}

// ConfigTemplates returns NZBGet configuration file template and
// also extracts configuration sections from all post-processing files.
// This information is for example used by web-interface to build settings
// page or page “Postprocess” in download details dialog.
// https://nzbget.net/api/configtemplates
func (n *NZBGet) ConfigTemplates(loadFromDisk bool) ([]*ConfigTemplate, error) {
	var output []*ConfigTemplate
	return output, n.GetInto("configtemplates", &output, loadFromDisk)
}

// ServerVolumes returns download volume statistics per news-server.
// https://nzbget.net/api/servervolumes
// NOTE: The first record (serverid=0) are totals for all servers.
func (n *NZBGet) ServerVolumes() ([]*ServerVolume, error) {
	var output []*ServerVolume
	return output, n.GetInto("servervolumes", &output)
}

// ResetServerVolume resets download volume statistics for a specified news-server.
// https://nzbget.net/api/resetservervolume
func (n *NZBGet) ResetServerVolume(serverID int64, sounter string) (bool, error) {
	var output bool
	return output, n.GetInto("resetservervolume", &output, serverID, sounter)
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
func (n *NZBGet) Append(appendReq *AppendInput) (int64, error) {
	var (
		output int64
		input  = []interface{}{
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
	)

	for _, param := range appendReq.Parameters {
		input = append(input, param)
	}

	return output, n.GetInto("append", &output, input...)
}

// EditQueue edits items in download queue or in history.
// Read the official docs for how to issue commands, and which commands are available.
// https://nzbget.net/api/editqueue
func (n *NZBGet) EditQueue(command, parameter string, ids []int64) (bool, error) {
	var output bool
	return output, n.GetInto("editqueue", &output, command, parameter, ids)
}
