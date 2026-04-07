package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"go.bug.st/serial"
)

// App struct
type App struct {
	ctx  context.Context
	port serial.Port
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetSerialNames() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		a.sendTowebStr("ERROR listing ports: " + err.Error() + "\n")
		return []string{}
	}

	return ports
}

func (a *App) PortOpen(serialName string, baudRate int, length int) {
	a.PortClose()
	mode := &serial.Mode{
		BaudRate: baudRate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(serialName, mode)
	if err != nil {
		a.sendTowebStr("ERROR opening " + serialName + ": " + err.Error() + "\n")
		return
	}

	a.port = port
	a.sendTowebStr(fmt.Sprintf("Connected to %s @ %d baud\n", serialName, baudRate))
	buff := make([]byte, length)
	go func() { // Goroutine untuk membaca port
		for {
			n, err := port.Read(buff)
			if err != nil {
				a.sendTowebStr("ERROR reading " + serialName + ": " + err.Error() + "\n")
				break
			}
			if n == 0 {
				continue
			}

			a.sendToweb(buff[:n])

		}
	}()
}

func (a *App) PortClose() {
	defer func() {
		if r := recover(); r != nil {
			errMsg := fmt.Sprintf("%v", r)
			a.sendTowebStr(errMsg)
		}
	}()

	if a.port == nil {
		return
	}

	if err := a.port.Close(); err != nil {
		a.sendTowebStr("ERROR closing port: " + err.Error() + "\n")
	}
	a.port = nil
}
func (a *App) sendToweb(buff []byte) {
	intArray := make([]int, len(buff))
	for i, b := range buff {
		intArray[i] = int(b)
	}

	// Emit array of int
	runtime.EventsEmit(a.ctx, "printToweb", intArray)
}

func (a *App) sendTowebStr(str string) {
	byteArray := []byte(str)
	// Emit array of int
	a.sendToweb(byteArray)
}

func (a *App) IsConnected() bool {
	return a.port != nil
}

func (a *App) OpenAbout() {
	runtime.BrowserOpenURL(a.ctx, "https://github.com/nnttoo/serial_reader")
}
