// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/repository.go.tpl

package repository

import (
    "context"
	"github.com/JonasMuehlmann/bntp.go/domain"
)

type DocumentRepository interface {
	New(...any) (DocumentRepository, error)

	Add(context.Context, []domain.Document) (numAffectedRecords int, newID int, err error)
	Replace(context.Context, []domain.Document) error
	UpdateWhere(context.Context, domain.DocumentFilter, map[domain.DocumentField]domain.DocumentUpdateOperation) (numAffectedRecords int, err error)
	Delete(context.Context, []domain.Document) error
	DeleteWhere(context.Context, domain.DocumentFilter) (numAffectedRecords int, err error)
	CountWhere(context.Context, domain.DocumentFilter) int
	CountAll(context.Context) int
	DoesExist(context.Context, domain.Document) bool
	DoesExistWhere(context.Context, domain.DocumentFilter) bool
	GetWhere(context.Context, domain.DocumentFilter) []domain.Document
	GetFirstWhere(context.Context, domain.DocumentFilter) domain.Document
	GetAll(context.Context) []domain.Document
}
