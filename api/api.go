package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tanookiben/benkyo/internal/config"
	"github.com/tanookiben/benkyo/internal/study/japanese"
)

// API ...
type API interface {
	Serve()
}

type apiImpl struct {
	c config.Config
	r *routerImpl
}

type routerImpl struct {
	Router *mux.Router
}

// NewAPI ...
func NewAPI(c config.Config) API {
	mux := mux.NewRouter()
	r := &routerImpl{
		Router: mux,
	}

	mux.HandleFunc("/", study)
	mux.HandleFunc("/answer", answer)
	mux.HandleFunc("/health", health)

	return &apiImpl{
		c: c,
		r: r,
	}
}

func (a *apiImpl) Serve() {
	s := &http.Server{
		Addr:    a.c.Addr,
		Handler: a.r.Router,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("API server error: %v", err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func study(w http.ResponseWriter, r *http.Request) {
	charset := r.URL.Query().Get("charset")
	if charset == "" {
		charset = japanese.HiraganaStudy
	}
	newPrompt(w, charset, "")
}

func answer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		write(w, "Error parsing request form")
	}
	charset := r.Form.Get("charset")
	prompt := r.Form.Get("prompt")
	answer := r.Form.Get("answer")
	if japanese.Check(charset, prompt, answer) {
		newPrompt(w, charset, "")
		return
	}
	newPrompt(w, charset, prompt)
	return
}

func newPrompt(w http.ResponseWriter, charset, prompt string) {
	var wrong string
	if prompt == "" {
		prompt = japanese.Prompt(charset)
	} else {
		wrong = " (Try again)"
	}
	write(w, fmt.Sprintf(`
<!DOCTYPE html>
<html>
	<body>
		<h1>%s%s</h1>
		<form action="/answer">
		<input type="hidden" id="charset" name="charset" value="%s"/>
		<input type="hidden" id="prompt" name="prompt" value="%s"/>
		<label for="answer">Answer</label><br>
		<input type="text" id="answer" name="answer"><br>
		<input type="submit" value="Check">
		</form> 
	</body>
</html>
	`, prompt, wrong, charset, prompt))
}

func write(w http.ResponseWriter, m string) {
	if _, err := w.Write([]byte(m)); err != nil {
		log.Printf("Error writing response: %v\n", err)
	}
}
