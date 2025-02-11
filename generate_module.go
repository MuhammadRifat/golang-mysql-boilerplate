package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generate_module.go <module_name>")
		os.Exit(1)
	}

	moduleName := os.Args[1]
	modulePath := fmt.Sprintf("src/modules/%s", moduleName)
	requestPath := fmt.Sprintf("%s/request", modulePath)
	modelPath := fmt.Sprintf("%s/model", modulePath)

	dirs := []string{modulePath, requestPath, modelPath}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			os.Exit(1)
		}
	}

	// Create files
	touchFiles := []string{
		fmt.Sprintf("%s/%s.request.go", requestPath, moduleName),
		fmt.Sprintf("%s/%s.model.go", modelPath, moduleName),
	}

	for _, file := range touchFiles {
		if _, err := os.Create(file); err != nil {
			fmt.Printf("Error creating file %s: %v\n", file, err)
			os.Exit(1)
		}
	}

	capModuleName := strings.ToUpper(moduleName[:1]) + moduleName[1:]

	// Create Router
	routerContent := fmt.Sprintf(`package %s

import "github.com/gin-gonic/gin"

func RoutesHandler(router *gin.RouterGroup) {
	group := router.Group("/%s")
	{
		group.POST("/", %sController.CreateOne)
		group.GET("/", %sController.GetAll)
		group.GET("/:id", %sController.GetOne)
		group.PUT("/:id", %sController.UpdateOne)
		group.DELETE("/:id", %sController.DeleteOne)
	}
}
`, moduleName, moduleName, capModuleName, capModuleName, capModuleName, capModuleName, capModuleName)

	writeToFile(fmt.Sprintf("%s/%s.router.go", modulePath, moduleName), routerContent)

	// Create Request
	requestContent := fmt.Sprintf(`package request

type %sRequest struct {
	Title string `+"`json:\"title\" binding:\"required,min=3,max=100\"`"+`
}

type Get%sRequest struct {
	Title string `+"`form:\"title\"`"+`
	Limit string `+"`form:\"limit\"`"+`
	Page  string `+"`form:\"page\"`"+`
}
`, capModuleName, capModuleName)

	writeToFile(fmt.Sprintf("%s/%s.request.go", requestPath, moduleName), requestContent)

	// Create Model
	modelContent := fmt.Sprintf(`package model

import "gorm.io/gorm"

type %s struct {
	gorm.Model
	Title string `+"`gorm:\"size:100;not null\"`"+`
}
`, capModuleName)

	writeToFile(fmt.Sprintf("%s/%s.model.go", modelPath, moduleName), modelContent)
}

func writeToFile(filename, content string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
		os.Exit(1)
	}
}
