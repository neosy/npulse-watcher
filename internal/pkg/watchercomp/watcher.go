package watchercomp

import (
	"fmt"
	"log"
	"os"
	"time"

	"git.n-hub.ru/neosy/npulse-watcher/internal/pkg/ntelegram"
)

type WatcherComp struct {
	fileLogNameBase   string
	fileLogFolderPath string
	fileLogSplitYear  bool
	tgToken           string
	tgChatId          string

	compIPs map[string]compState
}

type compState struct {
	ip       string
	name     string
	lastTime time.Time
	status   CompWatchStatus
	needInfo bool
}

type CompWatchStatus int

const (
	WatchStatusOK CompWatchStatus = iota
	WatchStatusFail
)

func (status CompWatchStatus) String() string {
	return [...]string{"OK", "Fail"}[status]
}

func New(fileLogName string, folderPath string, splitYear bool) *WatcherComp {
	wc := WatcherComp{
		fileLogNameBase:   fileLogName,
		fileLogFolderPath: folderPath,
		fileLogSplitYear:  splitYear,
	}

	wc.fileLogFolderPath, _ = homeFolderSignUpdate(wc.fileLogFolderPath)
	wc.compIPs, _ = getCompIPsFromFile(wc.FileLogPath())

	return &wc
}

func (wc *WatcherComp) TelegramConfigSet(token string, chatId string) {
	wc.tgToken = token
	wc.tgChatId = chatId
}

func (wc *WatcherComp) fileLogName() string {
	fileName := wc.fileLogNameBase

	if wc.fileLogSplitYear {
		fileName = fileNameAddYear(fileName)
	}

	return fileName
}

func (wc *WatcherComp) FileLogPath() string {
	return wc.fileLogFolderPath + "/" + wc.fileLogName()
}

func (wc *WatcherComp) Add(ip string, name string, time time.Time, status CompWatchStatus) (compState, error) {
	comp, exists := wc.compIPs[ip]

	if !exists {
		comp = compState{}
		comp.needInfo = status == WatchStatusFail
	} else {
		if comp.status == status {
			comp.lastTime = time
			wc.compIPs[ip] = comp
			return comp, nil
		}

		comp.needInfo = comp.needInfo || comp.status != status
	}

	comp.ip = ip
	comp.name = name
	comp.lastTime = time
	comp.status = status

	wc.compIPs[ip] = comp

	err := writeToFile(wc.FileLogPath(), comp.ip, comp.name, comp.lastTime, comp.status)

	return comp, err
}

func (wc *WatcherComp) Check(durationFailSeconds int) {
	const ch_br = "\n"
	var msg string

	for key, value := range wc.compIPs {
		if value.status == WatchStatusOK {
			timeNow := time.Now()
			if timeNow.Sub(value.lastTime).Seconds() > float64(durationFailSeconds) {
				value, _ = wc.Add(value.ip, value.name, time.Now(), WatchStatusFail)
			}
		}

		if value.needInfo {
			if msg == "" {
				hostName, _ := os.Hostname()
				msg = fmt.Sprintf("Server name '%s'", hostName)
				msg = msg + ch_br + "Check servers:"
			}

			msg = msg + ch_br + fmt.Sprintf("    '%s' %s", value.name, value.ip)
			switch value.status {
			case WatchStatusOK:
				msg += " - ✓"
			case WatchStatusFail:
				msg += " - ✗"
			}

			value.needInfo = false
			wc.compIPs[key] = value
		}
	}

	if msg != "" {
		log.Println(msg)

		tg := ntelegram.New(wc.tgToken)
		tg.Send(wc.tgChatId, msg)
	}
}
