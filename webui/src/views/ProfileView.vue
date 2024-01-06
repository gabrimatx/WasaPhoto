<template>
    <div class="user-profile">
        <div v-if="found" class="user-info">
            <h1>{{ userName }}</h1>
            <div>Followers: {{ followCount }}</div>
            <div>Followed: {{ followedCount }}</div>

            <div class="buttons" v-if="!isItMe">
                <div class="user-info">Banned: {{ isBanned ? 'Yes' : 'No' }}</div>
                <div class="user-info">Followed: {{ isFollowed ? 'Yes' : 'No' }}</div>
                <button @click="toggleFollow" class="btn btn-outline-warning">
                    {{ isFollowed ? 'Unfollow' : 'Follow' }} <svg class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#user-plus" />
                    </svg>
                </button>
                <button @click="toggleBan" class="btn btn-outline-danger">
                    {{ isBanned ? 'Unban' : 'Ban' }} <svg class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#slash" />
                    </svg>
                </button>
            </div>
            <hr />
        </div>
        <div v-else>
            <v-center>
                <h1>{{ userName }}</h1>
            </v-center>
        </div>
        <div class="photos">
            <PhotoCard v-for="photo in photoList" :key="photo.id" :photoId="photo.id" :authorName="userName"
                :likeCount="photo.likecount" :caption="photo.caption" />
        </div>
    </div>
</template>
  
<script>
import PhotoCard from '@/components/PhotoCard.vue';
const token = sessionStorage.getItem('authToken');

export default {
    data() {
        return {
            userName: '',
            found: false,
            followCount: 0,
            followedCount: 0,
            isBanned: false,
            isFollowed: false,
            isItMe: false,
            photoList: [],
            reloadFlag: true,
        };
    },
    mounted() {
        if (localStorage.getItem('reloaded')) {
            localStorage.removeItem('reloaded');
        } else {
            localStorage.setItem('reloaded', '1');
            location.reload();
        }
    },
    async created() {
        const userId = this.$route.params.userId;
        this.isItMe = (userId == token);
        this.fetchUserData();
    },
    methods: {
        async fetchUserData() {
            const userId = this.$route.params.userId;
            try {
                const response = await this.$axios.get(`/users/${userId}`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                this.found = true;
                this.userName = response.data.userName;
                this.followCount = response.data.followCount;
                this.followedCount = response.data.followedCount;
                this.isBanned = response.data.isBanned;
                this.isFollowed = response.data.isFollowed;
                this.photoList = response.data.PList;
            } catch (error) {
                if (error.response) {
                    const statusCode = error.response.status;
                    switch (statusCode) {
                        case 401:
                            console.error('Access Unauthorized:', error.response.data);
                            // unauthorized
                            this.userName = "You are not logged in"
                            break;
                        case 403:
                            console.error('Access Forbidden:', error.response.data);
                            // forbidden
                            this.userName = "You have been banned by the user"
                            break;
                        case 404:
                            console.error('Not Found:', error.response.data);
                            // not found
                            this.userName = "User not found"
                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                    }
                } else {
                    console.error('Error:', error);
                }
            }
        },
        async toggleFollow() {
            // frontend
            this.isFollowed = !this.isFollowed;
            // backend
            const userId = this.$route.params.userId;
            const token = sessionStorage.getItem('authToken');
            if (this.isFollowed) {
                this.followCount += 1;
                await this.$axios.put(`/users/${token}/follows/${userId}`, {
                }, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
            } else {
                this.followCount -= 1;
                await this.$axios.delete(`/users/${token}/follows/${userId}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });

            }
        },
        async toggleBan() {
            // frontend
            this.isBanned = !this.isBanned;
            // backend
            const userId = this.$route.params.userId;
            const token = sessionStorage.getItem('authToken');
            if (this.isBanned) {
                await this.$axios.put(`/users/${token}/bans/${userId}`, {
                }, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
            } else {
                await this.$axios.delete(`/users/${token}/bans/${userId}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });

            }

        },
    },
    components: {
        PhotoCard,
    },
};
</script>
  
<style scoped>
.user-profile {
    font-family: 'serif';
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
}

.user-info {
    text-align: center;
    font-size: 20px;
}

.user-info h1 {
    font-size: 24px;
    margin-bottom: 10px;
}

.buttons {
    margin-top: 10px;
}

.buttons button {
    margin-right: 10px;
}

hr {
    margin: 20px 0;
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
  
  