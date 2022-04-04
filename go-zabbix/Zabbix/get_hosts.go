package Zabbix

type HostsResult struct {
	Hostid int
	Name   string
	Ip     string
	Status int
}

func GetHosts() (result []HostsResult) {
	result = make([]HostsResult, 0)
	sql := "select a.hostid , a.name, b.ip, a.status  from hosts as a left join (select hostid, ip from interface GROUP BY hostid ) as b  on a.hostid = b.hostid where  a.\n\tstatus in (0,1,6) and a.flags = 0"
	DBCONNECT.Raw(sql).Scan(&result)
	return result
}
