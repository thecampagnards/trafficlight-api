package usb

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

/*
-n device use the USB-Switch with following serial number
0 Turn off(0 ist die number 0)
1 Turn on
R Turn on red traffic light
Y Turn on yellow traffic light
G Turn on green traffic light
O Turn off all traffic lights
-# Switch select Switch the case of multiple Switches, first=0
-i nnn interval test, Switch on and off in an endless loop, time interval nnn ms
-I nnn interval test, turn on, wait nnn ms and after this - turn off
-p t1 .. tn pulse mode, the Switch will be repeately turned on for 0,5 seconds, the time between the Switching operations will be set due to t1 to tn in seconds. At the end, the program will be stopped.
*/

type TrafficLight struct {
	driverExe string
	lastCmd   *exec.Cmd
}

func New() TrafficLight {
	var tl TrafficLight

	tl.driverExe = "drivers/linux/USBSwitchCMD"
	if runtime.GOOS == "windows" {
		tl.driverExe = "drivers\\windows\\USBSwitchCmd.exe"
	}

	log.Printf("Use driverExe: %s", tl.driverExe)
	return tl
}

/* Turn */

func (tl *TrafficLight) TurnOff() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tl.lastCmd = exec.Command(tl.driverExe, "O")
	log.Printf("TurnOff TrafficLight: %s", tl.lastCmd.String())
	return tl.lastCmd.Start()
}

func (tl *TrafficLight) TurnOnRed() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tl.lastCmd = exec.Command(tl.driverExe, "1", "-#", "0")
	log.Printf("TurnOnRed TrafficLight: %s", tl.lastCmd.String())
	err := tl.lastCmd.Start()
	/* fix idk why turn on return exit code 1 even when it's working
	$ ./drivers/windows/USBswitchCMD.exe 1 -# 2
	$ echo $?
	1
	*/
	if err != nil && strings.Contains(err.Error(), "exit code 1") {
		return nil
	}
	return err
}

func (tl *TrafficLight) TurnOnYellow() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tl.lastCmd = exec.Command(tl.driverExe, "1", "-#", "1")
	log.Printf("TurnOnYellow TrafficLight: %s", tl.lastCmd.String())
	err := tl.lastCmd.Start()
	/* fix idk why turn on return exit code 1 even when it's working
	$ ./drivers/windows/USBswitchCMD.exe 1 -# 2
	$ echo $?
	1
	*/
	if err != nil && strings.Contains(err.Error(), "exit code 1") {
		return nil
	}
	return err
}

func (tl *TrafficLight) TurnOnGreen() error {
	if tl.lastCmd != nil {
		test := tl.lastCmd.Process.Kill()
		log.Println(test.Error())
	}
	tl.lastCmd = exec.Command(tl.driverExe, "1", "-#", "2")
	log.Printf("TurnOnGreen TrafficLight: %s", tl.lastCmd.String())
	err := tl.lastCmd.Start()
	/* fix idk why turn on return exit code 1 even when it's working
	$ ./drivers/windows/USBswitchCMD.exe 1 -# 2
	$ echo $?
	1
	*/
	if err != nil && strings.Contains(err.Error(), "exit code 1") {
		return nil
	}
	return err
}

func (tl *TrafficLight) TurnOffRed() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tl.lastCmd = exec.Command(tl.driverExe, "0", "-#", "0")
	log.Printf("TurnOnRed TrafficLight: %s", tl.lastCmd.String())
	return tl.lastCmd.Start()
}

func (tl *TrafficLight) TurnOffYellow() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tl.lastCmd = exec.Command(tl.driverExe, "0", "-#", "1")
	log.Printf("TurnOnYellow TrafficLight: %s", tl.lastCmd.String())
	return tl.lastCmd.Start()
}

func (tl *TrafficLight) TurnOffGreen() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tl.lastCmd = exec.Command(tl.driverExe, "0", "-#", "2")
	log.Printf("TurnOnGreen TrafficLight: %s", tl.lastCmd.String())
	return tl.lastCmd.Start()
}

/* Switch */

func (tl *TrafficLight) SwitchRed() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tl.lastCmd = exec.Command(tl.driverExe, "R")
	log.Printf("SwitchRed TrafficLight: %s", tl.lastCmd.String())
	return tl.lastCmd.Start()
}

func (tl *TrafficLight) SwitchYellow() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tl.lastCmd = exec.Command(tl.driverExe, "Y")
	log.Printf("SwitchYellow TrafficLight: %s", tl.lastCmd.String())
	return tl.lastCmd.Start()
}

func (tl *TrafficLight) SwitchGreen() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tl.lastCmd = exec.Command(tl.driverExe, "G")
	log.Printf("SwitchGreen TrafficLight: %s", tl.lastCmd.String())
	return tl.lastCmd.Start()
}

/* Blink */

func (tl *TrafficLight) BlinkGreen() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tmp := []string{"-#", "2", "-p", "2"}
	for i := 0; i < 20; i++ {
		tmp = append(tmp, "2")
	}
	tl.lastCmd = exec.Command(tl.driverExe, tmp...)
	log.Printf("BlinkGreen TrafficLight: %s", tl.lastCmd.String())
	return tl.lastCmd.Start()
}

func (tl *TrafficLight) BlinkYellow() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tmp := []string{"-#", "1", "-p", "2"}
	for i := 0; i < 20; i++ {
		tmp = append(tmp, "2")
	}
	tl.lastCmd = exec.Command(tl.driverExe, tmp...)
	log.Printf("BlinkYellow TrafficLight: %s", tl.lastCmd.String())
	return tl.lastCmd.Start()
}

func (tl *TrafficLight) BlinkRed() error {
	if tl.lastCmd != nil {
		tl.lastCmd.Process.Kill()
	}
	tmp := []string{"-#", "0", "-p", "2"}
	for i := 0; i < 20; i++ {
		tmp = append(tmp, "2")
	}
	tl.lastCmd = exec.Command(tl.driverExe, tmp...)
	log.Printf("BlinkRed TrafficLight: %s", tl.lastCmd.String())
	return tl.lastCmd.Start()
}
