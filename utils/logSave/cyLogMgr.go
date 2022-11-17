package logSave

var CyMgr *CyLog

type CyLog struct {
	writeManage *WriteManage
}

func init() {
	CyMgr = new(CyLog)
}

func InitCyLog(dir string) {
	CyMgr.writeManage = NewWriteManage(dir)
}

func (c *CyLog) WriteLog(model string, log string) {
	c.writeManage.Write(model, log)
}
