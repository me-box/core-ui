<template>
	<div>
		<div v-if="status" style="display: flex; flex-wrap: wrap;">
			<icon v-bind:name="item.name"
			      v-bind:displayName="true"
			      v-bind:updating="(item.state !== 'running')"
			      v-bind:route="'/view/' + item.name"
			      :key="item"
			      v-for="item in status"
			      v-if="(item.type === 'app' || item.type === 'driver') && item.name !== 'core-ui'"
			      style="margin: 8px"/>
			<icon name="App Store"
			      v-bind:displayName="true"
			      v-bind:updating="false"
			      route="/store"
			      style="margin: 8px"/>
			<icon name="Settings"
			      v-bind:displayName="true"
			      v-bind:updating="false"
			      route="/settings"
			      icon="settings"
			      style="margin: 8px"/>
		</div>
	</div>
</template>
<script>
	import testdata from '../testData/status.json'
	import Icon from '../components/AppIcon.vue'

	export default {
		name: 'Home',
		props: {},
		components: {
			Icon,
		},
		data() {
			//get data from api later
			return {status: {}, timerID: 0}
		},
		mounted() {
			this.$parent.setTitle("Databox Dashboard", true);
			let devmode = localStorage.getItem('dev');
			if (this.$parent.authenticated === "false" && devmode !== "true") {
				this.$router.push('login')
			}

			this.loadData();
			this.timerID = setInterval(() => {
				this.loadData();
			}, 1000);
		},
		destroyed: function () {
			clearInterval(this.timerID)
		},
		methods: {
			loadData: function () {
				this.ApiGetRequest('/core-ui/ui/api/containerStatus', testdata)
					.then(json => {
						this.status = json;
					})
			},
			GoToUI: function (appName) {
				this.$router.push("view?ui=" + appName)
			},
			Restart: function (appName) {
				alert("Install " + appName)
			},
			Uninstall: function (appName) {
				alert("Uninstall" + appName)
			}
		}
	}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>