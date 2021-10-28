package model

import (
	"bubble/dao"
)

type Bubble struct {
	ID int `json:"id"` // gorm默认id为自增主键
	Title string `json:"title"`
	Status bool `json:"status"`
}

/**
	Bubble crud
 */

func CreateBubble(bubble *Bubble) error {
	return dao.DB.Debug().Create(bubble).Error
}

func RetrieveAllBubbles() ([]Bubble, error) {
	var bubbleList []Bubble
	err := dao.DB.Debug().Find(&bubbleList).Error
	return bubbleList, err
}

func UpdateBubbleById(bubble *Bubble) error {
	return dao.DB.Debug().Save(bubble).Error
}

func DeleteBubbleById(id string) error {
	return dao.DB.Debug().Where("id", id).Delete(&Bubble{}).Error
}
