import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/pages/Main'
import Settings from '@/pages/Settings'
import Landing from '@/pages/Landing'
import SystemLog from '@/pages/SystemLog'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Main',
      component: Main,
      children: [
        {
          path: '/settings',
          name: 'Settings',
          component: Settings
        },
        {
          path: '/home',
          name: 'Home',
          component: Landing
        },
        {
          path: '/log',
          name: 'Log',
          component: SystemLog
        }
      ]
    }
  ]
})
