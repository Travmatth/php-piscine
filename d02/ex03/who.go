package main

// #include <utmpx.h>
// #include <string.h>
import "C"

import (
	"fmt"
	"time"
)

func fromCString(str *C.char, size C.size_t) string {
	l := C.strnlen(str, size)
	if l == 0 {
		return ""
	}
	return C.GoStringN(str, C.int(l))
}

func main() {
	C.setutxent()
	for {
		centry := C.getutxent()
		if centry == nil {
			C.endutxent()
			break
		}
		user := fromCString(&centry.ut_user[0], C._UTX_USERSIZE)
		id := fromCString(&centry.ut_line[0], C._UTX_LINESIZE)
		if user == "" || id == "" {
			continue
		}
		sec := int64(centry.ut_tv.tv_sec)
		nsec := int64(centry.ut_tv.tv_usec) * 1000
		t := time.Unix(sec, nsec)
		timestamp := fmt.Sprintf("%s %02d %02d:%02d",
			t.Month().String()[0:3],
			t.Day(),
			t.Hour(),
			t.Minute(),
		)
		fmt.Printf("%-10s %-10s %s\n", user, id, timestamp)
	}
}
