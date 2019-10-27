package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

func copyfile(from string, to string, offset int64, limit int64) {
	openfile, err := os.Open(from)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("file is not exist: %v", err)
		}
		log.Fatalf("error: %v", err)
	}
	defer openfile.Close()

	fi, err := openfile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if fi.Size() < offset+limit {
		if fi.Size() < offset {
			log.Fatal("error: offset > filesize")
		}
		limit = fi.Size() - offset
	}
	if limit == 0 {
		limit = fi.Size() - offset
	}
	savefile, err := os.Create(to)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer savefile.Close()

	// creating progressbar
	bar := pb.Full.Start64(limit)

	limitreader := io.LimitReader(openfile, limit)
	_, err = openfile.Seek(offset, 0)
	if err != nil {
		log.Fatal(err)
	}
	bar.Finish()
	barReader := bar.NewProxyReader(limitreader)
	_, err = io.Copy(savefile, barReader)
	if err != nil {
		log.Fatal(err)
	}
	bar.Finish()
}

func main() {
	var limit int64
	flag.Int64Var(&limit, "limit", 0, "limit read bytes")
	var offset int64
	flag.Int64Var(&offset, "offset", 0, "offset of read bytes")
	var from string
	flag.StringVar(&from, "from", "in", "filename of read file")
	var to string
	flag.StringVar(&to, "to", "out", "filename of write file")
	flag.Parse()

	copyfile(from, to, offset, limit)
}
