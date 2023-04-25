<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {


	data: function () {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			photos: []
		}
	},
	methods: {
		getMainstream(userid) {
			axios.post('/get-my-mainstream/' + userid, {}, { responseType: 'arraybuffer' })
				.then(response => {
					// Trasforma l'array di byte in un oggetto Blob
					const blob = new Blob([response.data], { type: 'image/jpeg' }); // Modifica il tipo MIME a seconda del tipo di immagine restituito

					// Crea un URL di dati (data URL) dall'oggetto Blob
					const imageUrl = URL.createObjectURL(blob);

					// Assegna l'URL di dati all'array di immagini nel componente
					this.photos.push(imageUrl);
				})
				.catch(error => {
					console.error('Errore durante il caricamento delle immagini:', error);
				});
		}
	}
	,

		mounted() {
			this.getMainstream(this.$user.id);
		}

	}
</script>

<template>
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WasaPhoto</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse"
			data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6
						class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Home</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/mainstream/" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#home" />
								</svg>
								Home
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/my_profile" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#layout" />
								</svg>
								Search Profile
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/search/" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#key" />
								</svg>
								My Profile
							</RouterLink>
						</li>
					</ul>

					<h6
						class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Secondary menu</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink :to="'/settings/'" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#file-text" />
								</svg>
								Settings
							</RouterLink>
						</li>
						<li class="nav-item">
							New Photo
							<div>
    							<img v-bind:src="photoUrl" />
  							</div>
						</li>
					</ul>
				</div>
			</nav>

			<main>
				<RouterView />

				<div>
					<h1>Photos</h1>
					<div v-for="photo in photos" :key="photo">
    					<img :src="photo" alt="photo">
    				</div>
				</div>
			</main>
	</div>
</div></template>

<style></style>


