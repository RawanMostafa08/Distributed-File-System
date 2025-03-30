package models

type DataNode struct {
	NodeID          string
	IP              string
	Port            string
	IsDataNodeAlive bool
	HeartBeat       int
}

type FileData struct {
	Filename string
	FilePath string
	FileSize int64
	NodeID   string
}
