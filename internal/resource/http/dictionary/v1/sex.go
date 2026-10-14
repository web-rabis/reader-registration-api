package v1

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/web-rabis/httperrors"
	"github.com/web-rabis/reader-registration-api/internal/model"
	"github.com/web-rabis/reader-registration-api/internal/resource/http/dictionary/v1/dto"
)

// @Summary Dictionary Sex
// @Description Dictionary Sex
// @Accept json
// @Produce json
// @Tags Dictionary
// @Success 200 {object} model.ListResponse{result=[]dto.DSex}
// @Failure 400 {object} httperrors.Response
// @Failure 401 {object} httperrors.Response
// @Failure 403 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/dictionary/sex [get]
func (res *Resource) getSex(w http.ResponseWriter, r *http.Request) {
	count, sexList, err := res.dictMan.SexList(r.Context(), nil)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	render.JSON(w, r, model.ListResponse{Count: count, Result: dto.NewSexList(sexList)})
}
