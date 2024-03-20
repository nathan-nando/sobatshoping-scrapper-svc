package configs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"runtime"
	"strings"
)

func LoadConfig(mode AppMode) *Config {
	cfg := Config{}
	loc := strings.Split(getSourcePath(), "/configs")[0]
	data, err := os.ReadFile(fmt.Sprintf(FileLocationFormat, loc, mode))
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err != nil {
		logrus.Panicf("File not found %v", err)
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		logrus.Panicf("Failed to load file %v", err)
	}
	return &cfg
}

func getSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
