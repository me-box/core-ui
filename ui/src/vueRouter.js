import VueRouter from "vue-router";

import AppData from './pages/AppData.vue'
import AppStore from './pages/AppStore.vue'
import AppStoreToolbar from './pages/AppStoreToolbar.vue'
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

const router = new VueRouter({
	mode: 'history',
	base: '/core-ui/ui/',
	routes: [
		{path: '/', component: Home},
		{path: '/404', component: NotFound},
		{path: '/data', component: DataSources},
		{path: '/data/:app', component: AppData, props: true},
		{path: '/datasource/:datasource', component: DataSource, props: true},
		{path: '/login', component: Login},
		{path: '/store',  components: {
				default: AppStore,
				toolbar: AppStoreToolbar
			},
			props: {default: true, toolbar: false}
		},
		{path: '/settings', component: Settings},
		{path: '/settings/system', component: SettingsSystem},
		{path: '/install/:app', component: Install, props: true},
		{
			path: '/view/:app', components: {
				default: AppUI,
				toolbar: AppUIToolbar
			},
			props: {default: true, toolbar: true}
		},
		{
			path: '/view/:app/:path', components: {
				default: AppUI,
				toolbar: AppUIToolbar
			},
			props: {default: true, toolbar: true}
		}
	]
});

export default router;