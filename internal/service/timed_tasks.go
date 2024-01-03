package service

import "github.com/robfig/cron/v3"

type TimedTask struct {
	Spec string
	Fc   func()
}

func TimedTasksStart(c *cron.Cron, tts []TimedTask) {
	for _, tt := range tts {
		_, _ = c.AddFunc(tt.Spec, tt.Fc)
	}

	c.Start()
}
