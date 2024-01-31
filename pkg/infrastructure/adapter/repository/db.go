package infrastructure

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
	map_repo "github.com/octoposprime/op-be-logging/pkg/infrastructure/mapper/repository"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
	tserialize "github.com/octoposprime/op-be-shared/tool/serialize"
)

type DbAdapter struct {
	*tgorm.GormClient
}

func NewDbAdapter(dbClient *tgorm.GormClient) DbAdapter {
	adapter := DbAdapter{
		dbClient,
	}

	err := dbClient.DbClient.AutoMigrate(&map_repo.LogData{})
	if err != nil {
		panic(err)
	}

	return adapter
}

func (a DbAdapter) Log(ctx context.Context, logData me.LogData) {
	logDbMapper := map_repo.NewLogDataFromEntity(&logData)
	result := a.DbClient.Save(&logDbMapper)
	if result.Error != nil {
		fmt.Println(tserialize.NewSerializer(logData).ToJson())
		fmt.Println(result.Error)
	}
}

func (a DbAdapter) GetLogsByFilter(ctx context.Context, logDataFilter me.LogDataFilter) ([]me.LogData, error) {
	var logDatasDbMapper map_repo.LogDatas
	var filter map_repo.LogData
	qry := a.DbClient
	if logDataFilter.Id.String() != "" && logDataFilter.Id != (uuid.UUID{}) {
		filter.ID = logDataFilter.Id
	}
	if !logDataFilter.EventDateFrom.IsZero() && !logDataFilter.EventDateTo.IsZero() {
		qry = qry.Where("event_date between ? and ?", logDataFilter.EventDateFrom, logDataFilter.EventDateTo)
	}
	if logDataFilter.LogType != 0 {
		filter.LogType = int(logDataFilter.LogType)
	}
	if logDataFilter.ServiceName != "" {
		filter.ServiceName = logDataFilter.ServiceName
	}
	if logDataFilter.Path != "" {
		filter.Path = logDataFilter.Path
	}
	if logDataFilter.UserId != "" {
		filter.UserId = logDataFilter.UserId
	}
	if !logDataFilter.CreatedAtFrom.IsZero() && !logDataFilter.CreatedAtTo.IsZero() {
		qry = qry.Where("created_at between ? and ?", logDataFilter.CreatedAtFrom, logDataFilter.CreatedAtTo)
	}
	if !logDataFilter.UpdatedAtFrom.IsZero() && !logDataFilter.UpdatedAtTo.IsZero() {
		qry = qry.Where("updated_at between ? and ?", logDataFilter.UpdatedAtFrom, logDataFilter.UpdatedAtTo)
	}
	if logDataFilter.SearchText != "" {
		qry = qry.Where(
			qry.Where("UPPER(service_name) LIKE UPPER(?)", "%"+logDataFilter.SearchText+"%").
				Or("UPPER(message) LIKE UPPER(?)", "%"+logDataFilter.SearchText+"%"),
		)
	}
	qry = qry.Where(filter)
	if logDataFilter.Limit != 0 {
		qry = qry.Limit(logDataFilter.Limit)
	}
	if logDataFilter.Offset != 0 {
		qry = qry.Offset(logDataFilter.Offset)
	}
	if logDataFilter.SortType != "" && logDataFilter.SortField != 0 {
		sortStr := map_repo.LogSortMap[logDataFilter.SortField]
		if logDataFilter.SortType == "desc" || logDataFilter.SortType == "DESC" {
			sortStr += " desc"
		} else {
			sortStr += " asc"
		}
		qry = qry.Order(sortStr)
	}
	result := qry.Find(&logDatasDbMapper)
	if result.Error != nil {
		return []me.LogData{}, result.Error
	}
	return logDatasDbMapper.ToEntities(), nil
}
