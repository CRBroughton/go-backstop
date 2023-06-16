package config

import (
	"os"

	"github.com/crbroughton/go-backstop/utils"
)

var path = "config.json"

func CreateJSON() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)

		if utils.IsError(err) {
			return
		}
		defer file.Close()
	}
}
