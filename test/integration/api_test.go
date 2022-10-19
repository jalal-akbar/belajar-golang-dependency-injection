package integration

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/app"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/controller"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/helper"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/middleware"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/model/domain"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/repository"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/service"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	var (
		driverName     = "mysql"
		dataSourceName = "root:root@tcp(localhost:3306)/for_test"
	)
	db, err := sql.Open(driverName, dataSourceName)
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	return db

}

func setupTestRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	// Controller
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	// http Router impl
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	// Set Up
	db := setupTestDB()
	truncateCategory(db)
	router := setupTestRouter(db)
	// Arrange
	var (
		assert     = assert.New(t)
		method     = http.MethodPost
		target     = "http://localhost:3001/api/categories/"
		bodyReader = strings.NewReader(`{"Name": "Gadget"}`)
	)
	request := httptest.NewRequest(method, target, bodyReader)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()
	// Act
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	err := json.Unmarshal(body, &responseBody)
	if err != nil {
		panic(err)
	}
	// Assert
	fmt.Println(&responseBody)
	assert.Equal(200, response.StatusCode)
	assert.Equal(200, int(responseBody["code"].(float64)))
	assert.Equal("OK", responseBody["status"])
	assert.Equal(1, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal("Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	// Set Up
	db := setupTestDB()
	truncateCategory(db)
	router := setupTestRouter(db)
	// Arracnge
	var (
		method      = http.MethodPost
		target      = "http://localhost:3001/api/categories/"
		requestBody = strings.NewReader(`{"Name": ""}`)
	)
	request := httptest.NewRequest(method, target, requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()
	// Act
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	data, _ := io.ReadAll(response.Body)
	var v map[string]interface{}
	json.Unmarshal(data, &v)
	// Assert
	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, "400 Bad Request", response.Status)
}

func TestUpdateCategorySuccess(t *testing.T) {
	// Set Up
	db := setupTestDB()
	truncateCategory(db)
	router := setupTestRouter(db)
	// Create Data For Update
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Gadget"})
	tx.Commit()
	// Arrange
	var (
		assert      = assert.New(t)
		method      = http.MethodPut
		target      = "http://localhost:3001/api/categories/"
		requestBody = strings.NewReader(`{"Name":"Computer"}`)
		v           map[string]interface{}
	)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, target+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	// Act
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	err := json.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}
	// Assert
	assert.Equal(200, response.StatusCode)
	assert.Equal(200, int(v["code"].(float64)))
	assert.Equal("OK", v["status"])
	assert.Equal(category.Id, int(v["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal("Computer", v["data"].(map[string]interface{})["name"])
}
func TestUpdateCategoryFailed(t *testing.T) {
	// Set Up
	db := setupTestDB()
	router := setupTestRouter(db)
	truncateCategory(db)
	// Create Data For Update
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Alcohol"})
	tx.Commit()
	// Arrange
	var (
		assert      = assert.New(t)
		method      = http.MethodPut
		target      = "http://localhost:3001/api/categories/"
		requestBody = strings.NewReader(`{"Name":""}`)
		v           map[string]interface{}
	)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, target+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	// Act
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	err := json.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}
	// Assert
	assert.Equal(400, response.StatusCode)
	assert.Equal(400, int(v["code"].(float64)))
	assert.Equal("Bad Request", v["status"])
}
func TestGetCategorySuccess(t *testing.T) {
	// Set Up
	db := setupTestDB()
	router := setupTestRouter(db)
	truncateCategory(db)
	// Create Category To Get Id
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Alcohol"})
	tx.Commit()
	// Arrange
	var (
		assert = assert.New(t)
		method = http.MethodGet
		target = "http://localhost:3001/api/categories/"
		// requestBody = strings.NewReader(`{"Name":"Alcohol"}`)
		v map[string]interface{}
	)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, target+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	// Act
	router.ServeHTTP(recorder, request)
	respponse := recorder.Result()
	body, _ := io.ReadAll(respponse.Body)
	err := json.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}
	// Assert
	assert.Equal(200, respponse.StatusCode)
	assert.Equal(200, int(v["code"].(float64)))
	assert.Equal("OK", v["status"])
	assert.Equal(category.Id, int(v["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(category.Name, v["data"].(map[string]interface{})["name"])

}
func TestGetCategoryFailed(t *testing.T) {
	// Set Up
	db := setupTestDB()
	router := setupTestRouter(db)
	truncateCategory(db)
	// Arrange
	var (
		assert = assert.New(t)
		method = http.MethodGet
		target = "http://locahost:3001/api/categories/404"
		v      map[string]interface{}
	)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, target, nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	// Act
	router.ServeHTTP(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}
	// Assert
	assert.Equal(404, resp.StatusCode)
	assert.Equal("Not Found", v["status"])
}
func TestDeleteCategorySuccess(t *testing.T) {
	// Set Up
	db := setupTestDB()
	router := setupTestRouter(db)
	truncateCategory(db)
	// Create Category For Delete
	tx, _ := db.Begin()
	CategoryRepository := repository.NewCategoryRepository()
	category := CategoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Alcohol",
	})
	tx.Commit()
	// Arrange
	var (
		assert = assert.New(t)
		method = http.MethodDelete
		target = "http://localhost:3001/api/categories/"
		v      map[string]interface{}
	)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, target+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	// Act
	router.ServeHTTP(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}
	// Assert
	assert.Equal(200, resp.StatusCode)
	assert.Equal("OK", v["status"])
}
func TestDeleteCategoryFailed(t *testing.T) {
	// Set Up
	db := setupTestDB()
	router := setupTestRouter(db)
	truncateCategory(db)
	// Arrange
	var (
		assert = assert.New(t)
		method = http.MethodDelete
		target = "http://localhost:3001/api/categories/404"
		v      map[string]interface{}
	)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, target, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	// Act
	router.ServeHTTP(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}
	// Assert
	assert.Equal(404, resp.StatusCode)
	assert.Equal("Not Found", v["status"])
}
func TestListCategoriesSuccess(t *testing.T) {
	// Set Up
	db := setupTestDB()
	router := setupTestRouter(db)
	truncateCategory(db)
	// Create Category For Get All Category
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category1 := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Alcohol"})
	category2 := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Herbal"})
	tx.Commit()
	// Arrange
	var (
		assert = assert.New(t)
		method = http.MethodGet
		target = "http://localhost:3001/api/categories/"
	)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, target, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	// Act
	router.ServeHTTP(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)
	var responseBody map[string]interface{}
	err := json.Unmarshal(body, &responseBody)
	if err != nil {
		panic(err)
	}
	var categories = responseBody["data"].([]interface{})
	var categoryResp1 = categories[0].(map[string]interface{})
	var categoryResp2 = categories[1].(map[string]interface{})
	// Assert
	fmt.Println(categories)
	fmt.Println(categoryResp1)
	fmt.Println(categoryResp2)
	assert.Equal(200, int(responseBody["code"].(float64)))
	assert.Equal("OK", responseBody["status"])
	assert.Equal(category1.Id, int(categoryResp1["id"].(float64)))
	assert.Equal(category1.Name, categoryResp1["name"])
	assert.Equal(category2.Id, int(categoryResp2["id"].(float64)))
	assert.Equal(category2.Name, categoryResp2["name"])

}
func TestUnauthorizedCategorySuccess(t *testing.T) {
	// Set Up
	db := setupTestDB()
	router := setupTestRouter(db)
	truncateCategory(db)
	// Arrange
	var (
		assert = assert.New(t)
		method = http.MethodGet
		target = "http://localhost:3001/api/categories"
		v      map[string]interface{}
	)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, target, nil)
	request.Header.Add("X-API-Key", "Tidak Rahasia")
	// Act
	router.ServeHTTP(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}
	// Assert
	assert.Equal(401, int(v["code"].(float64)))
	assert.Equal("Unauthorized", v["status"])
}
