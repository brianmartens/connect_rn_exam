package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.POST("/handle1", Handle)
	log.Fatal(e.Start(":3000"))
}

func Handle(e echo.Context) error {
	req := Request{}
	if err := e.Bind(&req); err != nil {
		e.Error(err)
		return e.JSON(http.StatusInternalServerError, err)
	}

	resp := Response{}
	for _, data := range req.Data {

		birthDate, err := time.Parse("2006-01-02", data.DateOfBirth)
		if err != nil {
			e.Error(err)
			return e.JSON(http.StatusBadRequest, err)
		}
		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			e.Error(err)
			return e.JSON(http.StatusInternalServerError, err)
		}
		createTime := time.Unix(data.CreatedOn, 0).In(loc)
		resp.Data = append(resp.Data, ResponseItem{
			UserId:      data.UserId,
			Name:        data.Name,
			DateOfBirth: birthDate.Weekday().String(),
			CreatedOn:   createTime.Format(time.RFC3339),
		})
	}
	return e.JSON(http.StatusOK, resp)
}
