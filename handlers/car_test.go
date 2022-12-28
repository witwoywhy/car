package handlers

import (
	"bytes"
	"car/dtos"
	"car/orms"
	"car/repositories"
	"car/services"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var g = setupRouter()

func setupRouter() *gin.Engine {
	stage := os.Getenv("GO_TEST_STAGE")
	var r repositories.Car
	switch stage {
	case "mock":
		r = repositories.NewCarRepositoryMock()
	case "integration":
		dsn := "host=localhost user=postgres password=password dbname=integration port=5432 sslmode=disable"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		err = db.AutoMigrate(&orms.Car{})
		if err != nil {
			log.Fatal(err)
		}
		r = repositories.NewCarRepositoryDB(db)
	}
	s := services.NewCarService(r)
	h := Newhandler(s)

	g := gin.Default()
	g.POST("/car", h.Create)
	g.GET("/car", h.FindMany)
	g.GET("/car/:id", h.FindByID)
	g.PUT("/car/:id", h.Update)
	g.DELETE("/car/:id", h.Delete)

	return g
}

func TestCreate(t *testing.T) {
	want := dtos.Car{
		ID:       1,
		Brand:    "honda",
		Name:     "brv",
		Model:    "e",
		SubModel: "",
		Color:    "white",
		Price:    1000000,
	}

	data := map[string]any{
		"brand":    "honda",
		"name":     "brv",
		"model":    "e",
		"subModel": "",
		"color":    "white",
		"price":    1000000,
	}
	body, err := json.Marshal(data)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/car", bytes.NewBuffer(body))
	assert.Nil(t, err)
	g.ServeHTTP(w, req)

	var got dtos.Car
	json.Unmarshal(w.Body.Bytes(), &got)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, want, got)
}

func TestFindByID(t *testing.T) {
	want := dtos.Car{
		ID:       1,
		Brand:    "honda",
		Name:     "brv",
		Model:    "e",
		SubModel: "",
		Color:    "white",
		Price:    1000000,
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/car/1", nil)
	assert.Nil(t, err)
	g.ServeHTTP(w, req)

	var got dtos.Car
	json.Unmarshal(w.Body.Bytes(), &got)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want, got)
}

func TestFindMany(t *testing.T) {
	want := []dtos.Car{{
		ID:       1,
		Brand:    "honda",
		Name:     "brv",
		Model:    "e",
		SubModel: "",
		Color:    "white",
		Price:    1000000,
	}}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/car", nil)
	assert.Nil(t, err)
	g.ServeHTTP(w, req)

	var got []dtos.Car
	json.Unmarshal(w.Body.Bytes(), &got)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want, got)
}

func TestUpdate(t *testing.T) {
	want := dtos.Car{
		ID:       1,
		Brand:    "honda",
		Name:     "brv",
		Model:    "e",
		SubModel: "taffeta",
		Color:    "white",
		Price:    950000,
	}

	data := map[string]any{
		"brand":    "honda",
		"name":     "brv",
		"model":    "e",
		"subModel": "taffeta",
		"color":    "white",
		"price":    950000,
	}
	body, err := json.Marshal(data)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/car/1", bytes.NewBuffer(body))
	assert.Nil(t, err)
	g.ServeHTTP(w, req)

	var got dtos.Car
	json.Unmarshal(w.Body.Bytes(), &got)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want, got)
}

func TestDelete(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/car/1", nil)
	assert.Nil(t, err)
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
