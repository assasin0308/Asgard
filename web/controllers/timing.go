package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"Asgard/client"
	"Asgard/constants"
	"Asgard/models"
	"Asgard/services"
)

type TimingController struct {
	timingService  *services.TimingService
	agentService   *services.AgentService
	groupService   *services.GroupService
	moniterService *services.MonitorService
	archiveService *services.ArchiveService
}

func NewTimingController() *TimingController {
	return &TimingController{
		timingService:  services.NewTimingService(),
		agentService:   services.NewAgentService(),
		groupService:   services.NewGroupService(),
		moniterService: services.NewMonitorService(),
		archiveService: services.NewArchiveService(),
	}
}

func (c *TimingController) formatTiming(info *models.Timing) map[string]interface{} {
	data := map[string]interface{}{
		"ID":        info.ID,
		"Name":      info.Name,
		"GroupID":   info.GroupID,
		"AgentID":   info.AgentID,
		"Dir":       info.Dir,
		"Program":   info.Program,
		"Args":      info.Args,
		"StdOut":    info.StdOut,
		"StdErr":    info.StdErr,
		"Time":      info.Time.Format(TimeLayout),
		"Timeout":   info.Timeout,
		"IsMonitor": info.IsMonitor,
		"Status":    info.Status,
	}
	group := c.groupService.GetGroupByID(info.GroupID)
	if group != nil {
		data["GroupName"] = group.Name
	} else {
		data["GroupName"] = ""
	}
	agent := c.agentService.GetAgentByID(info.AgentID)
	if agent != nil {
		data["AgentName"] = fmt.Sprintf("%s:%s(%s)", agent.IP, agent.Port, agent.Alias)
	} else {
		data["AgentName"] = ""
	}
	return data
}

func (c *TimingController) List(ctx *gin.Context) {
	groupID := DefaultInt(ctx, "group_id", 0)
	agentID := DefaultInt(ctx, "agent_id", 0)
	status := DefaultInt(ctx, "status", -99)
	name := ctx.Query("name")
	page := DefaultInt(ctx, "page", 1)
	where := map[string]interface{}{
		"status": status,
	}
	querys := []string{}
	if groupID != 0 {
		where["group_id"] = groupID
		querys = append(querys, "group_id="+strconv.Itoa(groupID))
	}
	if agentID != 0 {
		where["agent_id"] = agentID
		querys = append(querys, "agent_id="+strconv.Itoa(agentID))
	}
	if status != -99 {
		querys = append(querys, "status="+strconv.Itoa(status))
	}
	if name != "" {
		where["name"] = name
		querys = append(querys, "name="+name)
	}
	timingList, total := c.timingService.GetTimingPageList(where, page, PageSize)
	if timingList == nil {
		APIError(ctx, "获取定时任务列表失败")
	}
	list := []map[string]interface{}{}
	for _, timing := range timingList {
		list = append(list, c.formatTiming(&timing))
	}
	mpurl := "/timing/list"
	if len(querys) > 0 {
		mpurl = "/timing/list?" + strings.Join(querys, "&")
	}
	ctx.HTML(StatusOK, "timing/list", gin.H{
		"Subtitle":   "定时任务列表",
		"List":       list,
		"Total":      total,
		"GroupList":  c.groupService.GetUsageGroup(),
		"AgentList":  c.agentService.GetUsageAgent(),
		"StatusList": constants.TIMING_STATUS,
		"GroupID":    groupID,
		"AgentID":    agentID,
		"Name":       name,
		"Status":     status,
		"Pagination": PagerHtml(total, page, mpurl),
	})
}

func (c *TimingController) Show(ctx *gin.Context) {
	id := DefaultInt64(ctx, "id", 0)
	if id == 0 {
		JumpError(ctx)
		return
	}
	timing := c.timingService.GetTimingByID(id)
	if timing == nil {
		JumpError(ctx)
		return
	}
	ctx.HTML(StatusOK, "timing/show", gin.H{
		"Subtitle": "查看定时任务",
		"Timing":   c.formatTiming(timing),
	})
}

func (c *TimingController) Monitor(ctx *gin.Context) {
	id := DefaultInt(ctx, "id", 0)
	if id == 0 {
		JumpError(ctx)
		return
	}
	cpus := []string{}
	memorys := []string{}
	times := []string{}
	moniters := c.moniterService.GetTimingMonitor(id, 100)
	for _, moniter := range moniters {
		cpus = append(cpus, FormatFloat(moniter.CPU))
		memorys = append(memorys, FormatFloat(moniter.Memory))
		times = append(times, FormatTime(moniter.CreatedAt))
	}
	ctx.HTML(StatusOK, "monitor/list", gin.H{
		"Subtitle": "定时任务监控信息",
		"BackUrl":  "/timing/list",
		"CPU":      cpus,
		"Memory":   memorys,
		"Time":     times,
	})
}

