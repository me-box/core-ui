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
import WStest from './pages/WStest.vue'

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
    { path: '/wstest', component: WStest },
  ]
})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

Vue.mixin({
  methods: {
    ApiGetRequest: function (url, cannedData) {
      let _this = this
      return fetch(url, {
        credentials: "same-origin"
      })
      .then((response) => {
            return response.json()
      })
      .then(json => {
        return json;
      })
      .catch(()=>{
          let devmode = localStorage.getItem('dev')
          if (devmode == "true") {
            return cannedData
          } else {
            localStorage.setItem('databoxAuthenticated', "false")
            _this.$parent.authenticated = "false"
            _this.$router.push("/login")
          }
      });
    }
  }
})

//Add in the event listener to redirect oauth requests
window.addEventListener('message', (event) => {
	if (event.data.type === 'databox_oauth_redirect') {
		window.location.href = event.data.url;
	}
});