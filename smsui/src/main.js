import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import 'element-ui/lib/theme-chalk/display.css';
import router from './router'
import SlideVerify from 'vue-monoplasty-slide-verify';
import axios from 'axios';
import { Message } from 'element-ui';

Vue.prototype.axios = axios;
Vue.config.productionTip = true;
Vue.prototype.Message=Message

Vue.use(ElementUI);
Vue.use(SlideVerify);

new Vue({
  el: '#app',
  router,
  render: h => h(App)
});
