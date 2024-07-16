package repositories

import (
	"sync"

	"member_church.com/core/interfaces"
	"member_church.com/infrastructure/database"
	"member_church.com/infrastructure/entities"
)

func GetRolInstance() interfaces.IRol {
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

func (db *OpenConnection) GetFindAll() ([]entities.Rol, error) {
	var roles []entities.Rol
	db.mux.Lock()
	err := db.connection.Order("id desc").Find(&roles).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return roles, err
}

func (db *OpenConnection) GetFindById(id int) (entities.Rol, error) {
	db.mux.Lock()
	var rol entities.Rol
	err := db.connection.First(&rol, id).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return rol, err
}

/*
param:rol is a struct
*/
func (db *OpenConnection) Create(rol entities.Rol) (entities.Rol, error) {
	db.mux.Lock()
	err := db.connection.Create(&rol).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return rol, err
}

/*
@Params: rol Rol is a struct, id is an integer
*/
func (db *OpenConnection) Update(id uint, rol entities.Rol) (entities.Rol, error) {
	db.mux.Lock()
	err := db.connection.Where("id=?", id).Updates(&rol).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return rol, err

}

/*
@param: id is an int
*/
func (db *OpenConnection) Delete(id uint) (bool, error) {
	db.mux.Lock()
	var rol entities.Rol
	err := db.connection.Where("id=?", id).Delete(&rol).Error

	defer database.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}
/*
@params: id is an uint number and name is a string
*/
func (db *OpenConnection) GetFindByName(id uint,name string) (bool, error) {
	db.mux.Lock()
	var rol entities.Rol
	
	query:=db.connection.Where("name=?",name)
	if id>0 {
		query = query.Where("id<>?",id).First(&rol)
	}else{
		query= query.First(&rol)
	}
	err:= query.Error
	defer database.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
	
}
