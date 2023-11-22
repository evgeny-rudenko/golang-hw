package main

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

const (
	pbSlleepTime = 100 // пауза для прогрес бара
	chunkSize    = 100
)

// для тестов убираем задержку при копировании и вывод прогресс бара.
var itsATest = true

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	fileInfo, err := fromFile.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	_, err = fromFile.Seek(offset, 0)
	if err != nil {
		return err
	}

	toFile, err := os.OpenFile(toPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		return err
	}
	defer toFile.Close()

	if limit > fileInfo.Size() || limit == 0 {
		limit = fileInfo.Size()
	}
	var bar *pb.ProgressBar
	buff := make([]byte, chunkSize)
	total := int64(0)

	if !itsATest {
		bar = pb.Full.Start64(100) // максимальное число в процентах
		bar.SetRefreshRate(time.Duration(100) * time.Millisecond)
		defer bar.Finish()
	}

	for {
		n, err := fromFile.Read(buff)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 || (limit > 0 && total >= limit) {
			break
		}

		if limit > 0 && total+int64(n) > limit {
			n = int(limit - total)
		}

		if !itsATest {
			time.Sleep(pbSlleepTime * time.Millisecond)
			bar.SetCurrent(int64(float64(total) / float64(limit) * 100))
		}
		_, err = toFile.Write(buff[:n])
		if err != nil {
			return err
		}

		total += int64(n)
	}
	if !itsATest {
		bar.SetCurrent(100)
	}
	return nil
}
