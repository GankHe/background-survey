package alipayMgr

import (
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/alipay"
)

type AlipayMgr struct {
	privateKey string
	client *alipay.Client
}

var(
	G_jobMgr *AlipayMgr
)

func InitAlipayMgr(isProd bool)(err error){
	var(
		privateKey string
		client *alipay.Client
	)
	privateKey="MIIEowIBAAKCAQEAjGzNIgk66PMTJzR7wd11thb16FtEFn6W8+kWHDAabgwih/5V6UkE6B9L59opT5tt+S7kspOwwVl8l0z5in3iYgEA/3mfscqh2xlH+EblGP/JTToJiVrhXYGJJsLiL5ynHChdsk00crqhpn1rUr2Tf7znh7oNHiMiph2LyXt+q7ThglaUvzCK3mWV7u/iZ2627S8LOgrYbA7v8VULKNMyvTiNSZyVl7pbigogBOyuowFPO38eMmm1O+PJTPVohbF9ms15AlOpz2X8CUwtovK/lKthZ6xk5ZWc7TAkkI/UVxXhfyfR9dDvIH0704F8nftIbyLP1E50kUFWcJqi3tQyaQIDAQABAoIBAEvSt9Tn/Popvi7OJQDiDpw6u58yIIqw4dtkAT9JLgw72y1pRIRJnC+mLntLjkDg4TnKe0x9eNR9+EtOEfQY0vMuLZbE2ljK3CiM9xJ1XWDpRI3Bh8K14rh1qFPYqP8ibA1FMoqDK43DfmOClz+/A3TW0IG2ET0zpZLio0cwuo4FuHxqhl/E17VsFNWvFL0c9rW+qQA6kZ6cLnb7zwCmp4jzX7aJBryQ8RZ7yFfVPw4cw+XQAfsOy+3f3p8rDamli0anzfKLnck+GvyCmymO91vfhHghI+kfNEkrV7kUdGXySoJ91Pq1R1wJpus2e9IpjJTseIbrnIStW47AroPrUTECgYEA0Scu3JZetltAJD2RwYPZH0bGTkLXUozNtGB2z6i+ExOQ5/XT3OqkH8GedQ+YZRimeFZTxe8EV5crdiZkFR+bJpIoeXNPHsRtOgc4e5Yv6U7Yxq8Tg5EKIBHZpTo9RLIdHXOJlbVFYM990C5TQpvLkV712pDLaSVwnZ4EpX/8rb0CgYEAq+DCDq3zaIROJk5S1SqOEBEmAscCbp6pRpjUwW7H0uq8OPra0dNSioZuJoIl1FXFIgMODCEKujdCZRz2f6VZhKi1v6BtNriPNTPb6+TifNYY3IMnLbGOLQnBfHZN1WqST3de1nu/rbQmoXR0ZiXE6sNlzUi/nxXjI4cUb5/k1B0CgYA6czPSa6d+WNP9DMPOC9XCutmyh1V7eWQHU6oAcIzl8eLeIEIcC+rgARESb9BL+G6VVLLGrvxMqi94LlRr3tBSICh1gHcNByHrRJdRWnB7SBCmvci7TO5CdhqX5UW4C/q8/0vk7aq82l+zrS6Rf86NCre0ZCCketUVB1MPGUx5FQKBgQCPln7IAK5eXFllSv2MM7bm2Uyl+Vegzk+gOMFUr5WkMJH6ECeKLz6/lzqiVQwBbX34whCUNW5ezxdanW2YiaxVVRQw+HcExf5tddt+IILNw5aVccPjOngKTCImNjcj2ZpNZO4HxO4G0X+MVlt/koIFliOP8fyTjJdq8Mgz0KI8SQKBgHpcBJ5h0j5qZVT8NoigkdPSQjCN+JheoO4y8UOz9voK2pZhc93WDkZckB/T+3Q4iSiRDCOPxrieFrDkeA/ielpH5goGgihnPOxpseuAkUsMqIa94mocfqdpreEO8DYte5r+IQz1KBjbiznRHGhvKZargQGkCNiL+R83z4031LN2"
	client=alipay.NewClient("2016101700711236", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl("https://www.gopay.ink")
	G_jobMgr=&AlipayMgr{
		privateKey:privateKey,
		client:client,
	}
	return
}


func (a AlipayMgr)TradePagePay(body gopay.BodyMap)(payUrl string,err error){
	payUrl, err = a.client.TradePagePay(body)
	return
}