<template>
  <div id="databoxStatus">
    <h1>Databox Data</h1>
    <p>
      Avalable data sources.
    </p>
    <table v-if="dataSources">
     <tr><th>Driver</th><th>Type</th><th>Id</th></tr>
     <icon v-for="(ds, index) in dataSources" :key="index" v-bind:hyperCat="ds" />
    </table>
  </div>
</template>
<script>
import testdata from '../testData/DataSources.json'
import Icon from '../components/renderDataSource.vue'

export default {
  name: 'databoxStatus',
  props: { },
  components: {
		Icon,
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
      this.ApiGetRequest('/core-ui/ui/api/dataSources',testdata)
      .then(json => {
        this.dataSources = json;
      })
    },
     GoToUI: function (appName) {
       alert("GoToUI " + appName)
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
