
<script setup>
import { RouterLink, RouterView } from 'vue-router'
import SideBarMenu from '../components/SideBar_Menu.vue'
</script>
<script>
export default {
	components: {
        SideBarMenu
  	},
	methods:{
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
		
		mounted() {

			this.getMainstream(this.$current_user.id);
			console.log("es: " + this.$current_user.id);
		},

	}
}
</script>

<template>	
	<main>
		<SideBarMenu></SideBarMenu>

		<RouterView :key="$route.fullPath"></RouterView>
	</main>
	
</template>

<style>

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


