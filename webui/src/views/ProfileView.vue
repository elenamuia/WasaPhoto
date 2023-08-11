<script setup>
import { RouterLink, RouterView } from 'vue-router'
import SideBarMenu from '../components/SideBar_Menu.vue'
</script>
<script>
export default {
    components: {
        SideBarMenu
    },
    data: function() {
      return{
        n_follower:0,
        n_followed:0,
        n_posts:0,
        my_profile: true,
        follower: [],
        followed:[],
        post:[],

      }
  },

  methods:{
    refresh(){
      this.profile_username = this.$route.params.userid;
      console.log(this.profile_username)
      
      if (this.$current_user.id == null) {
                this.$router.push("/");
                return
            }
      if (this.$current_user.id == this.$route.params.username) {
        this.my_profile = true;
      }

      //this.$axios.get("/users/" + this.profile_username + "/follower/").then(response => {
        //this.follower = response.data; // Salva la risposta del backend nell'array 'data'
        //this.n_follower = response.data.length; // Ottieni la lunghezza dell'array
        //});
      
      
      //for(let i = 0; i < this.re)
      // Get followers 
      this.$axios.get("/users/" + this.$current_user.id + "/profile/"+ this.profile_username).then(response => {
        //this.followed = response.data; // Salva la risposta del backend nell'array 'data'
        this.n_followed = response.data.Followed.length; // Ottieni la lunghezza dell'array
        });
      
      this.$axios.get("/users/" + this.$current_user.id + "/profile/"+ this.profile_username).then(response => {
      //this.followed = response.data; // Salva la risposta del backend nell'array 'data'
      this.n_follower = response.data.Follower.length; // Ottieni la lunghezza dell'array
      });

      this.$axios.get("/users/" + this.$current_user.id + "/profile/"+ this.profile_username).then(response => {
      //this.followed = response.data; // Salva la risposta del backend nell'array 'data'
      this.n_posts = response.data.Photos.length; // Ottieni la lunghezza dell'array
      });
      
      
    }
  },
  mounted(){
    this.refresh()
  }
}

</script>

<template >
      

    <main>
      
      <SideBarMenu>
        
      </SideBarMenu>
      <div class = "container-fluid" style="margin-left: 400px; margin-top: 30px; height: max-content;">
        <div class="row">
          
          <div class="col col-lg-3" style="font-size: large; ">
            <b>{{this.profile_username}}</b>
          
          </div>
        
          <div class="col col-sm-1" style="font-size: medium;">
            Num Follower: 
              <div></div>
          </div>
          <div class="col col-sm-1" style="font-size: medium; font-style: 
          ;">
            {{n_follower}}
          </div>
          <div class="col col-sm-1" style="font-size: medium;">
            Num Followed:
          </div>
          <div class="col col-sm-1" style="font-size: medium;">
            {{n_followed}}
          </div>
          <div class="col col-sm-1" style="font-size: medium;">
            Num Post:
          </div>
          <div class="col col-sm-1" style="font-size: medium;">
            {{n_posts}}
          </div>
        
          
        </div>      

      </div>   
		    <RouterView />
        
    
    </main>
	  
    
</template>
<style>


</style>