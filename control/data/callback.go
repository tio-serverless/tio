package data

import "time"

type CodeUpload struct {
	Message   string    `json:"Message"`
	Subject   string    `json:"Subject"`
	Timestamp time.Time `json:"Timestamp"`
}

type CodeUploadInfo struct {
	Event    string `json:"event"`
	Time     int64  `json:"time"`
	Reqid    string `json:"reqid"`
	SourceIP string `json:"source_ip"`
	Bucket   string `json:"bucket"`
	Key      string `json:"key"`
	Fsize    int    `json:"fsize"`
	Hash     string `json:"hash"`
	Ftype    int    `json:"ftype"`
	Mimetype string `json:"mimetype"`
}
