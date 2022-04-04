package plugin

func DoParse() {
	parselocal()
}

func parselocal() {
	Scanner.Flags().StringVarP(&localfile, "local", "l", "", "从本地文件读取资产，进行指纹识别，支持无协议，列如：192.168.1.1:9090 | http://192.168.1.1:9090")

}
