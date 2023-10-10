package server

import (
	"functional/prey"
	"functional/shark"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	shark shark.Shark
	prey  prey.Prey
}

func NewHandler(shark shark.Shark, prey prey.Prey) *Handler {
	return &Handler{shark: shark, prey: prey}
}

// PUT: /v1/shark

func (h *Handler) ConfigureShark() gin.HandlerFunc {
	type request struct {
		XPosition float64 `json:"x_position"`
		YPosition float64 `json:"y_position"`
		Speed     float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		req := request{}
		err := context.ShouldBindJSON(&req)

		if err != nil {
			body := response{Success: false}
			context.AbortWithStatusJSON(http.StatusUnprocessableEntity, body)
			return
		}

		h.shark.Configure([2]float64{req.XPosition, req.YPosition}, req.Speed)

		body := response{Success: true}
		context.JSON(http.StatusOK, body)
	}
}

// PUT: /v1/prey

func (h *Handler) ConfigurePrey() gin.HandlerFunc {
	type request struct {
		Speed float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		req := request{}
		err := context.ShouldBindJSON(&req)

		if err != nil {
			body := response{Success: false}
			context.AbortWithStatusJSON(http.StatusUnprocessableEntity, body)
			return
		}

		h.prey.SetSpeed(req.Speed)

		body := response{Success: true}
		context.JSON(http.StatusOK, body)
	}
}

// POST: /v1/simulate

func (h *Handler) SimulateHunt() gin.HandlerFunc {
	type response struct {
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Time    float64 `json:"time"`
	}

	return func(context *gin.Context) {
		err, time := h.shark.Hunt(h.prey)

		if err != nil {
			body := response{
				Success: false,
				Message: err.Error(),
				Time:    time,
			}
			context.AbortWithStatusJSON(
				http.StatusConflict,
				body,
			)
			return
		}

		body := response{
			Success: true,
			Message: "hunted successfully",
			Time:    time,
		}

		context.JSON(http.StatusOK, body)
	}
}
