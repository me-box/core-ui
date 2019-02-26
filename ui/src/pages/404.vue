<template>
	<div class="Error">
		<h1>404: Page Not Found</h1>
	</div>
</template>
<script>

	export default {
		name: '404',
		props: {},
		mounted() {
			this.$parent.title = 'Databox Dashboard';
			this.$parent.backRoute = '/';
		}
	}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
