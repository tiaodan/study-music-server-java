package service

import (
	"study-api-autotest-go/common"
	"study-api-autotest-go/mapper"
	"study-api-autotest-go/models"
)

type SongService struct {
	songMapper *mapper.SongMapper
}

func NewSongService() *SongService {
	return &SongService{
		songMapper: mapper.NewSongMapper(),
	}
}

func (s *SongService) AddSong(req *models.SongRequest) *common.Response {
	song := &models.Song{
		SingerId:    req.SingerId,
		Name:        req.Name,
		Introduction: req.Introduction,
		Pic:         req.Pic,
		Lyric:       req.Lyric,
		Url:         req.Url,
	}
	err := s.songMapper.Add(song)
	if err != nil {
		return common.Error("添加歌曲失败")
	}
	return common.SuccessWithData("添加成功", song)
}

func (s *SongService) UpdateSong(req *models.SongRequest) *common.Response {
	song, err := s.songMapper.FindById(req.ID)
	if err != nil {
		return common.Error("歌曲不存在")
	}
	song.Name = req.Name
	song.Introduction = req.Introduction
	song.Pic = req.Pic
	song.Lyric = req.Lyric
	song.Url = req.Url
	err = s.songMapper.Update(song)
	if err != nil {
		return common.Error("更新失败")
	}
	return common.Success("更新成功")
}

func (s *SongService) DeleteSong(id uint) *common.Response {
	err := s.songMapper.Delete(id)
	if err != nil {
		return common.Error("删除失败")
	}
	return common.Success("删除成功")
}

func (s *SongService) SongOfId(id uint) *common.Response {
	song, err := s.songMapper.FindById(id)
	if err != nil {
		return common.Error("歌曲不存在")
	}
	return common.SuccessWithData("获取成功", song)
}

func (s *SongService) SongOfSingerId(singerId uint) *common.Response {
	songs, err := s.songMapper.FindBySingerId(singerId)
	if err != nil {
		return common.Error("获取失败")
	}
	return common.SuccessWithData("获取成功", songs)
}

func (s *SongService) SongOfName(name string) *common.Response {
	songs, err := s.songMapper.FindByName(name)
	if err != nil {
		return common.Error("获取失败")
	}
	return common.SuccessWithData("获取成功", songs)
}

func (s *SongService) AllSong() *common.Response {
	songs, err := s.songMapper.FindAll()
	if err != nil {
		return common.Error("获取失败")
	}
	return common.SuccessWithData("获取成功", songs)
}
