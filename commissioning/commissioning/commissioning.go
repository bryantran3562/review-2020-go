package commissioning

import "fmt"

type Commissioning struct {
	Username  string
	Password  string
	Ipaddress string
}

func (c *Commissioning) login() string {
	URL := fmt.Sprintf("https://%s:443", c.Ipaddress)
	return URL
}
