# nzbget

Go Library used to interact with NZBGet. This library is not full-featured,
and probably only provides read-only methods to view the download list.

If you'd like new features, please open a GitHub issue or pull request.

## Methods

### Program control

- [x] [version](https://nzbget.net/api/version)
- [ ] [shutdown](https://nzbget.net/api/shutdown)
- [ ] [reload](https://nzbget.net/api/reload)

### Queue and history

- [x] [listgroups](https://nzbget.net/api/listgroups)
- [x] [listfiles](https://nzbget.net/api/listfiles)
- [x] [history](https://nzbget.net/api/history)
- [ ] [append](https://nzbget.net/api/append)
- [ ] [editqueue](https://nzbget.net/api/editqueue)
- [ ] [scan](https://nzbget.net/api/scan)

### Status, logging and statistics

- [x] [status](https://nzbget.net/api/status)
- [x] [log](https://nzbget.net/api/log)
- [ ] [writelog](https://nzbget.net/api/writelog)
- [ ] [loadlog](https://nzbget.net/api/loadlog)
- [ ] [servervolumes](https://nzbget.net/api/servervolumes)
- [ ] [resetservervolume](https://nzbget.net/api/resetservervolume)

### Pause and speed limit

- [ ] [rate](https://nzbget.net/api/rate)
- [ ] [pausedownload](https://nzbget.net/api/pausedownload)
- [ ] [resumedownload](https://nzbget.net/api/resumedownload)
- [ ] [pausepost](https://nzbget.net/api/pausepost)
- [ ] [resumepost](https://nzbget.net/api/resumepost)
- [ ] [pausescan](https://nzbget.net/api/pausescan)
- [ ] [resumescan](https://nzbget.net/api/resumescan)
- [ ] [scheduleresume](https://nzbget.net/api/scheduleresume)

### Configuration

- [ ] [config](https://nzbget.net/api/config)
- [ ] [loadconfig](https://nzbget.net/api/loadconfig)
- [ ] [saveconfig](https://nzbget.net/api/saveconfig)
- [ ] [configtemplates](https://nzbget.net/api/configtemplates)