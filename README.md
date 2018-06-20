# 任务管理-小程序
小程序端接口说明

注：若服务器内部出错，会返回
```
500 Internal Server Error
{
	"cnmsg":"很抱歉,服务器出错了"
}
```

## 新任务
### 获取新任务 [/task/miniprogram/new-task]
```
request:
response:
	200 OK
	{
		"count_data": 2,	// 总任务数
		"data": [
		{
			"task_id":		123456,			// 任务ID
			"title":		"集合演习",		// 任务主题
			"detail":		"任务详情",		// 任务详情
			"mem_count":	50, 			// 目标征集人数
			"launch_admin":	"海珠人武部",		// 发起单位/组织
			"launch_datetime":	"2018-06-20 10:00:00",	// 发起时间
			"gather_datetime": 	"2018-06-21 12:00:00",	// 集合时间
			"finish_datetime":	"2018-06-21 18:00:00",	// 结束时间
			"place_name":	"中山大学北门",		// 集合地点
			"place_lat":	23.333,			// 集合地点纬度
			"place_lng":	123.333			// 集合地点经度
		},
		...
		]
	}
```

### 接受任务 [/task/miniprogram/accept]
```
request:
{
	"task_id": 123456
}
response:
	200 OK
```

### 拒绝任务 [/task/miniprogram/refuse]
```
request:
{
	"task_id": 123456
}
response:
	200 OK
```

## 执行中的任务
### 获取执行中的任务 [/task/miniprogram/working]
```
request:
response:
	200 OK
	{
		"count_data": 2,	// 总任务数
		"data": [
		{
			"task_id":		123456,			// 任务ID
			"title":		"集合演习",		// 任务主题
			"detail":		"任务详情",		// 任务详情
			"mem_count":	50, 			// 目标征集人数
			"launch_admin":	"海珠人武部",		// 发起单位/组织
			"launch_datetime":	"2018-06-20 10:00:00",	// 发起时间
			"gather_datetime": 	"2018-06-21 12:00:00",	// 集合时间
			"finish_datetime":	"2018-06-21 18:00:00",	// 结束时间
			"place_name":	"中山大学北门",		// 集合地点
			"place_lat":	23.333,			// 集合地点纬度
			"place_lng":	123.333			// 集合地点经度
		},
		...
		]
	}
```

### 获取接受该任务的队友 [/task/miniprogram/soldiers/{task_id}]

返回同组织/单位的接受该任务的队友和负责人?



### 签到 [/task/miniprogram/check]
```
request:
{
	"task_id":		123456,
	"place_lat":	23.333,
	"place_lng":	123.333
}
response:
	200 OK
	{
		"enmsg":	"ok",
		"cnmsg":	"签到成功"
	}
	{
		"enmsg":	"error",
		"cnmsg":	"你没有成功签到噢"
	}
```

## 已完成任务
### 提交任务总结 [/task/miniprogram/summary]
request: task_id, summary
### 查看已完成任务列表 [/task/miniprogram/done]
### 查看已完成任务详情 [/task/miniprogram/done/{task_id}]