<template>
  <div>
    <el-table :data="resource_list" style="width: 100%" header-cell-class-name="header" height="90vh">
      <el-table-column prop="Id" label="ID" fixed="left" width="80"></el-table-column>
      <el-table-column prop="ResourceNo" label="编号" width="100"></el-table-column>
      <el-table-column prop="ResourceName" label="资源名称" width="120"></el-table-column>
      <el-table-column prop="ResourceUrl" label="资源书籍列表地址" min-width="350"></el-table-column>
      <el-table-column prop="ConfigName" label="爬虫配置" width="200"></el-table-column>
      <el-table-column prop="BookCount" label="书籍数量" width="100"></el-table-column>
      <el-table-column label="操作" fixed="right" width="120">
        <template slot-scope="scope">
          <el-button type="text" size="small">下载</el-button>
          <el-button type="text" size="small" @click="edit(scope.row)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      background
      layout="prev, pager, next"
      :total="1000">
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
      </el-form>
      <div v-show="resourceActive == 'config'">
        配置路径：{{form.ConfigName}}
        <codemirror v-model="code" 
          :options="{mode: 'javascript',extraKeys: {'Ctrl-Space': 'autocomplete'}}">
        </codemirror>
      </div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="resourceDialog = false">取 消</el-button>
        <el-button type="primary" @click="resourceDialog = false">确 定</el-button>
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
      title: "",
      form: {
        Id: 0,
        ResourceNo: "",
        ResourceName: "",
        ResourceUrl: "",
        ConfigName: "",
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
    },
    async edit(v) {
      this.resourceDialog = true
      const res = await http.get(`/cartoon/resource/${v.Id}`,)
      this.form = res.info
      this.code = res.config
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
  height: 430px !important;
}
.CodeMirror{
  height: 430px !important;
  border: 1px solid rgb(230, 230, 230);
}
.header{
  height: 5vh;
  color: rgb(0, 0, 0);
}
</style>
