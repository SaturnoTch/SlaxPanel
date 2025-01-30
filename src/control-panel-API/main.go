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

func main() {
	L_MaxRam := 5500
	UsedRam := 0
	go func() {
		for {
			_, UsedRam, _ = hwutils.GetRam()
			fmt.Println(UsedRam)
			if UsedRam >= L_MaxRam {
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
			"maxram": L_MaxRam,
		})
	})

	r.PUT("/home/dashboard/limiter/updatelimiter", func(c *gin.Context) {
		var request struct {
			Updateram int `json:"updateram"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "JSON inválido"})
			return
		}

		L_MaxRam = request.Updateram
		c.JSON(200, gin.H{"message": "Límite de RAM actualizado"})
	})

	r.Run()

}
