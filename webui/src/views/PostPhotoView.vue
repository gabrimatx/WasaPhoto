<template>
  <div class="custom-container">
    <h2 class="custom-title">Upload Photo</h2>
    <form @submit.prevent="uploadPhoto" class="custom-form">
      <div>
        <label for="photo">Select Photo:</label>
        <input type="file" id="photo" @change="onFileChange" />
        <span v-if="!photo">Photo is required</span>
      </div>

      <div>
        <label for="caption">Caption:</label>
        <textarea id="caption" v-model="caption"></textarea>
      </div>

      <button type="submit" class="btn btn-info" style="font-size: 30px;">Upload</button>
      <p v-if="uploadSuccess" class="success-message">{{ endText }}</p>
    </form>
  </div>
</template>

<script>
const token = sessionStorage.getItem('authToken');
export default {
  data() {
    return {
      photo: null,
      caption: '',
      uploadSuccess: false,
      endText: '',
    };
  },
  methods: {
    onFileChange(event) {
      this.photo = event.target.files[0];
    },
    async uploadPhoto() {
      if (!this.photo) {
        console.log('Photo is required');
        return;
      }

      const formData = new FormData();
      formData.append('file', this.photo);
      const additionalData = {
        Caption: this.caption,
      };

      formData.append('additionalData', JSON.stringify(additionalData));
      const config = {
        headers: {
          'Content-Type': 'multipart/form-data',
          'Authorization': `Bearer ${token}`,
        },
      };
      try {
        const response = await this.$axios.post(`/photos/`, formData, config);
        console.log('Photo uploaded successfully', response.data);
        this.endText = "Photo uploaded!";
        this.uploadSuccess = true;
      }
      catch (error) {
        const statusCode = error.response.status;
        switch (statusCode) {
                        case 401:
                            console.error('Access Unauthorized');
                            // unauthorized
                            this.endText = "You have to log in to post a photo";
                            this.uploadSuccess = true;
                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                    }
      }

    },
  },
};
</script>

<style scoped>
.custom-container {
  max-width: 600px;
  margin: 0 auto;
  text-align: center;
  justify-content: center;
  margin-top: 40px;
  font-family: 'Courier New', Courier, monospace;
}

.custom-title {
  font-size: 3em;
  margin-bottom: 20px;
}

.custom-form {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.custom-form div {
  margin-bottom: 15px;
}

.custom-form label {
  font-size: 2em;
}

.custom-form input,
.custom-form textarea {
  width: 100%;
  padding: 8px;
  font-size: 1em;
}

.success-message {
  color: green;
  margin-top: 10px;
}
</style>