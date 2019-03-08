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
          <el-button type="text" size="small">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      background
      layout="prev, pager, next"
      :total="1000">
    </el-pagination>
  </div>
</template>

<script>
import http from '@/lib/http'
export default {
  data() {
    return {
      resource_list: []
    }
  },
  mounted(){
    this.GetResource(1)
  },
  methods: {
    async GetResource(page){
      const res = await http.get('/cartoon/resource')
      this.resource_list = res.list
    }
  }
}
</script>
<style>
.header{
  height: 5vh;
  color: rgb(0, 0, 0);
}
</style>
