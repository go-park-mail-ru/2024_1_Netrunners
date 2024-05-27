// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package session

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FilmsClient is the client API for Films service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilmsClient interface {
	GetAllFilmsPreviews(ctx context.Context, in *AllFilmsPreviewsRequest, opts ...grpc.CallOption) (*AllFilmsPreviewsResponse, error)
	GetFilmDataByUuid(ctx context.Context, in *FilmDataByUuidRequest, opts ...grpc.CallOption) (*FilmDataByUuidResponse, error)
	GetFilmPreviewByUuid(ctx context.Context, in *FilmPreviewByUuidRequest, opts ...grpc.CallOption) (*FilmPreviewByUuidResponse, error)
	RemoveFilmByUuid(ctx context.Context, in *RemoveFilmByUuidRequest, opts ...grpc.CallOption) (*RemoveFilmByUuidResponse, error)
	GetActorDataByUuid(ctx context.Context, in *ActorDataByUuidRequest, opts ...grpc.CallOption) (*ActorDataByUuidResponse, error)
	GetActorsByFilm(ctx context.Context, in *ActorsByFilmRequest, opts ...grpc.CallOption) (*ActorsByFilmResponse, error)
	PutFavorite(ctx context.Context, in *PutFavoriteRequest, opts ...grpc.CallOption) (*PutFavoriteResponse, error)
	DeleteFavorite(ctx context.Context, in *DeleteFavoriteRequest, opts ...grpc.CallOption) (*DeleteFavoriteResponse, error)
	GetAllFavoriteFilms(ctx context.Context, in *GetAllFavoriteFilmsRequest, opts ...grpc.CallOption) (*GetAllFavoriteFilmsResponse, error)
	GetAllFilmsByGenre(ctx context.Context, in *GetAllFilmsByGenreRequest, opts ...grpc.CallOption) (*GetAllFilmsByGenreResponse, error)
	GetAllGenres(ctx context.Context, in *GetAllGenresRequest, opts ...grpc.CallOption) (*GetAllGenresResponse, error)
	AddFilm(ctx context.Context, in *AddFilmRequest, opts ...grpc.CallOption) (*AddFilmResponse, error)
	FindFilmsShort(ctx context.Context, in *FindFilmsShortRequest, opts ...grpc.CallOption) (*FindFilmsShortResponse, error)
	FindFilmsLong(ctx context.Context, in *FindFilmsShortRequest, opts ...grpc.CallOption) (*FindFilmsLongResponse, error)
	FindSerialsShort(ctx context.Context, in *FindFilmsShortRequest, opts ...grpc.CallOption) (*FindFilmsShortResponse, error)
	FindSerialsLong(ctx context.Context, in *FindFilmsShortRequest, opts ...grpc.CallOption) (*FindFilmsLongResponse, error)
	FindActorsShort(ctx context.Context, in *FindActorsShortRequest, opts ...grpc.CallOption) (*FindActorsShortResponse, error)
	FindActorsLong(ctx context.Context, in *FindActorsShortRequest, opts ...grpc.CallOption) (*FindActorsLongResponse, error)
	GetTopFilms(ctx context.Context, in *GetTopFilmsRequest, opts ...grpc.CallOption) (*GetTopFilmsResponse, error)
	GetAllFilmComments(ctx context.Context, in *AllFilmCommentsRequest, opts ...grpc.CallOption) (*AllFilmCommentsResponse, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error)
	RemoveComment(ctx context.Context, in *RemoveCommentRequest, opts ...grpc.CallOption) (*RemoveCommentResponse, error)
}

type filmsClient struct {
	cc grpc.ClientConnInterface
}

func NewFilmsClient(cc grpc.ClientConnInterface) FilmsClient {
	return &filmsClient{cc}
}

