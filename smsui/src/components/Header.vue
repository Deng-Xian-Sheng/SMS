<template>
  <div>
    <el-row type="flex" class="row-bg" justify="center">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          <el-row>
            <a href="/" target="_blank">
              <el-col :span="2" :push="10">
                <div class="grid-content bg-purple">
                  <div class="demo-image">
                    <div class="block">
                      <el-image
                        style="width: 50px; height: 50px;margin-top: 5px;"
                        :src="url"
                        :fit="fit"
                      >
                        <div slot="error" class="image-slot">
                          <i class="el-icon-picture-outline"></i>
                        </div>
                      </el-image>
                    </div>
                  </div>
                </div>
              </el-col>
            </a>
            <span>
              <el-col :span="12" :push="9" class="hidden-md-and-down">
                <div class="grid-content bg-purple-light">
                  <span v-on:click="TextLink" style="cursor:pointer;" >学生管理系统</span>
                </div>
              </el-col>
            </span>
          </el-row>
        </div>
      </el-col>
      <el-col :span="12">
        <div class="grid-content bg-purple-light">
          <el-menu
            :default-active="activeIndex"
            class="el-menu-demo"
            mode="horizontal"
            background-color="#d2e8ff"
            @select="handleSelect"
          >
            <el-menu-item index="1">首页</el-menu-item>
            <el-menu-item index="2">学生管理</el-menu-item>
            <el-menu-item index="3">公告中心</el-menu-item>
            <el-menu-item index="4">帮助说明</el-menu-item>
          </el-menu>
          <div class="line"></div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">demo</div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "Header",
  data() {
    return {
      activeIndex: "1",
      fit: "cover",
      url: "/favicon.png",
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      if (key == 1 && this.$route.path != "/") {
        this.$router.push({ path: "/" });
      } else if (key == 2 && this.$route.path != "/studentmanagement") {
        this.$router.push({ path: "/studentmanagement" });
      } else if (key == 3 && this.$route.path != "/announcementcenter") {
        this.$router.push({ path: "/announcementcenter" });
      } else if (key == 4 && this.$route.path != "/help") {
        this.$router.push({ path: "/help" });
      }
      console.log(key, keyPath);
    },
    TextLink:function(){
      window.open("/")
    },
  },
};
</script>