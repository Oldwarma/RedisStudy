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

//GetDate 获取日期
func GetDate(format string) string {
	return time.Now().Format(format)
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetLastDateOfMonth 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func GetDaysOfMonth() []string {
	days := make([]string, 0)

	year := time.Now().Year()

	for month := 1; month <= 12; month++ {
		for day := 1; day <= 31; day++ {
			//如果是2月
			if month == 2 {
				if isLeapYear(year) && day == 30 { //闰年2月29天
					break
				} else if !isLeapYear(year) && day == 29 { //平年2月28天
					break
				} else {
					days = append(days, fmt.Sprintf("%d-%02d-%02d", year, month, day))
				}
			} else if month == 4 || month == 6 || month == 9 || month == 11 { //小月踢出来
				if day == 31 {
					break
				}
				days = append(days, fmt.Sprintf("%d-%02d-%02d", year, month, day))
			} else {
				days = append(days, fmt.Sprintf("%d-%02d-%02d", year, month, day))
			}
		}
	}

	fmt.Println(days)
	return days
}

//判断是否为闰年
func isLeapYear(year int) bool {
	//y == 2000, 2004
	//判断是否为闰年
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}

func GetWeek(datetime string) (y, w int) {
	timeLayout := "20060102"
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	return tmp.ISOWeek()
}
