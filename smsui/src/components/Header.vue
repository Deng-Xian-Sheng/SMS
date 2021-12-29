<template>
  <div>
    <el-row type="flex" class="row-bg" justify="center">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          <el-row>
            <a href="/" target="_blank">
              <el-col :span="2" :push="10">
                <div class="grid-content bg-purple">
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
            class="el-menu"
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
          <el-button type="text" @click="loginDialogVisible = true">注册/登录</el-button>
        </div>
      </el-col>
    </el-row>
    <el-dialog
      title="登录"
      :visible.sync="loginDialogVisible"
      width="30%"
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
            <el-input placeholder="请输入用户名&邮箱" v-model="inputName" clearable></el-input>
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
      <el-row>
        <el-col :span="4">
          <div class="grid-content bg-purple">
            <el-link :underline="false" icon="el-icon-refresh">换一换</el-link>
          </div>
        </el-col>
        <el-col :span="16">
          <div class="grid-content bg-purple">
            <el-input placeholder="请输入人机验证码" v-model="inputManMachineVerificationCode" clearable></el-input>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="grid-content bg-purple">
            <el-image
              style="width: 80px; height: 40px;margin-top: 8px;"
              :src="urlManMachineVerificationCode"
              :fit="fitManMachineVerificationCode"
            >
              <div slot="error" class="image-slot">
                <i class="el-icon-picture-outline"></i>
              </div>
            </el-image>
          </div>
        </el-col>
      </el-row>
      <el-row :gutter="0">
        <el-col :span="12" :offset="9">
          <div class="grid-content bg-purple">
            <el-button
              type="text"
              @click="loginDialogVisible = false;orgetPasswdDialogVisible=false;RegisterDialogVisible = true"
            >注册</el-button>
            <span>/</span>
            <el-button
              type="text"
              @click="loginDialogVisible = false;RegisterDialogVisible = false;forgetPasswdDialogVisible=true;forgetPasswdTips()"
            >忘记密码</el-button>
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
      width="30%"
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
      <el-row>
        <el-col :span="4">
          <div class="grid-content bg-purple">邮箱：</div>
        </el-col>
        <el-col :span="20">
          <div class="grid-content bg-purple">
            <el-input
              placeholder="请输入邮箱"
              v-model="inputMailbox"
              clearable
              suffix-icon="el-icon-message"
            ></el-input>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="4">
          <div class="grid-content bg-purple">
            <el-link :underline="false" icon="el-icon-refresh">换一换</el-link>
          </div>
        </el-col>
        <el-col :span="16">
          <div class="grid-content bg-purple">
            <el-input placeholder="请输入人机验证码" v-model="inputManMachineVerificationCode" clearable></el-input>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="grid-content bg-purple">
            <el-image
              style="width: 80px; height: 40px;margin-top: 8px;"
              :src="urlManMachineVerificationCode"
              :fit="fitManMachineVerificationCode"
            >
              <div slot="error" class="image-slot">
                <i class="el-icon-picture-outline"></i>
              </div>
            </el-image>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="6">
          <div class="grid-content bg-purple">邮箱验证码：</div>
        </el-col>
        <el-col :span="14">
          <div class="grid-content bg-purple">
            <el-input placeholder="请输入邮箱验证码" v-model="inputMailboxVerification" clearable></el-input>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="grid-content bg-purple">
            <el-button size="small" round>获取</el-button>
          </div>
        </el-col>
      </el-row>
      <el-row :gutter="0">
        <el-col :span="12" :offset="9">
          <div class="grid-content bg-purple">
            <el-button
              type="text"
              @click="RegisterDialogVisible = false;forgetPasswdDialogVisible=false;loginDialogVisible = true"
            >登录</el-button>
            <span>/</span>
            <el-button
              type="text"
              @click="loginDialogVisible = false;RegisterDialogVisible = false;forgetPasswdDialogVisible=true;forgetPasswdTips()"
            >忘记密码</el-button>
          </div>
        </el-col>
      </el-row>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary">马上注册</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="忘记密码"
      :visible.sync="forgetPasswdDialogVisible"
      width="30%"
      center
      :destroy-on-close="true"
      custom-class="RegisterLogin"
    >
      <el-row>
        <el-col :span="4">
          <div class="grid-content bg-purple">邮箱：</div>
        </el-col>
        <el-col :span="20">
          <div class="grid-content bg-purple">
            <el-input
              placeholder="请输入邮箱"
              v-model="inputForgetPasswdMailbox"
              clearable
              suffix-icon="el-icon-message"
            ></el-input>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="4">
          <div class="grid-content bg-purple">
            <el-link :underline="false" icon="el-icon-refresh">换一换</el-link>
          </div>
        </el-col>
        <el-col :span="16">
          <div class="grid-content bg-purple">
            <el-input placeholder="请输入人机验证码" v-model="inputManMachineVerificationCode" clearable></el-input>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="grid-content bg-purple">
            <el-image
              style="width: 80px; height: 40px;margin-top: 8px;"
              :src="urlManMachineVerificationCode"
              :fit="fitManMachineVerificationCode"
            >
              <div slot="error" class="image-slot">
                <i class="el-icon-picture-outline"></i>
              </div>
            </el-image>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="6">
          <div class="grid-content bg-purple">邮箱验证码：</div>
        </el-col>
        <el-col :span="14">
          <div class="grid-content bg-purple">
            <el-input
              placeholder="请输入邮箱验证码"
              v-model="inputForgetPasswdMailboxVerification"
              clearable
            ></el-input>
          </div>
        </el-col>
        <el-col :span="4">
          <div class="grid-content bg-purple">
            <el-button size="small" round>获取</el-button>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="4">
          <div class="grid-content bg-purple">新密码：</div>
        </el-col>
        <el-col :span="20">
          <div class="grid-content bg-purple-light">
            <el-input placeholder="请输入新密码" v-model="inputNewPasswd" show-password></el-input>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="8">
          <div class="grid-content bg-purple">再次输入新密码：</div>
        </el-col>
        <el-col :span="16">
          <div class="grid-content bg-purple-light">
            <el-input placeholder="请再次输入新密码" v-model="inputNewPasswd2" show-password></el-input>
          </div>
        </el-col>
      </el-row>
      <el-row :gutter="0">
        <el-col :span="12" :offset="9">
          <div class="grid-content bg-purple">
            <el-button
              type="text"
              @click="loginDialogVisible = false;forgetPasswdDialogVisible=false;RegisterDialogVisible = true"
            >注册</el-button>
            <span>/</span>
            <el-button
              type="text"
              @click="RegisterDialogVisible = false;forgetPasswdDialogVisible=false;loginDialogVisible = true"
            >登录</el-button>
          </div>
        </el-col>
      </el-row>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary">确认更新密码</el-button>
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
      loginDialogVisible: false,
      inputName: "",
      inputPasswd: "",
      RegisterDialogVisible: false,
      inputNameRegister: "",
      inputPasswdRegister: "",
      inputMailbox: "",
      inputMailboxVerification: "",
      inputManMachineVerificationCode: "",
      urlManMachineVerificationCode: "/favicon.png",
      fitManMachineVerificationCode: "cover",
      forgetPasswdDialogVisible: false,
      inputForgetPasswdMailbox: "",
      inputForgetPasswdMailboxVerification: "",
      inputNewPasswd: "",
      inputNewPasswd2: "",
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
    forgetPasswdTips() {
      this.$alert(
        "“忘记密码”功能仅支持“校长”角色，如您是其他角色请联系上级修改密码！",
        "⚠️注意！",
        {
          confirmButtonText: "确定",
        }
      );
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