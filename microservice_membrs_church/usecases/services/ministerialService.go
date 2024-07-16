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

type ministerialService struct {
	Iministerial interfaces.IMinisterial
}

func NewMinisterialService() IServices.IMinisterialService {
	return &ministerialService{
		Iministerial: repositories.GetMinisterialInstance(),
	}
}
func (s *ministerialService) GetMinisterialFindAll(c *fiber.Ctx) error {
	result, err := s.Iministerial.GetMinisterialFindAll()
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
func (s *ministerialService) GetMinisterialFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.Iministerial.GetMinisterialFindById(uint(id))
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

func (s *ministerialService) CreateMinisterial(c *fiber.Ctx) error {
	var ministerialCreate entities.Ministerial
	ministerial := new(dto.MinisterialDTO)
	msgError := validateMinisterial(0, ministerial, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(ministerial).To(&ministerialCreate)
	result, err := s.Iministerial.CreateMinisterial(ministerialCreate)
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

// param c *fiber.Ctx
func (s *ministerialService) UpdateMinisterial(c *fiber.Ctx) error {

	var ministerialUpdate entities.Ministerial
	ministerial := new(dto.MinisterialDTO)
	id, _ := strconv.Atoi(c.Params(utils.ID))
	existId, err := s.Iministerial.GetMinisterialFindById(uint(id))
	if existId.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: err.Error(),
		})
	}

	msgError := validateMinisterial(uint(id), ministerial, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(ministerial).To(&ministerialUpdate)
	result, err := s.Iministerial.UpdateMinisterial(existId.Id, ministerialUpdate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.UPDATED,
		utils.DATA:    result,
	})
}

// param c *fiber.Ctx

func (s *ministerialService) DeleteMinisterial(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	existId, err := s.Iministerial.GetMinisterialFindById(uint(id))
	if existId.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: err.Error(),
		})
	}
	result, err := s.Iministerial.DeleteMinisterial(existId.Id)
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

// method private
// validateMinisterial validate ministerial
// param id uint
// param ministerialDto *dto.MinisterialDTO
// param s *ministerialService
// param c *fiber.Ctx
// return string
func validateMinisterial(id uint, ministerialDto *dto.MinisterialDTO, s *ministerialService, c *fiber.Ctx) string {

	if err := c.BodyParser(ministerialDto); err != nil {
		return err.Error()
	}
	v := validate.Struct(ministerialDto)
	if !v.Validate() {
		return v.Errors.Error()
	}
	existName, _ := s.Iministerial.GetMinisterialFindByName(id, ministerialDto.Name)

	if existName {
		return utils.NAME_ALREADY_EXIST
	}
	return ""
}
