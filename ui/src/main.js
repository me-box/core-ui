import App from './App.vue'
import Vue from "./vueCore";
import router from "./vueRouter"

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

new Vue({
	router,
	render: h => h(App)
}).$mount('#app');