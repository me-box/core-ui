<template>
	<div style="max-width: 100%; padding: 16px">
		<div v-if="manifest === null">
			<svg class="spinner" viewBox="0 0 66 66" xmlns="http://www.w3.org/2000/svg">
				<circle class="path" fill="none" stroke-width="6" stroke-linecap="round" cx="33" cy="33" r="30">
				</circle>
			</svg>
		</div>
		<div v-else>
			<div style="display: flex; margin-bottom: 8px">
				<icon :name="app" :displayName="false"/>
				<div style="margin-left: 8px">
					<div>{{ name }}</div>
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
			<div v-if="drivers === null || dataSources === null"
			     style="display: flex; justify-content: center;">
				<svg class="spinner" viewBox="0 0 66 66" xmlns="http://www.w3.org/2000/svg">
					<circle class="path" fill="none" stroke-width="6" stroke-linecap="round" cx="33" cy="33" r="30">
					</circle>
				</svg>
			</div>
			<div v-else>
				<template v-if="drivers && drivers.length > 0">
					<div style="display: flex;padding: 16px;align-items: center;background: #FFB; margin: 16px 0 8px">
						<span class="material-icons" style="padding-right: 16px; color: #990">warning</span>
						<div>Your Databox currently doesn't have any apps installed that supply the kinds of data that
							{{name}} uses. Please install one of the apps below so {{name}} will have the
							data it needs.
						</div>
					</div>
					<div v-for="driver in drivers" :key="driver.name" style="display: flex; align-items: center">
						<icon :name="driver.name" :displayName="false" :route="'/install/' + driver.name"
						      style="zoom: 0.4"/>
						<div style="flex-grow: 1; padding: 16px">{{ driver.name }}</div>
						<svg v-if="installingDriver === driver.name" class="spinner" viewBox="0 0 66 66" xmlns="http://www.w3.org/2000/svg">
							<circle class="path" fill="none" stroke-width="6" stroke-linecap="round" cx="33" cy="33" r="30">
							</circle>
						</svg>
						<button v-else class="mdc-button" :disabled="installingDriver !== null" @click="installDriver(driver)">Install</button>
					</div>
				</template>
				<div v-if="allowDatasources" style="display: flex; flex-direction: column">
					<h3>Assign Data Sources to App</h3>
					<DatasourceSelect v-for="ds in manifest.datasources"
					                  :dataSource="ds"
					                  :availableDataSources="dataSources"
					                  :key="ds.id"
					                  @selected="onDsSelected"/>
				</div>

			</div>

			<div style="display: flex; flex-direction: column; align-items: flex-end">
				<button class="mdc-button" :disabled="!allowInstall" @click="install">Install</button>
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
				installingDriver: null,
				selectedDataSources: new Map(),
				selectedDataSourceCount: 0,
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
			drivers: {
				lazy: true,
				get() {
					if (this.app == null || this.dataSources == null) {
						return Promise.resolve(null)
					}
					return this.$parent.apiRequest('/core-ui/ui/api/drivers/' + this.app, [])
				},
				default: null,
			}

		},
		computed: {
			name() {
				return this.manifest.displayName || this.app;
			},
			allowInstall() {
				return this.installingDriver === null && this.manifest !== null && (!this.manifest.datasources || this.selectedDataSourceCount === this.manifest.datasources.length);
			},
			allowDatasources() {
				if (this.manifest && this.dataSources) {
					let requiredSources = this.checkDatasource(this.dataSources);
					return requiredSources.length === 0;
				}
				return false
			}
		},
		destroyed() {
			if (this.timerID != null) {
				clearInterval(this.timerID)
			}
		},
		mounted() {
			this.$parent.setTitle("Install " + this.app);
		},
		beforeRouteUpdate(to, from, next) {
			if (to.name === from.name) {
				this.app = to.params.app;
				this.selectedDataSources.clear();
				this.selectedDataSourceCount = 0;
				if (this.timerID != null) {
					clearInterval(this.timerID);
					this.timerID = null;
				}
			}
			next()
		},
		methods: {
			onDsSelected(ds, Clientid) {
				if (ds != null) {
					this.selectedDataSources.set(Clientid, ds);
				} else {
					this.selectedDataSources.delete(Clientid);
				}
				this.selectedDataSourceCount = this.selectedDataSources.size;
			},
			install() {
				if (this.allowInstall) {
					if (Array.isArray(this.manifest.datasources)) {
						this.manifest.datasources.forEach((ds, index, arr) => {
							arr[index].hypercat = this.selectedDataSources.get(ds["clientid"])
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
				if(this.installingDriver === null) {
					this.installingDriver = driver.name;
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
								this.checkDriverInstalled(driver.name);
							}, 1000);
						})
				}
			},
			checkDriverInstalled(driver) {
				this.$parent.apiRequest('/core-ui/ui/api/containerStatus', testdata)
					.then(json => {
						for (const app of json) {
							if (app.name === driver) {
								clearInterval(this.timerID);
								this.checkDriverReady(driver);
								break;
							}
						}
					});
			},
			checkDriverReady(driver) {
				this.timerID = setInterval(() => {
					// TODO Timeout ==> configure
					this.$parent.apiRequest('/core-ui/ui/api/dataSources', testdata)
						.then((datasources) => {
							for (const datasource of datasources) {
								if (datasource.href.includes(driver)) {
									clearInterval(this.timerID);
									this.drivers = null;
									this.dataSources = datasources;
									this.installingDriver = null;
									break;
								}
							}
						});
				}, 1000);
			},
			checkDatasource(datasource) {
				const getValFromHypercat = (hypercat, match) => {
					return hypercat["item-metadata"].filter((itm) => {
						return itm.rel === match
					})[0].val;
				};

				let requiredSources = [];
				if (this.manifest.datasources) {
					for (const requiredDatasource of this.manifest.datasources) {
						if (requiredDatasource.required) {
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
				}

				return requiredSources;
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
