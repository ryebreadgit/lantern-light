package main

import (
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"github.com/getlantern/systray"
	"github.com/micmonay/keybd_event"
)

func main() {
	systray.Run(onReady, onExit)
}

func pressKey() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_F19)
	// Press the selected keys
	err = kb.Launching()
	if err != nil {
		panic(err)
	}
}

func setTime(item *systray.MenuItem, time int) {
	item.SetTitle(strconv.Itoa(time) + " seconds")
}

func onReady() {
	systray.SetTitle("Lantern Light")
	systray.SetIcon(iconData)
	waitstat := systray.AddMenuItem("Time Waiting", "Waiting x seconds before next button press")
	systray.AddSeparator()
	qmenu := systray.AddMenuItem("Quit", "Quit the app")
	waitstat.Disabled()

	go func() {
		for {
			pressKey()
			rand.Seed(time.Now().UnixNano())
			wait := 60 + rand.Intn(240-60)
			setTime(waitstat, wait)
			time.Sleep(time.Duration(wait) * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-qmenu.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// clean up here
}
