
import {createApp, reactive} from 'vue'
import Main from './views/Main.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import SideBar_Menu from './components/SideBar_Menu.vue'


import './assets/dashboard.css'
import './assets/main.css'



var user = {
    id: null,
}

axios.interceptors.request.use(config =>{
	const token = localStorage.getItem('token');
	if(token){
		config.headers.Authorization = `Bearer ${token}`
	}
	return config;
})

const main_back = createApp(Main);
main_back.config.globalProperties.$axios = axios;
main_back.config.globalProperties.$current_user = reactive(user);
main_back.component("ErrorMsg", ErrorMsg);
main_back.component("LoadingSpinner", LoadingSpinner);
main_back.component("SideBar_Menu", SideBar_Menu);
main_back.use(router)
main_back.mount('#main')