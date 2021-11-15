package service

import "errors"

type AuthRequest struct {
	AppKey    string
	AppSecret string
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth == nil {
		return errors.New("auth info does not exists")
	}
	return nil
}
