package service

import (
	"fmt"
	"testing"

	"github.com/14mdzk/dev_finance/internal/app/model"
	"github.com/14mdzk/dev_finance/internal/app/schema"
	"github.com/14mdzk/dev_finance/internal/app/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTransactionCategoryService_BrowseAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockITransactionCategoryRepository(ctrl)
	repo.EXPECT().Browse(fmt.Sprintf(`
		LIMIT %d 
		OFFSET %d
	`, 2, 0)).Return([]model.TransactionCategory{
		{
			ID:          1,
			Name:        "Food & Beverage",
			Description: "Food & Beverage description",
		},
		{
			ID:          2,
			Name:        "Shopping",
			Description: "Shopping description",
		},
	}, nil)

	categoryService := NewTransactionCategoryService(repo)
	categories, err := categoryService.BrowseAll(schema.PaginationReq{Page: 1, PageSize: 2})
	total := len(categories)
	assert.Equal(t, total, 2)
	assert.NoError(t, err)
}
