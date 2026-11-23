package model

import (
	"time"

	"github.com/web-rabis/reader-client/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Reader struct {
	GlobalTicketNumber int64         `json:"globalTicketNumber"`
	TicketNumber       int64         `json:"ticketNumber"`
	RFIDTag            string        `json:"rfidTag"`
	Barcode            string        `json:"barcode"`
	TypeCard           *TypeCard     `json:"typeCard"`
	RegistrationDate   *time.Time    `json:"registrationDate"`
	ReregistrationDate *time.Time    `json:"reregistrationDate"`
	ExpirationDate     *time.Time    `json:"expirationDate"`
	LastVisitDate      *time.Time    `json:"lastVisitDate"`
	CorrectionDate     *time.Time    `json:"correctionDate"`
	ReadersHall        string        `json:"readersHall"`
	LibraryId          int64         `json:"libraryId"`
	BranchId           int64         `json:"branchId"`
	Iin                string        `json:"iin"`
	Firstname          string        `json:"firstname"`
	Middlename         string        `json:"middlename"`
	Lastname           string        `json:"lastname"`
	BirthDateYear      *int32        `json:"birthDateYear"`
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
	Status             string        `json:"status"`
	IsEmployee         bool          `json:"isEmployee"`
	PhotoKey           string        `json:"photoKey"`
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
	PhotoKey       string     `json:"photoKey,omitempty"`
}

type ReaderApproveResult struct {
	OldTicketNumber int64
	TicketNumber    int64
	Barcode         string
	Status          string
	ExpirationDate  *time.Time
}

func ReaderApproveResultFromProto(r *protobuf.ReaderApproveResponse) *ReaderApproveResult {
	if r == nil {
		return nil
	}
	var expirationDate *time.Time
	if r.ExpirationDDate != nil {
		t := r.ExpirationDDate.AsTime()
		expirationDate = &t
	}
	return &ReaderApproveResult{
		OldTicketNumber: r.OldTicketNumber,
		TicketNumber:    r.TicketNumber,
		Barcode:         r.Barcode,
		Status:          r.Status,
		ExpirationDate:  expirationDate,
	}
}

func (r *Reader) ToProto() *protobuf.Reader {
	if r == nil {
		return nil
	}
	var registrationDate *timestamppb.Timestamp
	if r.RegistrationDate != nil {
		registrationDate = timestamppb.New(*r.RegistrationDate)
	}
	var reregistrationDate *timestamppb.Timestamp
	if r.ReregistrationDate != nil {
		reregistrationDate = timestamppb.New(*r.ReregistrationDate)
	}
	var expirationDate *timestamppb.Timestamp
	if r.ExpirationDate != nil {
		expirationDate = timestamppb.New(*r.ExpirationDate)
	}
	var lastVisitDate *timestamppb.Timestamp
	if r.LastVisitDate != nil {
		lastVisitDate = timestamppb.New(*r.LastVisitDate)
	}
	var correctionDate *timestamppb.Timestamp
	if r.CorrectionDate != nil {
		correctionDate = timestamppb.New(*r.CorrectionDate)
	}
	var birthDate *timestamppb.Timestamp
	if r.BirthDate != nil {
		birthDate = timestamppb.New(*r.BirthDate)
	}
	return &protobuf.Reader{
		TicketNumber:       r.TicketNumber,
		RfidTag:            r.RFIDTag,
		Barcode:            r.Barcode,
		TypeCard:           r.TypeCard.ToProto(),
		RegistrationDate:   registrationDate,
		ReregistrationDate: reregistrationDate,
		ExpirationDate:     expirationDate,
		LastVisitDate:      lastVisitDate,
		CorrectionDate:     correctionDate,
		ReadersHall:        r.ReadersHall,
		LibraryId:          r.LibraryId,
		BranchId:           r.BranchId,
		Iin:                r.Iin,
		Firstname:          r.Firstname,
		Middlename:         r.Middlename,
		Lastname:           r.Lastname,
		BirthDateYear:      r.BirthDateYear,
		BirthDate:          birthDate,
		Sex:                r.Sex.ToProto(),
		SocialStatus:       r.SocialStatus.ToProto(),
		Nationality:        r.Nationality.ToProto(),
		Speciality:         r.Speciality.ToProto(),
		Address:            r.Address,
		Email:              r.Email,
		Phone:              r.Phone,
		WorkPhone:          r.WorkPhone,
		Education:          r.Education.ToProto(),
		PlaceEmployment:    r.PlaceEmployment,
		Position:           r.Position,
		Institution:        r.Institution.ToProto(),
		AcademicDegree:     r.AcademicDegree.ToProto(),
		Course:             r.Course,
		Faculty:            r.Faculty.ToProto(),
		Note:               r.Note,
		Status:             r.Status,
		IsEmployee:         r.IsEmployee,
		PhotoKey:           r.PhotoKey,
	}
}
func ReaderFromProto(r *protobuf.Reader) *Reader {
	if r == nil {
		return nil
	}
	var registrationdate *time.Time
	var reregistrationdate *time.Time
	var expirationdate *time.Time
	var lastvisitdate *time.Time
	var correctiondate *time.Time
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
	if r.CorrectionDate != nil {
		t := r.CorrectionDate.AsTime()
		correctiondate = &t
	}
	if r.BirthDate != nil {
		t := r.BirthDate.AsTime()
		birthdate = &t
	}
	return &Reader{
		GlobalTicketNumber: r.GlobalTicketNumber,
		TicketNumber:       r.TicketNumber,
		RFIDTag:            r.RfidTag,
		Barcode:            r.Barcode,
		TypeCard:           TypeCardFromProto(r.TypeCard),
		RegistrationDate:   registrationdate,
		ReregistrationDate: reregistrationdate,
		ExpirationDate:     expirationdate,
		LastVisitDate:      lastvisitdate,
		CorrectionDate:     correctiondate,
		ReadersHall:        r.ReadersHall,
		LibraryId:          r.LibraryId,
		BranchId:           r.BranchId,
		Iin:                r.Iin,
		Firstname:          r.Firstname,
		Middlename:         r.Middlename,
		Lastname:           r.Lastname,
		BirthDateYear:      r.BirthDateYear,
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
		Status:             r.Status,
		IsEmployee:         r.IsEmployee,
		PhotoKey:           r.PhotoKey,
	}
}
func ReaderListFromProto(l []*protobuf.Reader) []*Reader {
	list := make([]*Reader, len(l))
	for i, d := range l {
		list[i] = ReaderFromProto(d)
	}
	return list
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
	}
}
