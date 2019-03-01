<template>
  <div class="main">
    <el-tabs v-model="tabsName" @tab-click="tabsClick">
      <el-tab-pane :label="v.ResourceName" :name="v.ResourceNo" v-for="(v, k) in tabsList" :key="k"></el-tab-pane>
    </el-tabs>
    <el-row :gutter="10">
      <el-col :xs="8" :sm="6" :lg="4" :xl="3" v-for="(v, k) in cartoonList" :key="k" class="info">
        <el-card>
          <img src="http://www.pptok.com/wp-content/uploads/2012/08/xunguang-4.jpg" style="width:100%">
          <div style="padding: 14px;">
            <span>{{v.ResourceName}}</span>
            <div class="bottom clearfix">
              <el-button type="text" class="button" @click="SelectCartoonChapter(v)">查看</el-button>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-dialog :title="`${dialogTitle}-章节`" :visible.sync="cartoonChapterShow" width="98%">
      <span>这是一段信息</span>
      <span slot="footer" class="dialog-footer"></span>
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
<style lang="less">
  .info{
    margin-bottom: 10px;
  }
</style>