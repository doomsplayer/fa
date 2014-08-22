package admin

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
	"fmt"
	"github.com/astaxie/beego"
	"math"
	"runtime"
	"strings"
	"time"
)

func init() {
	fmt.Print("")

}

type LoginController struct {
	lib.BaseController
}

func (this *LoginController) Prepare() {
	this.BaseController.Prepare()
	beego.ReadFromRequest(&this.Controller)
}

// @router /login [get]
func (this *LoginController) Login() {
	this.TplNames = "adminLayout/login.html"
	this.Render()
}

// @router /login [post]
func (this *LoginController) LoginPost() {
	u := this.GetString("username")
	p := this.GetString("password")
	if u == "123" && p == "123" {
		this.SetSession(`admin`, true)
		this.Redirect(this.UrlFor(`AdminController.Index`), 302)
		return
	} else {
		f := beego.NewFlash()
		f.Error("用户名或密码不正确")
		f.Store(&this.Controller)
		this.Redirect(this.UrlFor(`.Login`), 302)
		return
	}

}

type AdminController struct {
	lib.AuthController
}

func (this *AdminController) Prepare() {
	this.AuthController.Prepare()
	if this.GetSession(`admin`) == nil {
		this.Redirect(this.UrlFor(`LoginController.Login`), 302)
	}
	this.Data["title"] = "后台"
	this.Layout = "adminLayout/layout.html"
}

// @router / [get]
func (this *AdminController) Index() {

	this.TplNames = "adminLayout/index.html"
	this.Render()
}

// @router /learn [get]
func (this *AdminController) Learn() {
	t := &models.TutorialSlice{}
	t.All(100000, 0, ``)
	this.Data["learns"] = t
	this.TplNames = "adminLayout/learn.html"
	this.Render()
}

// @router /sysstatus [get]
func (this *AdminController) Sysstatus() {
	var startTime = time.Now()

	var sysStatus struct {
		Overall struct {
			OS        string `json:"操作系统"`
			Arch      string `json:"cpu架构"`
			CPUNum    int    `json:"cpu数量"`
			GoVersion string `json:"go版本"`

			Uptime       string `json:"已运行时间"`
			NumGoroutine int    `json:"goroutine数量"`
			OtherSys     string // other system allocations
		} `json:"总览"`

		MemStates struct {
			// General statistics.
			MemAllocated string `json:"已使用内存"`    // bytes allocated and still in use
			MemTotal     string `json:"已分配内存"`    // bytes allocated (even if freed)
			MemSys       string `json:"从系统中获得内存"` // bytes obtained from system (sum of XxxSys below)
			Lookups      uint64 `json:"指针查询数量"`   // number of pointer lookups
			MemMallocs   uint64 `json:"malloc数量"` // number of mallocs
			MemFrees     uint64 `json:"free数量"`   // number of frees
		} `json:"内存信息"`
		HeapStates struct {
			// Main allocation heap statistics.
			HeapAlloc    string `json:"堆内存分配"` // bytes allocated and still in use
			HeapSys      string `json:"堆内存获得"` // bytes obtained from system
			HeapIdle     string `json:"堆内存空闲"` // bytes in idle spans
			HeapInuse    string `json:"堆内存使用"` // bytes in non-idle span
			HeapReleased string `json:"堆释放"`   // bytes released to the OS
			HeapObjects  uint64 `json:"堆对象数量"` // total number of allocated objects
		} `json:"堆信息"`
		StackStates struct {
			// Low-level fixed-size structure allocator statistics.
			//	Inuse is bytes used now.
			//	Sys is bytes obtained from system.
			StackInuse  string `json:"栈内存使用"` // bootstrap stacks
			StackSys    string `json:"栈内存分配"`
			MSpanInuse  string // mspan structures
			MSpanSys    string
			MCacheInuse string // mcache structures
			MCacheSys   string
			BuckHashSys string // profiling bucket hash table

		} `json:"栈信息"`

		GCStates struct {
			// Garbage collector statistics.
			GCSys        string `json:"GC元数据大小"`  // GC metadata
			NextGC       string `json:"下一次GC大小"`  // next run in HeapAlloc time (bytes)
			LastGC       string `json:"上一次GC距现在"` // last run in absolute time (ns)
			PauseTotalNs string `json:"GC总暂停时间"`
			PauseNs      string `json:"上一次GC暂停时间"` // circular buffer of recent GC pause times, most recent at [(NumGC+255)%256]
			NumGC        uint32 `json:"GC次数"`
		} `json:"GC信息"`
	}

	sysStatus.Overall.GoVersion = runtime.Version()
	sysStatus.Overall.CPUNum = runtime.NumCPU()
	sysStatus.Overall.OS = runtime.GOOS
	sysStatus.Overall.Arch = runtime.GOARCH
	sysStatus.Overall.Uptime = TimeSincePro(startTime)

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	sysStatus.Overall.NumGoroutine = runtime.NumGoroutine()

	sysStatus.MemStates.MemAllocated = FileSize(int64(m.Alloc))
	sysStatus.MemStates.MemTotal = FileSize(int64(m.TotalAlloc))
	sysStatus.MemStates.MemSys = FileSize(int64(m.Sys))
	sysStatus.MemStates.Lookups = m.Lookups
	sysStatus.MemStates.MemMallocs = m.Mallocs
	sysStatus.MemStates.MemFrees = m.Frees

	sysStatus.HeapStates.HeapAlloc = FileSize(int64(m.HeapAlloc))
	sysStatus.HeapStates.HeapSys = FileSize(int64(m.HeapSys))
	sysStatus.HeapStates.HeapIdle = FileSize(int64(m.HeapIdle))
	sysStatus.HeapStates.HeapInuse = FileSize(int64(m.HeapInuse))
	sysStatus.HeapStates.HeapReleased = FileSize(int64(m.HeapReleased))
	sysStatus.HeapStates.HeapObjects = m.HeapObjects

	sysStatus.StackStates.StackInuse = FileSize(int64(m.StackInuse))
	sysStatus.StackStates.StackSys = FileSize(int64(m.StackSys))
	sysStatus.StackStates.MSpanInuse = FileSize(int64(m.MSpanInuse))
	sysStatus.StackStates.MSpanSys = FileSize(int64(m.MSpanSys))
	sysStatus.StackStates.MCacheInuse = FileSize(int64(m.MCacheInuse))
	sysStatus.StackStates.MCacheSys = FileSize(int64(m.MCacheSys))
	sysStatus.StackStates.BuckHashSys = FileSize(int64(m.BuckHashSys))
	sysStatus.GCStates.GCSys = FileSize(int64(m.GCSys))
	sysStatus.Overall.OtherSys = FileSize(int64(m.OtherSys))

	sysStatus.GCStates.NextGC = FileSize(int64(m.NextGC))
	sysStatus.GCStates.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	sysStatus.GCStates.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)
	sysStatus.GCStates.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	sysStatus.GCStates.NumGC = m.NumGC
	this.Data[`status`] = sysStatus
	this.TplNames = "adminLayout/sysStatus.html"
	this.Render()
	return
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+"%s", val, suffix)
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1024, sizes)
}

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

