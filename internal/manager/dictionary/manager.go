package dictionary

import (
	"context"

	readerclient "github.com/web-rabis/reader-client/client"
	rcmodel "github.com/web-rabis/reader-client/model"
)

type IManager interface {
	AcademicDegreeList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error)
	EducationList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error)
	FacultyList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error)
	InstitutionList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error)
	NationalityList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error)
	SexList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error)
	SocialStatusList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.SocialStatus, error)
	SpecialityList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error)
	TypeCardList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.TypeCard, error)
}
type Manager struct {
	dictCli readerclient.DictionaryService
}

func NewManager(dictCli readerclient.DictionaryService) *Manager {
	return &Manager{
		dictCli: dictCli,
	}
}

func (m *Manager) AcademicDegreeList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return m.dictCli.AcademicDegreeList(ctx, paging)
}
func (m *Manager) EducationList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return m.dictCli.EducationList(ctx, paging)
}
func (m *Manager) FacultyList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return m.dictCli.FacultyList(ctx, paging)
}
func (m *Manager) InstitutionList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return m.dictCli.InstitutionList(ctx, paging)
}
func (m *Manager) NationalityList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	count, list, err := m.dictCli.NationalityList(ctx, paging)
	if err != nil {
		return 0, nil, err
	}
	var priorityNationality []*rcmodel.Dictionary
	var finaly []*rcmodel.Dictionary
	for _, v := range list {
		if v.Name == "Қазақ" ||
			v.Name == "Русский" ||
			v.Name == "Украин" ||
			v.Name == "Узбек" ||
			v.Name == "Уйгур" ||
			v.Name == "Армян" ||
			v.Name == "Азер" ||
			v.Name == "Грек" ||
			v.Name == "Дунган" ||
			v.Name == "Кыpгыз" ||
			v.Name == "Курд" ||
			v.Name == "Немец" ||
			v.Name == "Корей" ||
			v.Name == "Татар" ||
			v.Name == "Турок" ||
			v.Name == "Чечен" {
			priorityNationality = append(priorityNationality, v)
		} else {
			finaly = append(finaly, v)
		}
	}
	return count, append(priorityNationality, finaly...), nil
}
func (m *Manager) SexList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return m.dictCli.SexList(ctx, paging)
}
func (m *Manager) SocialStatusList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.SocialStatus, error) {
	return m.dictCli.SocialStatusList(ctx, paging)
}
func (m *Manager) SpecialityList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return m.dictCli.SpecialityList(ctx, paging)
}
func (m *Manager) TypeCardList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.TypeCard, error) {
	return m.dictCli.TypeCardList(ctx, paging)
}
