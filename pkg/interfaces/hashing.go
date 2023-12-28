package interfaces

import "context"

type HashInteractor interface {
	HashFunction(ctx context.Context, inputToHash string) (string, bool)
}
