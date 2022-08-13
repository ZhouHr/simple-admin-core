package logic

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/model"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *core.IDReq) (*core.BaseResp, error) {
	var children []*model.Menu
	check := l.svcCtx.DB.Where("parent_id = ?", in.ID).Find(&children)
	if check.Error != nil {
		return nil, status.Error(codes.Internal, check.Error.Error())
	}
	if check.RowsAffected != 0 {
		return nil, status.Error(codes.InvalidArgument, "sys.menu.deleteChildrenDesc")
	}
	result := l.svcCtx.DB.Delete(&model.Menu{
		Model: gorm.Model{ID: uint(in.ID)},
	})
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "sys.menu.menuNotExists")
	}
	return &core.BaseResp{Msg: "common.deleteSuccess"}, nil
}
