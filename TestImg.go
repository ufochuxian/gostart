package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func main() {
	client, err := sdk.NewClientWithAccessKey("cn-shanghai", "LTAI4GKYXhzjRK5H9nhK48PA", "IQoeMGbYpczZbSphlmIo628ckV7MDe")
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "imageenhan.cn-shanghai.aliyuncs.com"
	request.Version = "2019-09-30"
	request.ApiName = "ImageBlindPicWatermark"
	request.QueryParams["RegionId"] = "cn-shanghai"
	request.QueryParams["FunctionType"] = "encode_pic"
	request.QueryParams["QualityFactor"] = "100"
	request.QueryParams["OutputFileType"] = "jpg"
	request.QueryParams["OriginImageURL"] = "https://tgm-test.oss-cn-shanghai.aliyuncs.com/p8p9-3.png?Expires=1611293104&OSSAccessKeyId=TMP.3Kj6Sy6Wy5A8DYdeNvqFPy5K3wBWkHnBwWT3tU5zoPtLsCjz4pcjRnNniCsKpFAUmcrHmxhkuUsyg2nij1iqh6V39YbPps&Signature=rLTlabQqK6pR7caf3GI7C5QztMY%3D&versionId=CAEQHhiBgMCc76yhuRciIGYxMzY5NzBiNWMxNTRmMGJhMWFhZTU4ZmQwYzk4ZmEw&response-content-type=application%2Foctet-stream"
	request.QueryParams["LogoURL"] = "https://tgm-test.oss-cn-shanghai.aliyuncs.com/tgm_code.png?Expires=1611293172&OSSAccessKeyId=TMP.3Kj6Sy6Wy5A8DYdeNvqFPy5K3wBWkHnBwWT3tU5zoPtLsCjz4pcjRnNniCsKpFAUmcrHmxhkuUsyg2nij1iqh6V39YbPps&Signature=4UdYLqbm3oLLD4vLk1QvlUsSGJg%3D&versionId=CAEQHhiBgIDNnK2huRciIGZjMDQwNGYyOWNjNjRjMGFiOTZkOGZmOGMxNTg3MTEw&response-content-type=application%2Foctet-stream"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpContentString())
}