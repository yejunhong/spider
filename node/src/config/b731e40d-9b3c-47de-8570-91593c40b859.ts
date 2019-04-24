import url from 'url';
// print 打印
/*
PHPSESSID=hh9i4d686jclve8t3v861mke71; 
Hm_lvt_15a0ce9456e6c6618d9cb8f8c96e30f9=1555725551; 
reaua=FYR2F2; 
reauava=http%3A%2F%2Fstore.cqhdx.com%2Fstatic%2Fserimgs%2Favatar%2F6.jpg; 
reauathf=40ae44844eab6102a09e8bddcd347873; 
reauathe=4823568e49d2de3f6f4947592984275a; 
reauathn=36c4deebe518e9b2b06be12e62d07823; 
reauathx=fb7f4c461609db34df008f0b54f6e482; 
reauath3=23d5a1d652525ee6835a8a69fe8fbdc8; 
reauatho=1284e68c7697b9d54363a4dab255a0c2; 
reauathp=4b7b31564653d4527fff0fd0be0de9cc; 
reauathq=c28865c745c5bb315789e9805596254a; 
reauathb=537d679636940dde506e6fd93f3a142c; 
reauathj=20526c058bd1d5d9eaf54a6fa6c02eec; 
history_cookie_ppm=null; 
Hm_lpvt_15a0ce9456e6c6618d9cb8f8c96e30f9=1555820011
*/
module.exports.Page = {
  name: '紫漫',
  mobile: "",
  cookie: [
    {name: "PHPSESSID", value:"qk0h2r1b608h4hfu2v7pfltskr", domain: "ttxiefu.com"},
  ],
  user_agent: "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
}
/*
<li>
  <a class="ImgA autoHeight" href="/100/comic/6421">
    <img src="http://store.cqhdx.com/static/serimgs/comics/1a/f0a373463d6fb911b57b48deac8172.jpg" width="100%" style="height: 148px;">
  </a>
  <a class="txtA">偷窥</a>
  <span class="info">作者：REC</span>
</li>
*/
// 获取书籍列表
module.exports.Book = {
  // selector: 'body', // 列表选择器
  print: true,
  selector: 'ul#classify_container li', // 列表选择器
  async handle(res: any, Element: any): Promise<any> {
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      resdata.push({
        tags: "",
        detail: "",
        resource_name: await e.Html('a.txtA'),
        resource_url: `http://ttxiefu.com` + await e.Attr('a', 'href'),
        resource_img_url: await e.Attr('img', 'src'),
        author: await e.Html('span.info')
      })
    }
    return resdata
  },
  scroll: true, // 是否滚动页面
}

/*
<li>
<div class="thumb_ep" style="background-image: url(http://store.cqhdx.com/static/serimgs/comics/6426/vol/115927/98/f782f799baff7c5d084d4b59c19c9d.jpg);"></div>
<a href="/100/100/view/6426/115927" onclick="chapterCookie(6426,115927,1,'第1话')">
  <span>第1话<br><span class="sub">2018-12-17</span></span>
</a>
<p class="coin_price"><span class="btn_free">免费</span></p>
</li>
*/
// 数据章节配置
module.exports.Chapter = {
  selector: 'body', // 列表选择器
  async event(page: any){
    if (await page.$('ul.autoHeight li.add') != null) {
      await page.tap('ul.autoHeight li.add');
      await page.waitFor(500);
    }
  },
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata: any = [];
    const list = await res[0].$$('ul.autoHeight li');
    for(const v of list){
      const e = new Element(v);
      let is_free = "1";
      if(await e.Attr('.coin_price span', 'class') == 'btn_free'){
        is_free = "0";
      }
      const resource_name = await e.Html('a > span');
      const resource_url = await e.Attr('div.thumb_ep', 'style')
      const c = resource_url.split("(")[1].split(")")[0]
      resdata.push({
        is_free: is_free,
        resource_name: resource_name.split("<br>")[0],
        resource_url: `http://ttxiefu.com` + await e.Attr('a', 'href'),
        resource_img_url: c,
      });
    }
    const info = await res[0].$eval('div.sub_r', (e: any) => {
      const tags = e.querySelectorAll('p.txtItme');
      const res: any = [];
      for (const t of tags[1].querySelectorAll("a.pd")) {
        res.push(t.innerHTML);
      }
      let is_end = "0";
      if (tags[2].innerHTML.indexOf("已完结") > 0) {
        is_end = "1";
      }
      return {
        tags: res.join(","),
        detail: "",
        is_end: is_end,
      };
    });
    info.detail = await res[0].$eval('p.txtDesc', (e: any) => e.innerHTML);
    return {list: resdata, detail: info};
  }
}

module.exports.Content = {
  selector: 'div.charpetBox', // 列表选择器
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    if(res[0] != undefined) {
      const resdata = await res[0].$$eval('img', (e: any) => {
        const res: any = [];
        for (const v of e) {
          res.push({
            resource_img_url: v.getAttribute('data-original'),
          });
        }
        return res;
      });
      return resdata
    }
    return [];
  },
}