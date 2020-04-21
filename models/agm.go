package models

import (
	"fmt"
)

// ElementosAGM ...
func ElementosAGM(recorridos []map[string]interface{}) (recorridosAGM []map[string]interface{}) {
	// fmt.Println("elementos")
	for i := 0; i < len(recorridos); i++ {
		Markers := ConstruccionMarkers(recorridos[i])
		Lineas := ConstruccionLineas(Markers)
		recorridos[i]["Markers"] = Markers
		recorridos[i]["Lineas"] = Lineas
	}
	return recorridos
}

// ConstruccionMarkers ...
func ConstruccionMarkers(recorrido map[string]interface{}) (markers []map[string]interface{}) {

	arrayMarkers := make([]map[string]interface{}, 0)
	datoContruirdo := make(map[string]interface{})
	datoContruirdo = map[string]interface{}{
		"lat":       recorrido["Inicio"].(map[string]interface{})["Latitud"],
		"lng":       recorrido["Inicio"].(map[string]interface{})["Longitud"],
		"label":     fmt.Sprintf("E:%v - M:%v", recorrido["Inicio"].(map[string]interface{})["ID"], recorrido["Inicio"].(map[string]interface{})["Magnitud"]),
		"draggable": false,
	}
	// logs.Warn(datoContruirdo)
	arrayMarkers = append(arrayMarkers, datoContruirdo)
	arrayRecorrido, _ := GetElementoMaptoStringToMapArray(recorrido["Recorrido"])
	// fmt.Println(errRecorrido)
	// logs.Error(len(arrayRecorrido))
	for i := 0; i < len(arrayRecorrido); i++ {
		// logs.Warning(arrayRecorrido[i]["Latitud"])
		// datoContruirdo2 := make(map[string]interface{})

		datoContruirdo = map[string]interface{}{
			"lat":       arrayRecorrido[i]["Latitud"],
			"lng":       arrayRecorrido[i]["Longitud"],
			"label":     fmt.Sprintf("E:%v - M:%v", arrayRecorrido[i]["ID"], arrayRecorrido[i]["Magnitud"]),
			"draggable": false,
		}
		arrayMarkers = append(arrayMarkers, datoContruirdo)

	}
	return arrayMarkers
}

// ConstruccionLineas ...
func ConstruccionLineas(markers []map[string]interface{}) (lineas []map[string]interface{}) {
	arrayLineas := make([]map[string]interface{}, 0)
	datoContruirdo := make(map[string]interface{})

	for i := 1; i < len(markers); i++ {
		datoContruirdo = map[string]interface{}{
			"latIni": markers[i-1]["lat"],
			"lngIni": markers[i-1]["lng"],
			"latFin": markers[i]["lat"],
			"lngFin": markers[i]["lng"],
		}
		arrayLineas = append(arrayLineas, datoContruirdo)
	}
	return arrayLineas

}
