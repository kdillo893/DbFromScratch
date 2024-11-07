package dbscratch

import (
	"fmt"
	"math/rand"
	"os"
)

// Save data to a file
func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)

	if err != nil {
		return err
	}
	//defer the file close until after this function...
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	return fp.Sync()
}

// Atomic write of a file over the existing file
func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	//defer discards temp file, close pointer for read old

	defer func() {
		fp.Close()
		if err != nil {
			os.Remove(tmp)
		}
	}()

	//write to tmp file
	if _, err = fp.Write(data); err != nil {
		return err
	}
	//fsync tmp file to disc
	if err = fp.Sync(); err != nil {
		return err
	}

	//replace tmp to path
	err = os.Rename(tmp, path)
	return err
}
