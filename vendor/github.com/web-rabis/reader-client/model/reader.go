package model

import (
	"time"

	"github.com/web-rabis/reader-client/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Reader struct {
	TicketNumber       int64         `json:"ticketNumber"`
	RFID               string        `json:"rfid"`
	Barcode            string        `json:"barcode"`
	TypeCard           *TypeCard     `json:"typeCard"`
	RegistrationDate   *time.Time    `json:"registrationDate"`
	ReregistrationDate *time.Time    `json:"reregistrationDate"`
	ExpirationDate     *time.Time    `json:"expirationDate"`
	LastVisitDate      *time.Time    `json:"lastVisitDate"`
	DateCorrection     *time.Time    `json:"dateCorrection"`
	ReadersHall        string        `json:"readersHall"`
	Status             string        `json:"status"`
	BranchId           string        `json:"branchId"`
	Iin                string        `json:"iin"`
	Firstname          string        `json:"firstname"`
	Middlename         string        `json:"middlename"`
	Lastname           string        `json:"lastname"`
	YearBirth          int64         `json:"yearBirth"`
	BirthDate          *time.Time    `json:"birthDate"`
	Sex                *Dictionary   `json:"sex"`
	SocialStatus       *SocialStatus `json:"socialStatus"`
	Nationality        *Dictionary   `json:"nationality"`
	Speciality         *Dictionary   `json:"speciality"`
	Address            string        `json:"address"`
	Email              string        `json:"email"`
	Phone              string        `json:"phone"`
	WorkPhone          string        `json:"workPhone"`
	Education          *Dictionary   `json:"education"`
	PlaceEmployment    string        `json:"placeEmployment"`
	Position           string        `json:"position"`
	Institution        *Dictionary   `json:"institution"`
	AcademicDegree     *Dictionary   `json:"academicDegree"`
	Course             string        `json:"course"`
	Faculty            *Dictionary   `json:"faculty"`
	Note               string        `json:"note"`
	PriznakUdal        string        `json:"priznakUdal"`
	Department         string        `json:"department"`
	IsEmployee         bool          `json:"isEmployee"`
}

type ReaderRegistrationRequest struct {
	Iin            string     `json:"iin,omitempty"`
	Lastname       string     `json:"lastname,omitempty"`
	Middlename     string     `json:"middlename,omitempty"`
	Firstname      string     `json:"firstname,omitempty"`
	BirthDate      *time.Time `json:"birthDate,omitempty"`
	Address        string     `json:"address,omitempty"`
	Phone          string     `json:"phone,omitempty"`
	Email          string     `json:"email,omitempty"`
	Sex            *int64     `json:"sex,omitempty"`
	Nationality    *int64     `json:"nationality,omitempty"`
	Education      *int64     `json:"education,omitempty"`
	SocialStatus   *int64     `json:"socialStatus,omitempty"`
	Speciality     *int64     `json:"speciality,omitempty"`
	Institution    *int64     `json:"institution,omitempty"`
	Faculty        *int64     `json:"faculty,omitempty"`
	TypeCard       *int64     `json:"typeCard,omitempty"`
	Course         string     `json:"course,omitempty"`
	WorkPlace      string     `json:"workPlace,omitempty"`
	Position       string     `json:"position,omitempty"`
	WorkPhone      string     `json:"workPhone,omitempty"`
	AcademicDegree *int64     `json:"academicDegree,omitempty"`
	Photo          []byte     `json:"photo,omitempty"`
}

func ReaderFromProto(r *protobuf.Reader) *Reader {
	if r == nil {
		return nil
	}
	var registrationdate *time.Time
	var reregistrationdate *time.Time
	var expirationdate *time.Time
	var lastvisitdate *time.Time
	var datecorrection *time.Time
	var birthdate *time.Time
	if r.RegistrationDate != nil {
		t := r.RegistrationDate.AsTime()
		registrationdate = &t
	}
	if r.ReregistrationDate != nil {
		t := r.ReregistrationDate.AsTime()
		reregistrationdate = &t
	}
	if r.ExpirationDate != nil {
		t := r.ExpirationDate.AsTime()
		expirationdate = &t
	}
	if r.LastVisitDate != nil {
		t := r.LastVisitDate.AsTime()
		lastvisitdate = &t
	}
	if r.DateCorrection != nil {
		t := r.DateCorrection.AsTime()
		datecorrection = &t
	}
	if birthdate != nil {
		t := r.BirthDate.AsTime()
		birthdate = &t
	}
	return &Reader{
		TicketNumber:       r.TicketNumber,
		RFID:               r.Rfid,
		Barcode:            r.Barcode,
		TypeCard:           TypeCardFromProto(r.TypeCard),
		RegistrationDate:   registrationdate,
		ReregistrationDate: reregistrationdate,
		ExpirationDate:     expirationdate,
		LastVisitDate:      lastvisitdate,
		DateCorrection:     datecorrection,
		ReadersHall:        r.ReadersHall,
		Status:             r.Status,
		BranchId:           r.BranchId,
		Iin:                r.Iin,
		Firstname:          r.Firstname,
		Middlename:         r.Middlename,
		Lastname:           r.Lastname,
		YearBirth:          r.YearBirth,
		BirthDate:          birthdate,
		Sex:                DictionaryFromProto(r.Sex),
		SocialStatus:       SocialStatusFromProto(r.SocialStatus),
		Nationality:        DictionaryFromProto(r.Nationality),
		Speciality:         DictionaryFromProto(r.Speciality),
		Address:            r.Address,
		Email:              r.Email,
		Phone:              r.Phone,
		WorkPhone:          r.WorkPhone,
		Education:          DictionaryFromProto(r.Education),
		PlaceEmployment:    r.PlaceEmployment,
		Position:           r.Position,
		Institution:        DictionaryFromProto(r.Institution),
		AcademicDegree:     DictionaryFromProto(r.AcademicDegree),
		Course:             r.Course,
		Faculty:            DictionaryFromProto(r.Faculty),
		Note:               r.Note,
		PriznakUdal:        r.PriznakUdal,
		Department:         r.Department,
		IsEmployee:         r.IsEmployee,
	}
}

func (r *ReaderRegistrationRequest) ToProto() *protobuf.ReaderRegistrationRequest {
	if r == nil {
		return nil
	}
	var birthDate *timestamppb.Timestamp
	if r.BirthDate != nil {
		birthDate = timestamppb.New(*r.BirthDate)
	}
	return &protobuf.ReaderRegistrationRequest{
		Iin:            r.Iin,
		Lastname:       r.Lastname,
		Middlename:     r.Middlename,
		Firstname:      r.Firstname,
		BirthDate:      birthDate,
		Address:        r.Address,
		Phone:          r.Phone,
		Email:          r.Email,
		Sex:            r.Sex,
		Nationality:    r.Nationality,
		Education:      r.Education,
		SocialStatus:   r.SocialStatus,
		Speciality:     r.Speciality,
		Institution:    r.Institution,
		Faculty:        r.Faculty,
		TypeCard:       r.TypeCard,
		Course:         r.Course,
		WorkPlace:      r.WorkPlace,
		Position:       r.Position,
		WorkPhone:      r.WorkPhone,
		AcademicDegree: r.AcademicDegree,
		Photo:          r.Photo,
	}
}

func ReaderListFromProto(l []*protobuf.Reader) []*Reader {
	list := make([]*Reader, len(l))
	for i, d := range l {
		list[i] = ReaderFromProto(d)
	}
	return list
}
