package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type Livro struct {
	Id        int
	Titulo    string
	Categoria string
	Autor     string
	Sinopse   string
}

func Read(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, titulo, categoria, autor, sinopse FROM livros")
	if err != nil {
		http.Error(w, "Erro ao consultar o banco de dados", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	data := make([]Livro, 0)

	for rows.Next() {
		livro := Livro{}
		err := rows.Scan(&livro.Id, &livro.Titulo, &livro.Categoria, &livro.Autor, &livro.Sinopse)
		if err != nil {
			http.Error(w, "Erro ao ler os dados do banco", http.StatusInternalServerError)
			return
		}
		data = append(data, livro)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Erro ao percorrer os dados", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Erro ao converter os dados para JSON", http.StatusInternalServerError)
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	l := Livro{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, "Erro ao ler dados", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO livros (titulo, categoria, autor, sinopse) VALUES ($1, $2, $3, $4)", l.Titulo, l.Categoria, l.Autor, l.Sinopse)
	if err != nil {
		http.Error(w, "Erro ao criar o livro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}

	lp := Livro{}
	err := json.NewDecoder(r.Body).Decode(&lp)
	if err != nil {
		http.Error(w, "Erro ao ler dados", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT id, titulo, categoria, autor, sinopse FROM livros WHERE id=$1", id)
	l := Livro{}
	err = row.Scan(&l.Id, &l.Titulo, &l.Categoria, &l.Autor, &l.Sinopse)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "Erro ao consultar o banco de dados", http.StatusInternalServerError)
		return
	}

	if lp.Titulo != "" {
		l.Titulo = lp.Titulo
	}
	if lp.Categoria != "" {
		l.Categoria = lp.Categoria
	}
	if lp.Autor != "" {
		l.Autor = lp.Autor
	}
	if lp.Sinopse != "" {
		l.Sinopse = lp.Sinopse
	}

	_, err = db.Exec("UPDATE livros SET titulo=$1, categoria=$2, autor=$3, sinopse=$4 WHERE id=$5", l.Titulo, l.Categoria, l.Autor, l.Sinopse, id)
	if err != nil {
		http.Error(w, "Erro ao atualizar o livro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(l)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	_, err := db.Exec("DELETE FROM livros WHERE id=$1;", id)
	if err != nil {
		http.Error(w, "Erro ao deletar o livro", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://stacoviaki:123456@postgres/crud?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("connected to database!")
}

func main() {
	http.HandleFunc("/livros/update", Update)
	http.HandleFunc("/livros/delete", Delete)
	http.HandleFunc("/livros/read", Read)
	http.HandleFunc("/livros/create", Create)
	http.ListenAndServe(":8080", nil)
}
