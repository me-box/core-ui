import App from './App.vue';
import Vue from "./vueCore";
import router from "./vueRouter";
import Scan from "./pages/Scan.vue";

router.addRoutes([
	{path: '/scan', component: Scan}
]);

Vue.mixin({
	data: function() {
		return {
			isMobile: true,
			isDev: true
		}
	}
});

new Vue({
	router,
	render: h => h(App)
}).$mount('#app');

window.addEventListener('message', (event) => {
	if (event.data.type === 'databox_oauth_redirect') {
		toolbar.showSpinner();
		SafariViewController.isAvailable((available) => {
			let url = event.data.url.replace('{callback}', 'databox://oauth');
			if (available) {
				SafariViewController.show({
						url: url,
						hidden: false,
						animated: false,
						enterReaderModeIfAvailable: false,
						tintColor: "#3f51b5",
					},
					(result) => {
						if (result.event === 'opened') {
						} else if (result.event === 'loaded') {
						} else if (result.event === 'closed') {
						}
					},
					(msg) => {
						console.log("KO: " + msg);
					})
			} else {
				window.open(url, '_blank', 'location=yes');
			}
		})
	}
});