func (c *TimingController) Archive(ctx *gin.Context) {
	id := DefaultInt(ctx, "id", 0)
	page := DefaultInt(ctx, "page", 1)
	where := map[string]interface{}{
		"type":       constants.TYPE_TIMING,
		"related_id": id,
	}
	if id == 0 {
		JumpError(ctx)
		return
	}
	archiveList, total := c.archiveService.GetArchivePageList(where, page, PageSize)
	if archiveList == nil {
		APIError(ctx, "获取归档列表失败")
	}
	list := []map[string]interface{}{}
	for _, archive := range archiveList {
		list = append(list, formatArchive(&archive))
	}
	mpurl := fmt.Sprintf("/timing/archive?id=%d", id)
	ctx.HTML(StatusOK, "archive/list", gin.H{
		"Subtitle":   "定时任务归档列表",
		"List":       list,
		"Total":      total,
		"Pagination": PagerHtml(total, page, mpurl),
	})
}

func (c *TimingController) OutLog(ctx *gin.Context) {
	id := DefaultInt64(ctx, "id", 0)
	lines := DefaultInt64(ctx, "lines", 10)
	if id == 0 {
		JumpError(ctx)
		return
	}
	timing := c.timingService.GetTimingByID(id)
	if timing == nil {
		JumpError(ctx)
		return
	}
	agent := c.agentService.GetAgentByID(timing.AgentID)
	if agent == nil {
		JumpError(ctx)
		return
	}
	content, err := client.GetAgentLog(agent, timing.StdOut, lines)
	if err != nil {
		JumpError(ctx)
		return
	}
	ctx.HTML(StatusOK, "log/list", gin.H{
		"Subtitle": "定时任务日志查看",
		"Path":     "/timing/out_log",
		"BackUrl":  "/timing/list",
		"ID":       id,
		"Lines":    lines,
		"Content":  content,
	})
}

func (c *TimingController) ErrLog(ctx *gin.Context) {
	id := DefaultInt64(ctx, "id", 0)
	lines := DefaultInt64(ctx, "lines", 10)
	if id == 0 {
		JumpError(ctx)
		return
	}
	timing := c.timingService.GetTimingByID(id)
	if timing == nil {
		JumpError(ctx)
		return
	}
	agent := c.agentService.GetAgentByID(timing.AgentID)
	if agent == nil {
		JumpError(ctx)
		return
	}
	content, err := client.GetAgentLog(agent, timing.StdErr, lines)
	if err != nil {
		JumpError(ctx)
		return
	}
	ctx.HTML(StatusOK, "log/list", gin.H{
		"Subtitle": "定时任务错误日志查看",
		"Path":     "/timing/err_log",
		"BackUrl":  "/timing/list",
		"ID":       id,
		"Lines":    lines,
		"Type":     "err_log",
		"Content":  content,
	})
}

func (c *TimingController) Add(ctx *gin.Context) {
	ctx.HTML(StatusOK, "timing/add", gin.H{
		"Subtitle":   "添加定时任务",
		"OutBaseDir": OutDir + "timer/",
		"GroupList":  c.groupService.GetUsageGroup(),
		"AgentList":  c.agentService.GetUsageAgent(),
	})
}

func (c *TimingController) Create(ctx *gin.Context) {
	groupID := FormDefaultInt64(ctx, "group_id", 0)
	name := ctx.PostForm("name")
	agentID := FormDefaultInt64(ctx, "agent_id", 0)
	dir := ctx.PostForm("dir")
	program := ctx.PostForm("program")
	args := ctx.PostForm("args")
	stdOut := ctx.PostForm("std_out")
	stdErr := ctx.PostForm("std_err")
	_time := ctx.PostForm("time")
	timeout := FormDefaultInt(ctx, "timeout", 0)
	isMonitor := ctx.PostForm("is_monitor")
	if !Required(ctx, &name, "名称不能为空") {
		return
	}
	if !Required(ctx, &dir, "执行目录不能为空") {
		return
	}
	if !Required(ctx, &program, "执行程序不能为空") {
		return
	}
	if !Required(ctx, &stdOut, "标准输出路径不能为空") {
		return
	}
	if !Required(ctx, &stdErr, "错误输出路径不能为空") {
		return
	}
	if !Required(ctx, &_time, "运行时间不能为空") {
		return
	}
	if agentID == 0 {
		APIBadRequest(ctx, "运行实例不能为空")
		return
	}
	timing := new(models.Timing)
	timing.GroupID = groupID
	timing.Name = name
	timing.AgentID = agentID
	timing.Dir = dir
	timing.Program = program
	timing.Args = args
	timing.StdOut = stdOut
	timing.StdErr = stdErr
	timing.Time, _ = parseTime(_time)
	timing.Timeout = int64(timeout)
	timing.Status = constants.TIMING_STATUS_STOP
	timing.Creator = GetUserID(ctx)
	if isMonitor != "" {
		timing.IsMonitor = 1
	}
	ok := c.timingService.CreateTiming(timing)
	if !ok {
		APIError(ctx, "创建定时任务失败")
		return
	}
	APIOK(ctx)
}

