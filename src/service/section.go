package service

import (
	"fmt"
	"forum/src/model"
)

func getSection(sectionID int) (*model.Section, error) {
	section := model.Section{}
	err := model.DB.Get(&section, "SELECT * FROM section WHERE section_id = ?", sectionID)
	if err != nil {
		return nil, fmt.Errorf("板块不存在")
	}
	return &section, nil
}
