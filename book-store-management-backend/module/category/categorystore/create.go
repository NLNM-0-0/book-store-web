package categorystore

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/category/categorymodel"
	"context"
)

func (s *sqlStore) CreateCategory(ctx context.Context, data *categorymodel.Category) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
			case "PRIMARY":
				return categorymodel.ErrCategoryIdDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
