package dummy

//import (
//	"fmt"
//	"log"
//)

type Config struct {
	P_setting string  `mapstructure:"p_setting"`
}


func (c *Config) Client() string {
	return c.P_setting
}
