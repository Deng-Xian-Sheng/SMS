<template>
  <div>
    <el-container>
      <el-header>
        <Header />
      </el-header>
      <el-main>
        <Main/>
      </el-main>
      <el-footer>
        <Footer />
      </el-footer>
    </el-container>
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
import Main from "./components/Main.vue";
export default {
  components: {
    Header,
    Footer,
    Main,
  },
};
</script>

<style>
.el-header,
.el-footer {
  background-color: #d2e8ff;
  color: #303133;
  text-align: center;
  line-height: 60px;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
}

.el-header{
  border-bottom-right-radius:0px;
  border-bottom-left-radius:0px;
}
.el-footer{
  border-top-right-radius:0px;
  border-top-left-radius:0px;
}

.el-aside {
  background-color: #d2e8ff;
  color: #303133;
  text-align: center;
  line-height: 200px;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
}

.el-main {
  background-color: #e9f4ff;
  color: #303133;
  text-align: center;
  line-height: 160px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
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