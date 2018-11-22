import VueRouter from "vue-router";

const Login = () => import('./pages/Login.vue');
import NotFound from './pages/404.vue'

import Home from './pages/Home.vue';
import AppUI from './pages/AppUI.vue'
import AppUIToolbar from './pages/AppUIToolbar.vue'

const AppStore = () => import(/* webpackChunkName: "group-store" */ './pages/AppStore.vue');
const AppStoreToolbar = () => import(/* webpackChunkName: "group-store" */ './pages/AppStoreToolbar.vue');
const Install = () => import(/* webpackChunkName: "group-store" */ './pages/Install.vue');

const Settings = () => import(/* webpackChunkName: "group-settings" */ './pages/Settings.vue');
const SettingsSystem = () => import(/* webpackChunkName: "group-settings" */ './pages/SettingsSystem.vue');
const DataSource = () => import(/* webpackChunkName: "group-settings" */ './pages/DataSource.vue');
const DataSources = () => import(/* webpackChunkName: "group-settings" */ './pages/DataSources.vue');
const AppData = () => import(/* webpackChunkName: "group-settings" */ './pages/AppData.vue');

const router = new VueRouter({
	mode: 'history',
	base: '/core-ui/ui/',
	routes: [
		{path: '/', component: Home},
		{path: '/404', component: NotFound},
		{path: '/data', component: DataSources},
		{path: '/data/:app', component: AppData, props: true},
		{path: '/data/:app/:datasource', component: DataSource, props: true},
		{path: '/login', component: Login},
		{
			path: '/store', components: {
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
		},
		{path: '*', component: NotFound}
	]
});

export default router;