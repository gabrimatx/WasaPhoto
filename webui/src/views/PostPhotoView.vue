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
      <p v-if="uploadSuccess" class="success-message">Photo uploaded successfully!</p>
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
        this.uploadSuccess = true;
      }
      catch (error) {
        console.error('Error', error)
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