package service

import (
	"time"

	"github.com/moshaoli688/miaospeed/service/taskpoll"
	"github.com/moshaoli688/miaospeed/utils"
)

var SpeedTaskPoll *taskpoll.TaskPollController
var ConnTaskPoll *taskpoll.TaskPollController

func StartTaskServer() {
	SpeedTaskPoll = taskpoll.NewTaskPollController("SpeedPoll", 1, time.Duration(utils.GCFG.PauseSecond)*time.Second, 200*time.Millisecond)
	ConnTaskPoll = taskpoll.NewTaskPollController("ConnPoll", utils.GCFG.ConnTaskTreading, 0, 200*time.Millisecond)

	go SpeedTaskPoll.Start()
	go ConnTaskPoll.Start()
}
