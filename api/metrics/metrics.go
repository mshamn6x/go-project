package metrics

import (
	"log"
	"time"

	"new/test/project/api/db"
	"new/test/project/api/model"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func PersistCPUPercentages() {

	CPUManger := db.NewCPUManger()

	for {
		time.Sleep(30 * time.Second)

		percentages, err := cpu.Percent(0, false)
		if err != nil {
			log.Println("Error in getting CPU percentage")
			continue
		}

		CPUManger.Insert(&model.CPU{
			CurrentPercent: float32(percentages[0]),
			TimeStamp:      time.Now(),
		})

	}
}

func PersistMemoryUsages() {

	RamManager := db.NewRamManger()

	for {
		time.Sleep(30 * time.Second)

		memstat, err := mem.VirtualMemory()
		if err != nil {
			log.Println("Error in getting RAM percentage")
			continue
		}

		RamManager.Insert(&model.RAM{
			MemoryUsage: float32(memstat.UsedPercent),
			TimeStamp:   time.Now(),
		})

	}
}
