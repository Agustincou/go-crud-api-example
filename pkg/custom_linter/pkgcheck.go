package main //nolint: pkgnamechecker

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var Analyzer = &analysis.Analyzer{
	Name: "pkgnamechecker",
	Doc:  "Verifica que el nombre del paquete coincide con el nombre del directorio",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	var errors []string

	for _, file := range pass.Files {
		// Obtener la ruta del archivo actual
		filePath := pass.Fset.File(file.Package).Name()

		// Verificar si el archivo tiene la extensión .go
		if filepath.Ext(filePath) != ".go" {
			continue // Ignora el archivo si no es un .go
		}

		// Crear un nuevo conjunto de archivos
		fset := token.NewFileSet()

		// Parsear el archivo Go
		node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("error al analizar el archivo %s: %w", filePath, err)
		}

		// Verificar si el archivo tiene el comentario de ignore
		if hasNolintComment(node.Comments) {
			continue // Ignorar este archivo
		}

		// Obtener el nombre del directorio
		dir := filepath.Base(filepath.Dir(filePath))

		// Obtener el nombre del paquete declarado
		packageName := node.Name.Name

		// Comparar el nombre del directorio con el nombre del paquete
		if packageName != dir {
			errors = append(errors, fmt.Sprintf("el nombre del paquete '%s' no coincide con el nombre del directorio '%s' en el archivo %s", packageName, dir, filePath))
		}
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("Se encontraron problemas:\n%s", strings.Join(errors, "\n"))
	}

	return nil, nil
}

// Función para verificar si hay un comentario que ignora el análisis
func hasNolintComment(comments []*ast.CommentGroup) bool {
	for _, group := range comments {
		for _, comment := range group.List {
			if strings.Contains(comment.Text, "//nolint: pkgnamechecker") {
				return true
			}
		}
	}

	return false
}

func main() {
	singlechecker.Main(Analyzer)
}
