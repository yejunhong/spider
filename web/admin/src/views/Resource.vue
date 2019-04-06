<template>
  <div>
    <el-button type="text" size="medium" @click="create">创建配置</el-button>
    <el-table :data="resource_list" header-cell-class-name="header" height="80vh">
      <el-table-column prop="Id" label="ID" fixed="left" width="50"></el-table-column>
      <el-table-column prop="ResourceNo" label="编号" width="80"></el-table-column>
      <el-table-column prop="ResourceName" label="资源名称" width="120"></el-table-column>
      <el-table-column prop="ResourceUrl" label="资源书籍列表地址" min-width="250"></el-table-column>
      <el-table-column prop="ConfigName" label="爬虫配置" width="150"></el-table-column>
      <el-table-column label="书籍数量">
        <template slot-scope="scope">{{book_count[scope.row.ResourceNo]}}</template>
      </el-table-column>
      <el-table-column label="操作" fixed="right" width="90">
        <template slot-scope="scope">
          <el-button type="text" size="small" @click="SpiderBook(scope.row)">下载书籍</el-button>
          <el-button type="text" size="small" @click="edit(scope.row)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      background
      layout="total, prev, pager, next"
      :page-size="pagesize"
      :total="total" style="margin-top: 10px;" @current-change="GetResource">
    </el-pagination>

    <el-dialog :title="title" :visible.sync="resourceDialog">
      <el-tabs v-model="resourceActive">
        <el-tab-pane label="资源基本信息" name="base"></el-tab-pane>
        <el-tab-pane label="爬虫配置" name="config"></el-tab-pane>
      </el-tabs>
      <el-form class="form" :model="form" v-show="resourceActive == 'base'">
        <el-form-item label="编号" :label-width="formLabelWidth">
          <el-input v-model="form.ResourceNo" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="资源名称" :label-width="formLabelWidth">
          <el-input v-model="form.ResourceName" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="书籍列表地址" :label-width="formLabelWidth">
          <el-input v-model="form.ResourceUrl" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="书籍类型" :label-width="formLabelWidth">
            <el-select v-model="form.BookType" placeholder="请选择">
              <el-option label="漫画" :value="1"></el-option>
              <el-option label="小说" :value="2"></el-option>
            </el-select>
        </el-form-item>
      </el-form>
      <div v-show="resourceActive == 'config'">
        配置路径：./node/src/config/{{form.ConfigName}}.ts
        <codemirror v-model="code" 
          :options="{mode: 'javascript',extraKeys: {'Ctrl-Space': 'autocomplete'}}">
        </codemirror>
      </div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="resourceDialog = false">取 消</el-button>
        <el-button type="primary" @click="SetResource">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
require('codemirror/mode/javascript/javascript')
require('codemirror/mode/vue/vue')
require('codemirror/addon/hint/show-hint.js')
require('codemirror/addon/hint/show-hint.css')
require('codemirror/addon/hint/javascript-hint.js')
import http from '@/lib/http'
import { codemirror } from 'vue-codemirror-lite'
export default {
  data() {
    return {
      resourceActive: "base",
      resource_list: [],
      resourceDialog: false,
      formLabelWidth: "100px",
      title: "创建配置",
      total: 0,
      page: 1,
      pagesize: 20,
      book_count: [],
      form: {
        Id: 0,
        ResourceNo: "",
        ResourceName: "",
        ResourceUrl: "",
        ConfigName: "",
        BookType: 1,
      },
      code: ""
    }
  },
  components: {
    codemirror
  },
  mounted(){
    this.GetResource(1)
  },
  methods: {
    async GetResource(p) {
      const res = await http.get('/cartoon/resource', {page: p})
      this.resource_list = res.list
      this.book_count = res.book_count
      this.total = res.count
    },
    async create(){
      this.title = "创建配置"
      this.resourceDialog = true
      this.form = { Id: 0, ResourceNo: "", ResourceName: "", ResourceUrl: "", ConfigName: "", BookType: 1}
    },
    async edit(v) {
      this.title = "修改配置"
      this.resourceDialog = true
      const res = await http.get(`/cartoon/resource/${v.Id}`,)
      this.form = res.info
      this.code = res.config
    },
    // 爬去 书籍章节
    async SpiderBook(resource) {
      await http.get('/download/book', {resourceId: resource.Id})
    },
    async SetResource() {
      this.form.ConfigText = this.code
      await http.post(`/cartoon/resource`, this.form)
      this.resourceDialog = false
      this.GetResource(1)
    }
  }
}
</script>
<style>
.el-dialog{
  width: 90vw !important;
  height: 90vh;
}
.el-form{
  height: 55vh !important;
}
.el-dialog__body{
  padding: 0 20px !important;
}
.form{
  height: 55vh !important;
}
.CodeMirror{
  height: 55vh !important;
  border: 1px solid rgb(230, 230, 230);
}
.header{
  height: 5vh;
  color: rgb(0, 0, 0);
}
</style>
