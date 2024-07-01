package handler

import (
	"GolangAPI"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createCar(c *gin.Context) {
	var input GolangAPI.Car
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Car.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Статус:":     201,
		"id(Created)": id,
	})
}
func (h *Handler) getCar(c *gin.Context) {
	type getCarResponse struct {
		Data GolangAPI.Car `json:"data"`
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	car, err := h.services.Car.GetCar(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getCarResponse{car})
	c.JSON(http.StatusOK, "Статус: 200")
}
func (h *Handler) updateCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input GolangAPI.UpdateCar
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Статус: 204(No content)")
}
func (h *Handler) updateCarPartly(c *gin.Context) {

}
func (h *Handler) deleteCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Car.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Статус: 204(No content)")
}

func (h *Handler) getAllCars(c *gin.Context) {
	type getAllCarsResponse struct {
		Data []GolangAPI.Car `json:"data"`
	}
	cars, err := h.services.Car.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCarsResponse{cars})
	c.JSON(http.StatusOK, "Статус: 200")
}
