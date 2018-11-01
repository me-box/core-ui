<template>
	<div id="appStore">
		<h3>Available Apps</h3>
		<div v-if="apps" id="appList">
			<icon v-for="item in apps"
			      :key="item"
			      :name="item"
			      :displayName="true"
			      :route="'/install/' + item"
			      style="margin: 8px"/>
		</div>
	</div>
</template>
<script>
	import testdata from '../testData/apps.json'
	import Icon from '../components/AppIcon.vue'

	export default {
		name: 'AppStore',
		props: {
			msg: String,
		},
		components: {
			Icon,
		},
		data() {
			//get data from api later
			return {apps: [], drivers: [], timerID: 0}
		},
		mounted() {
			this.loadData();
			this.timerID = setInterval(() => {
				this.loadData();
			}, 5000);

			this.$parent.setTitle("Databox App Store");
		},
		destroyed: function () {
			clearInterval(this.timerID)
		},
		methods: {
			loadData: function () {
				this.$parent.apiRequest("/core-ui/ui/api/appStore", testdata)
					.then((data) => {
						let appList = data.apps;
						Array.prototype.push.apply(appList,data.drivers);
						appList.sort((a, b) => {
							a.localeCompare(b);
						});
						this.apps = appList;
					})
			}
		}
	}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	#appList {
		display: flex;
		flex-wrap: wrap;
	}

	#appStore {
		display: flex;
		margin: 24px;
	}

	@media (max-width: 479px) {
		#appStore {
			margin: 16px;
		}
	}
</style>
