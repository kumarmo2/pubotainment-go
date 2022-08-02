package servicediscovery

type ServerInstances struct {
	Id  string   `cql:"id"`
	Ips []string `cql:"ips"`
}
