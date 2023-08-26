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
            //let response = await this.$axios.get("/users/" + this.username + "/photos/" + this.photo_id + "/listlike");
            //this.num_likes = response.data.LikeID.length;

            // Fetch comments
            //response = await this.$axios.get("/users/" + this.username + "/photos/" + this.photo_id + "/listcomment");           
        },

    },

    mounted() {

        this.initialize();

    }
}
</script>

<template>
    <div style="max-width: 500px;">
        <!-- Bordered Wrapper -->

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