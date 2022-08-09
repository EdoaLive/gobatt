package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/yusufpapurcu/wmi"
	"log"
	"os/exec"
	"time"
)

type BatteryStatus struct {
	Voltage       uint32
	DischargeRate int32
}

func main() {
	var dst []BatteryStatus
	q := wmi.CreateQuery(&dst, "")

	count := 0

	for {
		err := wmi.QueryNamespace(q, &dst, "\\ROOT\\WMI")
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range dst {
			fmt.Printf("V:\t%v\tP:\t%v\n", v.Voltage, v.DischargeRate)
			if v.Voltage < 9300 {
				//beeep.Alert("Title", "Message body", "assets/information.png")
				beeep.Beep(220, beeep.DefaultDuration)
				count++
				fmt.Printf("Voltage critical, count %v...\n", count)
				if count >= 5 {
					fmt.Println("Voltage critical, hibernating...")
					cmd := exec.Command("C:\\Windows\\System32\\shutdown.exe", "/h")
					err := cmd.Run()
					if err != nil {
						fmt.Println(err)
					} else {
						return
					}
				}
			} else {
				count = 0
			}
		}
		time.Sleep(3000 * time.Millisecond)
	}
}
