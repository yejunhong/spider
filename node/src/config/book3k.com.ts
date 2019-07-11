import url from 'url';
// print 打印
module.exports.Page = {
  name: '蓝光文学',
  mobile: "",
}

// 获取书籍列表
module.exports.Book = {
  selector: '没有书籍概念', // 列表选择器
  print: false,
  // selector: 'table.bookshelfListMain', // 列表选择器
  async handle(res: any): Promise<any> {
    return [];
  },
  scroll: true, // 是否滚动页面
}

// 数据章节配置
/*
<div class="post-column clearfix">
	<article id="post-14156" class="post-14156 post type-post status-publish format-standard hentry category-25 tag-17 tag-3 tag-107 tag-31 tag-37 tag-39 tag-32 tag-85 tag-9 tag-10 tag-45 tag-106 tag-104 tag-87">
		<header class="entry-header">
			<div class="entry-meta"><span class="meta-date"><a href="https://book3k.com/zh-cn/archives/14156" title="08:51" rel="bookmark"><time class="entry-date published updated" datetime="2019-07-08T08:51:38+00:00">08/07/2019</time></a></span><span class="meta-category"> <a href="https://book3k.com/zh-cn/archives/category/%e8%bf%91%e8%a6%aa%e4%ba%82%e5%80%ab" rel="category tag">近亲乱伦</a></span></div>
			<h2 class="entry-title"><a href="https://book3k.com/zh-cn/archives/14156" rel="bookmark">兄妹失乐园</a></h2>
		</header><!-- .entry-header -->
		<div class="entry-content entry-excerpt clearfix">
			<p>　　兄妹失乐园（１）——校车篇 当当当……放学的钟声跟往常一样的响起了，…放学了…终于结束了漫长的一天。我走出了教室，拖着沈重的脚步，一步一步的往校车的方向移动。 我远远看到了我妹从另一栋教室走了过来</p>
		<a href="https://book3k.com/zh-cn/archives/14156" class="more-link">Continue reading »</a>
				</div><!-- .entry-content -->
	</article>

</div>
*/
module.exports.Chapter = {
  selector: 'div#post-wrapper', // 列表选择器
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata = await res[0].$$eval('div.post-column', (e: any) => {
      let resdata: any = [];
      for(const v of e) {
        const a = v.querySelector('h2 a');
        let detail = "";
        if (v.querySelector('div.entry-content p')) {
          detail = v.querySelector('div.entry-content p').innerHTML;
        }
        resdata.push({
          is_free: "0",
          detail: detail,
          resource_name: a.innerHTML,
          resource_url: `${a.getAttribute("href")}`,
        })
      }
      return resdata
    });
    console.log(resdata)
    return resdata;
  },
  // 爬取 下一页数据
  next: {
    selector: 'body', // 列表选择器
    async handle(e: any, urlStr: string){
      const maxPage = 1272;
      const u = urlStr.split("/");
      let page = parseInt(u[u.length - 1]) + 1
      if (page > maxPage) {
        return false;
      }
      // https://book3k.com/zh-cn/page/1
      return `https://book3k.com/zh-cn/page/${page}`
    }
  },
}

module.exports.Content = {
  selector: 'body', // 列表选择器
  print: false,
  // browser_request: false, // 是通过浏览器请求
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata: any = [];
    const text = await res[0].$eval("div.entry-content", (e: any) => {
      if (e.querySelector("iframe")) {
        e.querySelector("iframe").remove();
      }
      return e.innerHTML;
    });
    resdata.push({
      resource_img_url: text,
    })
    return resdata
  },
}