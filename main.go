package main

import (
	"fmt"
	"strings"

	"double-list-go/structure"
)

func main() {
	imprimirTitulo("PRUEBA DE LISTA DOBLEMENTE ENLAZADA EN GO")

	// 1. Instanciación y comprobación de lista vacía
	lista := structure.NewDoubleList[string]()
	fmt.Println("1. Estado Inicial de la Lista:")
	fmt.Printf("   - ¿Está vacía? (IsEmpty): %t\n", lista.IsEmpty())
	fmt.Printf("   - Tamaño inicial (Size): %d\n", lista.Size())
	fmt.Printf("   - Representación visual: %s\n\n", lista)

	// 2. Probar el método Add
	imprimirSubtitulo("2. Probando el método Add (agregar elementos)")
	elementos := []string{"Nodo A", "Nodo B", "Nodo C", "Nodo D"}
	for _, elem := range elementos {
		agregado := lista.Add(elem)
		fmt.Printf("   -> Add(\"%s\") exitoso: %t | Lista actual: %s\n", elem, agregado, lista)
	}
	fmt.Printf("\n   - ¿Está vacía tras agregar? (IsEmpty): %t\n", lista.IsEmpty())
	fmt.Printf("   - Tamaño actual (Size): %d\n\n", lista.Size())

	// 3. Probar el método Contains
	imprimirSubtitulo("3. Probando el método Contains (búsqueda de elementos)")
	busquedas := []string{"Nodo B", "Nodo D", "Nodo Z", "Nodo Inexistente"}
	for _, busqueda := range busquedas {
		existe := lista.Contains(busqueda)
		fmt.Printf("   - ¿Contains(\"%s\")?: %t\n", busqueda, existe)
	}
	fmt.Println()

	// 4. Probar el método Remove en distintos casos
	imprimirSubtitulo("4. Probando el método Remove (eliminación en distintos puntos)")

	// Caso A: Eliminar nodo intermedio
	fmt.Println("   [Caso A: Eliminar nodo intermedio ('Nodo B')]")
	eliminadoIntermedio := lista.Remove("Nodo B")
	fmt.Printf("   -> Remove(\"Nodo B\"): %t\n", eliminadoIntermedio)
	fmt.Printf("   -> Estado de la lista: %s (Size: %d)\n\n", lista, lista.Size())

	// Caso B: Eliminar la cabeza
	fmt.Println("   [Caso B: Eliminar la cabeza ('Nodo A')]")
	eliminadoCabeza := lista.Remove("Nodo A")
	fmt.Printf("   -> Remove(\"Nodo A\"): %t\n", eliminadoCabeza)
	fmt.Printf("   -> Estado de la lista: %s (Size: %d)\n\n", lista, lista.Size())

	// Caso C: Eliminar el último nodo
	fmt.Println("   [Caso C: Eliminar el nodo final ('Nodo D')]")
	eliminadoFinal := lista.Remove("Nodo D")
	fmt.Printf("   -> Remove(\"Nodo D\"): %t\n", eliminadoFinal)
	fmt.Printf("   -> Estado de la lista: %s (Size: %d)\n\n", lista, lista.Size())

	// Caso D: Intentar eliminar un elemento que no existe
	fmt.Println("   [Caso D: Intentar eliminar elemento inexistente ('Nodo Fantasma')]")
	eliminadoInexistente := lista.Remove("Nodo Fantasma")
	fmt.Printf("   -> Remove(\"Nodo Fantasma\"): %t\n", eliminadoInexistente)
	fmt.Printf("   -> Estado de la lista: %s (Size: %d)\n\n", lista, lista.Size())

	// Caso E: Eliminar el único nodo restante y comprobar que queda vacía
	fmt.Println("   [Caso E: Eliminar el último nodo restante ('Nodo C')]")
	eliminadoUltimo := lista.Remove("Nodo C")
	fmt.Printf("   -> Remove(\"Nodo C\"): %t\n", eliminadoUltimo)
	fmt.Printf("   -> Estado de la lista: %s\n", lista)
	fmt.Printf("   -> ¿Está vacía ahora? (IsEmpty): %t\n", lista.IsEmpty())
	fmt.Printf("   -> Tamaño final (Size): %d\n\n", lista.Size())

	imprimirTitulo("FIN DE LAS PRUEBAS: TODOS LOS MÉTODOS VERIFICADOS")
}

func imprimirTitulo(texto string) {
	linea := strings.Repeat("=", 70)
	fmt.Println(linea)
	fmt.Printf(" %s\n", texto)
	fmt.Println(linea)
}

func imprimirSubtitulo(texto string) {
	fmt.Println(texto)
	fmt.Println(strings.Repeat("-", 70))
}
