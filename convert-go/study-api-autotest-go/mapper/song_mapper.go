package mapper

import (
	"study-api-autotest-go/models"
)

type SongMapper struct{}

func NewSongMapper() *SongMapper {
	return &SongMapper{}
}

func (*SongMapper) Add(song *models.Song) error {
	return DB.Create(song).Error
}

func (*SongMapper) FindById(id uint) (*models.Song, error) {
	var song models.Song
	err := DB.First(&song, id).Error
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (*SongMapper) FindAll() ([]models.Song, error) {
	var songs []models.Song
	err := DB.Find(&songs).Error
	return songs, err
}

func (*SongMapper) FindBySingerId(singerId uint) ([]models.Song, error) {
	var songs []models.Song
	err := DB.Where("singer_id = ?", singerId).Find(&songs).Error
	return songs, err
}

func (*SongMapper) FindByName(name string) ([]models.Song, error) {
	var songs []models.Song
	err := DB.Where("name LIKE ?", "%"+name+"%").Find(&songs).Error
	return songs, err
}

func (*SongMapper) Update(song *models.Song) error {
	return DB.Save(song).Error
}

func (*SongMapper) Delete(id uint) error {
	return DB.Delete(&models.Song{}, id).Error
}
