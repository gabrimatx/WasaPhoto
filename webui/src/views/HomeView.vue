<script>
import PhotoCanvas from "@/components/PhotoCanvas.vue";
export default {
	components: {
		PhotoCanvas,
	},
	data: function () {
		return {
			errormsg: null,
			loading: false,
			some_data: null,

			publisherName: "John Doe",
			likeCount: 42,
			isLiked: true,
		};
	},
	methods: {
		async loadPhoto() {
		this.loading = true
		this.errormsg = null
		let response = await this.$axios.get("/photos/1")
		},
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		fetchImage() {
			const imageUrl = 'http://your-backend-server.com/path/to/your/image.jpg';

			axios
				.get(imageUrl, {
					responseType: 'arraybuffer', // This is important for handling binary data like images
				})
				.then((response) => {
					const base64Image = btoa(
						new Uint8Array(response.data).reduce((data, byte) => data + String.fromCharCode(byte), '')
					);

					this.imageSrc = `data:${response.headers['content-type']};base64,${base64Image}`;
				})
				.catch((error) => {
					console.error('Error fetching image:', error);
				});
		},
	},
	mounted() {
		this.refresh()
		this.fetchImage()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>
		<photo-canvas :photo-url="imageSrc" :publisher-name="publisherName" :like-count="likeCount"
			:is-liked="isLiked"></photo-canvas>
	</div>
</template>

<style></style>
