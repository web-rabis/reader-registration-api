package v1

import (
	"github.com/go-chi/chi/v5"

	"github.com/web-rabis/reader-registration-api/internal/manager/dictionary"
)

type Resource struct {
	path    string
	dictMan dictionary.IManager
}

func NewDictionaryResource(basePath string, dictMan dictionary.IManager) *Resource {
	return &Resource{
		path:    basePath,
		dictMan: dictMan,
	}
}

func (res *Resource) Path() string {
	return res.path
}

func (res *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/academic-degree", res.getAcademicDegree)
		r.Get("/education", res.getEducation)
		r.Get("/faculty", res.getFaculty)
		r.Get("/institution", res.getInstitution)
		r.Get("/nationality", res.getNationality)
		r.Get("/sex", res.getSex)
		r.Get("/social-status", res.getSocialStatus)
		r.Get("/speciality", res.getSpeciality)
		r.Get("/type-card", res.getTypeCard)
	})

	return r
}
