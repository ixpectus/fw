package fw

import (
	"os"

	"github.com/wtertius/pp"
)

func Write(fileName string, bb []byte) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	f.Write(bb)
	f.Write([]byte("\n"))
	defer f.Close()
	return nil
}

func WriteStruct(fileName string, v interface{}) error {
	pp.ColoringEnabled = false
	return Write(fileName, []byte(pp.Sprint(v)))
}
