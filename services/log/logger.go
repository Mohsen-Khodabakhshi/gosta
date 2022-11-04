package log

import (
	"fmt"
	"os"
)

type (
	HttpLogs struct {
		Status int
		Method string
		Url    string
		Date   string
		Time   string
	}
)

func LoadHttp(method, url, date, time string, status int) *HttpLogs {
	return &HttpLogs{Status: status, Method: method, Url: url, Time: time, Date: date}
}

func HttpOnConsole(output *HttpLogs) {
	fmt.Printf("%v: uri: %v, status: %v, method: %v\n", output.Time, output.Url, output.Status, output.Method)
}

func HttpToFile(output *HttpLogs) {
	file, _ := os.OpenFile(fmt.Sprintf("./logs/%s", output.Date), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()

	_, _ = file.WriteString(fmt.Sprintf("%v %v: uri: %v, status: %v, method: %v\n", output.Date, output.Time, output.Url, output.Status, output.Method))

}
