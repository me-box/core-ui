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

//Add in the event listener to redirect oauth requests
window.addEventListener('message', (event) => {
	if (event.data.type === 'databox_oauth_redirect') {
		window.location.href = event.data.url;
	}
});

const router = new VueRouter({
	mode: 'history',
	base: '/core-ui/ui/',
	routes: routes
});
new Vue({
	router,
	render: h => h(App)
}).$mount('#app');