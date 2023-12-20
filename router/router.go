package router

import (
	"encoding/json"
	"go-parallel-transactions/domain"
	"go-parallel-transactions/repository"
	"go-parallel-transactions/util"
	"io/ioutil"
	"net/http"
)

type Handler struct{
	UserChannel *chan *domain.User
	TransactionChannel *chan *domain.NewTransactionRequest

	UserDB *[]*domain.User	
}

func New(userDb *[]*domain.User, userChannel *chan *domain.User, transactionChannel *chan *domain.NewTransactionRequest){
	h := Handler{
		UserChannel: userChannel,
		TransactionChannel: transactionChannel,
		UserDB: userDb,
	}


	http.HandleFunc("/api/v1/users", h.handleCreateUser)
	http.HandleFunc("/api/v1/users/all", h.handleFetchUsers)
	http.HandleFunc("/api/v1/transactions", h.handleTransaction)
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

	w.Header().Add("Content-Type", "application/json")
	w.Write(user)
}

func (h *Handler) handleTransaction(w http.ResponseWriter, r *http.Request) {
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

	var requestData domain.NewTransactionRequest
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	err = repository.NewTransaction(&requestData, h.TransactionChannel, h.UserDB)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// w.WriteHeader(http.StatusOK, "application/json")
	w.Write([]byte("OK\r\n"))
}

func (h *Handler) handleFetchUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}


	users, _ := json.Marshal(h.UserDB)
	w.Header().Add("Content-Type", "application/json")
	w.Write(users)
}