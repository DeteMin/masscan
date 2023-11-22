package tools

import (
	"encoding/json"
)

// ParseJson Parse takes a byte array of masscan json data and unmarshals it into a
// MasscanResult struct.
func ParseJson(content []byte) (*MasscanResult, error) {
	//fmt.Println(string(content))
	var m []Hosts

	err := json.Unmarshal(content, &m)
	if err != nil {
		return nil, err
	}

	var result MasscanResult

	var hostMap = make(map[string][]Ports)
	for i := range m {
		//fmt.Printf("%s\t", m[i].IP)
		//fmt.Printf("%d\n", m[i].Ports[0].Port)
		if _, ok := hostMap[m[i].IP]; !ok {
			hostMap[m[i].IP] = make([]Ports, 0)
		}

		hostMap[m[i].IP] = append(hostMap[m[i].IP], Ports{
			Port: m[i].Ports[0].Port,
			Proto:  m[i].Ports[0].Proto,
			Status: m[i].Ports[0].Status,
			Reason: m[i].Ports[0].Reason,
			TTL:    m[i].Ports[0].TTL,
		})
	}

	for hostIP, ports := range hostMap {
		result.Hosts = append(result.Hosts,
			Hosts{
				IP: hostIP,
				Ports: ports,
			})
	}

	return &result, nil
}

