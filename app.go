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
	ports, _ := serial.GetPortsList()

	return ports
}

func (a *App) PortOpen(serialName string, baudRate int, length int) {
	a.PortClose()
	mode := &serial.Mode{
		BaudRate: baudRate,
	}
	port, err := serial.Open(serialName, mode)
	if err != nil {
		runtime.EventsEmit(a.ctx, "printToweb", "ERROR : "+err.Error())
		return
	}

	a.port = port
	buff := make([]byte, length)
	go func() { // Goroutine untuk membaca port
		for {
			n, err := port.Read(buff)
			if err != nil {
				a.sendTowebStr((err.Error()))
				break
			}
			if n == 0 {
				fmt.Println("\nEOF")
				break
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

	a.port.Close()
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
