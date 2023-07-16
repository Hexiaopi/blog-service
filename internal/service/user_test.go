package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/pkg/auth"
	"github.com/hexiaopi/blog-service/internal/store"
)

func TestUserService_CheckAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	password, err := auth.Encrypt("123456")
	if err != nil {
		t.Fatal(err)
	}
	mockStore := store.NewMockUserStore(ctrl)
	mockStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(
		&model.User{ID: 1, Name: "admin", PassWord: password},
		nil,
	)
	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(mockStore).AnyTimes()
	userSrv := NewUserService(mockFactory)
	if err := userSrv.CheckAuth(
		context.Background(),
		&AuthRequest{
			UserId:   1,
			UserName: "admin",
			PassWord: "123456",
		}); err != nil {
		t.Fatal(err)
	}
}

func TestUserService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := store.NewMockUserStore(ctrl)
	mockStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(
		&model.User{ID: 1, Name: "admin", PassWord: "xxx"},
		nil,
	)
	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(mockStore).AnyTimes()
	userSrv := NewUserService(mockFactory)
	user, err := userSrv.Get(context.Background(), "admin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestUserService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockUserStore(ctrl)
	mockStore.EXPECT().Count(gomock.Any(), gomock.Any()).Return(
		int64(1), nil)
	mockStore.EXPECT().List(gomock.Any(), gomock.Any()).Return(
		[]model.User{
			{ID: 1, Name: "admin"},
		}, nil)

	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(mockStore).AnyTimes()

	userSrv := NewUserService(mockFactory)
	users, total, err := userSrv.List(context.Background(), &ListUserRequest{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(total)
	for _, user := range users {
		t.Log(user)
	}
}

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := store.NewMockUserStore(ctrl)
	mockStore.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(mockStore).AnyTimes()
	userSrv := NewUserService(mockFactory)
	if err := userSrv.Create(
		context.Background(),
		&CreateUserRequest{
			model.User{
				ID:       1,
				Name:     "admin",
				PassWord: "123456",
			}}); err != nil {
		t.Fatal(err)
	}
}
