<template>
	<div class="mdc-list" data-mdc-auto-init="MDCList">
		<div class="mdc-list-item" v-on:click="goto('/settings/system')">
			<span class="mdc-list-item__graphic material-icons" aria-hidden="true">info</span>
			System
		</div>
		<div class="mdc-list-item" v-on:click="goto('/data')">
			<span class="mdc-list-item__graphic material-icons" aria-hidden="true">data_usage</span>
			My Data
		</div>
	</div>
</template>
<script>
	export default {
		name: 'settings',
		props: {},
		created: function () {
			let devmode = localStorage.getItem('dev');
			if (this.$parent.authenticated === "false" && devmode !== "true") {
				this.$router.push('login')
			}
		},
		mounted: function () {
			this.$parent.setTitle("Databox Settings");
		},
		methods: {
			goto: function (path) {
				this.$router.push(path);
			}
		}
	}
</script>