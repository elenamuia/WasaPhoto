import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import MainstreamView from '../views/MainstreamView.vue'

import ProfileView from '../views/ProfileView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/mainstream/:userid', component: MainstreamView},
		{path: '/profile/:userid', component: ProfileView},		
	]
})
export default router
