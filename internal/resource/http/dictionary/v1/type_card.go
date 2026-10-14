package v1

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/web-rabis/httperrors"
	rcmodel "github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-registration-api/internal/model"
)

// @Summary Dictionary Type Card
// @Description Dictionary Type Card
// @Accept json
// @Produce json
// @Tags Dictionary
// @Success 200 {object} model.ListResponse{result=[]rcmodel.TypeCard}
// @Failure 400 {object} httperrors.Response
// @Failure 401 {object} httperrors.Response
// @Failure 403 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/dictionary/type-card [get]
func (res *Resource) getTypeCard(w http.ResponseWriter, r *http.Request) {
	paging, err := rcmodel.PagingParseFromHttp(r)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	paging.SortKey = "type"
	count, l, err := res.dictMan.TypeCardList(r.Context(), paging)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	render.JSON(w, r, model.ListResponse{Count: count, Result: l})
}
