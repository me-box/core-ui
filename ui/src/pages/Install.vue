<template>
	<div style="max-width: 100%; padding: 16px">
		<div v-if="manifest == null || dataSources == null">Loading .....</div>
		<div v-if="manifest != null && dataSources != null">
			<div style="display: flex; margin-bottom: 8px">
				<icon :name="app" :displayName="false"/>
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
			<div v-if="loading" style="display: flex; justify-content: center;">
				<svg class="spinner" viewBox="0 0 66 66" xmlns="http://www.w3.org/2000/svg">
					<circle class="path" fill="none" stroke-width="6" stroke-linecap="round" cx="33" cy="33" r="30">
					</circle>
				</svg>
			</div>
			<div v-else-if="drivers">
				<h3>Data Source Required</h3>
				<div>Your Databox currently doesn't have any apps that provide the kind of data that {{ manifest.name}}
				uses. Please install one of the apps below so {{ manifest.name}} will have the data it needs.</div>
				<div v-for="driver in drivers" style="display: flex; align-items: center">
					<icon :name="driver.name" :displayName="false" :route="'/install/' + driver.name" style="zoom: 0.4" />
					<div style="flex-grow: 1; padding: 16px">{{ driver.name }}</div>
					<div class="mdc-button" @click="installDriver(driver.name)">Install</div>
				</div>
			</div>
			<div v-else-if="manifest.datasources" style="display: flex; flex-direction: column">
				<h3>Assign Data Sources to App</h3>
				<DatasourceSelect v-for="ds in manifest.datasources"
				                  :dataSource="ds"
				                  :availableDataSources="dataSources"
				                  :key="ds.id"
				                  @selected="onDsSelected"/>
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
				drivers: null,
				loading: true,
				selectedDataSources: {},
				installEnabled: false,
				timerID: null
			}
		},
		asyncComputed: {
			manifest: {
				lazy: true,
				get() {
					if (this.app == null) {
						return Promise.resolve(null)
					}
					this.loading = true;
					this.drivers = null;
					this.selectedDataSources = {};
					this.installEnabled = false;
					if(this.timerID) {
						clearInterval(this.timerID);
						this.timerID = null;
					}
					return this.$parent.apiRequest('/core-ui/ui/api/manifest/' + this.app, testManifest)
				},
				default: null,
			},
			dataSources: {
				lazy: true,
				get() {
					return this.$parent.apiRequest('/core-ui/ui/api/dataSources', testdata)
				},
				default: null,
			},
		},
		updated() {
			if (this.loading && this.manifest && this.dataSources) {
				let requiredSources = this.checkDatasource(this.dataSources);
				if (requiredSources.length === 0) {
					this.loading = false;
				} else {
					this.$parent.apiRequest('/core-ui/ui/api/drivers', {}, {
						method: 'POST',
						body: JSON.stringify(requiredSources)
					})
						.then((drivers) => {
							this.drivers = drivers;
							this.loading = false;
						});
				}
			}
			this.updateInstallButtonState()
		},
		mounted() {
			//get manifest name from Query string
			//this.manifestName = this.$route.query['manifest'];
			this.$parent.setTitle("Install " + this.app);
			this.updateInstallButtonState();
		},
		methods: {
			onDsSelected(ds, Clientid) {
				if (ds != null) {
					this.selectedDataSources[Clientid] = ds
				} else {
					this.selectedDataSources[Clientid] = undefined;
					delete this.selectedDataSources[Clientid]
				}
				this.updateInstallButtonState()
			},
			install() {
				if (this.installEnabled) {
					if (Array.isArray(this.manifest.datasources)) {
						this.manifest.datasources.forEach((ds, index, arr) => {
							arr[index].hypercat = this.selectedDataSources[ds["clientid"]].hypercat
						});
					}
					this.$parent.apiRequest('/core-ui/ui/api/install', {}, {
						method: 'POST',
						headers: {
							'Accept': 'application/json, text/plain, */*',
							'Content-Type': 'application/json'
						},
						body: JSON.stringify({
							manifest: this.manifest,
						}),
					})
						.then(() => {
							this.$router.push("/")
						})
				}
			},
			installDriver(driver) {
				this.loading = true;
				this.$parent.apiRequest('/core-ui/ui/api/install', {}, {
					method: 'POST',
					headers: {
						'Accept': 'application/json, text/plain, */*',
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						manifest: driver,
					}),
				})
					.then(() => {
						this.timerID = setInterval(() => {
							this.$parent.apiRequest('/core-ui/ui/api/dataSources', testdata)
								.then((datasources) => {
									let requiredSources = this.checkDatasource(datasources);
									if(requiredSources.length === 0) {
										clearInterval(this.timerID);
										this.timerID = null;
										this.loading = false;
										this.drivers = null;
										this.dataSources = datasources;
									}
								})
						}, 1000);


					})
			},
			checkDatasource(datasource) {
				const getValFromHypercat = (hypercat, match) => {
					return hypercat["item-metadata"].filter((itm) => {
						return itm.rel === match
					})[0].val;
				};

				let requiredSources = [];
				if(this.manifest.datasources) {
					for (const requiredDatasource of this.manifest.datasources) {
						let found = false;
						for (const availDatasource of datasource) {
							const type = getValFromHypercat(availDatasource, 'urn:X-databox:rels:hasType');
							if (type === requiredDatasource.type) {
								found = true;
								break;
							}
						}
						if (!found) {
							requiredSources.push(requiredDatasource.type);
						}
					}
				}

				return requiredSources;
			},
			updateInstallButtonState() {
				//TODO deal with optional datasources
				//TODO Add functionality for min and max datasources of a given type
				this.installEnabled = this.manifest != null && (!this.manifest.datasources || Object.keys(this.selectedDataSources).length === this.manifest.datasources.length);
			}
		}
	}
</script>
<style lang="scss">
	.disabled {
		color: #999;
	}

	$offset: 187;
	$duration: 1.4s;

	.spinner {
		animation: rotator $duration linear infinite;
		width: 35px;
		height: 35px;
		padding: 16px;
	}

	@keyframes rotator {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(270deg);
		}
	}

	.path {
		stroke-dasharray: $offset;
		stroke-dashoffset: 0;
		transform-origin: center;
		animation: dash $duration ease-in-out infinite,
		colors ($duration*4) ease-in-out infinite;
	}

	@keyframes colors {
		100%, 0% {
			stroke: #536878;
		}
	}

	@keyframes dash {
		0% {
			stroke-dashoffset: $offset;
		}
		50% {
			stroke-dashoffset: $offset/4;
			transform: rotate(135deg);
		}
		100% {
			stroke-dashoffset: $offset;
			transform: rotate(450deg);
		}
	}
</style>
