package service

import (
	"study-api-autotest-go/common"
	"study-api-autotest-go/mapper"
	"study-api-autotest-go/models"
)

type SingerService struct {
	singerMapper *mapper.SingerMapper
}

func NewSingerService() *SingerService {
	return &SingerService{
		singerMapper: mapper.NewSingerMapper(),
	}
}

func (s *SingerService) AddSinger(req *models.SingerRequest) *common.Response {
	singer := &models.Singer{
		Name:        req.Name,
		Sex:         req.Sex,
		Pic:         req.Pic,
		Birth:       req.Birth,
		Location:    req.Location,
		Introduction: req.Introduction,
	}
	err := s.singerMapper.Add(singer)
	if err != nil {
		return common.Error("添加歌手失败")
	}
	return common.SuccessWithData("添加成功", singer)
}

func (s *SingerService) UpdateSinger(req *models.SingerRequest) *common.Response {
	singer, err := s.singerMapper.FindById(req.ID)
	if err != nil {
		return common.Error("歌手不存在")
	}
	singer.Name = req.Name
	singer.Sex = req.Sex
	singer.Pic = req.Pic
	singer.Birth = req.Birth
	singer.Location = req.Location
	singer.Introduction = req.Introduction
	err = s.singerMapper.Update(singer)
	if err != nil {
		return common.Error("更新失败")
	}
	return common.Success("更新成功")
}

func (s *SingerService) DeleteSinger(id uint) *common.Response {
	err := s.singerMapper.Delete(id)
	if err != nil {
		return common.Error("删除失败")
	}
	return common.Success("删除成功")
}

func (s *SingerService) SingerOfId(id uint) *common.Response {
	singer, err := s.singerMapper.FindById(id)
	if err != nil {
		return common.Error("歌手不存在")
	}
	return common.SuccessWithData("获取成功", singer)
}

func (s *SingerService) SingerOfName(name string) *common.Response {
	singers, err := s.singerMapper.FindByName(name)
	if err != nil {
		return common.Error("获取失败")
	}
	return common.SuccessWithData("获取成功", singers)
}

func (s *SingerService) AllSinger() *common.Response {
	singers, err := s.singerMapper.FindAll()
	if err != nil {
		return common.Error("获取失败")
	}
	return common.SuccessWithData("获取成功", singers)
}
