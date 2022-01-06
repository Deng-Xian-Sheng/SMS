import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import 'element-ui/lib/theme-chalk/display.css';
import router from './router'
import axios from 'axios';
import VueWechatTitle from 'vue-wechat-title'

Vue.prototype.axios = axios;
Vue.config.productionTip = true;

Vue.use(ElementUI);
Vue.use(VueWechatTitle)

new Vue({
    el: '#app',
    router,
    render: h => h(App)
});