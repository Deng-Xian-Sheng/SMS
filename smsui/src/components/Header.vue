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
                  <span v-on:click="TextLink" style="cursor:pointer;">学生管理系统</span>
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
        <div class="grid-content bg-purple">
          <el-button type="text" @click="centerDialogVisible = true">注册/登录</el-button>
        </div>
      </el-col>
    </el-row>
    <el-dialog
      title="登录"
      :visible.sync="centerDialogVisible"
      width="40%"
      center
      :destroy-on-close="true"
      custom-class="RegisterLogin"
    >
      <el-row>
        <el-col :span="4">
          <div class="grid-content bg-purple">用户名：</div>
        </el-col>
        <el-col :span="20">
          <div class="grid-content bg-purple-light">
            <el-input placeholder="请输入用户名" v-model="inputName" clearable></el-input>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="4">
          <div class="grid-content bg-purple">密码：</div>
        </el-col>
        <el-col :span="20">
          <div class="grid-content bg-purple-light">
            <el-input placeholder="请输入密码" v-model="inputPasswd" show-password></el-input>
          </div>
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="12" :offset="4">
          <div class="grid-content bg-purple">
            <slide-verify
              ref="slideblock"
              @again="onAgain"
              @fulfilled="onFulfilled"
              @success="onSuccess"
              @fail="onFail"
              @refresh="onRefresh"
              :slider-text="verificationCodeTips"
              :accuracy="accuracy"
              :imgs="slideblockImgs"
            ></slide-verify>
          </div>
        </el-col>
      </el-row>
      <el-row :gutter="40">
        <el-col :span="12" :offset="9">
          <div class="grid-content bg-purple">
            <el-button type="text" @click="centerDialogVisible = false;RegisterDialogVisible = true" >注册</el-button>
            <span>/</span>
            <el-button type="text">忘记密码</el-button>
          </div>
        </el-col>
      </el-row>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary">马上登录</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="注册"
      :visible.sync="RegisterDialogVisible"
      width="40%"
      center
      :destroy-on-close="true"
      custom-class="RegisterLogin"
    >
      <el-row>
        <el-col :span="4">
          <div class="grid-content bg-purple">用户名：</div>
        </el-col>
        <el-col :span="20">
          <div class="grid-content bg-purple-light">
            <el-input placeholder="请输入用户名" v-model="inputNameRegister" clearable></el-input>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="4">
          <div class="grid-content bg-purple">密码：</div>
        </el-col>
        <el-col :span="20">
          <div class="grid-content bg-purple-light">
            <el-input placeholder="请输入密码" v-model="inputPasswdRegister" show-password></el-input>
          </div>
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="12" :offset="4">
          <div class="grid-content bg-purple">
            <slide-verify
              ref="slideblock"
              @again="onAgain"
              @fulfilled="onFulfilled"
              @success="onSuccess"
              @fail="onFail"
              @refresh="onRefresh"
              :slider-text="verificationCodeTips"
              :accuracy="accuracy"
              :imgs="slideblockImgs"
            ></slide-verify>
          </div>
        </el-col>
      </el-row>
      <el-row :gutter="40">
        <el-col :span="12" :offset="9">
          <div class="grid-content bg-purple">
            <el-button type="text" @click="RegisterDialogVisible = false;centerDialogVisible = true">登录</el-button>
            <span>/</span>
            <el-button type="text">忘记密码</el-button>
          </div>
        </el-col>
      </el-row>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary">马上注册</el-button>
      </span>
    </el-dialog>
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
      centerDialogVisible: false,
      inputName: "",
      inputPasswd: "",
      verificationCodeStatus: "",
      verificationCodeTips: "向右滑动->",
      // 精确度小，可允许的误差范围小；为1时，则表示滑块要与凹槽完全重叠，才能验证成功。默认值为5
      accuracy: 1,
      slideblockImgs: [
        "/image/VerificationCodeImage/1023-310x155.jpg",
        "/image/VerificationCodeImage/32-310x155.jpg",
        "/image/VerificationCodeImage/433-310x155.jpg",
        "/image/VerificationCodeImage/622-310x155.jpg",
        "/image/VerificationCodeImage/87-310x155.jpg",
        "/image/VerificationCodeImage/1072-310x155.jpg",
        "/image/VerificationCodeImage/334-310x155.jpg",
        "/image/VerificationCodeImage/503-310x155.jpg",
        "/image/VerificationCodeImage/691-310x155.jpg",
        "/image/VerificationCodeImage/931-310x155.jpg",
        "/image/VerificationCodeImage/13-310x155.jpg",
        "/image/VerificationCodeImage/342-310x155.jpg",
        "/image/VerificationCodeImage/538-310x155.jpg",
        "/image/VerificationCodeImage/701-310x155.jpg",
        "/image/VerificationCodeImage/937-310x155.jpg",
        "/image/VerificationCodeImage/145-310x155.jpg",
        "/image/VerificationCodeImage/392-1-310x155.jpg",
        "/image/VerificationCodeImage/554-310x155.jpg",
        "/image/VerificationCodeImage/76-310x155.jpg",
        "/image/VerificationCodeImage/943-310x155.jpg",
        "/image/VerificationCodeImage/243-310x155.jpg",
        "/image/VerificationCodeImage/392-310x155.jpg",
        "/image/VerificationCodeImage/559-310x155.jpg",
        "/image/VerificationCodeImage/820-310x155.jpg",
        "/image/VerificationCodeImage/954-310x155.jpg",
        "/image/VerificationCodeImage/272-310x155.jpg",
        "/image/VerificationCodeImage/402-310x155.jpg",
        "/image/VerificationCodeImage/575-310x155.jpg",
        "/image/VerificationCodeImage/832-310x155.jpg",
        "/image/VerificationCodeImage/988-310x155.jpg",
      ],
      RegisterDialogVisible: false,
      inputNameRegister: "",
      inputPasswdRegister: "",
    };
  },
  methods: {
    handleSelect(key) {
      if (key == 1 && this.$route.path != "/") {
        this.$router.push({ path: "/" });
      } else if (key == 2 && this.$route.path != "/studentmanagement") {
        this.$router.push({ path: "/studentmanagement" });
      } else if (key == 3 && this.$route.path != "/announcementcenter") {
        this.$router.push({ path: "/announcementcenter" });
      } else if (key == 4 && this.$route.path != "/help") {
        this.$router.push({ path: "/help" });
      }
    },
    TextLink: function () {
      window.open("/");
    },
    onSuccess(times) {
      console.log("验证通过，耗时" + times + "毫秒");
      this.verificationCodeStatus =
        "login success, 耗时${(times / 1000).toFixed(1)}s";
    },
    onFail() {
      console.log("验证不通过");
      this.verificationCodeStatus = "";
    },
    onRefresh() {
      console.log("点击了刷新小图标");
      this.verificationCodeStatus = "";
    },
    onFulfilled() {
      console.log("刷新成功啦！");
    },
    onAgain() {
      console.log("检测到非人为操作的哦！");
      this.verificationCodeStatus = "try again";
      // 刷新
      this.$refs.slideblock.reset();
    },
    RefreshVerificationCode() {
      this.$refs.slideblock.reset();
    },
  },
};
</script>
<style>
.RegisterLogin {
  background-image: url("/image/RegisterLoginBackground.png") !important;
  background-size: cover !important;
}
</style>