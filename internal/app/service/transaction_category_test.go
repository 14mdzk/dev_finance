package service

import (
	"testing"

	"github.com/14mdzk/dev_finance/internal/app/service/mocks"
	"github.com/golang/mock/gomock"
)

func TestTransactionCategoryService_BrowseAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockITransactionCategoryRepository(ctrl)
	repo.EXPECT().Browse()
}
