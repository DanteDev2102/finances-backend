package enviroment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port   string
	HostDB string
	PassDB string
	UserDB string
	NameDB string
	PortDB string
}

func Enviroments() Env {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		return Env{}
	}

	return Env{
		Port:   os.Getenv("PORT_API"),
		HostDB: os.Getenv("HOST_DB"),
		UserDB: os.Getenv("USER_DB"),
		PassDB: os.Getenv("PASS_DB"),
		NameDB: os.Getenv("NAME_DB"),
		PortDB: os.Getenv("PORT_DB"),
	}
}
