<template>
  <div class="Footer">
    2021-现在｜学生管理系统 By FJD团队-DengXianSheng｜
    <el-button type="text" @click="drawer = true">免责声明</el-button>｜
    <el-link
      style="margin-top: -2px;"
      :underline="false"
      href="https://beian.miit.gov.cn/"
      target="_blank"
    >鲁ICP备2020042720号</el-link>
    <el-drawer
      title="免责声明"
      :visible.sync="drawer"
      :direction="direction"
      :destroy-on-close="true"
    >
      <span>本网站所提供的信息，只供参考之用。</span>
      <el-divider></el-divider>
      <span>本网站及其雇员一概毋须以任何方式就任何信息传递或传送的失误、不准确或错误对用户或任何其他人士负任何直接或间接的责任。</span>
      <el-divider></el-divider>
      <span>在法律允许的范围内，本网站在此声明,不承担用户或任何人士就使用或未能使用本网站所提供的信息或任何链接或项目所引致的任何直接、间接、附带、从属、特殊、惩罚性或惩戒性的损害赔偿（包括但不限于收益、预期利润的损失或失去的业务、未实现预期的节省）。</span>
      <el-divider></el-divider>
      <span>本网站所提供的信息，若在任何司法管辖地区供任何人士使用或分发给任何人士时会违反该司法管辖地区的法律或条例的规定或会导致本网站或其第三方代理人受限于该司法管辖地区内的任何监管规定时，则该等信息不宜在该司法管辖地区供该等任何人士使用或分发给该等任何人士。</span>
      <el-divider></el-divider>
      <span>用户须自行保证不会受限于任何限制或禁止用户使用或分发本网站所提供信息的当地的规定。</span>
      <el-divider></el-divider>
      <span>凡以任何方式登陆本网站或直接、间接使用本网站资料者，视为自愿接受本网站声明的约束。</span>
      <el-divider></el-divider>
      <span>本声明最后更新于2021年12月14日，最终解释权由FJD团队所有。</span>
    </el-drawer>
  </div>
</template>

<style>
.Footer {
  color: #606266;
  font-size: 14px;
}
</style>

<script>
export default {
  data() {
    return {
      drawer: false,
      direction: "rtl",
    };
  },
};
</script>
