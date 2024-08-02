package analyzer

import (
	commands "CLASE02/commands" // Importa el paquete "commands" desde el directorio "PRUEBA01/commands"
	"errors"                    // Importa el paquete "errors" para manejar errores
	"fmt"                       // Importa el paquete "fmt" para formatear e imprimir texto
	"os"                        // Importa el paquete "os" para interactuar con el sistema operativo
	"os/exec"                   // Importa el paquete "os/exec" para ejecutar comandos del sistema
	"strings"                   // Importa el paquete "strings" para manipulación de cadenas
)

// Analyzer analiza el comando de entrada y ejecuta la acción correspondiente
func Analyzer(input string) (interface{}, error) {
	// Divide la entrada en tokens usando espacios en blanco como delimitadores
	tokens := strings.Fields(input)

	// Si no se proporcionó ningún comando, devuelve un error
	if len(tokens) == 0 {
		return nil, errors.New("no se proporcionó ningún comando")
	}

	// Switch para manejar diferentes comandos
	switch tokens[0] {
	case "mkdisk":
		// Llama a la función ParseMkdisk del paquete commands con los argumentos restantes
		return commands.ParseMkdisk(tokens[1:])
	case "add":
		// Llama a la función ParseAdd del paquete commands con los argumentos restantes
		return commands.ParseAdd(tokens[1:])
	case "rep":
		// Llama a la función ParseRep del paquete commands
		return commands.ParseRep()
	case "clear":
		// Crea un comando para limpiar la terminal
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout // Redirige la salida del comando a la salida estándar
		err := cmd.Run()       // Ejecuta el comando
		if err != nil {
			// Si hay un error al ejecutar el comando, devuelve un error
			return nil, errors.New("no se pudo limpiar la terminal")
		}
		return nil, nil // Devuelve nil si el comando se ejecutó correctamente
	default:
		// Si el comando no es reconocido, devuelve un error
		return nil, fmt.Errorf("comando desconocido: %s", tokens[0])
	}
}