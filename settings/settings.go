package settings

import "gopkg.in/ini.v1"

func Init() (*ini.File, error) {
	conf, err := ini.Load("./conf/app.ini")
	if err != nil {
		return nil, err
	}
	return conf, nil
}
