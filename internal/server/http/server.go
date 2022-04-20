package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/wuyadong1990/grpc-demo-proto/user"

	//"github.com/wuyadong1990/grpc-demo-user/cinit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var opentracingHeaders = []string{
	"x-b3-traceid",
	"x-b3-spanid",
	"x-b3-parentspanid",
	"x-b3-sampled",
	"x-b3-flags"}

const (
	SN = "srv-user" // 定义services名称
)

func injectHeadersIntoMetadata(ctx context.Context, req *http.Request) metadata.MD {
	pairs := []string{}
	for _, h := range opentracingHeaders {
		if v := req.Header.Get(h); len(v) > 0 {
			pairs = append(pairs, h, v)
		}
	}
	return metadata.Pairs(pairs...)
}

type annotator func(context.Context, *http.Request) metadata.MD

func chainGrpcAnnotators(annotators ...annotator) annotator {
	return func(c context.Context, r *http.Request) metadata.MD {
		mds := []metadata.MD{}
		for _, a := range annotators {
			mds = append(mds, a(c, r))
		}
		return metadata.Join(mds...)
	}
}

func Serve(wg *sync.WaitGroup, grpc_port string, http_port string) {
	defer wg.Done()
	// 初始化,选着需要的组件
	//cinit.InitOption(SN, cinit.MySQL)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	annotators := []annotator{injectHeadersIntoMetadata}

	mux := runtime.NewServeMux(runtime.WithMetadata(chainGrpcAnnotators(annotators...)))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", grpc_port), opts)
	if err != nil {
		return
	}

	log.Printf("[User] Serving HTTP on localhost:%s ...", http_port)

	http.ListenAndServe(fmt.Sprintf(":%s", http_port), mux)
}
