package persistent

import (
	"asynq-implementation-example/persistent/notification"
	"asynq-implementation-example/persistent/repository"
	reqtest "asynq-implementation-example/persistent/reqTest"
)

type Persistent struct {
	Repository   repository.Repository
	Notification notification.Notification
	ReqTest      reqtest.ReqTest
}

func NewPersistent(repo repository.Repository, notif notification.Notification, rt reqtest.ReqTest) Persistent {
	return Persistent{
		Repository:   repo,
		Notification: notif,
		ReqTest:      rt,
	}
}
