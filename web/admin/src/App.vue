<template>
  <el-container id="app">
    <el-aside width="200px">
      <el-menu default-active="1-4-1" class="menu" router :collapse="isCollapse">
        <el-menu-item index="resource">
          <i class="el-icon-date"></i>
          <span slot="title">资源管理</span>
        </el-menu-item>
        <el-menu-item index="cartoon">
          <i class="el-icon-menu"></i>
          <span slot="title">书籍管理</span>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-main><router-view/></el-main>
  </el-container>
</template>
<script>
  export default {
    data() {
      return {
        isCollapse: false
      };
    },
    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      }
    }
  }
</script>
<style lang="less">
#app, div, html, body {
  margin: 0;
  padding: 0;
}
.menu{
  height: 100vh;
}
.el-main{
  padding: 10px !important;
}
</style>
