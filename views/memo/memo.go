package memo

import (
	"dapan/dbx"
	"dapan/model"
	"dapan/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func AddMemo(c *gin.Context) {
	// 获取当前用户
	user_id, ok := c.Get("user_id")
	id := user_id.(uint)

	var memoInfo model.MemoInfo

	err := c.ShouldBind(&memoInfo)

	if ok {
		memoInfo.Creator = id
	}

	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	if sqlErr := dbx.DB.Create(&memoInfo); sqlErr.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": sqlErr.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
} 

func UpdateMemo(c *gin.Context) {
	type MemoInfo struct {
		Status   string `form:"status" binding:"required"`
		MemoDesc string `form:"memo_desc" binding:"required"`
		MemoName string `form:"memo_name" binding:"required"`
	}

	var	memo_info MemoInfo
	err := c.ShouldBind(&memo_info)
	id := c.Param("id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if updateErr := dbx.DB.Where("id", id).Updates(memo_info); updateErr.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": updateErr.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
 }

func DeleteMemo(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	if (id == "") {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}

	if deleteErr := dbx.DB.Unscoped().Where("id", id).Delete(&model.MemoInfo{}); deleteErr.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": deleteErr.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

func GetMemo(c *gin.Context) {
	type Memo struct {
		ID uint `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		MemoName string `json:"memo_name"`
		MemoDesc string `json:"memo_desc"`
		Status string `json:"status"`
		Username string `json:"username"`
	}
	
	var memos []Memo

	if queryErr := dbx.DB.Table("memo_infos").Scopes(utils.Pagination(c)).Select("memo_infos.id, memo_infos.created_at, memo_infos.memo_name, memo_infos.memo_desc, memo_infos.status, user_infos.username").Joins("left join user_infos on memo_infos.creator = user_infos.id").Scan(&memos); queryErr.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": queryErr.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": memos,
	})
}