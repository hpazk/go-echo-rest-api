package controllers

import (
	"net/http"

	"github.com/hpazk/go-echo-rest-api/app/database"
	"github.com/hpazk/go-echo-rest-api/app/helpers"
	"github.com/hpazk/go-echo-rest-api/app/middlewares"
	ProductModel "github.com/hpazk/go-echo-rest-api/app/models/products"
	"github.com/labstack/echo/v4"
)

type (
	ProductsController struct {
	}

	AddProductRequest struct {
		Name          string `json:"name" validate:"required"`
		Description   string `json:"description" validate:"required"`
		Price         int    `json:"price" validate:"required"`
		ProductRating int    `json:"product_rating"`
		PicturePath   string `json:"picture_path" validate:"required"`
	}
)

func (controller ProductsController) Routes() []helpers.Route {
	return []helpers.Route{
		{
			Method:     echo.GET,
			Path:       "/products",
			Handler:    controller.GetProducts,
			Middleware: []echo.MiddlewareFunc{middlewares.JWTMiddleWare()},
		},
		{
			Method:     echo.POST,
			Path:       "/products",
			Handler:    controller.AddProduct,
			Middleware: []echo.MiddlewareFunc{middlewares.JWTMiddleWare()},
		},
		// {
		// 	Method: echo.GET,
		// 	Path:   "/blogs",
		// 	// Handler: controller.GetBlogs,
		// },
		// {
		// 	Method: echo.GET,
		// 	Path:   "/blog/:blogId",
		// 	// Handler: controller.GetBlog,
		// },
	}
}

func (controller ProductsController) GetProducts(c echo.Context) error {
	// db := database.GetInstance()
	// var product []ProductModel.Product
	// db.Find(&product)
	return c.String(http.StatusOK, "product")

	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(*middlewares.JWTCustomClaims)
	// // if claims.Role == "CUSTOMER" {
	// // 	db := database.GetInstance()
	// // 	var products []productModel.Product
	// // 	db.Find(&products)
	// // 	return c.JSON(http.StatusOK, products)
	// // }
	// return c.JSON(http.StatusOK, claims.Role)
}

func (controller ProductsController) AddProduct(c echo.Context) error {
	params := new(AddProductRequest)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := c.Validate(params); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := database.GetInstance()
	var product ProductModel.Product
	product.Name = params.Name
	product.Description = params.Description
	product.Price = params.Price
	product.ProductRating = params.ProductRating
	product.PicturePath = params.PicturePath
	db.Create(&product)

	return c.JSON(http.StatusOK, product)
}
