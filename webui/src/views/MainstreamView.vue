
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
			this.$axios.get("/users/"+ userid+"/mainstream/" , { responseType: 'arraybuffer' })
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
		},

		openForm() {
			document.getElementById("myForm").style.display = "block";
		},

		closeForm() {
			document.getElementById("myForm").style.display = "none";
		},

		postPhoto(userid){
			
			// disable the Upload button while the photo is uploaded to avoid user-made mistakes
			document.getElementById("submitBut").classList.add("disabled");
			// take the first element received from the fileInput (defined in the html as the form-container)
			const image = document.getElementById("fileInput").files[0];
			
			// define a new Reader object to read the URL and to turn it into 
			let reader = new FileReader();

			reader.readAsDataURL(image);
			reader.onload = (e) => {
				const fileContent = e.target.result;
				console.log(fileContent); // qui puoi usare la stringa URL generata dal file
			}
			
			let response =  this.$axios.post("/users/"+ userid +"/photos/");
			this.errormsg = null; 
			
		
			if (response.status == 204) {
                // manually restyle and rename the submit button
                const submit_button = document.getElementById("submitBut");
                submit_button.classList.remove("btn");
                submit_button.classList.remove("disabled");
                
                setTimeout(() => {
                    submit_button.innerHTML = "Submit";
                }, 3000);
                // Clear the form
                document.getElementById("fileInput").value = "";
                
            }
            else {
                alert("Error uploading photo")
            }
        
			}
 
		},

		Logout() {
			this.$current_user.id = null;
			this.$router.push("/");
		},
	

	mounted() {
		this.getMainstream(this.$current_user.id);
		console.log("es: " + this.$current_user.id);
	}

}
</script>

<template>
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow" style = "width:100%;">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WasaPhoto</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse"
			data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse" style= "margin-top:55px;">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6
						class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Home</span>
					</h6>
					<ul class="nav flex-column" >
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
									<use href="/feather-sprite-v4.29.0.svg#search" />
								</svg>
								Search Profile
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/search/" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#user" />
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
							<a class="nav-link"></a>
							<RouterLink :to="'/settings/'" class="nav-link">
								<svg class="feather">
									<use href="/feather-sprite-v4.29.0.svg#settings" />
								</svg>
								Settings
							</RouterLink>

						</li>

						<li class="nav-item" @click="openForm()">
				
							<div class="nav-link">
							<svg class="feather">
								<use href="/feather-sprite-v4.29.0.svg#plus" />
							</svg>
							New Photo
						</div>
						</li>
						<li class="nav-item" @click="Logout()">
				
							<div class="nav-link">
							<svg class="feather">
								<use href="/feather-sprite-v4.29.0.svg#log-out" />
							</svg>
							Logout
						</div>
						</li>
					</ul> 
				</div>
			</nav>

			<main>
				<RouterView />

				<div>
					
					<div v-for="photo in photos" :key="photo">
						<img :src="photo" alt="Photo not loaded">
					</div>
				</div>

				<div class="form-popup" id="myForm">
					<form @submit.prevent="postPhoto(this.$current_user.id)" class="form-container">
							<h3 style = "margin-left: 10px;">Upload Photo</h3>
							<input  class="form-control" type="file" id="fileInput" accept= "image/png, image/jpeg">
							<button type="submit" class="btn" id="submitBut">Upload</button>
							<button type="button" class="btn cancel" @click="closeForm()">Close</button>
					</form>
				</div>
			</main>
		</div>
	</div>
</template>

<style>

/* The popup form - hidden by default */
.form-popup {
	background-color:lightsteelblue;
	display: none;
	position: center;
	border: 3px solid black;
	z-index: 9;
	margin-top: 50px;

}

.form-control {
	margin-left: 10px;
	margin-top: 20px;
}

.form-container .btn {
  background-color:darkolivegreen;
  color: white;
  margin-top: 20px;
  margin-bottom: 20px;
  margin-right: 20px;
  margin-left: 10px;
  

}




.form-container .cancel {
	background-color:darkred;
	color: white;
	
}

/* Add some hover effects to buttons */
.form-container .btn:hover,
.open-button:hover {
	opacity: 1;
}
</style>


