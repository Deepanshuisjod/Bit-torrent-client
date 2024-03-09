package main

type File struct {
	announce      string
	comment       string
	creation_date int32
	httpseeds     []string
	info          Info
}

type Info struct {
	length       int64
	name         string
	piece_length int32
	pieces       []byte
}
