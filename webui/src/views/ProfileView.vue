<template>
    <div class="user-profile">
        <div class="user-info">
            <h1>{{ userName }}</h1>
            <div>Followers: {{ followCount }}</div>
            <div>Following: {{ followedCount }}</div>
            <div>Banned: {{ isBanned ? 'Yes' : 'No' }}</div>
            <div>Followed: {{ isFollowed ? 'Yes' : 'No' }}</div>
            <div class="buttons">
                <button @click="toggleFollow" class="btn btn-outline-warning">
                    {{ isFollowed ? 'Unfollow' : 'Follow' }}
                </button>
                <button @click="toggleBan" class="btn btn-outline-danger">
                    {{ isBanned ? 'Unban' : 'Ban' }}
                </button>
            </div>
            <hr />
        </div>
        <div class="photos">
            <PhotoCard v-for="photo in photoList" :key="photo.id" :photoId="photo.id" :authorName="userName"
                :likeCount="photo.likecount" :caption="photo.caption" />
        </div>
    </div>
</template>
  
<script>
import PhotoCard from '@/components/PhotoCard.vue'; // Adjust the path based on your project structure
const token = sessionStorage.getItem('authToken');

export default {
    data() {
        return {
            userName: '',
            followCount: 0,
            followedCount: 0,
            isBanned: false,
            isFollowed: false,
            photoList: [],
        };
    },
    async created() {
        // Assuming you have a method to fetch user data based on userId
        this.fetchUserData();
    },
    methods: {
        async fetchUserData() {
            const userId = this.$route.params.userId;
            const response = await this.$axios.get(`/users/${userId}`, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            this.userName = response.data.userName;
            this.followCount = response.data.followCount;
            this.followedCount = response.data.followedCount;
            this.isBanned = response.data.isBanned;
            this.isFollowed = response.data.isFollowed;
            this.photoList = response.data.PList;
        },
        async toggleFollow() {
            // frontend
            this.isFollowed = !this.isFollowed;
            // backend
            const userId = this.$route.params.userId;
            await this.$axios.put(`/users/${token}/follows/${userId}`, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
        },
        async toggleBan() {
            // frontend
            this.isBanned = !this.isBanned;
            // backend

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
  
  