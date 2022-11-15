package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

var token string = ""
var key string = ""
var baseUrl string = "http://test..com:58080/open/order/save"
var jobPrefix string = "tdx"

func main() {
	timeSpan64 := TimeToUnix(time.Now())
	timeSpanStr := strconv.FormatInt(timeSpan64, 10)
	base64Str := fmt.Sprintf("%s%s", baseUrl, timeSpanStr)
	encoded := base64.StdEncoding.EncodeToString([]byte(base64Str))
	encoded = fmt.Sprintf("%s%s", encoded, key)
	md5Encoded := MD5(encoded)

	request := YiYouMeiOrderRequest{}
	request.OrderNo = jobPrefix + "433308574647582939"
	request.Version = 1
	request.Priority = "NORMAL"
	request.PayDate = "2022-11-10 14:16:38"
	request.CreateDate = "2022-11-10 14:16:38"
	request.StoreId = "DX-56685"

	jobList := []YiYouMeiOrderJob{}
	pdfList := []YiYouMeiOrderJobPdf{}

	pdfList = append(pdfList, YiYouMeiOrderJobPdf{
		CheckNum:  "433308574647582936",
		GroupType: "TEXT_PRINT",
		PdfUrl:    "https://tunshu.oss-cn-zhangjiakou.aliyuncs.com/tunshu/PDF/433308574647582936.pdf",
	})
	pdfList = append(pdfList, YiYouMeiOrderJobPdf{
		CheckNum:  "",
		GroupType: "COVER_PRINT",
		PdfUrl:    "",
	})

	jobList = append(jobList, YiYouMeiOrderJob{
		JobId:   jobPrefix + "4333089",
		SkuCode: "TAILI",
		SkuName: "台历",
		QTY:     1,
		JobType: 0,
		Comment: "测试",
		PdfList: pdfList,
		CustomerAddress: YiYouMeiOrderJobAddress{
			Country:  "中国",
			Address:  "测试地址",
			City:     "济南",
			Name:     "测试店",
			Phone:    "18301431893",
			Province: "山东省",
			Region:   "历下区",
		},
		StoreAddress: YiYouMeiOrderJobAddress{
			Country:  "中国",
			Address:  "测试地址",
			City:     "济南",
			Name:     "测试店",
			Phone:    "18301431893",
			Province: "山东省",
			Region:   "历下区",
		},
	})

	request.JobList = jobList

	requestJson, _ := json.Marshal(request)
	requestJsonStr := string(requestJson)
	fmt.Println(requestJsonStr)

	response := YiYouMeiOrderResponse{}
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("token", token).
		SetHeader("sign", md5Encoded).
		SetHeader("timestamp", timeSpanStr).
		SetBody(requestJsonStr).
		SetResult(&response).
		Post("http://test.kencoo.com:58080/open/order/save")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func TimeToUnix(e time.Time) int64 {
	timeUnix, _ := time.Parse("2006-01-02 15:04:05", e.Format("2006-01-02 15:04:05"))
	return timeUnix.UnixNano() / 1e6
}

type YiYouMeiOrderRequest struct {
	OrderNo    string             `json:"orderNo"`
	Version    int                `json:"version"`
	Priority   string             `json:"priority"`
	PayDate    string             `json:"payDate"`
	CreateDate string             `json:"createDate"`
	StoreId    string             `json:"storeId"`
	JobList    []YiYouMeiOrderJob `json:"jobList"`
}
type YiYouMeiOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type YiYouMeiOrderJob struct {
	JobId           string                  `json:"jobId"`           //JobId
	SkuCode         string                  `json:"skuCode"`         //产品 sku
	SkuName         string                  `json:"skuName"`         //产品名称
	QTY             int                     `json:"qty"`             //数量
	JobType         int                     `json:"jobType"`         //0 定制产品（必须传递任务pdf）
	Comment         string                  `json:"comment"`         //
	PdfList         []YiYouMeiOrderJobPdf   `json:"pdfList"`         //
	CustomerAddress YiYouMeiOrderJobAddress `json:"customerAddress"` //用户信息
	StoreAddress    YiYouMeiOrderJobAddress `json:"storeAddress"`    //加盟店地址
}

type YiYouMeiOrderJobPdf struct {
	CheckNum  string `json:"checkNum"`
	GroupType string `json:"groupType"` //TEXT_PRINT
	PdfUrl    string `json:"pdfUrl"`
}

type YiYouMeiOrderJobAddress struct {
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Province string `json:"province"`
	Region   string `json:"region"`
}
