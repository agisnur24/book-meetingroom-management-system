package authenticate

import "context"

type AuthenthicationService interface {
	GetTokenUser(ctx context.Context, userName string, password string) (string, error)
}
