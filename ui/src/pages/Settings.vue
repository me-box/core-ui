<template>
	<div class="mdc-list mdc-list--two-line" data-mdc-auto-init="MDCList">
		<div class="mdc-list-item" v-on:click="goto('/settings/system')">
			<div class="mdc-list-item__graphic material-icons" aria-hidden="true">info</div>
			<div class="mdc-list-item__text">
				<div class="mdc-list-item__primary-text">System</div>
				<div class="mdc-list-item__secondary-text">Databox name, password</div>
			</div>
		</div>
		<div class="mdc-list-item" v-on:click="goto('/data')">
			<div class="mdc-list-item__graphic material-icons" aria-hidden="true">data_usage</div>
			<div class="mdc-list-item__text">
				<div class="mdc-list-item__primary-text">Data</div>
				<div class="mdc-list-item__secondary-text">Backup, export data</div>
			</div>
		</div>
		<div class="mdc-list-item" v-on:click="goto('/data')">
			<div class="mdc-list-item__graphic material-icons" aria-hidden="true">history</div>
			<div class="mdc-list-item__text">
				<div class="mdc-list-item__primary-text">History</div>
				<div class="mdc-list-item__secondary-text">App history</div>
			</div>
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