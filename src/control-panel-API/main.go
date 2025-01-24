package main

// Imports necesarios
import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func getRam() (ramTotal int, ramUsed int, ramFree int) {
	// crea 2 variables una de la VM y otro del error
	vm, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error: No se pudo obtener el valor de memoria ram")
		return
	}
	return int(vm.Total / 1024 / 1024), int(vm.Used / 1024 / 1024), int(vm.Free / 1024 / 1024)
}

func getCPU() (cpuCores int, cpuUsed int) {
	vmCores, err := cpu.Counts(false)
	vmUsage, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println("Error: No se pudo obtener el valor de la CPU")
		return
	}
	return vmCores, int(vmUsage[0])
}

func getDisk() (diskTotal int, diskUsed int, diskFree int) {
	vm, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Error: No se pudo obtener el valor del disco")
		return
	}
	return int(vm.Total / 1024 / 1024 / 1024), int(vm.Used / 1024 / 1024 / 1024), int(vm.Free / 1024 / 1024 / 1024)
}

func programShutdown(seconds int) {
	timer := strconv.Itoa(seconds)
	cmd := exec.Command("cmd", "/C", "shutdown", "/r", "/f", "/t", timer)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: No se pudo ejecutar el comando de apagado; ", err)
	}
}

func sendCommand(command string) {
	cmd := exec.Command("cmd", "/C", command)
	cmd.Run()
}

func main() {

	// Router
	r := gin.Default()
	r.LoadHTMLFiles("./pages/dashboard.html")
	// cargo statics
	r.Static("/static", "./pages/static")
	// Stats la ruta que muestra informacion del sistema
	// en la web
	r.GET("/home/dashboard", func(c *gin.Context) {

		info_DiskTotal, info_DiskUsed, info_DiskFree := getDisk()
		info_RamTotal, info_RamUsed, info_RamFree := getRam()
		info_CoresTotal, info_CPU := getCPU()

		c.HTML(200, "dashboard.html", gin.H{
			// Disk
			"diskTotal": info_DiskTotal,
			"diskUsed":  info_DiskUsed,
			"diskFree":  info_DiskFree,
			// CPU
			"coresTotal": info_CoresTotal,
			"cpu":        info_CPU,
			// RAM
			"ramTotal": info_RamTotal,
			"ramUsed":  info_RamUsed,
			"ramFree":  info_RamFree,
		})
	})
	r.POST("/shutdown", func(c *gin.Context) {
		timer := c.PostForm("timer")
		seconds, err := strconv.Atoi(timer)
		if err != nil {
			fmt.Println("Error: No se pudo convertir timer a int; ", err)
		}
		programShutdown(seconds)
	})
	r.POST("/cmd", func(c *gin.Context) {
		paramCommand := c.PostForm("command")
		sendCommand(paramCommand)
	})
	r.Run()
}
