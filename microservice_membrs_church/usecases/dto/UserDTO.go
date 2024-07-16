package dto

type UserDTO struct {
	Id                 uint   `json:"id"`
	FirstName          string `json:"first_name" validate:"required|min_len:3"  message:"required:{field} is required|min_len: min 3 caracters" label:"First Name"`
	LastName           string `json:"last_name" validate:"required|min_len:3" message:"required:{field} is required|min_len: min 3 caracters" label:"Last Name"`
	Identication       string `json:"identification" validate:"required|min_len:5" message:"required:{field} is required|min_len: min 5 caracters" label:"Identification"`
	TypeIdentification string `json:"type_identification" validate:"required|min_len:2|max_len:4" message:"required:{field} is required|min_le|max_len: min 2 and max 4 caracters" label:"Type Identification"`
	Birthdate          string `json:"birthdate" validate:"required" message:"required:{field} is required" label:"Birthdate"`
	PlaceOfBirth       string `json:"placeofbirth"`
	Address            string `json:"address" validate:"required" message:"required:{field} is required" label:"Address"`
	Phone              string `json:"phone" validate:"required" message:"required:{field} is required" label:"Phone"`
	Cellphone          string `json:"cellphone" validate:"required" message:"required:{field} is required" label:"Cellphone"`
	Email              string `json:"email" validate:"required" message:"required:{field} is required" label:"Email"`
	Neighborhood       string `json:"neighborhood"`
	Locality           string `json:"locality"`
	Socioeconomic      uint   `json:"socioeconomic"`
	Rh                 string `json:"rh" validate:"required|max_len:4" message:"required:{field} is required|max_len: max 4 caracters" label:"Rh"`
	Profession         string `json:"profession"`
	Company            string `json:"company"`
	Position           string `json:"position"`
	CivilStatus        string `json:"civil_status"`
	LeadersName        string `json:"leaders_name"`
	ConversionDate     string `json:"conversion_date"`
	WhoInvitedHim      string `json:"who_invited_him"`
	ChurchAttended     bool   `json:"century_attended"`
	ChurchName         string `json:"church_name"`
	MeetingDate        string `json:"meetings_date"`
	BaptismDate        string `json:"baptism_date"`
	YourSpouseName     string `json:"your_spouse_name"`
	SpouseMemberChurch bool   `json:"spouse_member_church"`
	DateMarriage       string `json:"date_marriage"`
	ParentsName        string `json:"parents_name"`
	IsMemberParent     bool   `json:"is_member_parent"`
	ChildrenaName      string `json:"childrena_name"`
	Allergy            bool   `json:"allergy"`
	Authorize          bool   `json:"authorize"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	ChurchId           uint64 `json:"church_id"`
	RolId              uint   `json:"rol_id"`
	TeamPescaId        uint   `json:"team_pesca_id"`
}


type UserUpdateDTO struct {
	Id                 uint   `json:"id"`
	FirstName          string `json:"first_name" validate:"required|min_len:3"  message:"required:{field} is required|min_len: min 3 caracters" label:"First Name"`
	LastName           string `json:"last_name" validate:"required|min_len:3" message:"required:{field} is required|min_len: min 3 caracters" label:"Last Name"`
	Identication       string `json:"identification" validate:"required|min_len:5" message:"required:{field} is required|min_len: min 5 caracters" label:"Identification"`
	TypeIdentification string `json:"type_identification" validate:"required|min_len:2|max_len:4" message:"required:{field} is required|min_le|max_len: min 2 and max 4 caracters" label:"Type Identification"`
	Birthdate          string `json:"birthdate" validate:"required" message:"required:{field} is required" label:"Birthdate"`
	PlaceOfBirth       string `json:"placeofbirth"`
	Address            string `json:"address" validate:"required" message:"required:{field} is required" label:"Address"`
	Phone              string `json:"phone" validate:"required" message:"required:{field} is required" label:"Phone"`
	Cellphone          string `json:"cellphone" validate:"required" message:"required:{field} is required" label:"Cellphone"`
	Email              string `json:"email" validate:"required" message:"required:{field} is required" label:"Email"`
	Neighborhood       string `json:"neighborhood"`
	Locality           string `json:"locality"`
	Socioeconomic      uint   `json:"socioeconomic"`
	Rh                 string `json:"rh" validate:"required|max_len:4" message:"required:{field} is required|max_len: max 4 caracters" label:"Rh"`
	Profession         string `json:"profession"`
	Company            string `json:"company"`
	Position           string `json:"position"`
	CivilStatus        string `json:"civil_status"`
	LeadersName        string `json:"leaders_name"`
	ConversionDate     string `json:"conversion_date"`
	WhoInvitedHim      string `json:"who_invited_him"`
	ChurchAttended     bool   `json:"century_attended"`
	ChurchName         string `json:"church_name"`
	MeetingDate        string `json:"meetings_date"`
	BaptismDate        string `json:"baptism_date"`
	YourSpouseName     string `json:"your_spouse_name"`
	SpouseMemberChurch bool   `json:"spouse_member_church"`
	DateMarriage       string `json:"date_marriage"`
	ParentsName        string `json:"parents_name"`
	IsMemberParent     bool   `json:"is_member_parent"`
	ChildrenaName      string `json:"childrena_name"`
	Allergy            bool   `json:"allergy"`
	ChurchId           uint64 `json:"church_id"`
	RolId              uint   `json:"rol_id"`
	TeamPescaId        uint   `json:"team_pesca_id"`
}