import Vue from 'vue'
import VueRouter from 'vue-router'
//allow computed properties to be fetched async
import AsyncComputed from 'vue-async-computed'

Vue.config.productionTip = false;

Vue.use(VueRouter);
Vue.use(AsyncComputed);

export default Vue;
