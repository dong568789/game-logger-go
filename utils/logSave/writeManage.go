package logSave

import (
	"game-logger-go/utils"
	"github.com/robfig/cron/v3"
	"os"
	"time"
)

type WriteManage struct {
	CountLimit    int
	FlushInterval time.Duration
	ModelMap      map[string]*WriteModel
	Path          string
}

const (
	CountLimit    = 10
	FlushInterval = 3 * time.Second
)

func NewWriteManage(path string) *WriteManage {
	if ok := utils.Exists(path); !ok {
		err := os.Mkdir(path, 0755)
		if err != nil {
			panic(err)
		}
	}
	w := &WriteManage{
		CountLimit:    CountLimit,
		FlushInterval: FlushInterval,
		ModelMap:      make(map[string]*WriteModel),
		Path:          path,
	}
	go w.scheduledTask()
	return w
}

func (m *WriteManage) scheduledTask() {

	local, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(local), cron.WithSeconds())
	intervalId, err := c.AddFunc("*/1 * * * * *", func() {
		now := time.Now()
		for _, model := range m.ModelMap {
			if model.Size() >= m.CountLimit || now.Sub(model.LastSaveTime) >= m.FlushInterval {
				model.Save()
			}
		}
	})

	if err != nil {
		utils.LogMgr.Error("全局定时器报错 error=%s", err)
		return
	}
	utils.LogMgr.Info("全局定时器已开启，Id=%d", intervalId)
	c.Start()

}

func (m *WriteManage) Write(modelName string, content string) {
	if _, ok := m.ModelMap[modelName]; !ok {
		m.addModel(modelName)
	}

	m.ModelMap[modelName].Write(content)
}

func (m *WriteManage) addModel(modelName string) {
	m.ModelMap[modelName] = &WriteModel{
		Path:      m.Path,
		ModelName: modelName,
	}
}
