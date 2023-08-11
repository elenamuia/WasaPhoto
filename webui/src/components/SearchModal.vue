
  
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

        var url = '/users/:userid/profile/:searcheduser';
        url.replace(':userid', this.searchQuery);
        url.replace(':searcheduser', 'Elena');
        
        // Effettua la chiamata API per cercare i profili
        axios.get(url, {})
        .then(response => {
          this.searchResults = response.data.User;
        })
        .catch(error => {
          console.error('Errore durante la ricerca:', error);
        });
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
        <button @click= "showmodal =false">Close</button>
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