package model

import "github.com/jinzhu/gorm"

type ExaFile struct {
	gorm.Model
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

type ExaFileChunk struct {
	gorm.Model
	ExaFileId       uint
	FileChunkNumber int
	FileChunkPath   string
}
