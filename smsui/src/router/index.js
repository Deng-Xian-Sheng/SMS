import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import StudentManagement from '../views/StudentManagement.vue'
import AnnouncementCenter from '../views/AnnouncementCenter.vue'
import Help from '../views/Help.vue'
import NotFoundComponent from '@/views/NotFoundComponent'

Vue.use(VueRouter)

const routes = [{
        path: '/',
        name: 'Home',
        components: {
            boss: Home,
        },
        meta: {
            // 页面标题title
            title: '首页',
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
        meta: {
            // 页面标题title
            title: '学生管理',
        },
    },
    {
        path: '/announcementcenter',
        name: 'AnnouncementCenter',
        components: {
            boss: AnnouncementCenter,
        },
        meta: {
            // 页面标题title
            title: '公告中心',
        },
    },
    {
        path: '/help',
        name: 'Help',
        components: {
            boss: Help,
        },
        meta: {
            // 页面标题title
            title: '帮助说明',
        },
    },
    {
        path: '*',
        name: 'NotFoundComponent',
        components: {
            boss: NotFoundComponent,
        },
        meta: {
            // 页面标题title
            title: '404',
        },
    },
]

const router = new VueRouter({
    mode: 'history',
    routes
})

router.beforeEach((to, from, next) => {
    /* 路由发生变化修改页面title */
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})

export default router