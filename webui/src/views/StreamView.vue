<script>
import PhotoCard from '@/components/PhotoCard.vue';
const token = sessionStorage.getItem('authToken');

export default {
    mounted() {
        if (localStorage.getItem('reloadedstream')) {
            localStorage.removeItem('reloadedstream');
        } else {
            localStorage.setItem('reloadedstream', '1');
            location.reload();
        }
    },
    data() {
        return {
            photoList: [],
            titlePage: "Your stream",
        }
    },
    async created() {
        const userId = this.$route.params.userId;
        this.fetchUserData();
    },
    methods: {
        async fetchUserData() {
            const userId = this.$route.params.userId;
            try {
                const response = await this.$axios.get(`/users/${userId}/stream/`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                this.photoList = response.data.PList;
            } catch (error) {
                if (error.response) {
                    const statusCode = error.response.status;
                    switch (statusCode) {
                        case 401:
                            console.error('Access Unauthorized:', error.response.data);
                            // unauthorized
                            this.titlePage = "You are not logged in"
                            break;
                        case 403:
                            console.error('Access Forbidden:', error.response.data);
                            // forbidden
                            this.titlePage = "You have been banned by the user"
                            break;
                        case 404:
                            console.error('Not Found:', error.response.data);
                            // not found
                            this.titlePage = "You are not logged in"
                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                    }
                } else {
                    console.error('Error:', error);
                }
            }
        },
    },
    components: {
        PhotoCard,
    },
}
</script>




<template>
    <h1 class="custom-title"> {{ titlePage }} </h1>
    <hr />
    <div class="photos">
        <PhotoCard v-for="photo in photoList" :key="photo.id" :photoId="photo.id" :authorName="photo.publisherName"
            :likeCount="photo.likecount" :caption="photo.caption" />
    </div>
</template>


<style scoped>
.custom-title {
    font-family: 'serif';
    font-size: 40px;
    margin: 40px;
    text-align: center;
}

.photos {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-around;
}

.photos .photo-card {
    width: 200px;
    margin-bottom: 30px;
}
</style>