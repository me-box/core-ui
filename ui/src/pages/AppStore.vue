<template>
	<div class="appStore">
		<div id="wrapper">
			<h3>Available Apps</h3>
			<div v-if="apps" style="display: flex; flex-wrap: wrap;">
				<icon v-for="item in apps"
				      :key="item"
				      :name="item"
				      :displayName="true"
				      :route="'/install/' + item"
				      style="margin: 8px"/>
				<icon v-for="item in drivers"
				      :key="item"
				      :name="item"
				      :displayName="true"
				      :route="'/install/' + item"
				      style="margin: 8px"/>
			</div>
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
				this.ApiGetRequest("/core-ui/ui/api/appStore", testdata)
					.then((data) => {
						this.apps = data.apps;
						this.drivers = data.drivers;
					})
			}
		}
	}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
