<template>
	<div style="max-width: 100%; padding: 16px">
		<div v-if="manifest === null">
			<Spinner/>
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
			<div v-if="drivers === null || dataSources === null || installing"
			     style="display: flex; justify-content: center; padding: 8px">
				<Spinner/>
			</div>
			<div v-else>
				<template v-if="!allowDatasources">
					<template v-if="drivers && drivers.length > 0">
						<div class="warning-block">
							<span class="material-icons">warning</span>
							<div>Your Databox currently doesn't have any apps installed that supply the kinds of data
								that {{name}} uses. Please install one of the apps below so {{name}} will have the data
								it needs.
							</div>
						</div>
						<div v-for="driver in drivers" :key="driver.name" style="display: flex; align-items: center">
							<icon :name="driver.name" :displayName="false" :route="'/install/' + driver.name"
							      style="zoom: 0.4"/>
							<div style="flex-grow: 1; padding: 16px">{{ driver.name }}</div>
							<Spinner v-if="installingDriver === driver.name" style="padding: 16px"/>
							<button v-else class="mdc-button" :disabled="installingDriver !== null"
							        @click="installDriver(driver)">Install
							</button>
						</div>
					</template>
					<div v-else>
						<div class="warning-block">
							<span class="material-icons">warning</span>
							<div>Your Databox currently doesn't have any apps installed that supply the kinds of data
								that {{name}} uses.
							</div>
						</div>
					</div>
				</template>
				<div v-else-if="manifest.datasources != null && manifest.datasources.length > 0"
				     style="display: flex; flex-direction: column">
					<h3>Assign Data Sources to App</h3>
					<DatasourceSelect v-for="ds in manifest.datasources"
					                  :dataSource="ds"
					                  :availableDataSources="dataSources"
					                  :key="ds.id"
					                  @selected="onDsSelected"/>
				</div>

				<div v-if="!installing" style="display: flex; flex-direction: column; align-items: flex-end">
					<button class="mdc-button" :disabled="!allowInstall" @click="install">Install</button>
				</div>
			</div>
		</div>
	</div>
</template>
<script>
	import Icon from '../components/AppIcon.vue'
	import DatasourceSelect from '../components/DataSourceSelect.vue'
	import Spinner from '../components/Spinner.vue'

	import testdata from '../testData/DataSources.json'
	import testManifest from '../testData/app-manifest.json'

	export default {
		name: 'install',
		props: ['app'],
		components: {Icon, DatasourceSelect, Spinner},
		data() {
			return {
				installingDriver: null,
				installing: false,
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
				return !this.installing && this.installingDriver === null && this.manifest !== null && (!this.manifest.datasources || this.selectedDataSourceCount === this.manifest.datasources.length);
			},
			allowDatasources() {
				if (this.manifest && this.dataSources) {
					let requiredSources = this.checkDatasource(this.dataSources);
					console.log(requiredSources);
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
			this.$parent.title = "Install " + this.app;
			this.$parent.backRoute = "/store"
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
					this.installing = true;
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
							this.timerID = setInterval(() => {
								this.checkInstalled();
							}, 1000);
						})
				}
			},
			installDriver(driver) {
				if (this.installingDriver === null) {
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
			checkInstalled() {
				this.$parent.apiRequest('/core-ui/ui/api/containerStatus', testdata)
					.then(json => {
						for (const app of json) {
							if (app.name === this.app) {
								clearInterval(this.timerID);
								this.$router.push("/");
								break;
							}
						}
					});
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

	.warning-block {
		display: flex;
		padding: 16px;
		align-items: center;
		background: #FFB;
		margin: 16px 0 8px
	}

	.warning-block > .material-icons {
		padding-right: 16px;
		color: #990
	}
</style>
