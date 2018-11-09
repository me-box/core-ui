<template>
	<div class="mdc-select" :id="'select-' + dsID" style="margin-bottom: 8px">
		<select class="mdc-select__native-control" :name="dsID" v-model="selected">
			<option v-if="filteredDataSources.length === 0" value="" disabled selected>
				No available data sources ...
			</option>
			<option v-for="(ds, index) in filteredDataSources" :key="index" :value="ds.href">
				{{ds.description}}
			</option>
		</select>
		<label class="mdc-floating-label">{{dsID}}</label>
		<div class="mdc-line-ripple"></div>
	</div>
</template>
<script>
	import {MDCSelect} from '@material/select';

	export default {
		name: 'datasourcesDDL',
		props: {
			dataSource: Object,
			availableDataSources: Array,
			value: Object,
		},
		computed: {
			dsID: function () {
				return this.dataSource.clientid
			},
			selected: {
				get() {
					return this.value
				},
				set(v) {
					let selectedDS = null;
					for (const ds of this.availableDataSources) {
						if (ds.href === v) {
							selectedDS = ds;
							break;
						}
					}
					this.$emit('selected', selectedDS, this.dsID)
				}
			},
			filteredDataSources: function () {
				let getValFromHypercat = (hypercat, match) => {
					return hypercat["item-metadata"].filter((itm) => {
						return itm.rel === match
					})[0].val;
				};
				let filteredDs = [];

				this.availableDataSources.forEach((ds) => {

					let dsType = getValFromHypercat(ds, "urn:X-databox:rels:hasType");
					if (dsType === this.dataSource.type) {
						filteredDs.push({
							description: getValFromHypercat(ds, "urn:X-hypercat:rels:hasDescription:en"),
							vendor: getValFromHypercat(ds, "urn:X-databox:rels:hasVendor"),
							dsType: dsType,
							href: ds["href"],
							hypercat: ds
						})
					}
				});

				return filteredDs
			}
		},
		mounted() {
			new MDCSelect(document.querySelector('#select-' + this.dsID));
		},
	}

</script>

<style scoped>
</style>
