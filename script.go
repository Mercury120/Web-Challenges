package main

import (
	"fmt"
	"time"
	"crypto/sha256"
	"strconv"
)

func main(){
	now := time.Now().Unix()
	for offset := range 10 {
		t := now - int64(offset)
		sessionID := fmt.Sprintf("%x", sha256.Sum256([]byte(strconv.FormatInt(t, 10))))
		fmt.Println(sessionID)
	}
}
