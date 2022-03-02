package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(StartTime string, machineID int64) (err error) {
	//machineID 分布式机器ID
	var st time.Time
	//指定起始时间
	st, err = time.Parse("2022-03-01", StartTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

func GenID() int64 {
	return
}
