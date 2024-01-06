import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import PostPhotoView from '../views/PostPhotoView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/session/', component: LoginView},
		{path: '/photos/', component: PostPhotoView},
		{path: '/users/:userId', component: ProfileView},
	]
})

export default router
