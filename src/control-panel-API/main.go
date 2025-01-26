package main

// Imports necesarios
import (
	"fmt"
	"os"
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

func sendCommand(command string) (out string) {
	cmd := exec.Command("cmd", "/C", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return string(output)
}

func setRamLimiter(max int) {
	_, usedRam, _ := getRam()
	for true {
		if usedRam >= max {
			cmd := exec.Command("cmd", "/C", "shutdown -r -f -t 100")
			cmd.Run()
			return
		}
	}
}
func main() {

	// Router
	r := gin.Default()
	r.LoadHTMLFiles("./pages/dashboard.html", "./pages/index.html")
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
		c.Redirect(301, "/home/dashboard")
	})
	r.POST("/cmd", func(c *gin.Context) {
		paramCommand := c.PostForm("command")
		output := sendCommand(paramCommand)
		file, err := os.OpenFile("console.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error: No se pudo abrir el archivo 'console.json'; ", err)
		}
		defer file.Close()
		file.WriteString(output)
		c.Redirect(301, "/home/dashboard")
	})

	r.GET("/home/dashboard/limiter", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"maxram": "0",
		})
	})

	r.POST("/home/dashboard/limiter/activate", func(c *gin.Context) {
		ramLimit := c.PostForm("limit")
		convertRamLimit, err := strconv.Atoi(ramLimit)
		if err != nil {
			fmt.Println("Error: no se pudo convertir la variable; ", err)
		}
		setRamLimiter(convertRamLimit)
	})
	r.Run()

}
