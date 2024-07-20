package api

import (
	CRUD "bank/sql_go"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	if err := server.validate_account_currency(ctx, request); err != nil {
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

func (server *Server) validate_account_currency(ctx *gin.Context, arg transferMoneyRequest) error {
	arg_from := CRUD.GetAccountForUpdateParam{
		ID: arg.From_account_id,
	}
	from_account, err := server.store.GetAccountForUpdate(ctx, arg_from)
	if err != nil {
		if err != nil {
			if err == sql.ErrNoRows {
				return err
			}
			return err
		}
	}

	arg_to := CRUD.GetAccountForUpdateParam{
		ID: arg.To_account_id,
	}
	to_account, err := server.store.GetAccountForUpdate(ctx, arg_to)
	if err != nil {
		if err == sql.ErrNoRows {
			err := fmt.Errorf("Account ID %d didn't register\n", arg_to.ID)
			return err
		}
		return err
	}
	if arg.Currency != from_account.Currency || arg.Currency != to_account.Currency {
		err := fmt.Errorf("Currency can't match!\n")
		return err
	}

	return nil
}
