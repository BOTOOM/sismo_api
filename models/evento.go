package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/logs"
)

// Clasificacion ...
func Clasificacion() {
	Eventos := LeerJSON()
	Matriz := MatrizCordenadas()
	ClasificacionEventos(Eventos, Matriz)
}

// LeerJSON ...
func LeerJSON() (dato []map[string]interface{}) {
	logs.Info("entro al leer")
	array := make([]map[string]interface{}, 0)
	// fmt.Println(array)
	raw, err := ioutil.ReadFile("./eventos1.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &array)
	logs.Info(array[0])
	return array
}

// MatrizCordenadas ...
func MatrizCordenadas() (MatrizCuadrantes []map[string]interface{}) {
	longitudBase := -90.0000
	latitudBase := 20.0000
	latitudIni := latitudBase

	matriz := make([]map[string]interface{}, 0)
	cuadrante := 0

	// logs.Info(array1)

	for i := 0; i < 8; i++ {
		fila := make([]map[string]interface{}, 0)
		datoFila := make(map[string]interface{})
		latitudFin := latitudIni - 9.9999
		longitudIni := longitudBase
		for j := 0; j < 5; j++ {
			longitudFin := longitudIni + 9.9999

			cuadrante = cuadrante + 1
			datoContruirdo := make(map[string]interface{})
			datoContruirdo = map[string]interface{}{
				// "Activo":    true,
				"fila":         j,
				"columna":      i,
				"cuadrante":    cuadrante,
				"longitud_ini": longitudIni,
				"longitud_fin": longitudFin,
				"latitud_ini":  latitudIni,
				"latitud_fin":  latitudFin,
			}

			fila = append(fila, datoContruirdo)
			longitudIni = longitudFin + 0.0001
		}
		datoFila = map[string]interface{}{
			"Fila": fila,
		}
		matriz = append(matriz, datoFila)
		latitudIni = latitudFin - 0.0001
	}

	// logs.Info(matriz[0])
	// logs.Info(matriz[len(matriz)-1])
	return matriz

}

// ClasificacionEventos ...
func ClasificacionEventos(eventos []map[string]interface{}, matriz []map[string]interface{}) {
	// if eventos[0]["Latitud"].(float64) < -5.0001 {
	// 	logs.Info(eventos[0]["Latitud"].(float64))

	// 	// logs.Info(matriz[0]["Fila"])

	// }
	// if  {
	fmt.Println(-81.2487 < -40)
	// }
	PrimerDatoMatriz, errPrimer := GetElementoMaptoStringToMapArray(matriz[0]["Fila"])
	fmt.Println(errPrimer)
	LatitudInicial := PrimerDatoMatriz[0]["latitud_ini"].(float64)
	LongitudInicial := PrimerDatoMatriz[0]["longitud_ini"].(float64)
	fmt.Println(LongitudInicial)
	UltimoDatoMatriz, errUltimo := GetElementoMaptoStringToMapArray(matriz[len(matriz)-1]["Fila"])
	fmt.Println(errUltimo)
	LatitudFinal := UltimoDatoMatriz[len(UltimoDatoMatriz)-1]["latitud_fin"].(float64)
	LongitudFinal := UltimoDatoMatriz[len(UltimoDatoMatriz)-1]["longitud_fin"].(float64)
	fmt.Println(LongitudFinal)
	logs.Info(len(eventos))
	cont := 0
	for i := 0; i < len(eventos)/10; i++ {
		// restriccion a matriz por latitud
		if (eventos[i]["Latitud"].(float64) <= LatitudInicial) && (eventos[0]["Latitud"].(float64) >= LatitudFinal) {
			// logs.Error(eventos[i])
			// (eventos[i]["Longitud"].(float64) >= LongitudInicial)
			// && (eventos[0]["Longitud"].(float64) <= LongitudFinal)
			if eventos[i]["Longitud"].(float64) >= LongitudInicial {
				if eventos[i]["Longitud"].(float64) <= LongitudFinal {
					logs.Error(eventos[i])
					cont++
				}
			}
		}
	}
	fmt.Println(cont)

}