func (c *TimingController) Edit(ctx *gin.Context) {
	id := DefaultInt64(ctx, "id", 0)
	if id == 0 {
		JumpError(ctx)
		return
	}
	timing := c.timingService.GetTimingByID(id)
	if timing == nil {
		JumpError(ctx)
		return
	}
	ctx.HTML(StatusOK, "timing/edit", gin.H{
		"Subtitle":  "编辑定时任务",
		"Info":      c.formatTiming(timing),
		"GroupList": c.groupService.GetUsageGroup(),
		"AgentList": c.agentService.GetUsageAgent(),
	})
}

func (c *TimingController) Update(ctx *gin.Context) {
	id := FormDefaultInt64(ctx, "id", 0)
	groupID := FormDefaultInt64(ctx, "group_id", 0)
	name := ctx.PostForm("name")
	agentID := FormDefaultInt64(ctx, "agent_id", 0)
	dir := ctx.PostForm("dir")
	program := ctx.PostForm("program")
	args := ctx.PostForm("args")
	stdOut := ctx.PostForm("std_out")
	stdErr := ctx.PostForm("std_err")
	_time := ctx.PostForm("time")
	timeout := FormDefaultInt64(ctx, "timeout", 0)
	isMonitor := ctx.PostForm("is_monitor")
	if id == 0 {
		APIBadRequest(ctx, "ID格式错误")
		return
	}
	if !Required(ctx, &name, "名称不能为空") {
		return
	}
	if !Required(ctx, &dir, "执行目录不能为空") {
		return
	}
	if !Required(ctx, &program, "执行程序不能为空") {
		return
	}
	if !Required(ctx, &stdOut, "标准输出路径不能为空") {
		return
	}
	if !Required(ctx, &stdErr, "错误输出路径不能为空") {
		return
	}
	if !Required(ctx, &_time, "运行时间不能为空") {
		return
	}
	if agentID == 0 {
		APIBadRequest(ctx, "运行实例不能为空")
		return
	}
	timing := c.timingService.GetTimingByID(id)
	if timing == nil {
		APIBadRequest(ctx, "定时任务不存在")
		return
	}
	timing.GroupID = groupID
	timing.Name = name
	timing.AgentID = agentID
	timing.Dir = dir
	timing.Program = program
	timing.Args = args
	timing.StdOut = stdOut
	timing.StdErr = stdErr
	timing.Time, _ = parseTime(_time)
	timing.Timeout = timeout
	timing.Updator = GetUserID(ctx)
	if isMonitor != "" {
		timing.IsMonitor = 1
	}
	ok := c.timingService.UpdateTiming(timing)
	if !ok {
		APIError(ctx, "更新定时任务失败")
		return
	}
	APIOK(ctx)
}

func (c *TimingController) Copy(ctx *gin.Context) {
	id := DefaultInt64(ctx, "id", 0)
	if id == 0 {
		APIBadRequest(ctx, "ID格式错误")
		return
	}
	timing := c.timingService.GetTimingByID(id)
	if timing == nil {
		APIError(ctx, "定时任务不存在")
		return
	}
	_timing := new(models.Timing)
	_timing.GroupID = timing.GroupID
	_timing.Name = timing.Name + "_copy"
	_timing.AgentID = timing.AgentID
	_timing.Dir = timing.Dir
	_timing.Program = timing.Program
	_timing.Args = timing.Args
	_timing.StdOut = timing.StdOut
	_timing.StdErr = timing.StdErr
	_timing.Time = timing.Time
	_timing.Timeout = timing.Timeout
	_timing.IsMonitor = timing.IsMonitor
	_timing.Status = constants.TIMING_STATUS_STOP
	_timing.Creator = GetUserID(ctx)
	ok := c.timingService.CreateTiming(_timing)
	if !ok {
		APIError(ctx, "复制定时任务失败")
		return
	}
	APIOK(ctx)
}

