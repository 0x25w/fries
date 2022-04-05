package plugin

type Finger struct {
	Url     string
	List    string
	Fofa    string
	Zoomeye string
	Urllist []string
	Timeout int64
	Output  string
}

// 默认线程 5
var Threads = 5
