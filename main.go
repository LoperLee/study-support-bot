package main

import (
	"fmt"

	"github.com/LoperLee/study-support-bot/config"
)

func main() {
	fmt.Println(config.LoadConfigration("config.yml"))
}
