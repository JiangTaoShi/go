// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"encoding/json"
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ice20201109 "github.com/alibabacloud-go/ice-20201109/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *ice20201109.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("ice.cn-shanghai.aliyuncs.com")
	_result = &ice20201109.Client{}
	_result, _err = ice20201109.NewClient(config)
	return _result, _err
}

type OutputMediaConfig struct {
	MediaURL        string
	StorageLocation string
	FileName        string
}

type ClipsParam struct {
	StartName string `json:"start_name"`
	Image01   string `json:"image01"`
	Image02   string `json:"image02"`
	Image03   string `json:"image03"`
}

type UserData struct {
	NotifyAddress string
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient(tea.String(""), tea.String(""))
	if _err != nil {
		return _err
	}

	outputMediaConfig := OutputMediaConfig{
		MediaURL:        "https://tunshu-shanghai.oss-cn-shanghai.aliyuncs.com/tunshu/MP4/331322332.mp4",
		StorageLocation: "tunshu-shanghai.oss-cn-shanghai.aliyuncs.com",
		FileName:        "331322332.mp4",
	}
	outputMediaConfigJson, _ := json.Marshal(outputMediaConfig)

	clipsParam := ClipsParam{
		StartName: "张三35",
		Image01:   "https://tunshu-shanghai.oss-cn-shanghai.aliyuncs.com/tunshu/Picture/433312596901277697.jpg",
		Image02:   "https://tunshu-shanghai.oss-cn-shanghai.aliyuncs.com/tunshu/Picture/433430246558052353.jpg",
		Image03:   "https://tunshu-shanghai.oss-cn-shanghai.aliyuncs.com/tunshu/Picture/433312823293030401.jpg",
	}
	clipsParamJson, _ := json.Marshal(clipsParam)

	userData := UserData{
		NotifyAddress: "https://api.t.dev.pay.fun/taikang/customer/mv_notify",
	}
	userDataJson, _ := json.Marshal(userData)

	submitMediaProducingJobRequest := &ice20201109.SubmitMediaProducingJobRequest{
		TemplateId:        tea.String("fc11dabd5f4a463c9401b28d5e678fd5"),
		ClipsParam:        tea.String(string(clipsParamJson)),
		OutputMediaConfig: tea.String(string(outputMediaConfigJson)),
		Source:            tea.String("OpenAPI"),
		UserData:          tea.String(string(userDataJson)),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.SubmitMediaProducingJobWithOptions(submitMediaProducingJobRequest, runtime)
	if _err != nil {
		return _err
	}

	console.Log(util.ToJSONString(tea.ToMap(resp)))
	return _err
}

func main() {
	// err := _main(tea.StringSlice(os.Args[1:]))
	// if err != nil {
	// 	panic(err)
	// }
	//获取单个合成任务

	client, _err := CreateClient(tea.String(""), tea.String(""))
	if _err != nil {
		fmt.Println(_err)
	}
	jobId := "8c1fda3dee714bd980012ff8de640174"
	getMediaProducingJobRequest := &ice20201109.GetMediaProducingJobRequest{
		JobId: &jobId,
	}
	getMediaProducingJobResult, err := client.GetMediaProducingJob(getMediaProducingJobRequest)
	fmt.Println(getMediaProducingJobResult, err)

}
