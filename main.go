package main

import (
	"flag"
	"log"
)

func main(){
	var period, category, totalPage int
	var website, cookie, url string
	flag.StringVar(&website, "w", "", "网站")
	flag.IntVar(&period, "p", 0, "期号")
	flag.IntVar(&category, "c", 0, "分类号")
	flag.IntVar(&totalPage, "t", 0, "分页总数")
	flag.StringVar(&cookie, "cookie", "", "Cookie值")
	flag.Parse()
	if website == "qds" {
		scraper := new(QdsScraper)
		if cookie == ""{
			cookie = "user_3746185747=ea2ce6577045e4015c595995ca219423; NTKF_T2D_CLIENTID=guestC183A1CC-152C-CAD5-629A-7F42E7E45630; anonymousId=1757f43e1255b5-02a4bd9d62358e-3e604000-921600-1757f43e1264c1; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%221757f43e1255b5-02a4bd9d62358e-3e604000-921600-1757f43e1264c1%22%2C%22%24device_id%22%3A%221757f43e1255b5-02a4bd9d62358e-3e604000-921600-1757f43e1264c1%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_referrer_host%22%3A%22%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%7D; nTalk_CACHE_DATA={uid:kf_9479_ISME9754_guestC183A1CC-152C-CA,tid:1604204777230454}; Hm_lvt_df2da21ec003ed3f44bbde6cbef22d1c=1604157893,1604204778; QDS_COOKIE=user%3Ainfo%3A384F3235-E388-0ED9-EF15-F1C4CD149308; INGRESSCOOKIE=1604205450.005.31301.463143; _csrf=fb491f27f9843c6108250f20c49d28d223e211df0bf502a08864973afd277c74a%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_csrf%22%3Bi%3A1%3Bs%3A32%3A%22Z9g9VlOnCOKGioJtTjLLBFe9xOaPF1VT%22%3B%7D; PHPSESSID=d2b7a92ab8ac63c248b10aa9d0bc3e1d; _pk_ref.3.7c47=%5B%22%22%2C%22%22%2C1604205452%2C%22https%3A%2F%2Fwww.quandashi.com%2F%22%5D; _pk_testcookie.3.7c47=1; _pk_id.3.7c47=ef48b1e4a74ab117.1604157942.3.1604205464.1604205452.; Hm_lpvt_df2da21ec003ed3f44bbde6cbef22d1c=1604205464"
		}
		url = "https://so.quandashi.com/search/notice/search-notice"
		scraper.SetCookie(cookie)
		scraper.SetUrl(url)
		err := scraper.Search(period, category, totalPage)
		if err != nil {
			log.Fatal(err)
		}
	}else if website == "qmx"{

		scraper := new(QmxScraper)
		if cookie == ""{
			cookie = "UM_distinctid=1757cd5d5581e2-045b7d7d920b1c-3e604000-e1000-1757cd5d55922a; Hm_lvt_dee908f345d388e5c360ef124ec330eb=1607065289,1607165818; CNZZDATA1276743869=1708044405-1604116748-%7C1607161864; serverToken=s%3Ad1dfb241-531e-4421-a673-eecfec502aba.8%2FYZ3UbkcaDieTP0%2FN012LkRe9xBAWnzHvyT5xVO6Kw; authUser=s%3Aj%3A%7B%22id%22%3A%22772069400386174976%22%2C%22username%22%3A%2215918710508%22%2C%22password%22%3A%22%242a%2410%24QPAr.Yq8OAlwOxnlarCAmudzimDBvZx3tBpKlzVcnjmgjSrur9h96%22%2C%22nickname%22%3A%22D..%22%2C%22headImgUrl%22%3A%22http%3A%2F%2Fthirdwx.qlogo.cn%2Fmmopen%2F3Lqm1xHojtbTDywcresMvEGIKD9Y4pov5cEjXiaPyZrV4TjmaibibLED3b1x4QkDUSBmQF6kicSXKSjN8iagILp2iat5kibs9KPeqbG%2F132%22%2C%22phone%22%3A%2215918710508%22%2C%22sex%22%3A1%2C%22enabled%22%3Atrue%2C%22type%22%3A%22APP_WECHATLOGIN%22%2C%22email%22%3Anull%2C%22userLevel%22%3A1%2C%22birthday%22%3Anull%2C%22province%22%3Anull%2C%22city%22%3Anull%2C%22county%22%3Anull%2C%22belongSystem%22%3Anull%2C%22isOpenReport%22%3Anull%2C%22recommendNum%22%3Anull%2C%22myRecommendNum%22%3Anull%2C%22discount%22%3Anull%2C%22phoneDistrict%22%3A%22%22%2C%22source%22%3A%220%22%2C%22serviceId%22%3Anull%2C%22serviceName%22%3Anull%2C%22clueType%22%3Anull%2C%22clueRemark%22%3Anull%2C%22clueRemarkDate%22%3Anull%2C%22remark%22%3Anull%2C%22createTime%22%3A1604117290000%2C%22updateTime%22%3A1604122178000%2C%22sysRoles%22%3A%5B%5D%2C%22permissions%22%3A%22%22%2C%22accountNonExpired%22%3Atrue%2C%22credentialsNonExpired%22%3Atrue%2C%22accountNonLocked%22%3Atrue%7D.zAdTaCC8mVzFcrfNARAi6kr8S168qDvwk502%2BIxddmg; Hm_lpvt_dee908f345d388e5c360ef124ec330eb=1607165928"
		}
		url = "https://sbgg.qmxip.com/x1"
		scraper.SetCookie(cookie)
		scraper.SetUrl(url)
		err := scraper.Search(period, category, totalPage)
		if err != nil {
			log.Fatal(err)
		}
	}
}
