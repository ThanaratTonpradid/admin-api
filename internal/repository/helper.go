package repository

import (
	"fmt"
	"time"
)

func GetUnixTimestamp() uint32 {
	unixTime := time.Now().Unix()
	return uint32(unixTime)
}

func GetKeyStaffSession(staffId uint32) string {
	return fmt.Sprintf("%s:%d", KeyStaffSession, staffId)
}
