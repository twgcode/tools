/**
@Author: guantingwei
@Email: guantingwei@sixents.com
@Date: 2022/12/9 11:12
@Description:
*/

package simplesql

import (
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewNullTimeFromTimestampPB 根据 *timestamppb.Timestamp 创建 sql.NullTime
func NewNullTimeFromTimestampPB(source *timestamppb.Timestamp) (result sql.NullTime) {
	if source == nil {
		result.Valid = false
		return
	}
	result.Valid = true
	result.Time = source.AsTime().Local()
	return
}
