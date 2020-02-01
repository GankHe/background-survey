package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gopay"
	"go-background-survey/alipayMgr"
	"net/http"
)

func GetTest(c *gin.Context){
	var(
		body gopay.BodyMap
		payUrl string
		err error
	)
	//请求参数
	body.Set("subject", "网站测试支付")
	body.Set("out_trade_no", "GZ201901301040355706100469")
	body.Set("total_amount", "88.88")
	body.Set("product_code", "FAST_INSTANT_TRADE_PAY")
	payUrl,err=alipayMgr.G_jobMgr.TradePagePay(body)
	if err !=nil{
		fmt.Println(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"status_code":http.StatusMovedPermanently,
		"data":payUrl,
	})
}