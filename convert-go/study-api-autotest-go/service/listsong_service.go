package service

import (
	"study-api-autotest-go/common"
	"study-api-autotest-go/mapper"
	"study-api-autotest-go/models"
)

type ListSongService struct {
	listSongMapper *mapper.ListSongMapper
}

func NewListSongService() *ListSongService {
	return &ListSongService{
		listSongMapper: mapper.NewListSongMapper(),
	}
}

func (s *ListSongService) AddListSong(req *models.ListSongRequest) *common.Response {
	listSong := &models.ListSong{
		SongId:     req.SongId,
		SongListId: req.SongListId,
	}
	err := s.listSongMapper.Add(listSong)
	if err != nil {
		return common.Error("添加歌曲到歌单失败")
	}
	return common.Success("添加成功")
}

func (s *ListSongService) DeleteListSong(id uint) *common.Response {
	err := s.listSongMapper.Delete(id)
	if err != nil {
		return common.Error("删除失败")
	}
	return common.Success("删除成功")
}

func (s *ListSongService) ListSongOfSongListId(songListId uint) *common.Response {
	listSongs, err := s.listSongMapper.FindBySongListId(songListId)
	if err != nil {
		return common.Error("获取失败")
	}
	return common.SuccessWithData("获取成功", listSongs)
}
