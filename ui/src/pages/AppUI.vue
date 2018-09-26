<template>
	<iframe :src="'/' + this.app + '/ui'"></iframe>
</template>
<script>
	export default {
		name: 'appUI',
		props: ['app'],
		data: function () {
			return {ui: ""}
		},
		mounted()
		{
			//get app name from Query string
			this.$parent.setTitle(this.app);
		},
	}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	iframe {
		width: 100%;
		height: 100%;
		border: none;
	}
</style>