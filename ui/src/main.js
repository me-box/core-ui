import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.config.productionTip = false

Vue.use(VueRouter)

import App from './App.vue'
import Home from './pages/Home.vue'
import AppStore from './pages/AppStore.vue'
import NotFound from './pages/404.vue'
import Settings from './pages/Settings.vue'
import MyApps from './pages/MyApps.vue'
import DataSources from './pages/DataSources.vue'
import Install from './pages/Install.vue'
import ViewAppUI from './pages/ViewAppUI.vue'
import Login from './pages/Login.vue'

//FontAwesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { faStar, faEye, faTimesCircle, faSyncAlt } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
library.add(faStar,faEye,faTimesCircle,faSyncAlt)
Vue.component('font-awesome-icon', FontAwesomeIcon)

//allow computed properties to be fetched async
import AsyncComputed from 'vue-async-computed'
Vue.use(AsyncComputed)

const router = new VueRouter({
  mode: 'history',
  base: '/core-ui/ui/',
  routes: [
    { path: '/', component: Home },
    { path: '/AppStore', component: AppStore },
    { path: '/404', component: NotFound },
    { path: '/Settings', component: Settings },
    { path: '/MyApps', component: MyApps },
    { path: '/DataSources', component: DataSources },
    { path: '/Install', component: Install },
    { path: '/view', component: ViewAppUI },
    { path: '/login', component: Login },
  ]
})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')