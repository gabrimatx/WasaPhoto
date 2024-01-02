<template>
    <div class="photo-canvas">
      <img :src="photoUrl" alt="Photo" />
      <div class="info-section">
        <p class="publisher">{{ publisherName }}</p>
        <div class="like-container">
          <button @click="toggleLike" class="like-button">{{ likeText }}</button>
          <span class="like-count">{{ likeCount }} Likes </span>
        </div>
        <button @click="showComments" class="comment-button">Comment</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      photoUrl: String,
      publisherName: String,
      likeCount: Number,
      isLiked: Boolean,
    },
    data() {
      return {
        liked: this.isLiked,
      };
    },
    computed: {
      likeText() {
        return this.liked ? 'Unlike' : 'Like';
      },
    },
    methods: {
      toggleLike() {
        this.liked = !this.liked;
      },
      showComments() {
        // Implement your logic to show comments
        console.log('Show comments clicked');
      },
    },
  };
  </script>
  
  <style scoped>
  .photo-canvas {
    width: 100%;
    max-width: 600px;
    overflow: hidden;
    border: 2px solid #ddd;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
    margin: 10px;
  }
  
  img {
    width: 100%;
    height: auto;
    display: block;
  }
  
  .info-section {
    background-color: #fff;
    padding: 20px;
  }
  
  .publisher {
    font-family: 'Arial', sans-serif;
    font-size: 18px;
    margin: 0;
  }
  
  .like-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 20px;
  }
  
  .like-button {
    cursor: pointer;
    background-color: #e74c3c; /* Change color when liked */
    border: #e74c3c;    
    border-radius: 5px;
    font-family: 'Arial', sans-serif;
  }

  .comment-button {
    cursor: pointer;
    background-color: #3498db;
    color: white;
    border: none;
    padding: 10px 15px;
    border-radius: 5px;
    font-family: 'Arial', sans-serif;
  }
  
  
  .like-count {
    margin-left: 10px;
    font-size: 16px;
  }
  </style>
  