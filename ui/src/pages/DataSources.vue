<template>
	<div>
		<div style="display: flex; justify-content: space-between;">
			<div class="mdc-button icon-button">
				<div class="material-icons">share</div>
				Share
			</div>
			<div class="mdc-button icon-button">
				<div class="material-icons">cloud_download</div>
				Export
			</div>
			<div class="mdc-button icon-button">
				<div class="material-icons">backup</div>
				Backup
			</div>
			<div class="mdc-button icon-button" style="color: firebrick">
				<div class="material-icons">delete</div>
				Delete
			</div>
		</div>
		<div class="mdc-list">
			<div class="mdc-list-item" v-for="source in dataSources" v-on:click="goto('/data/' + source)">
				<span class="mdc-list-item__graphic material-icons" aria-hidden="true">folder</span>
				{{source}}
			</div>
		</div>
	</div>
</template>
<script>
	import testdata from '../testData/DataSources.json'

	export default {
		name: 'databoxStatus',
		props: {},
		data() {
			//get data from api later
			return {
				dataSources: [],
				timerID: 0,
			}
		},
		mounted() {
			this.$parent.title = "Databox Data";
			this.$parent.backRoute = "/settings";
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
				this.$parent.apiRequest('/core-ui/ui/api/dataSources', testdata)
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

	.icon-button {
		flex-direction: column;
		height: auto;
		padding: 8px;
	}

	.icon-button > .material-icons {
		padding-bottom: 8px;
	}
</style>
