package config

import (
	"fmt"
	"strconv"

	"github.com/joho/godotenv"
)

type Setting struct {
	Version string
	Host    string
	Port    int
}

func (s *Setting) Dsn() string {
	return "book.db"
}

func (s *Setting) Url() string {
	return fmt.Sprintf("http://%s:%d", s.Host, s.Port)
}

func (s *Setting) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

var Cfg Setting

func init() {
	vars, err := godotenv.Read()
	if err != nil {
		panic(err)
	}
	port, _ := strconv.Atoi(vars["PORT"])
	Cfg = Setting{
		Version: vars["VERSION"],
		Host:    vars["HOST"],
		Port:    port,
	}
}
