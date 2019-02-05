package api

import (
	"strings"
	"os"
	"bufio"
	"encoding/json"
	"net/http"
	"github.com/luizhenriqueandrade/go-project/movies"
	"fmt"
)

type HomepageHandler struct{}

func(h *HomepageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bem vindo a página inicial!\nInsira /new para adicionar um filme!")
}


type MovieHandler struct{}

func (m *MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var movieName string
	var movieYear string
	var movieNac string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("nome filme: ")
	movieName, _ = reader.ReadString('\n')
	movieName = strings.TrimSuffix(movieName, "\n")

	fmt.Print("ano filme: ")
	movieYear, _ = reader.ReadString('\n')
	movieYear = strings.TrimSuffix(movieYear, "\n")
	fmt.Print("nacionalidade: ")
	movieNac, _ = reader.ReadString('\n')
	movieNac = strings.TrimSuffix(movieNac, "\n")

	movie := &movies.Movie{
		Name: movieName,
		Year: movieYear ,
		Nacionality: movieNac,
	}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}