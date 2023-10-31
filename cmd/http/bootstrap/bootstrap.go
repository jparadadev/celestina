package bootstrap

import (
	"celestina/internal/platform/server"
	"fmt"
	"reflect"
)

const (
	host = "localhost"
	port = 8080
)

func Run(configData map[string]interface{}, subscriptions map[string][]string) error {
	host := configData["server.host"].(string)
	port := configData["server.port"].(int)
	srv := server.New(host, port, subscriptions)
	return srv.Run()
}

func transformMap(inputMap map[string]interface{}) (map[string][]string, error) {
	transformedMap := make(map[string][]string)

	for key, value := range inputMap {
		// Verificar si el valor es un slice
		valRef := reflect.ValueOf(value)
		if valRef.Kind() != reflect.Slice {
			return nil, fmt.Errorf("el valor de la clave '%s' no es una lista", key)
		}

		// Iterar sobre los elementos del slice y agregarlos a la lista de strings
		var stringSlice []string
		for i := 0; i < valRef.Len(); i++ {
			element := valRef.Index(i).Interface()

			// Asegurar que el elemento sea un string
			str, ok := element.(string)
			if !ok {
				return nil, fmt.Errorf("el elemento en la clave '%s' no es un string", key)
			}

			stringSlice = append(stringSlice, str)
		}

		transformedMap[key] = stringSlice
	}

	return transformedMap, nil
}
