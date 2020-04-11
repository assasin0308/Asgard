package services

import (
	"fmt"

	"github.com/dalonghahaha/avenger/components/logger"
	"github.com/jinzhu/gorm"

	"Asgard/models"
)

type TimingService struct {
}

func NewTimingService() *TimingService {
	return &TimingService{}
}

func (s *TimingService) GetTimingCount(where map[string]interface{}) (count int) {
	err := models.Count(&models.Timing{}, where, &count)
	if err != nil {
		logger.Error("GetTimingCount Error:", err)
		return 0
	}
	return
}

func (s *TimingService) GetTimingPageList(where map[string]interface{}, page int, pageSize int) (list []models.Timing, count int) {
	condition := "1=1"
	for key, val := range where {
		if key == "status" {
			if val.(int) == -99 {
				condition += " and status != -1"
			} else {
				condition += fmt.Sprintf(" and %s=%v", key, val)
			}
		} else if key == "name" {
			condition += fmt.Sprintf(" and %s like '%%%v%%' ", key, val)
		} else {
			condition += fmt.Sprintf(" and %s=%v", key, val)
		}
	}
	err := models.PageListbyWhereString(&models.Timing{}, condition, page, pageSize, "created_at desc", &list, &count)
	if err != nil {
		logger.Error("GetTimingPageList Error:", err)
		return nil, 0
	}
	return
}

func (s *TimingService) GetTimingByID(id int64) *models.Timing {
	var timing models.Timing
	err := models.Get(id, &timing)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.Error("GetTimingID Error:", err)
		}
		return nil
	}
	return &timing
}

func (s *TimingService) GetTimingByAgentID(id int64) (list []models.Timing) {
	err := models.Where(&list, "agent_id = ? and status != ? and status != ?", id, models.STATUS_PAUSE, models.STATUS_FINISHED)
	if err != nil {
		logger.Error("GetTimingByAgentID Error:", err)
		return nil
	}
	return
}

func (s *TimingService) CreateTiming(timing *models.Timing) bool {
	err := models.Create(timing)
	if err != nil {
		logger.Error("CreateTiming Error:", err)
		return false
	}
	return true
}

func (s *TimingService) UpdateTiming(timing *models.Timing) bool {
	err := models.Update(timing)
	if err != nil {
		logger.Error("UpdateTiming Error:", err)
		return false
	}
	return true
}

func (s *TimingService) DeleteJobByID(id int64) bool {
	timing := new(models.Timing)
	timing.ID = id
	err := models.Delete(timing)
	if err != nil {
		logger.Error("DeleteAppByID Error:", err)
		return false
	}
	return true
}
