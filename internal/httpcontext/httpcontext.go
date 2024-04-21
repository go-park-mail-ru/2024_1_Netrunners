package httpcontext

import (
	"context"
	"math/rand"
)

type ContextKey string

const ReqIDKey ContextKey = "req_id"
const CtxEmail ContextKey = "email"

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateRequestID() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = symbols[rand.Int63()%int64(len(symbols))]
	}

	return string(b)
}

func GenerateReqIdCTX(reqId string) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ReqIDKey, reqId)

	return ctx
}
