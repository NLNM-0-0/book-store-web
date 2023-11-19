package inventorychecknotemodel

import (
	"book-store-management-backend/common"
	"book-store-management-backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
)

type InventoryCheckNoteCreate struct {
	Id                *string                                                        `json:"id" gorm:"column:id;"`
	AmountDifferent   float32                                                        `json:"-" gorm:"column:amountDifferent;"`
	AmountAfterAdjust float32                                                        `json:"-" gorm:"column:amountAfterAdjust;"`
	CreateBy          string                                                         `json:"-" gorm:"column:createBy;"`
	Details           []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate `json:"details" gorm:"-"`
}

func (*InventoryCheckNoteCreate) TableName() string {
	return common.TableInventoryCheckNote
}

func (data *InventoryCheckNoteCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrInventoryCheckNoteIdInvalid
	}

	mapExits := make(map[string]int)
	for _, detail := range data.Details {
		if err := detail.Validate(); err != nil {
			return err
		}
		mapExits[detail.BookId]++
		if mapExits[detail.BookId] >= 2 {
			return ErrInventoryCheckNoteExistDuplicateBook
		}
	}
	return nil
}