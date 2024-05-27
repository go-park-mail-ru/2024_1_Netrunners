package api

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

type UsersService interface {
	CreateUser(ctx context.Context, user domain.UserSignUp) error
	RemoveUser(ctx context.Context, email string) error
	HasUser(ctx context.Context, email, password string) error
	GetUser(ctx context.Context, email string) (domain.User, error)
	ChangeUserPassword(ctx context.Context, email, newPassword string) (domain.User, error)
	ChangeUserName(ctx context.Context, email, newName string) (domain.User, error)
	GetUserDataByUuid(ctx context.Context, uuid string) (domain.User, error)
	GetUserPreview(ctx context.Context, uuid string) (domain.UserPreview, error)
	ChangeUserPasswordByUuid(ctx context.Context, uuid, newPassword string) (domain.User, error)
	ChangeUserNameByUuid(ctx context.Context, uuid, newName string) (domain.User, error)
	ChangeUserAvatarByUuid(ctx context.Context, uuid, newAvatar string) (domain.User, error)
	HasSubscription(ctx context.Context, uuid string) (bool, error)
	PaySubscription(ctx context.Context, uuid, subId string) (string, error)
	GetSubscriptions(ctx context.Context) ([]domain.Subscription, error)
	GetSubscription(ctx context.Context, uuid string) (domain.Subscription, error)
}

type UsersServer struct {
	usersService UsersService
	logger       *zap.SugaredLogger
}

func NewUsersServer(service UsersService, logger *zap.SugaredLogger) *UsersServer {
	return &UsersServer{
		usersService: service,
		logger:       logger,
	}
}

func (server *UsersServer) CreateUser(ctx context.Context,
	req *session.CreateUserRequest) (res *session.CreateUserResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.usersService.CreateUser(ctx, convertUserSignUpToRegular(req.User))
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to create user: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to create user: %v\n", requestId, err)
	}
	return res, nil
}

func (server *UsersServer) RemoveUser(ctx context.Context,
	req *session.RemoveUserRequest) (res *session.RemoveUserResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.usersService.RemoveUser(ctx, req.Login)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to remove user: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to remove user: %v\n", requestId, err)
	}
	return res, nil
}

func (server *UsersServer) HasUser(ctx context.Context,
	req *session.HasUserRequest) (res *session.HasUserResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.usersService.HasUser(ctx, req.Login, req.Password)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to has user: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to has user: %v\n", requestId, err)
	}
	return res, nil
}

func (server *UsersServer) GetUser(ctx context.Context,
	req *session.GetUserRequest) (res *session.GetUserResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	user, err := server.usersService.GetUser(ctx, req.Login)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get user: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get user: %v\n", requestId, err)
	}
	return &session.GetUserResponse{
		User: convertUserToProto(user),
	}, nil
}

func (server *UsersServer) ChangeUserPassword(ctx context.Context,
	req *session.ChangeUserPasswordRequest) (res *session.ChangeUserPasswordResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	user, err := server.usersService.ChangeUserPassword(ctx, req.Login, req.NewPassword)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to change user: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to change user: %v\n", requestId, err)
	}
	return &session.ChangeUserPasswordResponse{
		User: convertUserToProto(user),
	}, nil
}

func (server *UsersServer) ChangeUserName(ctx context.Context,
	req *session.ChangeUserNameRequest) (res *session.ChangeUserNameResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	user, err := server.usersService.ChangeUserName(ctx, req.Login, req.NewUsername)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to change user: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to change user: %v\n", requestId, err)
	}
	return &session.ChangeUserNameResponse{
		User: convertUserToProto(user),
	}, nil
}

func (server *UsersServer) GetUserDataByUuid(ctx context.Context,
	req *session.GetUserDataByUuidRequest) (res *session.GetUserDataByUuidResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	user, err := server.usersService.GetUserDataByUuid(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get user: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get user: %v\n", requestId, err)
	}
	return &session.GetUserDataByUuidResponse{
		User: convertUserToProto(user),
	}, nil
}

func (server *UsersServer) GetUserPreview(ctx context.Context,
	req *session.GetUserPreviewRequest) (res *session.GetUserPreviewResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	user, err := server.usersService.GetUserPreview(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get user: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get user: %v\n", requestId, err)
	}
	return &session.GetUserPreviewResponse{
		User: convertUserPreviewToProto(user),
	}, nil
}

