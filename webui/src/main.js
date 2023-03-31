import {createApp, reactive} from 'vue'
import Login from './views/LoginView.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

import './assets/dashboard.css'
import './assets/main.css'

const login = createApp(Login)
login.config.globalProperties.$axios = axios;
login.component("ErrorMsg", ErrorMsg);
login.component("LoadingSpinner", LoadingSpinner);
login.use(router)
login.mount('#login')
