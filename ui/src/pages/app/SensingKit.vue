<template>
	<div id="sensorKit">
		<template v-if="installed == null || loading">
			<Spinner/>
		</template>
		<template v-else-if="installed === false">
			<div class="material-icons mdc-theme--primary" style="font-size: 64px; align-self: center">phonelink_ring</div>
			<div style="width: 200px; margin: 16px; align-self: center">The Sensor Kit allows you to capture data from your mobile and send it securely to your Databox</div>
			<button class="mdc-button mdc-button--unelevated" style="align-self: center" @click="installSensingKit">Install</button>
		</template>
		<div v-else class="mdc-list mdc-list--avatar-list">
			<div style="padding: 0 16px 16px 16px;">The Sensor Kit allows you to capture data from your mobile and send it securely to your Databox</div>
			<label :for="sensorId(sensor)" v-for="sensor in sensors" :key="sensor" class="mdc-list-item" role="checkbox"
			       :aria-checked="isEnabled(sensor)">
				{{ sensor }}
				<div class="mdc-checkbox mdc-list-item__meta">
					<input :id="sensorId(sensor)" type="checkbox" class="mdc-checkbox__native-control"
					       :name="sensor" :checked="isEnabled(sensor)" @change="checkboxChange($event)">
					<div class="mdc-checkbox__background">
						<svg class="mdc-checkbox__checkmark" viewBox="0 0 24 24">
							<path class="mdc-checkbox__checkmark-path"
							      fill="none"
							      d="M1.73,12.91 8.1,19.28 22.79,4.59">
							</path>
						</svg>
						<div class="mdc-checkbox__mixedmark"></div>
					</div>
				</div>
			</label>
		</div>
	</div>
</template>
<script>
	import Spinner from '../../components/Spinner.vue'
	import testdata from '../../testData/DataSources.json'
	import {MDCCheckbox} from '@material/checkbox';

	export default {
		name: 'install',
		components: {Spinner},
		data() {
			return {
				loading: false
			}
		},
		asyncComputed: {
			installed: {
				lazy: true,
				get() {
					return this.$parent.apiRequest('/core-ui/ui/api/containerStatus', testdata)
						.then(json => {
							for (const app of json) {
								if (app.name === "driver-sensingkit") {
									return true;
								}
							}
							return false;
						});
				},
				default: null,
			},
			sensors: {
				lazy: true,
				get() {
					if (!this.installed) {
						return Promise.resolve([])
					}
					return new Promise(function (resolve) {
						window.SensingKit.listSensors((result) => {
							result.sort();
							resolve(result);
						});
					});
				},
				default: [],
			}

		},
		computed: {
			enabledSensors() {
				let enabled_sensors = localStorage.getItem('sensors');
				if (!enabled_sensors) {
					return [];
				} else {
					return JSON.parse(enabled_sensors);
				}
			}
		},
		mounted() {
			this.$parent.title = "Sensing Kit";
			this.$parent.backRoute = "/"
		},
		updated() {
			const checkboxes = document.getElementsByClassName('mdc-checkbox');
			for (const checkboxElement of checkboxes) {
				new MDCCheckbox(checkboxElement);
			}
		},
		methods: {
			isEnabled(sensor) {
				return this.enabledSensors.indexOf(sensor) !== -1;
			},
			sensorId(sensor) {
				return 'sensor_' + sensor.toLowerCase().replace(/ /g, "_")
			},
			checkboxChange(event) {
				const name = event.target.name;
				if (event.target.checked) {
					if (this.enabledSensors.indexOf(name) === -1) {
						this.enabledSensors.push(name);
					}
				} else {
					for (let i = 0; i < this.enabledSensors.length; i++) {
						if (this.enabledSensors[i] === name) {
							this.enabledSensors.splice(i, 1);
						}
					}
				}
				const url = 'https://' + this.$parent.databoxUrl + '/driver-sensingkit';
				console.log(this.enabledSensors);
				window.SensingKit.startSensors(this.enabledSensors, url, () => {});
			},
			installSensingKit() {
				this.loading = true;
				this.$parent.apiRequest('/core-ui/ui/api/manifest/driver-sensingkit')
					.then((manifest) => {
						this.$parent.installAndWait(manifest)
							.then(() => {
								this.loading = false;
								this.installed = true;
							})
					});
			}
		}
	}
</script>
<style lang="scss">
	@import "~@material/checkbox/mdc-checkbox";
	#sensorKit {
		display: flex;
		flex-direction: column;
		margin: 24px;
		max-width: 100%;
	}

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
