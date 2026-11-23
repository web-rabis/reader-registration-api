package order_client

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
)

// Ключи grpc-metadata, которые ждёт reader (см.
// reader/internal/server/grpc/interceptor.go) - вызывающий сервис уже
// прошёл аутентификацию через SSO и передаёт активный tenant-контекст
// явно, без auth-интерсепторов на стороне reader.
const (
	metaLibraryID    = "library_id"
	metaBranchID     = "branch_id"
	metaDepartmentID = "department_id"
)

// WithTenant кладёт library_id/branch_id/department_id в исходящие
// grpc-metadata запроса к reader. Вызывающий сервис берёт эти значения
// из активного контекста JWT claims SSO (library_id/branch_id/
// department_id верхнего уровня, см. sso/CONTRACT.md) и оборачивает ctx
// этим вызовом перед каждым обращением к ReaderService/VisitService -
// сам reader-client ничего не проверяет и не выбирает "активный"
// контекст, просто передаёт то, что ему дали.
//
// departmentId - опционален; 0 означает "не передавать" (reader читает
// отсутствующий ключ как 0 = нет активного контекста, что для него не
// ошибка аутентификации, см. CONTRACT.md).
func WithTenant(ctx context.Context, libraryId, branchId, departmentId int64) context.Context {
	return metadata.AppendToOutgoingContext(ctx,
		metaLibraryID, strconv.FormatInt(libraryId, 10),
		metaBranchID, strconv.FormatInt(branchId, 10),
		metaDepartmentID, strconv.FormatInt(departmentId, 10),
	)
}
