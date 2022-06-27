package nzbget

// Version returns the NZBGet Version: https://nzbget.net/api/version.
func (n *NZBGet) Version() (string, error) {
	var output string
	return output, n.GetInto("version", &output)
}

// ListFiles returns the NZBGet Files for a download: https://nzbget.net/api/listfiles.
// nzbID is the NZBID of the group to be returned. Use 0 for all file groups.
func (n *NZBGet) ListFiles(nzbID int64) (*File, error) {
	var output File
	return &output, n.GetInto("listfiles", &output, 0, 0, nzbID)
}

// Status returns the NZBGet Status: https://nzbget.net/api/status.
func (n *NZBGet) Status() (*Status, error) {
	var output Status
	return &output, n.GetInto("status", &output)
}

// History returns the NZBGet Download History: https://nzbget.net/api/history.
func (n *NZBGet) History(hidden bool) ([]*History, error) {
	var output []*History
	return output, n.GetInto("history", &output, hidden)
}

// ListGroups returns the NZBGet Download list: https://nzbget.net/api/listgroups.
func (n *NZBGet) ListGroups(limit int64) ([]*Group, error) {
	var output []*Group
	return output, n.GetInto("listgroups", &output, limit)
}

// Log returns the NZBGet Logs.
// NOTE: only one parameter - either startID or limit - can be specified. The other parameter must be 0.
func (n *NZBGet) Log(startID, limit int64) (interface{}, error) {
	var output []*LogEntry
	return output, n.GetInto("log", &output, startID, limit)
}
