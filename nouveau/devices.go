package nouveau

import (
	"os"
	"strconv"
	"strings"
)

func NouveauGetDRIDevices() ([]int, error) {
	devices := make([]int, 0)
	for i := 0; i < 128; i++ {
		data, err := os.ReadFile("/sys/kernel/debug/dri/" + strconv.Itoa(i) + "/name")
		if os.IsNotExist(err) {
			break
		} else if err != nil {
			return nil, err
		} else {
			if strings.Fields(strings.Split(string(data), "\n")[0])[0] == "nouveau" {
				devices = append(devices, i)
			}
		}
	}
	return devices, nil
}
