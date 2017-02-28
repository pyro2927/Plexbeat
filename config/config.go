// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period      time.Duration `config:"period"`
  AuthToken   string        `config:"auth"`
  Host        string        `config:"host"`
  Port        string        `config:"port"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
  AuthToken: "",
  Host: "localhost",
  Port: "32400",
}
