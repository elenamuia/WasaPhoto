
<script setup>
import {RouterView } from 'vue-router'
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
			no_post:'',
			currentPage: 1,
      		photosPerPage: 5,
			photos: [],
		}
	},

	beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll);
  	},

	methods: {
		async initialize() {
			this.no_post = '';
			let response = await this.$axios.get("/users/" + this.$current_user.id + "/mainstream/",{
				params: {
					
					page: this.currentPage,
					perpage: this.photosPerPage
				}
          	});
			if (response.data != null){
				this.posts = response.data;
				for (const imageBytes of this.posts) {
					this.photos.push(imageBytes);
				}
				this.currentPage++;
			}
			
			if (this.photos.length == 0){
					this.no_post = 'There are no posts here yet, please come back later!';
				}
			
		},
		async delPost(post) {
			
			for (let i = 0; i < this.posts.length; i++) {
				if (this.posts[i].ID === post.ID) {
					this.posts.splice(i, 1);
					i--;
				}

			};
			this.n_posts -= 1;
		},
		async handleNewPhotoAdded() {
      		this.initialize();
   		},
		handleScroll() {
      		let nearBottom = window.innerHeight + window.scrollY >= document.documentElement.scrollHeight - 300;
      		if (nearBottom && !this.loading) {
      	  	this.initialize();
      		}
    	},	



	},
	mounted() {
		this.initialize();
		window.addEventListener('scroll', this.handleScroll);

	},
}
</script>

<template>
	<main>
		<SideBarMenu></SideBarMenu>
		<div style="margin-left: 450px; margin-top: 30px; font-size: large;">
			<strong> My Stream:</strong>
		</div>
		<div class="col" style="margin-left: 750px; margin-top: 30px;">
			{{no_post}}
			<Stream_Photo :posts="photos" @delete-post="delPost()"></Stream_Photo>
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


