<script>
export default {
	data: function() {
		return {
			errormsg: null,
			userid: "",
		}
	},
	methods: {
		async login() {
			
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/login/", {
					LoginName: this.userid
				});
				this.userid = response.data;
				this.$router.push('/mainstream')
			} catch (e) {
				this.errormsg = e.toString();
			}

			
		},
	}
}
</script>

<template >
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	<div style = "height: 300px; display: flex; justify-content: center; align-items: center;" >
		
		<div class="col-md-4">	
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
</template>

<style>
</style>



