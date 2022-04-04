package main

import (
	"fmt"
	"main/Zabbix"
)

func main() {
	hostresult := Zabbix.GetHosts()
	fmt.Println(hostresult)
}
