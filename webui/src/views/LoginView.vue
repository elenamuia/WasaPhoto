<script setup>
import { axios } from 'axios';
</script>

<script>

export default {
	data: function() {
		return {
			errormsg: null
		}
	},
	methods: {
		async login() {

			try {
				let response = await this.$axios.post("/login/", {
					"LoginName": this.userid
				});

				this.errormsg = null;
                this.$current_user.id = this.userid
                this.$router.push('/mainstream/' + this.userid);
			} catch (err) {
				this.errormsg = err.message;
			}
 
			
		},
	}
}
</script>

<template >
	<main>

	

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow" >
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WasaPhoto</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

	<div style = "height: 300px; display: flex; justify-content: center; align-items: center; " >
		
		<div class="col-md-2">	
			<form @submit.prevent="login()">
				<div style="padding:5px;">
					<input type="text" v-model="userid" placeholder="Username" style="border-color: #650C96;"/>
				</div>
				<div class="btn-toolbar mb-2 mb-md-0">
					<div class="btn-group me-2" style ="padding:5px;">
						<button type="submit" class="btn btn-sm btn-outline-secondary" style=" background-color: #B173BE; color:antiquewhite; border-color:#650C96;">
							Login
						</button>	
					</div>
				</div>
				
			</form>
		</div>
	</div>
	<Routerview />
</main>
</template>
<style>
</style>



