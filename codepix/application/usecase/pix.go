package usecase

import (
	"github.com/gabruw/codepix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInteface
}

func (puc *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := puc.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(account, key, kind)
	if err != nil {
		return nil, err
	}

	puc.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID != "" {
		return nil, err
	}

	return pixKey, nil
}

func (puc *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := puc.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
