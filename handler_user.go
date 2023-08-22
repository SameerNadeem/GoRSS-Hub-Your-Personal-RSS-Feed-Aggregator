package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SameerNadeem/rssag/internal/database"
	"github.com/google/uuid"
)

func (apiCfg*apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON:%v",err))
		return 
	}

	user, err := apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		
	})
	if err != nil {
		respondWithError(w, 400, fmt. Sprintf ("Couldn't create user:%v", err))
		 return
	}

	respondwithJSON(w, 200, user)
}