package types

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
)

type Process struct {
	Pid     int32                   `json:"pid"`
	Name    string                  `json:"name"`
	Cmdline string                  `json:"cmdline"`
	Cpu     float64                 `json:"cpu"`
	Mem     *process.MemoryInfoStat `json:"mem"`
}

type ProcessResponse struct {
	Processes []*Process `json:"processes"`
}

type CpuResponse struct {
	CpuInfos []cpu.InfoStat `json:"cpu_infos"`
}
