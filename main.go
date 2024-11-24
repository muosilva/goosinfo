package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	getSystemInfo()
}

func getSystemInfo() {
	hostname, _ := os.Hostname()
	osName := runtime.GOOS
	arch := runtime.GOARCH
	cpus := runtime.NumCPU()

	currentTime := time.Now().Format("Mon Jan 2 15:04:05 2006")

	fmt.Println("──────────────────────────────────")
	fmt.Printf("Host: %s\n", hostname)
	fmt.Printf("OS: %s\n", osName)
	fmt.Printf("Architecture: %s\n", arch)
	fmt.Printf("CPU cores: %d\n", cpus)
	fmt.Printf("Current time: %s\n", currentTime)
}

func getGpuInfo() {

}
