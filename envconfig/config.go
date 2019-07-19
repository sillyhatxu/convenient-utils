package envconfig

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func ParseConfig(configFile string, unmarshalfunc func([]byte)) {
	if fileInfo, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			log.Panicf("configuration file [%v] does not exist.", configFile)
		} else {
			log.Panicf("configuration file [%v] can not be stated. %v", configFile, err)
		}
	} else {
		if fileInfo.IsDir() {
			log.Panicf("%v is a directory name", configFile)
		}
	}

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panicf("read configuration file error. %v", err)
	}
	content = bytes.TrimSpace(content)
	unmarshalfunc(content)
}
