<template>
  <div>
  <select v-model="selected">
    <option v-if="filteredDataSources.length == 0" value="" disabled selected>No available data sources ...</option>
    <option v-else value="" selected>Select a data source</option>
    <option v-for="(ds, index) in filteredDataSources" v-bind:key="index" v-bind:value="ds">{{ds.description}}</option>
  </select>
  </div>
</template>
<script>

export default {
  name: 'datasourcesDDL',
  props: {
    dataSource: Object,
    avalableDataSources: Array,
    vlaue: Object,
  },
  computed: {
      dsID : function () {
          return this.dataSource.clientid
      },
      selected:{
      	get() {return this.value},
        set(v) {
                    this.$emit('selcted',v,this.dsID)
               }
      },
      filteredDataSources: function () {
          let getValFromHypercat = (hypercat,match) => {
              return hypercat["item-metadata"].filter((itm)=>{
                    return itm.rel == match
                })[0].val;
          }
          let filteredDs = []

          this.avalableDataSources.forEach( (ds) => {

            let dsType = getValFromHypercat(ds,"urn:X-databox:rels:hasType")
            if (dsType == this.dataSource.type) {
                filteredDs.push({
                    description: getValFromHypercat(ds,"urn:X-hypercat:rels:hasDescription:en"),
                    vendor: getValFromHypercat(ds,"urn:X-databox:rels:hasVendor"),
                    dsType: dsType,
                    href: ds["href"],
                    hypercat: ds
                })
            }
          })

        return filteredDs
      }
  }
}

</script>

<style scoped>
</style>
