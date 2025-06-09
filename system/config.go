package system

import "ntduncan.com/typer/utils"

type Config struct {
	TopScore float64        `json:"TopScore"`
	Mode     utils.TestMode `json:"Mode"`
	Size     int            `json:"Size"`
}
