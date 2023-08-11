
  
  <script>
  import axios from 'axios';
  
  export default {
    props: ['showModal'],
    data() {
      return {
        searchQuery: '',
        searchResults: ''
      };
    },
    methods: {
      searchProfiles() {
        console.log("searcheduser: " + this.searchQuery)
        console.log("userid: " + this.$current_user.id)

        var url = '/users/:userid/profile/:searcheduser';
        url.replace(':userid', this.$current_user.id);
        url.replace(':searcheduser', this.searchQuery);
        
        // Effettua la chiamata API per cercare i profili
        try {
				let response =  this.$axios.get(url, {});

				this.errormsg = null;
        this.$router.push('/profile/' + this.searchQuery);
			} catch (err) {
				this.errormsg = err.message;
			}
    },  
     
      closeModal() {
        this.$emit('close');
      }
    }
  }
  </script>



  <template>
    <div class="search-modal" v-if="showModal">
      <!-- Campi di ricerca -->
      <div class="search-modal-content">
        <input type="text" v-model="searchQuery" placeholder="Who are you looking for?">
          <!-- Bottone per eseguire la ricerca -->
          <button @click="searchProfiles">Search</button>
          
          <!-- Risultati della ricerca -->
          <div v-if="searchResults.length > 0">
            <h2>Risultati della ricerca:</h2>
            <ul>
              <li v-text="searchResults">
              </li>
            </ul>
          </div>
          
        <!-- Bottone per chiudere la modale -->
        <button @click= "closeModal">Close</button>
      </div>  
    </div>
  </template>
  
  <style>
  /* Stili per la modale di ricerca */
  .search-modal {
    position: fixed;
    top: 0;
    left:0;
    height:100%;
    width:100%;
    background-color: rgba(0,0,0,0.5);
    display:flex;
    justify-content: center;
    align-items: center;
    z-index: 9999;
    /* Stili per il posizionamento e overlay */
  }
  
  .search-modal-content{
    background-color: white;
    padding:20px;
    border-radius: 5px;

  }
  /* Stili per i campi di ricerca e risultati */
  </style>