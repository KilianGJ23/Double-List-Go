package main

import (
	"double-list-go/structure"
	"fmt"
	"strings"
)

func main() {
	imprimirTitulo("PRUEBA DE LISTA DOBLEMENTE ENLAZADA EN GO")

	// 1. Instanciación y comprobación de lista vacía
	lista := structure.NewDoubleList[string]()
	fmt.Println("1. Estado Inicial de la Lista:")
	fmt.Printf("   - ¿Está vacía? (IsEmpty): %t\n", lista.IsEmpty())
	fmt.Printf("   - Tamaño inicial (Size): %d\n", lista.Size())
	fmt.Printf("   - Representación visual: %s\n", lista)
	fmt.Printf("   - Intentar RemoveFirst() en lista vacía: %t\n", lista.RemoveFirst())
	fmt.Printf("   - Intentar RemoveFinal() en lista vacía: %t\n\n", lista.RemoveFinal())

	// 2. Probar el método AddFinal (agregar al final)
	imprimirSubtitulo("2. Probando el método AddFinal (agregar al final)")
	elementosFinal := []string{"Nodo B", "Nodo C"}
	for _, elem := range elementosFinal {
		agregado := lista.AddFinal(elem)
		fmt.Printf("   -> AddFinal(\"%s\") exitoso: %t | Lista actual: %s\n", elem, agregado, lista)
	}
	fmt.Printf("   - Tamaño actual: %d\n\n", lista.Size())

	// 3. Probar el método AddFirst (agregar al inicio)
	imprimirSubtitulo("3. Probando el método AddFirst (agregar al inicio)")
	elementosInicio := []string{"Nodo A", "Nodo Inicio"}
	for _, elem := range elementosInicio {
		agregado := lista.AddFirst(elem)
		fmt.Printf("   -> AddFirst(\"%s\") exitoso: %t | Lista actual: %s\n", elem, agregado, lista)
	}
	fmt.Printf("\n   - ¿Está vacía tras inserciones? (IsEmpty): %t\n", lista.IsEmpty())
	fmt.Printf("   - Tamaño actual (Size): %d\n", lista.Size())
	fmt.Printf("   - Estado de la lista completa: %s\n\n", lista)

	// 4. Probar el método Contains
	imprimirSubtitulo("4. Probando el método Contains (búsqueda de elementos)")
	busquedas := []string{"Nodo Inicio", "Nodo A", "Nodo B", "Nodo C", "Nodo Z", "Nodo Inexistente"}
	for _, busqueda := range busquedas {
		existe := lista.Contains(busqueda)
		fmt.Printf("   - ¿Contains(\"%s\")?: %t\n", busqueda, existe)
	}
	fmt.Println()

	// 5. Probar el método RemoveFirst (eliminación al inicio)
	imprimirSubtitulo("5. Probando el método RemoveFirst (eliminación al inicio)")

	fmt.Println("   [Caso A: Eliminar el primer nodo ('Nodo Inicio')]")
	eliminadoPrimero1 := lista.RemoveFirst()
	fmt.Printf("   -> RemoveFirst(): %t\n", eliminadoPrimero1)
	fmt.Printf("   -> Estado de la lista: %s (Size: %d)\n\n", lista, lista.Size())

	fmt.Println("   [Caso B: Eliminar el nuevo primer nodo ('Nodo A')]")
	eliminadoPrimero2 := lista.RemoveFirst()
	fmt.Printf("   -> RemoveFirst(): %t\n", eliminadoPrimero2)
	fmt.Printf("   -> Estado de la lista: %s (Size: %d)\n\n", lista, lista.Size())

	// 6. Probar el método RemoveFinal (eliminación al final)
	imprimirSubtitulo("6. Probando el método RemoveFinal (eliminación al final)")

	fmt.Println("   [Caso A: Eliminar el último nodo ('Nodo C')]")
	eliminadoFinal1 := lista.RemoveFinal()
	fmt.Printf("   -> RemoveFinal(): %t\n", eliminadoFinal1)
	fmt.Printf("   -> Estado de la lista: %s (Size: %d)\n\n", lista, lista.Size())

	fmt.Println("   [Caso B: Eliminar el único nodo restante ('Nodo B')]")
	eliminadoFinal2 := lista.RemoveFinal()
	fmt.Printf("   -> RemoveFinal(): %t\n", eliminadoFinal2)
	fmt.Printf("   -> Estado de la lista: %s\n", lista)
	fmt.Printf("   -> ¿Está vacía ahora? (IsEmpty): %t\n", lista.IsEmpty())
	fmt.Printf("   -> Tamaño final (Size): %d\n\n", lista.Size())

	// 7. Probar eliminaciones en lista nuevamente vacía
	imprimirSubtitulo("7. Probando eliminaciones cuando la lista ya está vacía")
	fmt.Printf("   -> RemoveFirst() en lista vacía: %t\n", lista.RemoveFirst())
	fmt.Printf("   -> RemoveFinal() en lista vacía: %t\n", lista.RemoveFinal())
	fmt.Printf("   -> Estado de la lista: %s (Size: %d)\n\n", lista, lista.Size())

	// 8. Prueba adicional: Inserción y eliminación de un único elemento
	imprimirSubtitulo("8. Prueba de ciclo de vida con un único elemento")
	lista.AddFirst("Elemento Solitario")
	fmt.Printf("   -> Tras AddFirst(\"Elemento Solitario\"): %s (Size: %d)\n", lista, lista.Size())
	lista.RemoveFirst()
	fmt.Printf("   -> Tras RemoveFirst(): %s (Size: %d, IsEmpty: %t)\n", lista, lista.Size(), lista.IsEmpty())

	lista.AddFinal("Otro Elemento")
	fmt.Printf("   -> Tras AddFinal(\"Otro Elemento\"): %s (Size: %d)\n", lista, lista.Size())
	lista.RemoveFinal()
	fmt.Printf("   -> Tras RemoveFinal(): %s (Size: %d, IsEmpty: %t)\n\n", lista, lista.Size(), lista.IsEmpty())

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
