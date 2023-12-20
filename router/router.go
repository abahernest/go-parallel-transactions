package router

import (
	"encoding/json"
	"go-parallel-transactions/domain"
	"go-parallel-transactions/repository"
	"go-parallel-transactions/util"
	"io/ioutil"
	"net/http"
	// "fmt"
)

type Handler struct{
	UserChannel *chan *domain.User
	UserDB *[]*domain.User	
}

func New(userDb *[]*domain.User, userChannel *chan *domain.User){
	h := Handler{
		UserChannel: userChannel,
		UserDB: userDb,
	}


	http.HandleFunc("/api/v1/user", h.handleCreateUser)
}


func (h *Handler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Error reading request body")
		return
	}

	var requestData domain.CreateUserRequest
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	newUser := repository.CreateUser(&requestData, h.UserChannel, h.UserDB)
	user, _ := json.Marshal(newUser)
	// userDb, _ := json.Marshal(h.UserDB)
	// fmt.Println(string(userDb), "\n")

	w.Write([]byte(user))
}