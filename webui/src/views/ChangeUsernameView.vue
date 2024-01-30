<template>
    <div class="container d-flex justify-content-center align-items-center vh-100">
        <form @submit.prevent="submitForm" class="border p-4 rounded">
            <h2 class="mb-4">Change your username</h2>
            <div class="mb-3">
                <label for="inputName" class="form-label">New Name</label>
                <input v-model="newname" type="text" class="form-control" id="inputName" required minlength="3" maxlength="16">
            </div>
            <button type="submit" class="btn btn-primary">Submit</button>
            <div class="alert alert-success" role="alert" v-if="changedSuccess" style="margin: 10px;">
                Name changed successfully!
            </div>
            <ErrorMsg :msg="error_msg" v-else-if="errore" style="margin: 10px;"/>
        </form>
    </div>
</template>
  
<script>
import ErrorMsg from '@/components/ErrorMsg.vue'
const token = sessionStorage.getItem('authToken');
export default {
    components: {
        ErrorMsg
    },
    data() {
        return {
            newname: '',
            changedSuccess: false,
            errore: false,
            error_msg: '',
        };
    },
    methods: {
        async submitForm() {
            try {
                const config = {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`,
                    },
                };
                const response = await this.$axios.put(`/users/${token}`, { username: this.newname }, config);
                console.log("Name changed");
                this.changedSuccess = true;
                this.errore = false;
            }
            catch (error) {
                console.error(error, "Error in changing name");
                const statusCode = error.response.status;
                switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized:', error.response.data);
                        this.error_msg = "You are not logged in"
                        break;
                    case 400:
                        console.error('Bad request:', error.response.data);
                        this.error_msg = "Name already in use"
                        break;
                    case 404:
                        console.error('Not found: ', error.response.data);
                        this.error_msg = "You are not logged in"
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                        this.error_msg = "You should login first"
                }
                this.changedSuccess = false;
                this.errore = true;
            }

            this.newname = '';

        },
    },
};
</script>
  