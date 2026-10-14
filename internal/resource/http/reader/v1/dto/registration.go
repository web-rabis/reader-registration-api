package dto

import (
	"errors"
	"strings"
	"time"

	rcmodel "github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-registration-api/internal/utils"
)

type RegistrationReaderUserRequest struct {
	Iin            string `json:"iin,omitempty"`
	Lastname       string `json:"lastname,omitempty"`
	Middlename     string `json:"middlename,omitempty"`
	Firstname      string `json:"firstname,omitempty"`
	BirthDate      string `json:"birthDate,omitempty"`
	Address        string `json:"address,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Email          string `json:"email,omitempty"`
	Sex            *int64 `json:"sex,omitempty"`
	Nationality    *int64 `json:"nationality,omitempty"`
	Education      *int64 `json:"education,omitempty"`
	SocialStatus   *int64 `json:"socialStatus,omitempty"`
	Speciality     *int64 `json:"speciality,omitempty"`
	Institution    *int64 `json:"institution,omitempty"`
	Faculty        *int64 `json:"faculty,omitempty"`
	Course         string `json:"course,omitempty"`
	WorkPlace      string `json:"workPlace,omitempty"`
	Position       string `json:"position,omitempty"`
	WorkPhone      string `json:"workPhone,omitempty"`
	AcademicDegree *int64 `json:"academicDegree,omitempty"`
	Photo          string `json:"photo,omitempty"`
}

type RegistrationReaderUserResponse struct {
	TicketNumber int64  `json:"ticketNumber"`
	Status       string `json:"status"`
}

func (r *RegistrationReaderUserRequest) Validate() error {
	if r.Photo == "" {
		return errors.New("загрузите фото")
	}
	if r.Iin == "" {
		return errors.New("введите ИИН")
	}
	if r.BirthDate == "" {
		return errors.New("введите дату рождения")
	}
	if r.Lastname == "" {
		return errors.New("введите фамилию")
	}
	if r.Firstname == "" {
		return errors.New("введите имя")
	}
	if r.Sex == nil || *r.Sex == 0 {
		return errors.New("укажите пол")
	}
	if r.Nationality == nil || *r.Nationality == 0 {
		return errors.New("укажите национальность")
	}
	if r.Phone == "" {
		return errors.New("введите номер телефона")
	}
	if r.Email == "" {
		return errors.New("введите email")
	}
	if r.Address == "" {
		return errors.New("введите адрес")
	}
	if len(strings.Trim(r.Iin, " ")) != 12 {
		return errors.New("иин неверный")
	}
	birthDate, err := time.Parse("2006-01-02", r.BirthDate)
	if err != nil {
		return errors.New("неверный формат даты рождения")
	}
	beginIin := birthDate.Format("060102")
	if len(strings.Trim(r.Iin, " ")) != 12 {
		return errors.New("иин неверный")
	}
	if beginIin != r.Iin[:6] {
		return errors.New("иин неверный.")
	}
	if !utils.ValidateIIN(r.Iin) {
		return errors.New("иин неверный..")
	}
	return nil
}
func (r *RegistrationReaderUserRequest) ReaderRegistrationRequest(typeCardId int64, photo []byte) (*rcmodel.ReaderRegistrationRequest, error) {
	birthDate, err := time.Parse("2006-01-02", r.BirthDate)
	if err != nil {
		return nil, err
	}
	return &rcmodel.ReaderRegistrationRequest{
		TypeCard:       &typeCardId,
		Iin:            r.Iin,
		Firstname:      r.Firstname,
		Middlename:     r.Middlename,
		Lastname:       r.Lastname,
		BirthDate:      &birthDate,
		Address:        r.Address,
		Email:          r.Email,
		Phone:          r.Phone,
		WorkPhone:      r.WorkPhone,
		WorkPlace:      r.WorkPlace,
		Position:       r.Position,
		Course:         r.Course,
		Faculty:        dictId(r.Faculty),
		AcademicDegree: dictId(r.AcademicDegree),
		Institution:    dictId(r.Institution),
		Education:      dictId(r.Education),
		Speciality:     dictId(r.Speciality),
		Nationality:    dictId(r.Nationality),
		SocialStatus:   dictId(r.SocialStatus),
		Sex:            dictId(r.Sex),
		Photo:          photo,
	}, nil
}
func dictId(id *int64) *int64 {
	if id == nil {
		return nil
	}
	return id
}
