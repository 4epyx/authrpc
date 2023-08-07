package util

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

func ParseMetadataFromCtx(ctx context.Context) (metadata.MD, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("error while getting metadata from context")
	}
	return md, nil
}

func GetHeaderValue(ctx context.Context, headerName string) (string, error) {
	md, err := ParseMetadataFromCtx(ctx)
	if err != nil {
		return "", err
	}
	header := md[headerName]
	if len(header) == 0 {
		return "", fmt.Errorf("header %s not found", headerName)
	}

	return header[0], nil
}

func GetAuthorizationToken(ctx context.Context) (string, error) {
	auth, err := GetHeaderValue(ctx, "authorization")
	if err != nil {
		return "", err
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	fmt.Println(token)

	return token, nil
}