// TimeSincePro calculates the time interval and generate full user-friendly string.
func TimeSincePro(then time.Time) string {
	now := time.Now()
	diff := now.Unix() - then.Unix()

	if then.After(now) {
		return "future"
	}

	var timeStr, diffStr string
	for {
		if diff == 0 {
			break
		}

		diff, diffStr = computeTimeDiff(diff)
		timeStr += ", " + diffStr
	}
	return strings.TrimPrefix(timeStr, ", ")
}

func computeTimeDiff(diff int64) (int64, string) {
	diffStr := ""
	switch {
	case diff <= 0:
		diff = 0
		diffStr = "now"
	case diff < 2:
		diff = 0
		diffStr = "1 second"
	case diff < 1*Minute:
		diffStr = fmt.Sprintf("%d seconds", diff)
		diff = 0

	case diff < 2*Minute:
		diff -= 1 * Minute
		diffStr = "1 minute"
	case diff < 1*Hour:
		diffStr = fmt.Sprintf("%d minutes", diff/Minute)
		diff -= diff / Minute * Minute

	case diff < 2*Hour:
		diff -= 1 * Hour
		diffStr = "1 hour"
	case diff < 1*Day:
		diffStr = fmt.Sprintf("%d hours", diff/Hour)
		diff -= diff / Hour * Hour

	case diff < 2*Day:
		diff -= 1 * Day
		diffStr = "1 day"
	case diff < 1*Week:
		diffStr = fmt.Sprintf("%d days", diff/Day)
		diff -= diff / Day * Day

	case diff < 2*Week:
		diff -= 1 * Week
		diffStr = "1 week"
	case diff < 1*Month:
		diffStr = fmt.Sprintf("%d weeks", diff/Week)
		diff -= diff / Week * Week

	case diff < 2*Month:
		diff -= 1 * Month
		diffStr = "1 month"
	case diff < 1*Year:
		diffStr = fmt.Sprintf("%d months", diff/Month)
		diff -= diff / Month * Month

	case diff < 2*Year:
		diff -= 1 * Year
		diffStr = "1 year"
	default:
		diffStr = fmt.Sprintf("%d years", diff/Year)
		diff = 0
	}
	return diff, diffStr
}
