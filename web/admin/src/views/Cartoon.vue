<template>
  <div class="main">
    <el-tabs v-model="tabsName" @tab-click="tabsClick">
      <el-tab-pane :label="v.ResourceName" :name="v.ResourceNo" v-for="(v, k) in tabsList" :key="k"></el-tab-pane>
    </el-tabs>
    <el-main class="cartoon-main">
      <el-row :gutter="20">
        <el-col :xs="12" :sm="8" :lg="6" :xl="4" v-for="(v, k) in cartoon_list" :key="k" class="info">
          <el-card>
            <div class="card-img">
              <div :class="['status', v.IsEnd == 0?'status-131345': 'status-01730']">
                {{v.IsEnd == 0 ? '连载': '完结'}} 
                {{chapters_count[v.UniqueId]?chapters_count[v.UniqueId].Number:0}}
              </div>
              <img :src="v.ResourceImgUrl" @click="SelectCartoonChapter(v)" style="height: 150%;position: relative;">
            </div>
            <span class="title">{{v.ResourceName}}</span>
            <div class="row">
              <el-button type="text" @click="SpiderBookChapter(v)" size="mini">更新</el-button>
              <span v-if="chapters_count[v.UniqueId] && chapters_count[v.UniqueId].NotNumber > 0">未抓取：{{chapters_count[v.UniqueId].NotNumber}}</span>
            </div>
            <div class="row">
              <el-button type="text" @click="DataAsyncProduce(v)" size="mini">同步到生产库</el-button>
            </div>
            <div class="row">{{v.CdateText}}</div>
          </el-card>
        </el-col>
      </el-row>
    </el-main>
    <el-pagination
      background
      layout="total, prev, pager, next"
      :page-size="pagesize"
      :total="total" style="margin-top: 10px;" @current-change="nextPage">
    </el-pagination>
    <el-dialog :title="`${dialogTitle}-章节`" :visible.sync="cartoonChapterShow" width="100%">
      <div class="chapter">
        <div style="width: 40vw;" class="scroll">
          <el-row :gutter="10">
            <el-col :xs="8" :sm="7" :lg="6" :xl="5" v-for="(v, k) in cartoon_chapter_list" :key="k" class="info">
              <el-button :type="v.Id == chapterId?'primary': ''" @click="SelectChapterContent(v)" plain>{{v.ResourceName}}</el-button>
            </el-col>
          </el-row>
        </div>
        <div v-if="this.resource.BookType == 1" style="width: 59vw;" class="scroll" ref="scrollDiv" v-loading="loading">
          <img class="chapter-img" :src="v.ResourceUrl" width="100%" v-for="(v, k) in cartoon_chapter_content" :key="k"/>
        </div>
        <div style="width: 59vw;" class="scroll" v-else v-html="cartoon_chapter.Content"></div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import http from '@/lib/http'
export default {
  name: 'cartoon',
  data () {
    return {
      loading: false,
      dialogTitle: '',
      tabsName: '',
      tabsList: [],
      total: 0,
      page: 1,
      pagesize: 20,
      cartoonChapterShow: false,
      resource: {},
      cartoon_list: [],
      cartoon_chapter: {},
      chapters_count: [],
      cartoon_chapter_list: [],
      cartoon_chapter_content: [],
      ws: {},
      chapterId: 0,
    }
  },
  mounted() {
    this.GetCartonnResource()
  },
  methods: {
    initWs(){
      const url = ''
      this.ws = new WebSocket(url)
      this.ws.onopen = () => {}
      this.ws.onmessage = this.wsMessage
      this.ws.onerror = () => {}
      this.ws.close = () => {}
    },
    nextPage(currentPage){
      this.page = currentPage
      this.GetCartoonData(this.tabsName)
    },
    async GetCartonnResource(){
      const res = await http.get('/cartoon/resource')
      this.tabsList = res.list
      this.tabsName = this.tabsList[0].ResourceNo
      this.tabsClick({name: this.tabsName})
    },
    // 获取漫画数据
    async GetCartoonData(no){
      const res = await http.get('/cartoon/list', {resource_no: no, page: this.page})
      this.chapters_count = res.chapters_count
      console.log(res.list)
      this.cartoon_list = res.list
      this.total = res.count
    },
    // 标签选择
    tabsClick(v) {
      this.GetCartoonData(v.name)
      this.tabsList.map(res => {
        if (res.ResourceNo === v.name) {
          this.resource = res
        }
      })
    },
    async SelectCartoonChapter(row){
      // this.cartoon_chapter_list = []
      this.cartoon_chapter_content = []
      this.cartoonChapterShow = true
      this.dialogTitle = row.ResourceName
      const res = await http.get('/cartoon/chapter', {list_unique_id: row.UniqueId})
      this.cartoon_chapter_list = res
      await this.SelectChapterContent(res[0])

    },
    async SelectChapterContent(row){
      this.chapterId = row.Id
      if (this.resource.BookType === 1) {
        this.cartoon_chapter_content = []
        this.loading = true
        this.$refs.scrollDiv.scrollTop = '0px'
        const res = await http.get('/cartoon/chapter/content', {chapter_unique_id: row.UniqueId})
        this.cartoon_chapter_content = res
        this.loading = false
        return
      }
      this.cartoon_chapter = row
    },
    // 爬去 书籍章节
    async SpiderBookChapter(row) {
      await http.get('/download/chapter', {bookId: row.Id})
    },
    // 同步书籍到生产库
    async DataAsyncProduce (row) {
      await http.get('/async/book', {bookId: row.Id})
    }
  }
}
</script>
<style lang="less" scope>
  .el-dialog{
    margin-top: 5vh !important;
  }
  .info{
    margin-bottom: 20px;
  }

  .el-card__body{
    padding: 5px !important;
  }
  .chapter{
    display: flex;
    flex-direction: row;
  }
  .title{
    display: flex;
    width: 100%;
    padding-top: 8px;
    padding-bottom: 8px;
    margin-top: 5px;
    margin-bottom: 5px;
    white-space:nowrap;
    font-size: 14px;
    overflow:hidden; //超出的文本隐藏
  }
  .card-img{
    cursor: pointer;
    display: flex;
    width: 100%;
    height: 230px;
    justify-content: center;
    align-content: center;
    overflow: hidden;
    position: relative;
    .status{
      position:absolute;
      /*transform: rotate(30deg);*/
      right: 0px;
      z-index: 999;
      display: flex;
      justify-content: center;
      align-items: center;
      color: #fff;
      
      font-size: 12px;
      padding: 3px;
      padding-left: 5px;
      padding-right: 5px;
    }
    .status-01730{
      background-color: rgba(0, 173, 0, 0.9);
    }
    .status-131345{
      background-color: rgba(131, 34, 5, 0.9);
    }
  }
  .cartoon-main{
    height: 80vh;
    overflow-x: hidden;
    overflow-y: scroll;
  }
  .chapter-img{
    vertical-align: bottom;
  }
  .row{
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    font-size: 14px;
    align-content: center;
    align-items: center;
    padding-top: 3px;
    padding-bottom: 3px;
    span {
      display: flex;
      align-content: center;
      font-size: 14px;
    }
  }
  .scroll {
    height: 70vh;
    overflow-x: hidden;
    overflow-y: scroll;
    padding: 20px;
  }
  
</style>