<script setup>
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
    async refresh() {

      this.profile_username = this.$route.params.userid;
      console.log(this.profile_username)

      if (this.$current_user.id == null) {
        this.$router.push("/");
        return
      }

      if (this.$current_user.id != this.$route.params.userid) {
        this.my_profile = false;
      }

      //response = this.$axios.get("/users/" + this.$current_user.id + "/banned");


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
      this.$axios.put("/users/" + this.$current_user.id + "/followed/" + this.profile_username);
      this.isFollower = true;
      this.n_follower += 1;

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

    }
},


mounted() {
    this.refresh()
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
                        <i class="bi-person-dash-fill"></i>
                        Unfollow
                      </button>
                    </div>
                    <div v-else>
                      <button class="btn btn-primary btn-lg" type="button" @click="Follow()">
                        <i class="bi-person-plus-fill"></i>
                        Follow
                      </button>
                    </div>
                  </Transition>
                </div>
                <div class="col-md-6 mb-2">
                  <Transition name="fade" mode="out-in">
                    <div v-if="isBanned">
                      <button class="btn btn-success btn-lg" type="button" @click="UnBan()">
                        <i class="bi-person-check-fill"></i>
                        Unban
                      </button>
                    </div>
                    <div v-else>
                      <button class=" btn btn-danger btn-lg" type="button" @click="Ban()">
                        <i class="bi-person-x-fill"></i>
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
        
      </div>
    </div>
    <RouterView :key="$route.fullPath"></RouterView>
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