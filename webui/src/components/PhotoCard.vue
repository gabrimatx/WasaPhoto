<template>
  <div class="center-container">
    <div class="photo-card">
      <img :src="imgSrc" alt="Photo" class="photo" />
      <div class="photo-details">
        <div class="author">Author: {{ authorName }}</div>
        <div class="caption">
          <div class="caption-border"></div>
          <div class="caption-text">{{ caption }}</div>
          <div class="caption-border"></div>
        </div>
        <div class="actions">
          <button @click="likePhoto" class="btn btn-sm btn-outline-primary">Like</button>
          <span class="like-counter">{{ likeCount }} Likes</span>
          <button @click="commentPhoto" class="btn btn-sm btn-outline-secondary">Comment</button>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
const token = sessionStorage.getItem('authToken');
export default {
  props: {
    photoId: Number, // Assuming you pass the photoId as a prop
    likeCount: Number,
    authorName: String,
    caption: Text,
  },
  data() {
    return {
      photoDetails: null,
      imgSrc: null,
    };
  },

  async created() {
    if (this.photoId) {
      try {
        const response = await this.$axios.get(`/photos/${this.photoId}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
          responseType: 'blob', 
        });
        const imageUrl = URL.createObjectURL(response.data);
        this.imgSrc = imageUrl; 
      } catch (error) {
        console.error('Failed to fetch photo details:', error);
      }
    }
  },
  computed: {

  },
  methods: {
    likePhoto() {
      // Implement the logic to handle the like button click
      // You may want to send a request to the backend to update the like count
    },
    commentPhoto() {
      // Implement the logic to handle the comment button click
      // You can navigate to a comment page or show a comment form, for example
    },
  },
};
</script>

<style scoped>
.center-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.photo-card {
  border: 3px solid #6d6969;
  border-radius: 4px;
  padding: 10px;
  width: 400px; /* Adjust the width as needed */
  text-align: center;
  font-family: 'Arial', sans-serif;
}

.photo {
  width: 100%;
  height: auto;
  border: 1px solid #ddd;
}

.photo-details {
  margin-top: 10px;
}

.author {
  font-weight: bold;
  font-size: 20px;
  margin-bottom: 5px;
}

.actions {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

.like-counter {
  margin-left: 2px;
  border: 2px solid #d102027a;
  border-radius: 4px;
  padding: 8px;
}

.caption {
  display: flex;
  align-items: center;
  margin-top: 10px;
}

.caption-border {
  flex: 1;
  height: 3px;
  background-color: #1a1212;
  padding: 4px;
  margin-top: 10px;
  margin-bottom: 10px;
  
}

.caption-text {
  padding: 0 10px;
}

</style>
