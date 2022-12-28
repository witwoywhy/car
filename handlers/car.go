package handlers

import (
	"car/dtos"
	"car/services"

	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	serv *services.CarService
}

var failed = map[string]string{
	"message": "failed",
}

var responseFailed = func() (int, map[string]string) {
	return http.StatusInternalServerError, failed
}

func must(c *gin.Context, err error) bool {
	if err != nil {
		log.Println(err)
		c.JSON(responseFailed())
		return false
	}

	return true
}

func (h *handler) Create(c *gin.Context) {
	var dto dtos.Car

	err := c.BindJSON(&dto)
	if !must(c, err) {
		return
	}

	car, err := h.serv.Create(dto.Brand, dto.Name, dto.Model, dto.SubModel, dto.Color, dto.Price)
	c.JSON(http.StatusCreated, car)
}

func (h *handler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if !must(c, err) {
		return
	}

	car, err := h.serv.FindByID(id)
	if !must(c, err) {
		return
	}

	c.JSON(http.StatusOK, car)
}

func (h *handler) FindMany(c *gin.Context) {
	car, err := h.serv.FindMany()
	if !must(c, err) {
		return
	}

	c.JSON(http.StatusOK, car)
}

func (h *handler) Update(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if !must(c, err) {
		return
	}

	var dto dtos.Car
	err = c.BindJSON(&dto)
	if !must(c, err) {
		return
	}

	car, err := h.serv.Update(id, dto.Brand, dto.Name, dto.Model, dto.SubModel, dto.Color, dto.Price)
	if !must(c, err) {
		return
	}

	c.JSON(http.StatusOK, car)
}

func (h *handler) Delete(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if !must(c, err) {
		return
	}

	err = h.serv.Delete(id)
	if !must(c, err) {
		return
	}
	c.JSON(http.StatusOK, nil)
}

func Newhandler(serv *services.CarService) *handler {
	return &handler{
		serv: serv,
	}
}
