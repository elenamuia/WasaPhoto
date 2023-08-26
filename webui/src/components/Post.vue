<script>
export default {

    props: {
        post: {
            type: Object
        }
    },

    data: function () {
        return {
            photo_id: null,
            liked: false,
            username: '',
            num_likes: 0,
            comments: [],
            
        }
    },

    methods: {

        async initialize() {
            this.username = this.post.User;
            console.log("username is " + this.username);
            this.photo_id = this.post.ID
            console.log("PhotoID is " + this.photo_id);
            //this.is_your_post = this.post_data.author_name["username-string"] == this.$user_state.username;

            // Fetch likes
            let response = await this.$axios.get("/users/" + this.username + "/photos/" + this.photo_id + "/listlike/");
            if (response.data == null) {
                this.num_likes = 0;
                this.liked = false;
                
            } else {
                
                this.num_likes = response.data.length;

            } 

            response = await this.$axios.get("/users/" + this.username + "/photos/" + this.photo_id + "/like/"+this.$current_user.id);
            this.liked = response.data;
            console.log("liked: " + this.liked)
            response = await this.$axios.get("/users/" + this.username + "/photos/" + this.photo_id + "/listcomment/");
            
        },

        toggleLike() {
            this.liked = !this.liked;
            if (!this.liked){
                this.Unlike();
            }else{
                this.Like();
            }  
        },

        async Like() {
            this.$axios.put("/users/" + this.username + "/photos/" + this.photo_id + "/like/" + this.$current_user.id);
            this.liked = true;
            this.num_likes += 1;

        },

        async Unlike() {
            this.$axios.delete("/users/" + this.username + "/photos/" + this.photo_id + "/like/" + this.$current_user.id);
            this.liked = false;
            this.num_likes -= 1;

        },
    },

    mounted() {

        this.initialize();
    }
}
</script>

<template>
    <div style="max-width: 500px;">

        <div class="rounded border shadow-lg">

            <div class="row align-content-between my-2">

                <div class="col">
                    <i class="bi-person-circle mx-1" style="font-size: 2em"></i>
                    <span class="col font-weight-bold h4">
                        {{ this.username }}
                    </span>
                </div>
            </div>

            <div class="row">
                <img :src="this.post.PhotoStructure" alt="Foto" class="uploaded-photo" />
            </div>

            <div class="row">
                <div @click="toggleLike">
                    <svg width="24" height="24" viewBox="0 0 24 24">
                        
                        <path v-if="liked"
                            d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"
                            fill="red"></path>
                        <path v-else
                            d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"
                            fill="none" stroke="red" ></path>
                    </svg>
                    {{this.num_likes}}
                </div>  
            </div>
        </div>
    </div>
</template>

<style>
.v-center {

    display: inline-block;
    vertical-align: middle;
    line-height: normal;

}

.comment-button {

    padding: 0.35rem 0.5rem;
    font-size: 0.8rem;
    width: 120px;

}

.uploaded-photo {
    max-width: 250px;
    max-height: 250px;
    margin: 15px;
    object-fit: cover;
}
</style>