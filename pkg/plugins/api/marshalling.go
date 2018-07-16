package api

import (
	"log"

	pluginTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"

	"github.com/vmihailenco/msgpack"
)

func marshallConfig(config pluginTypes.Config) (blob []byte) {
	blob, err := msgpack.Marshal(&config)
	if err != nil {
		log.Fatal("Config marshalling err:", err)
	}
	return
}

func marshallResource(resource kombustionTypes.TemplateObject, errs []error) (blob []byte) {
	result := pluginTypes.PluginResult{
		Data:   resource,
		Errors: errs,
	}
	blob, err := msgpack.Marshal(&result)
	if err != nil {
		log.Fatal("Resource marshalling err:", err)
	}
	return
}

func marshallMapping(mapping kombustionTypes.TemplateObject, errs []error) (blob []byte) {
	result := pluginTypes.PluginResult{
		Data:   mapping,
		Errors: errs,
	}
	blob, err := msgpack.Marshal(&result)
	if err != nil {
		log.Fatal("Mapping marshalling err:", err)
	}
	return
}

func marshallOutput(output kombustionTypes.TemplateObject, errs []error) (blob []byte) {
	result := pluginTypes.PluginResult{
		Data:   output,
		Errors: errs,
	}
	blob, err := msgpack.Marshal(&result)
	if err != nil {
		log.Fatal("Output marshalling err:", err)
	}
	return
}
