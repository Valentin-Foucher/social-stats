package common

import (
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

const envPrefix = "---ENV "

func ReadConfiguration[C any](filepath string) (*C, error) {
	conf := new(C)
	def, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(def, conf)
	if err != nil {
		return nil, err
	}

	return setEnvVariables(conf), nil
}

func setEnvVariables[C any](conf C) C {
	confValue := reflect.ValueOf(conf)
	confType := reflect.TypeOf(conf)

	if confType.Kind() == reflect.Ptr {
		confValue = confValue.Elem()
		confType = confType.Elem()
	}

	if confType.Kind() != reflect.Struct {

		if confType.Kind() == reflect.String {
			envVariable := confValue.String()

			if strings.HasPrefix(envVariable, envPrefix) {
				conf = reflect.ValueOf(os.Getenv(strings.Replace(envVariable, envPrefix, "", 1))).Interface().(C)
			}
		}

		return conf
	}

	for i := 0; i < confType.NumField(); i++ {
		fieldValue := confValue.Field(i)
		fieldValue.Set(reflect.ValueOf(setEnvVariables(fieldValue.Interface())))
	}

	return conf
}
