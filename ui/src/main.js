import Vue from 'vue'
import VueRouter from 'vue-router'

import App from './App.vue'
import AppData from './pages/AppData.vue'
import AppStore from './pages/AppStore.vue'
import AppUI from './pages/AppUI.vue'
import AppUIToolbar from './pages/AppUIToolbar.vue'
import DataSource from './pages/DataSource.vue'
import DataSources from './pages/DataSources.vue'
import Home from './pages/Home.vue'
import Install from './pages/Install.vue'
import Login from './pages/Login.vue'
import NotFound from './pages/404.vue'
import Settings from './pages/Settings.vue'
import SettingsSystem from './pages/SettingsSystem.vue'
import WStest from './pages/WStest.vue'
//allow computed properties to be fetched async
import AsyncComputed from 'vue-async-computed'

Vue.config.productionTip = false;

Vue.use(VueRouter);
Vue.use(AsyncComputed);

Vue.mixin({
	methods: {
		ApiGetRequest: function (url, cannedData) {
			let _this = this;
			return fetch(url, {
				credentials: "same-origin"
			})
				.then((response) => {
					return response.json()
				})
				.then(json => {
					return json;
				})
				.catch(() => {
					let devmode = localStorage.getItem('dev');
					if (devmode === "true") {
						return cannedData
					}
					else {
						// TODO If Authentication error
						localStorage.setItem('databoxAuthenticated', "false");
						_this.$parent.authenticated = false;
						_this.$router.push("/login")
					}
				});
		}
	}
});

const router = new VueRouter({
	//mode: 'history',
	base: '/core-ui/ui/',
	routes: [
		{path: '/', component: Home},
		{path: '/404', component: NotFound},
		{path: '/data', component: DataSources},
		{path: '/data/:app', component: AppData, props: true},
		{path: '/datasource/:datasource', component: DataSource, props: true},
		{path: '/login', component: Login},
		{path: '/store', component: AppStore},
		{path: '/settings', component: Settings},
		{path: '/settings/system', component: SettingsSystem},
		{path: '/install/:app', component: Install, props: true},
		{
			path: '/view/:app', components: {
				default: AppUI,
				toolbar: AppUIToolbar
			},
			props: {default: true, toolbar: false}
		},
		{path: '/wstest', component: WStest},
	]
});

new Vue({
	router,
	render: h => h(App)
}).$mount('#app');

//router.push('/');

//Add in the event listener to redirect oauth requests
window.addEventListener('message', (event) => {
	if (event.data.type === 'databox_oauth_redirect') {
		window.location.href = event.data.url;
	}
});