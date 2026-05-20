// Package cauth provides a gRPC per-RPC credentials implementation
// that sends an API key as a Bearer token in the authorization metadata.
package cauth

import (
	"context"
	"fmt"

	"google.golang.org/grpc/credentials"
)

// BearerToken implements credentials.PerRPCCredentials to send a
// static API key as a Bearer Authorization header on every RPC.
type BearerToken struct {
	Token string
}

// NewBearerToken creates a new BearerToken from the given API key.
// If the key is empty, GetRequestMetadata returns empty metadata.
func NewBearerToken(key string) credentials.PerRPCCredentials {
	return &BearerToken{Token: key}
}

// GetRequestMetadata returns the authorization metadata.
func (b *BearerToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	if b.Token == "" {
		return nil, nil
	}
	return map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", b.Token),
	}, nil
}

// RequireTransportSecurity indicates whether the credentials require TLS.
// We return false because token auth can be used with or without TLS;
// gRPC will still enforce TLS if transport credentials are also configured.
func (b *BearerToken) RequireTransportSecurity() bool {
	return false
}
