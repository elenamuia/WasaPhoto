<script setup>
import { RouterLink, RouterView } from 'vue-router'
import SideBarMenu from '../components/SideBar_Menu.vue'
</script>
<script>
export default {
    components: {
        SideBarMenu
  },
  data(){
    followerList = null,
    followedList = null,
    n_post = null
  },


mounted() {
    axios.get('/api/valore-dal-backend')
      .then(response => {
        this.valoreDallaChiamataAPI = response.data;
      });
  }
};
</script>

<template >
      

    <main>
      
      <SideBarMenu>
        
      </SideBarMenu>
      <div class = "container-fluid" style="margin-left: 400px; margin-top: 30px; height: max-content;">
       
        <div class="row">
          <div class="col col-lg-3" style="font-size: large;">
            {{this.$current_user.id}}
          </div>
          <div class="navbar col-sm-2 ">
            <ul class="nav" style="font-size: large;">
              <li class="nav-item active">
                

                <a class="nav-link" href="#">Follower </a>
              </li>
              <div v-if="followedList">{{followedList}}</div>
              <li class="nav-item active">
                <a class="nav-link" href="#">Followed </a>
              </li>
              <div v-if="followerList">{{followerList}}</div>
              <li class="nav-item active">
                <a class="nav-link" href="#">Num Post </a>
              </li>
              <div v-if="n_post">{{n_post}}</div>
            </ul>
          </div>
        </div>
        

            

      </div>   
		    <RouterView />
        
    
    </main>
	  
    
</template>
<style>


</style>