<script setup>
import Post from "./Post.vue"
</script>

<script>

export default {

    props: {
        posts: {
            type: Object // Photoposts to display
        }
    },

    watch: {
        posts: function (new_posts) {
            this.posts_ = new_posts;
        }
    },

    data: function () {
        return {
            posts_: this.posts
        }
    },

    

    methods: {

        async initialize() {

        },

        

    },
    mounted() {

        this.initialize();

    }
}
</script>

<template>

    <div v-for="(post, index) in posts_" :key="index" class="m-1">
        <Post :post="post"></Post>
    </div>

</template>

<style>

</style>