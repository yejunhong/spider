<template>
  <div class="main">
    <el-tabs v-model="tabsName" @tab-click="tabsClick">
      <el-tab-pane :label="v.ResourceName" :name="v.ResourceNo" v-for="(v, k) in tabsList" :key="k"></el-tab-pane>
    </el-tabs>
    <el-main class="cartoon-main">
      <el-row :gutter="10">
        <el-col :xs="8" :sm="6" :lg="4" :xl="3" v-for="(v, k) in cartoon_list" :key="k" class="info">
          <el-card>
            <div class="card-img">
              <img :src="v.ResourceImgUrl" @click="SelectCartoonChapter(v)" style="height: 150%;position: relative;">
            </div>
            <span class="title">{{v.ResourceName}}</span>
            <div class="row">
              <el-button type="text" @click="SelectCartoonChapter(v)" size="mini">下载章节列表</el-button>
            </div>
            <div class="row">
              <el-button type="text" @click="SelectCartoonChapter(v)" size="mini">下载章节内容</el-button>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </el-main>
    <el-dialog :title="`${dialogTitle}-章节`" :visible.sync="cartoonChapterShow" width="98%">
      <div class="chapter">
        <div style="width: 40vw; height: 80vh;overflow: scroll">
          <el-row :gutter="10">
            <el-col :xs="8" :sm="7" :lg="6" :xl="5" v-for="(v, k) in cartoon_chapter_list" :key="k" class="info">
              <el-button @click="SelectChapterContent(v)">{{v.ResourceName}}</el-button>
            </el-col>
          </el-row>
        </div>
        <el-main style="width: 59vw; height: 80vh;overflow: scroll" ref="scrollDiv" v-loading="loading">
          <img class="chapter-img" :src="v.ResourceUrl" width="100%" v-for="(v, k) in cartoon_chapter_content" :key="k"/>
        </el-main>
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
      cartoonChapterShow: false,
      cartoon_list: [],
      cartoon_chapter_list: [],
      cartoon_chapter_content: []
    }
  },
  mounted() {
    this.GetCartonnResource()
  },
  methods: {
    async GetCartonnResource(){
      const res = await http.get('/cartoon/resource')
      this.tabsList = res.list
      this.GetCartoonData(this.tabsList[0].ResourceNo)
      this.tabsName = this.tabsList[0].ResourceNo
    },
    // 获取漫画数据
    async GetCartoonData(no){
      const res = await http.get('/cartoon/list', {resource_no: no})
      this.cartoon_list = res.list
    },
    // 标签选择
    tabsClick(v) {
      this.GetCartoonData(v.name)
    },
    async SelectCartoonChapter(row){
      this.cartoon_chapter_content = []
      this.cartoonChapterShow = true
      this.dialogTitle = row.ResourceName
      const res = await http.get('/cartoon/chapter', {list_unique_id: row.UniqueId})
      this.cartoon_chapter_list = res
      await this.SelectChapterContent(res[0])
    },
    async SelectChapterContent(row){
      this.cartoon_chapter_content = []
      this.loading = true
      this.$refs.scrollDiv.scrollTop = '0px'
      const res = await http.get('/cartoon/chapter/content', {chapter_unique_id: row.UniqueId})
      this.cartoon_chapter_content = res
      this.loading = false
    }
  }
}
</script>
<style lang="less" scope>
  .el-dialog{
    margin-top: 5vh !important;
  }
  .info{
    margin-bottom: 10px;
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
    justify-content: center;
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
    height: 150px;
    justify-content: center;
    align-content: center;
    overflow: hidden;
  }
  .cartoon-main{
    height: 80vh;
    overflow: scroll-y;
    padding: 5px !important;
  }
  .chapter-img{
    vertical-align: bottom;
  }
</style>