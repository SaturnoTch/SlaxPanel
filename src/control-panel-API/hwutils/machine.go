package hwutils

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"os/exec"
	"strconv"
)

func GetRam() (ramTotal int, ramUsed int, ramFree int) {
	// crea 2 variables una de la VM y otro del error
	vm, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error: No se pudo obtener el valor de memoria ram")
		return
	}
	return int(vm.Total / 1024 / 1024), int(vm.Used / 1024 / 1024), int(vm.Free / 1024 / 1024)
}

func GetCPU() (cpuCores int, cpuUsed int) {
	vmCores, err := cpu.Counts(false)
	vmUsage, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println("Error: No se pudo obtener el valor de la CPU")
		return
	}
	return vmCores, int(vmUsage[0])
}

func GetDisk() (diskTotal int, diskUsed int, diskFree int) {
	vm, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Error: No se pudo obtener el valor del disco")
		return
	}
	return int(vm.Total / 1024 / 1024 / 1024), int(vm.Used / 1024 / 1024 / 1024), int(vm.Free / 1024 / 1024 / 1024)
}

func ProgramShutdown(seconds int) {
	timer := strconv.Itoa(seconds)
	cmd := exec.Command("cmd", "/C", "shutdown", "/r", "/f", "/t", timer)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: No se pudo ejecutar el comando de apagado; ", err)
	}
}

func SendCommand(command string) (out string) {
	cmd := exec.Command("cmd", "/C", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return string(output)
}

func SetRamLimiter(max int) {
	_, usedRam, _ := GetRam()
	for true {
		if usedRam >= max {
			cmd := exec.Command("cmd", "/C", "shutdown -r -f -t 100")
			cmd.Run()
			return
		}
	}
}
