package handler

import (
	"context"
	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/bwmarrin/snowflake"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snowflake/dao/config"
	"snowflake/proto"
	"time"
)

type SnowFlakeSrv struct {
	proto.UnimplementedSnowFlakeServer
}

func (s SnowFlakeSrv) GetId(ctx context.Context, n *proto.Num) (*proto.Ids, error) {
	e, b := api.Entry(config.Conf.Name, api.WithTrafficType(base.Inbound))
	if b != nil {
		return nil, status.Errorf(codes.ResourceExhausted, "Request blocked by Sentinel: %v", b.Error())
	}
	defer e.Exit()

	if n.Num < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "生成的ID个数必须大于0")
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
