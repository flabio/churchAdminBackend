package dto

type ChurchDTO struct {
	Id      uint   `json:"id"`
	Name    string `json:"name" validate:"required|min_len:3"  message:"required:{field} is required|min_len: min 3 caracters" label:"Name"`
	Address string `json:"address" validate:"required|min_len:3" message:"required:{field} is required|min_len: min 3 caracters" label:"Address"`
	Phone   string `json:"phone" validate:"required|min_len:3" message:"required:{field} is required|min" label:"Phone"`
	Email   string `json:"email" validate:"required|min_len:3" message:"required:{field} is required|min" label:"Email"`
	Active  bool   `json:"active" validate:"required" message:"required:{field} is required" label:"Active"`
}
