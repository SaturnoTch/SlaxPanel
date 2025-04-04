package main

// Imports necesarios
import (
	"crypto/rand"
	"fmt"
	"main/handlers"
	"main/hwutils"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var L_MaxRam int = 64000
var L_MaxDisk int = 655

var ID string = "73628"

func main() {
	timeNow := time.Now()
	r := gin.Default()
	/* Start Menu console
	Este menu se mostrara al inicio del servidor
	*/
	fmt.Println("--/Estos datos seran usados en /login para poder dar acceso total--/ \n Ingrese el Username")
	var Username_use string
	_, err := fmt.Scanln(&Username_use)
	if err != nil {
		fmt.Println("Error: No se pudo escanear el username")
	}
	fmt.Println("Indica tu contraseña")
	var Password_use string
	_, err = fmt.Scanln(&Password_use)
	if err != nil {
		fmt.Println("Error: No se pudo escanear la contraseña")
	}
	fmt.Println("Inicia sesion en: localhost:8080/login \n Precione cualquier tecla para continuar")
	fmt.Scan()
	UsedRam := 0
	UsedDisk := 0
	go func() {
		for {
			_, UsedRam, _ = hwutils.GetRam()
			_, UsedDisk, _ = hwutils.GetDisk()
			if UsedRam >= L_MaxRam {
				log := hwutils.SendCommand("shutdown -r -f -t 60")
				file_log, err := os.OpenFile("Limiter.json", os.O_CREATE, os.ModeAppend)
				if err != nil {
					fmt.Println("esto va en errors.json --Aun no esta listo --")
					return
				}
				defer file_log.Close()
				os.WriteFile(file_log.Name(), []byte(log), 0644)
			} else if UsedDisk >= L_MaxDisk {
				handlers.Cleanup(nil)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	// Router

	r.LoadHTMLFiles("./pages/dashboard.html", "./pages/limiter.html", "./pages/tools.html", "./pages/login.html")
	// cargo statics
	r.Static("/static", "./pages/static")
	// Stats la ruta que muestra informacion del sistema
	// en la web
	r.GET("/home/dashboard", func(c *gin.Context) {
		_, err := c.Cookie(ID)
		if err != nil {
			c.Redirect(301, "/login")
		}
		cache_ConsoleLog, err := os.ReadFile("./console.json")
		if err != nil {
			fmt.Println("Error: No se pudo leer el archivo; ", err)
		}
		cnv_ConsoleLog := string(cache_ConsoleLog)
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
			// Console
			"consolelog": cnv_ConsoleLog,
		})
	})

	// Programa un apagado
	r.POST("/shutdown", handlers.Shutdown)
	// Envia los comandos a la consola del
	// Servidor
	r.POST("/cmd", handlers.Cmd)

	r.GET("/home/dashboard/limiter", func(c *gin.Context) {
		fmt.Println(ID)
		_, err := c.Cookie(ID)
		if err != nil {
			c.Redirect(301, "/login")
		}
		c.HTML(200, "limiter.html", gin.H{
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
			return
		}
		filelog_admin, err := os.OpenFile("./admin/admin.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Error: No se pudo abrir el archivo json 'admin.json'; ", err)
			return
		}
		defer filelog_admin.Close()
		getip := c.ClientIP()
		filelog_admin.WriteString(fmt.Sprint(timeNow, "; Configuracion del limiter actualizada; ClientIP:", getip))
		L_MaxDisk = cnv_diskUpdated
		L_MaxRam = cnv_ramUpdated
		c.Redirect(301, "/home/dashboard/limiter")
	})

	r.GET("/home/dashboard/tools", func(c *gin.Context) {
		c.HTML(200, "tools.html", nil)
	})
	r.DELETE("/home/dashboard/tools/cleanup", handlers.Cleanup)

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.POST("/admin", func(c *gin.Context) {
		recover_username := c.PostForm("user")
		recover_password := c.PostForm("password")
		fileAdm, err := os.OpenFile("./admin/admin.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Error: No se puede abrir el archivo; ", err)
			c.Redirect(200, "/login")
		}
		defer fileAdm.Close()
		if recover_username == Username_use && recover_password == Password_use {
			randomNumber, err := rand.Int(rand.Reader, big.NewInt(1000000))
			if err != nil {
				fmt.Println("Error: No se pudo generar un número aleatorio; ", err)
				c.Redirect(301, "/login")
				return
			}
			ID = randomNumber.String()
			getip := c.ClientIP()
			c.SetCookie(ID, "", 900, "/", "", false, true)
			fileAdm.WriteString(fmt.Sprint(timeNow, "Warning: Cookie de entrada al panel concedido; ClientIP:", getip, "\n"))
			c.Redirect(301, "/home/dashboard")
		} else {
			filelog_admin, err := os.OpenFile("./admin/admin.json", os.O_APPEND|os.O_CREATE, os.ModeAppend)
			if err != nil {
				fmt.Println("Error: No se pudo abrir el archivo json 'admin.json'; ", err)
				return
			}
			defer filelog_admin.Close()
			getip := c.ClientIP()
			filelog_admin.WriteString(fmt.Sprint(timeNow, "Warning: Contraseña incorrecte; ClientIP:", getip, "\n"))
			c.Redirect(301, "/login")
		}
	})
	r.Run()
}
