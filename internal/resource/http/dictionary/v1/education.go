package v1

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/web-rabis/httperrors"
	rcmodel "github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-registration-api/internal/model"
	"github.com/web-rabis/reader-registration-api/internal/resource/http/dictionary/v1/dto"
)

// @Summary Dictionary Education
// @Description Dictionary Education
// @Accept json
// @Produce json
// @Tags Dictionary
// @Success 200 {object} model.ListResponse{result=[]dto.DEducation}
// @Failure 400 {object} httperrors.Response
// @Failure 401 {object} httperrors.Response
// @Failure 403 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/dictionary/education [get]
func (res *Resource) getEducation(w http.ResponseWriter, r *http.Request) {
	paging, err := rcmodel.PagingParseFromHttp(r)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	count, l, err := res.dictMan.EducationList(r.Context(), paging)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	render.JSON(w, r, model.ListResponse{Count: count, Result: dto.NewEducationList(l)})
}
