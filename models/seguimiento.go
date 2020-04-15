package models

import (
	"fmt"
	"strconv"
)

// BusquedaAdyacentes ...
func BusquedaAdyacentes(eventos []map[string]interface{}, matriz []map[string]interface{}) {
	for i := 0; i < len(eventos); i++ {
		FilaAuxiliar, _ := GetElementoMaptoStringToMapArray(matriz[0]["Fila"])
		celdas := BuscarCeldas(eventos[i], float64(len(matriz)), float64(len(FilaAuxiliar)))
		fmt.Println(celdas)
	}
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
