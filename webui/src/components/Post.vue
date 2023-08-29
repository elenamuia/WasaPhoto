<script>
export default {

    props: {
        post: {
            type: Object
        }
    },

    data: function () {
        return {
            my_post: false,
            photo_id: null,
            liked: false,
            username: '',
            num_likes: 0,
            comments: [],
            comment_open: false,
            newComment: '',
            datapost: null,

        }
    },

    emits: ["delete-post"],

    methods: {

        async initialize() {
            this.datapost = this.post.Datapost;
            this.username = this.post.User;
            console.log("username is " + this.username);
            this.photo_id = this.post.ID
            console.log("PhotoID is " + this.photo_id);
            if (this.username == this.$current_user.id) {
                this.my_post = true;
            } else {
                this.my_post = false;
            }
            this.datapost = this.datapost.split("T");
            let date = this.datapost[0].split("-");
            let time = this.datapost[1].split(":");
            time = time[0] + ":" + time[1];
            date = date[2] + "-" + date[1] + "-" + date[0] + " at " + time;

            this.datapost = date;


            // Fetch likes
            let response = await this.$axios.get("/users/" + this.username + "/photos/" + this.photo_id + "/listlike/");
            if (response.data == null) {
                this.num_likes = 0;
                this.liked = false;

            } else {

                this.num_likes = response.data.length;

            }

            response = await this.$axios.get("/users/" + this.username + "/photos/" + this.photo_id + "/like/" + this.$current_user.id);
            this.liked = response.data;

            response = await this.$axios.get("/users/" + this.username + "/photos/" + this.photo_id + "/listcomment/");
            if (response.data != null) {
                this.comments = response.data
            }
            console.log(this.comments)

        },

        toggleLike() {
            this.liked = !this.liked;
            if (!this.liked) {
                this.Unlike();
            } else {
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
        async deletePhoto() {

            this.$axios.delete("/users/" + this.username + "/photos/" + this.photo_id);
            this.$emit("delete-post", this.post);

        },
        async toggleCommentList() {
            this.comment_open = !this.comment_open;


        },
        async addComment() {

            if (this.newComment.trim() !== '') {
                const Body = [];
                Body.push(this.newComment)
                Body.push(this.$current_user.id)
                this.$axios.post("/users/" + this.username + "/photos/" + this.photo_id + "/comments/", Body);
                this.comments.push({
                    UserPut: this.$current_user.id,
                    CommMessage: this.newComment

                })
                this.newComment = '';
            }

        },

        async deleteComment(commentid){
            this.$axios.delete("/users/"+this.username+"/photos/"+this.photo_id+"/comments/"+commentid)
            for (let i = 0; i < this.comments.length; i++) {
                if (this.comments[i].CommentID === commentid) {
                this.comments.splice(i, 1);
                i--;
                }
        }
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

            <div class="row align-content-between my-2"
                style="display:flex; justify-content: space-between; align-items:center; ">

                <div class="col">
                    <i class="bi-person-circle mx-1" style="font-size: 2em"></i>

                    <span class="col font-weight-bold h4">
                        {{ this.username }}
                    </span>

                </div>
                <div class="col-auto" v-if="my_post">
                    <div @click="deletePhoto">
                        <div class="nav-link">
                            <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#trash-2" />
                            </svg>
                        </div>
                    </div>
                </div>
            </div>

            <div class="row">
                <img :src="this.post.PhotoStructure" alt="Foto" class="uploaded-photo" />
            </div>


            <div class="row">
                <div class="col-sm-2">
                    <div @click="toggleLike()" class="icon-container">
                        <svg width="24" height="24" viewBox="0 0 24 24">

                            <path v-if="liked"
                                d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"
                                fill="red"></path>
                            <path v-else
                                d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"
                                fill="none" stroke="red"></path>
                        </svg>
                        {{ this.num_likes }}
                    </div>
                </div>
                <div class="col">
                    <div @click="toggleCommentList()" class="icon-container">
                        <div class="nav-link">
                            <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#message-circle" />
                            </svg>
                        </div>
                    </div>
                    <div v-if="this.comment_open" class="comment-container">
                        <div class="comments">
                            <div v-for="(comment, index) in comments" :key="index" class="comment">
                                <div class="comments">
                                    <span class="comment-user"><strong>{{ comment.UserPut }}</strong>:</span>
                                    <span class="comment-text">{{ comment.CommMessage}}</span>
                                    <div v-if="comment.UserPut === this.$current_user.id">
                                        <div @click="deleteComment(comment.CommentID)">
                                            <div class="nav-link">
                                                <svg class="feather">
                                                    <use href="/feather-sprite-v4.29.0.svg#trash-2" />
                                                </svg>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <hr class="divider" />
                            </div>
                        </div>

                        <hr class="divider" />

                        <div class="add-comment">
                            <input v-model="newComment" placeholder="Type here your comment..." />
                            <button @click="addComment">Send</button>
                        </div>
                    </div>



                </div>
                <div class="col">
                    <span class="text-muted v-center" style="font-size: 0.8em, font-style: italic;">
                        {{ datapost }}
                    </span>

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

.divider {
    border-top: 1px solid #ccc;
    margin: 10px 0;
}

.comment-container {
    border: 1px solid #ccc;
    padding: 10px;
}

.add-comment {
    background-color: #f5f5f5;
    ;
    padding: 10px;
}

.comment-user {
    margin-right: 5px;
}

.comment-content {
    display: flex;
    align-items: center;
    font-weight: bold;
}

.comment-text {
    flex: 1;
}

.uploaded-photo {
    max-width: 250px;
    max-height: 250px;
    margin: 15px;
    object-fit: cover;
}

.delete-icon {
    position: absolute;
    top: 10px;
    right: 10px;
    cursor: pointer;

}

.icon-container{
    display: flex;
    align-items: center;
    gap:5px;
    cursor: pointer;
}


</style>