package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"member_church.com/core/interfaces"
	"member_church.com/core/repositories"
	"member_church.com/infrastructure/entities"
	"member_church.com/infrastructure/utils"
	"member_church.com/usecases/IServices"
	"member_church.com/usecases/dto"
)

type churchService struct {
	IChurch interfaces.IChurch
}

func NewChurchService() IServices.IChurchService {
	return &churchService{
		IChurch: repositories.GetChurchInstance(),
	}
}

func (s *churchService) GetChurchFindAll(c *fiber.Ctx) error {
	result, err := s.IChurch.GetChurchFindAll()
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
func (s *churchService) GetChurchFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	fmt.Println(c)
	result, err := s.IChurch.GetChurchFindById(uint(id))
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
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
func (s *churchService) CreateChurch(c *fiber.Ctx) error {
	var churchCreate entities.Church
	churchDto := new(dto.ChurchDTO)
	msgError := validateChurch(0, churchDto, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	err := c.BodyParser(&churchCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	result, err := s.IChurch.CreateChurch(churchCreate)
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

func (s *churchService) UpdateChurch(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	var churchUpdate entities.Church
	churchDto := new(dto.ChurchDTO)
	msgError := validateChurch(uint(id), churchDto, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	err := c.BodyParser(&churchUpdate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	result, err := s.IChurch.UpdateChurch(uint(id), churchUpdate)
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

func (s *churchService) DeleteChurch(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.IChurch.DeleteChurch(uint(id))
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

func validateChurch(id uint, churchDto *dto.ChurchDTO, s *churchService, c *fiber.Ctx) string {

	if err := c.BodyParser(churchDto); err != nil {
		return err.Error()
	}
	// v := validate.Struct(churchDto)
	// if !v.Validate() {
	// 	return v.Errors.Error()
	// }
	existName, _ := s.IChurch.GetChurchFindByName(id, churchDto.Name)
	if existName {
		return utils.NAME_ALREADY_EXIST
	}
	existEmail, _ := s.IChurch.GetChurchFindByEmail(id, churchDto.Email)
	if existEmail {
		return utils.EMAIL_ALREADY_EXIST
	}
	return ""
}
