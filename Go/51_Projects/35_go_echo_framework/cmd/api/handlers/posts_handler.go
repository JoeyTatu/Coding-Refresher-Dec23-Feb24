package handlers

import (
	"net/http"
	"strconv"

	"github.com/joeytatu/go-echo-framework/cmd/api/service"
	"github.com/labstack/echo"
)

func PostIndexHandler(ctx echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		ctx.String(http.StatusBadGateway, "Unable to proccess data")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data

	return ctx.JSON(http.StatusOK, res)
}

func PostSingleHandler(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.String(http.StatusNotFound, "Unable to get ID")
	}

	data, err := service.GetByID(id)
	if err != nil {
		ctx.String(http.StatusBadGateway, "Unable to proccess data")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data

	return ctx.JSON(http.StatusOK, res)
}
