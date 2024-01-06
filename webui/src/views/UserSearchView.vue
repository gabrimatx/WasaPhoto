<template>
    <div class="container mt-5 text-center">
      <h2 class="display-4 mb-4">User Search</h2>
      <form @submit.prevent="searchUsers" class="mb-4">
        <div class="form-group d-flex justify-content-center align-items-center">
          <label for="searchQuery" class="mr-3" style="font-size: 30px; margin: 20px;">Username: </label>
          <input type="text" id="searchQuery" v-model="searchQuery" class="form-control" placeholder="Enter username" />
          <button type="submit" class="btn btn-sm btn-outline-secondary ml-2" @click="searchUsers" style="margin: 20px; font-size: 30px;">Search</button>
        </div>
      </form>
      <p v-if="searchExecuted" class="mt-3" style="font-size: 25px;">{{ Text }}</p>
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
</style>
  