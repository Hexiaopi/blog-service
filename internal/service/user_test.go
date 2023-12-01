package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/pkg/auth"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

func TestUserService_CheckAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	password, err := auth.Encrypt("123456")
	if err != nil {
		t.Fatal(err)
	}
	userStore := store.NewMockUserStore(ctrl)
	userStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(
		&model.User{ID: 1, Name: "admin", PassWord: password},
		nil,
	)
	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(userStore).AnyTimes()
	userSrv := NewUserService(mockFactory, log.Std)
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
	userStore := store.NewMockUserStore(ctrl)
	userStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(
		&model.User{ID: 1, Name: "admin", PassWord: "xxx"},
		nil,
	)
	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(userStore).AnyTimes()
	userSrv := NewUserService(mockFactory, log.Std)
	user, err := userSrv.Get(context.Background(), "admin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestUserService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStore := store.NewMockUserStore(ctrl)
	userStore.EXPECT().Count(gomock.Any(), gomock.Any()).Return(
		int64(1), nil)
	userStore.EXPECT().List(gomock.Any(), gomock.Any()).Return(
		[]model.User{
			{ID: 1, Name: "admin"},
		}, nil)

	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(userStore).AnyTimes()

	userSrv := NewUserService(mockFactory, log.Std)
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
	userStore := store.NewMockUserStore(ctrl)
	userStore.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(userStore).AnyTimes()
	userSrv := NewUserService(mockFactory, log.Std)
	if err := userSrv.Create(
		context.Background(),
		&CreateUserRequest{
			model.User{
				ID:       1,
				Name:     "admin",
				PassWord: "123456",
				Roles: []model.Role{
					{ID: 1, Name: "admin"},
				},
			}}); err != nil {
		t.Fatal(err)
	}
}

func TestUserService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStore := store.NewMockUserStore(ctrl)
	userStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&model.User{ID: 1, Name: "admin", Roles: []model.Role{{ID: 1, Name: "admin"}}}, nil)
	userStore.EXPECT().Update(gomock.Any(), gomock.Any()).AnyTimes().Return(nil) //todo why anytime

	userRoleStore := store.NewMockUserRoleStore(ctrl)
	userRoleStore.EXPECT().Create(gomock.Any(), gomock.Any()).AnyTimes().Return(nil) //todo why anytime
	userRoleStore.EXPECT().Delete(gomock.Any(), gomock.Any()).AnyTimes().Return(nil) //todo why anytime

	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(userStore).AnyTimes()
	mockFactory.EXPECT().UserRole().Return(userRoleStore).AnyTimes()
	mockFactory.EXPECT().Tx(gomock.Any(), gomock.Any()).Return(nil)
	userSrv := NewUserService(mockFactory, log.Std)
	if err := userSrv.Update(
		context.Background(),
		&UpdateUserRequest{
			model.User{
				ID:   1,
				Name: "admin",
				Roles: []model.Role{
					{ID: 2, Name: "test"},
				},
			}}); err != nil {
		t.Fatal(err)
	}
}
