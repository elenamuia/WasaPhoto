<script setup>
import SideBarMenu from '../components/SideBar_Menu.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
import Stream_Photo from '../components/Stream_Photo.vue';
</script>
<script>
export default {
  components: {
    SideBarMenu,
    LoadingSpinner,
    Stream_Photo
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
      photos: [],
      currentPage: 1,
      photosPerPage: 5,
      profile_username: null,
      get_banned: [],
      loading: false,
      photoPartial: [],
      posts: [],
    }
  },


  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll);
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

      if (!this.my_profile) {
        await this.$axios.get("/users/" + this.$current_user.id + "/bannedby/" + this.profile_username).then(response => {
          this.has_banned_me = response.data
        });
        console.log(this.has_banned_me)
        await this.$axios.get("/users/" + this.$current_user.id + "/banned/" + this.profile_username).then(response => {
          this.isBanned = response.data
        });
        await this.$axios.get("/users/" + this.$current_user.id + "/followed/" + this.profile_username).then(response => {
          this.isFollower = response.data
        });

      }


      this.$axios.get("/users/" + this.$current_user.id + "/profile/" + this.profile_username).then(response => {
        if (response.data.Followed == null) {
          this.n_followed = 0;
        } else {
          this.n_followed = response.data.Followed.length;
        }
        if (response.data.Follower == null) {
          this.n_follower = 0;
        } else {
          this.n_follower = response.data.Follower.length;
        }
        if (response.data.Photos == null) {
          this.n_posts = 0;
        } else {
          this.n_posts = response.data.Photos.length;
        }
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

    async Ban() {
      this.$axios.put("/users/" + this.$current_user.id + "/banned/" + this.profile_username);
      this.isBanned = true;

    },

    async Unban() {
      this.$axios.delete("/users/" + this.$current_user.id + "/banned/" + this.profile_username);
      this.isBanned = false;

    },

    async bytesToBase64(bytes) {

      const binary = String.fromCharCode(...bytes);
      const binarystring = window.btoa(binary);
      return binarystring;
    },

    async handleNewPhotoAdded() {
      this.loadPhotos();

    },

    async loadPhotos(page) {
      this.loading = true;
      try {
        let response = await this.$axios.get('/users/' + this.$current_user.id + '/photos/', {
          params: {
            profile: this.profile_username,
            page: page,
            perpage: this.photosPerPage
          }
        });
        if (response.data == null){
          return
        }

        
        this.posts = response.data;
        for (const imageBytes of this.posts) {
          this.photos.push(imageBytes);
          
        }


        //var dataArray = response.data.map(item => item.PhotoStructure);
        //for (const imageBytes of dataArray) {
        //  this.photos.push(imageBytes);
        //}


        this.currentPage++;
        this.loading = false;
      } catch (error) {
        this.loading = false;
        console.error("Errore nel recuperare le foto:", error);
      }
    },
    handleScroll() {
      let nearBottom = window.innerHeight + window.scrollY >= document.documentElement.scrollHeight - 300;
      if (nearBottom && !this.loading) {
        this.loadPhotos(this.currentPage);
      }
    }


  },


  mounted() {
    this.refresh();
    this.loadPhotos(this.currentPage);
    window.addEventListener('scroll', this.handleScroll);
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
              No info available, you have been banned by this user!
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
                    <div v-if="this.isFollower">
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
                    <div v-if="this.isBanned">
                      <button class="btn btn-success btn-lg" type="button" @click="Unban()">
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

        <div v-if="!has_banned_me" class="row">

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
      <div>
        <Stream_Photo :posts="this.photos"></Stream_Photo>
        <div>
          <div v-if="loading">
            <LoadingSpinner></LoadingSpinner>
          </div>
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

.uploaded-photos {
  display: flex;

  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
}

.uploaded-photo {
  max-width: 500px;
  max-height: 500px;
  margin: 30px;
  object-fit: cover;
}
</style>