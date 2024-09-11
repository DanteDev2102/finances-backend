package enviroment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type TEnv struct {
	Port   string
	HostDB string
	PassDB string
	UserDB string
	NameDB string
	PortDB string
}

var err error = godotenv.Load()

var env TEnv

var Env *TEnv = &env

func init() {
	if err != nil {
		fmt.Println(err)
		env = TEnv{}
	}

	env = TEnv{
		Port:   os.Getenv("PORT_API"),
		HostDB: os.Getenv("HOST_DB"),
		UserDB: os.Getenv("USER_DB"),
		PassDB: os.Getenv("PASS_DB"),
		NameDB: os.Getenv("NAME_DB"),
		PortDB: os.Getenv("PORT_DB"),
	}
}
