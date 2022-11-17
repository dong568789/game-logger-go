package logSave

import (
	"fmt"
	"game-logger-go/utils"
	"os"
	"sync"
	"time"
)

type WriteModel struct {
	ModelName    string
	Path         string
	FilePath     string
	records      []string
	LastSaveTime time.Time
	writer       *os.File
	m            sync.RWMutex
}

func (w *WriteModel) Save() bool {
	if w.FilePath == "" {
		w.FilePath = w.genFilePath()
		w.resetWriter()
	}
	if w.writer == nil {
		return false
	}
	if w.Size() > 0 {
		for _, r := range w.records {
			w.writer.WriteString(r + "\r\n")
		}
	}

	w.records = nil
	w.LastSaveTime = time.Now()
	return true
}

func (w *WriteModel) Write(content string) {
	if w.FilePath == "" || w.FilePath != w.genFilePath() {
		w.FilePath = w.genFilePath()
		w.resetWriter()
	}
	w.records = append(w.records, content)
}

func (w *WriteModel) resetWriter() {
	w.m.Lock()
	defer w.m.Unlock()
	if w.writer != nil {
		w.writer = nil
	}

	f, err := os.OpenFile(w.genFilePath(), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		utils.Logger.Error("%v", err)
		return
	}
	w.writer = f
}

func (w *WriteModel) genFilePath() string {
	return fmt.Sprintf("%s/%s.log.%s", w.Path, w.ModelName, time.Now().Format("2006-01-02"))
}

func (w *WriteModel) Size() int {
	return len(w.records)
}
