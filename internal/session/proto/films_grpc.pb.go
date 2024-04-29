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
	GetAllFilmComments(ctx context.Context, in *AllFilmCommentsRequest, opts ...grpc.CallOption) (*AllFilmCommentsResponse, error)
	RemoveFilmByUuid(ctx context.Context, in *RemoveFilmByUuidRequest, opts ...grpc.CallOption) (*RemoveFilmByUuidResponse, error)
	GetActorDataByUuid(ctx context.Context, in *ActorDataByUuidRequest, opts ...grpc.CallOption) (*ActorDataByUuidResponse, error)
	GetActorsByFilm(ctx context.Context, in *ActorsByFilmRequest, opts ...grpc.CallOption) (*ActorsByFilmResponse, error)
	PutFavorite(ctx context.Context, in *PutFavoriteRequest, opts ...grpc.CallOption) (*PutFavoriteResponse, error)
	DeleteFavorite(ctx context.Context, in *DeleteFavoriteRequest, opts ...grpc.CallOption) (*DeleteFavoriteResponse, error)
	GetAllFavoriteFilms(ctx context.Context, in *GetAllFavoriteFilmsRequest, opts ...grpc.CallOption) (*GetAllFavoriteFilmsResponse, error)
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

func (c *filmsClient) GetAllFilmComments(ctx context.Context, in *AllFilmCommentsRequest, opts ...grpc.CallOption) (*AllFilmCommentsResponse, error) {
	out := new(AllFilmCommentsResponse)
	err := c.cc.Invoke(ctx, "/session.Films/GetAllFilmComments", in, out, opts...)
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

// FilmsServer is the server API for Films service.
// All implementations must embed UnimplementedFilmsServer
// for forward compatibility
type FilmsServer interface {
	GetAllFilmsPreviews(context.Context, *AllFilmsPreviewsRequest) (*AllFilmsPreviewsResponse, error)
	GetFilmDataByUuid(context.Context, *FilmDataByUuidRequest) (*FilmDataByUuidResponse, error)
	GetFilmPreviewByUuid(context.Context, *FilmPreviewByUuidRequest) (*FilmPreviewByUuidResponse, error)
	GetAllFilmComments(context.Context, *AllFilmCommentsRequest) (*AllFilmCommentsResponse, error)
	RemoveFilmByUuid(context.Context, *RemoveFilmByUuidRequest) (*RemoveFilmByUuidResponse, error)
	GetActorDataByUuid(context.Context, *ActorDataByUuidRequest) (*ActorDataByUuidResponse, error)
	GetActorsByFilm(context.Context, *ActorsByFilmRequest) (*ActorsByFilmResponse, error)
	PutFavorite(context.Context, *PutFavoriteRequest) (*PutFavoriteResponse, error)
	DeleteFavorite(context.Context, *DeleteFavoriteRequest) (*DeleteFavoriteResponse, error)
	GetAllFavoriteFilms(context.Context, *GetAllFavoriteFilmsRequest) (*GetAllFavoriteFilmsResponse, error)
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
func (UnimplementedFilmsServer) GetAllFilmComments(context.Context, *AllFilmCommentsRequest) (*AllFilmCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFilmComments not implemented")
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
			MethodName: "GetAllFilmComments",
			Handler:    _Films_GetAllFilmComments_Handler,
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/films.proto",
}
