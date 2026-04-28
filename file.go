package main

type FileMeta struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Offset int64  `json:"offset"`
	Size   int64  `json:"size"`
}

type FilesMeta struct {
	Filename string     `json:"filename"`
	Files    []FileMeta `json:"files"`
}
