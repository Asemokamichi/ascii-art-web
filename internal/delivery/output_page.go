package delivery

import (
	"ascii-art/internal/utils"
	"net/http"
	"strconv"
	"text/template"
)

// models move
type TextProcessing struct {
	ModifiedText      string
	ErrorModifiedText int
}

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Route() *http.ServeMux {
	s.mux.HandleFunc("/", HMainPage)
	s.mux.HandleFunc("/ascii-art", hAscii)
	s.mux.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	return s.mux
}

func hAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ServErrors(w, TextProcessing{http.StatusText(http.StatusNotFound), http.StatusNotFound})
		return
	}
	if r.Method != "POST" {
		ServErrors(w, TextProcessing{http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed})
		return
	}
	var (
		banner string
		text   string
		Text   TextProcessing
	)

	banner = r.FormValue("banner")
	text = r.FormValue("text")

	Text.ModifiedText, Text.ErrorModifiedText = utils.SetAsciiArt(text, banner)
	if Text.ErrorModifiedText != 0 {
		ServErrors(w, Text)
		return
	}
	if _, ok := r.Form["submit"]; ok {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			ServErrors(w, TextProcessing{http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError})
			return
		}
		err1 := tmpl.Execute(w, Text)
		if err1 != nil {
			ServErrors(w, TextProcessing{http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError})
			return
		}
	} else if _, ok := r.Form["export"]; ok {
		w.Header().Set("Content-Disposition", "attachment; filename=AsciiArt.txt")
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", strconv.Itoa(len(Text.ModifiedText)))

		_, err := w.Write([]byte(Text.ModifiedText))
		if err != nil {
			ServErrors(w, TextProcessing{http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError})
			return
		}
	} else {
		ServErrors(w, TextProcessing{http.StatusText(http.StatusBadRequest), http.StatusBadRequest})
		return
	}
}
