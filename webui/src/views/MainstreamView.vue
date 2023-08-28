
<script setup>
import { RouterLink, RouterView } from 'vue-router'
import SideBarMenu from '../components/SideBar_Menu.vue'
import Stream_Photo from '../components/Stream_Photo.vue';
</script>
<script>
export default {
	components: {
		SideBarMenu,
		Stream_Photo
	},
	data: function () {
		return {
			posts: [],
			loading: false,
		}
	},

	methods: {
		async initialize() {

			let response = await this.$axios.get("/users/" + this.$current_user.id + "/mainstream/");
			this.posts = response.data;
			console.log(this.posts)
		},
		async delPost(post) {
			console.log("post: " + post)
			for (let i = 0; i < this.posts.length; i++) {
				if (this.posts[i].ID === post.ID) {
					this.posts.splice(i, 1);
					i--;
				}

			};
			this.n_posts -= 1;
		},



	},
	mounted() {
		this.initialize();

	},
}
</script>

<template>
	<main>
		<SideBarMenu></SideBarMenu>
		<div class="col" style="margin-left: 750px;">
			<Stream_Photo :posts="this.posts" @delete-post="delPost()"></Stream_Photo>
			<div>
				<div v-if="loading">
					<LoadingSpinner></LoadingSpinner>
				</div>
			</div>
		</div>
		<RouterView :key="$route.fullPath"></RouterView>
	</main>
</template>

<style></style>


