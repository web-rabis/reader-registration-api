package v1

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/web-rabis/httperrors"
	rcmodel "github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-registration-api/internal/model"
)

// @Summary Dictionary Nationality
// @Description Dictionary Nationality
// @Accept json
// @Produce json
// @Tags Dictionary
// @Success 200 {object} model.ListResponse{result=[]rcmodel.Dictionary}
// @Failure 400 {object} httperrors.Response
// @Failure 401 {object} httperrors.Response
// @Failure 403 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/dictionary/nationality [get]
func (res *Resource) getNationality(w http.ResponseWriter, r *http.Request) {
	paging, err := rcmodel.PagingParseFromHttp(r)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	count, l, err := res.dictMan.NationalityList(r.Context(), paging)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	render.JSON(w, r, model.ListResponse{Count: count, Result: l})
}
