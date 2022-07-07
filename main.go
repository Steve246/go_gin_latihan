package main

import (
	"go_gin_latihan/model"
	"go_gin_latihan/repository"
	"go_gin_latihan/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Healthy Check")
	})

	//default localhost:8080

	//POST Category
	router.POST("/product/PostCategory", PostCategory)
	// http://localhost:8080/product/PostCategory

	//bikin di body json
	// {
	//   "id": "C0002",
	//   "name": "Food"
	// }

	//GET Category
	router.GET("/product/GetCategory", Retreive)
	// http://localhost:8080/product/GetCategory


	//POST Produk

	err := router.Run()

	if err != nil {
		panic(err)
	}

}

// func GetCategory(c *gin.Context) {

// }

func Retreive(c *gin.Context) {
	// var newCategory *model.Category
	displayRepo := repository.NewCategoryRepository()

	ucDisplayAll := usecase.NewCreateDisplayCategory(displayRepo)

	hasilDisplay, err := ucDisplayAll.DisplayAll()


	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "Error when creating Product",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "list of products",
		"data":    hasilDisplay,
	})

}

func PostCategory(c *gin.Context) {
	var newCategory *model.Category

	categoryRepo := repository.NewCategoryRepository()

	categoryUc := usecase.NewCreateProductUseCase(categoryRepo)

	if err := c.BindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {

		err := categoryUc.CreateCategory(newCategory)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating Product",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": newCategory,
			})
		}
	}
}
