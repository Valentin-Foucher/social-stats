package commonutils

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func EnrichError(err error) error {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return err
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	function := funcname[strings.LastIndex(funcname, ".")+1:]
	return fmt.Errorf("%s - %s - %s", filename, function, err.Error())
}
