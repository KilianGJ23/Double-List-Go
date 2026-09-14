package main

import (
	"double-list-go/structure"
	"fmt"
)

func main() {
	lista := structure.NewDoubleList[int]()
	fmt.Println("Lista inicial:", lista)

	// Inserciones
	fmt.Println("\n--- Inserciones ---")
	lista.AddFirst(10)
	fmt.Println("Insertar al inicio (10):", lista)

	lista.AddFirst(5)
	fmt.Println("Insertar al inicio (5): ", lista)

	lista.AddFinal(20)
	fmt.Println("Insertar al final (20): ", lista)

	lista.AddFinal(30)
	fmt.Println("Insertar al final (30): ", lista)

	// Eliminaciones
	fmt.Println("\n--- Eliminaciones ---")
	lista.RemoveFirst()
	fmt.Println("Eliminar al inicio:     ", lista)

	lista.RemoveFinal()
	fmt.Println("Eliminar al final:      ", lista)

	// Búsqueda (Contains)
	fmt.Println("\n--- Comprobación (Contains) ---")
	fmt.Println("¿Contiene el 20?:", lista.Contains(20))
	fmt.Println("¿Contiene el 99?:", lista.Contains(99))
}
