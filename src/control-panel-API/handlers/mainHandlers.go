package handlers

import (
	"fmt"
	"main/hwutils"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Func
func Cleanup(c *gin.Context) {
	var file_size int64
	// Contiene todos los archivos json de logs
	All_Logs := []string{
		"console.json",
		"errors.json",
		"Limiter.json",
	}
	// itera sobre el slice
	// para borrar y crear los archivos
	for _, removeFile := range All_Logs {
		deleted_file, err := os.Open(removeFile)
		if err != nil {
			fmt.Println("Error: No se pudo abrir el archivo; ", err)
		}
		defer deleted_file.Close()
		files_Stats, err := deleted_file.Stat()
		if err != nil {
			fmt.Println("Error: no se puede ver la informacion; ", err)
		}
		file_size += files_Stats.Size() / 1024
		os.Remove(removeFile)
		os.Create(removeFile)
	}
	c.JSON(200, gin.H{
		"cleanupmessage": fmt.Sprintf("Limpieza de tu servidor terminada, %d kb Eliminados.", file_size),
	})
}

func Cmd(c *gin.Context) {
	// Extrae el el valor command del form
	paramCommand := c.PostForm("command")
	// Usa la libreria hwutils para enviar un comando
	// al servidor
	output := hwutils.SendCommand(paramCommand)
	// Crea el log del comando
	file, err := os.OpenFile("console.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// Verifica si hay un error al abrirlo
	if err != nil {
		fmt.Println("Error: No se pudo abrir el archivo 'console.json'; ", err)
	}
	// Al finalizar cierra el archivo
	defer file.Close()
	// Escribe dentro del log el output
	file.WriteString(" ")
	file.WriteString(paramCommand)
	file.WriteString(output)
	// Redirecciona a /home/dashboard
	// con un codigo 301(Permanentemente Removido)
	c.Redirect(301, "/home/dashboard")
}

func Shutdown(c *gin.Context) {
	// extrae el valor del form
	timer := c.PostForm("timer")
	// convierte los segundos a int
	seconds, err := strconv.Atoi(timer)
	// verifica si hay un error al convertirlo
	if err != nil {
		fmt.Println("Error: No se pudo convertir timer a int; ", err)
	}
	// Programa un apagado
	hwutils.ProgramShutdown(seconds)
	// Redirecciona a /home/dashboard
	// con el codigo 301(Permanentemente Removido)
	c.Redirect(301, "/home/dashboard")
}



