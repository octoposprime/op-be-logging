package infrastructure

import (
	mo "github.com/octoposprime/op-be-logging/internal/domain/model/object"
)

var LogSortMap map[mo.LogSortField]string = map[mo.LogSortField]string{
	mo.LogSortFieldId:          "id",
	mo.LogSortFieldServiceName: "service_name",
	mo.LogSortFieldEventDate:   "verified_at",
	mo.LogSortFieldCreatedAt:   "created_at",
	mo.LogSortFieldUpdatedAt:   "updated_at",
}
