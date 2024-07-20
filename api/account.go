package api

import (
	CRUD "bank/sql_go"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=NT USD"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var request createAccountRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := CRUD.CreateAccountParam{
		Owner:    request.Owner,
		Currency: request.Currency,
		Balance:  0,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var request getAccountRequest
	ctx.Writer.Header().Set("Content-Type", "application/json")

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := CRUD.GetAccountParam{
		ID: request.ID,
	}

	account, err := server.store.GetAccount(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSZ int32 `form:"page_sz" binding:"required,min=5,max=10"`
}

func (server *Server) listAccount(ctx *gin.Context) {
	var request listAccountRequest

	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := CRUD.ListAccountParam{
		Limit:  request.PageSZ,
		Offset: (request.PageID - 1) * request.PageSZ,
	}

	account, err := server.store.ListAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type transferMoneyRequest struct {
	From_account_id int64  `json:"from_account_id" binding:"required,min=1"`
	To_account_id   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount          int64  `json:"amount" binding:"required,min=1"`
	Currency        string `json:"currency" binding:"required,oneof=NT USD"`
}

func (server *Server) transferMoney(ctx *gin.Context) {
	var request transferMoneyRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := CRUD.TransactionTxParam{
		From_account_id: request.From_account_id,
		To_account_id:   request.To_account_id,
		Amount:          request.Amount,
	}

	account, err := server.store.TransactionTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}
