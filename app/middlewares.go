package app

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
)

func accessLogs(next http.Handler) http.Handler {
	var writer io.Writer = os.Stdout

	return middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger: log.New(writer, "", log.LstdFlags),
	})(next)
}
