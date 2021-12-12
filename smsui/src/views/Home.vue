<template>
  <div>
    <el-container>
      <el-header><Header /></el-header>
      <el-main><Main /></el-main>
      <el-footer><Footer /></el-footer>
    </el-container>
  </div>
</template>

<script>
import Header from "./header.vue";
import Main from "./main.vue";
import Footer from "./footer.vue";
export default {
  name: "Home",
  components: {
    Main,
    Header,
    Footer,
  },
};
</script>

<style>
.el-header,
.el-footer {
  /* background-color: #d2e8ff; */
  /* color: #303133; */
  text-align: center;
  line-height: 60px;
}

.el-aside {
  /* background-color: #d2e8ff; */
  /* color: #303133; */
  text-align: center;
  line-height: 200px;
}

.el-main {
  /* background-color: #e9f4ff; */
  /* color: #303133; */
  text-align: center;
  line-height: 160px;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>