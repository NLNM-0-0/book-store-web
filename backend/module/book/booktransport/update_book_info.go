package booktransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/module/book/bookbiz"
	"book-store-management-backend/module/book/bookmodel"
	"book-store-management-backend/module/book/bookrepo"
	"book-store-management-backend/module/book/bookstore"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// UpdateBookInfo
// @BasePath /v1
// @Security BearerAuth
// @Summary Update Book Info
// @Tags books
// @Accept json
// @Produce json
// @Param book body bookmodel.ReqUpdateBook true "Update Book"
// @Response 200 {object} common.ResSuccess "book id"
// @Router /books [patch]
func UpdateBookInfo(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		id := strings.Trim(c.Param("id"), " ")
		var reqData bookmodel.ReqUpdateBook

		if err := c.ShouldBind(&reqData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if id == "" {
			panic(common.ErrInvalidRequest(fmt.Errorf("id is empty")))
			return
		}

		requester := c.MustGet(common.CurrentUserStr).(middleware.Requester)

		db := appCtx.GetMainDBConnection().Begin()

		store := bookstore.NewSQLStore(db)

		repo := bookrepo.NewUpdateBookRepo(store)

		biz := bookbiz.NewUpdateBookInfoBiz(repo, requester)

		err := biz.UpdateBook(c.Request.Context(), id, &reqData)
		if err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			db.Rollback()
			panic(err)
		}

		c.JSON(http.StatusOK, common.ResSuccess{IsSuccess: true})
	}
}
