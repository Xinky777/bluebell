package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(StartTime string, machineID int64) (err error) {
	//machineID 分布式机器ID
	var st time.Time
	//指定起始时间
	st, err = time.Parse("2006-01-02", StartTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
