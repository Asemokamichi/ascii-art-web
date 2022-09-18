package delivery

import (
	"net/http"
	"text/template"
)

func ServErrors(w http.ResponseWriter, Error TextProcessing) {
	temp, err1 := template.ParseFiles("templates/error.html")
	w.WriteHeader(Error.ErrorModifiedText)
	if err1 != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	err2 := temp.Execute(w, Error)
	if err2 != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}
