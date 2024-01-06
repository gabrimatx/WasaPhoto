<template>
    <div class="custom-container">
        <h2 class="custom-title">User Search</h2>
        <form @submit.prevent="searchUsers">
            <label for="searchQuery" style="margin-right: 5px;">Search Users: </label>
            <input type="text" id="searchQuery" v-model="searchQuery" placeholder="Enter username" />
            <button type="submit" class="btn btn-sm btn-outline-secondary" style="margin-left: 10px;"
                @click="searchUsers">Search</button>
        </form>
        <p v-if="searchExecuted" style="margin-top: 30px;">{{ Text }}</p>
    </div>
</template>
  
<script>
const token = sessionStorage.getItem('authToken');

export default {
    data() {
        return {
            searchQuery: '',
            searchExecuted: false,
            Text: '',
        };
    },
    methods: {
        async searchUsers() {
            try {
                console.log("search started")
                const response = await this.$axios.get(`/users/`, {
                    params: { userName: this.searchQuery }, 
                    headers: {
                        'Authorization': `Bearer ${token}`, 
                        'Accept': 'application/json', 
                    },
                });
                console.log("search finished")
                this.searchExecuted = true;
                this.userId = response.data.userId;
                this.$router.push('/users/' + this.userId);
            }
            catch (error) {
                console.error(error, "mannaggia")
                this.searchExecuted = true;
                this.Text = "No users found with that username"
            }
        },
    },
};
</script>

<style scoped>
.custom-container {
    max-width: 600px;
    text-align: center;
    margin: auto;
    margin-top: 40px;
    justify-content: center;
    font-family: 'Courier New', Courier, monospace;
}

.custom-title {
    font-size: 3em;
    margin-bottom: 20px;
}
</style>
  