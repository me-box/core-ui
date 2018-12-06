<template>
	<div>
		<div v-if="apps" id="appList">
			<icon v-bind:name="item.name"
			      v-bind:displayName="true"
			      v-bind:updating="(item.state !== 'running')"
			      v-bind:route="item.route"
			      v-bind:icon="item.icon"
			      :key="item.name"
			      v-for="item in apps"
			      v-if="(item.type === 'app' || item.type === 'driver') && !exclusions.includes(item.name)"
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
			return {
				apps: [],
				timerID: 0,
				exclusions: ["core-ui", "app-store", "driver-sensingkit"]
			}
		},
		mounted() {
			this.$parent.title = 'Databox Dashboard';
			this.$parent.backRoute = null;
			this.loadData();
			this.timerID = setInterval(() => {
				this.loadData();
			}, 2000);
		},
		destroyed() {
			clearInterval(this.timerID)
		},
		methods: {
			loadData() {
				this.$parent.apiRequest('/core-ui/ui/api/containerStatus', testdata)
					.then(json => {
						let changed = false;
						json.push({
							name: "Settings",
							type: "app",
							state: "running",
							icon: "settings",
							route: "/settings"
						});
						json.push({
							name: "App Store",
							type: "app",
							state: "running",
							icon: "apps",
							route: "/store"
						});
						if(this.isMobile) {
							json.push({
								name: "Sensing Kit",
								type: "app",
								state: "running",
								icon: "phonelink_ring",
								route: "/sensing"
							});
						}
						json.sort((a, b) => {
							return a.name.toLowerCase().localeCompare(b.name.toLowerCase());
						});
						if (this.apps.length === json.length) {
							for (let step = 0; step < this.apps.length; step++) {
								let app1 = this.apps[step];
								let app2 = json[step];
								if ((app1.name !== app2.name) || (app1.status !== app2.status)) {
									changed = true;
									break;
								}
							}
						} else {
							changed = true;
						}
						if (changed) {
							for (const app of json) {
								if (app.route == null) {
									app.route = '/view/' + app.name;
								}
							}
							this.apps = json;
						}
					})
			},
			GoToUI(appName) {
				this.$router.push("view?ui=" + appName)
			},
			Restart(appName) {
				alert("Install " + appName)
			},
			Uninstall(appName) {
				alert("Uninstall" + appName)
			}
		}
	}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	#appList {
		display: flex;
		flex-wrap: wrap;
		margin: 24px;
	}

	@media (max-width: 479px) {
		#appList {
			margin: 16px;
		}
	}
</style>