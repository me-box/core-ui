<template>
	<div class="mdc-list" data-mdc-auto-init="MDCList">
		<div class="mdc-list-item" v-on:click="goto('/datasource/All Data')">
			<span class="mdc-list-item__graphic material-icons" aria-hidden="true">folder</span>
			All Data
		</div>
		<div class="mdc-list-item" v-for="source in dataSources" v-on:click="goto('/data/' + source)">
			<span class="mdc-list-item__graphic material-icons" aria-hidden="true">folder</span>
			{{source}}
		</div>
	</div>
</template>
<script>
	import testdata from '../testData/DataSources.json'

	export default {
		name: 'databoxStatus',
		props: {},
		created() {
			this.$parent.setTitle("Databox Data");
		},
		data() {
			//get data from api later
			return {
				dataSources: {},
				timerID: 0,
			}
		},
		mounted() {
			this.loadData();
			this.timerID = setInterval(() => {
				this.loadData();
			}, 5000);
		},
		destroyed: function () {
			clearInterval(this.timerID)
		},
		methods: {
			loadData: function () {
				this.ApiGetRequest('/core-ui/ui/api/dataSources', testdata)
					.then(json => {
						const extractDriverName = (hyperCat) => {
							//the URL api will only pars http/s urs so replace tcp: with http:
							const url = new URL(hyperCat.href.replace("tcp:", "http:"));
							return url.hostname.replace("-core-store", "");
						};

						let sourceNames = [];
						json.forEach((ds) => {
							let name = extractDriverName(ds);
							if (sourceNames.indexOf(name) === -1) {
								sourceNames.push(name);
							}
						});
						this.dataSources = sourceNames;
					})
			},
			goto: function (route) {
				this.$router.push(route);
			},
		}
	}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	table {
		border-collapse: collapse;
	}

	table, th, td {
		border: 1px solid black;
		margin: 5px;

	}

	th {
		text-align: center;
	}

	h3 {
		margin: 40px 0 0;
	}

	li {
		display: block;
		padding: .5em;
	}
</style>
