package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/store"
)

func TestUserService_CheckAuth(t *testing.T) {
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
