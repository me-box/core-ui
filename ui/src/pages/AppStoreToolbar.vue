<template>
	<div>
		<div class="material-icons mdc-top-app-bar__action-item" @click="gotoView">
			settings
		</div>
	</div>
</template>
<script>
	export default {
		name: 'storeToolbar',
		methods: {
			gotoView: function () {
				this.$router.push('/view/app-store');
			},
		}
	}
</script>