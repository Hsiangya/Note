package handler

import (
	"context"
	"errors"
	"github.com/bwmarrin/snowflake"
	"snowflake/dao/config"
	"snowflake/proto"
	"time"
)

type SnowFlakeSrv struct {
	proto.UnimplementedSnowFlakeServer
}

func (s SnowFlakeSrv) GetId(ctx context.Context, n *proto.Num) (*proto.Ids, error) {

	if n.Num < 1 {
		return nil, errors.New("生成的ID个数必须大于0")
	}

	snowflake.Epoch = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	node, err := snowflake.NewNode(config.Conf.MachineID)
	if err != nil {
		return nil, err
	}
	result := &proto.Ids{}
	for i := int64(0); i < n.Num; i++ {
		id := node.Generate().Int64()
		result.Id = append(result.Id, &proto.OrderId{Id: id})
	}
	return result, nil
}
