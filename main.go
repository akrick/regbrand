package main

import (
	"flag"
	"log"
)

func main(){
	var period, category, totalPage int
	cookie := "user_3746185747=ea2ce6577045e4015c595995ca219423; NTKF_T2D_CLIENTID=guestC183A1CC-152C-CAD5-629A-7F42E7E45630; anonymousId=1757f43e1255b5-02a4bd9d62358e-3e604000-921600-1757f43e1264c1; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%221757f43e1255b5-02a4bd9d62358e-3e604000-921600-1757f43e1264c1%22%2C%22%24device_id%22%3A%221757f43e1255b5-02a4bd9d62358e-3e604000-921600-1757f43e1264c1%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_referrer_host%22%3A%22%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%7D; nTalk_CACHE_DATA={uid:kf_9479_ISME9754_guestC183A1CC-152C-CA,tid:1604204777230454}; Hm_lvt_df2da21ec003ed3f44bbde6cbef22d1c=1604157893,1604204778; QDS_COOKIE=user%3Ainfo%3A384F3235-E388-0ED9-EF15-F1C4CD149308; INGRESSCOOKIE=1604205450.005.31301.463143; _csrf=fb491f27f9843c6108250f20c49d28d223e211df0bf502a08864973afd277c74a%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_csrf%22%3Bi%3A1%3Bs%3A32%3A%22Z9g9VlOnCOKGioJtTjLLBFe9xOaPF1VT%22%3B%7D; PHPSESSID=d2b7a92ab8ac63c248b10aa9d0bc3e1d; _pk_ref.3.7c47=%5B%22%22%2C%22%22%2C1604205452%2C%22https%3A%2F%2Fwww.quandashi.com%2F%22%5D; _pk_testcookie.3.7c47=1; _pk_id.3.7c47=ef48b1e4a74ab117.1604157942.3.1604205464.1604205452.; Hm_lpvt_df2da21ec003ed3f44bbde6cbef22d1c=1604205464"
	url := "https://so.quandashi.com/search/notice/search-notice"
	flag.IntVar(&period, "p", 0, "期号")
	flag.IntVar(&category, "c", 0, "分类号")
	flag.IntVar(&totalPage, "t", 0, "分页总数")
	flag.StringVar(&cookie, "cookie", cookie, "Cookie值")
	flag.Parse()

	scraper := new(QdsScraper)
	scraper.SetCookie(cookie)
	scraper.SetUrl(url)
	err := scraper.Search(period, category, totalPage)
	if err != nil {
		log.Fatal(err)
	}
}
