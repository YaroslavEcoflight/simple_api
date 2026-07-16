package config

import (
	"fmt"
	"os"
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
	godotenv.Load()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	Cfg = Setting{
		Version: os.Getenv("VERSION"),
		Host:    os.Getenv("HOST"),
		Port:    port,
	}
}
