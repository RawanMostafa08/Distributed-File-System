package models

type DataNode struct {
	NodeID          string
	IP              string
	Port            []string
	IsPortBusy      []bool
	IsDataNodeAlive bool
	HeartBeat       int
}

type FileData struct {
	Filename string
	FilePath string
	FileSize int64
	NodeID   string
}
