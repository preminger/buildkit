package secrets

import (
	"context"

	"github.com/pkg/errors"
	"github.com/preminger/buildkit/session"
	"github.com/preminger/buildkit/util/grpcerrors"
	"google.golang.org/grpc/codes"
)

type SecretStore interface {
	GetSecret(context.Context, string) ([]byte, error)
}

var ErrNotFound = errors.Errorf("not found")

func GetSecret(ctx context.Context, c session.Caller, id string) ([]byte, error) {
	client := NewSecretsClient(c.Conn())
	resp, err := client.GetSecret(ctx, &GetSecretRequest{
		ID: id,
	})
	if err != nil {
		if code := grpcerrors.Code(err); code == codes.Unimplemented || code == codes.NotFound {
			return nil, errors.Wrapf(ErrNotFound, "secret %s", id)
		}
		return nil, err
	}
	return resp.Data, nil
}
