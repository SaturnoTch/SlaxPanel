package main

// Imports necesarios
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/hwutils"
	"os"
	"strconv"
)

func main() {

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
			"maxram": "0",
		})
	})

	r.POST("/home/dashboard/limiter/activate", func(c *gin.Context) {
		ramLimit := c.PostForm("limit")
		convertRamLimit, err := strconv.Atoi(ramLimit)
		if err != nil {
			fmt.Println("Error: no se pudo convertir la variable; ", err)
		}
		hwutils.SetRamLimiter(convertRamLimit)
	})
	r.Run()

}
