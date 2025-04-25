package main

import (
	"errors"
	"fmt"
	"net/http"
	"task-app/internal/repository"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"task-app/internal/auth"
)

type application struct {
	config config
	store repository.TaskRepo
}

type config struct {
	maxOpenConns int
	maxIdleConns int
	maxIdleTime string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	
	r.Route("/health", func(r chi.Router) {
		r.Get("/", app.healthCheckHandler)
	})
	r.Route("/v1/tasks", func(r chi.Router) {
		r.Get("/", app.getTaskHandler)
		r.Post("/", app.CreateTaskHandler)
		r.Put("/{id}", app.toggleTaskHandler)
		r.Delete("/{id}", app.deleteHandler)
	})

	r.Route("/v1/auth", func(r chi.Router) {
		r.Get("/google", func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, auth.GoogleLoginURL(), http.StatusTemporaryRedirect)
		})
		r.Get("/google/callback", func(w http.ResponseWriter, r *http.Request) {
				code := r.URL.Query().Get("code")
				email, err := auth.HandleGoogleCallback(code)
				if err != nil {
						http.Error(w, "Failed to authenticate", http.StatusInternalServerError)
						return
				}

				token, err := auth.GenerateJWT(email)
				if err != nil {
						http.Error(w, "Failed to generate token", http.StatusInternalServerError)
						return
				}

				http.SetCookie(w, &http.Cookie{
						Name:     "token",
						Value:    token,
						HttpOnly: true,
				})
				w.Write([]byte("Login successful"))
		})
		r.Get("/google/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		email, err := auth.HandleGoogleCallback(code)
		if err != nil {
				http.Error(w, "Failed to authenticate", http.StatusInternalServerError)
				return
		}

		token, err := auth.GenerateJWT(email)
		if err != nil {
				http.Error(w, "Failed to generate token", http.StatusInternalServerError)
				return
		}

		// Trả token về frontend
		http.Redirect(w, r, "http://localhost:5173/dashboard?token="+token, http.StatusSeeOther)
	})
	})

	return r
}


func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	fmt.Println("Served has started", "addr", ":8080")	

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}



