package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"

	"github.com/web-rabis/httperrors"
	"github.com/web-rabis/reader-registration-api/internal/resource/http/reader/v1/dto"
	"github.com/web-rabis/reader-registration-api/internal/utils"
)

// @Summary Create Reader User
// @Description Create Reader User
// @Accept json
// @Produce json
// @Tags ReaderUser
// @Param body body dto.RegistrationReaderUserRequest true "Request"
// @Success 200 {object} dto.RegistrationReaderUserResponse
// @Failure 400 {object} httperrors.Response
// @Failure 401 {object} httperrors.Response
// @Failure 403 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /api/v1/reader/registration [post]
func (res *Resource) registrationReader(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request dto.RegistrationReaderUserRequest
	var err error

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err := request.Validate(); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	photoBytes, err := utils.ResizeImage(request.Photo, 1024, 1024, utils.ModeFit, 90)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	req, err := request.ReaderRegistrationRequest(res.readerTypeCard, photoBytes)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	ticketNumber, status, err := res.readerMan.ReaderRegistration(ctx, req)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	render.JSON(w, r, dto.RegistrationReaderUserResponse{
		TicketNumber: ticketNumber,
		Status:       status,
	})
}
