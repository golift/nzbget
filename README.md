# NZBGet Go Module

[![GoDoc](https://godoc.org/golift.io/nzbget/svc?status.svg)](https://pkg.go.dev/golift.io/nzbget)
[![Go Report Card](https://goreportcard.com/badge/golift.io/nzbget)](https://goreportcard.com/report/golift.io/nzbget)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://github.com/golift/nzbget/blob/main/LICENSE)
[![discord](https://badgen.net/badge/icon/Discord?color=0011ff&label&icon=https://simpleicons.now.sh/discord/eee "GoLift Discord")](https://golift.io/discord)

Go Library to interact with NZBGet. Almost full-featured. See exposed [methods](#methods) below.

If you'd like new features, please open a GitHub issue or pull request.

## Example

Simple example to print some logs.

```golang
package main

import (
	"fmt"

	"golift.io/nzbget"
)

func main() {
	nzb := nzbget.New(&nzbget.Config{
		URL:  "http://nzbget.server.io:6789/",
		User: "userName",
		Pass: "passWord",
	})

	events, err := nzb.Log(0, 100)
	if err != nil {
		panic(err)
	}

	for _, event := range events {
		fmt.Println(event.ID, event.Kind, event.Time, event.Text)
	}
}
// Output examples:
// 47 INFO 2022-06-27 01:42:19 -0700 PDT Renaming 0b710bf619488ca0a1b5f83f53fde577.15 to eQ7Aq0DBEhHGCgSXy3PZ.part16.rar
// 103 INFO 2022-06-27 01:42:23 -0700 PDT Unrar: Extracting from eQ7Aq0DBEhHGCgSXy3PZ.part28.rar
// 104 INFO 2022-06-27 01:42:24 -0700 PDT Unrar: All OK
// 105 INFO 2022-06-27 01:42:24 -0700 PDT Deleting archive files
// 106 INFO 2022-06-27 01:42:24 -0700 PDT Deleting file eQ7Aq0DBEhHGCgSXy3PZ.part21.rar
```

## Methods

Official NZBGet API reference can be [found here](https://nzbget.net/api/).

### Program Control

- [x] [version](https://nzbget.net/api/version)
- [x] [shutdown](https://nzbget.net/api/shutdown)
- [x] [reload](https://nzbget.net/api/reload)

### Queue and History

- [x] [listgroups](https://nzbget.net/api/listgroups)
- [x] [listfiles](https://nzbget.net/api/listfiles)
- [x] [history](https://nzbget.net/api/history)
- [ ] [append](https://nzbget.net/api/append) (_needed_)
- [ ] [editqueue](https://nzbget.net/api/editqueue)
- [x] [scan](https://nzbget.net/api/scan)

### Status, Logging and Statistics

- [x] [status](https://nzbget.net/api/status)
- [x] [log](https://nzbget.net/api/log)
- [x] [writelog](https://nzbget.net/api/writelog)
- [x] [loadlog](https://nzbget.net/api/loadlog)
- [ ] [servervolumes](https://nzbget.net/api/servervolumes)
- [ ] [resetservervolume](https://nzbget.net/api/resetservervolume)

### Pause and Speed Limit

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