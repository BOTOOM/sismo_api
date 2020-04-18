package models

import (
	"fmt"
	"strconv"
)

// BusquedaAdyacentes ...
func BusquedaAdyacentes(eventos []map[string]interface{}, matriz []map[string]interface{}) (resultadoRecorrido []map[string]interface{}) {
	// logs.Informational("tamaño de eventos", len(eventos))
	Recorridos := make([]map[string]interface{}, 0)
	for i := 0; i < len(eventos); i++ {
		arrayRecorrido := make([]map[string]interface{}, 0)
		FilaAuxiliar, _ := GetElementoMaptoStringToMapArray(matriz[0]["Fila"])
		CantidadFilas := float64(len(matriz))
		CantidadColumnas := float64(len(FilaAuxiliar))
		celdas := BuscarCeldas(eventos[i], CantidadFilas, CantidadColumnas)
		// logs.Warn(eventos[i])
		// fmt.Println(i)
		if (i+1 < len(eventos)) && (eventos[i]["usado"] == false) {
			eventos[i]["celdas"] = celdas
			eventos[i]["usado"] = true

			eventos, arrayRecorrido = BuscarEventos(celdas, eventos, i, arrayRecorrido, CantidadColumnas, CantidadFilas)
			datoContruirdo := make(map[string]interface{})
			datoContruirdo = map[string]interface{}{
				"Inicio":    eventos[i],
				"Recorrido": arrayRecorrido,
			}
			Recorridos = append(Recorridos, datoContruirdo)
			// logs.Informational("tamaño de EVENTOS DESPUES", len(eventos))

			// logs.Info(i)
			// logs.Error(eventos[i])
		}

	}
	return Recorridos
}

// BuscarCeldas ... busca que celdas tiene alrededor segun la matriz
func BuscarCeldas(evento map[string]interface{}, cantidadFilas float64, cantidadColumnas float64) []string {
	celdas := make([]string, 0)
	// logs.Warning(evento["Fila"])
	Fila, _ := strconv.ParseFloat(fmt.Sprintf("%v", evento["Fila"]), 64)
	Columna, _ := strconv.ParseFloat(fmt.Sprintf("%v", evento["Columna"]), 64)
	// ubicacionEvento := fmt.Sprintf("%v", evento["Ubicacion"])
	celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila, Columna))

	if Fila-1 >= 0 {
		celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila-1, Columna))
		if Columna-1 > 0 {
			celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila-1, Columna-1))

		}
		if Columna+1 < cantidadColumnas {
			celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila-1, Columna+1))
		}
	}
	if Fila+1 < cantidadFilas {
		celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila+1, Columna))
		if Columna-1 > 0 {
			celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila+1, Columna-1))

		}
		if Columna+1 < cantidadColumnas {
			celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila+1, Columna+1))
		}
	}
	if Columna-1 >= 0 {
		celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila, Columna-1))

	}
	if Columna+1 < cantidadColumnas {
		celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila, Columna+1))
	}
	return celdas
}

// BuscarEventos ...
func BuscarEventos(celdas []string, eventos []map[string]interface{}, indice int, recorrido []map[string]interface{}, numCol float64, numFilas float64) (eventosMod []map[string]interface{}, eventosBuscados []map[string]interface{}) {
	// func BuscarEventos(celdas []string, eventos []map[string]interface{}, indice int, recorrido []map[string]interface{}) (eventosMod []map[string]interface{}, eventosBuscados map[string]interface{}) {
	for i := indice + 1; i < len(eventos); i++ {
		for j := 0; j < len(celdas); j++ {
			ubicacion := fmt.Sprintf("%v", eventos[i]["Ubicacion"])
			if celdas[j] == ubicacion {
				listaEventos := eventos
				if listaEventos[i]["usado"] == false {
					listaEventos[i]["usado"] = true
					// listaEventos[i]["ubicacion_usada"] = listaEventos[i]["Ubicacion"]
					celdasActuales := BuscarCeldas(listaEventos[i], numFilas, numCol)
					listaEventos[i]["celdas"] = celdas
					recorrido = append(recorrido, listaEventos[i])
					// return listaEventos, recorrido
					return BuscarEventos(celdasActuales, listaEventos, i, recorrido, numCol, numFilas)
				}

			}
		}
	}
	return eventos, recorrido

}
