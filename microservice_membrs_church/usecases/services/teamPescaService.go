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

type teamPescaService struct {
	ITeamPesca interfaces.ITeamPesca
}

func NewTeamPescaService() IServices.ITeamPescaService {
	return &teamPescaService{
		ITeamPesca: repositories.GetTeamPescaInstance(),
	}
}

func (s *teamPescaService) GetTeamPescaFindAll(c *fiber.Ctx) error {
	result, err := s.ITeamPesca.GetTeamPescaFindAll()
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
func (s *teamPescaService) GetTeamPescaFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.ITeamPesca.GetTeamPescaFindById(uint(id))
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
func (s *teamPescaService) CreateTeamPesca(c *fiber.Ctx) error {
	var teamPescaCreate entities.TeamPesca
	teamPescaDto := new(dto.TeamPescaDTO)
	msgError := validateTeamPesca(0, teamPescaDto, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(teamPescaDto).To(&teamPescaCreate)
	result, err := s.ITeamPesca.CreateTeamPesca(teamPescaCreate)
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
func (s *teamPescaService) UpdateTeamPesca(c *fiber.Ctx) error {
	var teamPescaUpdate entities.TeamPesca
	teamPescaDto := new(dto.TeamPescaDTO)
	id, _ := strconv.Atoi(c.Params(utils.ID))
	existId, err := s.ITeamPesca.GetTeamPescaFindById(uint(id))
	if existId.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: err.Error(),
		})
	}
	msgError := validateTeamPesca(existId.Id, teamPescaDto, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(teamPescaDto).To(&teamPescaUpdate)
	result, err := s.ITeamPesca.UpdateTeamPesca(existId.Id, teamPescaUpdate)
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

func (s *teamPescaService) DeleteTeamPesca(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	existId, err := s.ITeamPesca.GetTeamPescaFindById(uint(id))
	if existId.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: err.Error(),
		})
	}
	result, err := s.ITeamPesca.DeleteTeamPesca(existId.Id)
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
// validateTeamPesca validate teamPesca
// param id uint
// param teamPescaDto *dto.TeamPescaDTO
// param s *teamPescaService
// param c *fiber.Ctx
// return string
func validateTeamPesca(id uint, teamPescaDto *dto.TeamPescaDTO, s *teamPescaService, c *fiber.Ctx) string {
	if err := c.BodyParser(teamPescaDto); err != nil {
		return err.Error()
	}
	v := validate.Struct(teamPescaDto)
	if !v.Validate() {
		return v.Errors.Error()
	}
	existName, _ := s.ITeamPesca.GetTeamPescaFindByName(id, teamPescaDto.Name)
	if existName {
		return utils.NAME_ALREADY_EXIST
	}
	return ""
}
