<script setup>
import {  RouterView } from 'vue-router'
</script>
<script>
export default {


	data: function () {
		return {
			errormsg: null,
			loading: false,
		}
	},
	methods: {
        async refresh() {
			if (this.$current_user.id == null) {
				console.log("Empty username, redirecting to login")
				this.$router.push("/");
			}
        }
		

    },

    mounted() {
		this.refresh()

    }
}

</script>

<template>
    <main >
		<RouterView :key="$route.fullPath"></RouterView>
    </main>
</template>


<style>
</style>