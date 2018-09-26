<template>
	<div style="max-width: 100%">
		<div v-if="manifest == null || dataSources == null">Loading .....</div>
		<div v-if="manifest != null && dataSources != null">
			<div style="display: flex; margin-bottom: 8px">
				<icon v-bind:name="app" v-bind:displayName="false"/>
				<div style="margin-left: 8px">
					<div>{{ manifest.name }}</div>
					<div style="font-size: small">{{ manifest.author }}</div>
					<div style="font-size: small">{{ manifest.version }}</div>
				</div>
			</div>
			<div style="display: flex; align-items: center; margin-bottom: 8px">
				<span class="material-icons">star</span>
				<span class="material-icons">star</span>
				<span class="material-icons">star</span>
				<span class="material-icons">star</span>
				<span class="material-icons disabled">star</span>
				&nbsp; Moderate Risk
			</div>
			<div>{{ manifest.description }}</div>
			<div v-if="manifest.datasources" style="display: flex; flex-direction: column">
				<h3>Assignee Data Sources to App</h3>
				<DatasourceSelect v-for="ds in manifest.datasources" v-bind:dataSource="ds"
				                v-bind:availableDataSources="dataSources" v-bind:key="ds.id"
				                v-on:selected="onDsSelected"/>
			</div>

			<div style="display: flex; flex-direction: column; align-items: flex-end">
				<button class="mdc-button" :disabled="!installEnabled" @click="install">Install</button>
			</div>
		</div>
	</div>
</template>
<script>
	import Icon from '../components/AppIcon.vue'
	import DatasourceSelect from '../components/DataSourceSelect.vue'

	import testdata from '../testData/DataSources.json'
	import testManifest from '../testData/app-manifest.json'

	export default {
		name: 'install',
		props: ['app'],
		components: {Icon, DatasourceSelect},
		data() {
			return {
				selectedDataSources: {},
				installEnabled: false,
				loadingMan: "true",
				loadingData: "true"
			}
		},
		asyncComputed: {
			manifest: {
				lazy: true,
				get() {
					if (this.app == null) {
						return Promise.resolve(null)
					}

					console.log(this.app);
					return fetch('/core-ui/ui/api/manifest/' + this.app, {credentials: "same-origin"})
						.then((response) => {
							if (response.status === 401) {
								localStorage.setItem('databoxAuthenticated', false);
								this.$router.push('/');
								return
							}
							return response.json()
						})
						.then((json) => {
							return json;
						})
						.catch(() => {
							return new Promise((resolve, reject) => {
								setTimeout(() => {
									resolve(testManifest)
								}, 500)
							});
						});
				},
				default: null,
			},
			dataSources: {
				lazy: true,
				get() {
					return fetch('/core-ui/ui/api/dataSources', {credentials: "same-origin"})
						.then((response) => {
							if (response.status === 401) {
								localStorage.setItem('databoxAuthenticated', false);
								this.$router.push('/');
								return
							}
							return response.json()
						})
						.then((json) => {
							return json;
						})
						.catch(() => {
							return new Promise((resolve, reject) => {
								setTimeout(() => {
									resolve(testdata)
								}, 500)
							});
						});
				},
				default: null,
			},
		},
		updated() {
			this.updateInstallButtonState()
		},
		mounted() {
			//get manifest name from Query string
			//this.manifestName = this.$route.query['manifest'];
			this.$parent.setTitle("Install " + this.app);
			this.updateInstallButtonState();
		},
		methods: {
			onDsSelected: function (ds, Clientid) {
				if (ds != null) {
					this.selectedDataSources[Clientid] = ds
				} else {
					this.selectedDataSources[Clientid] = undefined;
					delete this.selectedDataSources[Clientid]
				}
				this.updateInstallButtonState()
			},
			install: function () {
				if (this.installEnabled) {
					if (Array.isArray(this.manifest.datasources)) {
						this.manifest.datasources.forEach((ds, index, arr) => {
							arr[index].hypercat = this.selectedDataSources[ds["clientid"]].hypercat
						});
					}
					let data = {
						manifest: this.manifest,
					};
					fetch(`/core-ui/ui/api/install`, {
						method: "POST",
						credentials: "same-origin",
						headers: {
							'Accept': 'application/json, text/plain, */*',
							'Content-Type': 'application/json'
						},
						body: JSON.stringify(data),
					})
						.then((response) => {
							if (response.status === 401) {
								localStorage.setItem('databoxAuthenticated', false);
								this.$router.push('/');
								return
							}
							if (response.status !== 200) {
								alert("Error installing app. Sorry.")
							} else {
								this.$router.push("MyApps")
							}
						})
						.catch(error => console.error(error));
				}

				//else
				//alert("Please select the required data sources")

			},
			cancel: function () {
				this.$router.push("AppStore")
			},
			updateInstallButtonState: function () {
				//TODO deal with optional datasources
				//TODO Add functionality for min and max datasources of a given type
				this.installEnabled = this.manifest != null && (!this.manifest.datasources || Object.keys(this.selectedDataSources).length === this.manifest.datasources.length);
			}
		}
	}
</script>
<style>
	.disabled {
		color: #909090;
	}
</style>
