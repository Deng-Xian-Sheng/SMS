import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import StudentManagement from '../views/StudentManagement.vue'
import AnnouncementCenter from '../views/AnnouncementCenter.vue'
import Help from '../views/Help.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    components: {
      boss: Home,
    },
  },
  {
    path: '/studentmanagement',
    name: 'StudentManagement',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    components: {
      boss: StudentManagement,
    },
  },
  {
    path: '/announcementcenter',
    name: 'AnnouncementCenter',
    components:{
      boss:AnnouncementCenter,
    },
  },
  {
    path:'/help',
    name:'Help',
    components:{
      boss:Help,
    },
  },
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
