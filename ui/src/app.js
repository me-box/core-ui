import App from './App.vue';
import Vue from "./vueCore";
import routes from "./routes";
import Scan from "./pages/app/Scan";
import SensingKit from "./pages/app/SensingKit";
import VueRouter from "vue-router";

Vue.mixin({
	data: function() {
		return {
			isMobile: true,
			isDev: false
		}
	}
});

routes.push({path: '/scan', component: Scan});
routes.push({path: '/sensing', component: SensingKit});
const router = new VueRouter({
	routes: routes
});
new Vue({
	router,
	render: h => h(App),
	mounted() {
		document.addEventListener("pause", this.onPause, false);
		document.addEventListener("resume", this.onResume, false);
		this.$children[0].$watch('authenticated', () => {
			this.onResume();
		});
	},
	methods: {
		onPause() {
			window.SensingKit.stop();
		},
		onResume() {
			const app = this.$children[0];
			if(app.authenticated) {
				app.request('/driver-sensingkit/ui')
					.then(() => {
						const enabled_sensors = localStorage.getItem('sensors');
						const url = 'https://' + app.databoxUrl + '/driver-sensingkit';
						if (enabled_sensors) {
							window.SensingKit.startSensors(JSON.parse(enabled_sensors), url, () => {});
						}
					});
			}
		}
	}
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