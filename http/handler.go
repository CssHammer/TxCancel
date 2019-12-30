package http

import (
	"txCancel/dbqueries"
)

type Handler struct {
	userQ      dbqueries.UserQI
	compositeQ dbqueries.CompositeQI
}

func NewHandler(userQ dbqueries.UserQI,
	compositeQ dbqueries.CompositeQI) *Handler {

	return &Handler{
		userQ:      userQ,
		compositeQ: compositeQ,
	}

}
