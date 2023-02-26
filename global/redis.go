package global

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

var GoRedisClient *redis.Client

const (
	Following       = "following:"       // 关注集合Key
	Followers       = "followers:"       // 粉丝集合key
	FollowingFeeds  = "following_feeds:" //我关注的好友的FeedsKey
	CommonKey       = "COMMON_FOLLOWING"
	PageSize        = 10
	DateFormat1     = "2006-01-02"
	DateFormat2     = "2006-01-02 15:04:05"
	SignFormat      = "200601"
	AccountPoints   = "account:points"    //用户的积分Key
	AccountLocation = "account:location"  //用户地理位置Key
	Product         = "product:"          //产品的Key
	ProductComment  = "product:comments:" //商品评论Key
)

func InitGoRedisClient() {
	GoRedisClient = redis.NewClient(&redis.Options{
		Addr:     HOST + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//r:=GoRedisClient.Ping(context.Background())
	//r.Err()
}
