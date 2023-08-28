<script>

export default {

    data: function () {
        return {
            searchQuery: '',
            searchResults: '',
            isSearchModalOpen: false,
            isPhotoModalOpen: false,
            errormsg: null,
            loading: false,
            some_data: null,
            url_profile: null,
            photos: [],
            img: {
                id: "",
                UserID: 0,
                postingdate: null,
                file: "",
            },
        }
    },
    methods: {
        openSearchModal() {
            this.isSearchModalOpen = true;

        },
        closeSearchModal() {
            console.log("closeSearchModal")
            this.isSearchModalOpen = false;

        },

        openPhotoModal() {
            console.log(this.isPhotoModalOpen)
            this.isPhotoModalOpen = true;
        },

        closePhotoModal() {
            this.isPhotoModalOpen = false;
        },

        openForm() {
            document.getElementById("myForm").style.display = "block";
        },

        closeForm() {
            document.getElementById("myForm").style.display = "none";
        },



        async postPhoto(userid, event) {
            this.loading = true;
            this.errormsg = null;
            const image = document.getElementById("fileInput").files[0];
            try {
                console.log(image)
                let fd = new FormData();
                fd.append("photo", image);
                let response = this.$axios.post("/users/" + userid + "/photos/", fd).then(res => res);
                this.img = response.data;
                this.closePhotoModal();
                this.$router.push('/mainstream/' + this.$current_user.id);
                request.onreadystatechange = function () {
                    if (request.readyState === 4) {
                        if (request.status === 200 && request.status.text === "OK") {
                            console.log("successful");

                        }
                        else {
                            console.log("failed")
                        }
                    }
                }
            } catch (e) {
                if (e.response && e.response.status === 404) {
                    errormsg.msg = ""
                }
                this.errormsg = e.toString();
            }
            this.loading = false;



        },


        async searchProfiles() {

            console.log("searcheduser: " + this.searchQuery);
            console.log("userid: " + this.$current_user.id);

            this.$axios.get('/users/' + this.$current_user.id + '/profile/' + this.searchQuery).then(response => {
                console.log("searchresults:" + response.data.User);
                if (response.data.User == null) {

                    this.searchResults = "No profile has been found";
                    this.searchQuery = '';

                } else {
                    this.searchQuery = '';
                    this.errormsg = null;

                    this.$router.push('/profile/' + response.data.User);
                }
            }).catch(err => {
                this.errormsg = err.message;
                console.log(this.errormsg)
            });

        },

        async Logout() {
            this.$current_user.id = null;
            this.$router.push("/");
        },

        async deleteProfile(){
            this.$axios.delete("/users/"+this.$current_user.id);
            this.$router.push('/');
            
        },

    },
    mounted() {

    }

}

</script>

<template>
    <header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow" style="width:100%;">
        <a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WasaPhoto</a>
        <button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse"
            data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
    </header>

    <div class="container-fluid">
        <div class="row">
            <div class="col-md-3">
                <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse"
                    style="margin-top:50px;">
                    <div class="position-sticky pt-3 sidebar-sticky">
                        <h6
                            class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">

                        </h6>
                        <ul class="nav flex-column">
                            <li class="nav-item">
                                <RouterLink :to="'/mainstream/' + this.$current_user.id" class="nav-link">
                                    <svg class="feather">
                                        <use href="/feather-sprite-v4.29.0.svg#home" />
                                    </svg>
                                    Home
                                </RouterLink>
                            </li>
                            <li class="nav-item" @click="openSearchModal">

                                <div class="nav-link">
                                    <svg class="feather">
                                        <use href="/feather-sprite-v4.29.0.svg#search" />
                                    </svg>
                                    Search Profile
                                </div>

                            </li>
                            <li class="nav-item">
                                <RouterLink :to="'/profile/' + this.$current_user.id" class="nav-link">
                                    <svg class="feather">
                                        <use href="/feather-sprite-v4.29.0.svg#user" />
                                    </svg>
                                    My Profile
                                </RouterLink>
                            </li>

                            <li class="nav-item" @click="openPhotoModal()">

                                <div class="nav-link">
                                    <svg class="feather">
                                        <use href="/feather-sprite-v4.29.0.svg#plus" />
                                    </svg>
                                    New Photo
                                </div>
                            </li>
                            <li class="nav-item" @click="Logout()">

                                <div class="nav-link">
                                    <svg class="feather">
                                        <use href="/feather-sprite-v4.29.0.svg#log-out" />
                                    </svg>
                                    Logout
                                </div>
                            </li>
                            <li class="nav-item" @click="deleteProfile()">

                                <div class="nav-link">
                                    <svg class="feather">
                                        <use href="/feather-sprite-v4.29.0.svg#user-x" />
                                    </svg>
                                    Delete Profile
                                </div>
                            </li>
                        </ul>
                    </div>
                </nav>
            </div>
            <main class="col-md-8" style="width:500px; height:auto">
                <div>
                    <div v-for="photo in photos" :key="photo">
                        <img :src="photo" alt="Photo not loaded">
                    </div>
                </div>


                <div class="photoModal" v-if="isPhotoModalOpen">
                    <div class="photo-modal-content">
                        <form @submit.prevent="postPhoto(this.$current_user.id, event)" class="form-container"
                            enctype="multipart/form-data">
                            <h3 style="margin-left: 10px;">Upload Photo</h3>
                            <input class="form-control" type="file" id="fileInput" accept="image/jpeg, image/png"
                                style="width:fit-content;">
                            <button type="submit" class="btn" id="submitBut">Upload</button>
                            <button type="button" class="btn cancel" @click="closePhotoModal()">Close</button>
                        </form>
                    </div>
                </div>

                <!--    <div class="form-popup" id="myForm">
                    <form @submit.prevent="postPhoto(this.$current_user.id, event)" class="form-container"
                        enctype="multipart/form-data">
                        <h3 style="margin-left: 10px;">Upload Photo</h3>
                        <input class="form-control" type="file" id="fileInput" accept="image/jpeg, image/png"
                            style="width:fit-content;">
                        <button type="submit" class="btn" id="submitBut">Upload</button>
                        <button type="button" class="btn cancel" @click="closeForm()">Close</button>
                    </form>

                </div> -->



                <div class="searchModal" v-if="isSearchModalOpen">
                    <!-- Campi di ricerca -->
                    <div class="search-modal-content">
                        <input type="text" v-model="searchQuery" placeholder="Who are you looking for?">
                        <!-- Bottone per eseguire la ricerca -->
                        <button @click="searchProfiles">Search</button>
                        <div>
                            <p v-text="searchResults"></p>
                        </div>

                        <!-- Bottone per chiudere la modale -->
                        <button @click="closeSearchModal">Close</button>
                    </div>
                </div>
                <RouterView :key="$route.fullPath"></RouterView>
            </main>

        </div>
    </div>
</template>
<style>
.searchModal {
    position: fixed;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999;
    /* Stili per il posizionamento e overlay */
}

.search-modal-content {
    background-color: white;
    padding: 20px;
    border-radius: 5px;

}

.photoModal {

    position: fixed;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    background-color: rgba(0, 0, 0, 0.5);


}

.photo-modal-content {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background-color: white;
    padding: 20px;
    border-radius: 5px;
}
</style>