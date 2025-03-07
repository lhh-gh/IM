package myMongo

import (
	"context"
	"fmt"
	"time"

	"github/lhh-gh/IM/internal/common/constant"
	"github/lhh-gh/IM/internal/common/dao/myMongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetBigGroupMessage 从Timeline数据库获得 读扩散 的大群组消息
func (db *DB) GetBigGroupMessage(ctx context.Context, groupID uint32, clientTimestamp time.Time) ([]models.UserTimeline, error) {
	filter := mongo.Pipeline{}

	matchStage := bson.D{
		{"$match", bson.D{
			{"$and", bson.A{
				bson.D{{"receiver_id", groupID}},
				bson.D{{"type", constant.BIG_GROUP_CHAT}},
				bson.D{{"group_id", groupID}},
			}},
		}},
	}
	filter = append(filter, matchStage)

	if clientTimestamp != time.Unix(0, 0) {
		timestampMatch := bson.D{
			{"$match", bson.D{
				{"timestamp", bson.D{{"$gt", clientTimestamp}}},
			}},
		}
		filter = append(filter, timestampMatch)
	}

	sortStage := bson.D{
		{"$sort", bson.D{
			{"timestamp", -1},
		}},
	}
	filter = append(filter, sortStage)

	var timelines []models.UserTimeline
	if err := db.TimeLine.Aggregate(ctx, &timelines, filter); err != nil {
		return nil, fmt.Errorf("failed to retrieve user timeline: %w", err)
	}
	return timelines, nil
}
