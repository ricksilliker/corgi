import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';
import Dashboard from "@/pages/Dashboard";
import VueRouter from 'vue-router';

// configure router
Vue.use(VueRouter);
const routes = [
	{ path: "/", component: Dashboard }
]
const router = new VueRouter({
	mode: "abstract",
	linkActiveClass: "active",
	routes,
});

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Wails.Init(() => {
	new Vue({
		router,
		render: h => h(App)
	}).$mount('#app');
});
