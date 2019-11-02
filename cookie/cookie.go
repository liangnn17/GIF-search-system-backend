package cookie

import(
	gocache "github.com/patrickmn/go-cache"
	// "fmt"
	"time"
	"backend/utils"
	"net/http"
	"crypto/sha1"
	// "crypto/md5"
	"encoding/binary"
	"bytes"
	// "encoding/json"
)

func CookieCacheInit() *gocache.Cache{
	c:=gocache.New(utils.COOKIE_EXPIRE*time.Second, 600*time.Second)
	return c
}

func CookieSet(user string, c *gocache.Cache){
	c.Set(user, ShaConvert(user), gocache.DefaultExpiration)
}

func RootCookieSet(user string, c *gocache.Cache){
	c.Set(user, "", gocache.DefaultExpiration)
}

func CookieCheck(req *http.Request, c *gocache.Cache)bool{
	cookie, _:=req.Cookie("user_name")
	if(cookie== nil){
		return false;
	}else{
		_, b:=c.Get(cookie.Value)
		return b
		// if(b == false){
		// 	return b;
		// }else{
			
		// }
	}
}

func ShaConvert(user string) []uint8{
	b0:=sha1.Sum([]byte(utils.COOKIE_SALT+user))
	var t0 []uint8
	t0=make([]uint8,20)
	binary.Read(bytes.NewBuffer(b0[0:20]), binary.BigEndian, &t0)
	// fmt.Println(string(t0))
	// fmt.Println(json.Marshal(t0))
	return t0
}