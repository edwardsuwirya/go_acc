package api

import (
	"enigmacamp.com/goacc/delivery/appreq"
	"enigmacamp.com/goacc/usecase"
	"github.com/gin-gonic/gin"
)

type ProductApi struct {
	BaseApi
	productRegistrationUseCase usecase.ProductRegistrationUseCase
}

func (p *ProductApi) createProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productReq appreq.ProductRequest
		err := p.ParseRequestBody(c, &productReq)
		if err != nil {
			p.FailedRequest(c, "api-createProduct", "02", "Can not parse body")
			return
		}
		productRegistered, err := p.productRegistrationUseCase.Register(productReq.ProductCode, productReq.ProductName, "b4a467a2-fc43-4e18-85c2-02dab217a730")
		if err != nil {
			p.Failed(c, "api-createProduct", "01", "Can not create product")
			return
		}
		p.Success(c, "Product", productRegistered)
	}
}

func NewProductApi(productRoute *gin.RouterGroup, productReqistrationUseCase usecase.ProductRegistrationUseCase) {
	api := ProductApi{
		productRegistrationUseCase: productReqistrationUseCase,
	}
	productRoute.POST("", api.createProduct())
}
