package services

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/ulule/deepcopier"
	"member_church.com/core/interfaces"
	"member_church.com/core/repositories"
	"member_church.com/infrastructure/entities"
	"member_church.com/infrastructure/utils"
	"member_church.com/usecases/IServices"
	"member_church.com/usecases/dto"
)

type rolService struct {
	Irol interfaces.IRol
}

func NewRolService() IServices.IRolService {
	return &rolService{
		Irol: repositories.GetRolInstance(),
	}
}

func (rolService *rolService) GetFindAll(c *fiber.Ctx) error {
	result, err := rolService.Irol.GetFindAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}
func (rolService *rolService) GetFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := rolService.Irol.GetFindById(id)
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: err.Error(),
		})
	}
	
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}
func (rolService *rolService) Create(c *fiber.Ctx)error {
	var rolCreate entities.Rol
	rol := new(dto.RolDTO)

	msgError := validateRol(0, rol, rolService, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
		
	}

	deepcopier.Copy(rol).To(&rolCreate)
	result, err := rolService.Irol.Create(rolCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}
func (rolService *rolService) Update(c *fiber.Ctx)error {

	var rolEntity entities.Rol
	rolDto := new(dto.RolDTO)

	id, _ := strconv.Atoi(c.Params(utils.ID))
	rol, err := rolService.Irol.GetFindById(id)
	if rol.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: err.Error(),
		})
	}

	msgError := validateRol(rol.Id, rolDto, rolService, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}

	deepcopier.Copy(rolDto).To(&rolEntity)
	result, err := rolService.Irol.Update(rol.Id, rolEntity)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS: http.StatusCreated,
		utils.DATA:   result,
	})
}

func (rolService *rolService) Delete(c *fiber.Ctx)error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	rol, err := rolService.Irol.GetFindById(id)
	if rol.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: err.Error(),
		})
	}
	result, err := rolService.Irol.Delete(rol.Id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS: http.StatusCreated,
		utils.DATA:   result,
	})
}
func validateRol(id uint, rolDto *dto.RolDTO, rolService *rolService, c *fiber.Ctx) string {

	if err := c.BodyParser(rolDto); err != nil {
		return err.Error()
	}
	v := validate.Struct(rolDto)
	if !v.Validate() {
		return v.Errors.Error()
	}
	existName, _ := rolService.Irol.GetFindByName(id, rolDto.Name)
	if existName {
		return utils.NAME_ALREADY_EXIST
	}
	return ""
}

// func MapRolEntityToDTO(rol entities.Rol) dto.RolDTO {
// 	return dto.RolDTO{
// 		Id:     rol.Id,
// 		Name:   rol.Name,
// 		Active: rol.Active,
// 	}
// }