func (c *filmsClient) GetAllFilmsPreviews(ctx context.Context, in *AllFilmsPreviewsRequest, opts ...grpc.CallOption) (*AllFilmsPreviewsResponse, error) {
	out := new(AllFilmsPreviewsResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetAllFilmsPreviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) GetFilmDataByUuid(ctx context.Context, in *FilmDataByUuidRequest, opts ...grpc.CallOption) (*FilmDataByUuidResponse, error) {
	out := new(FilmDataByUuidResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetFilmDataByUuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) GetFilmPreviewByUuid(ctx context.Context, in *FilmPreviewByUuidRequest, opts ...grpc.CallOption) (*FilmPreviewByUuidResponse, error) {
	out := new(FilmPreviewByUuidResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetFilmPreviewByUuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) RemoveFilmByUuid(ctx context.Context, in *RemoveFilmByUuidRequest, opts ...grpc.CallOption) (*RemoveFilmByUuidResponse, error) {
	out := new(RemoveFilmByUuidResponse)
	err := c.cc.Invoke(ctx, "/session.Films/RemoveFilmByUuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) GetActorDataByUuid(ctx context.Context, in *ActorDataByUuidRequest, opts ...grpc.CallOption) (*ActorDataByUuidResponse, error) {
	out := new(ActorDataByUuidResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetActorDataByUuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) GetActorsByFilm(ctx context.Context, in *ActorsByFilmRequest, opts ...grpc.CallOption) (*ActorsByFilmResponse, error) {
	out := new(ActorsByFilmResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetActorsByFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) PutFavorite(ctx context.Context, in *PutFavoriteRequest, opts ...grpc.CallOption) (*PutFavoriteResponse, error) {
	out := new(PutFavoriteResponse)
	err := c.cc.Invoke(ctx, "/session.Films/PutFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) DeleteFavorite(ctx context.Context, in *DeleteFavoriteRequest, opts ...grpc.CallOption) (*DeleteFavoriteResponse, error) {
	out := new(DeleteFavoriteResponse)
	err := c.cc.Invoke(ctx, "/session.Films/DeleteFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) GetAllFavoriteFilms(ctx context.Context, in *GetAllFavoriteFilmsRequest, opts ...grpc.CallOption) (*GetAllFavoriteFilmsResponse, error) {
	out := new(GetAllFavoriteFilmsResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetAllFavoriteFilms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) GetAllFilmsByGenre(ctx context.Context, in *GetAllFilmsByGenreRequest, opts ...grpc.CallOption) (*GetAllFilmsByGenreResponse, error) {
	out := new(GetAllFilmsByGenreResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetAllFilmsByGenre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) GetAllGenres(ctx context.Context, in *GetAllGenresRequest, opts ...grpc.CallOption) (*GetAllGenresResponse, error) {
	out := new(GetAllGenresResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetAllGenres", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) AddFilm(ctx context.Context, in *AddFilmRequest, opts ...grpc.CallOption) (*AddFilmResponse, error) {
	out := new(AddFilmResponse)
	err := c.cc.Invoke(ctx, "/session.Films/AddFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) FindFilmsShort(ctx context.Context, in *FindFilmsShortRequest, opts ...grpc.CallOption) (*FindFilmsShortResponse, error) {
	out := new(FindFilmsShortResponse)
	err := c.cc.Invoke(ctx, "/session.Films/FindFilmsShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) FindFilmsLong(ctx context.Context, in *FindFilmsShortRequest, opts ...grpc.CallOption) (*FindFilmsLongResponse, error) {
	out := new(FindFilmsLongResponse)
	err := c.cc.Invoke(ctx, "/session.Films/FindFilmsLong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) FindSerialsShort(ctx context.Context, in *FindFilmsShortRequest, opts ...grpc.CallOption) (*FindFilmsShortResponse, error) {
	out := new(FindFilmsShortResponse)
	err := c.cc.Invoke(ctx, "/session.Films/FindSerialsShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) FindSerialsLong(ctx context.Context, in *FindFilmsShortRequest, opts ...grpc.CallOption) (*FindFilmsLongResponse, error) {
	out := new(FindFilmsLongResponse)
	err := c.cc.Invoke(ctx, "/session.Films/FindSerialsLong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) FindActorsShort(ctx context.Context, in *FindActorsShortRequest, opts ...grpc.CallOption) (*FindActorsShortResponse, error) {
	out := new(FindActorsShortResponse)
	err := c.cc.Invoke(ctx, "/session.Films/FindActorsShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) FindActorsLong(ctx context.Context, in *FindActorsShortRequest, opts ...grpc.CallOption) (*FindActorsLongResponse, error) {
	out := new(FindActorsLongResponse)
	err := c.cc.Invoke(ctx, "/session.Films/FindActorsLong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) GetTopFilms(ctx context.Context, in *GetTopFilmsRequest, opts ...grpc.CallOption) (*GetTopFilmsResponse, error) {
	out := new(GetTopFilmsResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetTopFilms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) GetAllFilmComments(ctx context.Context, in *AllFilmCommentsRequest, opts ...grpc.CallOption) (*AllFilmCommentsResponse, error) {
	out := new(AllFilmCommentsResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetAllFilmComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error) {
	out := new(AddCommentResponse)
	err := c.cc.Invoke(ctx, "/session.Films/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsClient) RemoveComment(ctx context.Context, in *RemoveCommentRequest, opts ...grpc.CallOption) (*RemoveCommentResponse, error) {
	out := new(RemoveCommentResponse)
	err := c.cc.Invoke(ctx, "/session.Films/RemoveComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilmsServer is the server API for Films service.
// All implementations must embed UnimplementedFilmsServer
// for forward compatibility
type FilmsServer interface {
	GetAllFilmsPreviews(context.Context, *AllFilmsPreviewsRequest) (*AllFilmsPreviewsResponse, error)
	GetFilmDataByUuid(context.Context, *FilmDataByUuidRequest) (*FilmDataByUuidResponse, error)
	GetFilmPreviewByUuid(context.Context, *FilmPreviewByUuidRequest) (*FilmPreviewByUuidResponse, error)
	RemoveFilmByUuid(context.Context, *RemoveFilmByUuidRequest) (*RemoveFilmByUuidResponse, error)
	GetActorDataByUuid(context.Context, *ActorDataByUuidRequest) (*ActorDataByUuidResponse, error)
	GetActorsByFilm(context.Context, *ActorsByFilmRequest) (*ActorsByFilmResponse, error)
	PutFavorite(context.Context, *PutFavoriteRequest) (*PutFavoriteResponse, error)
	DeleteFavorite(context.Context, *DeleteFavoriteRequest) (*DeleteFavoriteResponse, error)
	GetAllFavoriteFilms(context.Context, *GetAllFavoriteFilmsRequest) (*GetAllFavoriteFilmsResponse, error)
	GetAllFilmsByGenre(context.Context, *GetAllFilmsByGenreRequest) (*GetAllFilmsByGenreResponse, error)
	GetAllGenres(context.Context, *GetAllGenresRequest) (*GetAllGenresResponse, error)
	AddFilm(context.Context, *AddFilmRequest) (*AddFilmResponse, error)
	FindFilmsShort(context.Context, *FindFilmsShortRequest) (*FindFilmsShortResponse, error)
	FindFilmsLong(context.Context, *FindFilmsShortRequest) (*FindFilmsLongResponse, error)
	FindSerialsShort(context.Context, *FindFilmsShortRequest) (*FindFilmsShortResponse, error)
	FindSerialsLong(context.Context, *FindFilmsShortRequest) (*FindFilmsLongResponse, error)
	FindActorsShort(context.Context, *FindActorsShortRequest) (*FindActorsShortResponse, error)
	FindActorsLong(context.Context, *FindActorsShortRequest) (*FindActorsLongResponse, error)
	GetTopFilms(context.Context, *GetTopFilmsRequest) (*GetTopFilmsResponse, error)
	GetAllFilmComments(context.Context, *AllFilmCommentsRequest) (*AllFilmCommentsResponse, error)
	AddComment(context.Context, *AddCommentRequest) (*AddCommentResponse, error)
	RemoveComment(context.Context, *RemoveCommentRequest) (*RemoveCommentResponse, error)
}

// UnimplementedFilmsServer must be embedded to have forward compatible implementations.
type UnimplementedFilmsServer struct {
}

func (UnimplementedFilmsServer) GetAllFilmsPreviews(context.Context, *AllFilmsPreviewsRequest) (*AllFilmsPreviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFilmsPreviews not implemented")
}
func (UnimplementedFilmsServer) GetFilmDataByUuid(context.Context, *FilmDataByUuidRequest) (*FilmDataByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilmDataByUuid not implemented")
}
func (UnimplementedFilmsServer) GetFilmPreviewByUuid(context.Context, *FilmPreviewByUuidRequest) (*FilmPreviewByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilmPreviewByUuid not implemented")
}
func (UnimplementedFilmsServer) RemoveFilmByUuid(context.Context, *RemoveFilmByUuidRequest) (*RemoveFilmByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFilmByUuid not implemented")
}
func (UnimplementedFilmsServer) GetActorDataByUuid(context.Context, *ActorDataByUuidRequest) (*ActorDataByUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActorDataByUuid not implemented")
}
func (UnimplementedFilmsServer) GetActorsByFilm(context.Context, *ActorsByFilmRequest) (*ActorsByFilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActorsByFilm not implemented")
}
func (UnimplementedFilmsServer) PutFavorite(context.Context, *PutFavoriteRequest) (*PutFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutFavorite not implemented")
}
func (UnimplementedFilmsServer) DeleteFavorite(context.Context, *DeleteFavoriteRequest) (*DeleteFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFavorite not implemented")
}
func (UnimplementedFilmsServer) GetAllFavoriteFilms(context.Context, *GetAllFavoriteFilmsRequest) (*GetAllFavoriteFilmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFavoriteFilms not implemented")
}
func (UnimplementedFilmsServer) GetAllFilmsByGenre(context.Context, *GetAllFilmsByGenreRequest) (*GetAllFilmsByGenreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFilmsByGenre not implemented")
}
func (UnimplementedFilmsServer) GetAllGenres(context.Context, *GetAllGenresRequest) (*GetAllGenresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGenres not implemented")
}
func (UnimplementedFilmsServer) AddFilm(context.Context, *AddFilmRequest) (*AddFilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFilm not implemented")
}
func (UnimplementedFilmsServer) FindFilmsShort(context.Context, *FindFilmsShortRequest) (*FindFilmsShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFilmsShort not implemented")
}
func (UnimplementedFilmsServer) FindFilmsLong(context.Context, *FindFilmsShortRequest) (*FindFilmsLongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFilmsLong not implemented")
}
func (UnimplementedFilmsServer) FindSerialsShort(context.Context, *FindFilmsShortRequest) (*FindFilmsShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSerialsShort not implemented")
}
func (UnimplementedFilmsServer) FindSerialsLong(context.Context, *FindFilmsShortRequest) (*FindFilmsLongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSerialsLong not implemented")
}
func (UnimplementedFilmsServer) FindActorsShort(context.Context, *FindActorsShortRequest) (*FindActorsShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindActorsShort not implemented")
}
func (UnimplementedFilmsServer) FindActorsLong(context.Context, *FindActorsShortRequest) (*FindActorsLongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindActorsLong not implemented")
}
func (UnimplementedFilmsServer) GetTopFilms(context.Context, *GetTopFilmsRequest) (*GetTopFilmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopFilms not implemented")
}
func (UnimplementedFilmsServer) GetAllFilmComments(context.Context, *AllFilmCommentsRequest) (*AllFilmCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFilmComments not implemented")
}
func (UnimplementedFilmsServer) AddComment(context.Context, *AddCommentRequest) (*AddCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedFilmsServer) RemoveComment(context.Context, *RemoveCommentRequest) (*RemoveCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveComment not implemented")
}
func (UnimplementedFilmsServer) mustEmbedUnimplementedFilmsServer() {}

// UnsafeFilmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilmsServer will
// result in compilation errors.
type UnsafeFilmsServer interface {
	mustEmbedUnimplementedFilmsServer()
}

func RegisterFilmsServer(s grpc.ServiceRegistrar, srv FilmsServer) {
	s.RegisterService(&Films_ServiceDesc, srv)
}

func _Films_GetAllFilmsPreviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllFilmsPreviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetAllFilmsPreviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/GetAllFilmsPreviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetAllFilmsPreviews(ctx, req.(*AllFilmsPreviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_GetFilmDataByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmDataByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetFilmDataByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/GetFilmDataByUuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetFilmDataByUuid(ctx, req.(*FilmDataByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_GetFilmPreviewByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmPreviewByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetFilmPreviewByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/GetFilmPreviewByUuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetFilmPreviewByUuid(ctx, req.(*FilmPreviewByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_RemoveFilmByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFilmByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).RemoveFilmByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/RemoveFilmByUuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).RemoveFilmByUuid(ctx, req.(*RemoveFilmByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_GetActorDataByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActorDataByUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetActorDataByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/GetActorDataByUuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetActorDataByUuid(ctx, req.(*ActorDataByUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_GetActorsByFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActorsByFilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetActorsByFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/GetActorsByFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetActorsByFilm(ctx, req.(*ActorsByFilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_PutFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).PutFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/PutFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).PutFavorite(ctx, req.(*PutFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_DeleteFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).DeleteFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/DeleteFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).DeleteFavorite(ctx, req.(*DeleteFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_GetAllFavoriteFilms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllFavoriteFilmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetAllFavoriteFilms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/GetAllFavoriteFilms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetAllFavoriteFilms(ctx, req.(*GetAllFavoriteFilmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_GetAllFilmsByGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllFilmsByGenreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetAllFilmsByGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/GetAllFilmsByGenre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetAllFilmsByGenre(ctx, req.(*GetAllFilmsByGenreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_GetAllGenres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllGenresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetAllGenres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/GetAllGenres",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetAllGenres(ctx, req.(*GetAllGenresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_AddFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).AddFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/AddFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).AddFilm(ctx, req.(*AddFilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_FindFilmsShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFilmsShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).FindFilmsShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/FindFilmsShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).FindFilmsShort(ctx, req.(*FindFilmsShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_FindFilmsLong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFilmsShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).FindFilmsLong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/FindFilmsLong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).FindFilmsLong(ctx, req.(*FindFilmsShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_FindSerialsShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFilmsShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).FindSerialsShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/FindSerialsShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).FindSerialsShort(ctx, req.(*FindFilmsShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_FindSerialsLong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFilmsShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).FindSerialsLong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/FindSerialsLong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).FindSerialsLong(ctx, req.(*FindFilmsShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_FindActorsShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindActorsShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).FindActorsShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/FindActorsShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).FindActorsShort(ctx, req.(*FindActorsShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_FindActorsLong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindActorsShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).FindActorsLong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/FindActorsLong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).FindActorsLong(ctx, req.(*FindActorsShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_GetTopFilms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopFilmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetTopFilms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/GetTopFilms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetTopFilms(ctx, req.(*GetTopFilmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_GetAllFilmComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllFilmCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).GetAllFilmComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/GetAllFilmComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).GetAllFilmComments(ctx, req.(*AllFilmCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Films_RemoveComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServer).RemoveComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.Films/RemoveComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServer).RemoveComment(ctx, req.(*RemoveCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Films_ServiceDesc is the grpc.ServiceDesc for Films service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Films_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "session.Films",
	HandlerType: (*FilmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllFilmsPreviews",
			Handler:    _Films_GetAllFilmsPreviews_Handler,
		},
		{
			MethodName: "GetFilmDataByUuid",
			Handler:    _Films_GetFilmDataByUuid_Handler,
		},
		{
			MethodName: "GetFilmPreviewByUuid",
			Handler:    _Films_GetFilmPreviewByUuid_Handler,
		},
		{
			MethodName: "RemoveFilmByUuid",
			Handler:    _Films_RemoveFilmByUuid_Handler,
		},
		{
			MethodName: "GetActorDataByUuid",
			Handler:    _Films_GetActorDataByUuid_Handler,
		},
		{
			MethodName: "GetActorsByFilm",
			Handler:    _Films_GetActorsByFilm_Handler,
		},
		{
			MethodName: "PutFavorite",
			Handler:    _Films_PutFavorite_Handler,
		},
		{
			MethodName: "DeleteFavorite",
			Handler:    _Films_DeleteFavorite_Handler,
		},
		{
			MethodName: "GetAllFavoriteFilms",
			Handler:    _Films_GetAllFavoriteFilms_Handler,
		},
		{
			MethodName: "GetAllFilmsByGenre",
			Handler:    _Films_GetAllFilmsByGenre_Handler,
		},
		{
			MethodName: "GetAllGenres",
			Handler:    _Films_GetAllGenres_Handler,
		},
		{
			MethodName: "AddFilm",
			Handler:    _Films_AddFilm_Handler,
		},
		{
			MethodName: "FindFilmsShort",
			Handler:    _Films_FindFilmsShort_Handler,
		},
		{
			MethodName: "FindFilmsLong",
			Handler:    _Films_FindFilmsLong_Handler,
		},
		{
			MethodName: "FindSerialsShort",
			Handler:    _Films_FindSerialsShort_Handler,
		},
		{
			MethodName: "FindSerialsLong",
			Handler:    _Films_FindSerialsLong_Handler,
		},
		{
			MethodName: "FindActorsShort",
			Handler:    _Films_FindActorsShort_Handler,
		},
		{
			MethodName: "FindActorsLong",
			Handler:    _Films_FindActorsLong_Handler,
		},
		{
			MethodName: "GetTopFilms",
			Handler:    _Films_GetTopFilms_Handler,
		},
		{
			MethodName: "GetAllFilmComments",
			Handler:    _Films_GetAllFilmComments_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _Films_AddComment_Handler,
		},
		{
			MethodName: "RemoveComment",
			Handler:    _Films_RemoveComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/films.proto",
}
