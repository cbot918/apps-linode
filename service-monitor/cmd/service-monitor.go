package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

const interval = 20

type SM struct {
	C *resty.Client
}

func NewSM() *SM {
	return &SM{
		C: resty.New(),
	}
}

func (sm *SM) checkLinode(interval int32) error {

	for {
		resp, err := sm.C.R().
			EnableTrace().
			Get("https://getsub.fiveplanet.online/ping")
		if err != nil {
			return err
		}
		fmt.Println(resp)
		fmt.Println(resp.StatusCode())
		break
		time.Sleep(30 * time.Second)
	}
	return nil
}

func main() {
	sm := NewSM()
	if err := sm.checkLinode(30); err != nil {
		log.Fatal(err)
	}
}
