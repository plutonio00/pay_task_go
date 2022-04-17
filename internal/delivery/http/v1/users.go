package v1

import (
	"fmt"
// 	validation "github.com/go-ozzo/ozzo-validation"
// 	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"net/http"
)

func (h *Handler) initUsersRoutes(router *mux.Router) {
	usersRouter := router.PathPrefix(h.baseEndpoint + "users").Subrouter()

	usersRouter.HandleFunc("/replenish", h.Replenish).Methods("POST")
	usersRouter.HandleFunc("/transfer", h.Transfer).Methods("POST")
}

type UserReplenishInput struct {
	UserId int
	Sum    int
}

// @Summary Replenish balance of user
// @Tags users
// @Description endpoint to replenish balance
// @Produce json
// @Param userId formData integer true "User's id"
// @Param sum formData integer true "Amount of money"
// @Success 200 {object} ApiResponse{result=string}
// @Failure 400 {object} ApiResponse{result=string}
// @Failure 404 {object} ApiResponse{result=string}
// @Failure 500 {object} ApiResponse{result=string}
// @Failure default {object} ApiResponse{result=string}
// @Router /users/replenish [post]
func (h *Handler) Replenish(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var userReplenishInput UserReplenishInput
	var decoder = schema.NewDecoder()

	err := decoder.Decode(&userReplenishInput, r.PostForm)

	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Users.Replenish(userReplenishInput.UserId, userReplenishInput.Sum)

	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("Balance for user id was replenished successfully"))
	return
}

type UserTransferInput struct {
	SenderId    int
	RecipientId int
	Sum         int
}

// @Summary Transfer money from one user to another
// @Tags users
// @Description endpoint to transfer money
// @Produce json
// @Param senderId formData integer true "Sender's id"
// @Param recipientId formData integer true "Recipient's title"
// @Param sum formData integer true "Amount of money"
// @Success 200 {object} ApiResponse{result=string}
// @Failure 400 {object} ApiResponse{result=string}
// @Failure 404 {object} ApiResponse{result=string}
// @Failure 500 {object} ApiResponse{result=string}
// @Failure default {object} ApiResponse{result=string}
// @Router /users/transfer [post]
func (h *Handler) Transfer(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var userTransferInput UserTransferInput
	var decoder = schema.NewDecoder()

	err := decoder.Decode(&userTransferInput, r.PostForm)

	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Users.Transfer(userTransferInput.SenderId, userTransferInput.RecipientId, userTransferInput.Sum)

	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("Balance for user id was replenished successfully"))
	return
}
