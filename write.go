package fw

import (
	"fmt"
	"net/http"
	"os"

	"strconv"

	"github.com/moul/http2curl"
	"github.com/wtertius/pp"
)

func WriteStructByMask(root, path string, v interface{}) error {
	pp.ColoringEnabled = false
	return WriteNewFileByMask(root, path, []byte(pp.Sprint(v)))
}

func WriteCurlReqByMask(root, path string, req *http.Request) error {
	return WriteNewFileByMask(root, path, []byte(RequestToCurl(req)))
}

func RequestToCurl(req *http.Request) (curl string) {
	command, _ := http2curl.GetCurlCommand(req)
	return command.String()
}

func WriteNewFileByMask(root, path string, bb []byte) error {
	if freeFileName, err := findFileByMask(root, path); err != nil {
		fmt.Printf(">>> %s <<< debug\n", err)
		return err
	} else {
		fmt.Printf(">>> %s <<< debug\n", freeFileName)

		return Write(freeFileName, bb)
	}
}

func findFileByMask(root, path string) (string, error) {
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return "", fmt.Errorf("root directory not exists %s", root)
	}
	freeNameFound := false
	i := 0
	for !freeNameFound {
		var freeName string
		if i == 0 {
			freeName = root + "/" + path
		} else {
			freeName = root + "/" + path + strconv.Itoa(i)
		}
		if _, err := os.Stat(freeName); os.IsNotExist(err) {
			return freeName, nil
		}
		i++
	}
	return "", fmt.Errorf("unexpected error")
}

func Write(fileName string, bb []byte) error {
	fmt.Printf("fw debug >>> %s <<< debug\n", fileName)

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
