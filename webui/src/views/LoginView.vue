<template>
    <div class="login-container">
        <LoadingSpinner v-if="loading"></LoadingSpinner>
        <div class="login-form">
            <h2>Login into your account</h2>
            <form @submit.prevent="login">
                <label class="login-label" for="username">Username:</label>
                <input type="text" id="username" v-model="username" required minlength="3" maxlength="16"
                    style="padding: 6px;" />
                <button type="submit" class="btn btn-sm btn-outline-primary"
                    style="padding: 8px; float: right; font-size: 20px;">Login <svg class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#key" />
                    </svg></button>
            </form>
            <div v-if="identifier !== null">
                <p>Login successful! User identifier: {{ identifier.userId }}</p>
            </div>
        </div>
    </div>
</template>
  
<script>
export default {
    data() {
        return {
            username: "",
            identifier: null,
            loading: false,
        };
    },

    methods: {
        async login() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.post('/session/', { username: this.username }, {
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json',
                    },
                });
                this.identifier = response.data
                this.saveTokenToSessionStorage()
            } catch (error) {
                console.error("Error while logging in!");
            }
            this.loading = false;
            this.navigateToMyPage()
        },
        navigateToMyPage() {
            this.$router.push('/users/' + this.identifier.userId);
        },
        saveTokenToSessionStorage() {
            const bearerToken = `${this.identifier.userId}`;
            sessionStorage.setItem('authToken', bearerToken);
        },
    },
};
</script>
  
<style scoped>
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.login-form {
    padding: 20px;
    border: 1px solid #000000;
    align-items: center;
    border-radius: 8px;
}

.login-label {
    padding: 3px;
    display: block;
    margin-bottom: 8px;
}
</style>
  