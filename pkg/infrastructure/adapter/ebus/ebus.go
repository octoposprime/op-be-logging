package infrastructure

import (
	"context"
	"time"

	me "github.com/octoposprime/op-be-logging/internal/domain/model/entity"
	map_ebus "github.com/octoposprime/op-be-logging/pkg/infrastructure/mapper/ebus"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tredis "github.com/octoposprime/op-be-shared/tool/redis"
	tserialize "github.com/octoposprime/op-be-shared/tool/serialize"
)

type EBusAdapter struct {
	redisClient *tredis.RedisClient
}

func NewEBusAdapter(redisClient *tredis.RedisClient) EBusAdapter {
	adapter := EBusAdapter{
		redisClient: redisClient,
	}
	return adapter
}

// Listen listens to the redis messaging queue and calls the given callBack function for each received log.
func (a EBusAdapter) Listen(ctx context.Context, channelName string, callBack func(channelName string, logData me.LogData)) {
	for {
		result, err := a.redisClient.BLPop(ctx, 0*time.Second, channelName).Result()
		if err != nil {
			continue
		}
		inChannelName := result[0]
		logData := tserialize.NewSerializer(me.LogData{}).FormJson(result[1]).(*pb.LogData)
		go callBack(inChannelName, *map_ebus.NewLogData(logData).ToEntity())
	}
}
