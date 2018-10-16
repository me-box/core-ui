import Vue from 'vue'
import VueRouter from 'vue-router'
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
					if (!response.ok) {
						throw Error(response.statusText);
					}
					return response;
				})
				.then((response) => {
					return response.json()
				})
				.catch(() => {
					if (_this.isDev) {
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

export default Vue;

//router.push('/');