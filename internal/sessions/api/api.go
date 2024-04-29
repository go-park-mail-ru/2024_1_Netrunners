package api

import (
	"context"
	"fmt"
	"go.uber.org/zap"

	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

type SessionService interface {
	Add(ctx context.Context, login string, token string, version uint32) (err error)
	DeleteSession(ctx context.Context, login string, token string) (err error)
	Update(ctx context.Context, login string, token string) (err error)
	CheckVersion(ctx context.Context, login string, token string, usersVersion uint32) (hasSession bool, err error)
	GetVersion(ctx context.Context, login string, token string) (version uint32, err error)
	HasSession(ctx context.Context, login string, token string) error
	CheckAllUserSessionTokens(ctx context.Context, login string) error
	GenerateTokens(login string, isAdmin bool, version uint32) (tokenSigned string, err error)
}

type SessionSever struct {
	sessionsService SessionService
	logger          *zap.SugaredLogger
}

func NewSessionServer(sessionsService SessionService, logger *zap.SugaredLogger) *SessionSever {
	return &SessionSever{
		sessionsService: sessionsService,
		logger:          logger,
	}
}

func (server *SessionSever) Add(ctx context.Context, req *session.AddRequest) (res *session.AddResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.sessionsService.Add(ctx, req.Login, req.Token, req.Version)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to add session: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to add session: %v\n", requestId, err)
	}
	return &session.AddResponse{}, nil
}

func (server *SessionSever) DeleteSession(ctx context.Context,
	req *session.DeleteSessionRequest) (res *session.DeleteSessionResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.sessionsService.DeleteSession(ctx, req.Login, req.Token)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to delete session: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to delete session: %v\n", requestId, err)
	}
	return &session.DeleteSessionResponse{}, nil
}

func (server *SessionSever) Update(ctx context.Context,
	req *session.UpdateRequest) (res *session.UpdateRequestResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.sessionsService.Update(ctx, req.Login, req.Token)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to update session: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to update session: %v\n", requestId, err)
	}
	return &session.UpdateRequestResponse{}, nil
}

func (server *SessionSever) CheckVersion(ctx context.Context,
	req *session.CheckVersionRequest) (res *session.CheckVersionResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	has, err := server.sessionsService.CheckVersion(ctx, req.Login, req.Token, req.Version)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to check session version: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to check session version: %v\n", requestId, err)
	}
	return &session.CheckVersionResponse{
		HasSession: has,
	}, nil
}

func (server *SessionSever) GetVersion(ctx context.Context,
	req *session.GetVersionRequest) (res *session.GetVersionResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	version, err := server.sessionsService.GetVersion(ctx, req.Login, req.Token)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to check session version: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to check session version: %v\n", requestId, err)
	}
	return &session.GetVersionResponse{
		Version: version,
	}, nil
}

func (server *SessionSever) HasSession(ctx context.Context,
	req *session.HasSessionRequest) (res *session.HasSessionResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.sessionsService.HasSession(ctx, req.Login, req.Token)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to check session: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to check session: %v\n", requestId, err)
	}
	return &session.HasSessionResponse{}, nil
}

func (server *SessionSever) CheckAllUserSessionTokens(ctx context.Context,
	req *session.CheckAllUserSessionTokensRequest) (res *session.CheckAllUserSessionTokensResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	err = server.sessionsService.CheckAllUserSessionTokens(ctx, req.Login)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to check all user session tokens: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to check all user session tokens: %v\n", requestId, err)
	}
	return &session.CheckAllUserSessionTokensResponse{}, nil
}

func (server *SessionSever) GenerateToken(ctx context.Context,
	req *session.GenerateTokenRequest) (res *session.GenerateTokenResponse, err error) {
	requestId := ctx.Value(reqid.ReqIDKey)
	token, err := server.sessionsService.GenerateTokens(req.Login, req.IsAdmin, req.Version)
	if err != nil {
		server.logger.Errorf("[reqid=%s] failed to generate token: %v\n", requestId, err)
		return nil, fmt.Errorf("[reqid=%s] failed to generate token: %v\n", requestId, err)
	}
	return &session.GenerateTokenResponse{
		TokenSigned: token,
	}, nil
}
