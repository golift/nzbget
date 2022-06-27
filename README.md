# nzbget

[![GoDoc](https://godoc.org/golift.io/nzbget/svc?status.svg)](https://pkg.go.dev/golift.io/nzbget)
[![Go Report Card](https://goreportcard.com/badge/golift.io/nzbget)](https://goreportcard.com/report/golift.io/nzbget)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://github.com/golift/nzbget/blob/main/LICENSE)
[![discord](https://badgen.net/badge/icon/Discord?color=0011ff&label&icon=https://simpleicons.now.sh/discord/eee "GoLift Discord")](https://golift.io/discord)

Go Library used to interact with NZBGet. Almost full-featured. See below.

If you'd like new features, please open a GitHub issue or pull request.

## Methods

### Program control

- [x] [version](https://nzbget.net/api/version)
- [x] [shutdown](https://nzbget.net/api/shutdown)
- [x] [reload](https://nzbget.net/api/reload)

### Queue and history

- [x] [listgroups](https://nzbget.net/api/listgroups)
- [x] [listfiles](https://nzbget.net/api/listfiles)
- [x] [history](https://nzbget.net/api/history)
- [ ] [append](https://nzbget.net/api/append)
- [ ] [editqueue](https://nzbget.net/api/editqueue)
- [x] [scan](https://nzbget.net/api/scan)

### Status, logging and statistics

- [x] [status](https://nzbget.net/api/status)
- [x] [log](https://nzbget.net/api/log)
- [x] [writelog](https://nzbget.net/api/writelog)
- [x] [loadlog](https://nzbget.net/api/loadlog)
- [ ] [servervolumes](https://nzbget.net/api/servervolumes)
- [ ] [resetservervolume](https://nzbget.net/api/resetservervolume)

### Pause and speed limit

- [x] [rate](https://nzbget.net/api/rate)
- [x] [pausedownload](https://nzbget.net/api/pausedownload)
- [x] [resumedownload](https://nzbget.net/api/resumedownload)
- [x] [pausepost](https://nzbget.net/api/pausepost)
- [x] [resumepost](https://nzbget.net/api/resumepost)
- [x] [pausescan](https://nzbget.net/api/pausescan)
- [x] [resumescan](https://nzbget.net/api/resumescan)
- [x] [scheduleresume](https://nzbget.net/api/scheduleresume)

### Configuration

- [x] [config](https://nzbget.net/api/config)
- [x] [loadconfig](https://nzbget.net/api/loadconfig)
- [x] [saveconfig](https://nzbget.net/api/saveconfig)
- [ ] [configtemplates](https://nzbget.net/api/configtemplates)