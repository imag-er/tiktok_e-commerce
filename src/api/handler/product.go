package handler

import (
	"context"
	"src/api/def"
	"src/kitex_gen/product"
	service "src/kitex_gen/product/productcatalogservice"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	"log"

	"github.com/cloudwego/kitex/client"
)



func CreateProduct(ctx context.Context, apireq *app.RequestContext) {
	var rpcreq product.CreateProductReq

	if err := apireq.BindAndValidate(&rpcreq); err != nil {
		// request data error
		apireq.JSON(consts.StatusBadRequest, utils.H{"message": "param error"})
		return
	}
	
	log.Printf("CreateProduct: %v", rpcreq)
	rpcclient := service.MustNewClient("product",client.WithResolver(def.EtcdResolver))
	rpcresp, err := rpcclient.CreateProduct(ctx, &rpcreq)
	if err != nil {
		apireq.JSON(consts.StatusInternalServerError, utils.H{"message": err.Error()})
		return
	}
	 
	apireq.JSON(consts.StatusOK, rpcresp)
}

func DeleteProduct(ctx context.Context, apireq *app.RequestContext) {
	var rpcreq product.DeleteProductReq

	id, err := strconv.Atoi(apireq.Param("id"))
	if err != nil {
		apireq.JSON(consts.StatusBadRequest, utils.H{"message": "param error"})
		return
	}
	rpcreq.Id = uint32(id)

	rpcclient := service.MustNewClient("product",client.WithResolver(def.EtcdResolver))
	rpcresp, err := rpcclient.DeleteProduct(ctx, &rpcreq)
	if err != nil {
		apireq.JSON(consts.StatusInternalServerError, utils.H{"message": err.Error()})
		return
	}
	 
	apireq.JSON(consts.StatusOK, rpcresp)
}

func GetProduct(ctx context.Context, apireq *app.RequestContext) {

	var rpcreq product.GetProductReq
	
	id, err := strconv.Atoi(apireq.Param("id"))
	if err != nil {
		apireq.JSON(consts.StatusBadRequest, utils.H{"message": "invalid id"})
		return
	}
	
	rpcreq.Id = uint32(id)

	rpcclient := service.MustNewClient("product",client.WithResolver(def.EtcdResolver))
	rpcresp, err := rpcclient.GetProduct(ctx, &rpcreq)
	if err != nil {
		apireq.JSON(consts.StatusInternalServerError, utils.H{"message": err.Error()})
		return
	}
	 
	apireq.JSON(consts.StatusOK, rpcresp)
}
