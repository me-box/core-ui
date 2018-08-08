<template>
      <tr><td>{{ driver }}</td>
      <td>{{ dsType }} </td>
      <td>  {{ datasourceid }}</td></tr>
</template>
<script>

export default {

  name: 'dataSourceIcon',
  props: {
    hyperCat: Object,
  },
  computed: {
      dsType: function () { return filterItemMetadata(this.hyperCat,"urn:X-databox:rels:hasType")},
      datasourceid: function () { return filterItemMetadata(this.hyperCat,"urn:X-databox:rels:hasDatasourceid") },
      vendor: function () { return filterItemMetadata(this.hyperCat,"urn:X-databox:rels:hasVendor") },
      driver: function () { return extractDriverName(this.hyperCat) },
  },
}

const extractDriverName =  (hyperCat) => {
        //the URL api will only pars http/s urs so replace tcp: with http:
        const url = new URL(hyperCat.href.replace("tcp:","http:"))
        return url.hostname.replace("-core-store","")
    }

const filterItemMetadata =  (hyperCat,match) => {
    return hyperCat["item-metadata"].filter((itm) =>
        {
            return itm.rel.includes(match)
        }
    )[0].val;
}

export {extractDriverName}
export {filterItemMetadata}

</script>

<!--  "scoped" CSS -->
<style scoped>

td {
    border: 1px solid black;
    margin: 5px;
    padding:5px;
}

</style>
