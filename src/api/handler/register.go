package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"src/api/def"
	"src/kitex_gen/user"
	"src/kitex_gen/user/userservice"	
	"github.com/cloudwego/kitex/client"

)

func Register(ctx context.Context, req *app.RequestContext) {
	var loginVals def.Register

	if err := req.BindAndValidate(&loginVals); err != nil {
		// request data error
		req.JSON(consts.StatusBadRequest, utils.H{"message": err.Error()})
	} 

	rpcreq := user.RegisterReq{
		Email: loginVals.Email,
		Password: loginVals.Password,
		Username: loginVals.Username,
	}

	rpcclient := userservice.MustNewClient("user",client.WithResolver(def.EtcdResolver))
	rpcresp, err := rpcclient.Register(ctx, &rpcreq)
	if err != nil {
		req.JSON(consts.StatusInternalServerError, utils.H{"message": err.Error()})
		return
	}
	
	req.JSON(consts.StatusOK, rpcresp)


}
