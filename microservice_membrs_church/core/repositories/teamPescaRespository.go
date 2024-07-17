package repositories

import (
	"sync"

	"member_church.com/core/interfaces"
	"member_church.com/infrastructure/database"
	"member_church.com/infrastructure/entities"
)

func GetTeamPescaInstance() interfaces.ITeamPesca {
	var (
		_OPEN *OpenConnection
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.DatabaseConnection(),
		}
	})
	return _OPEN
}

func (db *OpenConnection) GetTeamPescaFindAll() ([]entities.TeamPesca, error) {
	var teamPesca []entities.TeamPesca
	db.mux.Lock()
	err := db.connection.Order("id desc").Find(&teamPesca).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return teamPesca, err
}
func (db *OpenConnection) GetTeamPescaFindById(id uint) (entities.TeamPesca, error) {
	var teamPesca entities.TeamPesca
	db.mux.Lock()
	err := db.connection.Where("id=?", id).First(&teamPesca).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return teamPesca, err
}

func (db *OpenConnection) CreateTeamPesca(teamPesca entities.TeamPesca) (entities.TeamPesca, error) {
	db.mux.Lock()
	err := db.connection.Create(&teamPesca).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return teamPesca, err
}

func (db *OpenConnection) UpdateTeamPesca(id uint, teamPesca entities.TeamPesca) (entities.TeamPesca, error) {
	db.mux.Lock()
	err := db.connection.Where("id=?", id).Updates(&teamPesca).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return teamPesca, err
}

func (db *OpenConnection) DeleteTeamPesca(id uint) (bool, error) {
	db.mux.Lock()
	var teamPesca entities.TeamPesca
	err := db.connection.Where("id=?", id).Delete(&teamPesca).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return true, err
}
func (db *OpenConnection) GetTeamPescaFindByName(id uint, name string) (bool, error) {
	db.mux.Lock()
	var teamPesca entities.TeamPesca
	query := db.connection.Where("name=?", name)

	if id>0 {
		query = query.Where("id<>?",id).First(&teamPesca)
	}else{
		query= query.First(&teamPesca)
	}
	err := query.Error
	defer db.mux.Unlock()
	defer database.Closedb()
	if err == nil {
		return true, err
	}
	return false, err

}
