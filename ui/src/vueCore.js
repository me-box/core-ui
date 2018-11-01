import Vue from 'vue'
import VueRouter from 'vue-router'
//allow computed properties to be fetched async
import AsyncComputed from 'vue-async-computed'

Vue.config.productionTip = false;

Vue.use(VueRouter);
Vue.use(AsyncComputed);

Vue.mixin({
	created() {
		let authenticated = localStorage.getItem("databoxAuthenticated") === "true";
		if (!authenticated) {
			console.log("login2");
			this.$router.replace("/login");
		}
	}
});

export default Vue;

//router.push('/');