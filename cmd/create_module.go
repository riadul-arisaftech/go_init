package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var module name string
	var moduleName string

	//flag.StringVar(&modelName, "model", "", "GORM model name")
	flag.StringVar(&moduleName, "name", "example", "Module name")
	flag.Parse()

	modName := capitalizeFirstLetter(strings.ToLower(moduleName))
	handlerInterfaceName := fmt.Sprintf("I%sHandlers", modName)
	serviceInterfaceName := fmt.Sprintf("I%sService", modName)

	interfacesContent := getInterfaceContent(handlerInterfaceName, serviceInterfaceName)
	handlerContent := getHandlerContent(modName, handlerInterfaceName, serviceInterfaceName)
	serviceContent := getServiceContent(modName, serviceInterfaceName)
	routeContent := getRouteContent(modName)

	createInterfaceContent(moduleName, interfacesContent)
	createHandlerContent(moduleName, handlerContent)
	createServiceContent(moduleName, serviceContent)
	createRouteContent(moduleName, routeContent)

	fmt.Printf("all files success created")
}

func createInterfaceContent(name, content string) {
	// Create interface file
	fileName := fmt.Sprintf("%s/%s_interface.go", "./api/interfaces", name)
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating interface file: %v\n", err)
		os.Exit(1)
	}
}

func createHandlerContent(name, content string) {
	// Create interface file
	fileName := fmt.Sprintf("%s/%s_handler.go", "./api/handlers", name)
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating handler file: %v\n", err)
		os.Exit(1)
	}
}

func createServiceContent(name, content string) {
	// Create service file
	fileName := fmt.Sprintf("%s/%s_service.go", "./api/services", name)
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating service file: %v\n", err)
		os.Exit(1)
	}
}

func createRouteContent(name, content string) {
	// Create route file
	fileName := fmt.Sprintf("%s/%s_routes.go", "./api/routes", name)
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating routes file: %v\n", err)
		os.Exit(1)
	}
}

func getInterfaceContent(handler, service string) string {
	return fmt.Sprintf(
		`package interfaces

import "github.com/labstack/echo/v4"

type %s interface {
	Demo(ctx echo.Context) error
}

type %s interface {
	CreateDemo(params any) error
}`, handler, service)
}

func getHandlerContent(modeName, iHandlerName, iServiceName string) string {
	return fmt.Sprintf(
		`package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/go_sample/api/interfaces"
)

type %sHandlers struct {
	validator   *validator.Validate
	service interfaces.%s
}

func New%sHandler(val *validator.Validate, as interfaces.%s) interfaces.%s {
	return &%sHandlers{
		validator:   val,
		service: as,
	}
}

func (h %sHandlers) Demo(ctx echo.Context) error {
	fmt.Printf("hello world")
	return nil
}`, modeName, iServiceName, modeName, iServiceName, iHandlerName, modeName, modeName)
}

func getServiceContent(modeName, iServiceName string) string {
	return fmt.Sprintf(
		`package services

import "github.com/go_sample/api/interfaces"

type %sService struct {
}

func New%sService() interfaces.%s {
	return &%sService{}
}

func (s %sService) CreateDemo(params any) error {
	return nil
}`, modeName, modeName, iServiceName, modeName, modeName)
}

func getRouteContent(modeName string) string {
	return fmt.Sprintf(`package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/go_sample/api/handlers"
	"github.com/go_sample/api/services"
)

func %sRoutes(route *echo.Group) {
	val := validator.New()
	service := services.New%sService()
	handlers := handlers.New%sHandler(val, service)

	// routes
	route.GET("/", handlers.Demo)
}`, modeName, modeName, modeName)
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
