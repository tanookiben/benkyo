package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	newPrompt(w, charset, "", 0)
}

func answer(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		write(w, "Error parsing request form")
	}
	charset := r.Form.Get("charset")
	prompt := r.Form.Get("prompt")
	answer := r.Form.Get("answer")
	attempts, err := strconv.Atoi(r.Form.Get("attempts"))
	if err != nil {
		attempts = 0
	}
	if japanese.Check(charset, prompt, answer) {
		newPrompt(w, charset, "", 0)
		return
	}
	newPrompt(w, charset, prompt, attempts+1)
	return
}

func newPrompt(w http.ResponseWriter, charset, prompt string, attempts int) {
	var message string
	if prompt == "" {
		prompt = japanese.Prompt(charset)
	} else {
		message = " (try again)"
	}
	if attempts > 4 {
		message = fmt.Sprintf(" (%q: %q)", prompt, japanese.Hint(charset, prompt))
		prompt = japanese.Prompt(charset)
	}
	swap := japanese.KatakanaStudy
	if charset == japanese.KatakanaStudy {
		swap = japanese.HiraganaStudy
	}
	write(w, fmt.Sprintf(`
<!DOCTYPE html>
<html>
	<body>
		<h1><a href="?charset=%s">%s</a></h1>
		<h2>%s%s</h2>
		<form action="/answer">
		<input type="hidden" id="charset" name="charset" value="%s"/>
		<input type="hidden" id="prompt" name="prompt" value="%s"/>
		<input type="hidden" id="attempts" name="attempts" value="%d"/>
		<label for="answer">Answer</label><br>
		<input type="text" id="answer" name="answer"><br>
		<input type="submit" value="Check">
		</form> 
	</body>
</html>
	`, swap, strings.Title(charset), prompt, message, charset, prompt, attempts))
}

func write(w http.ResponseWriter, m string) {
	if _, err := w.Write([]byte(m)); err != nil {
		log.Printf("Error writing response: %v\n", err)
	}
}
