package controller

import (
	"forum/src/service"

	"github.com/gin-gonic/gin"
)

func GetAllSection(c *gin.Context) {
	data, err := service.GetAllSection()
	if err != nil {
		apiErr(c, err.Error())
		return
	}

	type listItem struct {
		SectionID int    `json:"section_id"`
		Name      string `json:"name"`
	}
	list := make([]listItem, len(data))
	for i := range data {
		list[i].SectionID = data[i].SectionID
		list[i].Name = data[i].Name
	}

	apiOK(c, gin.H{
		"count": len(list),
		"list":  list,
	}, "获取所有板块成功")
}
