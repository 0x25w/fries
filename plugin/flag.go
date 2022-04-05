package plugin

import "flag"

func Flag(f *Finger) {
	Banner()
	flag.StringVar(&f.Url, "u", "", "-u http://127.0.0.1/")
	flag.StringVar(&f.List, "l", "", "exec command (ssh)")
	flag.StringVar(&f.Fofa, "f", "", "sshkey file (id_rsa)")
	flag.StringVar(&f.Zoomeye, "domain", "", "smb domain")
	flag.Int64Var(&f.Timeout, "time", 3, "Set timeout")
	flag.StringVar(&f.Output, "o", "", "-o ./path.txt")
	flag.IntVar(&Threads, "t", 600, "Thread nums")
	flag.Int64Var(&f.Timeout, "wt", 5, "Set web timeout")
	flag.Parse()
}
