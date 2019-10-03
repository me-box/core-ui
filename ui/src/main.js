import App from './App.vue'
import Vue from "./vueCore";
import routes from "./routes"
import './registerServiceWorker'
import VueRouter from "vue-router";

Vue.mixin({
	data: function() {
		return {
			isMobile: false,
			isDev: false
		}
	},
});

window.getOauthCallbackURL = () => {
	const databoxUrl = localStorage.getItem("databoxURL");
	if(window.location.pathname.startsWith('/core-ui/ui/view/')) {
		const appname = window.location.pathname.split('/')[4];
		return 'https://' + databoxUrl + '/core-ui/ui/view/' + appname + '/oauth';
	}
	return 'https://' + databoxUrl + '/ui/oauth';
};
window.startOauth = (url) => {
	window.location.href = url;
};

const router = new VueRouter({
	mode: 'history',
	base: '/core-ui/ui/',
	routes: routes
});
new Vue({
	router,
	render: h => h(App)
}).$mount('#app');