func (server *UsersServer) ChangeUserPasswordByUuid(ctx context.Context,
	req *session.ChangeUserPasswordByUuidRequest) (res *session.ChangeUserPasswordByUuidResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	user, err := server.usersService.ChangeUserPasswordByUuid(ctx, req.Uuid, req.NewPassword)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to change user password: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to change user password: %v\n", requestId, err)
	}
	return &session.ChangeUserPasswordByUuidResponse{
		User: convertUserToProto(user),
	}, nil
}

func (server *UsersServer) ChangeUserNameByUuid(ctx context.Context,
	req *session.ChangeUserNameByUuidRequest) (res *session.ChangeUserNameByUuidResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	user, err := server.usersService.ChangeUserNameByUuid(ctx, req.Uuid, req.NewUsername)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to change username: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to change username: %v\n", requestId, err)
	}
	return &session.ChangeUserNameByUuidResponse{
		User: convertUserToProto(user),
	}, nil
}

func (server *UsersServer) ChangeUserAvatarByUuid(ctx context.Context,
	req *session.ChangeUserAvatarByUuidRequest) (res *session.ChangeUserAvatarByUuidResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	user, err := server.usersService.ChangeUserAvatarByUuid(ctx, req.Uuid, req.NewAvatar)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to change user avatar: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to change user avatar: %v\n", requestId, err)
	}
	return &session.ChangeUserAvatarByUuidResponse{
		User: convertUserToProto(user),
	}, nil
}

func (server *UsersServer) HasSubscription(ctx context.Context,
	req *session.HasSubscriptionRequest) (*session.HasSubscriptionResponse, error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	stat, err := server.usersService.HasSubscription(ctx, req.Uuid)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to check user subscription: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to check user subscription: %v\n", requestId, err)
	}
	return &session.HasSubscriptionResponse{
		Status: stat,
	}, nil
}

func (server *UsersServer) GetSubscriptions(ctx context.Context,
	req *session.GetSubscriptionsRequest) (*session.GetSubscriptionsResponse, error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	subs, err := server.usersService.GetSubscriptions(ctx)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to get subs: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to get subs: %v\n", requestId, err)
	}
	return &session.GetSubscriptionsResponse{
		Subscriptions: convertSubsToProto(subs),
	}, nil
}

func (server *UsersServer) PaySubscription(ctx context.Context,
	req *session.PaySubscriptionRequest) (*session.PaySubscriptionResponse, error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	response, err := server.usersService.PaySubscription(ctx, req.Uuid, req.SubId)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to pay: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to pay: %v\n", requestId, err)
	}
	return &session.PaySubscriptionResponse{
		PaymentResponse: response,
	}, nil
}

func convertUserSignUpToRegular(user *session.UserSignUp) domain.UserSignUp {
	return domain.UserSignUp{
		Email:    user.Email,
		Name:     user.Username,
		Password: user.Password,
	}
}

func convertUserToProto(user domain.User) *session.User {
	return &session.User{
		Uuid:            user.Uuid,
		Email:           user.Email,
		Username:        user.Name,
		Password:        user.Password,
		IsAdmin:         user.IsAdmin,
		Version:         user.Version,
		RegisteredAt:    convertTimeToProto(user.RegisteredAt),
		Birthday:        convertTimeToProto(user.Birthday),
		Avatar:          user.Avatar,
		HasSubscription: user.HasSubscription,
	}
}

func convertTimeToProto(time time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: time.Unix(),
		Nanos:   int32(time.Nanosecond()),
	}
}

func convertUserPreviewToProto(user domain.UserPreview) *session.UserPreview {
	return &session.UserPreview{
		Uuid:     user.Uuid,
		Username: user.Name,
		Avatar:   user.Avatar,
	}
}

func convertSubsToProto(subs []domain.Subscription) []*session.Subscription {
	protoSubs := make([]*session.Subscription, 0, len(subs))
	for _, sub := range subs {
		protoSubs = append(protoSubs, &session.Subscription{
			Uuid:        sub.Uuid,
			Title:       sub.Title,
			Description: sub.Description,
			Price:       sub.Amount,
		})
	}
	return protoSubs
}
