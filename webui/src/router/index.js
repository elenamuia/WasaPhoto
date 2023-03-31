import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginView from '../views/LoginView.vue'
import MainstreamView from '../views/MainstreamView.vue'

Vue.use(VueRouter);


const routers = [
		{
			path: '/',
			name: 'login',
			component: LoginView
		},
		{
			path: '/mainstream', 
			name:'mainstream',
			component: MainstreamView
		}
];

const router = new VueRouter({
    routers
});

export default router
