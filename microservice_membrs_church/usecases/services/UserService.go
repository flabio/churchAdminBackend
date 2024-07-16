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

type UserService struct {
	IUser interfaces.IUser
}

func NewUserService() IServices.IUserService {
	return &UserService{
		IUser: repositories.GetUserInstance(),
	}
}

func (s *UserService) GetFindUserAll(c *fiber.Ctx) error {
	result, err := s.IUser.GetFindUserAll()
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
func (s *UserService) GetFindUserById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, _ := s.IUser.GetFindUserById(uint(id))
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}

func (s *UserService) CreateUser(c *fiber.Ctx) error {
	var createUser entities.User
	user := new(dto.UserDTO)
	msgError := validateUser(0, user, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(user).To(&createUser)
	result, err := s.IUser.CreateUser(createUser)
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
func (s *UserService) UpdateUser(c *fiber.Ctx) error {
	var updateUser entities.User
	user := new(dto.UserUpdateDTO)
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, _ := s.IUser.GetFindUserById(uint(id))
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	msgError := validateUpdateUser(uint(id), user, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(user).To(&updateUser)
	result, err := s.IUser.UpdateUser(uint(id), updateUser)
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
func (s *UserService) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	user, _ := s.IUser.GetFindUserById(uint(id))
	if user.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.IUser.DeleteUser(uint(id))
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

// region
// update user
// param id type uint
// param userDto type struct
// param userService type struct
// param c type fiber
// return a string

func validateUpdateUser(id uint, userDto *dto.UserUpdateDTO, userService *UserService, c *fiber.Ctx) string {

	if err := c.BodyParser(userDto); err != nil {
		return err.Error()
	}
	v := validate.Struct(userDto)
	if !v.Validate() {
		return v.Errors.Error()
	}
	result, _ := userService.IUser.GetFindUserByEmail(userDto.Email)
	if result.Email != "" {
		return utils.EMAIL_ALREADY_EXIST
	}
	existIdentification, _ := userService.IUser.GetFindUserByIdentification(userDto.Identication)
	if existIdentification.Identication != "" {
		return utils.IDENTIFICATION_ALREADY_EXIST
	}
	return ""
}

// update user
// param id type uint
// param userDto type struct
// param userService type struct
// param c type fiber
// return a string

func validateUser(id uint, userDto *dto.UserDTO, userService *UserService, c *fiber.Ctx) string {

	if err := c.BodyParser(userDto); err != nil {
		return err.Error()
	}
	v := validate.Struct(userDto)
	if !v.Validate() {
		return v.Errors.Error()
	}
	result, _ := userService.IUser.GetFindUserByEmail(userDto.Email)
	if result.Email != "" {
		return utils.EMAIL_ALREADY_EXIST
	}
	existIdentification, _ := userService.IUser.GetFindUserByIdentification(userDto.Identication)
	if existIdentification.Identication != "" {
		return utils.IDENTIFICATION_ALREADY_EXIST
	}
	return ""
}

// endregion
