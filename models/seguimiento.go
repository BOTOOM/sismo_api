package models

import (
	"fmt"
	"strconv"
)

// BusquedaAdyacentes ...
func BusquedaAdyacentes(eventos []map[string]interface{}, matriz []map[string]interface{}) (resultadoRecorrido []map[string]interface{}) {
	arrayRecorrido := make([]map[string]interface{}, 0)
	// logs.Informational("tamaño de eventos", len(eventos))
	for i := 0; i < len(eventos); i++ {
		FilaAuxiliar, _ := GetElementoMaptoStringToMapArray(matriz[0]["Fila"])
		celdas := BuscarCeldas(eventos[i], float64(len(matriz)), float64(len(FilaAuxiliar)))
		eventos[i]["celdas"] = celdas
		// logs.Warn(eventos[i])
		// fmt.Println(i)
		if (i+1 < len(eventos)) && (eventos[i]["usado"] == false) {
			eventos, arrayRecorrido = BuscarEventos(celdas, eventos, i, arrayRecorrido)
			// logs.Informational("tamaño de EVENTOS DESPUES", len(eventos))

			// logs.Info(i)
			// logs.Error(eventos[i])
		}

	}
	return arrayRecorrido
}

// BuscarCeldas ... busca que celdas tiene alrededor segun la matriz
func BuscarCeldas(evento map[string]interface{}, cantidadFilas float64, cantidadColumnas float64) []string {
	celdas := make([]string, 0)
	// logs.Warning(evento["Fila"])
	Fila, _ := strconv.ParseFloat(fmt.Sprintf("%v", evento["Fila"]), 64)
	Columna, _ := strconv.ParseFloat(fmt.Sprintf("%v", evento["Columna"]), 64)
	if Fila-1 > 0 {
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
	if Columna-1 > 0 {
		celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila, Columna-1))

	}
	if Columna+1 < cantidadColumnas {
		celdas = append(celdas, fmt.Sprintf("F%v-C%v", Fila, Columna+1))
	}
	return celdas
}

// BuscarEventos ...
func BuscarEventos(celdas []string, eventos []map[string]interface{}, indice int, recorrido []map[string]interface{}) (eventosMod []map[string]interface{}, eventosBuscados []map[string]interface{}) {
	// func BuscarEventos(celdas []string, eventos []map[string]interface{}, indice int, recorrido []map[string]interface{}) (eventosMod []map[string]interface{}, eventosBuscados map[string]interface{}) {
	for i := indice + 1; i < len(eventos); i++ {
		for j := 0; j < len(celdas); j++ {
			ubicacion := fmt.Sprintf("%v", eventos[i]["Ubicacion"])
			if celdas[j] == ubicacion {
				listaEventos := eventos
				listaEventos[i]["usado"] = true
				recorrido = append(recorrido, listaEventos[i])
				return listaEventos, recorrido

			}
		}
	}
	return eventos, recorrido

}
