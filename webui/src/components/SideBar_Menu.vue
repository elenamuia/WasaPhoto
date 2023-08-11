<script>
import SearchModal from './SearchModal.vue';
export default {

    components: {
        SearchModal
    },

    data: function () {
        return {
            isSearchModalOpen: false,
            errormsg: null,
            loading: false,
            some_data: null,
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
            this.isSearchModalOpen = false;
        },

        openForm() {
            document.getElementById("myForm").style.display = "block";
        },

        closeForm() {
            document.getElementById("myForm").style.display = "none";
        },
        openSearchBar(){
            document.getElementById("search").style.display = "block";
        },

        Search() {
            let search = document.querySelector("input").value;
			search = search.trim();
			if (search.length > 0) {
				// query ./users for results
				
				this.$router.push('/profile/' + search);
            }

        },


        postPhoto(userid, event) {
            this.loading = true;
            this.errormsg = null;
            const image = document.getElementById("fileInput").files[0];
            try {
                console.log(image)
                let fd = new FormData();
                fd.append("photo", image);

                let response = this.$axios.post("/users/" + userid + "/photos/", fd).then(res => res);

                this.img = response.data
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
            document.getElementById("myForm").style.display = "none";

        },

       



        Logout() {
            this.$current_user.id = null;
            this.$router.push("/");
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
                <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse" style="margin-top:50px;">
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
                            <li class="nav-item" @click="openSearchBar()">
                            <!--
                                <div class="nav-link">
                                    <svg class="feather">
                                        <use href="/feather-sprite-v4.29.0.svg#search" />
                                    </svg>
                                    Search Profile
                                </div>
                            -->
                                <div>
                                    <button @click="openSearchModal">Cerca Profili</button>
                                    <search-modal :showModal="isSearchModalOpen" @close="closeSearchModal" />
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
                        </ul>

                        <h6
                            class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
                            <span>Secondary menu</span>
                        </h6>
                        <ul class="nav flex-column">
                            <li class="nav-item">
                                <a class="nav-link"></a>
                                <RouterLink :to="'/settings/'" class="nav-link">
                                    <svg class="feather">
                                        <use href="/feather-sprite-v4.29.0.svg#settings" />
                                    </svg>
                                    Settings
                                </RouterLink>

                            </li>

                            <li class="nav-item" @click="openForm()">

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
                        </ul>
                    </div>
                </nav>
            </div>
                <main class= "col-md-8" style ="width:500px; height:auto">
                    <div >
                        <div v-for="photo in photos" :key="photo" >
                            <img :src="photo" alt="Photo not loaded">
                        </div>
                    </div>

                    <div class="form-popup" id="myForm">
                        <form @submit.prevent="postPhoto(this.$current_user.id, event)" class="form-container"
                            enctype="multipart/form-data">
                            <h3 style="margin-left: 10px;">Upload Photo</h3>
                            <input class="form-control" type="file" id="fileInput" accept="image/jpeg, image/png" style ="width:fit-content;">
                            <button type="submit" class="btn" id="submitBut">Upload</button>
                            <button type="button" class="btn cancel" @click="closeForm()">Close</button>
                        </form>

                    </div>
                <!--
                    <div>
                        <form>

                            <input v-model="searchQuery" type="text" id="search">
                            <button type="submit" class="btn" id="searchBut" @click="Search()">Search</button>
                                
                                <datalist class="list-group custom-select w-25 dropdown mt-5 position-absolute">

                                    <option class=" list-group-item align-middle" v-for="user in search_results"
                                        :key="user">

                                        <i class="bi-person-circle m-2 fa-lg" style="font-size: 1.5rem;"></i>

                                    </option>
                                </datalist>
                        </form>
                    </div>
                -->
                </main>

        </div>
    </div>
</template>
<style></style>