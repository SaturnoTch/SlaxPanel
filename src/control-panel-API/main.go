package main

// Imports necesarios
import (
	"fmt"
	"main/hwutils"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var L_MaxRam int = 64000
var L_MaxDisk int = 655

func main() {
	UsedRam := 0
	UsedDisk := 0
	go func() {
		for {
			_, UsedRam, _ = hwutils.GetRam()
			_, UsedDisk, _ = hwutils.GetDisk()
			if UsedRam >= L_MaxRam || UsedDisk >= L_MaxDisk {
				log := hwutils.SendCommand("shutdown -r -f -t 60")
				file_log, err := os.OpenFile("Limiter.json", os.O_CREATE, os.ModeAppend)
				if err != nil {
					fmt.Println("esto va en errors.json --Aun no esta listo --")
					defer file_log.Close()
					return
				}
				defer file_log.Close()
				os.WriteFile(file_log.Name(), []byte(log), 0644)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Router
	r := gin.Default()
	r.LoadHTMLFiles("./pages/dashboard.html", "./pages/index.html")
	// cargo statics
	r.Static("/static", "./pages/static")
	// Stats la ruta que muestra informacion del sistema
	// en la web
	r.GET("/home/dashboard", func(c *gin.Context) {

		info_DiskTotal, info_DiskUsed, info_DiskFree := hwutils.GetDisk()
		info_RamTotal, info_RamUsed, info_RamFree := hwutils.GetRam()
		info_CoresTotal, info_CPU := hwutils.GetCPU()

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
		hwutils.ProgramShutdown(seconds)
		c.Redirect(301, "/home/dashboard")
	})
	r.POST("/cmd", func(c *gin.Context) {
		paramCommand := c.PostForm("command")
		output := hwutils.SendCommand(paramCommand)
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
			"maxram":  L_MaxRam,
			"maxdisk": L_MaxDisk,
		})
	})

	r.POST("/home/dashboard/limiter/updatelimiter", func(c *gin.Context) {
		fmt.Println("Changes Succesfull")
		ramUpdated := c.PostForm("updateram")
		diskUpdated := c.PostForm("updatedisk")
		cnv_ramUpdated, err := strconv.Atoi(ramUpdated)
		cnv_diskUpdated, err := strconv.Atoi(diskUpdated)
		if err != nil {
			fmt.Println("Error: No se pudo convertir; ", err)
		}
		L_MaxDisk = cnv_diskUpdated
		L_MaxRam = cnv_ramUpdated
		c.Redirect(301, "/home/dashboard/limiter")
	})

	r.Run()

}
