package transport

import (
	"encoding/json"
	"github.com/Varsilias/bytesizego-course/internal/todo"
	"log"
	"net/http"
)

type TodoItems struct {
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		todoItems, err := todoSvc.GetAll()

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		todoList, err := json.Marshal(todoItems)

		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(todoList)

		if err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /todos", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItems
		err := json.NewDecoder(request.Body).Decode(&t)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = todoSvc.Add(t.Item)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("GET /todos/search/{s}", func(w http.ResponseWriter, request *http.Request) {
		s := request.PathValue("s")
		if s == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		todoItems, err := todoSvc.Search(s)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		todos, err := json.Marshal(todoItems)

		if err != nil {
			log.Println(err)
		}

		_, err = w.Write(todos)

		if err != nil {
			log.Println(err)
		}
	})

	return &Server{
		mux,
	}
}

func (s *Server) ServeHTTP() error {
	return http.ListenAndServe(":9000", s.mux)
}
