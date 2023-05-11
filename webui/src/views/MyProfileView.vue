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
      }
  },

}

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