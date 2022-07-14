package logic

import (
	"book/service/user/rpc/types/user"
	"context"
	"encoding/json"
	"fmt"

	"book/service/search/api/internal/svc"
	"book/service/search/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchReply, err error) {
	// todo: add your logic here and delete this line
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %v", userIdNumber) // 这里的key和生成jwt token时传入的key一致
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, err
	}

	//	使用user rpc
	_, err = l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.SearchReply{
		Name:  req.Name,
		Count: 100,
	}, nil
}
