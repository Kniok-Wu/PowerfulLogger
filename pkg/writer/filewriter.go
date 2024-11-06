package writer

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

type FileWriter lumberjack.Logger

func NewFileWriter(logPath string) *FileWriter {
	return &FileWriter{
		Filename:   logPath,
		MaxSize:    256,
		MaxBackups: 3,
		MaxAge:     32,
		Compress:   true,
	}
}

func (fw *FileWriter) SetMaxSize(maxSize int) {
	fw.MaxSize = maxSize
}

func (fw *FileWriter) SetMaxBackups(maxBackups int) {
	fw.MaxBackups = maxBackups
}

func (fw *FileWriter) SetCompress(compress bool) {
	fw.Compress = compress
}