func (c *TimingController) Delete(ctx *gin.Context) {
	id := DefaultInt64(ctx, "id", 0)
	if id == 0 {
		APIBadRequest(ctx, "ID格式错误")
		return
	}
	timing := c.timingService.GetTimingByID(id)
	if timing == nil {
		APIError(ctx, "定时任务不存在")
		return
	}
	if timing.Status == 1 {
		APIError(ctx, "定时任务正在运行不能删除")
		return
	}
	timing.Status = -1
	timing.Updator = GetUserID(ctx)
	ok := c.timingService.UpdateTiming(timing)
	if !ok {
		APIError(ctx, "删除定时任务失败")
		return
	}
	APIOK(ctx)
}

func (c *TimingController) Start(ctx *gin.Context) {
	id := DefaultInt64(ctx, "id", 0)
	if id == 0 {
		APIBadRequest(ctx, "ID格式错误")
		return
	}
	timing := c.timingService.GetTimingByID(id)
	if timing == nil {
		APIError(ctx, "定时任务不存在")
		return
	}
	if timing.Status == constants.TIMING_STATUS_RUNNING {
		APIError(ctx, "定时任务已经启动")
		return
	}
	agent := c.agentService.GetAgentByID(timing.AgentID)
	if agent == nil {
		APIError(ctx, "定时任务对应实例获取异常")
		return
	}
	_timing, err := client.GetAgentTiming(agent, id)
	if err != nil {
		APIError(ctx, fmt.Sprintf("获取定时任务情况异常:%s", err.Error()))
		return
	}
	if _timing == nil {
		err = client.AddAgentTiming(agent, timing)
		if err != nil {
			APIError(ctx, fmt.Sprintf("添加计划任务异常:%s", err.Error()))
			return
		}
		timing.Status = constants.TIMING_STATUS_RUNNING
		c.timingService.UpdateTiming(timing)
		APIOK(ctx)
		return
	}
	timing.Status = constants.TIMING_STATUS_RUNNING
	timing.Updator = GetUserID(ctx)
	c.timingService.UpdateTiming(timing)
	APIOK(ctx)
}

func (c *TimingController) ReStart(ctx *gin.Context) {
	id := DefaultInt64(ctx, "id", 0)
	if id == 0 {
		APIBadRequest(ctx, "ID格式错误")
		return
	}
	timing := c.timingService.GetTimingByID(id)
	if timing == nil {
		APIError(ctx, "定时任务不存在")
		return
	}
	agent := c.agentService.GetAgentByID(timing.AgentID)
	if agent == nil {
		APIError(ctx, "定时任务对应实例获取异常")
		return
	}
	_timing, err := client.GetAgentTiming(agent, id)
	if err != nil {
		APIError(ctx, fmt.Sprintf("获取定时任务情况异常:%s", err.Error()))
		return
	}
	if _timing == nil {
		err = client.AddAgentTiming(agent, timing)
		if err != nil {
			APIError(ctx, fmt.Sprintf("重启定时任务异常:%s", err.Error()))
			return
		}
		APIOK(ctx)
	}
	err = client.UpdateAgentTiming(agent, timing)
	if err != nil {
		APIError(ctx, fmt.Sprintf("重启定时任务异常:%s", err.Error()))
		return
	}
	APIOK(ctx)
}

func (c *TimingController) Pause(ctx *gin.Context) {
	id := DefaultInt64(ctx, "id", 0)
	if id == 0 {
		APIBadRequest(ctx, "ID格式错误")
		return
	}
	timing := c.timingService.GetTimingByID(int64(id))
	if timing == nil {
		APIError(ctx, "定时任务不存在")
		return
	}
	agent := c.agentService.GetAgentByID(timing.AgentID)
	if agent == nil {
		APIError(ctx, "定时任务对应实例获取异常")
		return
	}
	_timing, err := client.GetAgentTiming(agent, id)
	if err != nil {
		APIError(ctx, fmt.Sprintf("获取定时任务情况异常:%s", err.Error()))
		return
	}
	if _timing == nil {
		APIOK(ctx)
		return
	}
	err = client.RemoveAgentTiming(agent, int64(id))
	if err != nil {
		APIError(ctx, fmt.Sprintf("停止定时任务异常:%s", err.Error()))
		return
	}
	timing.Status = constants.TIMING_STATUS_PAUSE
	timing.Updator = GetUserID(ctx)
	c.timingService.UpdateTiming(timing)
	APIOK(ctx)
}
