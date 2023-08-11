import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import MainstreamView from '../views/MainstreamView.vue'
import SettingsView from '../views/SettingsView.vue'
import ProfileView from '../views/ProfileView.vue'
//import SearchView from '../views/SearchModal.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/mainstream/:userid', component: MainstreamView},
		{path: '/settings/', component: SettingsView},
		{path: '/profile/:userid', component: ProfileView},
		//{path: '/search/', component: SearchView},
			
		
	]
})
export default router
