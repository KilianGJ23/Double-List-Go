# Implementación de Lista Doblemente Enlazada en Go

## Descripción del Proyecto

Este proyecto consiste en el diseño e implementación de una **Lista Doblemente Enlazada (`DoubleList`)** desarrollada de manera nativa en el lenguaje de programación **Go (Golang)**.

A diferencia de las listas simples, una lista doblemente enlazada es una estructura de datos lineal y dinámica compuesta por una serie de nodos independientes (`DoubleNode`), donde cada uno contiene no solo el dato a almacenar sino también dos punteros o referencias en memoria: uno hacia el elemento sucesor (`next`) y otro hacia el elemento predecesor (`previous`). Esta doble conectividad ofrece una navegación bidireccional y una reconfiguración eficiente de enlaces al momento de insertar o suprimir elementos. En esta implementación, la lista es gestionada bajo una arquitectura basada en un único punto de acceso principal: el puntero **cabeza (`head`)**, centralizando las operaciones de adición al final (`Add`), búsqueda lineal (`Contains`), retiro seguro de nodos (`Remove`), verificación de contenido (`IsEmpty`) y cálculo de longitud (`Size`).

### Paradigma de Programación Aplicado

El proyecto se fundamenta en tres pilares del paradigma de programación adoptado por Go:

* **Paradigma Imperativo y Estructurado:** El flujo de los algoritmos se expresa mediante secuencias de instrucciones, condicionales y ciclos iterativos explícitos encargados de manipular directamente los punteros de memoria y actualizar las referencias bidireccionales de los nodos.

* **Orientación a Objetos Basada en Tipos y Métodos:** Go prescinde de la herencia y de las clases tradicionales. En su lugar, el encapsulamiento se logra definiendo estructuras de datos (`struct`) asociadas a métodos con receptores (*receiver methods*), permitiendo tratar a `DoubleList` y `DoubleNode` como tipos de datos abstractos (TDA) autónomos.

## Requisitos de Software e Instalación

Para compilar y ejecutar este proyecto se requiere tener instalado el compilador de Go.

* **Versión de Go utilizada en el desarrollo:** **Go 1.27** *(o cualquier versión compatible igual o superior a Go 1.18, requerida para el soporte de genéricos)*.

### ¿Cómo instalar Go?

1. **Descarga:** Ingrese a [go.dev/dl](https://go.dev/dl/) y descargue el instalador correspondiente a su sistema operativo.

2. **Instalación:**

   * **Linux:** Descomprima el archivo descargado en `/usr/local` (por ejemplo: `sudo tar -C /usr/local -xzf go1.27.linux-amd64.tar.gz`) y agregue `export PATH=$PATH:/usr/local/go/bin` en su archivo `~/.bashrc` o `~/.zshrc`.

   * **macOS:** Ejecute el paquete instalador `.pkg` descargado o instálelo vía Homebrew con `brew install go`.
   
   * **Windows:** Ejecute el instalador `.msi` y siga las instrucciones del asistente (el instalador configura automáticamente las variables de entorno).

3. **Verificación:** Abra una terminal o consola y compruebe la instalación ejecutando:

   ```bash
   go version
   ```

## Instrucciones de Ejecución

Primero, clone el repositorio público en su máquina local:

```bash
git clone https://github.com/KilianGJ23/Double-List-Go.git
cd Double-List-Go
```

Dependiendo de su sistema operativo, ejecute la aplicación de pruebas `main.go` siguiendo los pasos descritos a continuación:

### En Linux / macOS

Abra su terminal favorita (Bash, Zsh, etc.) dentro del directorio del proyecto y ejecute:

```bash
# Ejecución directa del archivo de pruebas
go run main.go

# (Opcional) Compilar un binario ejecutable y correrlo
go build -o double-list-go main.go
./double-list-go
```

### En Windows

Abra **PowerShell** o el **Símbolo del sistema (CMD)** en la carpeta del repositorio y ejecute:

```powershell
# Ejecución directa en PowerShell o CMD
go run main.go

# (Opcional) Compilar el ejecutable .exe
go build -o double-list-go.exe main.go
.\double-list-go.exe
```

## Autores

* **KILIAN DAVID GOMEZ JOROPA** — Código: `202410372`
* **CRISTIAN CAMILO BLANCO CASTILLO** — Código: `202410713`