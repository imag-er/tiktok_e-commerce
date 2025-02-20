package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"src/api/def"
	"src/kitex_gen/user"
	service "src/kitex_gen/user/userservice"	
	"github.com/cloudwego/kitex/client"

)

func Register(ctx context.Context, apireq *app.RequestContext) {
	var rpcreq user.RegisterReq

	if err := apireq.BindAndValidate(&rpcreq); err != nil {
		// request data error
		apireq.JSON(consts.StatusBadRequest, utils.H{"message": err.Error()})
	} 

	rpcclient := service.MustNewClient("user",client.WithResolver(def.EtcdResolver))
	rpcresp, err := rpcclient.Register(ctx, &rpcreq)
	if err != nil {
		apireq.JSON(consts.StatusInternalServerError, utils.H{"message": err.Error()})
		return
	}
	
	apireq.JSON(consts.StatusOK, rpcresp)


}
