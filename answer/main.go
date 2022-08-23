package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	e := echo.New()
	e.POST("/users", UsersHandle)
	e.POST("/image", ImageHandle)
	log.Fatal(e.Start(":3000"))
}

func UsersHandle(e echo.Context) error {
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

func ImageHandle(e echo.Context) error {
	file, err := e.FormFile("image.jpg")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	jpg, err := io.ReadAll(src)
	if err != nil {
		e.Error(err)
		return e.JSON(http.StatusInternalServerError, err)
	}

	if err := convertImage(jpg); err != nil {
		e.Error(err)
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.Attachment("image.png", "image.png")
}

func convertImage(jpg []byte) error {
	if err := os.WriteFile("image.jpg", jpg, 0666); err != nil {
		return err
	}
	cmd := exec.Command("resize.sh", "image")
	if err := cmd.Run(); err != nil {
		return err
	}
	if err := os.Remove("image.jpg"); err != nil {
		return err
	}
	return nil
}
