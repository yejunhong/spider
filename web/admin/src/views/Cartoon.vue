<template>
  <div class="main">
    <el-tabs v-model="tabsName" @tab-click="tabsClick">
      <el-tab-pane :label="v.ResourceName" :name="v.ResourceNo" v-for="(v, k) in tabsList" :key="k"></el-tab-pane>
    </el-tabs>
    <el-main class="cartoon-main">
      <el-row :gutter="10">
        <el-col :xs="8" :sm="6" :lg="4" :xl="3" v-for="(v, k) in cartoonList" :key="k" class="info">
          <el-card>
            <img class="card-img" src="http://www.pptok.com/wp-content/uploads/2012/08/xunguang-4.jpg" @click="SelectCartoonChapter(v)" style="width:100%">
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
        <div style="width: 40vw">
          <el-row :gutter="10">
            <el-col :xs="7" :sm="5" :lg="4" :xl="3" v-for="(v, k) in cartoonList" :key="k" class="info">
              <el-button >{{v.Id}}</el-button>
            </el-col>
          </el-row>
        </div>
        <div style="width: 59vw; height: 70vh;overflow: scroll">
        
          123123<br/>
          123123<br/>
          123123<br/>
          123123<br/>
          123123<br/>
          123123<br/>
          123123<br/>
          123123<br/>
          123123<br/>
          123123<br/>
          123123<br/>
          123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>123123<br/>
          123123<br/>

        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
// @ is an alias to /src
export default {
  name: 'cartoon',
  data () {
    return {
      dialogTitle: '',
      tabsName: '',
      tabsList: [],
      cartoonChapterShow: false,
      cartoonList: []
    }
  },
  mounted() {
    this.GetCartonnResource()
  },
  components: {
  },
  methods: {
    GetCartonnResource(){
      for(var i = 0; i < 6; i++){
        this.tabsList.push({
          Id: i,
          ResourceNo: `C${i}`,
          ResourceName: `资源名称${i}`,
        })
      }
    },
    // 获取漫画数据
    GetCartoonData(){
      this.cartoonList = []
      for(var i = 0; i < 17; i++){
        this.cartoonList.push({
          Id: i,
          ResourceName: `资源名称${i}`,
        })
      }
    },
    // 标签选择
    tabsClick() {
      this.GetCartoonData()
    },
    SelectCartoonChapter(row){
      this.cartoonChapterShow = true
      this.dialogTitle = row.ResourceName
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
    font-size: 14px;
  }
  .card-img{
    cursor: pointer;
  }
  .cartoon-main{
    height: 80vh;
    overflow: scroll-y;
    padding: 5px !important;
  }
</style>