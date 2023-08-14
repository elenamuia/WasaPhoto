<script setup>
import { RouterLink, RouterView } from 'vue-router'
import SideBarMenu from '../components/SideBar_Menu.vue'
</script>
<script>
export default {
  components: {
    SideBarMenu
  },
  data: function () {
    return {
      isFollower: false,
      isBanned: false,
      has_banned_me: false,
      n_follower: 0,
      n_followed: 0,
      n_posts: 0,
      my_profile: true,
      follower: [],
      followed: [],
      post: [],
      profile_username: null,

    }
  },

  methods: {
    refresh() {

      this.profile_username = this.$route.params.userid;
      console.log(this.profile_username)

      if (this.$current_user.id == null) {
        this.$router.push("/");
        return
      }

      if (this.$current_user.id != this.$route.params.userid) {
        this.my_profile = false;
      }

      let response =  this.$axios.get("/users/" + this.username + "/bans");
      this.has_banned_you = response.data.users.map(x => x["username-string"]).includes(this.$current_user.id);



      this.$axios.get("/users/" + this.$current_user.id + "/profile/" + this.profile_username).then(response => {
        this.n_followed = response.data.Followed.length;
      });

      this.$axios.get("/users/" + this.$current_user.id + "/profile/" + this.profile_username).then(response => {
        this.n_follower = response.data.Follower.length;
      });

      this.$axios.get("/users/" + this.$current_user.id + "/profile/" + this.profile_username).then(response => {
        this.n_posts = response.data.Photos.length;
      });
    },

    async Follow() {
      console.log("before following: " + this.isFollower)
      this.$axios.put("/users/" + this.$current_user.id + "/followed/" + this.profile_username);
      this.isFollower = true;
      this.n_follower += 1;
      console.log("after following: " + this.isFollower)
    },

    async Unfollow() {
      this.$axios.delete("/users/" + this.$current_user.id + "/followed/" + this.profile_username);
      this.isFollower = false;
      this.n_follower -= 1;
    },

    async Ban(){
      this.$axios.put("/users/" + this.$current_user.id + "/banned/" + this.profile_username);
      this.isBanned = true;

    },

    async Unban(){
      this.$axios.delete("/users/" + this.$current_user.id + "/banned/" + this.profile_username);
      this.isBanned = false;

    },

  

},





mounted() {
  this.refresh();

},

}

</script>

<template >
  <main>

    <SideBarMenu>

    </SideBarMenu>
    <div class="container-fluid" style="margin-left: 400px; margin-top: 30px; height: max-content;">
      <div class="row">

        <div class="col col-lg-3" style="font-size: large; ">
          <div class="profile-info">
            <b style="margin-bottom: 10px">{{ this.profile_username }}</b>

            <div v-if="has_banned_me">
              Unavailable information
            </div>

            <div v-else-if="my_profile">
              <button class="btn btn-primary btn-md" type="button" @click="changeUsername()">
                Change Name
              </button>

            </div>

            <div v-else>

              <!-- Follow Button -->
              <div class="row">
                <div class="col-md-6 mb-2">
                  <Transition name="fade" mode="out-in">
                    <div v-if="isFollower">
                      <button class="btn btn-warning btn-lg" type="button" @click="Unfollow()">
                        
                        Unfollow
                      </button>
                    </div>
                    <div v-else>
                      <button class="btn btn-primary btn-lg" type="button" @click="Follow()">
                        
                        Follow
                      </button>
                    </div>
                  </Transition>
                </div>
                <div class="col-md-6 mb-2">
                  <Transition name="fade" mode="out-in">
                    <div v-if="isBanned">
                      <button class="btn btn-success btn-lg" type="button" @click="Unban()">
                        
                        Unban
                      </button>
                    </div>
                    <div v-else>
                      <button class=" btn btn-danger btn-lg" type="button" @click="Ban()">
                        
                        Ban
                      </button>
                    </div>
                  </Transition>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="col col-sm-1" style="font-size: medium;">
          Num Follower:
        </div>
        <div class="col col-sm-1" style="font-size: medium;">
          {{ this.n_follower }}
        </div>

        <div class="col col-sm-1" style="font-size: medium;">
          Num Followed:
        </div>
        <div class="col col-sm-1" style="font-size: medium;">
          {{ this.n_followed }}
        </div>

        <div class="col col-sm-1" style="font-size: medium;">
          Num Post:
        </div>
        <div class="col col-sm-1" style="font-size: medium;">
          {{ this.n_posts }}
        </div>



        <RouterView />
      </div>
    </div>

  </main>
</template>
<style scoped>
.profile-info {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.fade-enter-active,
.fade-leave-active {
    transition: opacity cubic-bezier(0.4, 0, 0.2, 1) 0.1s
}

.fade-enter,
.fade-leave-to {
    opacity: 0
}
</style>