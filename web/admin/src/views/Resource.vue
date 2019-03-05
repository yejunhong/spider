<template>
  <div>
    <el-table :data="resource_list" style="width: 100%" header-cell-class-name="header" height="90vh">
      <el-table-column prop="id" label="ID" fixed="left" width="80"></el-table-column>
      <el-table-column prop="resource_no" label="编号" width="100"></el-table-column>
      <el-table-column prop="resource_name" label="资源名称" width="120"></el-table-column>
      <el-table-column prop="resource_url" label="资源书籍列表地址" min-width="350"></el-table-column>
      <el-table-column prop="book_count" label="书籍数量" width="100"></el-table-column>
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
  export default {
    data() {
      return {
        resource_list: [{
          id: '1',
          resource_no: 'C001',
          resource_name: '漫画书籍',
          resource_url: 'http://127.0.0.1:8080/resource',
          book_count: 20
        }, {
          id: '1',
          resource_no: 'C001',
          resource_name: '漫画书籍',
          resource_url: 'http://127.0.0.1:8080/resource',
          book_count: 20
        }, {
          id: '1',
          resource_no: 'C001',
          resource_name: '漫画书籍',
          resource_url: 'http://127.0.0.1:8080/resource',
          book_count: 20
        }, {
          id: '1',
          resource_no: 'C001',
          resource_name: '漫画书籍',
          resource_url: 'http://127.0.0.1:8080/resource',
          book_count: 20
        }]
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
