package mClock

import (
	"github.com/robfig/cron/v3"
)

/*

go mClock.New(mClock.OptType{
		Func: SysReStart,
		Spec: "0 18 5 3,9,15,21,27 * ? ", // 每个月的 3 日、9 日 每隔5天的凌晨 5:18 重启一次 Linux 系统
})

*/

// http://cron.ciding.cc
type OptType struct {
	Func func()
	Spec string // cron 表达式
}

func New(opt OptType) {
	if opt.Func == nil {
		panic("Func 不能为空")
	}

	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	// crontab := cron.New()  默认从分开始进行时间调度
	crontab := cron.New(cron.WithSeconds()) // 精确到秒

	// 添加定时任务,
	_, err := crontab.AddFunc(opt.Spec, opt.Func)
	if err != nil {
		panic("定时任务创建失败 , Spec 不合法")
	}
	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {} // 阻塞主线程停止
}
