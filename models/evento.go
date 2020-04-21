package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/logs"
)

// Clasificacion ...
func Clasificacion() (resultadoRecorrido []map[string]interface{}) {
	Eventos := LeerJSON()
	Matriz := MatrizCordenadas()
	EventosClasificados := ClasificacionEventos(Eventos, Matriz)
	// logs.Warn(EventosClasificados)
	// recorrido := BusquedaAdyacentes(EventosClasificados, Matriz)
	recorrido := ElementosAGM(BusquedaAdyacentes(EventosClasificados, Matriz))

	return recorrido
}

// ClasificacionV2 ...
func ClasificacionV2(datosEventos []map[string]interface{}) (resultadoRecorrido []map[string]interface{}) {
	Eventos := datosEventos
	Matriz := MatrizCordenadas()
	EventosClasificados := ClasificacionEventos(Eventos, Matriz)
	// logs.Warn(EventosClasificados)
	recorrido := BusquedaAdyacentes(EventosClasificados, Matriz)
	ElementosAGM(recorrido)
	return recorrido
}

// LeerJSON ...
func LeerJSON() (dato []map[string]interface{}) {
	// logs.Info("entro al leer")
	array := make([]map[string]interface{}, 0)
	// fmt.Println(array)
	raw, err := ioutil.ReadFile("./eventos1.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &array)
	// logs.Info(array[0])
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
func ClasificacionEventos(eventos []map[string]interface{}, matriz []map[string]interface{}) (eventosCLasificados []map[string]interface{}) {
	arrayEventos := make([]map[string]interface{}, 0)
	PrimerDatoMatriz, errPrimer := GetElementoMaptoStringToMapArray(matriz[0]["Fila"])
	fmt.Println(errPrimer)
	LatitudInicial := PrimerDatoMatriz[0]["latitud_ini"].(float64)
	LongitudInicial := PrimerDatoMatriz[0]["longitud_ini"].(float64)
	// fmt.Println(LongitudInicial)
	UltimoDatoMatriz, errUltimo := GetElementoMaptoStringToMapArray(matriz[len(matriz)-1]["Fila"])
	fmt.Println(errUltimo)
	LatitudFinal := UltimoDatoMatriz[len(UltimoDatoMatriz)-1]["latitud_fin"].(float64)
	LongitudFinal := UltimoDatoMatriz[len(UltimoDatoMatriz)-1]["longitud_fin"].(float64)
	// fmt.Println(LongitudFinal)
	logs.Info(len(eventos))
	cont := 0
	for i := 0; i < len(eventos)/10; i++ {
		// restriccion a matriz por latitud
		if (eventos[i]["Latitud"].(float64) <= LatitudInicial) && (eventos[i]["Latitud"].(float64) >= LatitudFinal) {
			// restriccion por longitud
			if eventos[i]["Longitud"].(float64) >= LongitudInicial {
				if eventos[i]["Longitud"].(float64) <= LongitudFinal {
					eventos[i]["ID"] = cont
					eventos[i] = ClasificacionCuadrante(eventos[i], matriz)
					arrayEventos = append(arrayEventos, eventos[i])
					// logs.Error(eventos[i])
					cont++
				}
			}
		}
	}
	fmt.Println(cont)
	return arrayEventos
}

// ClasificacionCuadrante ...
func ClasificacionCuadrante(evento map[string]interface{}, matriz []map[string]interface{}) (eventoClasificado map[string]interface{}) {
	for i := 0; i < len(matriz); i++ {
		FilaActual, _ := GetElementoMaptoStringToMapArray(matriz[i]["Fila"])
		// logs.Error(FilaActual[0])
		// fmt.Println(errFila)
		for j := 0; j < len(FilaActual); j++ {
			// logs.Error(FilaActual[j])
			LatitudInicial := FilaActual[j]["latitud_ini"].(float64)
			LongitudInicial := FilaActual[j]["longitud_ini"].(float64)
			LatitudFinal := FilaActual[j]["latitud_fin"].(float64)
			LongitudFinal := FilaActual[j]["longitud_fin"].(float64)
			if (evento["Latitud"].(float64) <= LatitudInicial) && (evento["Latitud"].(float64) >= LatitudFinal) {
				// restriccion por longitud
				if evento["Longitud"].(float64) >= LongitudInicial {
					if evento["Longitud"].(float64) <= LongitudFinal {
						evento["Cuadrante"] = FilaActual[j]["cuadrante"]
						evento["Ubicacion"] = fmt.Sprintf("F%v-C%v", i, j)
						evento["Fila"] = i
						evento["Columna"] = j
						evento["usado"] = false
						// logs.Error(evento)
						return evento
						// cont++
					}
				}
			}
		}
	}
	return nil
}
